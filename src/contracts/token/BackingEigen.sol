// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin-upgrades/contracts/token/ERC20/extensions/ERC20VotesUpgradeable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";

contract BackingEigen is OwnableUpgradeable, ERC20VotesUpgradeable {
    /// CONSTANTS & IMMUTABLES
    /// @notice the address of the wrapped Eigen token EIGEN
    IERC20 public immutable EIGEN;

    /// STORAGE
    /// @dev Do not remove, deprecated storage.
    /// @notice the timestamp after which transfer restrictions are disabled
    uint256 internal __dreprecated_transferRestrictionsDisabledAfter;
    /// @dev Do not remove, deprecated storage.
    /// @notice mapping of addresses that are allowed to transfer tokens to any address
    mapping(address => bool) internal __dreprecated_allowedFrom;
    /// @dev Do not remove, deprecated storage.
    /// @notice mapping of addresses that are allowed to receive tokens from any address
    mapping(address => bool) internal __dreprecated_allowedTo;

    // @notice whether or not an address is allowed to mint new bEIGEN tokens
    mapping(address => bool) public isMinter;

    event Backed();
    // @notice event emitted when the `isMinter` mapping is modified
    event IsMinterModified(address indexed minterAddress, bool newStatus);

    constructor(
        IERC20 _EIGEN
    ) {
        EIGEN = _EIGEN;
        _disableInitializers();
    }

    // @notice Allows the contract owner to modify an entry in the `isMinter` mapping.
    function setIsMinter(address minterAddress, bool newStatus) external onlyOwner {
        emit IsMinterModified(minterAddress, newStatus);
        isMinter[minterAddress] = newStatus;
    }

    /**
     * @notice Allows any privileged address to mint `amount` new tokens to the address `to`.
     * @dev Callable only by an address that has `isMinter` set to true.
     */
    function mint(address to, uint256 amount) external {
        require(isMinter[msg.sender], "BackingEigen.mint: caller is not a minter");
        _mint(to, amount);
    }

    /**
     * @dev Destroys `amount` tokens from the caller.
     *
     * See {ERC20-_burn}.
     */
    function burn(
        uint256 amount
    ) public virtual {
        _burn(_msgSender(), amount);
    }

    /**
     * @notice An initializer function that sets initial values for the contract's state variables.
     */
    function initialize(
        address initialOwner
    ) public initializer {
        __Ownable_init();
        __ERC20_init("Backing Eigen", "bEIGEN");
        _transferOwnership(initialOwner);
        __ERC20Permit_init("bEIGEN");

        // Mint the entire supply of EIGEN - this is a one-time event that
        // ensures bEIGEN fully backs EIGEN.
        _mint(address(EIGEN), 1_673_646_668_284_660_000_000_000_000);
        emit Backed();
    }

    /// VIEW FUNCTIONS

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
