// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin-upgrades/contracts/token/ERC20/extensions/ERC20VotesUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract Eigen is OwnableUpgradeable, ERC20VotesUpgradeable {
    /// CONSTANTS & IMMUTABLES
    /// @notice the address of the backing Eigen token bEIGEN
    IERC20 public immutable bEIGEN;

    /// STORAGE
    /// @dev Do not remove, deprecated storage.
    /// @notice mapping of minter addresses to the timestamp after which they are allowed to mint
    mapping(address => uint256) internal __deprecated_mintAllowedAfter;
    /// @dev Do not remove, deprecated storage.
    /// @notice mapping of minter addresses to the amount of tokens they are allowed to mint
    mapping(address => uint256) internal __deprecated_mintingAllowance;
    /// @dev Do not remove, deprecated storage.
    /// @notice the timestamp after which transfer restrictions are disabled
    uint256 internal __deprecated_transferRestrictionsDisabledAfter;
    /// @dev Do not remove, deprecated storage.
    /// @notice mapping of addresses that are allowed to transfer tokens to any address
    mapping(address => bool) internal __deprecated_allowedFrom;
    /// @dev Do not remove, deprecated storage.
    /// @notice mapping of addresses that are allowed to receive tokens from any address
    mapping(address => bool) internal __deprecated_allowedTo;

    constructor(
        IERC20 _bEIGEN
    ) {
        bEIGEN = _bEIGEN;
        _disableInitializers();
    }

    /**
     * @notice An initializer function that sets initial values for the contract's state variables.
     */
    function initialize(
        address initialOwner
    ) public initializer {
        __Ownable_init();
        __ERC20_init("Eigen", "EIGEN");
        _transferOwnership(initialOwner);
        __ERC20Permit_init("EIGEN");
    }

    /**
     * @notice This function allows bEIGEN holders to wrap their tokens into Eigen
     */
    function wrap(
        uint256 amount
    ) external {
        require(bEIGEN.transferFrom(msg.sender, address(this), amount), "Eigen.wrap: bEIGEN transfer failed");
        _mint(msg.sender, amount);
    }

    /**
     * @notice This function allows Eigen holders to unwrap their tokens into bEIGEN
     */
    function unwrap(
        uint256 amount
    ) external {
        _burn(msg.sender, amount);
        require(bEIGEN.transfer(msg.sender, amount), "Eigen.unwrap: bEIGEN transfer failed");
    }

    /**
     * @notice Allows the sender to transfer tokens to multiple addresses in a single transaction
     */
    function multisend(address[] calldata receivers, uint256[] calldata amounts) public {
        require(receivers.length == amounts.length, "Eigen.multisend: receivers and amounts must be the same length");
        for (uint256 i = 0; i < receivers.length; i++) {
            _transfer(msg.sender, receivers[i], amounts[i]);
        }
    }

    /**
     * @notice Overridden to return the total bEIGEN supply instead.
     * @dev The issued supply of EIGEN should match the bEIGEN balance of this contract,
     * less any bEIGEN tokens that were sent directly to the contract (rather than being wrapped)
     */
    function totalSupply() public view override returns (uint256) {
        return bEIGEN.totalSupply();
    }

    /**
     * @dev Clock used for flagging checkpoints. Has been overridden to implement timestamp based
     * checkpoints (and voting).
     */
    function clock() public view override returns (uint48) {
        return SafeCastUpgradeable.toUint48(block.timestamp);
    }

    /**
     * @dev Machine-readable description of the clock as specified in EIP-6372.
     * Has been overridden to inform callers that this contract uses timestamps instead of block numbers, to match `clock()`
     */
    // solhint-disable-next-line func-name-mixedcase
    function CLOCK_MODE() public pure override returns (string memory) {
        return "mode=timestamp";
    }
}
