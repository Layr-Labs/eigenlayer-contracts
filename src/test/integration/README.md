## EigenLayer Core Integration Testing

### What the Hell?

Good question.

This folder contains the integration framework and tests for Eigenlayer core, which orchestrates the deployment of all EigenLayer core contracts to fuzz high-level user flows across multiple user and asset types, and supports time-travelling state lookups to quickly compare past and present states (please try to avoid preventing your own birth).

**If you want to know how to run the tests**:

* Local: `forge t --mc Integration`
* Mainnet fork tests: `env FOUNDRY_PROFILE=forktest forge t --mc Integration`

Note that for mainnet fork tests, you'll need to set the `RPC_MAINNET` environment variable to your RPC provider of choice!

**If you want to know where the tests are**, take a look at `/tests`. We're doing one test contract per top-level flow, and defining multiple test functions for variants on that flow. 

e.g. if you're testing the flow "deposit into strategies -> delegate to operator -> queue withdrawal -> complete withdrawal", that's it's own test contract. For variants where withdrawals are completed "as tokens" vs "as shares," those are their own functions inside that contract.

Looking at the current tests is a good place to start.

**If you want to know how we're fuzzing these flows**, take a look at how we're using the `_configRand` method at the start of each test, which accepts bitmaps for the types of users and assets you want to spawn during the test.

During the test, the config passed into `_configRand` will randomly generate only the values you configure:
* `assetTypes` affect the assets granted to Users when they are first created. You can use this to ensure your flows and assertions work when users are holding only LSTs, native ETH, or some combination.
* `userTypes` affect the actual `User` contract being deployed. The `DEFAULT` flag deploys the base `User` contract, while `ALT_METHODS` deploys a version that derives from the same contract, but overrides some methods to use "functionWithSignature" and other variants.

Here's an example:

```solidity
function testFuzz_deposit_delegate_EXAMPLE(uint24 _random) public {   
    // When new Users are created, they will choose a random configuration from these params.
    // `_randomSeed` will be the starting seed for all random lookups.
    _configRand({
        _randomSeed: _random,
        _assetTypes: HOLDS_LST,
        _userTypes: DEFAULT | ALT_METHODS
    });

    // Because of the `assetTypes` flags above, this will create two Users for our test,
    // each of which holds some random assortment of LSTs.
    (
        User staker,
        IStrategy[] memory strategies, 
        uint[] memory tokenBalances
    ) = _newRandomStaker();
    (User operator, ,) = _newRandomOperator();

    // Because of the `userTypes` flags above, this user might be using either:
    // - `strategyManager.depositIntoStrategy`
    // - `strategyManager.depositIntoStrategyWithSignature`
    staker.depositIntoEigenlayer(strategies, tokenBalances);
    // assertions go here
}
```

### `print` Library Overview

The `print` library provides utilities for structured and stylized logging to streamline debugging in EigenLayer contracts. It includes helpers for common tasks such as logging user actions, integration test details. The library works in conjunction with the `Logger` abstract contract, which must be inherited by contracts that use `print`.

NOTE: `print` is intended to be used in conjunction with `console.log`.

```solidity
import "src/test/utils/Logger.t.sol";

// The `User` contract inherits `Logger` to enable `print` functionalities.
contract User is Logger {
    string public _name;

    constructor(string memory name) {
        _name = name;
    }

    function NAME() public override view returns (string memory) {
        return _name;
    }

    function vote(User who, uint256 votes) public {
        print.method(
            "vote", 
            string.concat(
                "{who: ", who.NAME_COLORED(), 
                ", votes: ", votes.asWad(), 
            "}"
            )
        );
    }
}

User alice = User("Alice");
User bob = User("Bob");

// Will emit a log equal to: `console.log("Alice.vote({who: Bob, votes: 1.0 (wad)})");`
alice.vote(bob, 1 ether);
```

**If you want to know about the time travel**, there's a few things to note:

The main feature we're using is foundry's `cheats.snapshotState()` and `cheats.revertToState(snapshot)` to zip around in time. You can look at the [Cheatcodes Reference](https://book.getfoundry.sh/cheatcodes/#cheatcodes-interface) to get some idea, but the docs aren't actually correct. The best thing to do is look through our tests and see how it's being used. If you see an assertion called `assert_Snap_...`, that's using the `TimeMachine` under the hood. 

Speaking of, the `TimeMachine` is a global contract that controls the time, fate, and destiny of all who use it.
* `Users` use the `TimeMachine` to snapshot chain state *before* every action they perform. (see the [`User.createSnapshot`](https://github.com/layr-labs/eigenlayer-contracts/blob/c5193f7bff00903a4323be2a1500cbf7137a83e9/src/test/integration/User.t.sol#L43-L46) modifier).
* `IntegrationBase` uses a `timewarp` modifier to quickly fetch state "from before the last user action". These are leveraged within various `assert_Snap_XYZ` methods to allow the test to quickly compare previous and current values. ([example assertion method](https://github.com/layr-labs/eigenlayer-contracts/blob/c99e847709852d7246c73b7d72d44bba368b760e/src/test/integration/IntegrationBase.t.sol#L146-L148))

This means that tests can perform user actions with very little setup or "reading prior state", and perform all the important assertions after each action. For example:

```solidity
function testFuzz_deposit_delegate_EXAMPLE(uint24 _random) public {   
    // ... test setup goes above here
    
    // This snapshots state before the deposit.
    staker.depositIntoEigenlayer(strategies, tokenBalances);
    // This checks the staker's shares from before `depositIntoEigenlayer`, and compares
    // them to their shares after `depositIntoEigenlayer`.
    assert_Snap_AddedStakerShares(staker, strategies, expectedShares, "failed to award staker shares");

    // This snapshots state before delegating.
    staker.delegateTo(operator);
    // This checks the operator's `operatorShares` before the staker delegated to them, and
    // compares those shares to the `operatorShares` after the staker delegated.
    assert_Snap_AddedOperatorShares(operator, strategies, expectedShares, "failed to award operator shares");
}
```

### Additional Important Concepts

* Most testing logic and checks are performed at the test level. `IntegrationBase` has primarily helpers and a few sanity checks, but the current structure exists to make it clear what's being tested by reading the test itself.
* Minimal logic/assertions/cheats used in User contract. These are for carrying out user behaviors, only. Exception:
    * User methods snapshot state before performing actions
* Top-level error messages are passed into helper assert methods so that it's always clear where an error came from
* User contract should have an interface as similar as possible to the contract interfaces, so it feels like calling an EigenLayer method rather than some weird abstraction. Exceptions for things like:
    * `user.depositIntoEigenLayer(strats, tokenBalances)` - because this deposits all strategies/shares and may touch either Smgr or Emgr

### What needs to be done?

* Suggest or PR cleanup if you have ideas. Currently, the `IntegrationDeployer` contract is pretty messy.
* Coordinate in Slack to pick out some user flows to write tests for!

#### Reduce RPC spam for fork tests

Currently our mainnet fork tests spam whatever RPC we use. We can improve this in the future - apparently the meta is:

> Use an anvil node to fork the network, you can write a script to make some changes to the forked network for setup etc, then fork your local node in your test.
> Effectively you just setup an anvil node with the command
`anvil -f RPC_URL `
You can use `anvil -h` for more info on what it can do.

> Then in your test you use the vm.createSelectFork command in your setup with the argument to point to your local anvil node which is basically a copy of the rpc you set it up as. 
>  If you want to do some setup before running your tests you can write a script file and broadcast the setup transactions to your local anvil node (make sure to use one of the private keys anvil gives you)