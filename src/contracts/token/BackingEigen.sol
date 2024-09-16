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
    /// @notice the timestamp after which transfer restrictions are disabled
    uint256 public transferRestrictionsDisabledAfter;
    /// @notice mapping of addresses that are allowed to transfer tokens to any address
    mapping(address => bool) public allowedFrom;
    /// @notice mapping of addresses that are allowed to receive tokens from any address
    mapping(address => bool) public allowedTo;
    // @notice whether or not an address is allowed to mint new bEIGEN tokens
    mapping(address => bool) public isMinter;

    /// @notice event emitted when the allowedFrom status of an address is set
    event SetAllowedFrom(address indexed from, bool isAllowedFrom);
    /// @notice event emitted when the allowedTo status of an address is set
    event SetAllowedTo(address indexed to, bool isAllowedTo);
    /// @notice event emitted when the transfer restrictions are disabled
    event TransferRestrictionsDisabled();
    /// @notice event emitted when the EIGEN token is backed
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

        // set transfer restrictions to be disabled at type(uint256).max to be set down later
        transferRestrictionsDisabledAfter = type(uint256).max;

        // the EIGEN contract should be allowed to transfer tokens to any address for unwrapping
        // likewise, anyone should be able to transfer bEIGEN to EIGEN for wrapping
        _setAllowedFrom(address(EIGEN), true);
        _setAllowedTo(address(EIGEN), true);

        // Mint the entire supply of EIGEN - this is a one-time event that
        // ensures bEIGEN fully backs EIGEN.
        _mint(address(EIGEN), 1_673_646_668_284_660_000_000_000_000);
        emit Backed();
    }

    /// EXTERNAL FUNCTIONS

    /**
     * @notice This function allows the owner to set the allowedFrom status of an address
     * @param from the address whose allowedFrom status is being set
     * @param isAllowedFrom the new allowedFrom status
     */
    function setAllowedFrom(address from, bool isAllowedFrom) external onlyOwner {
        _setAllowedFrom(from, isAllowedFrom);
    }

    /**
     * @notice This function allows the owner to set the allowedTo status of an address
     * @param to the address whose allowedTo status is being set
     * @param isAllowedTo the new allowedTo status
     */
    function setAllowedTo(address to, bool isAllowedTo) external onlyOwner {
        _setAllowedTo(to, isAllowedTo);
    }

    /**
     * @notice Allows the owner to disable transfer restrictions
     */
    function disableTransferRestrictions() external onlyOwner {
        require(
            transferRestrictionsDisabledAfter == type(uint256).max,
            "BackingEigen.disableTransferRestrictions: transfer restrictions are already disabled"
        );
        transferRestrictionsDisabledAfter = 0;
        emit TransferRestrictionsDisabled();
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

    /// INTERNAL FUNCTIONS

    function _setAllowedFrom(address from, bool isAllowedFrom) internal {
        allowedFrom[from] = isAllowedFrom;
        emit SetAllowedFrom(from, isAllowedFrom);
    }

    function _setAllowedTo(address to, bool isAllowedTo) internal {
        allowedTo[to] = isAllowedTo;
        emit SetAllowedTo(to, isAllowedTo);
    }

    /**
     * @notice Overrides the beforeTokenTransfer function to enforce transfer restrictions
     * @param from the address tokens are being transferred from
     * @param to the address tokens are being transferred to
     * @param amount the amount of tokens being transferred
     */
    function _beforeTokenTransfer(address from, address to, uint256 amount) internal override {
        // if transfer restrictions are enabled
        if (block.timestamp <= transferRestrictionsDisabledAfter) {
            // if both from and to are not whitelisted
            require(
                allowedFrom[from] || allowedTo[to] || from == address(0),
                "BackingEigen._beforeTokenTransfer: from or to must be whitelisted"
            );
        }
        super._beforeTokenTransfer(from, to, amount);
    }
}
