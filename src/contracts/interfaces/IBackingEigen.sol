// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IBackingEigen is IERC20 {
    /**
     * @notice This function allows the owner to set the allowedFrom status of an address
     * @param from the address whose allowedFrom status is being set
     * @param isAllowedFrom the new allowedFrom status
     */
    function setAllowedFrom(address from, bool isAllowedFrom) external;

    /**
     * @notice This function allows the owner to set the allowedTo status of an address
     * @param to the address whose allowedTo status is being set
     * @param isAllowedTo the new allowedTo status
     */
    function setAllowedTo(address to, bool isAllowedTo) external;

    /**
     * @notice Allows the owner to disable transfer restrictions
     */
    function disableTransferRestrictions() external;

    /**
     * @notice An initializer function that sets initial values for the contract's state variables.
     */
    function initialize(
        address initialOwner
    ) external;

    // @notice Allows the contract owner to modify an entry in the `isMinter` mapping.
    function setIsMinter(address minterAddress, bool newStatus) external;

    /**
     * @notice Allows any privileged address to mint `amount` new tokens to the address `to`.
     * @dev Callable only by an address that has `isMinter` set to true.
     */
    function mint(address to, uint256 amount) external;

    /**
     * @dev Destroys `amount` tokens from the caller.
     *
     * See {ERC20-_burn}.
     */
    function burn(
        uint256 amount
    ) external;

    /// @notice the address of the wrapped Eigen token EIGEN
    function EIGEN() external view returns (IERC20);

    /// @notice the timestamp after which transfer restrictions are disabled
    function transferRestrictionsDisabledAfter() external view returns (uint256);

    /**
     * @dev Clock used for flagging checkpoints. Has been overridden to implement timestamp based
     * checkpoints (and voting).
     */
    function clock() external view returns (uint48);

    /**
     * @dev Machine-readable description of the clock as specified in EIP-6372.
     * Has been overridden to inform callers that this contract uses timestamps instead of block numbers, to match `clock()`
     */
    // solhint-disable-next-line func-name-mixedcase
    function CLOCK_MODE() external pure returns (string memory);
}
