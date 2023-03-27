// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../../contracts/pods/BeaconChainOracle.sol";

import "forge-std/Test.sol";

contract BeaconChainOracleUnitTests is Test {

    Vm cheats = Vm(HEVM_ADDRESS);

    BeaconChainOracle public beaconChainOracle;

    address public initialBeaconChainOwner = address(this);
    uint256 public initialBeaconChainOracleThreshold = 2;
    uint256 public minThreshold;

    mapping(address => bool) public addressIsExcludedFromFuzzedInputs;

    // static values reused across several tests
    uint256 numberPotentialOracleSigners = 16;
    address[] public potentialOracleSigners;
    uint64 public blockNumberToVoteFor = 5151;
    bytes32 public stateRootToVoteFor = bytes32(uint256(987));

    modifier filterFuzzedAddressInputs(address fuzzedAddress) {
        cheats.assume(!addressIsExcludedFromFuzzedInputs[fuzzedAddress]);
        _;
    }

    function setUp() external {
        address[] memory initialOracleSignersArray = new address[](0);
        beaconChainOracle = new BeaconChainOracle(initialBeaconChainOwner, initialBeaconChainOracleThreshold, initialOracleSignersArray);
        minThreshold = beaconChainOracle.MINIMUM_THRESHOLD();

        // set up array for use in testing
        for (uint256 i = 0; i < numberPotentialOracleSigners; ++i) {
            potentialOracleSigners.push(address(uint160(777 + i)));
        }
    }

    function testConstructor_RevertsOnThresholdTooLow() external {
         address[] memory initialOracleSignersArray = new address[](0);
       // check that deployment fails when trying to set threshold below `MINIMUM_THRESHOLD`
        cheats.expectRevert(bytes("BeaconChainOracle._setThreshold: cannot set threshold below MINIMUM_THRESHOLD"));
        new BeaconChainOracle(initialBeaconChainOwner, minThreshold - 1, initialOracleSignersArray);

        // check that deployment succeeds when trying to set threshold *at* (i.e. equal to) `MINIMUM_THRESHOLD`
        beaconChainOracle = new BeaconChainOracle(initialBeaconChainOwner, minThreshold, initialOracleSignersArray);
    }

    function testSetThreshold(uint256 newThreshold) public {
        // filter out disallowed inputs
        cheats.assume(newThreshold >= minThreshold);

        cheats.startPrank(beaconChainOracle.owner());
        beaconChainOracle.setThreshold(newThreshold);
        cheats.stopPrank();

        assertEq(newThreshold, beaconChainOracle.threshold());
    }

    function testSetThreshold_RevertsOnThresholdTooLow() external {
        cheats.startPrank(beaconChainOracle.owner());
        cheats.expectRevert(bytes("BeaconChainOracle._setThreshold: cannot set threshold below MINIMUM_THRESHOLD"));
        beaconChainOracle.setThreshold(minThreshold - 1);
        cheats.stopPrank();

        // make sure it works *at* (i.e. equal to) the threshold
        testSetThreshold(minThreshold);
    }

    function testSetThreshold_RevertsOnCallingFromNotOwner(address notOwner) external {
        cheats.assume(notOwner != beaconChainOracle.owner());

        cheats.startPrank(notOwner);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        beaconChainOracle.setThreshold(minThreshold);
        cheats.stopPrank();
    }

    function testAddOracleSigner(address signerToAdd) public {
        uint256 totalSignersBefore = beaconChainOracle.totalOracleSigners();
        bool alreadySigner = beaconChainOracle.isOracleSigner(signerToAdd);

        address[] memory signerArray = new address[](1);
        signerArray[0] = signerToAdd;
        cheats.startPrank(beaconChainOracle.owner());
        beaconChainOracle.addOracleSigners(signerArray);
        cheats.stopPrank();

        uint256 totalSignersAfter = beaconChainOracle.totalOracleSigners();
        require(beaconChainOracle.isOracleSigner(signerToAdd), "signer not added");
        if (alreadySigner) {
            require(totalSignersAfter == totalSignersBefore, "totalSigners incremented incorrectly");
        } else {
            require(totalSignersAfter == totalSignersBefore + 1, "totalSigners did not increment correctly");
        }
    }

    function testAddOracleSigners(uint8 amountSignersToAdd) public {
        cheats.assume(amountSignersToAdd <= numberPotentialOracleSigners);
        uint256 totalSignersBefore = beaconChainOracle.totalOracleSigners();

        // copy array to memory
        address[] memory signerArray = new address[](amountSignersToAdd);
        for (uint256 i = 0; i < amountSignersToAdd; ++i) {
            signerArray[i] = potentialOracleSigners[i];
        }

        cheats.startPrank(beaconChainOracle.owner());
        beaconChainOracle.addOracleSigners(signerArray);
        cheats.stopPrank();

        // check post conditions
        uint256 totalSignersAfter = beaconChainOracle.totalOracleSigners();
        for (uint256 i = 0; i < amountSignersToAdd; ++i) {
            require(beaconChainOracle.isOracleSigner(signerArray[i]), "signer not added");
        }
        require(totalSignersAfter == totalSignersBefore + amountSignersToAdd, "totalSigners did not increment correctly");
    }

    function testAddOracleSigners_SignerAlreadyInSet() external {
        address oracleSigner = potentialOracleSigners[0];
        address[] memory signerArray = new address[](1);
        signerArray[0] = oracleSigner;
        testAddOracleSigner(oracleSigner);

        cheats.startPrank(beaconChainOracle.owner());
        beaconChainOracle.addOracleSigners(signerArray);
        cheats.stopPrank();

        require(beaconChainOracle.isOracleSigner(oracleSigner), "signer improperly removed");
    }

    function testAddOracleSigners_RevertsOnCallingFromNotOwner(address notOwner) external {
        cheats.assume(notOwner != beaconChainOracle.owner());
        address oracleSigner = potentialOracleSigners[0];
        address[] memory signerArray = new address[](1);
        signerArray[0] = oracleSigner;

        cheats.startPrank(notOwner);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        beaconChainOracle.addOracleSigners(signerArray);
        cheats.stopPrank();

        require(!beaconChainOracle.isOracleSigner(oracleSigner), "signer improperly added");
    }

    function testRemoveOracleSigner(address signerToRemove) public {
        uint256 totalSignersBefore = beaconChainOracle.totalOracleSigners();
        bool alreadySigner = beaconChainOracle.isOracleSigner(signerToRemove);

        address[] memory signerArray = new address[](1);
        signerArray[0] = signerToRemove;
        cheats.startPrank(beaconChainOracle.owner());
        beaconChainOracle.removeOracleSigners(signerArray);
        cheats.stopPrank();

        uint256 totalSignersAfter = beaconChainOracle.totalOracleSigners();
        require(!beaconChainOracle.isOracleSigner(signerToRemove), "signer not removed");
        if (alreadySigner) {
            require(totalSignersAfter == totalSignersBefore - 1, "totalSigners did not decrement correctly");
        } else {
            require(totalSignersAfter == totalSignersBefore, "totalSigners decremented incorrectly");
        }
    }

    function testRemoveOracleSigners(uint8 amountSignersToAdd, uint8 amountSignersToRemove) external {
        cheats.assume(amountSignersToAdd <= numberPotentialOracleSigners);
        cheats.assume(amountSignersToRemove <= numberPotentialOracleSigners);
        testAddOracleSigners(amountSignersToAdd);

        uint256 totalSignersBefore = beaconChainOracle.totalOracleSigners();

        // copy array to memory
        address[] memory signerArray = new address[](amountSignersToRemove);
        for (uint256 i = 0; i < amountSignersToRemove; ++i) {
            signerArray[i] = potentialOracleSigners[i];
        }

        cheats.startPrank(beaconChainOracle.owner());
        beaconChainOracle.removeOracleSigners(signerArray);
        cheats.stopPrank();

        // check post conditions
        uint256 totalSignersAfter = beaconChainOracle.totalOracleSigners();
        for (uint256 i = 0; i < amountSignersToRemove; ++i) {
            require(!beaconChainOracle.isOracleSigner(signerArray[i]), "signer not removed");
        }
        uint256 amountThatShouldHaveBeenRemoved = amountSignersToRemove > amountSignersToAdd ? amountSignersToAdd : amountSignersToRemove;
        require(totalSignersAfter + amountThatShouldHaveBeenRemoved == totalSignersBefore, "totalSigners did not decrement correctly");
    }

    function testRemoveOracleSigners_SignerAlreadyNotInSet() external {
        address oracleSigner = potentialOracleSigners[0];
        address[] memory signerArray = new address[](1);
        signerArray[0] = oracleSigner;

        cheats.startPrank(beaconChainOracle.owner());
        beaconChainOracle.removeOracleSigners(signerArray);
        cheats.stopPrank();

        require(!beaconChainOracle.isOracleSigner(oracleSigner), "signer improperly added");
    }

    function testRemoveOracleSigners_RevertsOnCallingFromNotOwner(address notOwner) external {
        cheats.assume(notOwner != beaconChainOracle.owner());
        address oracleSigner = potentialOracleSigners[0];
        address[] memory signerArray = new address[](1);
        signerArray[0] = oracleSigner;

        cheats.startPrank(notOwner);
        cheats.expectRevert(bytes("Ownable: caller is not the owner"));
        beaconChainOracle.removeOracleSigners(signerArray);
        cheats.stopPrank();

        require(!beaconChainOracle.isOracleSigner(oracleSigner), "signer improperly added");
    }

    function testVoteForBeaconChainStateRoot(address oracleSigner, uint64 _blockNumber, bytes32 _stateRoot) public {
        uint256 votesBefore = beaconChainOracle.stateRootVotes(_blockNumber, _stateRoot);

        testAddOracleSigner(oracleSigner);
        cheats.startPrank(oracleSigner);
        beaconChainOracle.voteForBeaconChainStateRoot(_blockNumber, _stateRoot);
        cheats.stopPrank();

        uint256 votesAfter = beaconChainOracle.stateRootVotes(_blockNumber, _stateRoot);
        require(votesAfter == votesBefore + 1, "votesAfter != votesBefore + 1");
        require(beaconChainOracle.hasVoted(_blockNumber, oracleSigner), "vote not recorded as being cast");
        if (votesAfter == beaconChainOracle.threshold()) {
            assertEq(beaconChainOracle.beaconStateRootAtBlockNumber(_blockNumber), _stateRoot, "state root not confirmed when it should be");
        } else {
            require(beaconChainOracle.beaconStateRootAtBlockNumber(_blockNumber) == bytes32(0), "state root improperly confirmed");
        }
    }

    function testVoteForBeaconChainStateRoot_VoteDoesNotCauseConfirmation() public {
        address _oracleSigner = potentialOracleSigners[0];
        testVoteForBeaconChainStateRoot(_oracleSigner, blockNumberToVoteFor, stateRootToVoteFor);
    }

    function testVoteForBeaconChainStateRoot_VoteCausesConfirmation(uint64 _blockNumber, bytes32 _stateRoot) public {
        uint64 latestConfirmedOracleBlockNumberBefore = beaconChainOracle.latestConfirmedOracleBlockNumber();

        uint256 votesBefore = beaconChainOracle.stateRootVotes(_blockNumber, _stateRoot);
        require(votesBefore == 0, "something is wrong, state root should have zero votes before voting");

        for (uint256 i = 0; i < beaconChainOracle.threshold(); ++i) {
            testVoteForBeaconChainStateRoot(potentialOracleSigners[i], _blockNumber, _stateRoot);
        }

        assertEq(beaconChainOracle.beaconStateRootAtBlockNumber(_blockNumber), _stateRoot, "state root not confirmed when it should be");
        assertEq(beaconChainOracle.threshold(), beaconChainOracle.stateRootVotes(_blockNumber, _stateRoot), "state root confirmed with incorrect votes");

        if (_blockNumber > latestConfirmedOracleBlockNumberBefore) {
            assertEq(_blockNumber, beaconChainOracle.latestConfirmedOracleBlockNumber(), "latestConfirmedOracleBlockNumber did not update appropriately");
        } else {
            assertEq(latestConfirmedOracleBlockNumberBefore, beaconChainOracle.latestConfirmedOracleBlockNumber(), "latestConfirmedOracleBlockNumber updated inappropriately");
        }
    }

    function testVoteForBeaconChainStateRoot_VoteCausesConfirmation_latestOracleBlockNumberDoesNotIncrease() external {
        testVoteForBeaconChainStateRoot_VoteCausesConfirmation(blockNumberToVoteFor + 1, stateRootToVoteFor);
        uint64 latestConfirmedOracleBlockNumberBefore = beaconChainOracle.latestConfirmedOracleBlockNumber();
        testVoteForBeaconChainStateRoot_VoteCausesConfirmation(blockNumberToVoteFor, stateRootToVoteFor);
        assertEq(latestConfirmedOracleBlockNumberBefore, beaconChainOracle.latestConfirmedOracleBlockNumber(), "latestConfirmedOracleBlockNumber updated inappropriately");
    }

    function testVoteForBeaconChainStateRoot_RevertsWhenCallerHasVoted() external {
        address _oracleSigner = potentialOracleSigners[0];
        testVoteForBeaconChainStateRoot(_oracleSigner, blockNumberToVoteFor, stateRootToVoteFor);

        cheats.startPrank(_oracleSigner);
        cheats.expectRevert(bytes("BeaconChainOracle.voteForBeaconChainStateRoot: Signer has already voted"));
        beaconChainOracle.voteForBeaconChainStateRoot(blockNumberToVoteFor, stateRootToVoteFor);
        cheats.stopPrank();
    }

    function testVoteForBeaconChainStateRoot_RevertsWhenStateRootAlreadyConfirmed() external {
        address _oracleSigner = potentialOracleSigners[potentialOracleSigners.length - 1];
        testAddOracleSigner(_oracleSigner);
        testVoteForBeaconChainStateRoot_VoteCausesConfirmation(blockNumberToVoteFor, stateRootToVoteFor);

        cheats.startPrank(_oracleSigner);
        cheats.expectRevert(bytes("BeaconChainOracle.voteForBeaconChainStateRoot: State root already confirmed"));
        beaconChainOracle.voteForBeaconChainStateRoot(blockNumberToVoteFor, stateRootToVoteFor);
        cheats.stopPrank();
    }

    function testVoteForBeaconChainStateRoot_RevertsWhenCallingFromNotOracleSigner(address notOracleSigner) external {
        cheats.startPrank(notOracleSigner);
        cheats.expectRevert(bytes("BeaconChainOracle.onlyOracleSigner: Not an oracle signer"));
        beaconChainOracle.voteForBeaconChainStateRoot(blockNumberToVoteFor, stateRootToVoteFor);
        cheats.stopPrank();
    }
}