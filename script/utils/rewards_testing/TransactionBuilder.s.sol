// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./TransactionSubmitter.sol";
import "script/utils/ExistingDeploymentParser.sol";
import "@openzeppelin/contracts/utils/Strings.sol";

/**
 * Operator Indices: Derived from `MNEMONIC`. Indices 0-1000
 * AVS Indices: Contract 0-10 in `transaction_submitter` 
 * Strategy Indices: Contract 0-15 in `transaction_submitter`
 * Staker Indices: Derived from `MNEMONIC`. Indices 750-100_000 (500 operators have also deposited)
 * AVS Indices: Contract 0-5 in `transaction_submitter`
 */
contract TransactionBuilder is ExistingDeploymentParser {
    Vm cheats = Vm(VM_ADDRESS);
    using Strings for uint256;

    // Addresses set by _parseState function
    ProxyAdmin proxyAdmin;
    TransactionSubmitter transactionSubmitter;
    uint256 operatorsRegistered;
    string MNEMONIC;

    // Path to EigenLayer contracts
    string contractsPath = "script/configs/holesky/eigenlayer_addresses_preprod.config.json";
    // Path to state so we do not duplicate transactions deployed directly from this contract
    string statePath = "script/utils/rewards_testing/preprod_state.json";

    // Operator Indices
    uint256 minOperatorIndex = 0;
    uint256 maxOperatorIndex = 1000;
    uint256 numOperatorsRegisteredToAVSs = 800;

    // Staker Indices
    uint256 minStakerIndex = 750;
    uint256 stakerNonOperatorStartIndex = 1001;
    uint256 maxStakerIndex = 51_000; // 50000 pure stakers. 250 opStakers
    uint256 firstHalfStakerIndex = 26_000;

    // AVS Indices
    uint16 minAVSIndex = 0;
    uint16 maxAVSIndex = 5;


    /**
     * @notice Seeds operators with 0.0003 ETH
     */
    function seedOperators() external parseState {
        address[] memory operators = _getAllOperators();
        uint256 amountToSend = 0.0003 ether;

        uint256 batchSize = 100;

        // Split up operators into batches of `batchSize`
        for(uint256 i = 0; i < operators.length; i += batchSize){
            address[] memory batch = new address[](batchSize);
            for(uint256 j = 0; j < batchSize; j++){
                batch[j] = operators[i + j];
            }

            vm.startBroadcast();
            transactionSubmitter.sendEth(batch, amountToSend);
            vm.stopBroadcast();
        }
    }

    /**
     * @notice Registers operators with the DelegationManager
     * @dev This function is not deployed by the Submitter contract since we save no gas by doing so
     */
    function registerOperators() external parseState {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: address(0x000000000000000000000000000000000000dEaD),
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });

        for(uint256 i = operatorsRegistered; i < maxOperatorIndex; i++){
            (address operator, uint256 privateKey) = deriveRememberKey(MNEMONIC, uint32(i));
            
            vm.startBroadcast(operator);
            IDelegationManager(delegationManager).registerAsOperator(operatorDetails, "");
            vm.stopBroadcast();

            operatorsRegistered++;
        }

        emit log_named_uint("Operators Deployed", operatorsRegistered);
    }

    /**
     * @notice Registers operators to AVSs
     * @dev A random number of AVSs are registered to for each operator 
     */
    function registerOperatorsToAVSs() external parseState {
        address[] memory avss = transactionSubmitter.getAllAVSs();
        uint256 batchSize = 40;

        for (uint256 i = 0; i < numOperatorsRegisteredToAVSs; i += batchSize) {
            TransactionSubmitter.OperatorAVSRegistration[] memory registrations = new TransactionSubmitter.OperatorAVSRegistration[](batchSize);

            for (uint256 j = 0; j < batchSize; j++) {
                // Get operator 
                (address operator, uint256 privateKey) = deriveRememberKey(MNEMONIC, uint32(i + j));

                // Get number of avss to register
                address[] memory avssToRegister = _getRandomAVSs(avss);
                ISignatureUtils.SignatureWithSaltAndExpiry[] memory signatures = _getOperatorSignatures(privateKey, operator, avssToRegister);              
                
                // Push currRegistration to registrations
                registrations[j] = TransactionSubmitter.OperatorAVSRegistration({
                    operator: operator,
                    avss: avssToRegister,
                    sigs: signatures
                });
            }

            vm.startBroadcast();
            transactionSubmitter.registerOperatorsToAVSs(registrations);
            vm.stopBroadcast();
        }
    }

    function registerOperatorsToAVSsByIndex(uint32 operatorIndex) external parseState {
        address[] memory avss = transactionSubmitter.getAllAVSs();

        (address operator, uint256 privateKey) = deriveRememberKey(MNEMONIC, uint32(operatorIndex));

        // Get number of avss to register
        address[] memory avssToRegister = _getRandomAVSs(avss);
        ISignatureUtils.SignatureWithSaltAndExpiry[] memory signatures = _getOperatorSignatures(privateKey, operator, avssToRegister);              
        
        // Push currRegistration to registrations
        TransactionSubmitter.OperatorAVSRegistration[] memory registrations = new TransactionSubmitter.OperatorAVSRegistration[](1);
        registrations[0] = TransactionSubmitter.OperatorAVSRegistration({
            operator: operator,
            avss: avssToRegister,
            sigs: signatures
        });

        vm.startBroadcast();
        transactionSubmitter.registerOperatorsToAVSs(registrations);
        vm.stopBroadcast();
    }

    /**
     * @notice Deploys TransactionSubmitter contract based on existing network config file
     */
    function deployTransactionSubmitter() external parseState {
        vm.startBroadcast();

        // Deploy TransactionSubmitterImpl
        TransactionSubmitter transactionSubmitterImpl = new TransactionSubmitter(
            IDelegationManager(delegationManager),
            IAVSDirectory(avsDirectory),
            IStrategyManager(strategyManager),
            IRewardsCoordinator(rewardsCoordinator),
            IStrategyFactory(strategyFactory)
        );

        // Deploy ProxyAdmin
        ProxyAdmin transactionSubmitterProxyAdmin = new ProxyAdmin();

        // Deploy Proxy
        TransparentUpgradeableProxy transactionSubmitterProxy = new TransparentUpgradeableProxy(
            address(transactionSubmitterImpl),
            address(proxyAdmin),
            ""
        );

        vm.stopBroadcast();
        

        console.log("ProxyAdmin deployed at: ", address(transactionSubmitterProxyAdmin));
        console.log("TransactionSubmitterProxy deployed at: ", address(transactionSubmitterProxy));
        console.log("TransactionSubmitterImpl deployed at: ", address(transactionSubmitterImpl));
    }

    /**
     * @notice Upgrades TransactionSubmitter contract based on existing network config file
     * @dev Expects `transactionSubmitter` and `proxyAdmin` to be set locally
     */
    function upgradeTransactionSubmitter() external parseState {

        vm.startBroadcast();

        // Deploy new TransactionSubmitterImpl
        TransactionSubmitter transactionSubmitterImpl = new TransactionSubmitter(
            IDelegationManager(delegationManager),
            IAVSDirectory(avsDirectory),
            IStrategyManager(strategyManager),
            IRewardsCoordinator(rewardsCoordinator),
            IStrategyFactory(0xad4A89E3cA9b3dc25AABe0aa7d72E61D2Ec66052)
        );

        // Upgrade Proxy
        proxyAdmin.upgrade(TransparentUpgradeableProxy(payable(address(transactionSubmitter))), address(transactionSubmitterImpl));

        vm.stopBroadcast();

        console.log("TransactionSubmitterProxy upgraded to: ", address(transactionSubmitterImpl));
    }

    function deployCustomStrategies() external parseState {
        string[] memory names = new string[](14);
        string[] memory symbols = new string[](14);
        uint256 startIndex = 1;

        for (uint256 i = 0; i < 14; i++) {
            names[i] = string.concat("StrategyToken_", startIndex.toString());
            symbols[i] = string.concat("ST_", startIndex.toString());
            startIndex++;
        }
        vm.startBroadcast();
        transactionSubmitter.deployCustomStrategies(names, symbols);
        vm.stopBroadcast();
    }

    function delegateFirstHalfOfStakers() external parseState {
        uint256 firstStaker = stakerNonOperatorStartIndex;
        uint256 lastStaker = firstHalfStakerIndex;

        uint256 batchSize = 40;

        for (uint256 i = firstStaker; i < lastStaker; i += batchSize) {
            TransactionSubmitter.StakerDelegation[] memory delegations = new TransactionSubmitter.StakerDelegation[](batchSize);

            for (uint256 j = 0; j < batchSize; j++) {
                uint256 stakerPrivateKey = vm.deriveKey(MNEMONIC, uint32(i + j));
                address staker = vm.addr(stakerPrivateKey);

                address operator = vm.addr(vm.deriveKey(MNEMONIC, uint32(_getOperatorIndexFirstHalf(i + j))));

                ISignatureUtils.SignatureWithExpiry memory stakerSignature = _getStakerDelegationSignature(stakerPrivateKey, staker, operator);

                delegations[j] = TransactionSubmitter.StakerDelegation({
                    staker: staker,
                    operator: operator,
                    stakerSignatureAndExpiry: stakerSignature,
                    approverSignatureAndExpiry: ISignatureUtils.SignatureWithExpiry({
                        signature: "",
                        expiry: 0
                    }),
                    approverSalt: bytes32(0)
                });
            }

            vm.startBroadcast();
            transactionSubmitter.delegateStakers(delegations);
            vm.stopBroadcast();
        }
    }

    function registerOperatorByIndex(uint32 operatorIndex) public parseState {
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            __deprecated_earningsReceiver: address(0x000000000000000000000000000000000000dEaD),
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });

        uint256 opKey = vm.deriveKey(MNEMONIC, uint32(operatorIndex));
        address operator = vm.addr(opKey);

        vm.startBroadcast();
        operator.call{value: 0.0003 ether}("");
        vm.stopBroadcast();

        (address operatorSame, uint256 privateKey) = deriveRememberKey(MNEMONIC, operatorIndex);

        vm.startBroadcast(operatorSame);
        IDelegationManager(delegationManager).registerAsOperator(operatorDetails, "");
        vm.stopBroadcast();
    }

    function undelegateStakers(uint32 startIndex, uint32 endIndex) external parseState {
        for (uint32 i = startIndex; i <= endIndex; i++) {
            (address staker, uint256 stakerPrivateKey) = deriveRememberKey(MNEMONIC, i);

            if (!delegationManager.isDelegated(staker)) {
                continue;
            }

            vm.startBroadcast();
            staker.call{value: 0.0003 ether}("");
            vm.stopBroadcast();

            vm.startBroadcast(staker);
            delegationManager.undelegate(staker);
            vm.stopBroadcast();
        }
    }

    function delegateSecondHalfOfStakers() external parseState {
        uint256 firstStaker = firstHalfStakerIndex + 1 + 2080;
        uint256 lastStaker = maxStakerIndex;

        uint256 batchSize = 40;

        for (uint256 i = firstStaker; i < lastStaker; i += batchSize) {
            TransactionSubmitter.StakerDelegation[] memory delegations = new TransactionSubmitter.StakerDelegation[](batchSize);

            for (uint256 j = 0; j < batchSize; j++) {
                uint256 stakerPrivateKey = vm.deriveKey(MNEMONIC, uint32(i + j));
                address staker = vm.addr(stakerPrivateKey);

                uint256 operatorIndex = vm.randomUint(5, 1000);

                address operator = vm.addr(vm.deriveKey(MNEMONIC, uint32(operatorIndex)));

                if(delegationManager.isDelegated(staker)) {
                    continue;
                }

                if(!delegationManager.isOperator(operator)){
                    registerOperatorByIndex(uint32(operatorIndex));
                }   

                ISignatureUtils.SignatureWithExpiry memory stakerSignature = _getStakerDelegationSignature(stakerPrivateKey, staker, operator);

                delegations[j] = TransactionSubmitter.StakerDelegation({
                    staker: staker,
                    operator: operator,
                    stakerSignatureAndExpiry: stakerSignature,
                    approverSignatureAndExpiry: ISignatureUtils.SignatureWithExpiry({
                        signature: "",
                        expiry: 0
                    }),
                    approverSalt: bytes32(0)
                });
            }

            vm.startBroadcast();
            transactionSubmitter.delegateStakers(delegations);
            vm.stopBroadcast();
        }
    }

    function depositStakersSingleStrat(uint256 startStaker, uint256 endStaker) external parseState {
        uint256 batchSize = 20;
        require ((endStaker - startStaker) % batchSize == 0, "Batch size must be a factor of the total stakers");

        // Get strategies
        IStrategy[] memory strategies = transactionSubmitter.getAllStrategies();
        IERC20[] memory tokens = new IERC20[](strategies.length);
        for(uint256 i = 0; i < strategies.length; i++){
            tokens[i] = IERC20(strategies[i].underlyingToken());
        }

        for(uint256 i = startStaker; i < endStaker; i += batchSize) {
            TransactionSubmitter.StakerDeposit[] memory deposits = new TransactionSubmitter.StakerDeposit[](batchSize);

            for(uint256 j = 0; j < batchSize; j++){
                // Initialize struct params
                TransactionSubmitter.StrategyInfo[] memory strategyInfo = new TransactionSubmitter.StrategyInfo[](1);

                // Get key/address
                uint256 stakerPrivateKey = vm.deriveKey(MNEMONIC, uint32(i + j));
                address staker = vm.addr(stakerPrivateKey);

                // Get random strategy
                uint256 strategyIndex = vm.randomUint(1, strategies.length - 1);
                IStrategy strategy = strategies[strategyIndex];
                IERC20 token = tokens[strategyIndex];

                // Get random amount
                uint256 amount = vm.randomUint(1, 25_000_000e18);

                // Get StakerInfo
                strategyInfo[0] = _getStakerStrategyInfo(
                    stakerPrivateKey,
                    staker,
                    strategy,
                    token,
                    amount
                );

                deposits[j] = TransactionSubmitter.StakerDeposit({
                    strategyInfos: strategyInfo
                });
            }

            vm.startBroadcast();
            transactionSubmitter.depositStakers(deposits);
            vm.stopBroadcast();
        }
    }

    function depositStakersMultiStrat(uint256 startStaker, uint256 endStaker) external parseState {
        uint256 batchSize = 1;
        // require ((endStaker - startStaker) % batchSize == 0, "Batch size must be a factor of the total stakers");

        // Get strategies
        IStrategy[] memory strategies = transactionSubmitter.getAllStrategies();
        IERC20[] memory tokens = new IERC20[](strategies.length);
        for(uint256 i = 0; i < strategies.length; i++){
            tokens[i] = IERC20(strategies[i].underlyingToken());
        }

        for(uint256 i = startStaker; i < endStaker; i += batchSize) {
            TransactionSubmitter.StakerDeposit[] memory deposits = new TransactionSubmitter.StakerDeposit[](batchSize);

            for(uint256 j = 0; j < batchSize; j++){
                // Get key/address
                uint256 stakerPrivateKey = vm.deriveKey(MNEMONIC, uint32(i + j));
                address staker = vm.addr(stakerPrivateKey);

                // Get random number of strategies to deposit for staker
                uint256 numStrategies = vm.randomUint(2, 4);

                // Initialize struct params
                TransactionSubmitter.StrategyInfo[] memory strategyInfos = new TransactionSubmitter.StrategyInfo[](numStrategies);

                uint256 nonce = strategyManager.nonces(staker);

                for (uint256 k = 0; k < numStrategies; k++) {
                    // Get random strategy
                    uint256 strategyIndex = vm.randomUint(1, strategies.length - 1);
                    IStrategy strategy = strategies[strategyIndex];
                    IERC20 token = tokens[strategyIndex];

                    // Get random amount
                    uint256 amount = vm.randomUint(1, 25_000_000e18);

                    // Get StakerInfo
                    strategyInfos[k] = _getStakerStrategyInfoManualNonce(
                        stakerPrivateKey,
                        staker,
                        strategy,
                        token,
                        amount,
                        nonce
                    );

                    nonce++;
                }

                deposits[j] = TransactionSubmitter.StakerDeposit({
                    strategyInfos: strategyInfos
                });
            }

            vm.startBroadcast();
            transactionSubmitter.depositStakers(deposits);
            vm.stopBroadcast();
        }
    }

    function depositStakerByIndex(uint32 index) external parseState {
            // Get strategies
            IStrategy[] memory strategies = transactionSubmitter.getAllStrategies();
            IERC20[] memory tokens = new IERC20[](strategies.length);
            for(uint256 i = 0; i < strategies.length; i++){
                tokens[i] = IERC20(strategies[i].underlyingToken());
            }

            // Initialize struct params
            TransactionSubmitter.StrategyInfo[] memory strategyInfo = new TransactionSubmitter.StrategyInfo[](1);

            // Get key/address
            uint256 stakerPrivateKey = vm.deriveKey(MNEMONIC, uint32(index));
            address staker = vm.addr(stakerPrivateKey);

            // Get random strategy
            uint256 strategyIndex = vm.randomUint(1, strategies.length - 1);
            IStrategy strategy = strategies[strategyIndex];
            IERC20 token = tokens[strategyIndex];

            // Get random amount
            uint256 amount = vm.randomUint(1, 25_000_000e18);

            // Get StakerInfo
            strategyInfo[0] = _getStakerStrategyInfo(
                stakerPrivateKey,
                staker,
                strategy,
                token,
                amount
            );

            TransactionSubmitter.StakerDeposit[] memory deposits = new TransactionSubmitter.StakerDeposit[](1);

            deposits[0] = TransactionSubmitter.StakerDeposit({
                strategyInfos: strategyInfo
            });

            vm.startBroadcast();
            transactionSubmitter.depositStakers(deposits);
            vm.stopBroadcast();
    }
    
    function deployAVSs() external parseState {
        vm.startBroadcast();
        transactionSubmitter.deployAVSs(maxAVSIndex);
        vm.stopBroadcast();
    }

    function delegateByIndex(uint32 stakerIndex, uint32 operatorIndex) public parseState {
        (address staker, uint256 stakerPrivateKey) = deriveRememberKey(MNEMONIC, stakerIndex);
        
        if (delegationManager.isDelegated(staker)) {
            return;
        }

        vm.startBroadcast();
        staker.call{value: 0.0003 ether}("");
        vm.stopBroadcast();
        
        address operator = vm.addr(vm.deriveKey(MNEMONIC, operatorIndex));
        ISignatureUtils.SignatureWithExpiry memory approverSignatureAndExpiry;

        vm.startBroadcast(staker);
        delegationManager.delegateTo(operator, approverSignatureAndExpiry, "");
        vm.stopBroadcast();
    }

    function delegateByIndexMultiple(uint32 stakerIndex, uint32 numberStaker, uint32 operatorIndex) public parseState {
        for(uint32 i = 0; i < numberStaker; i++){

            delegateByIndex(stakerIndex + i, operatorIndex);
        }
    }

    function completeAndDelegate(uint32 stakerIndex, uint32 operatorIndex) public parseState {
        IStrategy[] memory strategies = new IStrategy[](1);
        uint256[] memory shares = new uint256[](1);
        IERC20[] memory tokens = new IERC20[](1);
        strategies[0] = IStrategy(0x77335a08a877cd874165bDC766feeD951a0d84c8);
        shares[0] = 4157021384921519315796965;
        tokens[0] = IERC20(0x4Ce198f835e3f0bb18D4775946fd87f5654B0b49);

        IDelegationManager.Withdrawal memory withdrawal = IDelegationManager.Withdrawal({
            staker: 0xc17e7649c237013603BeD09Be3647d23575a1042,
            delegatedTo: 0x4e1Efdbb25446Aa7C9fAc4731BE159b650E9eE5f,
            withdrawer: 0xc17e7649c237013603BeD09Be3647d23575a1042,
            nonce: 0,
            startBlock: 2329626,
            strategies: strategies,
            shares: shares
        });

        
        (address staker, uint256 stakerPrivateKey) = deriveRememberKey(MNEMONIC, stakerIndex);

        vm.startBroadcast(staker);
        delegationManager.completeQueuedWithdrawal(
            withdrawal,
            tokens,
            0,
            false
        );
        delegationManager.delegateTo(
            vm.addr(vm.deriveKey(MNEMONIC, operatorIndex)),
            ISignatureUtils.SignatureWithExpiry({
                signature: "",
                expiry: 0
            }),
            ""
        );
        vm.stopBroadcast();
    }

    /**
     *
     *                         HELPER FUNCTIONS
     *
     */

    modifier parseState() {
        _parseDeployedContracts(contractsPath);
        _parseState(statePath);
        _;
    }

    function _parseState(string memory statePathToParse) internal {
        // READ JSON CONFIG DATA
        string memory stateData = vm.readFile(statePathToParse);
        emit log_named_string("Using state file", statePathToParse);

        transactionSubmitter = TransactionSubmitter(payable(stdJson.readAddress(stateData, ".submitterProxy")));
        proxyAdmin = ProxyAdmin(stdJson.readAddress(stateData, ".submitterProxyAdmin"));
        operatorsRegistered = stdJson.readUint(stateData, ".operatorsRegistered");
        MNEMONIC = vm.envString("MNEMONIC");
    }

    function _getAllOperators() internal returns(address[] memory){
        address[] memory operators = new address[](maxOperatorIndex);
        for(uint256 i = 0; i < operators.length; i++){
            (address operator, /* privateKey */) = deriveRememberKey(MNEMONIC, uint32(i));
            operators[i] = operator;
        }
        return operators;
    }

    /// @notice helper to get operator Signature
    function _getOperatorSignature(
        uint256 _operatorPrivateKey,
        address operator,
        address avs
    ) internal returns (ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) {
        operatorSignature.expiry = type(uint32).max;
        operatorSignature.salt = bytes32(vm.randomUint());
        {
            bytes32 digestHash = avsDirectory.calculateOperatorAVSRegistrationDigestHash(operator, avs, operatorSignature.salt, operatorSignature.expiry);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(_operatorPrivateKey, digestHash);
            operatorSignature.signature = abi.encodePacked(r, s, v);
        }
        return operatorSignature;
    }

    function _getStakerDelegationSignature(
        uint256 _stakerPrivateKey,
        address staker,
        address operator
    ) internal returns (ISignatureUtils.SignatureWithExpiry memory stakerSignature) {
        stakerSignature.expiry = type(uint32).max;
        {
            bytes32 digestHash = delegationManager.calculateCurrentStakerDelegationDigestHash(staker, operator, stakerSignature.expiry);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(_stakerPrivateKey, digestHash);
            stakerSignature.signature = abi.encodePacked(r, s, v);
        }
        return stakerSignature;
    }

    function _getRandomAVSs(address[] memory avss) internal returns (address[] memory) {
        uint256 min = 1;
        uint256 max = 5;
        uint256 randomNumber = vm.randomUint(min, max);

        address[] memory randomAVSs = new address[](randomNumber);

        for (uint256 i = 0; i < randomNumber; i++) {
            randomAVSs[i] = avss[i];
        }

        return randomAVSs;
    }

    function _getOperatorSignatures(
        uint256 _operatorPrivateKey,
        address operator,
        address[] memory avss
    ) internal returns (ISignatureUtils.SignatureWithSaltAndExpiry[] memory operatorSignatures) {
        ISignatureUtils.SignatureWithSaltAndExpiry[] memory signatures = new ISignatureUtils.SignatureWithSaltAndExpiry[](avss.length);
        for (uint256 i = 0; i < avss.length; i++) {
            signatures[i] = _getOperatorSignature(_operatorPrivateKey, operator, avss[i]);
        }
        return signatures;
    }

    bytes32 private constant PERMIT_TYPEHASH =
        keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

    function _getStakerPermitInfo(
        uint256 stakerPrivateKey,
        address staker,
        IERC20 token,
        address spender,
        uint256 value
    ) internal returns (TransactionSubmitter.PermitInfo memory permitInfo) {
        uint256 nonce = StrategyToken(address(token)).nonces(staker);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(
            stakerPrivateKey,
            keccak256(
                abi.encodePacked(
                    "\x19\x01",
                    StrategyToken(address(token)).DOMAIN_SEPARATOR(),
                    keccak256(abi.encode(PERMIT_TYPEHASH, staker, spender, value, nonce, type(uint256).max))
                )
            )
        );

        return TransactionSubmitter.PermitInfo({
            owner: staker,
            spender: spender,
            value: value,
            deadline: type(uint256).max,
            v: v,
            r: r,
            s: s
        });
    }

    bytes32 public constant DEPOSIT_TYPEHASH =
        keccak256("Deposit(address staker,address strategy,address token,uint256 amount,uint256 nonce,uint256 expiry)");

    bytes32 public constant DOMAIN_SEPARATOR_SM = 0xcd2248fa45cb388d7d043522f71c7934b72e935295a7e04afdf7794a7bf18df2;

    function _getStakerStrategyInfoManualNonce(
        uint256 stakerPrivateKey,
        address staker,
        IStrategy strategy,
        IERC20 token,
        uint256 amount,
        uint256 nonce
    ) internal view returns (TransactionSubmitter.StrategyInfo memory strategyInfo) {
        uint256 expiry = type(uint32).max;
        bytes memory signature;
        {
            bytes32 structHash = keccak256(
                abi.encode(DEPOSIT_TYPEHASH, staker, strategy, address(token), amount, nonce, expiry)
            );
            bytes32 digestHash = keccak256(abi.encodePacked("\x19\x01", DOMAIN_SEPARATOR_SM, structHash));

            (uint8 v, bytes32 r, bytes32 s) = cheats.sign(stakerPrivateKey, digestHash);

            signature = abi.encodePacked(r, s, v);
        }

        return TransactionSubmitter.StrategyInfo({
            strategy: strategy,
            token: token,
            amount: amount,
            staker: staker,
            expiry: expiry,
            signature: signature
        });
    }

    function _getStakerStrategyInfo(
        uint256 stakerPrivateKey,
        address staker,
        IStrategy strategy,
        IERC20 token,
        uint256 amount
    ) internal view returns (TransactionSubmitter.StrategyInfo memory strategyInfo) {
        return _getStakerStrategyInfoManualNonce(stakerPrivateKey, staker, strategy, token, amount, strategyManager.nonces(staker));
    }

    function _getOperatorIndexFirstHalf(uint256 stakerIndex) internal returns (uint256) {
        uint256 lastIndex = firstHalfStakerIndex;
        if (stakerIndex < 11_000) {
            return 0;
        } else if (stakerIndex < 16_000) {
            return 1;
        } else if (stakerIndex < 21_000) {
            return 2;
        } else if (stakerIndex < 23_500) {
            return 3;
        } else if (stakerIndex < 26_000) {
            return 4;
        } else {
            return 5;
        }
    }

    function setWhitelistedStrategies() external parseState {
        IStrategy[] memory strategies = transactionSubmitter.getAllStrategies();
        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](strategies.length);
        IStrategyFactory stratFactory = IStrategyFactory(0xad4A89E3cA9b3dc25AABe0aa7d72E61D2Ec66052);
        vm.startBroadcast();
        stratFactory.whitelistStrategies(strategies, thirdPartyTransfersForbiddenValues);
        vm.stopBroadcast();
    }

    function addAddressToJSON() external parseState {
        // string memory key = "addresses";
        for (uint256 i = 0; i < 51021; i++) {
            address staker = vm.addr(vm.deriveKey(MNEMONIC, uint32(i)));
            // Add to JSON file
            emit log_named_address("Staker", staker);
        }
        // vm.writeJson(vm.serializeString(key, "", ""), "script/utils/rewards_testing/addresses.json");
    }

    function printPrivateKey(uint256 index) external parseState {
        (address addr, uint256 privateKey) = deriveRememberKey(MNEMONIC, uint32(index));
        emit log_named_uint("Private Key", privateKey);
        emit log_named_address("Address", addr);
        // Encoded private key (in hex)
        // emit log_named_string("Private Key (hex)", abi.encodePacked(privateKey).toHexString());
    }
} 