// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "../../contracts/pods/DelayedWithdrawalRouter.sol";
import "../../contracts/permissions/PauserRegistry.sol";

import "../mocks/EigenPodManagerMock.sol";
import "../mocks/Reenterer.sol";

import "forge-std/Test.sol";

contract DelayedWithdrawalRouterUnitTests is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;

    EigenPodManagerMock public eigenPodManagerMock;

    DelayedWithdrawalRouter public delayedWithdrawalRouterImplementation;
    DelayedWithdrawalRouter public delayedWithdrawalRouter;

    Reenterer public reenterer;

    address public pauser = address(555);
    address public unpauser = address(999);

    address initialOwner = address(this);

    uint224[] public delayedWithdrawalAmounts;

    uint256 internal _pseudorandomNumber;

    // index for flag that pauses withdrawals (i.e. 'delayedWithdrawal claims') when set
    uint8 internal constant PAUSED_PAYMENT_CLAIMS = 0;

    mapping(address => bool) public addressIsExcludedFromFuzzedInputs;

    modifier filterFuzzedAddressInputs(address fuzzedAddress) {
        cheats.assume(!addressIsExcludedFromFuzzedInputs[fuzzedAddress]);
        _;
    }

    function setUp() external {
        proxyAdmin = new ProxyAdmin();

        pauserRegistry = new PauserRegistry(pauser, unpauser);

        eigenPodManagerMock = new EigenPodManagerMock();

        delayedWithdrawalRouterImplementation = new DelayedWithdrawalRouter(eigenPodManagerMock);

        uint256 initPausedStatus = 0;
        uint256 withdrawalDelayBlocks = delayedWithdrawalRouterImplementation.MAX_WITHDRAWAL_DELAY_BLOCKS();
        delayedWithdrawalRouter = DelayedWithdrawalRouter(
            address(
                new TransparentUpgradeableProxy(
                    address(delayedWithdrawalRouterImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(DelayedWithdrawalRouter.initialize.selector, initialOwner, pauserRegistry, initPausedStatus, withdrawalDelayBlocks)
                )
            )
        );

        reenterer = new Reenterer();

        // exclude the zero address, the proxyAdmin, and this contract from fuzzed inputs
        addressIsExcludedFromFuzzedInputs[address(0)] = true;
        addressIsExcludedFromFuzzedInputs[address(proxyAdmin)] = true;
        addressIsExcludedFromFuzzedInputs[address(this)] = true;
    }

    function testCannotReinitialize() external {
        uint256 initPausedStatus = 0;
        uint256 withdrawalDelayBlocks = delayedWithdrawalRouter.MAX_WITHDRAWAL_DELAY_BLOCKS();
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        delayedWithdrawalRouter.initialize(initialOwner, pauserRegistry, initPausedStatus, withdrawalDelayBlocks);
    }

    function testCreateDelayedWithdrawalNonzeroAmount(uint224 delayedWithdrawalAmount, address podOwner, address recipient) public filterFuzzedAddressInputs(podOwner) {
        cheats.assume(delayedWithdrawalAmount != 0);
        cheats.assume(recipient != address(0));

        IDelayedWithdrawalRouter.UserDelayedWithdrawals memory userWithdrawalsBefore = delayedWithdrawalRouter.userWithdrawals(recipient);

        address podAddress = address(eigenPodManagerMock.getPod(podOwner));
        cheats.deal(podAddress, delayedWithdrawalAmount);
        cheats.startPrank(podAddress);
        delayedWithdrawalRouter.createDelayedWithdrawal{value: delayedWithdrawalAmount}(podOwner, recipient);
        cheats.stopPrank();

        IDelayedWithdrawalRouter.UserDelayedWithdrawals memory userWithdrawalsAfter = delayedWithdrawalRouter.userWithdrawals(recipient);

        require(userWithdrawalsAfter.delayedWithdrawals.length == userWithdrawalsBefore.delayedWithdrawals.length + 1,
            "userWithdrawalsAfter.delayedWithdrawals.length != userWithdrawalsBefore.delayedWithdrawals.length + 1");

        IDelayedWithdrawalRouter.DelayedWithdrawal memory delayedWithdrawal = userWithdrawalsAfter.delayedWithdrawals[userWithdrawalsAfter.delayedWithdrawals.length - 1];
        require(delayedWithdrawal.amount == delayedWithdrawalAmount, "delayedWithdrawal.amount != delayedWithdrawalAmount");
        require(delayedWithdrawal.blockCreated == block.number, "delayedWithdrawal.blockCreated != block.number");
    }

    function testCreateDelayedWithdrawalZeroAmount(address podOwner, address recipient) public filterFuzzedAddressInputs(podOwner) {
        cheats.assume(recipient != address(0));
        IDelayedWithdrawalRouter.UserDelayedWithdrawals memory userWithdrawalsBefore = delayedWithdrawalRouter.userWithdrawals(recipient);
        uint224 delayedWithdrawalAmount = 0;

        address podAddress = address(eigenPodManagerMock.getPod(podOwner));
        cheats.deal(podAddress, delayedWithdrawalAmount);
        cheats.startPrank(podAddress);
        delayedWithdrawalRouter.createDelayedWithdrawal{value: delayedWithdrawalAmount}(podOwner, recipient);
        cheats.stopPrank();

        IDelayedWithdrawalRouter.UserDelayedWithdrawals memory userWithdrawalsAfter = delayedWithdrawalRouter.userWithdrawals(recipient);

        // verify that no new 'delayedWithdrawal' was created
        require(userWithdrawalsAfter.delayedWithdrawals.length == userWithdrawalsBefore.delayedWithdrawals.length,
            "userWithdrawalsAfter.delayedWithdrawals.length != userWithdrawalsBefore.delayedWithdrawals.length");
    }

    function testCreateDelayedWithdrawalZeroAddress(address podOwner) external {
        uint224 delayedWithdrawalAmount = 0;
        address podAddress = address(eigenPodManagerMock.getPod(podOwner));
        cheats.startPrank(podAddress);
        cheats.expectRevert(bytes("DelayedWithdrawalRouter.createDelayedWithdrawal: recipient cannot be zero address"));
        delayedWithdrawalRouter.createDelayedWithdrawal{value: delayedWithdrawalAmount}(podOwner, address(0));
    }

    function testClaimDelayedWithdrawals(uint8 delayedWithdrawalsToCreate, uint8 maxNumberOfDelayedWithdrawalsToClaim, uint256 pseudorandomNumber_, address recipient, bool useOverloadedFunction)
        public filterFuzzedAddressInputs(recipient)
    {
        // filter contracts out of fuzzed recipient input, since most don't implement a payable fallback function
        cheats.assume(!Address.isContract(recipient));
        // filter out precompile addresses (they won't accept delayedWithdrawal either)
        cheats.assume(uint160(recipient) > 256);
        // filter fuzzed inputs to avoid running out of gas & excessive test run-time
        cheats.assume(delayedWithdrawalsToCreate <= 32);

        address podOwner = address(88888);

        // create the delayedWithdrawals
        _pseudorandomNumber = pseudorandomNumber_;
        uint8 delayedWithdrawalsCreated;
        for (uint256 i = 0; i < delayedWithdrawalsToCreate; ++i) {
            uint224 delayedWithdrawalAmount = uint224(_getPseudorandomNumber());
            if (delayedWithdrawalAmount != 0) {
                testCreateDelayedWithdrawalNonzeroAmount(delayedWithdrawalAmount, podOwner, recipient);
                delayedWithdrawalAmounts.push(delayedWithdrawalAmount);
                delayedWithdrawalsCreated += 1;
            }
        }

        IDelayedWithdrawalRouter.UserDelayedWithdrawals memory userWithdrawalsBefore = delayedWithdrawalRouter.userWithdrawals(recipient);
        uint256 userBalanceBefore = recipient.balance;

        // pre-condition check
        require(userWithdrawalsBefore.delayedWithdrawals.length == delayedWithdrawalsCreated, "userWithdrawalsBefore.delayedWithdrawals.length != delayedWithdrawalsCreated");

        // roll forward the block number enough to make the delayedWithdrawals claimable
        cheats.roll(block.number + delayedWithdrawalRouter.withdrawalDelayBlocks());

        // claim the delayedWithdrawals
        if (delayedWithdrawalsCreated != 0) {
            if (useOverloadedFunction) {
                delayedWithdrawalRouter.claimDelayedWithdrawals(recipient, maxNumberOfDelayedWithdrawalsToClaim);                
            } else {
                cheats.startPrank(recipient);
                delayedWithdrawalRouter.claimDelayedWithdrawals(maxNumberOfDelayedWithdrawalsToClaim);
                cheats.stopPrank();
            }
        }

        IDelayedWithdrawalRouter.UserDelayedWithdrawals memory userWithdrawalsAfter = delayedWithdrawalRouter.userWithdrawals(recipient);
        uint256 userBalanceAfter = recipient.balance;

        // post-conditions
        uint256 numberOfDelayedWithdrawalsThatShouldBeCompleted = (maxNumberOfDelayedWithdrawalsToClaim > delayedWithdrawalsCreated) ? delayedWithdrawalsCreated : maxNumberOfDelayedWithdrawalsToClaim;
        require(userWithdrawalsAfter.delayedWithdrawalsCompleted == userWithdrawalsBefore.delayedWithdrawalsCompleted + numberOfDelayedWithdrawalsThatShouldBeCompleted,
            "userWithdrawalsAfter.delayedWithdrawalsCompleted != userWithdrawalsBefore.delayedWithdrawalsCompleted + numberOfDelayedWithdrawalsThatShouldBeCompleted");
        uint256 totalDelayedWithdrawalAmount = 0;
        for (uint256 i = 0; i < numberOfDelayedWithdrawalsThatShouldBeCompleted; ++i) {
            totalDelayedWithdrawalAmount += delayedWithdrawalAmounts[i];
        }
        require(userBalanceAfter == userBalanceBefore + totalDelayedWithdrawalAmount,
            "userBalanceAfter != userBalanceBefore + totalDelayedWithdrawalAmount");
    }

    /// @notice This function is used to test the getter function 'getClaimableDelayedWithdrawals'
    function testGetClaimableDelayedWithdrawals(uint8 delayedWithdrawalsToCreate, uint224 delayedWithdrawalAmount, address recipient)
        public filterFuzzedAddressInputs(recipient)
    {
        cheats.assume(delayedWithdrawalAmount != 0);
        cheats.assume(delayedWithdrawalsToCreate > 5);
        // filter contracts out of fuzzed recipient input, since most don't implement a payable fallback function
        cheats.assume(!Address.isContract(recipient));
        // filter out precompile addresses (they won't accept delayedWithdrawal either)
        cheats.assume(uint160(recipient) > 256);
        // filter fuzzed inputs to avoid running out of gas & excessive test run-time
        cheats.assume(delayedWithdrawalsToCreate <= 32);

        address podOwner = address(88888);

        // create the delayedWithdrawals
        uint8 delayedWithdrawalsCreated;
        for (uint256 i = 0; i < delayedWithdrawalsToCreate; ++i) {   
            testCreateDelayedWithdrawalNonzeroAmount(delayedWithdrawalAmount, podOwner, recipient);
            delayedWithdrawalAmounts.push(delayedWithdrawalAmount);
            delayedWithdrawalsCreated += 1;
            cheats.roll(block.number + 1); // make sure each delayedWithdrawal has a unique block number
        }

        cheats.roll(block.number + delayedWithdrawalRouter.withdrawalDelayBlocks() + 1 - delayedWithdrawalsToCreate);
        for (uint i = 1; i <= delayedWithdrawalsToCreate; ++i) {
            emit log_named_uint("i", delayedWithdrawalRouter.getClaimableUserDelayedWithdrawals(recipient).length);
            require(delayedWithdrawalRouter.getClaimableUserDelayedWithdrawals(recipient).length == i);
            cheats.roll(block.number + 1);
            
        }
        require(delayedWithdrawalRouter.getClaimableUserDelayedWithdrawals(recipient).length == delayedWithdrawalsCreated, "");

    }


    /**
     * @notice Creates a set of delayedWithdrawals of length (2 * `delayedWithdrawalsToCreate`) where only the first half is claimable, claims using `maxNumberOfDelayedWithdrawalsToClaim` input,
     * and checks that only appropriate delayedWithdrawals were claimed.
     */
    function testClaimDelayedWithdrawalsSomeUnclaimable(uint8 delayedWithdrawalsToCreate, uint8 maxNumberOfDelayedWithdrawalsToClaim, bool useOverloadedFunction) external {
        // filter fuzzed inputs to avoid running out of gas & excessive test run-time
        cheats.assume(delayedWithdrawalsToCreate <= 32);

        address podOwner = address(88888);
        address recipient = address(22222);

        // create the first set of delayedWithdrawals
        _pseudorandomNumber = 1554;
        uint256 delayedWithdrawalsCreated_1;
        for (uint256 i = 0; i < delayedWithdrawalsToCreate; ++i) {
            uint224 delayedWithdrawalAmount = uint224(_getPseudorandomNumber());
            if (delayedWithdrawalAmount != 0) {
                testCreateDelayedWithdrawalNonzeroAmount(delayedWithdrawalAmount, podOwner, recipient);
                delayedWithdrawalAmounts.push(delayedWithdrawalAmount);
                delayedWithdrawalsCreated_1 += 1;
            }
        }

        // roll forward the block number half of the delay window
        cheats.roll(block.number + delayedWithdrawalRouter.withdrawalDelayBlocks() / 2);

        // create the second set of delayedWithdrawals
        uint256 delayedWithdrawalsCreated_2;
        for (uint256 i = 0; i < delayedWithdrawalsToCreate; ++i) {
            uint224 delayedWithdrawalAmount = uint224(_getPseudorandomNumber());
            if (delayedWithdrawalAmount != 0) {
                testCreateDelayedWithdrawalNonzeroAmount(delayedWithdrawalAmount, podOwner, recipient);
                delayedWithdrawalAmounts.push(delayedWithdrawalAmount);
                delayedWithdrawalsCreated_2 += 1;
            }
        }

        // roll forward the block number half of the delay window -- the first set of delayedWithdrawals should now be claimable, while the second shouldn't be!
        cheats.roll(block.number + delayedWithdrawalRouter.withdrawalDelayBlocks() / 2);

        IDelayedWithdrawalRouter.UserDelayedWithdrawals memory userWithdrawalsBefore = delayedWithdrawalRouter.userWithdrawals(recipient);
        uint256 userBalanceBefore = recipient.balance;

        // pre-condition check
        require(userWithdrawalsBefore.delayedWithdrawals.length == delayedWithdrawalsCreated_1 + delayedWithdrawalsCreated_2,
            "userWithdrawalsBefore.delayedWithdrawals.length != delayedWithdrawalsCreated_1 + delayedWithdrawalsCreated_2");

        // claim the delayedWithdrawals
        if (delayedWithdrawalsCreated_1 + delayedWithdrawalsCreated_2 != 0) {
            if (useOverloadedFunction) {
                delayedWithdrawalRouter.claimDelayedWithdrawals(recipient, maxNumberOfDelayedWithdrawalsToClaim);                
            } else {
                cheats.startPrank(recipient);
                delayedWithdrawalRouter.claimDelayedWithdrawals(maxNumberOfDelayedWithdrawalsToClaim);
                cheats.stopPrank();
            }
        }

        IDelayedWithdrawalRouter.UserDelayedWithdrawals memory userWithdrawalsAfter = delayedWithdrawalRouter.userWithdrawals(recipient);
        uint256 userBalanceAfter = recipient.balance;

        // post-conditions
        uint256 numberOfDelayedWithdrawalsThatShouldBeCompleted = (maxNumberOfDelayedWithdrawalsToClaim > delayedWithdrawalsCreated_1) ? delayedWithdrawalsCreated_1 : maxNumberOfDelayedWithdrawalsToClaim;
        require(userWithdrawalsAfter.delayedWithdrawalsCompleted == userWithdrawalsBefore.delayedWithdrawalsCompleted + numberOfDelayedWithdrawalsThatShouldBeCompleted,
            "userWithdrawalsAfter.delayedWithdrawalsCompleted != userWithdrawalsBefore.delayedWithdrawalsCompleted + numberOfDelayedWithdrawalsThatShouldBeCompleted");
        uint256 totalDelayedWithdrawalAmount = 0;
        for (uint256 i = 0; i < numberOfDelayedWithdrawalsThatShouldBeCompleted; ++i) {
            totalDelayedWithdrawalAmount += delayedWithdrawalAmounts[i];
        }
        require(userBalanceAfter == userBalanceBefore + totalDelayedWithdrawalAmount,
            "userBalanceAfter != userBalanceBefore + totalDelayedWithdrawalAmount");
    }

    function testClaimDelayedWithdrawals_NoneToClaim_AttemptToClaimZero(uint256 pseudorandomNumber_, bool useOverloadedFunction) external {
        uint8 delayedWithdrawalsToCreate = 0;
        uint8 maxNumberOfDelayedWithdrawalsToClaim = 0;
        address recipient = address(22222);
        testClaimDelayedWithdrawals(delayedWithdrawalsToCreate, maxNumberOfDelayedWithdrawalsToClaim, pseudorandomNumber_, recipient, useOverloadedFunction);
    }

    function testClaimDelayedWithdrawals_NoneToClaim_AttemptToClaimNonzero(uint256 pseudorandomNumber_, bool useOverloadedFunction) external {
        uint8 delayedWithdrawalsToCreate = 0;
        uint8 maxNumberOfDelayedWithdrawalsToClaim = 2;
        address recipient = address(22222);
        testClaimDelayedWithdrawals(delayedWithdrawalsToCreate, maxNumberOfDelayedWithdrawalsToClaim, pseudorandomNumber_, recipient, useOverloadedFunction);
    }

    function testClaimDelayedWithdrawals_NonzeroToClaim_AttemptToClaimZero(uint256 pseudorandomNumber_, bool useOverloadedFunction) external {
        uint8 delayedWithdrawalsToCreate = 2;
        uint8 maxNumberOfDelayedWithdrawalsToClaim = 0;
        address recipient = address(22222);
        testClaimDelayedWithdrawals(delayedWithdrawalsToCreate, maxNumberOfDelayedWithdrawalsToClaim, pseudorandomNumber_, recipient, useOverloadedFunction);
    }

    function testClaimDelayedWithdrawals_NonzeroToClaim_AttemptToClaimNonzero(uint8 maxNumberOfDelayedWithdrawalsToClaim, uint256 pseudorandomNumber_, bool useOverloadedFunction) external {
        uint8 delayedWithdrawalsToCreate = 2;
        address recipient = address(22222);
        testClaimDelayedWithdrawals(delayedWithdrawalsToCreate, maxNumberOfDelayedWithdrawalsToClaim, pseudorandomNumber_, recipient, useOverloadedFunction);
    }

    function testClaimDelayedWithdrawals_RevertsOnAttemptingReentrancy(bool useOverloadedFunction) external {
        uint8 maxNumberOfDelayedWithdrawalsToClaim = 1;
        address recipient = address(reenterer);
        address podOwner = address(reenterer);

        // create the delayedWithdrawal
        uint224 delayedWithdrawalAmount = 123;
        testCreateDelayedWithdrawalNonzeroAmount(delayedWithdrawalAmount, podOwner, recipient);

        // roll forward the block number enough to make the delayedWithdrawal claimable
        cheats.roll(block.number + delayedWithdrawalRouter.withdrawalDelayBlocks());

        // prepare the Reenterer contract
        address targetToUse = address(delayedWithdrawalRouter);
        uint256 msgValueToUse = 0;
        bytes memory expectedRevertDataToUse = bytes("ReentrancyGuard: reentrant call");
        bytes memory callDataToUse;
        if (useOverloadedFunction) {
            callDataToUse = abi.encodeWithSignature(
                "claimDelayedWithdrawals(address,uint256)", address(22222), maxNumberOfDelayedWithdrawalsToClaim);
        } else {
            callDataToUse  = abi.encodeWithSignature(
                "claimDelayedWithdrawals(uint256)", maxNumberOfDelayedWithdrawalsToClaim);
        }
        reenterer.prepare(targetToUse, msgValueToUse, callDataToUse, expectedRevertDataToUse);

        if (useOverloadedFunction) {
            delayedWithdrawalRouter.claimDelayedWithdrawals(recipient, maxNumberOfDelayedWithdrawalsToClaim);                
        } else {
            cheats.startPrank(recipient);
            delayedWithdrawalRouter.claimDelayedWithdrawals(maxNumberOfDelayedWithdrawalsToClaim);
            cheats.stopPrank();
        }
    }

    function testClaimDelayedWithdrawals_RevertsWhenPaused(bool useOverloadedFunction) external {
        uint8 maxNumberOfDelayedWithdrawalsToClaim = 1;
        address recipient = address(22222);

        // pause delayedWithdrawal claims
        cheats.startPrank(delayedWithdrawalRouter.pauserRegistry().pauser());
        delayedWithdrawalRouter.pause(2 ** PAUSED_PAYMENT_CLAIMS);
        cheats.stopPrank();

        cheats.expectRevert(bytes("Pausable: index is paused"));
        if (useOverloadedFunction) {
            delayedWithdrawalRouter.claimDelayedWithdrawals(recipient, maxNumberOfDelayedWithdrawalsToClaim);                
        } else {
            cheats.startPrank(recipient);
            delayedWithdrawalRouter.claimDelayedWithdrawals(maxNumberOfDelayedWithdrawalsToClaim);
            cheats.stopPrank();
        }
    }

    function testSetWithdrawalDelayBlocks(uint16 valueToSet) external {
        // filter fuzzed inputs to allowed amounts
        cheats.assume(valueToSet <= delayedWithdrawalRouter.MAX_WITHDRAWAL_DELAY_BLOCKS());

        // set the `withdrawalDelayBlocks` variable
        cheats.startPrank(delayedWithdrawalRouter.owner());
        delayedWithdrawalRouter.setWithdrawalDelayBlocks(valueToSet);
        cheats.stopPrank();
        require(delayedWithdrawalRouter.withdrawalDelayBlocks() == valueToSet, "delayedWithdrawalRouter.withdrawalDelayBlocks() != valueToSet");
    }

    function testSetWithdrawalDelayBlocksRevertsWhenCalledByNotOwner(address notOwner) filterFuzzedAddressInputs(notOwner) external {
        cheats.assume(notOwner != delayedWithdrawalRouter.owner());

        uint256 valueToSet = 1;
        // set the `withdrawalDelayBlocks` variable
        cheats.startPrank(notOwner);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        delayedWithdrawalRouter.setWithdrawalDelayBlocks(valueToSet);
        cheats.stopPrank();
    }

    function testSetWithdrawalDelayBlocksRevertsWhenInputValueTooHigh(uint256 valueToSet) external {
        // filter fuzzed inputs to disallowed amounts
        cheats.assume(valueToSet > delayedWithdrawalRouter.MAX_WITHDRAWAL_DELAY_BLOCKS());

        // attempt to set the `withdrawalDelayBlocks` variable
        cheats.startPrank(delayedWithdrawalRouter.owner());
        cheats.expectRevert(bytes("DelayedWithdrawalRouter._setWithdrawalDelayBlocks: newValue too large"));
        delayedWithdrawalRouter.setWithdrawalDelayBlocks(valueToSet);
    }

    function _getPseudorandomNumber() internal returns (uint256) {
        _pseudorandomNumber = uint256(keccak256(abi.encode(_pseudorandomNumber)));
        return _pseudorandomNumber;
    }
}