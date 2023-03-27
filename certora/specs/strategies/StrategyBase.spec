using StrategyManager as strategyManager
methods {
    // external calls to StrategyManager
    stakerStrategyShares(address, address) returns (uint256) => DISPATCHER(true)
    
    // external calls to PauserRegistry
    pauser() returns (address) => DISPATCHER(true)
	unpauser() returns (address) => DISPATCHER(true)

    // external calls to ERC20
    balanceOf(address) returns (uint256) => DISPATCHER(true)
    transfer(address, uint256) returns (bool) => DISPATCHER(true)
    transferFrom(address, address, uint256) returns (bool) => DISPATCHER(true)

	// external calls from StrategyManager to Slasher
    isFrozen(address) returns (bool) => DISPATCHER(true)
	canWithdraw(address,uint32,uint256) returns (bool) => DISPATCHER(true)

    // envfree functions
    totalShares() returns (uint256) envfree
    underlyingToken() returns (address) envfree
    sharesToUnderlyingView(uint256) returns (uint256) envfree
    sharesToUnderlying(uint256) returns (uint256) envfree
    underlyingToSharesView(uint256) returns (uint256) envfree
    underlyingToShares(uint256) returns (uint256) envfree
    shares(address) returns (uint256) envfree
}

/**
* Verifies that `totalShares` is always in the set {0, [MIN_NONZERO_TOTAL_SHARES, type(uint256).max]}
* i.e. that `totalShares` is *never* in the range [1, MIN_NONZERO_TOTAL_SHARES - 1]
* Note that this uses that MIN_NONZERO_TOTAL_SHARES = 1e9
*/
invariant totalSharesNeverTooSmall()
    // CVL doesn't appear to parse 1e9, so the literal value is typed out instead.
    (totalShares() == 0) || (totalShares() >= 1000000000)

// // idea based on OpenZeppelin invariant -- see https://github.com/OpenZeppelin/openzeppelin-contracts/blob/formal-verification/certora/specs/ERC20.spec#L8-L22
// ghost sumOfShares() returns uint256 {
//   init_state axiom sumOfShares() == 0;
// }

// hook Sstore currentContract.strategyManager.stakerStrategyShares[KEY address staker][KEY address strategy] uint256 newValue (uint256 oldValue) STORAGE {
//     havoc sumOfShares assuming sumOfShares@new() == sumOfShares@old() + newValue - oldValue;
// }

// invariant totalSharesIsSumOfShares()
//     totalShares() == sumOfShares()