// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import "../Env.sol";

import "@openzeppelin/contracts/governance/TimelockController.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "src/test/mocks/EmptyContract.sol";

contract Deploy is EOADeployer {
    using Env for *;

    function _runAsEOA() internal override {
        vm.startBroadcast();

        deployTimelockControllers();
        deployProtocolMultisigs();
        configureTimelockController(Env.timelockController());
        configureTimelockController(Env.beigenTimelockController());
        deployTokensAndStrategy();

        vm.stopBroadcast();
    }

    function testDeploy() public virtual {
        _runAsEOA();
        checkGovernanceConfiguration();
    }

    // set up initially with deployer as a proposer & executor, to be renounced prior to finalizing deployment
    function deployTimelockControllers() public {
        address[] memory proposers = new address[](1);
        proposers[0] = msg.sender;

        address[] memory executors = new address[](1);
        executors[0] = msg.sender;

        vm.startBroadcast();
        deployImpl({
            name: "timelockController",
            deployedTo: address(new TimelockController({
                minDelay: 0, // no delay for setup
                proposers: proposers, 
                executors: executors,
                admin: address(0)
            }))
        });
        deployImpl({
            name: "beigenTimelockController",
            deployedTo: address(new TimelockController({
                minDelay: 0, // no delay for setup
                proposers: proposers, 
                executors: executors,
                admin: address(0)
            }))
        });
    }

    function deployProtocolMultisigs() public {
        // deploy multisigs that simply have the deployer as their initial owner
        address[] memory singleOwner = new address[](1);
        singleOwner[0] = msg.sender;
        deployImpl({
            name: "pauserMultisig",
            deployedTo: deployMultisig({
                initialOwners: singleOwner,
                initialThreshold: 1
            })
        });
        deployImpl({
            name: "opsMultisig",
            deployedTo: deployMultisig({
                initialOwners: singleOwner,
                initialThreshold: 1
            })
        });
        deployImpl({
            name: "protocolCouncilMultisig",
            deployedTo: deployMultisig({
                initialOwners: singleOwner,
                initialThreshold: 1
            })
        });
        deployImpl({
            name: "communityMultisig",
            deployedTo: deployMultisig({
                initialOwners: singleOwner,
                initialThreshold: 1
            })
        });

        // deploy primary executorMultisig
        require(address(Env.timelockController()) != address(0),
            "must deploy timelockController before executorMultisig");
        address[] memory owners_executorMultisig = new address[](2);
        owners_executorMultisig[0] = address(Env.timelockController());
        owners_executorMultisig[1] = Env.communityMultisig();
        deployImpl({
            name: "executorMultisig",
            deployedTo: deployMultisig({
                initialOwners: owners_executorMultisig,
                initialThreshold: 1
            })
        });

        // deploy beigenExecutorMultisig
        require(address(Env.beigenTimelockController()) != address(0),
            "must deploy beigenTokenTimelockController before beigenExecutorMultisig");
        address[] memory owners_beigenExecutorMultisig = new address[](2);
        owners_beigenExecutorMultisig[0] = address(Env.beigenTimelockController());
        owners_beigenExecutorMultisig[1] = Env.communityMultisig();
        deployImpl({
            name: "beigenExecutorMultisig",
            deployedTo: deployMultisig({
                initialOwners: owners_beigenExecutorMultisig,
                initialThreshold: 1
            })
        });
    }

    function deployMultisig(address[] memory initialOwners, uint256 initialThreshold) public returns (address) {
        // TODO: solution for local networks / those that do not have Safe deployed on them?
        // addresses taken from https://github.com/safe-global/safe-smart-account/blob/main/CHANGELOG.md#expected-addresses-with-deterministic-deployment-proxy-default
        address safeFactory = 0xa6B71E26C5e0845f74c812102Ca7114b6a896AB2;
        address safeSingleton = 0xd9Db270c1B5E3Bd161E8c8503c55cEABeE709552;
        address safeFallbackHandler = 0xf48f2B2d2a534e402487b3ee7C18c33Aec0Fe5e4;

        bytes memory emptyData;
        // TODO: is implementing a nonzero salt useful at all? if yes, this should be an input
        uint256 salt = 0;

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

        bytes memory calldataToFactory = abi.encodeWithSignature(
            "createProxyWithNonce(address,bytes,uint256)",
            safeSingleton,
            initializerData,
            salt
        );

        (bool success, bytes memory returndata) = safeFactory.call(calldataToFactory);
        require(success, "multisig deployment failed");
        address deployedMultisig = abi.decode(returndata, (address));
        require(deployedMultisig != address(0), "something wrong in multisig deployment, zero address returned");
        return deployedMultisig;
    }

    function configureTimelockController(TimelockController timelockController) public {
        uint256 tx_array_length = 10;
        address[] memory targets = new address[](tx_array_length);
        for (uint256 i = 0; i < targets.length; ++i) {
            targets[i] = address(timelockController);
        }

        uint256[] memory values = new uint256[](tx_array_length);

        bytes[] memory payloads = new bytes[](tx_array_length);
        // 1. remove sender as canceller
        payloads[0] = abi.encodeWithSelector(AccessControl.revokeRole.selector, timelockController.CANCELLER_ROLE(), msg.sender);
        // 2. remove sender as executor
        payloads[1] = abi.encodeWithSelector(AccessControl.revokeRole.selector, timelockController.EXECUTOR_ROLE(), msg.sender);
        // 3. remove sender as proposer
        payloads[2] = abi.encodeWithSelector(AccessControl.revokeRole.selector, timelockController.PROPOSER_ROLE(), msg.sender);
        // 4. remove sender as admin
        payloads[3] = abi.encodeWithSelector(AccessControl.revokeRole.selector, timelockController.TIMELOCK_ADMIN_ROLE(), msg.sender);

        // 5. add operationsMultisig as canceller
        payloads[4] = abi.encodeWithSelector(AccessControl.grantRole.selector, timelockController.CANCELLER_ROLE(), Env.opsMultisig());
        // 6. add operationsMultisig as proposer
        payloads[5] = abi.encodeWithSelector(AccessControl.grantRole.selector, timelockController.PROPOSER_ROLE(), Env.opsMultisig());

        // 7. add protocolCouncilMultisig as proposer
        payloads[6] = abi.encodeWithSelector(AccessControl.grantRole.selector, timelockController.PROPOSER_ROLE(), Env.protocolCouncilMultisig());
        // 8. add protocolCouncilMultisig as executor
        payloads[7] = abi.encodeWithSelector(AccessControl.grantRole.selector, timelockController.EXECUTOR_ROLE(), Env.protocolCouncilMultisig());

        // 9. add communityMultisig as admin
        payloads[8] = abi.encodeWithSelector(AccessControl.grantRole.selector, timelockController.TIMELOCK_ADMIN_ROLE(), Env.communityMultisig());

        uint256 delayToSet;
        if (block.chainid == 1) {
            if (timelockController == Env.timelockController()) {
                delayToSet = 10 days;
            } else if (timelockController == Env.beigenTimelockController()) {
                delayToSet = 24 days;
            } else {
                revert("error in setting timelock delay");
            }
        } else {
            delayToSet = 1;
        }
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

    function deployTokensAndStrategy() public {
        deployImpl({
            name: "proxyAdmin",
            deployedTo: address(new ProxyAdmin())
        });
        deployImpl({
            name: "beigenProxyAdmin",
            deployedTo: address(new ProxyAdmin())
        });
        ProxyAdmin proxyAdmin = ProxyAdmin(Env.proxyAdmin());
        ProxyAdmin beigenProxyAdmin = ProxyAdmin(Env.beigenProxyAdmin());

        // TODO: decide if this should be in env
        // placeholder used for initial proxy deployments since initial implementation must be a contract
        EmptyContract emptyContract = new EmptyContract();

        deployProxy({
            name: type(BackingEigen).name,
            deployedTo: address(new TransparentUpgradeableProxy({
                _logic: address(emptyContract),
                admin_: address(beigenProxyAdmin),
                _data: ""
            }))
        });
        deployProxy({
            name: type(Eigen).name,
            deployedTo: address(new TransparentUpgradeableProxy({
                _logic: address(emptyContract),
                admin_: address(proxyAdmin),
                _data: ""
            }))
        });

        deployImpl({
            name: type(BackingEigen).name,
            deployedTo: address(new BackingEigen({
                _EIGEN: IERC20(Env.proxy.eigen())
            }))
        });
        deployImpl({
            name: type(Eigen).name,
            deployedTo: address(new Eigen({
                _bEIGEN: IERC20(Env.proxy.beigen())
            }))
        });

        // use deployer as initial owner, for disabling transfer restrictions prior to transferring ownership
        address initialOwner = msg.sender;
        address[] memory minters;
        uint256[] memory mintingAllowances;
        uint256[] memory mintAllowedAfters;
        proxyAdmin.upgradeAndCall({
            proxy: ITransparentUpgradeableProxy(address(Env.proxy.eigen())),
            implementation: address(Env.impl.eigen()),
            data: abi.encodeWithSelector(
                Eigen.initialize.selector,
                initialOwner,
                minters,
                mintingAllowances,
                mintAllowedAfters
            )
        });
        Eigen(address(Env.proxy.eigen())).disableTransferRestrictions();
        Eigen(address(Env.proxy.eigen())).transferOwnership(Env.executorMultisig());

        // use deployer as initial owner, for editing minting permissions prior to transferring ownership
        proxyAdmin.upgradeAndCall({
            proxy: ITransparentUpgradeableProxy(address(Env.proxy.beigen())),
            implementation: address(Env.impl.beigen()),
            data: abi.encodeWithSelector(
                BackingEigen.initialize.selector,
                initialOwner
            )
        });
        // TODO: get correct minterAddress here!
        BackingEigen(address(Env.proxy.beigen())).setIsMinter({
            minterAddress: address(0),
            newStatus: true
        });
        BackingEigen(address(Env.proxy.beigen())).transferOwnership(Env.beigenExecutorMultisig());

        proxyAdmin.changeProxyAdmin({
            proxy: ITransparentUpgradeableProxy(address(Env.proxy.eigen())),
            newAdmin: address(proxyAdmin)
        });

        proxyAdmin.changeProxyAdmin({
            proxy: ITransparentUpgradeableProxy(address(Env.proxy.beigen())),
            newAdmin: address(proxyAdmin)
        });
    }

    function checkGovernanceConfiguration() public view {
        ProxyAdmin proxyAdmin = ProxyAdmin(Env.proxyAdmin());
        ProxyAdmin beigenProxyAdmin = ProxyAdmin(Env.beigenProxyAdmin());

        assertEq(proxyAdmin.owner(), Env.executorMultisig(),
            "proxyAdmin.owner() != executorMultisig");
        require(address(Env.proxyAdmin()) != address(Env.beigenProxyAdmin()),
            "tokens must have different proxy admins to allow different timelock controllers");
        require(address(Env.timelockController()) != address(Env.beigenTimelockController()),
            "tokens must have different timelock controllers");

        // note that proxy admin owners are different but _token_ owners per se are the same
        assertEq(Ownable(address(Env.proxy.eigen())).owner(), address(Env.executorMultisig()),
            "EIGEN.owner() != executorMultisig");
        assertEq(Ownable(address(Env.proxy.beigen())).owner(), address(Env.executorMultisig()),
            "bEIGEN.owner() != executorMultisig");
        assertEq(proxyAdmin.owner(), address(Env.executorMultisig()),
            "proxyAdmin.owner() != executorMultisig");
        assertEq(beigenProxyAdmin.owner(), address(Env.beigenExecutorMultisig()),
            "beigenProxyAdmin.owner() != beigenExecutorMultisig");

        assertEq(proxyAdmin.getProxyAdmin(ITransparentUpgradeableProxy(payable(address(Env.proxy.eigen())))),
            Env.proxyAdmin(),
            "proxyAdmin is not actually the admin of the EIGEN token");
        assertEq(beigenProxyAdmin.getProxyAdmin(ITransparentUpgradeableProxy(payable(address(Env.proxy.beigen())))),
            Env.beigenProxyAdmin(),
            "beigenProxyAdmin is not actually the admin of the bEIGEN token");

        // check that community multisig and protocol timelock are the owners of the executorMultisig
        checkExecutorMultisigOwnership(Env.executorMultisig(), address(Env.timelockController()));
        // check that community multisig and bEIGEN protocol timelock are the owners of the beigenExecutorMultisig
        checkExecutorMultisigOwnership(Env.beigenExecutorMultisig(), address(Env.beigenTimelockController()));

        checkTimelockControllerConfig(Env.timelockController());
        checkTimelockControllerConfig(Env.beigenTimelockController());

        // TODO: this block commented-out because these contracts aren't deployed yet! move to another script?
        // assertEq(delegationManager.owner(), Env.executorMultisig(),
        //     "delegationManager.owner() != executorMultisig");
        // assertEq(strategyManager.owner(), Env.executorMultisig(),
        //     "strategyManager.owner() != executorMultisig");
        // assertEq(strategyManager.strategyWhitelister(), address(strategyFactory),
        //     "strategyManager.strategyWhitelister() != address(strategyFactory)");
        // assertEq(strategyFactory.owner(), Env.opsMultisig(),
        //     "strategyFactory.owner() != operationsMultisig");
        // assertEq(avsDirectory.owner(), executorMultisig,
        //     "avsDirectory.owner() != executorMultisig");
        // assertEq(rewardsCoordinator.owner(), Env.opsMultisig(),
        //     "rewardsCoordinator.owner() != operationsMultisig");
        // assertEq(eigenLayerPauserReg.unpauser(), Env.executorMultisig(),
        //     "eigenLayerPauserReg.unpauser() != operationsMultisig");
        // require(eigenLayerPauserReg.isPauser(Env.opsMultisig()),
        //     "operationsMultisig does not have pausing permissions");
        // require(eigenLayerPauserReg.isPauser(Env.executorMultisig()),
        //     "executorMultisig does not have pausing permissions");
        // require(eigenLayerPauserReg.isPauser(pauserMultisig),
        //     "pauserMultisig does not have pausing permissions");
        // require(eigenPodBeacon.owner() == Env.executorMultisig(), "eigenPodBeacon: owner not set correctly");
        // require(strategyBeacon.owner() == Env.executorMultisig(), "strategyBeacon: owner not set correctly");
    }

    function checkExecutorMultisigOwnership(address _executorMultisig, address timelockControllerAddress) public view {
        (bool success, bytes memory returndata) = _executorMultisig.staticcall(abi.encodeWithSignature("getOwners()"));
        require(success, "call to _executorMultisig.getOwners() failed");
        address[] memory _executorMultisigOwners = abi.decode(returndata, (address[]));
        require(_executorMultisigOwners.length == 2,
            "executorMultisig owners wrong length");
        bool timelockControllerInOwners;
        bool communityMultisigInOwners;
        for (uint256 i = 0; i < 2; ++i) {
            if (_executorMultisigOwners[i] == address(timelockControllerAddress)) {
                timelockControllerInOwners = true;
            }
            if (_executorMultisigOwners[i] == Env.communityMultisig()) {
                communityMultisigInOwners = true;
            }
        }
        require(timelockControllerInOwners, "timelockControllerAddress not in _executorMultisig owners");
        require(communityMultisigInOwners, "communityMultisig not in _executorMultisig owners");
    }

    function checkTimelockControllerConfig(TimelockController timelockController) public view {
        // check for proposer + executor rights on Protocol Council multisig
        require(timelockController.hasRole(timelockController.PROPOSER_ROLE(), Env.protocolCouncilMultisig()),
            "protocolCouncilMultisig does not have PROPOSER_ROLE on timelockController");
        require(timelockController.hasRole(timelockController.EXECUTOR_ROLE(), Env.protocolCouncilMultisig()),
            "protocolCouncilMultisig does not have EXECUTOR_ROLE on timelockController");

        // check for proposer + canceller rights on ops multisig
        require(timelockController.hasRole(timelockController.PROPOSER_ROLE(), Env.opsMultisig()),
            "operationsMultisig does not have PROPOSER_ROLE on timelockController");
        require(timelockController.hasRole(timelockController.CANCELLER_ROLE(), Env.opsMultisig()),
            "operationsMultisig does not have CANCELLER_ROLE on timelockController");

        // check that community multisig has admin rights
        require(timelockController.hasRole(timelockController.TIMELOCK_ADMIN_ROLE(), Env.communityMultisig()),
            "communityMultisig does not have TIMELOCK_ADMIN_ROLE on timelockController");

        // check for self-administration
        require(timelockController.hasRole(timelockController.TIMELOCK_ADMIN_ROLE(), address(timelockController)),
            "timelockController does not have TIMELOCK_ADMIN_ROLE on itself");

        // check that deployer has no rights
        require(!timelockController.hasRole(timelockController.TIMELOCK_ADMIN_ROLE(), msg.sender),
            "deployer erroenously retains TIMELOCK_ADMIN_ROLE on timelockController");
        require(!timelockController.hasRole(timelockController.PROPOSER_ROLE(), msg.sender),
            "deployer erroenously retains PROPOSER_ROLE on timelockController");
        require(!timelockController.hasRole(timelockController.EXECUTOR_ROLE(), msg.sender),
            "deployer erroenously retains EXECUTOR_ROLE on timelockController");
        require(!timelockController.hasRole(timelockController.CANCELLER_ROLE(), msg.sender),
            "deployer erroenously retains CANCELLER_ROLE on timelockController");
    }
}
