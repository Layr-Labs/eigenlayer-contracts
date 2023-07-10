// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./IServiceManager.sol";

/**
 * @title Draft interface for a contract that helps structure the delegation relationship between operators and stakers. May be merged with another interface.
 * @author Layr Labs, Inc.
 */
interface IDelegationDetails {
    // @notice Struct used for storing a single operator's fee for a single service
    struct OperatorFee {
        // operator fee in basis points
        uint16 feeBips;
        // boolean flag used for tracking whether or not the operator has ever called `setOperatorFeeForService` for the specific service in question 
        bool feeSet;
    }

    // @notice Event emitted when an `operator` modifies their fee for a service specified by `serviceManager`, from `previousFeeBips` to `newFeeBips`
    event OperatorFeeModifiedForService(address indexed operator, IServiceManager indexed serviceManager, uint16 previousFeeBips, uint16 newFeeBips);

    // @notice Mapping operator => ServiceManager => operator fee in basis points + whether or not they have set their fee for the service
    // mapping(address => IServiceManager => OperatorFee) public operatorFeeBipsByService;
    function operatorFeeBipsByService(address operator, IServiceManager serviceManager) external view returns (OperatorFee memory);

    // uint256 public constant MAX_BIPS = 10000;
    function MAX_BIPS() external view returns (uint256);

    // uint256 public constant MAX_OPERATOR_FEE_BIPS = 1500;
    function MAX_OPERATOR_FEE_BIPS() external view returns (uint256);

    // @notice Called by an operator to set their fee bips for the specified service. Can only be called once by each operator, for each `serviceManager`
    function setOperatorFeeForService(uint16 feeBips, IServiceManager serviceManager) external;

    // @notice Called by an operator to reduce their fee bips for the specified service.
    function decreaseOperatorFeeForService(uint16 newFeeBips,  IServiceManager serviceManager) external;
}