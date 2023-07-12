// // SPDX-License-Identifier: BUSL-1.1
// pragma solidity =0.8.12;

// import "./EigenLayerTestHelper.t.sol";

// import "@openzeppelin/contracts/utils/math/Math.sol";

// import "../contracts/libraries/BytesLib.sol";

// import "./mocks/MiddlewareVoteWeigherMock.sol";
// import "./mocks/ServiceManagerMock.sol";
// import "./mocks/PublicKeyCompendiumMock.sol";
// import "./mocks/StrategyManagerMock.sol";

// import "../../src/contracts/middleware/BLSRegistry.sol";
// import "../../src/contracts/middleware/BLSPublicKeyCompendium.sol";



// contract RegistrationTests is EigenLayerTestHelper {

//     BLSRegistry public dlRegImplementation;
//     BLSPublicKeyCompendiumMock public pubkeyCompendium;

//     BLSPublicKeyCompendiumMock public pubkeyCompendiumImplementation;
//     BLSRegistry public dlReg;
//     ProxyAdmin public dataLayrProxyAdmin;

//     ServiceManagerMock public dlsm;
//     StrategyManagerMock public strategyManagerMock;

//     function setUp() public virtual override {
//         EigenLayerDeployer.setUp();
//         initializeMiddlewares();
//     }


//     function initializeMiddlewares() public {
//         dataLayrProxyAdmin = new ProxyAdmin();
//         fuzzedAddressMapping[address(dataLayrProxyAdmin)] = true;

//         pubkeyCompendium = new BLSPublicKeyCompendiumMock();

//         strategyManagerMock = new StrategyManagerMock();
//         strategyManagerMock.setAddresses(delegation, eigenPodManager, slasher);

//         dlsm = new ServiceManagerMock(slasher);

//         dlReg = BLSRegistry(
//             address(new TransparentUpgradeableProxy(address(emptyContract), address(dataLayrProxyAdmin), ""))
//         );

//         dlRegImplementation = new BLSRegistry(
//                 strategyManagerMock,
//                 dlsm,
//                 pubkeyCompendium
//             );

//         uint256[] memory _quorumBips = new uint256[](2);
//         // split 60% ETH quorum, 40% EIGEN quorum
//         _quorumBips[0] = 6000;
//         _quorumBips[1] = 4000;

//         VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] memory ethStratsAndMultipliers =
//                 new VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[](1);
//             ethStratsAndMultipliers[0].strategy = wethStrat;
//             ethStratsAndMultipliers[0].multiplier = 1e18;
//         VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[] memory eigenStratsAndMultipliers =
//                 new VoteWeigherBaseStorage.StrategyAndWeightingMultiplier[](1);
//             eigenStratsAndMultipliers[0].strategy = eigenStrat;
//             eigenStratsAndMultipliers[0].multiplier = 1e18;
        
//         dataLayrProxyAdmin.upgradeAndCall(
//                 TransparentUpgradeableProxy(payable(address(dlReg))),
//                 address(dlRegImplementation),
//                 abi.encodeWithSelector(BLSRegistry.initialize.selector, address(this), false, _quorumBips, ethStratsAndMultipliers, eigenStratsAndMultipliers)
//             );

//     }


//     function testRegisterOperator(address operator, uint32 operatorIndex, string calldata socket) public fuzzedAddress(operator) {
//         cheats.assume(operatorIndex < 15);
//         BN254.G1Point memory pk = getOperatorPubkeyG1(operatorIndex);

//         //register as both ETH and EIGEN operator
//         uint256 wethToDeposit = 1e18;
//         uint256 eigenToDeposit = 1e18;
//         _testDepositWeth(operator, wethToDeposit);
//         _testDepositEigen(operator, eigenToDeposit);
//         _testRegisterAsOperator(operator, IDelegationTerms(operator));
        
//         cheats.startPrank(operator);
//         slasher.optIntoSlashing(address(dlsm));
//         pubkeyCompendium.registerPublicKey(pk);
//         dlReg.registerOperator(1, pk, socket);
//         cheats.stopPrank();

//         bytes32 pubkeyHash = BN254.hashG1Point(pk);
        
//         (uint32 toBlockNumber, uint32 index) = dlReg.pubkeyHashToIndexHistory(pubkeyHash,0);

//         assertTrue(toBlockNumber == 0, "block number set when it shouldn't be");
//         assertTrue(index == 0, "index has been set incorrectly");
//         assertTrue(dlReg.operatorList(0) == operator, "incorrect operator added");
//     }

//     function testDeregisterOperator(address operator, uint32 operatorIndex, string calldata socket) public fuzzedAddress(operator) {
//         cheats.assume(operatorIndex < 15);
//         BN254.G1Point memory pk = getOperatorPubkeyG1(operatorIndex);

//         testRegisterOperator(operator, operatorIndex, socket);
//         cheats.startPrank(operator);
//         dlReg.deregisterOperator(pk, 0);
//         cheats.stopPrank();

//         bytes32 pubkeyHash = BN254.hashG1Point(pk);
//         (uint32 toBlockNumber, /*uint32 index*/) = dlReg.pubkeyHashToIndexHistory(pubkeyHash,0);
//         assertTrue(toBlockNumber == block.number, "toBlockNumber has been set incorrectly");
//     }


// }