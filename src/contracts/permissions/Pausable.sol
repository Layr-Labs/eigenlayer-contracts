// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "../interfaces/IPausable.sol";

/**
 * @title Adds pausability to a contract, with pausing & unpausing controlled by the `pauser` and `unpauser` of a PauserRegistry contract.
 * @author Layr Labs, Inc.
 * @notice Terms of Service: https://docs.eigenlayer.xyz/overview/terms-of-service
 * @notice Contracts that inherit from this contract may define their own `pause` and `unpause` (and/or related) functions.
 * These functions should be permissioned as "onlyPauser" which defers to a `PauserRegistry` for determining access control.
 * @dev Pausability is implemented using a uint256, which allows up to 256 different single bit-flags; each bit can potentially pause different functionality.
 * Inspiration for this was taken from the NearBridge design here https://etherscan.io/address/0x3FEFc5A4B1c02f21cBc8D3613643ba0635b9a873#code.
 * For the `pause` and `unpause` functions we've implemented, if you pause, you can only flip (any number of) switches to on/1 (aka "paused"), and if you unpause,
 * you can only flip (any number of) switches to off/0 (aka "paused").
 * If you want a pauseXYZ function that just flips a single bit / "pausing flag", it will:
 * 1) 'bit-wise and' (aka `&`) a flag with the current paused state (as a uint256)
 * 2) update the paused state to this new value
 * @dev We note as well that we have chosen to identify flags by their *bit index* as opposed to their numerical value, so, e.g. defining `DEPOSITS_PAUSED = 3`
 * indicates specifically that if the *third bit* of `_paused` is flipped -- i.e. it is a '1' -- then deposits should be paused
 */
abstract contract Pausable is IPausable {
    /// Constants

    uint256 internal constant _UNPAUSE_ALL = 0;

    uint256 internal constant _PAUSE_ALL = type(uint256).max;

    /// @notice Address of the `PauserRegistry` contract that this contract defers to for determining access control (for pausing).
    IPauserRegistry public immutable pauserRegistry;

    /// Storage

    /// @dev Do not remove, deprecated storage.
    IPauserRegistry private __deprecated_pauserRegistry;

    /// @dev Returns a bitmap representing the paused status of the contract.
    uint256 private _paused;

    /// Modifiers

    /// @dev Thrown if the caller is not a valid pauser according to the pauser registry.
    modifier onlyPauser() {
        require(pauserRegistry.isPauser(msg.sender), OnlyPauser());
        _;
    }

    /// @dev Thrown if the caller is not a valid unpauser according to the pauser registry.
    modifier onlyUnpauser() {
        require(msg.sender == pauserRegistry.unpauser(), OnlyUnpauser());
        _;
    }

    /// @dev Thrown if the contract is paused, i.e. if any of the bits in `_paused` is flipped to 1.
    modifier whenNotPaused() {
        require(_paused == 0, CurrentlyPaused());
        _;
    }

    /// @dev Thrown if the `indexed`th bit of `_paused` is 1, i.e. if the `index`th pause switch is flipped.
    modifier onlyWhenNotPaused(
        uint8 index
    ) {
        require(!paused(index), CurrentlyPaused());
        _;
    }

    /// Construction

    constructor(
        IPauserRegistry _pauserRegistry
    ) {
        require(address(_pauserRegistry) != address(0), InputAddressZero());
        pauserRegistry = _pauserRegistry;
    }

    /// @inheritdoc IPausable
    function pause(
        uint256 newPausedStatus
    ) external onlyPauser {
        uint256 currentPausedStatus = _paused;
        // verify that the `newPausedStatus` does not *unflip* any bits (i.e. doesn't unpause anything, all 1 bits remain)
        require((currentPausedStatus & newPausedStatus) == currentPausedStatus, InvalidNewPausedStatus());
        _setPausedStatus(newPausedStatus);
    }

    /// @inheritdoc IPausable
    function pauseAll() external onlyPauser {
        _setPausedStatus(_PAUSE_ALL);
    }

    /// @inheritdoc IPausable
    function unpause(
        uint256 newPausedStatus
    ) external onlyUnpauser {
        uint256 currentPausedStatus = _paused;
        // verify that the `newPausedStatus` does not *flip* any bits (i.e. doesn't pause anything, all 0 bits remain)
        require(((~currentPausedStatus) & (~newPausedStatus)) == (~currentPausedStatus), InvalidNewPausedStatus());
        _paused = newPausedStatus;
        emit Unpaused(msg.sender, newPausedStatus);
    }

    /// @inheritdoc IPausable
    function paused() public view virtual returns (uint256) {
        return _paused;
    }

    /// @inheritdoc IPausable
    function paused(
        uint8 index
    ) public view virtual returns (bool) {
        uint256 mask = 1 << index;
        return ((_paused & mask) == mask);
    }

    /// @dev Internal helper for setting the paused status, and emitting the corresponding event.
    function _setPausedStatus(
        uint256 pausedStatus
    ) internal {
        _paused = pausedStatus;
        emit Paused(msg.sender, pausedStatus);
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[48] private __gap;
}
