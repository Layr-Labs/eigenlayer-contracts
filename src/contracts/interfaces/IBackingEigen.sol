// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IBackingEigen is IERC20 {
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
