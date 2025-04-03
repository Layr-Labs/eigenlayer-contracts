// SPDX-License-Identifier: BUSL-1.1
pragma solidity >=0.5.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IEigen is IERC20 {
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
