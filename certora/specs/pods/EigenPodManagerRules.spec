using EigenPodManager as EigenPodManager;

methods {
    //// External Calls
	// external calls to DelegationManager 
    function _.undelegate(address) external;
    function _.decreaseDelegatedShares(address,address,uint256) external;
	function _.increaseDelegatedShares(address,address,uint256) external;

    // external calls from DelegationManager to ServiceManager
    function _.updateStakes(address[]) external => NONDET;

	// external calls to StrategyManager
    function _.getDeposits(address) external => DISPATCHER(true);

    // external calls to PauserRegistry
    function _.isPauser(address) external => DISPATCHER(true);
	function _.unpauser() external => DISPATCHER(true);
	
    // envfree functions
    function ownerToPod(address podOwner) external returns (address) envfree;
    function getPod(address podOwner) external returns (address) envfree;
    function ethPOS() external returns (address) envfree;
    function eigenPodBeacon() external returns (address) envfree;
    function hasPod(address podOwner) external returns (bool) envfree;
    function numPods() external returns (uint256) envfree;
    function beaconChainETHStrategy() external returns (address) envfree;

    function Create2.deploy(uint256,bytes32,bytes memory) internal returns (address) => NONDET;
}

// -----------------------------------------
// Constants and Definitions
// -----------------------------------------

definition WAD() returns uint64 = 1000000000000000000; // definition uint64 WAD = 1e18
// -----------------------------------------
// Ghost Mappings and Mirroring
// -----------------------------------------

// Mirror `_beaconChainSlashingFactor`
ghost mapping(address =>  uint64) slashingFactorsGhost {
    init_state axiom forall address staker . slashingFactorsGhost[staker] == 0;
}

// -----------------------------------------
// Hooks for Synchronizing Ghost Mappings
// -----------------------------------------

// Sync slashingFactor when `_beaconChainSlashingFactor` is updated
hook Sstore EigenPodManager._beaconChainSlashingFactor[KEY address operator].slashingFactor uint64 newSlashingFactor (uint64 oldSlashingFactor) {
    require oldSlashingFactor == slashingFactorsGhost[operator];
    slashingFactorsGhost[operator] = newSlashingFactor;
}
hook Sload uint64 slashingFactor EigenPodManager._beaconChainSlashingFactor[KEY address operator].slashingFactor {
    require slashingFactor == slashingFactorsGhost[operator];
}

// -----------------------------------------
// Invariants And Protocol Rules
// -----------------------------------------

/// @title Verifies that ownerToPod[podOwner] is set once (when podOwner deploys a pod), and can otherwise never be updated
/// @property Pod address immutability
rule podAddressNeverChanges(address podOwner) {
   address podAddressBefore = EigenPodManager.ownerToPod[podOwner];
   // perform arbitrary function call
   method f;
   env e;
   calldataarg args;
   f(e,args);
   address podAddressAfter = EigenPodManager.ownerToPod[podOwner];
   assert(podAddressBefore == 0 || podAddressBefore == podAddressAfter,
       "pod address changed after being set!");
}

/// @title BeaconChainSlashingFactor is initialized as WAD and only decreases.
/// @property BeaconChainSlashingFactor is montonically decreasing
rule BeaconChainSlashingFactorMonotonicDec(method f, address staker) {
    env e;
    calldataarg args;

    uint64 BeaconChainSlashingFactorPre = slashingFactorsGhost[staker];
    require BeaconChainSlashingFactorPre <= WAD();

        f(e, args);

    uint64 BeaconChainSlashingFactorPost = slashingFactorsGhost[staker];

    assert BeaconChainSlashingFactorPre >= BeaconChainSlashingFactorPost;
}

/// @title Verifies that podOwnerDepositShares can't go negative, assuming the current balance is a positive int256 value in podOwnerDepositShares
/// @property Pod owner deposited shares remain positive
invariant podOwnerDepositSharesRemainPositive(address podOwner) 
    EigenPodManager.podOwnerDepositShares[podOwner] >= 0;

/// @title This rule verifies that when a specified amount of deposit shares is removed from a staker's account, the staker's total deposited shares decrease exactly by that amount.
/// @property Remove Deposit Shares Integrity 
rule integrityOfRemoveDepositShares(address staker, address strategy, uint256 depositSharesToRemove) {
    env e;

    int256 depositedSharesPre = EigenPodManager.podOwnerDepositShares[staker];

    // avoid overflows and underflows
    require depositedSharesPre - depositSharesToRemove > -max_uint128;
    require depositSharesToRemove + depositedSharesPre < max_uint128;

        removeDepositShares(e, staker, strategy, depositSharesToRemove);

    int256 depositedSharesPost = EigenPodManager.podOwnerDepositShares[staker];

    assert depositedSharesPost == depositedSharesPre - depositSharesToRemove;
}

/// @title 
/*
This rule verifies that when shares are added via the addShares function, 
the staker’s deposit shares increase by exactly the specified amount. 
It also confirms that the function returns the expected number of added shares and 
that the recorded previous deposit shares match the pre-operation value.
*/
/// @property Add Shares Integrity
rule integrityOfAddShares(address staker, address strategy, uint256 shares) {
    env e;

    int256 prevDepositSharesPre = EigenPodManager.podOwnerDepositShares[staker];

    uint256 returnedPrevDeposit; 
    uint256 returnedShares;

    (returnedPrevDeposit, returnedShares) = addShares(e, staker, strategy, shares);

    int256 updatedDepositSharesPost = EigenPodManager.podOwnerDepositShares[staker];

    // Verify that deposit shares have been updated correctly
    assert updatedDepositSharesPost == prevDepositSharesPre + shares;

    // If the updated deposit shares are non-positive, the function should return (0, 0)
    if (updatedDepositSharesPost <= 0) {
        assert returnedPrevDeposit == 0 && returnedShares == 0;
    } else {
        // Otherwise, if the pre-existing deposit shares were negative, the returned previous deposit should be 0 and the shares should be the updated deposited shares;
        // if non-negative, it should equal the pre-state value.
        if (prevDepositSharesPre < 0) {
            assert returnedPrevDeposit == 0 && returnedShares == updatedDepositSharesPost;
        } else {
            assert returnedPrevDeposit == prevDepositSharesPre && returnedShares == shares;
        }
    }
}

/// @title 
/*
This rule ensures correct behavior during token withdrawal. 
It checks that if a staker’s deposited shares are negative before the withdrawal, 
then afterward they either become zero or are increased by exactly the withdrawn shares. Conversely, 
if the deposited shares are non-negative, the withdrawal operation does not alter the staker’s deposited shares.
*/
/// @property Integrity Of Withdraw Shares As Tokens
rule integrityOfWithdrawSharesAsTokens(address staker, address strategy, address erc20, uint256 shares) {
    env e;

    int256 depositedSharesPre = EigenPodManager.podOwnerDepositShares[staker];

        withdrawSharesAsTokens(e, staker, strategy, erc20, shares);

    int256 depositedSharesPost = EigenPodManager.podOwnerDepositShares[staker];

    assert depositedSharesPre < 0 => depositedSharesPost == 0 || depositedSharesPost == depositedSharesPre + shares;
    assert depositedSharesPre >= 0 => depositedSharesPre == depositedSharesPost;
}

/// @title This rule checks that if shares are first added and then removed with the same amount, the staker's total deposited shares remain unchanged, confirming that the operations cancel each other out.
/// @property Add-Then-Remove Shares Neutrality
rule addThenRemoveSharesNoEffect(address staker, address strategy, uint256 depositSharesToRemove) {
    env e;

    int256 depositedSharesPre = EigenPodManager.podOwnerDepositShares[staker];

        addShares(e, staker, strategy, depositSharesToRemove);

        removeDepositShares(e, staker, strategy, depositSharesToRemove);

    int256 depositedSharesPost = EigenPodManager.podOwnerDepositShares[staker];
    
    assert depositedSharesPost == depositedSharesPre;
}

/// @title This rule confirms that if shares are removed and subsequently re-added with the same amount, the overall balance of the staker's deposited shares is preserved, ensuring no net effect from the operations.
/// @property Remove-Then-Add Shares Neutrality
rule removeThenAddSharesNoEffect(address staker, address strategy, uint256 depositSharesToRemove) {
    env e;

    int256 depositedSharesPre = EigenPodManager.podOwnerDepositShares[staker];

        removeDepositShares(e, staker, strategy, depositSharesToRemove);

        addShares(e, staker, strategy, depositSharesToRemove);

    int256 depositedSharesPost = EigenPodManager.podOwnerDepositShares[staker];
    
    assert depositedSharesPost == depositedSharesPre;
}