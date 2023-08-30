
methods {
    //// External Calls
	// external calls to StakeRegistry
    function _.quorumCount() external => DISPATCHER(true);
    function _.getCurrentTotalStakeForQuorum(uint8) external => DISPATCHER(true);
    function _.getCurrentOperatorStakeForQuorum(bytes32, uint8) external => DISPATCHER(true);
    function _.registerOperator(address, bytes32, bytes) external => DISPATCHER(true);
    function _.deregisterOperator(bytes32, bytes) external => DISPATCHER(true);

	// external calls to Slasher
    function _.contractCanSlashOperatorUntilBlock(address, address) external => DISPATCHER(true);

    // external calls to BLSPubkeyRegistry
    function _.registerOperator(address, bytes, BN254.G1Point) external => DISPATCHER(true);
    function _.deregisterOperator(address, bytes) external => DISPATCHER(true);

	// external calls to IndexRegistry
    function _.latestServeUntilBlock() external => DISPATCHER(true);
    function _.recordLastStakeUpdateAndRevokeSlashingAbility(address, uint256) external => DISPATCHER(true);

	// external calls to ServiceManager
    function _.registerOperator(bytes32, bytes) external => DISPATCHER(true);
    function _.deregisterOperator(address, bytes, bytes32[]) external => DISPATCHER(true);

    // external calls to ERC1271 (can import OpenZeppelin mock implementation)
    // isValidSignature(bytes32 hash, bytes memory signature) returns (bytes4 magicValue) => DISPATCHER(true)
    function _.isValidSignature(bytes32, bytes) external => DISPATCHER(true);

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
    function function getQuorumBitmapIndicesByOperatorIdsAtBlockNumber(uint32 blockNumber, bytes32[] memory operatorIds)
        external returns (uint32[]) envfree;
    getQuorumBitmapByOperatorIdAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) external returns (uint192) envfree;
    function getQuorumBitmapUpdateByOperatorIdByIndex(bytes32 operatorId, uint256 index)
        external returns (IRegistryCoordinator.QuorumBitmapUpdate) envfree;
    function getCurrentQuorumBitmapByOperatorId(bytes32 operatorId) external returns (uint192) envfree;
    function getQuorumBitmapUpdateByOperatorIdLength(bytes32 operatorId) external returns (uint256) envfree;
    function numRegistries() external returns (uint256) envfree;
    function calculateOperatorChurnApprovalDigestHash(
        bytes32 registeringOperatorId,
        bytes calldata quorumNumbers,
        OperatorKickParam[] memory operatorKickParams,
        bytes32 salt,
        uint256 expiry
    ) external returns (bytes32) envfree;
}

// TODOs: add properties in English + rules in CVL