// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IEigen is IERC20 {
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
     * @notice This function allows minter to mint tokens
     */
    function mint() external;

    /**
     * @notice This function allows bEIGEN holders to wrap their tokens into Eigen
     */
    function wrap(
        uint256 amount
    ) external;

    /**
     * @notice This function allows Eigen holders to unwrap their tokens into bEIGEN
     */
    function unwrap(
        uint256 amount
    ) external;

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
