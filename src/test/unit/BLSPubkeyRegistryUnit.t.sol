//SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;


import "forge-std/Test.sol";
import "../../contracts/middleware/BLSPubkeyRegistry.sol";
import "../../contracts/interfaces/IRegistryCoordinator.sol";
import "../mocks/BLSPublicKeyCompendiumMock.sol";
import "../mocks/RegistryCoordinatorMock.sol";


contract BLSPubkeyRegistryUnitTests is Test {
    using BN254 for BN254.G1Point;
    Vm cheats = Vm(HEVM_ADDRESS);

    address defaultOperator = address(4545);
    address defaultOperator2 = address(4546);

    bytes32 internal constant ZERO_PK_HASH = hex"ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5";

    BLSPubkeyRegistry public blsPubkeyRegistry;
    BLSPublicKeyCompendiumMock public pkCompendium;
    RegistryCoordinatorMock public registryCoordinator;

    BN254.G1Point internal defaultPubKey =  BN254.G1Point(18260007818883133054078754218619977578772505796600400998181738095793040006897,3432351341799135763167709827653955074218841517684851694584291831827675065899);

    uint8 internal defaultQuorumNumber = 0;

    function setUp() external {
        registryCoordinator = new RegistryCoordinatorMock();
        pkCompendium = new BLSPublicKeyCompendiumMock();
        blsPubkeyRegistry = new BLSPubkeyRegistry(registryCoordinator, pkCompendium);
    }

    function testConstructorArgs() public view {
        require(blsPubkeyRegistry.registryCoordinator() == registryCoordinator, "registryCoordinator not set correctly");
        require(blsPubkeyRegistry.pubkeyCompendium() == pkCompendium, "pubkeyCompendium not set correctly");
    }

    function testCallRegisterOperatorFromNonCoordinatorAddress(address nonCoordinatorAddress) public {
        cheats.assume(nonCoordinatorAddress != address(registryCoordinator));

        cheats.startPrank(nonCoordinatorAddress);
        cheats.expectRevert(bytes("BLSPubkeyRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"));
        blsPubkeyRegistry.registerOperator(nonCoordinatorAddress, new bytes(0), BN254.G1Point(0, 0));
        cheats.stopPrank();
    }

    function testCallDeregisterOperatorFromNonCoordinatorAddress(address nonCoordinatorAddress) public {
        cheats.assume(nonCoordinatorAddress != address(registryCoordinator));

        cheats.startPrank(nonCoordinatorAddress);
        cheats.expectRevert(bytes("BLSPubkeyRegistry.onlyRegistryCoordinator: caller is not the registry coordinator"));
        blsPubkeyRegistry.deregisterOperator(nonCoordinatorAddress, new bytes(0), BN254.G1Point(0, 0));
        cheats.stopPrank();
    }

    function testOperatorDoesNotOwnPubKeyRegister() public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSPubkeyRegistry.registerOperator: operator does not own pubkey"));
        blsPubkeyRegistry.registerOperator(defaultOperator, new bytes(1), BN254.G1Point(1, 0));
        cheats.stopPrank();
    }

    function testOperatorRegisterZeroPubkey() public {
        cheats.startPrank(address(registryCoordinator));
        cheats.expectRevert(bytes("BLSPubkeyRegistry.registerOperator: cannot register zero pubkey"));
        blsPubkeyRegistry.registerOperator(defaultOperator, new bytes(1), BN254.G1Point(0, 0));
        cheats.stopPrank();
    }

    function testRegisterOperatorBLSPubkey(address operator, bytes32 x) public returns(bytes32){
        
        BN254.G1Point memory pubkey = BN254.hashToG1(x);
        bytes32 pkHash = BN254.hashG1Point(pubkey);

        cheats.startPrank(operator);
        pkCompendium.registerPublicKey(pubkey);
        cheats.stopPrank();

        //register for one quorum
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        
        cheats.startPrank(address(registryCoordinator));
        bytes32 registeredpkHash = blsPubkeyRegistry.registerOperator(operator, quorumNumbers, pubkey);
        cheats.stopPrank();


        require(registeredpkHash == pkHash, "registeredpkHash not set correctly");
        emit log("ehey");

        return pkHash;
    }

    function testQuorumApkUpdates(uint8 quorumNumber1, uint8 quorumNumber2) public {
        cheats.assume(quorumNumber1 != quorumNumber2);

        bytes memory quorumNumbers = new bytes(2);
        quorumNumbers[0] = bytes1(quorumNumber1);
        quorumNumbers[1] = bytes1(quorumNumber2);

        BN254.G1Point[] memory quorumApksBefore = new BN254.G1Point[](quorumNumbers.length);
        for(uint8 i = 0; i < quorumNumbers.length; i++){
            quorumApksBefore[i] = blsPubkeyRegistry.getApkForQuorum(uint8(quorumNumbers[i]));
        }

        cheats.startPrank(defaultOperator);
        pkCompendium.registerPublicKey(defaultPubKey);
        cheats.stopPrank();
        
        cheats.startPrank(address(registryCoordinator));
        blsPubkeyRegistry.registerOperator(defaultOperator, quorumNumbers, defaultPubKey);
        cheats.stopPrank();

        //check quorum apk updates
        for(uint8 i = 0; i < quorumNumbers.length; i++){
            BN254.G1Point memory quorumApkAfter = blsPubkeyRegistry.getApkForQuorum(uint8(quorumNumbers[i]));
            bytes32 temp = BN254.hashG1Point(BN254.plus(quorumApkAfter, BN254.negate(quorumApksBefore[i])));
            require(temp == BN254.hashG1Point(defaultPubKey), "quorum apk not updated correctly");
        }
    }

    function testRegisterWithNegativeQuorumApk(address operator, bytes32 x) external {
        testRegisterOperatorBLSPubkey(defaultOperator, x);

        BN254.G1Point memory quorumApk = blsPubkeyRegistry.getApkForQuorum(defaultQuorumNumber);

        BN254.G1Point memory negatedQuorumApk = BN254.negate(quorumApk);

        //register for one quorum
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        cheats.startPrank(operator);
        pkCompendium.registerPublicKey(negatedQuorumApk);
        cheats.stopPrank();

        cheats.startPrank(address(registryCoordinator));
        blsPubkeyRegistry.registerOperator(operator, quorumNumbers, negatedQuorumApk);
        cheats.stopPrank();

        require(BN254.hashG1Point(blsPubkeyRegistry.getApkForQuorum(defaultQuorumNumber)) == ZERO_PK_HASH, "quorumApk not set correctly");
    }
    
    function testQuorumApkUpdatesDeregistration(uint8 quorumNumber1, uint8 quorumNumber2) external {
        cheats.assume(quorumNumber1 != quorumNumber2);
        bytes memory quorumNumbers = new bytes(2);
        quorumNumbers[0] = bytes1(quorumNumber1);
        quorumNumbers[1] = bytes1(quorumNumber2);

        testQuorumApkUpdates(quorumNumber1, quorumNumber2);

        BN254.G1Point[] memory quorumApksBefore = new BN254.G1Point[](2);
        for(uint8 i = 0; i < quorumNumbers.length; i++){
            quorumApksBefore[i] = blsPubkeyRegistry.getApkForQuorum(uint8(quorumNumbers[i]));
        }

        cheats.startPrank(address(registryCoordinator));
        blsPubkeyRegistry.deregisterOperator(defaultOperator, quorumNumbers, defaultPubKey);
        cheats.stopPrank();

        
        BN254.G1Point memory quorumApkAfter;
        for(uint8 i = 0; i < quorumNumbers.length; i++){
            quorumApkAfter = blsPubkeyRegistry.getApkForQuorum(uint8(quorumNumbers[i]));
            require(BN254.hashG1Point(quorumApksBefore[i].plus(defaultPubKey.negate())) == BN254.hashG1Point(quorumApkAfter), "quorum apk not updated correctly");
        }
    }

    function testDeregisterOperatorWithQuorumApk(bytes32 x1, bytes32 x2) external {
        testRegisterOperatorBLSPubkey(defaultOperator, x1);
        testRegisterOperatorBLSPubkey(defaultOperator2, x2);

        BN254.G1Point memory quorumApksBefore= blsPubkeyRegistry.getApkForQuorum(defaultQuorumNumber);

        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);

        cheats.startPrank(defaultOperator);
        pkCompendium.registerPublicKey(quorumApksBefore);
        cheats.stopPrank();

        cheats.prank(address(registryCoordinator));
        blsPubkeyRegistry.deregisterOperator(defaultOperator, quorumNumbers, quorumApksBefore);

        BN254.G1Point memory pk = blsPubkeyRegistry.getApkForQuorum(defaultQuorumNumber);
        require(pk.X == 0, "quorum apk not set to zero");
        require(pk.Y == 0, "quorum apk not set to zero");
    }

    function testQuorumApkUpdatesAtBlockNumber(uint256 numRegistrants, uint256 blockGap) external{
        cheats.assume(numRegistrants > 0 && numRegistrants <  100);
        cheats.assume(blockGap < 100);

        BN254.G1Point memory quorumApk = BN254.G1Point(0,0);
        bytes24 quorumApkHash;
        for (uint256 i = 0; i < numRegistrants; i++) {
            bytes32 pk = _getRandomPk(i);
            testRegisterOperatorBLSPubkey(defaultOperator, pk);
            quorumApk = quorumApk.plus(BN254.hashToG1(pk));
            quorumApkHash = bytes24(BN254.hashG1Point(quorumApk));
            require(quorumApkHash == blsPubkeyRegistry.getApkHashForQuorumAtBlockNumberFromIndex(defaultQuorumNumber, uint32(block.number + blockGap) , i), "incorrect quorum aok updates");
            cheats.roll(block.number + 100);
            if(_generateRandomNumber(i) % 2 == 0){
               _deregisterOperator(pk);
               quorumApk = quorumApk.plus(BN254.hashToG1(pk).negate());
               quorumApkHash = bytes24(BN254.hashG1Point(quorumApk));
                require(quorumApkHash == blsPubkeyRegistry.getApkHashForQuorumAtBlockNumberFromIndex(defaultQuorumNumber, uint32(block.number + blockGap) , i + 1), "incorrect quorum aok updates");
                cheats.roll(block.number + 100);
                i++;
            }
        }
    }

    function testIncorrectBlockNumberForQuorumApkUpdates(uint256 numRegistrants, uint32 indexToCheck, uint32 wrongBlockNumber) external {
        cheats.assume(numRegistrants > 0 && numRegistrants <  100);
        cheats.assume(indexToCheck < numRegistrants - 1);

        uint256 startingBlockNumber = block.number;

        for (uint256 i = 0; i < numRegistrants; i++) {
            bytes32 pk = _getRandomPk(i);
            testRegisterOperatorBLSPubkey(defaultOperator, pk);
            cheats.roll(block.number + 100);
        }
        if(wrongBlockNumber < startingBlockNumber + indexToCheck*100){
            cheats.expectRevert(bytes("BLSPubkeyRegistry._validateApkHashForQuorumAtBlockNumber: index too recent"));
            blsPubkeyRegistry.getApkHashForQuorumAtBlockNumberFromIndex(defaultQuorumNumber, wrongBlockNumber, indexToCheck);
        } 
         if (wrongBlockNumber >= startingBlockNumber + (indexToCheck+1)*100){
            cheats.expectRevert(bytes("BLSPubkeyRegistry._validateApkHashForQuorumAtBlockNumber: not latest apk update"));
            blsPubkeyRegistry.getApkHashForQuorumAtBlockNumberFromIndex(defaultQuorumNumber, wrongBlockNumber, indexToCheck);
        }
    }

    function _getRandomPk(uint256 seed) internal view returns (bytes32) {
        return keccak256(abi.encodePacked(block.timestamp, seed));
    }

    function _generateRandomNumber(uint256 seed) internal view returns (uint256) {
        uint256 randomNumber = uint256(keccak256(abi.encodePacked(block.timestamp, seed)));
        return (randomNumber % 100) + 1; 
    }

    function _deregisterOperator(bytes32 pk) internal {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(defaultQuorumNumber);
        cheats.startPrank(address(registryCoordinator));
        blsPubkeyRegistry.deregisterOperator(defaultOperator, quorumNumbers, BN254.hashToG1(pk));
        cheats.stopPrank();
    }

}