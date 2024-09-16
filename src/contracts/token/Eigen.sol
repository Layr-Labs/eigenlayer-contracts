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
    /// @notice mapping of minter addresses to the timestamp after which they are allowed to mint
    mapping(address => uint256) public mintAllowedAfter;
    /// @notice mapping of minter addresses to the amount of tokens they are allowed to mint
    mapping(address => uint256) public mintingAllowance;

    /// @notice the timestamp after which transfer restrictions are disabled
    uint256 public transferRestrictionsDisabledAfter;
    /// @notice mapping of addresses that are allowed to transfer tokens to any address
    mapping(address => bool) public allowedFrom;
    /// @notice mapping of addresses that are allowed to receive tokens from any address
    mapping(address => bool) public allowedTo;

    /// @notice event emitted when the allowedFrom status of an address is set
    event SetAllowedFrom(address indexed from, bool isAllowedFrom);
    /// @notice event emitted when the allowedTo status of an address is set
    event SetAllowedTo(address indexed to, bool isAllowedTo);
    /// @notice event emitted when a minter mints
    event Mint(address indexed minter, uint256 amount);
    /// @notice event emitted when the transfer restrictions disabled
    event TransferRestrictionsDisabled();

    constructor(
        IERC20 _bEIGEN
    ) {
        bEIGEN = _bEIGEN;
        _disableInitializers();
    }

    /**
     * @notice An initializer function that sets initial values for the contract's state variables.
     * @param minters the addresses that are allowed to mint
     * @param mintingAllowances the amount of tokens that each minter is allowed to mint
     */
    function initialize(
        address initialOwner,
        address[] memory minters,
        uint256[] memory mintingAllowances,
        uint256[] memory mintAllowedAfters
    ) public initializer {
        __Ownable_init();
        __ERC20_init("Eigen", "EIGEN");
        _transferOwnership(initialOwner);
        __ERC20Permit_init("EIGEN");

        require(
            minters.length == mintingAllowances.length,
            "Eigen.initialize: minters and mintingAllowances must be the same length"
        );
        require(
            minters.length == mintAllowedAfters.length,
            "Eigen.initialize: minters and mintAllowedAfters must be the same length"
        );
        // set minting allowances for each minter
        for (uint256 i = 0; i < minters.length; i++) {
            mintingAllowance[minters[i]] = mintingAllowances[i];
            mintAllowedAfter[minters[i]] = mintAllowedAfters[i];
            // allow each minter to transfer tokens
            allowedFrom[minters[i]] = true;
            emit SetAllowedFrom(minters[i], true);
        }

        // set transfer restrictions to be disabled at type(uint256).max to be set down later
        transferRestrictionsDisabledAfter = type(uint256).max;
    }

    /**
     * @notice This function allows the owner to set the allowedFrom status of an address
     * @param from the address whose allowedFrom status is being set
     * @param isAllowedFrom the new allowedFrom status
     */
    function setAllowedFrom(address from, bool isAllowedFrom) external onlyOwner {
        allowedFrom[from] = isAllowedFrom;
        emit SetAllowedFrom(from, isAllowedFrom);
    }

    /**
     * @notice This function allows the owner to set the allowedTo status of an address
     * @param to the address whose allowedTo status is being set
     * @param isAllowedTo the new allowedTo status
     */
    function setAllowedTo(address to, bool isAllowedTo) external onlyOwner {
        allowedTo[to] = isAllowedTo;
        emit SetAllowedTo(to, isAllowedTo);
    }

    /**
     * @notice Allows the owner to disable transfer restrictions
     */
    function disableTransferRestrictions() external onlyOwner {
        require(
            transferRestrictionsDisabledAfter == type(uint256).max,
            "Eigen.disableTransferRestrictions: transfer restrictions are already disabled"
        );
        transferRestrictionsDisabledAfter = 0;
        emit TransferRestrictionsDisabled();
    }

    /**
     * @notice This function allows minter to mint tokens
     */
    function mint() external {
        require(mintingAllowance[msg.sender] > 0, "Eigen.mint: msg.sender has no minting allowance");
        require(block.timestamp > mintAllowedAfter[msg.sender], "Eigen.mint: msg.sender is not allowed to mint yet");
        uint256 amount = mintingAllowance[msg.sender];
        mintingAllowance[msg.sender] = 0;
        _mint(msg.sender, amount);
        emit Mint(msg.sender, amount);
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
                from == address(0) || to == address(0) || allowedFrom[from] || allowedTo[to],
                "Eigen._beforeTokenTransfer: from or to must be whitelisted"
            );
        }
        super._beforeTokenTransfer(from, to, amount);
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
