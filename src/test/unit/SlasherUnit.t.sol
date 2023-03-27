// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import "forge-std/Test.sol";

import "../../contracts/core/Slasher.sol";
import "../../contracts/permissions/PauserRegistry.sol";
import "../../contracts/strategies/StrategyWrapper.sol";

import "../mocks/DelegationMock.sol";
import "../mocks/EigenPodManagerMock.sol";
import "../mocks/StrategyManagerMock.sol";
import "../mocks/Reenterer.sol";
import "../mocks/Reverter.sol";

import "../mocks/ERC20Mock.sol";

contract SlasherUnitTests is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    uint256 private constant HEAD = 0;

    uint256 private constant _NULL = 0;
    uint256 private constant _HEAD = 0;

    bool private constant _PREV = false;
    bool private constant _NEXT = true;

    uint8 internal constant PAUSED_OPT_INTO_SLASHING = 0;
    uint8 internal constant PAUSED_FIRST_STAKE_UPDATE = 1;
    uint8 internal constant PAUSED_NEW_FREEZING = 2;

    uint32 internal constant MAX_CAN_SLASH_UNTIL = type(uint32).max;

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;

    Slasher public slasherImplementation;
    Slasher public slasher;
    StrategyManagerMock public strategyManagerMock;
    DelegationMock public delegationMock;
    EigenPodManagerMock public eigenPodManagerMock;

    Reenterer public reenterer;

    address public pauser = address(555);
    address public unpauser = address(999);

    address initialOwner = address(this);





    IERC20 public dummyToken;
    StrategyWrapper public dummyStrat;







    uint256[] public emptyUintArray;

    // used as transient storage to fix stack-too-deep errors
    uint32 contractCanSlashOperatorUntilBefore;
    uint256 linkedListLengthBefore;
    uint256 middlewareTimesLengthBefore;
    bool nodeExists;
    uint256 prevNode;
    uint256 nextNode;
    ISlasher.MiddlewareDetails middlewareDetailsBefore;
    ISlasher.MiddlewareDetails middlewareDetailsAfter;

    mapping(address => bool) public addressIsExcludedFromFuzzedInputs;

    modifier filterFuzzedAddressInputs(address fuzzedAddress) {
        cheats.assume(!addressIsExcludedFromFuzzedInputs[fuzzedAddress]);
        _;
    }

    function setUp() virtual public {
        proxyAdmin = new ProxyAdmin();

        pauserRegistry = new PauserRegistry(pauser, unpauser);

        delegationMock = new DelegationMock();
        eigenPodManagerMock = new EigenPodManagerMock();
        strategyManagerMock = new StrategyManagerMock();
        slasherImplementation = new Slasher(strategyManagerMock, delegationMock);
        slasher = Slasher(
            address(
                new TransparentUpgradeableProxy(
                    address(slasherImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(Slasher.initialize.selector, initialOwner, pauserRegistry, 0/*initialPausedStatus*/)
                )
            )
        );
        dummyToken = new ERC20Mock();
        dummyStrat = new StrategyWrapper(strategyManagerMock, dummyToken);

        // excude the zero address and the proxyAdmin from fuzzed inputs
        addressIsExcludedFromFuzzedInputs[address(0)] = true;
        addressIsExcludedFromFuzzedInputs[address(proxyAdmin)] = true;
    }

    function testCannotReinitialize() public {
        cheats.expectRevert(bytes("Initializable: contract is already initialized"));
        slasher.initialize(initialOwner, pauserRegistry, 0);
    }

    function testOptIntoSlashing(address operator, address contractAddress)
        public
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        delegationMock.setIsOperator(operator, true);

        cheats.startPrank(operator);
        slasher.optIntoSlashing(contractAddress);
        cheats.stopPrank();

        assertEq(slasher.contractCanSlashOperatorUntil(operator, contractAddress), MAX_CAN_SLASH_UNTIL);
        require(slasher.canSlash(operator, contractAddress), "contract was not properly granted slashing permission");
    }

    function testOptIntoSlashing_RevertsWhenPaused() public {
        address operator = address(this);
        address contractAddress = address(this);

        // pause opting into slashing
        cheats.startPrank(pauser);
        slasher.pause(2 ** PAUSED_OPT_INTO_SLASHING);
        cheats.stopPrank();

        cheats.startPrank(operator);
        cheats.expectRevert(bytes("Pausable: index is paused"));
        slasher.optIntoSlashing(contractAddress);
        cheats.stopPrank();
    }

    function testOptIntoSlashing_RevertsWhenCallerNotOperator(address notOperator) public filterFuzzedAddressInputs(notOperator) {
        require(!delegationMock.isOperator(notOperator), "caller is an operator -- this is assumed false");
        address contractAddress = address(this);

        cheats.startPrank(notOperator);
        cheats.expectRevert(bytes("Slasher.optIntoSlashing: msg.sender is not a registered operator"));
        slasher.optIntoSlashing(contractAddress);
        cheats.stopPrank();
    }

    function testFreezeOperator(address toBeFrozen, address freezingContract) public
        filterFuzzedAddressInputs(toBeFrozen)
        filterFuzzedAddressInputs(freezingContract)
    {
        testOptIntoSlashing(toBeFrozen, freezingContract);
        cheats.startPrank(freezingContract);
        slasher.freezeOperator(toBeFrozen);
        cheats.stopPrank();

        require(slasher.isFrozen(toBeFrozen), "operator not properly frozen");
    }

    function testFreezeOperatorTwice(address toBeFrozen, address freezingContract) public {
        testFreezeOperator(toBeFrozen, freezingContract);
        testFreezeOperator(toBeFrozen, freezingContract);
    }

    function testFreezeOperator_RevertsWhenPaused(address toBeFrozen, address freezingContract) external
        filterFuzzedAddressInputs(toBeFrozen)
        filterFuzzedAddressInputs(freezingContract)
    {
        testOptIntoSlashing(toBeFrozen, freezingContract);

        // pause freezing
        cheats.startPrank(pauser);
        slasher.pause(2 ** PAUSED_NEW_FREEZING);
        cheats.stopPrank();

        cheats.startPrank(freezingContract);
        cheats.expectRevert(bytes("Pausable: index is paused"));
        slasher.freezeOperator(toBeFrozen);
        cheats.stopPrank();
    }

    function testFreezeOperator_WhenCallerDoesntHaveSlashingPermission(address toBeFrozen, address freezingContract) external
        filterFuzzedAddressInputs(toBeFrozen)
        filterFuzzedAddressInputs(freezingContract)
    {
        cheats.startPrank(freezingContract);
        cheats.expectRevert(bytes("Slasher.freezeOperator: msg.sender does not have permission to slash this operator"));
        slasher.freezeOperator(toBeFrozen);
        cheats.stopPrank();
    }

    function testResetFrozenStatus(uint8 numberOfOperators, uint256 pseudorandomInput) external {
        // sanity filtering
        cheats.assume(numberOfOperators <= 16);

        address contractAddress = address(this);

        address[] memory operatorAddresses = new address[](numberOfOperators);
        bool[] memory operatorFrozen = new bool[](numberOfOperators);
        for (uint256 i = 0; i < numberOfOperators; ++i) {
            address operatorAddress = address(uint160(8888 + i));
            operatorAddresses[i] = operatorAddress;
            testOptIntoSlashing(operatorAddress, contractAddress);
            bool freezeOperator = (pseudorandomInput % 2 == 0) ? false : true;
            pseudorandomInput = uint256(keccak256(abi.encodePacked(pseudorandomInput)));
            operatorFrozen[i] = freezeOperator;
            if (freezeOperator) {
                testFreezeOperator(operatorAddress, contractAddress);
            }
        }

        cheats.startPrank(slasher.owner());
        slasher.resetFrozenStatus(operatorAddresses);
        cheats.stopPrank();

        for (uint256 i = 0; i < numberOfOperators; ++i) {
            require(!slasher.isFrozen(operatorAddresses[i]), "operator frozen improperly (not unfrozen when should be)");
        }
    }

    function testResetFrozenStatus_RevertsWhenCalledByNotOwner(address notOwner) external filterFuzzedAddressInputs(notOwner)  {        
        // sanity filtering
        cheats.assume(notOwner != slasher.owner());

        address[] memory operatorAddresses = new address[](1);

        cheats.startPrank(notOwner);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        slasher.resetFrozenStatus(operatorAddresses);
        cheats.stopPrank();
    }

    function testRecordFirstStakeUpdate(address operator, address contractAddress, uint32 serveUntil)
        public
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        testOptIntoSlashing(operator, contractAddress);
        _testRecordFirstStakeUpdate(operator, contractAddress, serveUntil);
    }

    // internal function corresponding to the bulk of the logic in `testRecordFirstStakeUpdate`, so we can reuse it elsewhere without calling `testOptIntoSlashing`
    function _testRecordFirstStakeUpdate(address operator, address contractAddress, uint32 serveUntil)
        internal
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {

        linkedListLengthBefore = slasher.operatorWhitelistedContractsLinkedListSize(operator);
        middlewareDetailsBefore = slasher.whitelistedContractDetails(operator, contractAddress);
        contractCanSlashOperatorUntilBefore = slasher.contractCanSlashOperatorUntil(operator, contractAddress);

        ISlasher.MiddlewareTimes memory mostRecentMiddlewareTimesStructBefore;
        // fetch the most recent struct, if at least one exists (otherwise leave the struct uninitialized)
        middlewareTimesLengthBefore = slasher.middlewareTimesLength(operator);
        if (middlewareTimesLengthBefore != 0) {
            mostRecentMiddlewareTimesStructBefore = slasher.operatorToMiddlewareTimes(operator, middlewareTimesLengthBefore - 1);
        }

        cheats.startPrank(contractAddress);
        slasher.recordFirstStakeUpdate(operator, serveUntil);
        cheats.stopPrank();

        ISlasher.MiddlewareTimes memory mostRecentMiddlewareTimesStructAfter = slasher.operatorToMiddlewareTimes(operator, slasher.middlewareTimesLength(operator) - 1);

        // check that linked list size increased appropriately
        require(slasher.operatorWhitelistedContractsLinkedListSize(operator) == linkedListLengthBefore + 1, "linked list length did not increase when it should!");
        // get the linked list entry for the `contractAddress`
        (nodeExists, prevNode, nextNode) = slasher.operatorWhitelistedContractsLinkedListEntry(operator, contractAddress);
        // verify that the node exists
        require(nodeExists, "node does not exist");

        // if the `serveUntil` time is greater than the previous maximum, then an update must have been pushed and the `latestServeUntil` must be equal to the `serveUntil` input
        if (serveUntil > mostRecentMiddlewareTimesStructBefore.latestServeUntil) {
            require(slasher.middlewareTimesLength(operator) == middlewareTimesLengthBefore + 1, "MiddlewareTimes struct not pushed to array");
            require(mostRecentMiddlewareTimesStructAfter.latestServeUntil == serveUntil, "latestServeUntil not updated correctly");
        } else {
            require(mostRecentMiddlewareTimesStructAfter.latestServeUntil  == mostRecentMiddlewareTimesStructBefore.latestServeUntil, "latestServeUntil updated incorrectly");
        }
        // if this is the first MiddlewareTimes struct in the array, then an update must have been pushed and the `stalestUpdateBlock` *must* be equal to the current block
        if (middlewareTimesLengthBefore == 0) {
            require(slasher.middlewareTimesLength(operator) == middlewareTimesLengthBefore + 1,
                "MiddlewareTimes struct not pushed to array");
            require(mostRecentMiddlewareTimesStructAfter.stalestUpdateBlock == block.number,
                "stalestUpdateBlock not updated correctly -- contractAddress is first list entry");
        // otherwise, we check if the `contractAddress` is the head of the list. If it *is*, then prevNode will be _HEAD, and...
        } else if (prevNode == _HEAD) {
            // if nextNode is _HEAD, then the this indicates that `contractAddress` is actually the only list entry
            if (nextNode != _HEAD) {
                // if nextNode is not the only list entry, then `latestUpdateBlock` should update to a more recent time (if applicable!)
                uint32 nextMiddlewaresLeastRecentUpdateBlock = slasher.whitelistedContractDetails(operator, _uintToAddress(nextNode)).latestUpdateBlock;
                uint32 newValue = (nextMiddlewaresLeastRecentUpdateBlock < block.number) ? nextMiddlewaresLeastRecentUpdateBlock: uint32(block.number);
                require(mostRecentMiddlewareTimesStructAfter.stalestUpdateBlock == newValue,
                    "stalestUpdateBlock not updated correctly -- should have updated to newValue");                
            } else {
                require(mostRecentMiddlewareTimesStructAfter.stalestUpdateBlock == block.number,
                    "stalestUpdateBlock not updated correctly -- contractAddress is only list entry");
            }
        }

        middlewareDetailsAfter = slasher.whitelistedContractDetails(operator, contractAddress);
        require(middlewareDetailsAfter.latestUpdateBlock == block.number,
            "latestUpdateBlock not updated correctly");
        require(middlewareDetailsAfter.contractCanSlashOperatorUntil == middlewareDetailsBefore.contractCanSlashOperatorUntil,
            "contractCanSlashOperatorUntil changed unexpectedly");
        // check that `contractCanSlashOperatorUntil` did not change
        require(slasher.contractCanSlashOperatorUntil(operator, contractAddress) == contractCanSlashOperatorUntilBefore, "contractCanSlashOperatorUntil changed unexpectedly");
    }

    function testRecordFirstStakeUpdate_RevertsWhenPaused() external {
        address operator = address(this);
        address contractAddress = address(this);
        uint32 serveUntil = 0;
        testOptIntoSlashing(operator, contractAddress);

        // pause first stake updates
        cheats.startPrank(pauser);
        slasher.pause(2 ** PAUSED_FIRST_STAKE_UPDATE);
        cheats.stopPrank();

        cheats.startPrank(contractAddress);
        cheats.expectRevert(bytes("Pausable: index is paused"));
        slasher.recordFirstStakeUpdate(operator, serveUntil);
        cheats.stopPrank();
    }

    function testRecordFirstStakeUpdate_RevertsWhenCallerDoesntHaveSlashingPermission() external {
        address operator = address(this);
        address contractAddress = address(this);
        uint32 serveUntil = 0;

        require(!slasher.canSlash(operator, contractAddress), "improper slashing permission has been given");

        cheats.startPrank(contractAddress);
        cheats.expectRevert(bytes("Slasher.onlyRegisteredForService: Operator has not opted into slashing by caller"));
        slasher.recordFirstStakeUpdate(operator, serveUntil);
        cheats.stopPrank();
    }

    function testRecordFirstStakeUpdate_RevertsWhenCallerAlreadyInList() external {
        address operator = address(this);
        address contractAddress = address(this);
        uint32 serveUntil = 0;

        testRecordFirstStakeUpdate(operator, contractAddress, serveUntil);

        cheats.startPrank(contractAddress);
        cheats.expectRevert(bytes("Slasher.recordFirstStakeUpdate: Appending middleware unsuccessful"));
        slasher.recordFirstStakeUpdate(operator, serveUntil);
        cheats.stopPrank();
    }

    function testRecordStakeUpdate(
        address operator,
        address contractAddress,
        uint32 prevServeUntil,
        uint32 updateBlock,
        uint32 serveUntil,
        uint256 insertAfter
    )
        public
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        testRecordFirstStakeUpdate(operator, contractAddress, prevServeUntil);
        _testRecordStakeUpdate(operator, contractAddress, updateBlock, serveUntil, insertAfter);
    }

    // internal function corresponding to the bulk of the logic in `testRecordStakeUpdate`, so we can reuse it elsewhere without calling `testOptIntoSlashing`
    function _testRecordStakeUpdate(
        address operator,
        address contractAddress,
        uint32 updateBlock,
        uint32 serveUntil,
        uint256 insertAfter
    )
        internal
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        // filter out invalid fuzzed inputs. "cannot provide update for future block"
        cheats.assume(updateBlock <= block.number);

        linkedListLengthBefore = slasher.operatorWhitelistedContractsLinkedListSize(operator);
        middlewareDetailsBefore = slasher.whitelistedContractDetails(operator, contractAddress);
        contractCanSlashOperatorUntilBefore = slasher.contractCanSlashOperatorUntil(operator, contractAddress);

        // filter out invalid fuzzed inputs. "can't push a previous update"
        cheats.assume(updateBlock >= middlewareDetailsBefore.latestUpdateBlock);

        // fetch the most recent struct
        middlewareTimesLengthBefore = slasher.middlewareTimesLength(operator);
        ISlasher.MiddlewareTimes memory mostRecentMiddlewareTimesStructBefore = slasher.operatorToMiddlewareTimes(operator, middlewareTimesLengthBefore - 1);

        cheats.startPrank(contractAddress);
        slasher.recordStakeUpdate(operator, updateBlock, serveUntil, insertAfter);
        cheats.stopPrank();

        ISlasher.MiddlewareTimes memory mostRecentMiddlewareTimesStructAfter = slasher.operatorToMiddlewareTimes(operator, slasher.middlewareTimesLength(operator) - 1);

        // check that linked list size remained the same appropriately
        require(slasher.operatorWhitelistedContractsLinkedListSize(operator) == linkedListLengthBefore, "linked list length did increased inappropriately");
        // get the linked list entry for the `contractAddress`
        (nodeExists, prevNode, nextNode) = slasher.operatorWhitelistedContractsLinkedListEntry(operator, contractAddress);
        // verify that the node exists
        require(nodeExists, "node does not exist");

        // if the `serveUntil` time is greater than the previous maximum, then an update must have been pushed and the `latestServeUntil` must be equal to the `serveUntil` input
        if (serveUntil > mostRecentMiddlewareTimesStructBefore.latestServeUntil) {
            require(slasher.middlewareTimesLength(operator) == middlewareTimesLengthBefore + 1, "MiddlewareTimes struct not pushed to array");
            require(mostRecentMiddlewareTimesStructAfter.latestServeUntil == serveUntil, "latestServeUntil not updated correctly");
        } else {
            require(mostRecentMiddlewareTimesStructAfter.latestServeUntil  == mostRecentMiddlewareTimesStructBefore.latestServeUntil, "latestServeUntil updated incorrectly");
        }
        // if this is the first MiddlewareTimes struct in the array, then an update must have been pushed and the `stalestUpdateBlock` *must* be equal to the current block
        if (middlewareTimesLengthBefore == 0) {
            require(slasher.middlewareTimesLength(operator) == middlewareTimesLengthBefore + 1,
                "MiddlewareTimes struct not pushed to array");
            require(mostRecentMiddlewareTimesStructAfter.stalestUpdateBlock == updateBlock,
                "stalestUpdateBlock not updated correctly -- contractAddress is first list entry");
        // otherwise, we check if the `contractAddress` is the head of the list. If it *is*, then prevNode will be _HEAD, and...
        } else if (prevNode == _HEAD) {
            // if nextNode is _HEAD, then the this indicates that `contractAddress` is actually the only list entry
            if (nextNode != _HEAD) {
                // if nextNode is not the only list entry, then `latestUpdateBlock` should update to a more recent time (if applicable!)
                uint32 nextMiddlewaresLeastRecentUpdateBlock = slasher.whitelistedContractDetails(operator, _uintToAddress(nextNode)).latestUpdateBlock;
                uint32 newValue = (nextMiddlewaresLeastRecentUpdateBlock < updateBlock) ? nextMiddlewaresLeastRecentUpdateBlock: uint32(updateBlock);
                require(mostRecentMiddlewareTimesStructAfter.stalestUpdateBlock == newValue,
                    "stalestUpdateBlock not updated correctly -- should have updated to newValue");                
            } else {
                require(mostRecentMiddlewareTimesStructAfter.stalestUpdateBlock == updateBlock,
                    "stalestUpdateBlock not updated correctly -- contractAddress is only list entry");
            }
        }

        middlewareDetailsAfter = slasher.whitelistedContractDetails(operator, contractAddress);
        require(middlewareDetailsAfter.latestUpdateBlock == updateBlock,
            "latestUpdateBlock not updated correctly");
        require(middlewareDetailsAfter.contractCanSlashOperatorUntil == middlewareDetailsBefore.contractCanSlashOperatorUntil,
            "contractCanSlashOperatorUntil changed unexpectedly");
        // check that `contractCanSlashOperatorUntil` did not change
        require(slasher.contractCanSlashOperatorUntil(operator, contractAddress) == contractCanSlashOperatorUntilBefore, "contractCanSlashOperatorUntil changed unexpectedly");
    }

    function testRecordStakeUpdate_MultipleLinkedListEntries(
        address operator,
        address contractAddress,
        uint32 prevServeUntil,
        uint32 updateBlock,
        uint32 serveUntil,
        uint256 insertAfter
    )
        public
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        address _contractAddress = address(this);
        cheats.assume(contractAddress != _contractAddress);
        testRecordFirstStakeUpdate(operator, _contractAddress, prevServeUntil);
        testRecordStakeUpdate(operator, contractAddress, prevServeUntil, updateBlock, serveUntil, insertAfter);
    }

    function testRecordStakeUpdate_RevertsWhenCallerNotAlreadyInList(
        address operator,
        address contractAddress,
        uint32 serveUntil,
        uint256 insertAfter
    )
        public
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        uint32 updateBlock = 0;

        testOptIntoSlashing(operator, contractAddress);

        cheats.expectRevert(bytes("Slasher.recordStakeUpdate: Removing middleware unsuccessful"));
        cheats.startPrank(contractAddress);
        slasher.recordStakeUpdate(operator, updateBlock, serveUntil, insertAfter);
        cheats.stopPrank();
    }

    function testRecordStakeUpdate_RevertsWhenCallerNotAlreadyInList_MultipleLinkedListEntries(
        address operator,
        address contractAddress,
        uint32 prevServeUntil,
        uint32 serveUntil,
        uint256 insertAfter
    )
        public
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        address _contractAddress = address(this);
        uint32 updateBlock = 0;

        cheats.assume(contractAddress != _contractAddress);

        testRecordFirstStakeUpdate(operator, _contractAddress, prevServeUntil);
        testOptIntoSlashing(operator, contractAddress);

        cheats.expectRevert(bytes("Slasher.recordStakeUpdate: Caller is not the list entrant"));
        cheats.startPrank(contractAddress);
        slasher.recordStakeUpdate(operator, updateBlock, serveUntil, insertAfter);
        cheats.stopPrank();
    }

    function testRecordStakeUpdate_RevertsWhenCallerCannotSlash(
        address operator,
        address contractAddress,
        uint32 serveUntil,
        uint256 insertAfter
    )
        public
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        uint32 updateBlock = 0;
        cheats.expectRevert(bytes("Slasher.onlyRegisteredForService: Operator has not opted into slashing by caller"));
        cheats.startPrank(contractAddress);
        slasher.recordStakeUpdate(operator, updateBlock, serveUntil, insertAfter);
        cheats.stopPrank();
    }


    function testRecordStakeUpdate_RevertsWhenUpdateBlockInFuture(
        address operator,
        address contractAddress,
        uint32 prevServeUntil,
        uint32 updateBlock,
        uint32 serveUntil,
        uint256 insertAfter
    )
        public
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        // filter to appropriate fuzzed inputs (appropriate for causing reverts!)
        cheats.assume(updateBlock > block.number);

        testRecordFirstStakeUpdate(operator, contractAddress, prevServeUntil);

        cheats.expectRevert(bytes("Slasher.recordStakeUpdate: cannot provide update for future block"));
        cheats.startPrank(contractAddress);
        slasher.recordStakeUpdate(operator, updateBlock, serveUntil, insertAfter);
        cheats.stopPrank();
    }

    function testRecordLastStakeUpdateAndRevokeSlashingAbility(
        address operator,
        address contractAddress,
        uint32 prevServeUntil,
        uint32 updateBlock,
        uint32 serveUntil,
        uint256 insertAfter
    )
        public
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        // filter out setting the `serveUntil` time to the MAX, since the contract will revert in this instance.
        cheats.assume(serveUntil != MAX_CAN_SLASH_UNTIL);

        testRecordStakeUpdate_MultipleLinkedListEntries(operator, contractAddress, prevServeUntil, updateBlock, serveUntil, insertAfter);

        linkedListLengthBefore = slasher.operatorWhitelistedContractsLinkedListSize(operator);
        middlewareDetailsBefore = slasher.whitelistedContractDetails(operator, contractAddress);

        ISlasher.MiddlewareTimes memory mostRecentMiddlewareTimesStructBefore;
        // fetch the most recent struct, if at least one exists (otherwise leave the struct uninitialized)
        middlewareTimesLengthBefore = slasher.middlewareTimesLength(operator);
        if (middlewareTimesLengthBefore != 0) {
            mostRecentMiddlewareTimesStructBefore = slasher.operatorToMiddlewareTimes(operator, middlewareTimesLengthBefore - 1);
        }

        // get the linked list entry for the `contractAddress`
        (nodeExists, prevNode, nextNode) = slasher.operatorWhitelistedContractsLinkedListEntry(operator, contractAddress);
        require(nodeExists, "node does not exist when it should");

        // perform the last stake update and revoke slashing ability, from the `contractAddress`
        cheats.startPrank(contractAddress);
        slasher.recordLastStakeUpdateAndRevokeSlashingAbility(operator, serveUntil);
        cheats.stopPrank();

        ISlasher.MiddlewareTimes memory mostRecentMiddlewareTimesStructAfter = slasher.operatorToMiddlewareTimes(operator, slasher.middlewareTimesLength(operator) - 1);

        // check that linked list size decrease appropriately
        require(slasher.operatorWhitelistedContractsLinkedListSize(operator) == linkedListLengthBefore - 1, "linked list length did not decrease when it should!");
        // verify that the node no longer exists
        (nodeExists, /*prevNode*/, /*nextNode*/) = slasher.operatorWhitelistedContractsLinkedListEntry(operator, contractAddress);
        require(!nodeExists, "node exists when it should have been deleted");

        // if the `serveUntil` time is greater than the previous maximum, then an update must have been pushed and the `latestServeUntil` must be equal to the `serveUntil` input
        if (serveUntil > mostRecentMiddlewareTimesStructBefore.latestServeUntil) {
            require(slasher.middlewareTimesLength(operator) == middlewareTimesLengthBefore + 1, "MiddlewareTimes struct not pushed to array");
            require(mostRecentMiddlewareTimesStructAfter.latestServeUntil == serveUntil, "latestServeUntil not updated correctly");
        } else {
            require(mostRecentMiddlewareTimesStructAfter.latestServeUntil  == mostRecentMiddlewareTimesStructBefore.latestServeUntil, "latestServeUntil updated incorrectly");
        }
        // if this is the first MiddlewareTimes struct in the array, then an update must have been pushed and the `stalestUpdateBlock` *must* be equal to the current block
        if (middlewareTimesLengthBefore == 0) {
            require(slasher.middlewareTimesLength(operator) == middlewareTimesLengthBefore + 1,
                "MiddlewareTimes struct not pushed to array");
            require(mostRecentMiddlewareTimesStructAfter.stalestUpdateBlock == block.number,
                "stalestUpdateBlock not updated correctly -- contractAddress is first list entry");
        // otherwise, we check if the `contractAddress` is the head of the list. If it *is*, then prevNode will be _HEAD, and...
        } else if (prevNode == _HEAD) {
            // if nextNode is _HEAD, then the this indicates that `contractAddress` is actually the only list entry
            if (nextNode != _HEAD) {
                // if nextNode is not the only list entry, then `latestUpdateBlock` should update to a more recent time (if applicable!)
                uint32 nextMiddlewaresLeastRecentUpdateBlock = slasher.whitelistedContractDetails(operator, _uintToAddress(nextNode)).latestUpdateBlock;
                uint32 newValue = (nextMiddlewaresLeastRecentUpdateBlock < block.number) ? nextMiddlewaresLeastRecentUpdateBlock: uint32(block.number);
                require(mostRecentMiddlewareTimesStructAfter.stalestUpdateBlock == newValue,
                    "stalestUpdateBlock not updated correctly -- should have updated to newValue");                
            } else {
                require(mostRecentMiddlewareTimesStructAfter.stalestUpdateBlock == block.number,
                    "stalestUpdateBlock not updated correctly -- contractAddress is only list entry");
            }
        }

        middlewareDetailsAfter = slasher.whitelistedContractDetails(operator, contractAddress);
        require(middlewareDetailsAfter.latestUpdateBlock == block.number,
            "latestUpdateBlock not updated correctly");
        // check that slashing ability was revoked after `serveUntil`
        require(slasher.contractCanSlashOperatorUntil(operator, contractAddress) == serveUntil, "contractCanSlashOperatorUntil not set correctly");
    }

    function testRecordLastStakeUpdateAndRevokeSlashingAbility_RevertsWhenCallerCannotSlash(
        address operator,
        address contractAddress,
        uint32 serveUntil
    )
        public
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        // filter out setting the `serveUntil` time to the MAX, since the contract will revert in this instance.
        cheats.assume(serveUntil != MAX_CAN_SLASH_UNTIL);

        cheats.expectRevert(bytes("Slasher.onlyRegisteredForService: Operator has not opted into slashing by caller"));
        cheats.startPrank(contractAddress);
        slasher.recordLastStakeUpdateAndRevokeSlashingAbility(operator, serveUntil);
        cheats.stopPrank();
    }

    function testRecordLastStakeUpdateAndRevokeSlashingAbility_RevertsWhenCallerNotAlreadyInList(
        address operator,
        address contractAddress,
        uint32 serveUntil
    )
        public
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        // filter out setting the `serveUntil` time to the MAX, since the contract will revert in this instance.
        cheats.assume(serveUntil != MAX_CAN_SLASH_UNTIL);

        testOptIntoSlashing(operator, contractAddress);

        cheats.expectRevert(bytes("Slasher.recordLastStakeUpdateAndRevokeSlashingAbility: Removing middleware unsuccessful"));
        cheats.startPrank(contractAddress);
        slasher.recordLastStakeUpdateAndRevokeSlashingAbility(operator, serveUntil);
        cheats.stopPrank();
    }

    function testRecordLastStakeUpdateAndRevokeSlashingAbility_RevertsWhenServeUntilInputIsMax(
        address operator,
        address contractAddress,
        uint32 prevServeUntil,
        uint32 updateBlock,
        uint256 insertAfter
    )
        public
        filterFuzzedAddressInputs(operator)
        filterFuzzedAddressInputs(contractAddress)
    {
        uint32 serveUntil = MAX_CAN_SLASH_UNTIL;

        testOptIntoSlashing(operator, contractAddress);

        _testRecordFirstStakeUpdate(operator, contractAddress, prevServeUntil);
        _testRecordStakeUpdate(operator, contractAddress, updateBlock, prevServeUntil, insertAfter);

        // perform the last stake update and revoke slashing ability, from the `contractAddress`
        cheats.startPrank(contractAddress);
        cheats.expectRevert(bytes("Slasher._revokeSlashingAbility: serveUntil time must be limited"));
        slasher.recordLastStakeUpdateAndRevokeSlashingAbility(operator, serveUntil);
        cheats.stopPrank();
    }

    function _addressToUint(address addr) internal pure returns(uint256) {
        return uint256(uint160(addr));
    }

    function _uintToAddress(uint256 x) internal pure returns(address) {
        return address(uint160(x));
    }
}