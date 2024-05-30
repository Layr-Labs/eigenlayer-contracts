// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "script/deploy/holesky/Deploy_Test_RewardsCoordinator.s.sol";
import "./ServiceManagerMock.sol";

/**
 * @notice Script used for testing payments on preprod. Includes user actions such as
 * deploying from scratch new ServiceManagers
 * call payForRange through ServiceManagers
 * registering as operators and to AVSs
 * depositing into strategies
 * delegating to operators
 *
 */
contract TestPreprodRewardsCoordinator is ExistingDeploymentParser {
    // 0-20: operators
    // 21-100: stakers
    string internal constant TEST_MNEMONIC = "hundred february vast fluid produce radar notice ridge armed glare panther balance";

    address testAddress = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;
    ProxyAdmin avsProxyAdmin = ProxyAdmin(0x1BEF05C7303d44e0E2FCD2A19d993eDEd4c51b5B);

    /// Constants: contract addresses
    ServiceManagerMock serviceManager;
    IStrategy wethStrategy = IStrategy(0xD523267698C81a372191136e477fdebFa33D9FB4);
    IERC20 eigen = IERC20(0xD58f6844f79eB1fbd9f7091d05f7cb30d3363926);
    IERC20 weth = IERC20(0x94373a4919B3240D86eA41593D5eBa789FEF3848);

    // Temp storage for managing stack in processClaim
    bytes32 merkleRoot;
    uint32 earnerIndex;
    bytes earnerTreeProof;
    address proofEarner;
    bytes32 earnerTokenRoot;

    function _setupScript() internal virtual {
        _parseInitialDeploymentParams("script/configs/holesky/Deploy_RewardsCoordinator.holesky.config.json");
        _parseDeployedContracts("script/output/holesky/M2_deploy_preprod.output.json");

        // require(
        //     msg.sender == testAddress,
        //     "Only the test address can run this script"
        // );
        require(
            address(eigenStrategy) == 0xdcCF401fD121d8C542E96BC1d0078884422aFAD2,
            "Eigen strategy address incorrect"
        );
        require(
            address(wethStrategy) == 0xD523267698C81a372191136e477fdebFa33D9FB4,
            "WETH strategy address incorrect"
        );

        // Overwrite testAddress and multisigs to be EOAowner
        executorMultisig = testAddress;
        operationsMultisig = testAddress;
        pauserMultisig = testAddress;
        communityMultisig = testAddress;
        STRATEGY_MANAGER_WHITELISTER = testAddress;
    }

    /// @notice helper to get operator Signature
    function _getOperatorSignature(
        uint256 _operatorPrivateKey,
        address operator,
        address avs,
        bytes32 salt,
        uint256 expiry
    ) internal view returns (ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) {
        operatorSignature.expiry = expiry;
        operatorSignature.salt = salt;
        {
            bytes32 digestHash = avsDirectory.calculateOperatorAVSRegistrationDigestHash(operator, avs, salt, expiry);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(_operatorPrivateKey, digestHash);
            operatorSignature.signature = abi.encodePacked(r, s, v);
        }
        return operatorSignature;
    }

    /**
        Deploys a new AVS ServiceManager
        ========ANVIL========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvvv \
            --sig "deployNewAVS()"

        ========HOLESKY========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv --verify \
            --sig "deployNewAVS()"
     */
    function deployNewAVS() external virtual {
        _setupScript();
        require(
            address(rewardsCoordinator) == 0xb22Ef643e1E067c994019A4C19e403253C05c2B0,
            "PaymentCoordinator address incorrect"
        );
        require(
            address(avsDirectory) == 0x141d6995556135D4997b2ff72EB443Be300353bC,
            "AVSDirectory address incorrect"
        );
        address emptyContract = 0xc08b788d587F927b49665b90ab35D5224965f3d9;

        vm.startBroadcast();
        ServiceManagerMock serviceManagerImplementation = new ServiceManagerMock(avsDirectory, rewardsCoordinator);

        serviceManager = ServiceManagerMock(
            address(
                new TransparentUpgradeableProxy(emptyContract, address(avsProxyAdmin), "")
            )
        );
        avsProxyAdmin.upgradeAndCall(   
            TransparentUpgradeableProxy(payable(address(serviceManager))),
            address(serviceManagerImplementation),
            abi.encodeWithSelector(
                ServiceManagerMock.initialize.selector,
                testAddress // initialOwner
            )
        );

        require(
            address(serviceManager.rewardsCoordinator()) == address(rewardsCoordinator),
            "PaymentCoordinator not set in ServiceManager"
        );

        vm.stopBroadcast();
    }

    /**
        Deploys a new AVS ServiceManager
        ========ANVIL========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvvv \
            --sig "upgradeAVSs(address[])" "[0x5d78d44aE92bF5C588930B4948D8DbCb93344830]"

        ========HOLESKY========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv --verify \
            --sig "upgradeAVSs(address[])" "[0x5d78d44aE92bF5C588930B4948D8DbCb93344830]"
     */
    function upgradeAVSs(address[] calldata avss) external virtual {
        _setupScript();

        require(
            address(rewardsCoordinator) == 0xb22Ef643e1E067c994019A4C19e403253C05c2B0,
            "PaymentCoordinator address incorrect"
        );
        require(
            address(avsDirectory) == 0x141d6995556135D4997b2ff72EB443Be300353bC,
            "AVSDirectory address incorrect"
        );

        vm.startBroadcast();
        for (uint256 i = 0; i < avss.length; i++) {
            ServiceManagerMock serviceManagerImplementation = new ServiceManagerMock(avsDirectory, rewardsCoordinator);

            serviceManager = ServiceManagerMock(avss[i]);
            avsProxyAdmin.upgrade(   
                TransparentUpgradeableProxy(payable(address(serviceManager))),
                address(serviceManagerImplementation)
            );

            require(
                address(serviceManager.rewardsCoordinator()) == address(rewardsCoordinator),
                "PaymentCoordinator not set in ServiceManager"
            );
        }


        vm.stopBroadcast();
    }

    /**
        Submit a RewardsSubmission to the PaymentCoordinator
        ========ANVIL========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "createAVSRewardsSubmission(string memory startTimestampType, address avsServiceManager, address paymentTokenAddress, uint256 amount)" \
            "genesis" 0x5d78d44aE92bF5C588930B4948D8DbCb93344830 0x0000000000000000000000000000000000000000 1000000000000000000000000000000

        ========HOLESKY========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "createAVSRewardsSubmission(string memory startTimestampType, address avsServiceManager, address paymentTokenAddress, uint256 amount)" \
            "genesis" 0x5384B6701DbBAAa0a78A4F20cc48f150f7b1526D 0x0000000000000000000000000000000000000000 1000000000000000000000000000000
     */
    function createAVSRewardsSubmission(
        string memory startTimestampType,
        address avsServiceManager,
        address paymentTokenAddress,
        uint256 amount
    ) external virtual {
        _setupScript();

        uint256 mockTokenInitialSupply = 1e32;
        uint32 startTimestamp;
        uint32 duration;

        IERC20 paymentToken = IERC20(paymentTokenAddress);
        // if paymentToken is address(0) deploy a new mocktoken
        if (paymentTokenAddress == address(0)) {
            vm.startBroadcast();
            paymentToken = new ERC20PresetFixedSupply(
                "mockpayment token",
                "MOCK1",
                mockTokenInitialSupply,
                msg.sender
            );
            vm.stopBroadcast();
        }

        if (keccak256(abi.encode(startTimestampType)) == keccak256(abi.encode("genesis"))) {
            /// 1710979200 unix timestamp, genesis at March 21st
            startTimestamp = REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP;
            duration = 10 weeks;
        } else if (keccak256(abi.encode(startTimestampType)) == keccak256(abi.encode("current"))) {
            /// 1715212800 unix timestamp, May 9th
            startTimestamp = REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP + 9 weeks;
            duration = 4 weeks;
        } else if (keccak256(abi.encode(startTimestampType)) == keccak256(abi.encode("future"))) {
            /// 1715817600 unix timestamp, May 16th
            startTimestamp = REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP + 8 weeks;
            duration = 2 weeks;
        }

        IRewardsCoordinator.StrategyAndMultiplier[] memory strategiesAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](2);
        strategiesAndMultipliers[0] = IRewardsCoordinator.StrategyAndMultiplier({
            strategy: wethStrategy,
            multiplier: 2e18
        });
        strategiesAndMultipliers[1] = IRewardsCoordinator.StrategyAndMultiplier({
            strategy: eigenStrategy,
            multiplier: 1e18
        });

        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: strategiesAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: startTimestamp,
            duration: duration
        });

        vm.startBroadcast();
        paymentToken.approve(avsServiceManager, amount);
        ServiceManagerMock(avsServiceManager).createAVSRewardsSubmission(rewardsSubmissions);
        vm.stopBroadcast();
    }

    /**
        Submit a PayForAll to the PaymentCoordinator
        Address associated with `PRIVATE_KEY` must be a payForAllSubmitter
        ========ANVIL========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "payForAll(string memory startTimestampType, address paymentTokenAddress, uint256 amount)" \
            "genesis" 0x0000000000000000000000000000000000000000 1000000000000000000000000000000

        ========HOLESKY========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig " "payForAll(string memory startTimestampType, address paymentTokenAddress, uint256 amount)" \
            "genesis" 0x0000000000000000000000000000000000000000 1000000000000000000000000000000
     */
    function payForAll(
        string memory startTimestampType,
        address paymentTokenAddress,
        uint256 amount
    ) external {
        _setupScript();

        uint256 mockTokenInitialSupply = 1e32;
        uint32 startTimestamp;
        uint32 duration;

        IERC20 paymentToken = IERC20(paymentTokenAddress);
        // if paymentToken is address(0) deploy a new mocktoken
        if (paymentTokenAddress == address(0)) {
            vm.startBroadcast();
            paymentToken = new ERC20PresetFixedSupply(
                "mockpayment token",
                "MOCK1",
                mockTokenInitialSupply,
                msg.sender
            );
            vm.stopBroadcast();
        }

        if (keccak256(abi.encode(startTimestampType)) == keccak256(abi.encode("genesis"))) {
            /// 1710979200 unix timestamp, genesis at March 21st
            startTimestamp = REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP;
            duration = 10 weeks;
        } else if (keccak256(abi.encode(startTimestampType)) == keccak256(abi.encode("current"))) {
            /// 1715212800 unix timestamp, May 9th
            startTimestamp = REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP + 11 weeks;
            duration = 3 weeks;
        } else if (keccak256(abi.encode(startTimestampType)) == keccak256(abi.encode("future"))) {
            /// 1715817600 unix timestamp, May 16th
            startTimestamp = REWARDS_COORDINATOR_GENESIS_REWARDS_TIMESTAMP + 8 weeks;
            duration = 2 weeks;
        }

        IRewardsCoordinator.StrategyAndMultiplier[] memory strategiesAndMultipliers = new IRewardsCoordinator.StrategyAndMultiplier[](2);
        strategiesAndMultipliers[0] = IRewardsCoordinator.StrategyAndMultiplier({
            strategy: wethStrategy,
            multiplier: 1e18
        });
        strategiesAndMultipliers[1] = IRewardsCoordinator.StrategyAndMultiplier({
            strategy: eigenStrategy,
            multiplier: 2e18
        });

        IRewardsCoordinator.RewardsSubmission[] memory rewardsSubmissions = new IRewardsCoordinator.RewardsSubmission[](1);
        rewardsSubmissions[0] = IRewardsCoordinator.RewardsSubmission({
            strategiesAndMultipliers: strategiesAndMultipliers,
            token: paymentToken,
            amount: amount,
            startTimestamp: startTimestamp,
            duration: duration
        });

        vm.startBroadcast();
        paymentToken.approve(address(rewardsCoordinator), amount);
        rewardsCoordinator.createRewardsForAllSubmission(rewardsSubmissions);
        vm.stopBroadcast();
    }

    /**
     * @notice Takes the latest distributionRoot and uses the claim against it. Broadcasts with earnerIndex and the test mnemonic
        ========ANVIL========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "processClaim(string memory processClaimsPath, uint8 earnerIndexMnemonic)" \
            "script/utils/rewardsCoordinator/claimProofs/processClaim_2.json" 31

        ========HOLESKY========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "processClaim(string memory processClaimsPath, uint8 earnerIndexMnemonic)" \
            "script/utils/rewardsCoordinator/claimProofs/processClaim_2.json" 31
     */
    function processClaim(string memory processClaimsPath, uint8 earnerIndexMnemonic) external virtual {
        _setupScript();
        string memory claimProofData = vm.readFile(processClaimsPath);

        (address earner, /* privateKey */) = deriveRememberKey(TEST_MNEMONIC, uint32(earnerIndexMnemonic));

        // Parse PaymentMerkleClaim
        merkleRoot = abi.decode(stdJson.parseRaw(claimProofData, ".proof.root"), (bytes32));
        earnerIndex = abi.decode(stdJson.parseRaw(claimProofData, ".proof.earnerIndex"), (uint32));
        earnerTreeProof = abi.decode(stdJson.parseRaw(claimProofData, ".proof.earnerTreeProof"), (bytes));
        proofEarner = stdJson.readAddress(claimProofData, ".proof.earnerLeaf.earner");
        require(earner == proofEarner, "earner index and in json file do not match");
        earnerTokenRoot = abi.decode(stdJson.parseRaw(claimProofData, ".proof.earnerLeaf.earnerTokenRoot"), (bytes32));
        uint256 numTokenLeaves = stdJson.readUint(claimProofData, ".proof.tokenLeavesNum");
        uint256 numTokenTreeProofs = stdJson.readUint(claimProofData, ".proof.tokenTreeProofsNum");
 
        IRewardsCoordinator.TokenTreeMerkleLeaf[] memory tokenLeaves = new IRewardsCoordinator.TokenTreeMerkleLeaf[](
            numTokenLeaves
        );
        uint32[] memory tokenIndices = new uint32[](numTokenLeaves);
        for (uint256 i = 0; i < numTokenLeaves; ++i) {
            string memory tokenKey = string.concat(".proof.tokenLeaves[", vm.toString(i), "].token");
            string memory amountKey = string.concat(".proof.tokenLeaves[", vm.toString(i), "].cumulativeEarnings");
            string memory leafIndicesKey = string.concat(".proof.leafIndices[", vm.toString(i), "]");

            IERC20 token = IERC20(stdJson.readAddress(claimProofData, tokenKey));
            uint256 cumulativeEarnings = stdJson.readUint(claimProofData, amountKey);
            tokenLeaves[i] = IRewardsCoordinator.TokenTreeMerkleLeaf({
                token: token,
                cumulativeEarnings: cumulativeEarnings
            });
            tokenIndices[i] = uint32(stdJson.readUint(claimProofData, leafIndicesKey));
        }
        bytes[] memory tokenTreeProofs = new bytes[](numTokenTreeProofs);
        for (uint256 i = 0; i < numTokenTreeProofs; ++i) {
            string memory tokenTreeProofKey = string.concat(".proof.tokenTreeProofs[", vm.toString(i), "]");
            tokenTreeProofs[i] = abi.decode(stdJson.parseRaw(claimProofData, tokenTreeProofKey), (bytes));
        }

        IRewardsCoordinator.RewardsMerkleClaim memory newClaim = IRewardsCoordinator.RewardsMerkleClaim({
            rootIndex: uint32(rewardsCoordinator.getDistributionRootsLength() - 1),
            earnerIndex: earnerIndex,
            earnerTreeProof: earnerTreeProof,
            earnerLeaf: IRewardsCoordinator.EarnerTreeMerkleLeaf({earner: earner, earnerTokenRoot: earnerTokenRoot}),
            tokenIndices: tokenIndices,
            tokenTreeProofs: tokenTreeProofs,
            tokenLeaves: tokenLeaves
        });

        vm.startBroadcast(earner);
        rewardsCoordinator.processClaim(newClaim, earner);
        vm.stopBroadcast();
    }

    /**
        ========ANVIL========
        Deposit WETH strategy:
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "depositIntoStrategy(uint8 numStakers, address strategy)" 20 0xD523267698C81a372191136e477fdebFa33D9FB4

        Deposit EIGEN strategy:
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "depositIntoStrategy(uint8 numStakers, address strategy)" 20 0xdcCF401fD121d8C542E96BC1d0078884422aFAD2

        ========HOLESKY========
        Deposit WETH strategy:
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "depositIntoStrategy(uint8 numStakers, address strategy)" 20 0xD523267698C81a372191136e477fdebFa33D9FB4

        Deposit EIGEN strategy:
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "depositIntoStrategy(uint8 numStakers, address strategy)" 20 0xdcCF401fD121d8C542E96BC1d0078884422aFAD2
     */
    function depositIntoStrategy(uint8 numStakers, address strategy) public virtual {
        _setupScript();
        require(
            numStakers <= 80,
            "numStakers should not be larger than 80, indexes 20-100"
        );
        require(
            strategyManager.strategyIsWhitelistedForDeposit(IStrategy(strategy)),
            "Strategy not whitelisted for deposit"
        );
        IERC20 token = IStrategy(strategy).underlyingToken();
        if (strategy == address(eigenStrategy)) {
            // Use EIGEN instead of bEIGEN if its EIGEN strategy
            token = IERC20(0xD58f6844f79eB1fbd9f7091d05f7cb30d3363926);
        }

        for (uint256 i = 20; i < 20 + numStakers; ++i) {
            (address staker, /* privateKey */) = deriveRememberKey(TEST_MNEMONIC, uint32(i));

            // Transfer strategy token to staker
            vm.startBroadcast();
            uint256 depositAmount = 1e16;
            token.transfer(staker, depositAmount);
            vm.stopBroadcast();
            
            // Deposit into strategy as the staker
            vm.startBroadcast(staker);
            token.approve(address(strategyManager), depositAmount);
            strategyManager.depositIntoStrategy(IStrategy(strategy), token, depositAmount);
            vm.stopBroadcast();
        }
    }

    /**
        Deposit into WETH and EIGEN strategies
        ========ANVIL========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "depositWETHAndEIGEN(uint8 numStakers)" 80

        ========HOLESKY========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "depositWETHAndEIGEN(uint8 numStakers)" 10
     */
    function depositWETHAndEIGEN(uint8 numStakers) external virtual {
        _setupScript();
        require(
            numStakers <= 80,
            "numStakers should not be larger than 80, indexes 20-100"
        );

        for (uint256 i = 20; i < 20 + numStakers; ++i) {
            (address staker, /* privateKey */) = deriveRememberKey(TEST_MNEMONIC, uint32(i));

            // Transfer strategy token to staker
            vm.startBroadcast();
            uint256 depositAmount = 1e16;
            eigen.transfer(staker, depositAmount);
            weth.transfer(staker, depositAmount);
            vm.stopBroadcast();
            
            // Deposit into strategy as the staker
            vm.startBroadcast(staker);
            eigen.approve(address(strategyManager), depositAmount);
            strategyManager.depositIntoStrategy(eigenStrategy, eigen, depositAmount);

            weth.approve(address(strategyManager), depositAmount);
            strategyManager.depositIntoStrategy(wethStrategy, weth, depositAmount);
            vm.stopBroadcast();
        }
    }

    /**
        Register Operators in the DM
        ========ANVIL========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "registerOperators(uint8 numOperators, address serviceManagerAddress)" 5 0xa5d5E9bcdDC1dACe96E5d8f7536A97900550BbB2

        ========HOLESKY========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "registerOperators(uint8 numOperators, address serviceManagerAddress)" 1 0xa5d5E9bcdDC1dACe96E5d8f7536A97900550BbB2
     */
    function registerOperators(uint8 numOperators, address serviceManagerAddress) external virtual {
        _setupScript();
        ServiceManagerMock serviceManager = ServiceManagerMock(serviceManagerAddress);
        require(
            numOperators <= 20,
            "numStakers should not be larger than 80, indexes 20-100"
        );

        for (uint256 i = 0; i < numOperators; ++i) {
            (address operator, uint256 privateKey) = deriveRememberKey(TEST_MNEMONIC, uint32(i));
            // get signature
            ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _getOperatorSignature(
                privateKey,
                operator,
                address(serviceManagerAddress),
                bytes32(0), // salt
                type(uint256).max //expiry
            );

            // Transfer strategy token to staker
            vm.startBroadcast(operator);
            IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
                earningsReceiver: address(operator),
                delegationApprover: address(0),
                stakerOptOutWindowBlocks: 0
            });
            delegationManager.registerAsOperator(
                operatorDetails,
                ""
            );

            serviceManager.registerOperatorToAVS(operator, operatorSignature);
            vm.stopBroadcast();
        }
    }

    /**
        Register Operators in the DM
        ========ANVIL========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "registerOperatorByIndex(uint8 operatorIndex, address serviceManagerAddress)" 5 0xa5d5E9bcdDC1dACe96E5d8f7536A97900550BbB2

        ========HOLESKY========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "registerOperatorByIndex(uint8 operatorIndex, address serviceManagerAddress)" 1 0xa5d5E9bcdDC1dACe96E5d8f7536A97900550BbB2
     */
    function registerOperatorByIndex(uint8 operatorIndex, address serviceManagerAddress) external virtual {
        _setupScript();
        ServiceManagerMock serviceManager = ServiceManagerMock(serviceManagerAddress);

        (address operator, uint256 privateKey) = deriveRememberKey(TEST_MNEMONIC, uint32(operatorIndex));
        // get signature
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature = _getOperatorSignature(
            privateKey,
            operator,
            address(serviceManagerAddress),
            bytes32(0), // salt
            type(uint256).max //expiry
        );

        // Transfer strategy token to staker
        vm.startBroadcast(operator);
        IDelegationManager.OperatorDetails memory operatorDetails = IDelegationManager.OperatorDetails({
            earningsReceiver: address(operator),
            delegationApprover: address(0),
            stakerOptOutWindowBlocks: 0
        });
        delegationManager.registerAsOperator(
            operatorDetails,
            ""
        );

        serviceManager.registerOperatorToAVS(operator, operatorSignature);
        vm.stopBroadcast();

    }

    /**
        Delegate Stakers to Operators
        ========ANVIL========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "delegateTo(uint8 stakerId, uint8 operatorId)" 40 10

        ========HOLESKY========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "delegateTo(uint8 stakerId, uint8 operatorId)" 40 10
     */
    function delegateTo(uint8 stakerId, uint8 operatorId) external virtual {
        _setupScript();
        (address staker, /* privateKey */) = deriveRememberKey(TEST_MNEMONIC, uint32(stakerId));
        (address operator, /* privateKey */) = deriveRememberKey(TEST_MNEMONIC, uint32(operatorId));

        ISignatureUtils.SignatureWithExpiry memory signature = ISignatureUtils.SignatureWithExpiry({
            signature: abi.encodePacked("0x"),
            expiry: 0
        });

        vm.startBroadcast(staker);
        delegationManager.delegateTo(operator, signature, bytes32(0));
        vm.stopBroadcast();
    }

    /**
        Transfer ETH to the test addresses so they can submit txs
        ========ANVIL========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "transferETH()"

        ========HOLESKY========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "transferETH()"

        currently transferred 0-16
    */
    function transferETH() external virtual {
        for (uint256 i = 20; i < 40; ++i) {

            vm.startBroadcast();
            (address user, /*privateKey*/) = deriveRememberKey(TEST_MNEMONIC, uint32(i));
            payable(user).call{value: 1 ether}("");
            vm.stopBroadcast();
            require(user.balance > 0, "Should be nonzero balance of eth");
        }
    }

    /**
        ========HOLESKY========
        forge script script/utils/rewardsCoordinator/TestPreprodRewardsCoordinator.s.sol:TestPreprodRewardsCoordinator \
            --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --broadcast -vvvv \
            --sig "deployStablecoin()"
     */
    function deployStablecoin() external virtual {
        vm.startBroadcast();
        TestStablecoin stablecoin = new TestStablecoin();
        vm.stopBroadcast();
    }
}


contract TestStablecoin is ERC20PresetFixedSupply {
    constructor() ERC20PresetFixedSupply("My Stablecoin", "TSC", 10000000000 * 10**6, msg.sender) {}

    function decimals() public pure override returns (uint8) {
        return 6;
    }
}