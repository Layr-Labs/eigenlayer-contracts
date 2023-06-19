// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "../mocks/BLSPubkeyRegistryMock.sol";
import "../mocks/RegistryCoordinatorMock.sol";
import "../mocks/StakeRegistryMock.sol";
import "../mocks/BLSSignatureCheckerMock.sol";

import "../../contracts/middleware/BLSSignatureChecker.sol";

import "forge-std/Test.sol";


contract BLSSignatureCheckerUnitTests is Test {
    Vm cheats = Vm(HEVM_ADDRESS);

    BLSPubkeyRegistryMock public blsPubkeyRegistry;
    StakeRegistryMock public stakeRegistry;
    RegistryCoordinatorMock public coordinator;
    BLSSignatureCheckerWrapper public signatureChecker;

    BLSSignatureChecker.NonSignerStakesAndSignature public defaultNonSignerStakesAndSignature;

    function setUp() public {

        blsPubkeyRegistry = new BLSPubkeyRegistryMock();
        stakeRegistry = new StakeRegistryMock();
        coordinator = new RegistryCoordinatorMock(stakeRegistry, blsPubkeyRegistry);
        signatureChecker = new BLSSignatureCheckerWrapper(coordinator);
    }

    function testMismatchedAPKHash(uint8 quorumNumber, bytes32 apkHash) public {
        bytes memory quorumNumbers = new bytes(1);
        quorumNumbers[0] = bytes1(quorumNumber);

        uint32 referenceBlockNumber = 0;

        BLSSignatureChecker.NonSignerStakesAndSignature memory defaultNonSignerStakesAndSignature = nonSignerStakesAndSignature();

        emit log_named_uint("apk.x", defaultNonSignerStakesAndSignature.quorumApks[0].X);
        cheats.expectRevert(bytes("BLSSignatureChecker.checkSignatures: apkIndex does not match apk"));
        signatureChecker.checkSignatures(bytes32(0), quorumNumbers, referenceBlockNumber, defaultNonSignerStakesAndSignature);

    }



    function nonSignerStakesAndSignature() public returns(BLSSignatureChecker.NonSignerStakesAndSignature memory) {
        BLSSignatureChecker.NonSignerStakesAndSignature memory defaultNonSignerStakesAndSignature;
        defaultNonSignerStakesAndSignature.nonSignerPubkeys =  new BN254.G1Point[](1);
        defaultNonSignerStakesAndSignature.nonSignerPubkeys[0] = BN254.G1Point(uint256(0), uint256(0));

        defaultNonSignerStakesAndSignature.quorumApks =  new BN254.G1Point[](1);
        defaultNonSignerStakesAndSignature.quorumApks[0] = BN254.G1Point(uint256(0), uint256(0));

        defaultNonSignerStakesAndSignature.apkG2 =  BN254.G2Point([uint256(0), uint256(0)], [uint256(0), uint256(0)]);
        defaultNonSignerStakesAndSignature.sigma = BN254.G1Point(uint256(0), uint256(0));
        defaultNonSignerStakesAndSignature.apkIndexes = new uint32[](1);
        defaultNonSignerStakesAndSignature.totalStakeIndexes = new uint32[](1);
        defaultNonSignerStakesAndSignature.nonSignerStakeIndexes = new uint32[][](1);


        return defaultNonSignerStakesAndSignature;
    }

}