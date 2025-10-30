// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {DeployGovernance} from "./1-deployGovernance.s.sol";
import {DeployPauser} from "./2-deployPauser.s.sol";
import "../Env.sol";

/// @dev This script is used to deploy the token contracts on a testnet environment.
contract DeployToken is DeployPauser {
    using Env for *;

    function _runAsEOA() internal virtual override {
        vm.startBroadcast();

        // 0. Deploy the empty contract
        address emptyContract = address(new EmptyContract());

        // 1. Deploy the bEIGEN and normal proxy admins
        ProxyAdmin beigenProxyAdmin = new ProxyAdmin();
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        zUpdate("beigenProxyAdmin", address(beigenProxyAdmin));
        zUpdate("proxyAdmin", address(proxyAdmin));

        // 2. Deploy the bEIGEN and Eigen proxies
        deployProxy({
            name: type(BackingEigen).name,
            deployedTo: address(
                new TransparentUpgradeableProxy({
                    _logic: address(emptyContract),
                    admin_: address(beigenProxyAdmin),
                    _data: ""
                })
            )
        });

        deployProxy({
            name: type(Eigen).name,
            deployedTo: address(
                new TransparentUpgradeableProxy({_logic: address(emptyContract), admin_: address(proxyAdmin), _data: ""})
            )
        });

        // 3. Deploy the bEIGEN and Eigen implementations
        deployImpl({
            name: type(BackingEigen).name,
            deployedTo: address(new BackingEigen({_EIGEN: IERC20(address(Env.proxy.eigen()))}))
        });
        deployImpl({
            name: type(Eigen).name,
            deployedTo: address(new Eigen({_bEIGEN: IERC20(address(Env.proxy.beigen())), _version: Env.deployVersion()}))
        });

        // 4. Upgrade the EIGEN Proxy to point to the implementation
        address initialOwner = msg.sender;
        address[] memory minters;
        uint256[] memory mintingAllowances;
        uint256[] memory mintAllowedAfters;
        proxyAdmin.upgradeAndCall({
            proxy: ITransparentUpgradeableProxy(address(Env.proxy.eigen())),
            implementation: address(Env.impl.eigen()),
            data: abi.encodeWithSelector(
                Eigen.initialize.selector, initialOwner, minters, mintingAllowances, mintAllowedAfters
            )
        });

        // Disable transfer restrictions and transfer ownership
        Env.proxy.eigen().disableTransferRestrictions();
        Env.proxy.eigen().transferOwnership(Env.executorMultisig());

        // 5. Upgrade the bEIGEN Proxy to point to the implementation
        beigenProxyAdmin.upgradeAndCall({
            proxy: ITransparentUpgradeableProxy(address(Env.proxy.beigen())),
            implementation: address(Env.impl.beigen()),
            data: abi.encodeWithSelector(BackingEigen.initialize.selector, initialOwner)
        });

        // Set minter, disable transfer restrictions and transfer ownership
        Env.proxy.beigen().setIsMinter({minterAddress: TESTNET_OWNER, newStatus: true});
        Env.proxy.beigen().disableTransferRestrictions();
        Env.proxy.beigen().transferOwnership(Env.executorMultisig());

        // 6. Transfer ownership of the proxy admins to the executor multisig
        proxyAdmin.transferOwnership(Env.executorMultisig());
        beigenProxyAdmin.transferOwnership(Env.beigenExecutorMultisig());

        vm.stopBroadcast();
    }

    function testScript() public virtual override {
        // Deploy older contracts. We have to manually set the EOA mode so we don't revert
        _mode = OperationalMode.EOA;
        DeployGovernance._runAsEOA();
        DeployPauser._runAsEOA();

        // Run the deploy token script
        runAsEOA();
        Eigen eigen = Env.proxy.eigen();
        BackingEigen beigen = Env.proxy.beigen();

        // Check proxyAdmin owner
        assertEq(ProxyAdmin(address(Env.proxyAdmin())).owner(), Env.executorMultisig());
        assertEq(ProxyAdmin(address(Env.beigenProxyAdmin())).owner(), Env.beigenExecutorMultisig());

        // Check the proxy impl and proxy admin for EIGEn
        assertEq(Env._getProxyImpl(address(eigen)), address(Env.impl.eigen()));
        assertEq(Env._getProxyAdmin(address(eigen)), address(Env.proxyAdmin()));

        // Check the proxy impl and proxy admin for bEIGEN - we can't use the Env helpers here because the proxy admin is a different contract
        ProxyAdmin beigenProxyAdmin = ProxyAdmin(Env.beigenProxyAdmin());
        assertEq(
            beigenProxyAdmin.getProxyImplementation(ITransparentUpgradeableProxy(address(beigen))),
            address(Env.impl.beigen())
        );
        assertEq(
            beigenProxyAdmin.getProxyAdmin(ITransparentUpgradeableProxy(address(beigen))), address(beigenProxyAdmin)
        );

        // Assert that transfer restrictions are disabled
        assertEq(beigen.transferRestrictionsDisabledAfter(), 0);
        assertEq(eigen.transferRestrictionsDisabledAfter(), 0);

        // Assert that the owner of the contracts is the executor multisig
        assertEq(Ownable(address(beigen)).owner(), Env.executorMultisig());
        assertEq(Ownable(address(eigen)).owner(), Env.executorMultisig());

        // Assert that the minter of bEIGEN is the testnet owner
        assertEq(beigen.isMinter(TESTNET_OWNER), true);

        // Check EIGEN immutables + storage
        assertEq(address(eigen.bEIGEN()), address(beigen));
        assertTrue(Env._strEq(Env.proxy.eigen().version(), Env.deployVersion()));

        // Check bEIGEN immutables + storage
        assertEq(address(beigen.EIGEN()), address(eigen));
    }
}
