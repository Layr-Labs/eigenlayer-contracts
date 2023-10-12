// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "./FFIBase.sol";
import "./util/BLSSigCheckerExperimental.sol";

contract OperatorCapAnalysisFFI is FFIBase {
    using BN254 for BN254.G1Point;

    BLSSigCheckerExperimental blsSignatureChecker;

    function testOneQuorumLinearNonSigningCost() public { 
        for(uint64 i = 0; i < 10; i++) {
            _deployMockEigenLayerAndAVS();
            defaultMaxOperatorCount = type(uint32).max;
            blsSignatureChecker = new BLSSigCheckerExperimental(registryCoordinator);
            _compareScalarMuls(i, 10, i, 1, 1);
        }        
    }

    function _compareScalarMuls(
        uint64 pseudoRandomNumber,
        uint64 numOperators, 
        uint64 numNonSigners, 
        uint64 numQuorums,
        uint256 setQuorumBitmap
    ) internal {
        _setNonSignerPrivKeys(numNonSigners, pseudoRandomNumber);

        (
            bytes32 msgHash, 
            bytes memory quorumNumbers, 
            uint32 referenceBlockNumber, 
            BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
        ) = _getRandomNonSignerStakeAndSignatures(
            pseudoRandomNumber, 
            numOperators, 
            numNonSigners, 
            numQuorums,
            setQuorumBitmap
        );
        
        uint256 gasBeforeReg = gasleft();
        blsSignatureChecker.checkSignatures_reg(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
        uint256 gasAfterReg = gasleft();
        uint256 regCost = gasBeforeReg - gasAfterReg;

        uint256 gasBeforeTiny = gasleft();
        blsSignatureChecker.checkSignatures_tiny(
            msgHash, 
            quorumNumbers,
            referenceBlockNumber, 
            nonSignerStakesAndSignature
        );
        uint256 gasAfterTiny = gasleft();
        uint256 tinyCost = gasBeforeTiny - gasAfterTiny;

        emit log_named_uint("Operators", numOperators);
        emit log_named_uint("NonSigners", numNonSigners);
        emit log_named_uint("Quorums", numQuorums);
        emit log_named_uint("scalar_mul", regCost);
        emit log_named_uint("scalar_mul_tiny", tinyCost);
        
        if(tinyCost < regCost){
            emit log_named_uint("-", regCost - tinyCost);
        } else {
            emit log_named_uint("+", tinyCost - regCost);
        }
    }

}