// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/security/ReentrancyGuardUpgradeable.sol";
import "../permissions/Pausable.sol";
import "../libraries/EIP1271SignatureUtils.sol";
import "./AVSDirectoryStorage.sol";

contract AVSDirectory is
    Initializable,
    OwnableUpgradeable,
    Pausable,
    AVSDirectoryStorage,
    ReentrancyGuardUpgradeable
{
    // @dev Index for flag that pauses operator register/deregister to avs when set.
    uint8 internal constant PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS = 0;

    // @dev Chain ID at the time of contract deployment
    uint256 internal immutable ORIGINAL_CHAIN_ID;

    /// @notice Canonical, virtual beacon chain ETH strategy
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    /**
     *
     *                         INITIALIZING FUNCTIONS
     *
     */

    /**
     * @dev Initializes the immutable addresses of the strategy mananger, delegationManager, slasher,
     * and eigenpodManager contracts
     */
    constructor(
        IDelegationManager _delegation,
        IStrategyManager _strategyManager
    ) AVSDirectoryStorage(_delegation, _strategyManager) {
        _disableInitializers();
        ORIGINAL_CHAIN_ID = block.chainid;
    }

    /**
     * @dev Initializes the addresses of the initial owner, pauser registry, and paused status.
     * minWithdrawalDelayBlocks is set only once here
     */
    function initialize(
        address initialOwner,
        IPauserRegistry _pauserRegistry,
        uint256 initialPausedStatus
    ) external initializer {
        _initializePauser(_pauserRegistry, initialPausedStatus);
        _DOMAIN_SEPARATOR = _calculateDomainSeparator();
        _transferOwnership(initialOwner);
    }

    /**
     *
     *                         EXTERNAL FUNCTIONS
     *
     */

    /**
     * @notice Called by AVSs to add an operator to an operator set
     *
     * @param operator the address of the operator to be added to the operator set
     * @param operatorSetID the ID of the operator set
     * @param operatorSignature the signature of the operator on their intent to register
     *
     * @dev msg.sender is used as the AVS
     * @dev operator must not have a deregistration from the operator set
     * @dev if this is the first operator set in the AVS that the operator is
     * registering for, a OperatorAVSRegistrationStatusUpdated event is emitted with
     * a REGISTERED status
     * =
     */
    function registerOperatorToOperatorSet(
        address operator,
        uint32 operatorSetID,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        require(
            operatorSignature.expiry >= block.timestamp,
            "AVSDirectory.registerOperatorToOperatorSet: operator signature expired"
        );
        require(
            operatorSetRegistrations[msg.sender][operator][operatorSetID] != true,
            "AVSDirectory.registerOperatorToOperatorSet: operator already registered to operator set"
        );
        require(
            !operatorSaltIsSpent[operator][operatorSignature.salt],
            "AVSDirectory.registerOperatorToAVS: salt already spent"
        );
        require(
            delegation.isOperator(operator),
            "AVSDirectory.registerOperatorToAVS: operator not registered to EigenLayer yet"
        );

        // Calculate the digest hash
        bytes32 operatorSetRegistrationDigestHash = calculateOperatorAVSRegistrationDigestHash({
            operator: operator,
            avs: msg.sender,
            salt: operatorSignature.salt,
            expiry: operatorSignature.expiry
        });

        // Check that the signature is valid
        EIP1271SignatureUtils.checkSignature_EIP1271(
            operator, operatorSetRegistrationDigestHash, operatorSignature.signature
        );

        // Set as Operator Set AVS, preventing any further legacy registrations
        if (!isOperatorSetAVS[msg.sender]) {
            isOperatorSetAVS[msg.sender] = true;
        }

        // Update operator set registration
        operatorSetRegistrations[msg.sender][operator][operatorSetID] = true;
        operatorAVSOperatorSetCount[msg.sender][operator] += 1;

        // Mark the salt as spent
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;

        // Set the operator as registered if not already
        if (avsOperatorStatus[msg.sender][operator] != OperatorAVSRegistrationStatus.REGISTERED) {
            avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.REGISTERED;
            emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.REGISTERED);
        }

        emit OperatorAddedToOperatorSet(operator, OperatorSet(msg.sender, operatorSetID));
    }

    /**
     * @notice Called by AVSs or operators to remove an operator to from operator set
     *
     * @param operator the address of the operator to be removed from the
     * operator set
     * @param operatorSetID the ID of the operator set
     *
     * @dev msg.sender is used as the AVS
     * @dev operator must be registered for msg.sender AVS and the given
     * operator set
     */
    function deregisterOperatorFromOperatorSet(
        address operator,
        uint32 operatorSetID
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        require(
            operatorSetRegistrations[msg.sender][operator][operatorSetID] == true,
            "AVSDirectory.deregisterOperatorFromOperatorSet: operator not registered for operator set"
        );

        // Update operator set registration
        operatorSetRegistrations[msg.sender][operator][operatorSetID] = false;
        operatorAVSOperatorSetCount[msg.sender][operator] -= 1;

        emit OperatorRemovedFromOperatorSet(operator, OperatorSet(msg.sender, operatorSetID));

        // Set the operator as deregistered if no longer registered for any operator sets
        if (operatorAVSOperatorSetCount[msg.sender][operator] == 0) {
            avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.UNREGISTERED;
            emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED);
        }
    }

    /**
     * @notice Called by the AVS's service manager contract to register an operator with the avs.
     * @param operator The address of the operator to register.
     * @param operatorSignature The signature, salt, and expiry of the operator's signature.
     */
    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS) {
        require(
            operatorSignature.expiry >= block.timestamp,
            "AVSDirectory.registerOperatorToAVS: operator signature expired"
        );
        require(
            avsOperatorStatus[msg.sender][operator] != OperatorAVSRegistrationStatus.REGISTERED,
            "AVSDirectory.registerOperatorToAVS: operator already registered"
        );
        require(
            !operatorSaltIsSpent[operator][operatorSignature.salt],
            "AVSDirectory.registerOperatorToAVS: salt already spent"
        );
        require(
            delegation.isOperator(operator),
            "AVSDirectory.registerOperatorToAVS: operator not registered to EigenLayer yet"
        );
        require(!isOperatorSetAVS[msg.sender]);

        // Calculate the digest hash
        bytes32 operatorRegistrationDigestHash = calculateOperatorAVSRegistrationDigestHash({
            operator: operator,
            avs: msg.sender,
            salt: operatorSignature.salt,
            expiry: operatorSignature.expiry
        });

        // Check that the signature is valid
        EIP1271SignatureUtils.checkSignature_EIP1271(
            operator, operatorRegistrationDigestHash, operatorSignature.signature
        );

        // Set the operator as registered
        avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.REGISTERED;

        // Mark the salt as spent
        operatorSaltIsSpent[operator][operatorSignature.salt] = true;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.REGISTERED);
    }

    /**
     * @notice Called by an avs to deregister an operator with the avs.
     * @param operator The address of the operator to deregister.
     */
    function deregisterOperatorFromAVS(address operator)
        external
        onlyWhenNotPaused(PAUSED_OPERATOR_REGISTER_DEREGISTER_TO_AVS)
    {
        require(
            avsOperatorStatus[msg.sender][operator] == OperatorAVSRegistrationStatus.REGISTERED,
            "AVSDirectory.deregisterOperatorFromAVS: operator not registered"
        );

        // Set the operator as deregistered
        avsOperatorStatus[msg.sender][operator] = OperatorAVSRegistrationStatus.UNREGISTERED;

        emit OperatorAVSRegistrationStatusUpdated(operator, msg.sender, OperatorAVSRegistrationStatus.UNREGISTERED);
    }

    /**
     * @notice Called by an avs to emit an `AVSMetadataURIUpdated` event indicating the information has updated.
     * @param metadataURI The URI for metadata associated with an avs
     */
    function updateAVSMetadataURI(string calldata metadataURI) external {
        emit AVSMetadataURIUpdated(msg.sender, metadataURI);
    }

    /**
     * @notice Adds strategies to an operator set
     * @param operatorSetID The ID of the operator set
     * @param strategies The strategies to add to the operator set
     * TODO: align with team on if we want to keep the two require statements
     */
    function addStrategiesToOperatorSet(uint32 operatorSetID, IStrategy[] calldata strategies) external {
        uint256 strategiesToAdd = strategies.length;
        for (uint256 i = 0; i < strategiesToAdd; i++) {
            // Require that the strategy is valid
            require(
                strategyManager.strategyIsWhitelistedForDeposit(strategies[i])
                    || strategies[i] == beaconChainETHStrategy,
                "AVSDirectory.addStrategiesToOperatorSet: invalid strategy considered"
            );
            require(
                !operatorSetStrategies[msg.sender][operatorSetID][strategies[i]],
                "AVSDirectory.addStrategiesToOperatorSet: strategy already added to operator set"
            );
            operatorSetStrategies[msg.sender][operatorSetID][strategies[i]] = true;
            emit OperatorSetStrategyAdded(OperatorSet(msg.sender, operatorSetID), strategies[i]);
        }
    }

    /**
     * @notice Removes strategies from an operator set
     * @param operatorSetID The ID of the operator set
     * @param strategies The strategies to remove from the operator set
     */
    function removeStrategiesFromOperatorSet(uint32 operatorSetID, IStrategy[] calldata strategies) external {
        uint256 strategiesToRemove = strategies.length;
        for (uint256 i = 0; i < strategiesToRemove; i++) {
            require(
                operatorSetStrategies[msg.sender][operatorSetID][strategies[i]],
                "AVSDirectory.removeStrategiesFromOperatorSet: strategy not a member of operator set"
            );
            operatorSetStrategies[msg.sender][operatorSetID][strategies[i]] = false;
            emit OperatorSetStrategyRemoved(OperatorSet(msg.sender, operatorSetID), strategies[i]);
        }
    }

    /**
     * @notice Called by an operator to cancel a salt that has been used to register with an AVS.
     * @param salt A unique and single use value associated with the approver signature.
     */
    function cancelSalt(bytes32 salt) external {
        require(!operatorSaltIsSpent[msg.sender][salt], "AVSDirectory.cancelSalt: cannot cancel spent salt");
        operatorSaltIsSpent[msg.sender][salt] = true;
    }

    /**
     *
     *                         VIEW FUNCTIONS
     *
     */

    /**
     * @notice Calculates the digest hash to be signed by an operator to register with an AVS
     * @param operator The account registering as an operator
     * @param avs The address of the service manager contract for the AVS that the operator is registering to
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid
     */
    function calculateOperatorAVSRegistrationDigestHash(
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        // calculate the struct hash
        bytes32 structHash = keccak256(abi.encode(OPERATOR_AVS_REGISTRATION_TYPEHASH, operator, avs, salt, expiry));
        // calculate the digest hash
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", domainSeparator(), structHash));
        return digestHash;
    }

    /**
     * @notice Calculates the digest hash to be signed by an operator to register with an operator set
     * @param operator The operator set that the operator is registering to
     * @param operatorSet A struct containing info about a given operator set. 
     * @param salt A unique and single use value associated with the approver signature.
     * @param expiry Time after which the approver's signature becomes invalid
     */
    function calculateOperatorSetRegistrationDigestHash(
        address operator,
        OperatorSet memory operatorSet,
        bytes32 salt,
        uint256 expiry
    ) public view returns (bytes32) {
        // calculate the struct hash
        bytes32 structHash = keccak256(
            abi.encode(OPERATOR_SET_REGISTRATION_TYPEHASH, operator, operatorSet.avs, operatorSet.id, salt, expiry)
        );
        // calculate the digest hash
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", domainSeparator(), structHash));
        return digestHash;
    }

    /**
     * @notice Getter function for the current EIP-712 domain separator for this contract.
     * @dev The domain separator will change in the event of a fork that changes the ChainID.
     */
    function domainSeparator() public view returns (bytes32) {
        if (block.chainid == ORIGINAL_CHAIN_ID) {
            return _DOMAIN_SEPARATOR;
        } else {
            return _calculateDomainSeparator();
        }
    }

    // @notice Internal function for calculating the current domain separator of this contract
    function _calculateDomainSeparator() internal view returns (bytes32) {
        return keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes("EigenLayer")), block.chainid, address(this)));
    }
}
