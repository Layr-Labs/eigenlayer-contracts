using StrategyManager as strategyManager;
methods {
    // external calls to StrategyManager
    function _.stakerStrategyShares(address, address) external => DISPATCHER(true);
    
    // external calls to PauserRegistry
    function _.pauser() external => DISPATCHER(true);
	function _.unpauser() external => DISPATCHER(true);

    // external calls to ERC20
    function _.balanceOf(address) external => DISPATCHER(true);
    function _.transfer(address, uint256) external => DISPATCHER(true);
    function _.transferFrom(address, address, uint256) external => DISPATCHER(true);

	// external calls from StrategyManager to Slasher
    function _.isFrozen(address) external => DISPATCHER(true);
	function _.canWithdraw(address,uint32,uint256) external => DISPATCHER(true);

    // envfree functions
    function totalShares() external returns (uint256) envfree;
    function underlyingToken() external returns (address) envfree;
    function sharesToUnderlyingView(uint256) external returns (uint256) envfree;
    function sharesToUnderlying(uint256) external returns (uint256) envfree;
    function underlyingToSharesView(uint256) external returns (uint256) envfree;
    function underlyingToShares(uint256) external returns (uint256) envfree;
    function shares(address) external returns (uint256) envfree;
}

/**
* Verifies that `totalShares` is always in the set {0, [MIN_NONZERO_TOTAL_SHARES, type(uint256).max]}
* i.e. that `totalShares` is *never* in the range [1, MIN_NONZERO_TOTAL_SHARES - 1]
* Note that this uses that MIN_NONZERO_TOTAL_SHARES = 1e9
*/
invariant totalSharesNeverTooSmall()
    // CVL doesn't appear to parse 1e9, so the literal value is typed out instead.
    (totalShares() == 0) || (totalShares() >= 1000000000);

// // idea based on OpenZeppelin invariant -- see https://github.com/OpenZeppelin/openzeppelin-contracts/blob/formal-verification/certora/specs/ERC20.spec#L8-L22
// ghost sumOfShares() returns uint256 {
//   init_state axiom sumOfShares() == 0;
// }

// hook Sstore currentContract.strategyManager.stakerStrategyShares[KEY address staker][KEY address strategy] uint256 newValue (uint256 oldValue) STORAGE {
//     havoc sumOfShares assuming sumOfShares@new() == sumOfShares@old() + newValue - oldValue;
// }

// invariant totalSharesIsSumOfShares()
//     totalShares() == sumOfShares()