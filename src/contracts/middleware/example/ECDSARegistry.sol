// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../RegistryBase.sol";

/**
 * @title A Registry-type contract identifying operators by their Ethereum address, with only 1 quorum.
 * @author Layr Labs, Inc.
 * @notice This contract is used for
 * - registering new operators
 * - committing to and finalizing de-registration as an operator
 * - updating the stakes of the operator
 */
contract ECDSARegistry is RegistryBase {

    /// @notice the address that can whitelist people
    address public operatorWhitelister;
    /// @notice toggle of whether the operator whitelist is on or off 
    bool public operatorWhitelistEnabled;
    /// @notice operator => are they whitelisted (can they register with the middleware)
    mapping(address => bool) public whitelisted;

    // EVENTS
    /**
     * @notice Emitted upon the registration of a new operator for the middleware
     * @param operator Address of the new operator
     * @param socket The ip:port of the operator
     */
    event Registration(
        address indexed operator,
        string socket
    );

    /// @notice Emitted when the `operatorWhitelister` role is transferred.
    event OperatorWhitelisterTransferred(address previousAddress, address newAddress);

    /// @notice Modifier that restricts a function to only be callable by the `whitelister` role.
    modifier onlyOperatorWhitelister {
        require(operatorWhitelister == msg.sender, "BLSRegistry.onlyOperatorWhitelister: not operatorWhitelister");
        _;
    }

    constructor(
        IStrategyManager _strategyManager,
        IServiceManager _serviceManager
    )
        RegistryBase(
            _strategyManager,
            _serviceManager
        )
    {}

    /// @notice Initialize whitelister and the quorum strategies + multipliers.
    function initialize(
        address _operatorWhitelister,
        bool _operatorWhitelistEnabled,
        uint256[] memory _quorumBips,
        uint96[] memory _minimumStakeForQuorums,
        StrategyAndWeightingMultiplier[][] memory _quorumStrategiesConsideredAndMultipliers
    ) public virtual initializer {
        _setOperatorWhitelister(_operatorWhitelister);
        operatorWhitelistEnabled = _operatorWhitelistEnabled;

        RegistryBase._initialize(
            _quorumBips,
            _minimumStakeForQuorums,
            _quorumStrategiesConsideredAndMultipliers
        );
    }

    /**
     * @notice Called by the service manager owner to transfer the whitelister role to another address 
     */
    function setOperatorWhitelister(address _operatorWhitelister) external onlyServiceManagerOwner {
        _setOperatorWhitelister(_operatorWhitelister);
    }

    /**
     * @notice Callable only by the service manager owner, this function toggles the whitelist on or off
     * @param _operatorWhitelistEnabled true if turning whitelist on, false otherwise
     */
    function setOperatorWhitelistStatus(bool _operatorWhitelistEnabled) external onlyServiceManagerOwner {
        operatorWhitelistEnabled = _operatorWhitelistEnabled;
    }

    /**
     * @notice Called by the operatorWhitelister, adds a list of operators to the whitelist
     * @param operators the operators to add to the whitelist
     */
    function addToOperatorWhitelist(address[] calldata operators) external onlyOperatorWhitelister {
        for (uint i = 0; i < operators.length; i++) {
            whitelisted[operators[i]] = true;
        }
    }

    /**
     * @notice Called by the operatorWhitelister, removes a list of operators to the whitelist
     * @param operators the operators to remove from the whitelist
     */
    function removeFromWhitelist(address[] calldata operators) external onlyOperatorWhitelister {
        for (uint i = 0; i < operators.length; i++) {
            whitelisted[operators[i]] = false;
        }
    }
    /**
     * @notice called for registering as an operator
     * @param quorumBitmap is the bitmap of quorums that the operator is registering for
     * @param socket is the socket address of the operator
     */
    function registerOperator(uint256 quorumBitmap, string calldata socket) external virtual {
        _registerOperator(msg.sender, quorumBitmap, socket);
    }

    /**
     * @param operator is the node who is registering to be a operator
     * @param socket is the socket address of the operator
     */
    function _registerOperator(address operator, uint256 quorumBitmap, string calldata socket)
        internal
    {
        if(operatorWhitelistEnabled) {
            require(whitelisted[operator], "BLSRegistry._registerOperator: not whitelisted");
        }

        // add the operator to the list of registrants and do accounting
        _addRegistrant(operator, bytes32(uint256(uint160(operator))), quorumBitmap);

        emit Registration(operator, socket);
    }

    /**
     * @notice Used by an operator to de-register itself from providing service to the middleware.
     * @param index is the sender's location in the dynamic array `operatorList`
     */
    function deregisterOperator(uint32 index) external virtual returns (bool) {
        _deregisterOperator(msg.sender, index);
        return true;
    }

    /**
     * @notice Used to process de-registering an operator from providing service to the middleware.
     * @param operator The operator to be deregistered
     * @param index is the sender's location in the dynamic array `operatorList`
     */
    function _deregisterOperator(address operator, uint32 index) internal {
        // verify that the `operator` is an active operator and that they've provided the correct `index`
        _deregistrationCheck(operator, index);

        // Perform necessary updates for removing operator, including updating operator list and index histories
        _removeOperator(operator, bytes32(uint256(uint160(operator))), index);
    }

    /**
     * @notice Used for updating information on deposits of nodes.
     * @param operators are the nodes whose deposit information is getting updated
     * @param prevElements are the elements before this middleware in the operator's linked list within the slasher
     */
    function updateStakes(address[] memory operators, uint256[] memory prevElements) external {
        // load all operator structs into memory
        Operator[] memory operatorStructs = new Operator[](operators.length);
        for (uint i = 0; i < operators.length;) {
            operatorStructs[i] = registry[operators[i]];
            unchecked {
                ++i;
            }
        }

        // load all operator quorum bitmaps into memory
        uint256[] memory quorumBitmaps = new uint256[](operators.length);
        for (uint i = 0; i < operators.length;) {
            quorumBitmaps[i] = pubkeyHashToQuorumBitmap[operatorStructs[i].pubkeyHash];
            unchecked {
                ++i;
            }
        }

        // for each quorum, loop through operators and see if they are apart of the quorum
        // if they are, get their new weight and update their individual stake history and the
        // quorum's total stake history accordingly
        for (uint8 quorumNumber = 0; quorumNumber < quorumCount;) {
            OperatorStakeUpdate memory totalStakeUpdate;
            // for each operator
            for(uint i = 0; i < operators.length;) {
                // if the operator is apart of the quorum
                if (quorumBitmaps[i] >> quorumNumber & 1 == 1) {
                    // if the total stake has not been loaded yet, load it
                    if (totalStakeUpdate.updateBlockNumber == 0) {
                        totalStakeUpdate = totalStakeHistory[quorumNumber][totalStakeHistory[quorumNumber].length - 1];
                    }
                    // get the operator's pubkeyHash
                    bytes32 pubkeyHash = operatorStructs[i].pubkeyHash;
                    // get the operator's current stake
                    OperatorStakeUpdate memory prevStakeUpdate = pubkeyHashToStakeHistory[pubkeyHash][quorumNumber][pubkeyHashToStakeHistory[pubkeyHash][quorumNumber].length - 1];
                    // update the operator's stake based on current state
                    OperatorStakeUpdate memory newStakeUpdate = _updateOperatorStake(operators[i], pubkeyHash, quorumBitmaps[i], quorumNumber, prevStakeUpdate.updateBlockNumber);
                    // calculate the new total stake for the quorum
                    totalStakeUpdate.stake = totalStakeUpdate.stake - prevStakeUpdate.stake + newStakeUpdate.stake;
                }
                unchecked {
                    ++i;
                }
            }
            // if the total stake for this quorum was updated, record it in storage
            if (totalStakeUpdate.updateBlockNumber != 0) {
                // update the total stake history for the quorum
                _recordTotalStakeUpdate(quorumNumber, totalStakeUpdate);
            }
            unchecked {
                ++quorumNumber;
            }
        }

        // record stake updates in the EigenLayer Slasher
        for (uint i = 0; i < operators.length;) {
            serviceManager.recordStakeUpdate(operators[i], uint32(block.number), serviceManager.latestServeUntilBlock(), prevElements[i]);
            unchecked {
                ++i;
            }
        }     
    }

    function _setOperatorWhitelister(address _operatorWhitelister) internal {
        require(_operatorWhitelister != address(0), "BLSRegistry.initialize: cannot set operatorWhitelister to zero address");
        emit OperatorWhitelisterTransferred(operatorWhitelister, _operatorWhitelister);
        operatorWhitelister = _operatorWhitelister;
    }
}
