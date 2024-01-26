// contract DelegationManagerUnitTests_operatorAVSRegisterationStatus is DelegationManagerUnitTests {
//     // @notice Tests that an avs who calls `updateAVSMetadataURI` will correctly see an `AVSMetadataURIUpdated` event emitted with their input
//     function testFuzz_UpdateAVSMetadataURI(string memory metadataURI) public {
//         // call `updateAVSMetadataURI` and check for event
//         cheats.expectEmit(true, true, true, true, address(delegationManager));
//         cheats.prank(defaultAVS);
//         emit AVSMetadataURIUpdated(defaultAVS, metadataURI);
//         delegationManager.updateAVSMetadataURI(metadataURI);
//     }

//     // @notice Verifies an operator registers successfull to avs and see an `OperatorAVSRegistrationStatusUpdated` event emitted
//     function testFuzz_registerOperatorToAVS(bytes32 salt) public {
//         address operator = cheats.addr(delegationSignerPrivateKey);
//         assertFalse(delegationManager.isOperator(operator), "bad test setup");
//         _registerOperatorWithBaseDetails(operator);

//         cheats.expectEmit(true, true, true, true, address(delegationManager));
//         emit OperatorAVSRegistrationStatusUpdated(operator, defaultAVS, OperatorAVSRegistrationStatus.REGISTERED);

//         uint256 expiry = type(uint256).max;

//         cheats.prank(defaultAVS);
//         ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _getOperatorSignature(
//             delegationSignerPrivateKey,
//             operator,
//             defaultAVS,
//             salt,
//             expiry
//         );

//         delegationManager.registerOperatorToAVS(operator, operatorSignature);
//     }

//     // @notice Verifies an operator registers successfull to avs and see an `OperatorAVSRegistrationStatusUpdated` event emitted
//     function testFuzz_revert_whenOperatorNotRegisteredToEigenLayerYet(bytes32 salt) public {
//         address operator = cheats.addr(delegationSignerPrivateKey);
//         assertFalse(delegationManager.isOperator(operator), "bad test setup");

//         cheats.prank(defaultAVS);
//         uint256 expiry = type(uint256).max;
//         ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _getOperatorSignature(
//             delegationSignerPrivateKey,
//             operator,
//             defaultAVS,
//             salt,
//             expiry
//         );

//         cheats.expectRevert("DelegationManager.registerOperatorToAVS: operator not registered to EigenLayer yet");
//         delegationManager.registerOperatorToAVS(operator, operatorSignature);
//     }

//     // @notice Verifies an operator registers fails when the signature is not from the operator
//     function testFuzz_revert_whenSignatureAddressIsNotOperator(bytes32 salt) public {
//         address operator = cheats.addr(delegationSignerPrivateKey);
//         assertFalse(delegationManager.isOperator(operator), "bad test setup");
//         _registerOperatorWithBaseDetails(operator);

//         uint256 expiry = type(uint256).max;
//         ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _getOperatorSignature(
//             delegationSignerPrivateKey,
//             operator,
//             defaultAVS,
//             salt,
//             expiry
//         );

//         cheats.expectRevert("EIP1271SignatureUtils.checkSignature_EIP1271: signature not from signer");
//         cheats.prank(operator);
//         delegationManager.registerOperatorToAVS(operator, operatorSignature);
//     }

//     // @notice Verifies an operator registers fails when the signature expiry already expires
//     function testFuzz_revert_whenExpiryHasExpired(bytes32 salt, uint256 expiry, ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) public {
//         address operator = cheats.addr(delegationSignerPrivateKey);
//         cheats.assume(operatorSignature.expiry < block.timestamp);

//         cheats.expectRevert("DelegationManager.registerOperatorToAVS: operator signature expired");
//         delegationManager.registerOperatorToAVS(operator, operatorSignature);
//     }

//     // @notice Verifies an operator registers fails when it's already registered to the avs
//     function testFuzz_revert_whenOperatorAlreadyRegisteredToAVS(bytes32 salt) public {
//         address operator = cheats.addr(delegationSignerPrivateKey);
//         assertFalse(delegationManager.isOperator(operator), "bad test setup");
//         _registerOperatorWithBaseDetails(operator);

//         uint256 expiry = type(uint256).max;
//         ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _getOperatorSignature(
//             delegationSignerPrivateKey,
//             operator,
//             defaultAVS,
//             salt,
//             expiry
//         );

//         cheats.startPrank(defaultAVS);
//         delegationManager.registerOperatorToAVS(operator, operatorSignature);

//         cheats.expectRevert("DelegationManager.registerOperatorToAVS: operator already registered");
//         delegationManager.registerOperatorToAVS(operator, operatorSignature);
//         cheats.stopPrank();
//     }
// }