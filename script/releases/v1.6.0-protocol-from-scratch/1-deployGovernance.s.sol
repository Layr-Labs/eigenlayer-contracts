// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {MultisigDeployLib} from "../MultisigDeployLib.sol";
import "../Env.sol";

import "@openzeppelin/contracts/governance/TimelockController.sol";

/// @notice This script is used to deploy the governance contracts on a testnet environment.
/// @dev Before running this script, please ensure that you have deployed the MultichainDeployer multisig
/// This script deploys the following contracts/msigs:
/// - TimelockController
/// - protocolCouncilMultisig
/// - communityMultisig
/// - executorMultisig
/// - operationsMultisig
contract DeployGovernance is EOADeployer {
    using Env for *;

    // Testnet constants
    address public constant TESTNET_OWNER = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;
    uint256 public constant TESTNET_THRESHOLD = 1;

    function _runAsEOA() internal virtual override {
        // Skip if not on source chain
        if (!Env.isSourceChain()) {
            return;
        }

        vm.startBroadcast();

        deployTimelockControllers();
        deployProtocolMultisigs();
        configureTimelockController(Env.timelockController());
        configureTimelockController(Env.beigenTimelockController());

        vm.stopBroadcast();
    }

    function testScript() public virtual {
        // Skip if not on source chain
        if (!Env.isSourceChain()) {
            return;
        }

        runAsEOA();

        // Assert that the multisigs have the proper owners - protocolCouncilMultisig, communityMultisig, and operationsMultisig have the same owners & threshold
        checkMultisig(Env.protocolCouncilMultisig());
        checkMultisig(Env.communityMultisig());
        checkMultisig(Env.opsMultisig());

        // Assert that the executorMultisig and beigenExecutorMultisig have the proper configuration
        checkExecutorMultisig();
        checkBeigenExecutorMultisig();

        // Assert that the timelock controller is configured correctly
        checkTimelockControllerConfig(Env.timelockController());
        checkTimelockControllerConfig(Env.beigenTimelockController());
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

        TimelockController beigenTimelockController = new TimelockController({
            minDelay: 0, // no delay for setup
            proposers: proposers,
            executors: executors,
            admin: address(0)
        });

        zUpdate("timelockController", address(timelockController));
        zUpdate("beigenTimelockController", address(beigenTimelockController));
    }

    ///
    function deployProtocolMultisigs() public {
        // pseudorandom number for salt
        uint256 salt = uint256(keccak256(abi.encode(block.chainid, block.timestamp))); // Pseudo-random salt

        // Set the initial owner of the multisig to the testnet owner
        address[] memory initialOwners = new address[](1);
        initialOwners[0] = TESTNET_OWNER;

        // 1. Deploy protocolCouncilMultisig
        address protocolCouncilMultisig = MultisigDeployLib.deployMultisig({
            initialOwners: initialOwners,
            initialThreshold: TESTNET_THRESHOLD,
            salt: ++salt
        });
        zUpdate("protocolCouncilMultisig", protocolCouncilMultisig);

        // 2. Deploy communityMultisig
        address communityMultisig = MultisigDeployLib.deployMultisig({
            initialOwners: initialOwners,
            initialThreshold: TESTNET_THRESHOLD,
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

        address executorMultisig = MultisigDeployLib.deployMultisig({
            initialOwners: owners_executorMultisig,
            initialThreshold: TESTNET_THRESHOLD,
            salt: ++salt
        });
        zUpdate("executorMultisig", executorMultisig);

        // 4. Deploy beigenExecutorMultisig
        require(
            address(Env.beigenTimelockController()) != address(0),
            "must deploy beigenTimelockController before beigenExecutorMultisig"
        );
        address[] memory owners_beigenExecutorMultisig = new address[](2);
        owners_beigenExecutorMultisig[0] = address(Env.beigenTimelockController());
        owners_beigenExecutorMultisig[1] = Env.communityMultisig();
        address beigenExecutorMultisig = MultisigDeployLib.deployMultisig({
            initialOwners: owners_beigenExecutorMultisig,
            initialThreshold: TESTNET_THRESHOLD,
            salt: ++salt
        });
        zUpdate("beigenExecutorMultisig", beigenExecutorMultisig);

        // 5. Deploy operationsMultisig
        address operationsMultisig = MultisigDeployLib.deployMultisig({
            initialOwners: initialOwners,
            initialThreshold: TESTNET_THRESHOLD,
            salt: ++salt
        });
        zUpdate("operationsMultisig", operationsMultisig);
    }

    function configureTimelockController(
        TimelockController timelockController
    ) public {
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

        // For testnet, the delay is 1 second
        uint256 delayToSet = 1;

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
        address multisig
    ) public view {
        // Check threshold for testnet
        assertEq(MultisigDeployLib.getThreshold(multisig), TESTNET_THRESHOLD);

        // Check owners
        address[] memory expectedOwners = new address[](1);
        expectedOwners[0] = TESTNET_OWNER;

        for (uint256 i = 0; i < expectedOwners.length; i++) {
            assertTrue(MultisigDeployLib.isOwner(multisig, expectedOwners[i]), "owner mismatch");
        }

        // Assert that the owner counts are correct
        assertEq(expectedOwners.length, MultisigDeployLib.getOwners(multisig).length, "owner count mismatch");
    }

    function checkExecutorMultisig() public view {
        // Check threshold
        assertEq(MultisigDeployLib.getThreshold(Env.executorMultisig()), TESTNET_THRESHOLD);

        // Check owner count
        assertEq(MultisigDeployLib.getOwners(Env.executorMultisig()).length, 2, "executorMultisig owner count mismatch");

        // Check owners
        assertTrue(
            MultisigDeployLib.isOwner(Env.executorMultisig(), address(Env.timelockController())),
            "timelockController not in executorMultisig"
        );
        assertTrue(
            MultisigDeployLib.isOwner(Env.executorMultisig(), Env.communityMultisig()),
            "communityMultisig not in executorMultisig"
        );
    }

    function checkBeigenExecutorMultisig() public view {
        // Check threshold
        assertEq(MultisigDeployLib.getThreshold(Env.beigenExecutorMultisig()), TESTNET_THRESHOLD);

        // Check owner count
        assertEq(
            MultisigDeployLib.getOwners(Env.beigenExecutorMultisig()).length,
            2,
            "beigenExecutorMultisig owner count mismatch"
        );

        // Check owners
        assertTrue(
            MultisigDeployLib.isOwner(Env.beigenExecutorMultisig(), address(Env.beigenTimelockController())),
            "timelockController not in beigenExecutorMultisig"
        );
        assertTrue(
            MultisigDeployLib.isOwner(Env.beigenExecutorMultisig(), Env.communityMultisig()),
            "communityMultisig not in beigenExecutorMultisig"
        );
    }

    function checkTimelockControllerConfig(
        TimelockController timelockController
    ) public view {
        // check for proposer + executor rights on Protocol Council multisig
        require(
            timelockController.hasRole(timelockController.PROPOSER_ROLE(), Env.protocolCouncilMultisig()),
            "protocolCouncilMultisig does not have PROPOSER_ROLE on timelockController"
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

        // Check the delay for testnet (1 second)
        assertEq(timelockController.getMinDelay(), 1);
    }
}
