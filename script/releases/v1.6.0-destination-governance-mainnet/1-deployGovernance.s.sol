// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/governance/TimelockController.sol";

// For TOML parsing
import {stdToml} from "forge-std/StdToml.sol";

/// @notice This script is used to deploy the governance contracts for the mainnet destination chain.
contract DeployGovernance is EOADeployer {
    using Env for *;
    using stdToml for string;

    // Mainnet constants
    uint256 public constant PROTOCOL_COUNCIL_AND_COMMUNITY_THRESHOLD = 3;
    uint256 public constant EXECUTOR_THRESHOLD = 1;

    function _runAsEOA() internal override {
        if (!Env.isDestinationChain() || !Env._strEq(Env.env(), "base")) {
            return;
        }

        vm.startBroadcast();

        deployTimelockControllers();
        deployProtocolMultisigs();
        configureTimelockController();

        vm.stopBroadcast();
    }

    function testDeploy() public virtual {
        if (!Env.isDestinationChain() || !Env._strEq(Env.env(), "base")) {
            return;
        }
        
        runAsEOA();

        // Assert that the multisigs have the proper owners - protocolCouncilMultisig and communityMultisig have the same owners & threshold
        checkMultisig(Multisig(Env.protocolCouncilMultisig()));
        checkMultisig(Multisig(Env.communityMultisig()));

        // Assert that the executorMultisig has the proper configuration
        checkExecutorMultisig();

        // Assert that the timelock controller is configured correctly
        checkTimelockControllerConfig();
    }

    // set up initially with deployer as a proposer & executor, to be renounced prior to finalizing deployment
    function deployTimelockControllers() public {
        address[] memory proposers = new address[](1);
        proposers[0] = msg.sender;

        address[] memory executors = new address[](1);
        executors[0] = msg.sender;

        TimelockController timelockController = new TimelockController({
            minDelay: 0, // no delay for setup
            proposers: proposers,
            executors: executors,
            admin: address(0)
        });

        zUpdate("timelockController", address(timelockController));
    }

    ///
    function deployProtocolMultisigs() public {
        // pseudorandom number for salt
        uint256 salt = uint256(keccak256(abi.encode(block.chainid, block.timestamp))); // Pseudo-random salt

        // Set the initial owner of the multisig from TOML file
        address[] memory initialOwners = _getMultisigOwner();

        // 1. Deploy protocolCouncilMultisig
        address protocolCouncilMultisig = deployMultisig({
            initialOwners: initialOwners,
            initialThreshold: PROTOCOL_COUNCIL_AND_COMMUNITY_THRESHOLD,
            salt: ++salt
        });
        zUpdate("protocolCouncilMultisig", protocolCouncilMultisig);

        // 2. Deploy communityMultisig
        address communityMultisig = deployMultisig({
            initialOwners: initialOwners,
            initialThreshold: PROTOCOL_COUNCIL_AND_COMMUNITY_THRESHOLD,
            salt: ++salt
        });
        zUpdate("communityMultisig", communityMultisig);

        // 3. Deploy primary executorMultisig
        require(
            address(Env.timelockController()) != address(0), "must deploy timelockController before executorMultisig"
        );
        address[] memory owners_executorMultisig = new address[](2);
        owners_executorMultisig[0] = address(Env.timelockController());
        owners_executorMultisig[1] = Env.communityMultisig();

        address executorMultisig =
            deployMultisig({initialOwners: owners_executorMultisig, initialThreshold: EXECUTOR_THRESHOLD, salt: ++salt});
        zUpdate("executorMultisig", executorMultisig);
    }

    function deployMultisig(
        address[] memory initialOwners,
        uint256 initialThreshold,
        uint256 salt
    ) internal returns (address) {
        // addresses taken from https://github.com/safe-global/safe-smart-account/blob/main/CHANGELOG.md#expected-addresses-with-deterministic-deployment-proxy-default
        // NOTE: double check these addresses are correct on each chain
        address safeFactory = 0x4e1DCf7AD4e460CfD30791CCC4F9c8a4f820ec67;
        address safeSingleton = 0x29fcB43b46531BcA003ddC8FCB67FFE91900C762; // Gnosis safe L2 singleton
        address safeFallbackHandler = 0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99;

        bytes memory emptyData;

        bytes memory initializerData = abi.encodeWithSignature(
            "setup(address[],uint256,address,bytes,address,address,uint256,address)",
            initialOwners, /* signers */
            initialThreshold, /* threshold */
            address(0), /* to (used in setupModules) */
            emptyData, /* data (used in setupModules) */
            safeFallbackHandler,
            address(0), /* paymentToken */
            0, /* payment */
            payable(address(0)) /* paymentReceiver */
        );

        bytes memory calldataToFactory =
            abi.encodeWithSignature("createProxyWithNonce(address,bytes,uint256)", safeSingleton, initializerData, salt);

        (bool success, bytes memory returndata) = safeFactory.call(calldataToFactory);
        require(success, "multisig deployment failed");
        address deployedMultisig = abi.decode(returndata, (address));
        require(deployedMultisig != address(0), "something wrong in multisig deployment, zero address returned");
        return deployedMultisig;
    }

    function configureTimelockController() public {
        // Get the timelock controller
        TimelockController timelockController = Env.timelockController();

        // We have 10 actions to perform on the timelock controller
        uint256 tx_array_length = 10;
        address[] memory targets = new address[](tx_array_length);
        for (uint256 i = 0; i < targets.length; ++i) {
            targets[i] = address(timelockController);
        }
        uint256[] memory values = new uint256[](tx_array_length);
        bytes[] memory payloads = new bytes[](tx_array_length);

        // 1. remove sender as canceller
        payloads[0] =
            abi.encodeWithSelector(AccessControl.revokeRole.selector, timelockController.CANCELLER_ROLE(), msg.sender);
        // 2. remove sender as executor
        payloads[1] =
            abi.encodeWithSelector(AccessControl.revokeRole.selector, timelockController.EXECUTOR_ROLE(), msg.sender);
        // 3. remove sender as proposer
        payloads[2] =
            abi.encodeWithSelector(AccessControl.revokeRole.selector, timelockController.PROPOSER_ROLE(), msg.sender);
        // 4. remove sender as admin
        payloads[3] = abi.encodeWithSelector(
            AccessControl.revokeRole.selector, timelockController.TIMELOCK_ADMIN_ROLE(), msg.sender
        );

        // 5. add operationsMultisig as canceller
        payloads[4] = abi.encodeWithSelector(
            AccessControl.grantRole.selector, timelockController.CANCELLER_ROLE(), Env.opsMultisig()
        );
        // 6. add operationsMultisig as proposer
        payloads[5] = abi.encodeWithSelector(
            AccessControl.grantRole.selector, timelockController.PROPOSER_ROLE(), Env.opsMultisig()
        );

        // 7. add protocolCouncilMultisig as proposer
        payloads[6] = abi.encodeWithSelector(
            AccessControl.grantRole.selector, timelockController.PROPOSER_ROLE(), Env.protocolCouncilMultisig()
        );
        // 8. add protocolCouncilMultisig as executor
        payloads[7] = abi.encodeWithSelector(
            AccessControl.grantRole.selector, timelockController.EXECUTOR_ROLE(), Env.protocolCouncilMultisig()
        );

        // 9. add communityMultisig as admin
        payloads[8] = abi.encodeWithSelector(
            AccessControl.grantRole.selector, timelockController.TIMELOCK_ADMIN_ROLE(), Env.communityMultisig()
        );

        // For mainnet, the delay is 1 day
        uint256 delayToSet = 1 days;

        require(delayToSet != 0, "delay not calculated");
        // 10. set min delay to appropriate length
        payloads[9] = abi.encodeWithSelector(timelockController.updateDelay.selector, delayToSet);

        // schedule the batch
        timelockController.scheduleBatch(
            targets,
            values,
            payloads,
            bytes32(0), // no predecessor needed
            bytes32(0), // no salt
            0 // 0 enforced delay
        );

        // execute the batch
        timelockController.executeBatch(
            targets,
            values,
            payloads,
            bytes32(0), // no predecessor needed
            bytes32(0) // no salt
        );
    }

    /**
     *
     *                         CHECKING FUNCTIONS
     *
     */

    /// @dev Used to check the configuration of the protocolCouncilMultisig and communityMultisig
    function checkMultisig(
        Multisig multisig
    ) public view {
        // Check threshold
        assertEq(multisig.getThreshold(), PROTOCOL_COUNCIL_AND_COMMUNITY_THRESHOLD);

        // Check owners
        address[] memory expectedOwners = _getMultisigOwner();
        for (uint256 i = 0; i < expectedOwners.length; i++) {
            assertTrue(multisig.isOwner(expectedOwners[i]), "owner mismatch");
        }

        // Assert that the owner counts are correct
        assertEq(expectedOwners.length, multisig.getOwners().length, "owner count mismatch");
    }

    function checkExecutorMultisig() public view {
        // Check threshold
        assertEq(Multisig(Env.executorMultisig()).getThreshold(), EXECUTOR_THRESHOLD);

        // Check owner count
        assertEq(Multisig(Env.executorMultisig()).getOwners().length, 2, "executorMultisig owner count mismatch");

        // Check owners
        assertTrue(
            Multisig(Env.executorMultisig()).isOwner(address(Env.timelockController())),
            "timelockController not in executorMultisig"
        );
        assertTrue(
            Multisig(Env.executorMultisig()).isOwner(Env.communityMultisig()),
            "communityMultisig not in executorMultisig"
        );
    }

    function checkTimelockControllerConfig() public view {
        TimelockController timelockController = Env.timelockController();

        // check for proposer + executor rights on Protocol Council multisig
        require(
            timelockController.hasRole(timelockController.PROPOSER_ROLE(), Env.protocolCouncilMultisig()),
            "protocolCouncilMultisig does not have PROPOSER_ROLE on timelockController"
        );
        require(
            timelockController.hasRole(timelockController.EXECUTOR_ROLE(), Env.protocolCouncilMultisig()),
            "protocolCouncilMultisig does not have EXECUTOR_ROLE on timelockController"
        );

        // check for proposer + canceller rights on ops multisig
        require(
            timelockController.hasRole(timelockController.PROPOSER_ROLE(), Env.opsMultisig()),
            "operationsMultisig does not have PROPOSER_ROLE on timelockController"
        );
        require(
            timelockController.hasRole(timelockController.CANCELLER_ROLE(), Env.opsMultisig()),
            "operationsMultisig does not have CANCELLER_ROLE on timelockController"
        );

        // check that community multisig has admin rights
        require(
            timelockController.hasRole(timelockController.TIMELOCK_ADMIN_ROLE(), Env.communityMultisig()),
            "communityMultisig does not have TIMELOCK_ADMIN_ROLE on timelockController"
        );

        // check for self-administration
        require(
            timelockController.hasRole(timelockController.TIMELOCK_ADMIN_ROLE(), address(timelockController)),
            "timelockController does not have TIMELOCK_ADMIN_ROLE on itself"
        );

        // check that deployer has no rights
        require(
            !timelockController.hasRole(timelockController.TIMELOCK_ADMIN_ROLE(), msg.sender),
            "deployer erroenously retains TIMELOCK_ADMIN_ROLE on timelockController"
        );
        require(
            !timelockController.hasRole(timelockController.PROPOSER_ROLE(), msg.sender),
            "deployer erroenously retains PROPOSER_ROLE on timelockController"
        );
        require(
            !timelockController.hasRole(timelockController.EXECUTOR_ROLE(), msg.sender),
            "deployer erroenously retains EXECUTOR_ROLE on timelockController"
        );
        require(
            !timelockController.hasRole(timelockController.CANCELLER_ROLE(), msg.sender),
            "deployer erroenously retains CANCELLER_ROLE on timelockController"
        );

        // Check the delay for mainnet (1 day)
        assertEq(timelockController.getMinDelay(), 1 days);
    }

    function _getMultisigOwner() internal view returns (address[] memory) {
        // Read the TOML file
        string memory path = "script/releases/v1.6.0-destination-governance-mainnet/pcAndCommunityOwners.toml";
        string memory root = vm.projectRoot();
        string memory fullPath = string.concat(root, "/", path);
        string memory toml = vm.readFile(fullPath);

        // Parse the owners array from the TOML
        return toml.readAddressArray(".owners");
    }
}

interface Multisig {
    function getThreshold() external view returns (uint256);
    function getOwners() external view returns (address[] memory);
    function isOwner(
        address owner
    ) external view returns (bool);
}
