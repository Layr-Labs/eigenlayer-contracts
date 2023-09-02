
methods {
    //// External Calls
	// external calls to StakeRegistry
    function _.quorumCount() external => DISPATCHER(true);
    function _.getCurrentTotalStakeForQuorum(uint8 quorumNumber) external => DISPATCHER(true);
    function _.getCurrentOperatorStakeForQuorum(bytes32 operatorId, uint8 quorumNumber) external => DISPATCHER(true);
    function _.registerOperator(address, bytes32, bytes) external => DISPATCHER(true);
    function _.deregisterOperator(bytes32, bytes) external => DISPATCHER(true);

	// external calls to Slasher
    function _.contractCanSlashOperatorUntilBlock(address, address) external => DISPATCHER(true);

    // external calls to BLSPubkeyRegistry
    function _.registerOperator(address, bytes, BN254.G1Point) external => DISPATCHER(true);
    function _.deregisterOperator(address, bytes, BN254.G1Point) external => DISPATCHER(true);

	// external calls to ServiceManager
    function _.latestServeUntilBlock() external => DISPATCHER(true);
    function _.recordLastStakeUpdateAndRevokeSlashingAbility(address, uint256) external => DISPATCHER(true);

	// external calls to IndexRegistry
    function _.registerOperator(bytes32, bytes) external => DISPATCHER(true);
    function _.deregisterOperator(bytes32, bytes, bytes32[]) external => DISPATCHER(true);

    // external calls to ERC1271 (can import OpenZeppelin mock implementation)
    // isValidSignature(bytes32 hash, bytes memory signature) returns (bytes4 magicValue) => DISPATCHER(true)
    function _.isValidSignature(bytes32, bytes) external => DISPATCHER(true);

	// external calls to BLSPubkeyCompendium
    function _.pubkeyHashToOperator(bytes32) external => DISPATCHER(true);

    //envfree functions
    function OPERATOR_CHURN_APPROVAL_TYPEHASH() external returns (bytes32) envfree;
    function slasher() external returns (address) envfree;
    function serviceManager() external returns (address) envfree;
    function blsPubkeyRegistry() external returns (address) envfree;
    function stakeRegistry() external returns (address) envfree;
    function indexRegistry() external returns (address) envfree;
    function registries(uint256) external returns (address) envfree;
    function churnApprover() external returns (address) envfree;
    function isChurnApproverSaltUsed(bytes32) external returns (bool) envfree;
    function getOperatorSetParams(uint8 quorumNumber) external returns (IBLSRegistryCoordinatorWithIndices.OperatorSetParam) envfree;
    function getOperator(address operator) external returns (IRegistryCoordinator.Operator) envfree;
    function getOperatorId(address operator) external returns (bytes32) envfree;
    function getOperatorStatus(address operator) external returns (IRegistryCoordinator.OperatorStatus) envfree;
    function getQuorumBitmapIndicesByOperatorIdsAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds)
        external returns (uint32[]) envfree;
    function getQuorumBitmapByOperatorIdAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) external returns (uint192) envfree;
    function getQuorumBitmapUpdateByOperatorIdByIndex(bytes32 operatorId, uint256 index)
        external returns (IRegistryCoordinator.QuorumBitmapUpdate) envfree;
    function getCurrentQuorumBitmapByOperatorId(bytes32 operatorId) external returns (uint192) envfree;
    function getQuorumBitmapUpdateByOperatorIdLength(bytes32 operatorId) external returns (uint256) envfree;
    function numRegistries() external returns (uint256) envfree;
    function calculateOperatorChurnApprovalDigestHash(
        bytes32 registeringOperatorId,
        bytes quorumNumbers,
        IBLSRegistryCoordinatorWithIndices.OperatorKickParam[] operatorKickParams,
        bytes32 salt,
        uint256 expiry
    ) external returns (bytes32) envfree;
}

// TODOs: add properties in English + rules in CVL

// If my Operator status is REGISTERED ⇔ my quorum bitmap MUST BE nonzero
invariant registeredOperatorsHaveNonzeroBitmaps(address operator)
    getOperatorStatus(operator) == IRegistryCoordinator.OperatorStatus.REGISTERED <=>
        getCurrentQuorumBitmapByOperatorId(getOperatorId(operator)) != 0;

// if two operators have different addresses, then they have different IDs
// excludes the case in which the operator is not registered, since then they can both have ID zero (the default)
// TODO: seems like this needs to be more carefully defined, might need to assume something about existing IDs first?
invariant operatorIdIsUnique(address operator1, address operator2)
    operator1 != operator2 =>
        (getOperatorStatus(operator1) == IRegistryCoordinator.OperatorStatus.REGISTERED => getOperatorId(operator1) != getOperatorId(operator2));

definition methodCanModifyBitmap(method f) returns bool =
    f.selector == sig:registerOperatorWithCoordinator(bytes, bytes).selector
    || f.selector == sig:registerOperatorWithCoordinator(bytes, BN254.G1Point, string).selector
    || f.selector == sig:registerOperatorWithCoordinator(
        bytes,
        BN254.G1Point,
        string,
        IBLSRegistryCoordinatorWithIndices.OperatorKickParam[],
        ISignatureUtils.SignatureWithSaltAndExpiry
    ).selector
    || f.selector == sig:deregisterOperatorWithCoordinator(bytes, bytes).selector
    || f.selector == sig:deregisterOperatorWithCoordinator(bytes, BN254.G1Point, bytes32[]).selector;

definition methodCanAddToBitmap(method f) returns bool =
    f.selector == sig:registerOperatorWithCoordinator(bytes, bytes).selector
    || f.selector == sig:registerOperatorWithCoordinator(bytes, BN254.G1Point, string).selector
    || f.selector == sig:registerOperatorWithCoordinator(
        bytes,
        BN254.G1Point,
        string,
        IBLSRegistryCoordinatorWithIndices.OperatorKickParam[],
        ISignatureUtils.SignatureWithSaltAndExpiry
    ).selector;

// `registerOperatorWithCoordinator` with kick params also meets this definition due to the 'churn' mechanism
definition methodCanRemoveFromBitmap(method f) returns bool =
    f.selector == sig:registerOperatorWithCoordinator(
        bytes,
        BN254.G1Point,
        string,
        IBLSRegistryCoordinatorWithIndices.OperatorKickParam[],
        ISignatureUtils.SignatureWithSaltAndExpiry
    ).selector
    || f.selector == sig:deregisterOperatorWithCoordinator(bytes, bytes).selector
    || f.selector == sig:deregisterOperatorWithCoordinator(bytes, BN254.G1Point, bytes32[]).selector;

/* TODD: this is a Work In Progress
rule canOnlyModifyBitmapWithSpecificFunctions(address operator) {
    requireInvariant registeredOperatorsHaveNonzeroBitmaps(operator);
    uint256 bitmapBefore = getCurrentQuorumBitmapByOperatorId(getOperatorId(operator));
    // prepare to perform arbitrary function call
    method f;
    env e;
    // TODO: need to ensure that if the function can modify the bitmap, then we are using the operator as an input
    if (!methodCanModifyBitmap(f)) {
        // perform arbitrary function call
        calldataarg arg;
        f(e, arg);
        uint256 bitmapAfter = getCurrentQuorumBitmapByOperatorId(getOperatorId(operator));
        assert(bitmapAfter == bitmapBefore);
    } else if (
        f.selector == sig:registerOperatorWithCoordinator(bytes, bytes).selector
    ) {
        if (e.msg.sender != operator) {
            uint256 bitmapAfter = getCurrentQuorumBitmapByOperatorId(getOperatorId(operator));
            assert(bitmapAfter == bitmapBefore);
        }
    }

        // if method did not remove from bitmap, it must have added
        if (bitmapAfter & bitmapBefore == bitmapBefore) {
            assert(methodCanAddToBitmap(f));
        } else {
            assert(methodCanRemoveFromBitmap(f));
        }
    }
}
*/