using StrategyManager as strategyManager;
methods {
    // external calls to PauserRegistry
    function _.isPauser(address) external => DISPATCHER(true);
	function _.unpauser() external => DISPATCHER(true);

    // external calls to ERC20
    function _.balanceOf(address) external => DISPATCHER(true);
    function _.transfer(address, uint256) external => DISPATCHER(true);
    function _.decimals() external => DISPATCHER(true);

    // envfree functions
    function totalShares() external returns (uint256) envfree;
    function underlyingToken() external returns (address) envfree;
    function sharesToUnderlyingView(uint256) external returns (uint256) envfree;
    function sharesToUnderlying(uint256) external returns (uint256) envfree;
    function underlyingToSharesView(uint256) external returns (uint256) envfree;
    function underlyingToShares(uint256) external returns (uint256) envfree;
    function shares(address) external returns (uint256) envfree;
}

// // idea based on OpenZeppelin invariant -- see https://github.com/OpenZeppelin/openzeppelin-contracts/blob/formal-verification/certora/specs/ERC20.spec#L8-L22
// ghost sumOfShares() returns uint256 {
//   init_state axiom sumOfShares() == 0;
// }

// hook Sstore currentContract.strategyManager.stakerStrategyShares[KEY address staker][KEY address strategy] uint256 newValue (uint256 oldValue) STORAGE {
//     havoc sumOfShares assuming sumOfShares@new() == sumOfShares@old() + newValue - oldValue;
// }

// invariant totalSharesIsSumOfShares()
//     totalShares() == sumOfShares()

use builtin rule sanity;