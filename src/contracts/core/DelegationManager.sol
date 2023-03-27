// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/interfaces/IERC1271.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "./DelegationManagerStorage.sol";
import "../permissions/Pausable.sol";
import "./Slasher.sol";

/**
 * @title The primary delegation contract for EigenLayer.
 * @author Layr Labs, Inc.
 * @notice  This is the contract for delegation in EigenLayer. The main functionalities of this contract are
 * - enabling anyone to register as an operator in EigenLayer
 * - allowing new operators to provide a DelegationTerms-type contract, which may mediate their interactions with stakers who delegate to them
 * - enabling any staker to delegate its stake to the operator of its choice
 * - enabling a staker to undelegate its assets from an operator (performed as part of the withdrawal process, initiated through the StrategyManager)
 */
contract DelegationManager is Initializable, OwnableUpgradeable, Pausable, DelegationManagerStorage {
    // index for flag that pauses new delegations when set
    uint8 internal constant PAUSED_NEW_DELEGATION = 0;
    // bytes4(keccak256("isValidSignature(bytes32,bytes)")
    bytes4 constant internal ERC1271_MAGICVALUE = 0x1626ba7e;

    /// @notice Simple permission for functions that are only callable by the StrategyManager contract.
    modifier onlyStrategyManager() {
        require(msg.sender == address(strategyManager), "onlyStrategyManager");
        _;
    }

    // INITIALIZING FUNCTIONS
    constructor(IStrategyManager _strategyManager, ISlasher _slasher) 
        DelegationManagerStorage(_strategyManager, _slasher)
    {
        _disableInitializers();
    }

    /// @dev Emitted when a low-level call to `delegationTerms.onDelegationReceived` fails, returning `returnData`
    event OnDelegationReceivedCallFailure(IDelegationTerms indexed delegationTerms, bytes32 returnData);

    /// @dev Emitted when a low-level call to `delegationTerms.onDelegationWithdrawn` fails, returning `returnData`
    event OnDelegationWithdrawnCallFailure(IDelegationTerms indexed delegationTerms, bytes32 returnData);

    function initialize(address initialOwner, IPauserRegistry _pauserRegistry, uint256 initialPausedStatus)
        external
        initializer
    {
        _initializePauser(_pauserRegistry, initialPausedStatus);
        DOMAIN_SEPARATOR = keccak256(abi.encode(DOMAIN_TYPEHASH, bytes("EigenLayer"), block.chainid, address(this)));
        _transferOwnership(initialOwner);
    }

    // EXTERNAL FUNCTIONS
    /**
     * @notice This will be called by an operator to register itself as an operator that stakers can choose to delegate to.
     * @param dt is the `DelegationTerms` contract that the operator has for those who delegate to them.
     * @dev An operator can set `dt` equal to their own address (or another EOA address), in the event that they want to split payments
     * in a more 'trustful' manner.
     * @dev In the present design, once set, there is no way for an operator to ever modify the address of their DelegationTerms contract.
     */
    function registerAsOperator(IDelegationTerms dt) external {
        require(
            address(delegationTerms[msg.sender]) == address(0),
            "DelegationManager.registerAsOperator: operator has already registered"
        );
        // store the address of the delegation contract that the operator is providing.
        delegationTerms[msg.sender] = dt;
        _delegate(msg.sender, msg.sender);
    }

    /**
     *  @notice This will be called by a staker to delegate its assets to some operator.
     *  @param operator is the operator to whom staker (msg.sender) is delegating its assets
     */
    function delegateTo(address operator) external {
        _delegate(msg.sender, operator);
    }

    /**
     * @notice Delegates from `staker` to `operator`.
     * @dev requires that:
     * 1) if `staker` is an EOA, then `signature` is valid ECSDA signature from `staker`, indicating their intention for this action
     * 2) if `staker` is a contract, then `signature` must will be checked according to EIP-1271
     */
    function delegateToBySignature(address staker, address operator, uint256 expiry, bytes memory signature)
        external
    {
        require(expiry >= block.timestamp, "DelegationManager.delegateToBySignature: delegation signature expired");

        // calculate struct hash, then increment `staker`'s nonce
        uint256 nonce = nonces[staker];
        bytes32 structHash = keccak256(abi.encode(DELEGATION_TYPEHASH, staker, operator, nonce, expiry));
        unchecked {
            nonces[staker] = nonce + 1;
        }
        bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", DOMAIN_SEPARATOR, structHash));

        /**
         * check validity of signature:
         * 1) if `staker` is an EOA, then `signature` must be a valid ECSDA signature from `staker`,
         * indicating their intention for this action
         * 2) if `staker` is a contract, then `signature` must will be checked according to EIP-1271
         */
        if (Address.isContract(staker)) {
            require(IERC1271(staker).isValidSignature(digestHash, signature) == ERC1271_MAGICVALUE,
                "DelegationManager.delegateToBySignature: ERC1271 signature verification failed");
        } else {
            require(ECDSA.recover(digestHash, signature) == staker,
                "DelegationManager.delegateToBySignature: sig not from staker");
        }

        _delegate(staker, operator);
    }

    /**
     * @notice Undelegates `staker` from the operator who they are delegated to.
     * @notice Callable only by the StrategyManager
     * @dev Should only ever be called in the event that the `staker` has no active deposits in EigenLayer.
     */
    function undelegate(address staker) external onlyStrategyManager {
        require(!isOperator(staker), "DelegationManager.undelegate: operators cannot undelegate from themselves");
        delegatedTo[staker] = address(0);
    }

    /**
     * @notice Increases the `staker`'s delegated shares in `strategy` by `shares, typically called when the staker has further deposits into EigenLayer
     * @dev Callable only by the StrategyManager
     */
    function increaseDelegatedShares(address staker, IStrategy strategy, uint256 shares)
        external
        onlyStrategyManager
    {
        //if the staker is delegated to an operator
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];

            // add strategy shares to delegate's shares
            operatorShares[operator][strategy] += shares;

            //Calls into operator's delegationTerms contract to update weights of individual staker
            IStrategy[] memory stakerStrategyList = new IStrategy[](1);
            uint256[] memory stakerShares = new uint[](1);
            stakerStrategyList[0] = strategy;
            stakerShares[0] = shares;

            // call into hook in delegationTerms contract
            IDelegationTerms dt = delegationTerms[operator];
            _delegationReceivedHook(dt, staker, stakerStrategyList, stakerShares);
        }
    }

    /**
     * @notice Decreases the `staker`'s delegated shares in each entry of `strategies` by its respective `shares[i]`, typically called when the staker withdraws from EigenLayer
     * @dev Callable only by the StrategyManager
     */
    function decreaseDelegatedShares(
        address staker,
        IStrategy[] calldata strategies,
        uint256[] calldata shares
    )
        external
        onlyStrategyManager
    {
        if (isDelegated(staker)) {
            address operator = delegatedTo[staker];

            // subtract strategy shares from delegate's shares
            uint256 stratsLength = strategies.length;
            for (uint256 i = 0; i < stratsLength;) {
                operatorShares[operator][strategies[i]] -= shares[i];
                unchecked {
                    ++i;
                }
            }

            // call into hook in delegationTerms contract
            IDelegationTerms dt = delegationTerms[operator];
            _delegationWithdrawnHook(dt, staker, strategies, shares);
        }
    }

    // INTERNAL FUNCTIONS

    /** 
     * @notice Makes a low-level call to `dt.onDelegationReceived(staker, strategies, shares)`, ignoring reverts and with a gas budget 
     * equal to `LOW_LEVEL_GAS_BUDGET` (a constant defined in this contract).
     * @dev *If* the low-level call fails, then this function emits the event `OnDelegationReceivedCallFailure(dt, returnData)`, where
     * `returnData` is *only the first 32 bytes* returned by the call to `dt`.
     */
    function _delegationReceivedHook(
        IDelegationTerms dt,
        address staker,
        IStrategy[] memory strategies,
        uint256[] memory shares
    )
        internal
    {
        /**
         * We use low-level call functionality here to ensure that an operator cannot maliciously make this function fail in order to prevent undelegation.
         * In particular, in-line assembly is also used to prevent the copying of uncapped return data which is also a potential DoS vector.
         */
        // format calldata
        bytes memory lowLevelCalldata = abi.encodeWithSelector(IDelegationTerms.onDelegationReceived.selector, staker, strategies, shares);
        // Prepare memory for low-level call return data. We accept a max return data length of 32 bytes
        bool success;
        bytes32[1] memory returnData;
        // actually make the call
        assembly {
            success := call(
                // gas provided to this context
                LOW_LEVEL_GAS_BUDGET,
                // address to call
                dt,
                // value in wei for call
                0,
                // memory location to copy for calldata
                add(lowLevelCalldata, 32),
                // length of memory to copy for calldata
                mload(lowLevelCalldata),
                // memory location to copy return data
                returnData,
                // byte size of return data to copy to memory
                32
            )
        }
        // if the call fails, we emit a special event rather than reverting
        if (!success) {
            emit OnDelegationReceivedCallFailure(dt, returnData[0]);
        }
    }

    /** 
     * @notice Makes a low-level call to `dt.onDelegationWithdrawn(staker, strategies, shares)`, ignoring reverts and with a gas budget 
     * equal to `LOW_LEVEL_GAS_BUDGET` (a constant defined in this contract).
     * @dev *If* the low-level call fails, then this function emits the event `OnDelegationReceivedCallFailure(dt, returnData)`, where
     * `returnData` is *only the first 32 bytes* returned by the call to `dt`.
     */
    function _delegationWithdrawnHook(
        IDelegationTerms dt,
        address staker,
        IStrategy[] memory strategies,
        uint256[] memory shares
    )
        internal
    {
        /**
         * We use low-level call functionality here to ensure that an operator cannot maliciously make this function fail in order to prevent undelegation.
         * In particular, in-line assembly is also used to prevent the copying of uncapped return data which is also a potential DoS vector.
         */
        // format calldata
        bytes memory lowLevelCalldata = abi.encodeWithSelector(IDelegationTerms.onDelegationWithdrawn.selector, staker, strategies, shares);
        // Prepare memory for low-level call return data. We accept a max return data length of 32 bytes
        bool success;
        bytes32[1] memory returnData;
        // actually make the call
        assembly {
            success := call(
                // gas provided to this context
                LOW_LEVEL_GAS_BUDGET,
                // address to call
                dt,
                // value in wei for call
                0,
                // memory location to copy for calldata
                add(lowLevelCalldata, 32),
                // length of memory to copy for calldata
                mload(lowLevelCalldata),
                // memory location to copy return data
                returnData,
                // byte size of return data to copy to memory
                32
            )
        }
        // if the call fails, we emit a special event rather than reverting
        if (!success) {
            emit OnDelegationWithdrawnCallFailure(dt, returnData[0]);
        }
    }

    /**
     * @notice Internal function implementing the delegation *from* `staker` *to* `operator`.
     * @param staker The address to delegate *from* -- this address is delegating control of its own assets.
     * @param operator The address to delegate *to* -- this address is being given power to place the `staker`'s assets at risk on services
     * @dev Ensures that the operator has registered as a delegate (`address(dt) != address(0)`), verifies that `staker` is not already
     * delegated, and records the new delegation.
     */ 
    function _delegate(address staker, address operator) internal onlyWhenNotPaused(PAUSED_NEW_DELEGATION) {
        IDelegationTerms dt = delegationTerms[operator];
        require(
            address(dt) != address(0), "DelegationManager._delegate: operator has not yet registered as a delegate"
        );

        require(isNotDelegated(staker), "DelegationManager._delegate: staker has existing delegation");
        // checks that operator has not been frozen
        require(!slasher.isFrozen(operator), "DelegationManager._delegate: cannot delegate to a frozen operator");

        // record delegation relation between the staker and operator
        delegatedTo[staker] = operator;

        // retrieve list of strategies and their shares from strategy manager
        (IStrategy[] memory strategies, uint256[] memory shares) = strategyManager.getDeposits(staker);

        // add strategy shares to delegate's shares
        uint256 stratsLength = strategies.length;
        for (uint256 i = 0; i < stratsLength;) {
            // update the share amounts for each of the operator's strategies
            operatorShares[operator][strategies[i]] += shares[i];
            unchecked {
                ++i;
            }
        }

        // call into hook in delegationTerms contract
        _delegationReceivedHook(dt, staker, strategies, shares);
    }

    // VIEW FUNCTIONS

    /// @notice Returns 'true' if `staker` *is* actively delegated, and 'false' otherwise.
    function isDelegated(address staker) public view returns (bool) {
        return (delegatedTo[staker] != address(0));
    }

    /// @notice Returns 'true' if `staker` is *not* actively delegated, and 'false' otherwise.
    function isNotDelegated(address staker) public view returns (bool) {
        return (delegatedTo[staker] == address(0));
    }

    /// @notice Returns if an operator can be delegated to, i.e. it has called `registerAsOperator`.
    function isOperator(address operator) public view returns (bool) {
        return (address(delegationTerms[operator]) != address(0));
    }
}
