// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "src/contracts/interfaces/IETHPOSDeposit.sol";

import "src/test/mocks/BeaconChainMock.t.sol";
import "src/test/mocks/EIP_4788_Oracle_Mock.t.sol";
import "src/test/utils/TimeMachine.t.sol";

import "src/contracts/interfaces/IStrategy.sol";

/// -----------------------------------------------------------------------
/// General Constants
/// -----------------------------------------------------------------------

/// @dev Returns the address of the Foundry Hevm cheatcodes.
Vm constant cheats = Vm(0x7109709ECfa91a80626fF3989D68f67F5b1DD12D); // Vm(address(uint160(uint256(keccak256("hevm cheat code")))))
/// @dev Returns the address that has pauser privileges.
address constant PAUSER = address(0x0000000000000000000000000000706175736572); // bytes32("pauser")
/// @dev Returns the address that has unpauser privileges.
address constant UNPAUSER = address(0x000000000000000000000000756E706175736572); // bytes32("unpauser")

/// -----------------------------------------------------------------------
/// Beacon Chain Constants
/// -----------------------------------------------------------------------

/// @dev Returns the address of the mock implementation of the Ethereum beacon chain.
BeaconChainMock constant beaconChain = BeaconChainMock(address(0x0000000000000000626561636f6e636861696e6D6f636b)); // bytes32("beacon chain")
/// @dev Returns the address of the mock beacon chain ETH strategy.
IStrategy constant BEACONCHAIN_ETH_STRAT = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
/// @dev Returns the address of the mock beacon chain ETH token.
IERC20 constant NATIVE_ETH = IERC20(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);
/// @dev Returns the address of the mock beacon chain POS deposit contract.
IETHPOSDeposit constant DEPOSIT_CONTRACT = IETHPOSDeposit(0x00000000219ab540356cBB839Cbe05303d7705Fa);
/// @dev Returns the address of the mock beacon chain EIP-4788 oracle.
EIP_4788_Oracle_Mock constant EIP_4788_ORACLE = EIP_4788_Oracle_Mock(0x000F3df6D732807Ef1319fB7B8bB8522d0Beac02);

/// -----------------------------------------------------------------------
/// Integration Test Constants
/// -----------------------------------------------------------------------

/// @dev Returns you to 1985 if you reach 88mph, but we just use it for fetching previous state.
TimeMachine constant timeMachine = TimeMachine(address(0x000000000000000074696D65206D616368696e65)); // bytes32("time machine")
/// @dev Returns the genesis time for the beacon chain when running locally.
uint64 constant GENESIS_TIME_LOCAL = 1 hours * 12;
/// @dev Returns the genesis time for the beacon chain when forking mainnet.
uint64 constant GENESIS_TIME_MAINNET = 1_606_824_023;
/// @dev Returns the block number to fork from mainnet.
/// @dev Current value is post "Protocol Council" upgrade.
uint constant MAINNET_FORK_BLOCK = 21_616_692;
/// @dev Returns the number of LST strategies to deploy.
uint8 constant NUM_LST_STRATS = 8;

/// @dev Returns the minimum balance a user can hold.
uint constant MIN_BALANCE = 1e6;
/// @dev Returns the maximum balance a user can hold.
uint constant MAX_BALANCE = 5e6;
/// @dev Returns the conversion factor from Gwei to Wei.
uint constant GWEI_TO_WEI = 1e9;

/// @dev Types representing the different types of assets a ranomized users can hold.
uint constant NO_ASSETS = (1 << 0); // will have no assets
uint constant HOLDS_LST = (1 << 1); // will hold some random amount of LSTs
uint constant HOLDS_ETH = (1 << 2); // will hold some random amount of ETH
uint constant HOLDS_ALL = (1 << 3); // will always hold ETH, and some LSTs
uint constant HOLDS_MAX = (1 << 4); // will hold every LST and ETH (used for testing max strategies)

/// @dev Types representing the different types of users that can be created.
uint constant DEFAULT = (1 << 0);
uint constant ALT_METHODS = (1 << 1);

/// @dev Types representing the different types of forks that can be simulated.
uint constant LOCAL = (1 << 0);
uint constant MAINNET = (1 << 1);
uint constant HOLESKY = (1 << 2);
