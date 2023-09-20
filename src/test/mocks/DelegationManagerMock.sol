// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "forge-std/Test.sol";
import "../../contracts/interfaces/IDelegationManager.sol";


contract DelegationManagerMock is IDelegationManager, Test {
    mapping(address => bool) public isOperator;

    function setIsOperator(address operator, bool _isOperatorReturnValue) external {
        isOperator[operator] = _isOperatorReturnValue;
    }

    mapping(address => address) public delegatedTo;

    function registerAsOperator(OperatorDetails calldata /*registeringOperatorDetails*/, string calldata /*metadataURI*/) external pure {}
    
    function updateOperatorMetadataURI(string calldata /*metadataURI*/) external pure {}

    function delegateTo(address operator, SignatureWithExpiry memory /*approverSignatureAndExpiry*/, bytes32 /*approverSalt*/) external {
        delegatedTo[msg.sender] = operator;
    }

    function modifyOperatorDetails(OperatorDetails calldata /*newOperatorDetails*/) external pure {}

    function delegateToBySignature(
        address /*staker*/,
        address /*operator*/,
        SignatureWithExpiry memory /*stakerSignatureAndExpiry*/,
        SignatureWithExpiry memory /*approverSignatureAndExpiry*/,
        bytes32 /*approverSalt*/
    ) external pure {}

    function undelegate(address staker) external returns (bytes32 queuedWithdrawal) {
        delegatedTo[staker] = address(0);
    }

    function increaseDelegatedShares(address /*staker*/, IStrategy /*strategy*/, uint256 /*shares*/) external pure {}

    function decreaseDelegatedShares(address /*staker*/, IStrategy[] calldata /*strategies*/, uint256[] calldata /*shares*/, bool /*undelegateIfPossible*/) external pure {}

    function operatorDetails(address operator) external pure returns (OperatorDetails memory) {
        OperatorDetails memory returnValue = OperatorDetails({
            earningsReceiver: operator,
            delegationApprover: operator,
            stakerOptOutWindowBlocks: 0
        });
        return returnValue;
    }

    function earningsReceiver(address operator) external pure returns (address) {
        return operator;
    }

    function delegationApprover(address operator) external pure returns (address) {
        return operator;
    }

    function stakerOptOutWindowBlocks(address /*operator*/) external pure returns (uint256) {
        return 0;
    }

    function operatorShares(address /*operator*/, IStrategy /*strategy*/) external pure returns (uint256) {}

    function isDelegated(address staker) external view returns (bool) {
        return (delegatedTo[staker] != address(0));
    }

    function isNotDelegated(address /*staker*/) external pure returns (bool) {}

    // function isOperator(address /*operator*/) external pure returns (bool) {}

    function stakerNonce(address /*staker*/) external pure returns (uint256) {}

    function delegationApproverSaltIsSpent(address /*delegationApprover*/, bytes32 /*salt*/) external pure returns (bool) {}

    function calculateCurrentStakerDelegationDigestHash(address /*staker*/, address /*operator*/, uint256 /*expiry*/) external view returns (bytes32) {}

    function calculateStakerDelegationDigestHash(address /*staker*/, uint256 /*stakerNonce*/, address /*operator*/, uint256 /*expiry*/) external view returns (bytes32) {}

    function calculateDelegationApprovalDigestHash(
        address /*staker*/,
        address /*operator*/,
        address /*_delegationApprover*/,
        bytes32 /*approverSalt*/,
        uint256 /*expiry*/
    ) external view returns (bytes32) {}

    function calculateStakerDigestHash(address /*staker*/, address /*operator*/, uint256 /*expiry*/)
        external pure returns (bytes32 stakerDigestHash) {}

    function calculateApproverDigestHash(address /*staker*/, address /*operator*/, uint256 /*expiry*/) external pure returns (bytes32 approverDigestHash) {}

    function DOMAIN_TYPEHASH() external view returns (bytes32) {}

    function STAKER_DELEGATION_TYPEHASH() external view returns (bytes32) {}

    function DELEGATION_APPROVAL_TYPEHASH() external view returns (bytes32) {}

    function domainSeparator() external view returns (bytes32) {}

    /** 
     * @notice Returns 'true' if the staker can undelegate or  if the staker is already undelegated, and 'false' otherwise
     * @dev A staker can only undelegate if they have no "active" shares in EigenLayer and are not themselves an operator
     */
    mapping(address => bool) public stakerCanUndelegate;

    function setStakerCanUndelegate(address staker, bool valueToSet) external {
        stakerCanUndelegate[staker] = valueToSet;
    }

    /**
     * @notice Returns 'true' if `staker` has "active" shares in EigenLayer (i.e. the staker has shares which are currently in the StrategyManager
     * or in the EigenPodManager + not in "undelegation limbo"), and returns 'false' otherwise.
     */
    mapping(address => bool) public stakerHasActiveShares;

    function setStakerHasActiveShares(address staker, bool valueToSet) external {
        stakerHasActiveShares[staker] = valueToSet;
    }
}