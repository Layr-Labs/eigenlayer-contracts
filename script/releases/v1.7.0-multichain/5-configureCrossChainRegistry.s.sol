// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {DeploySourceChain} from "./1-deploySourceChain.s.sol";
import {CrosschainDeployLib} from "script/releases/CrosschainDeployLib.sol";
import {DeployDestinationChainProxies} from "./2-deployDestinationChainProxies.s.sol";
import "../Env.sol";

/**
 * Purpose: Call the cross chain registry to configure chainIDs
 */
contract ConfigureCrossChainRegistry is DeploySourceChain, DeployDestinationChainProxies {
    using Env for *;

    function _runAsMultisig() internal override prank(Env.opsMultisig()) {
        // If we're not on a source chain, we don't need to run this
        if (!Env.isSourceChain()) {
            return;
        }

        // Update the call to the cross chain registry based on the environment
        address operatorTableUpdater = address(Env.proxy.operatorTableUpdater());
        if (Env._strEq(Env.env(), "preprod")) {
            _callPreprod(operatorTableUpdater);
        } else if (Env._strEq(Env.env(), "testnet-sepolia")) {
            _callTestnet(operatorTableUpdater);
        } else if (Env._strEq(Env.env(), "mainnet")) {
            _callMainnet(operatorTableUpdater);
        }
    }

    function testScript() public virtual override(DeploySourceChain, DeployDestinationChainProxies) {
        if (!Env.isSourceChain()) {
            return;
        }

        // Deploy source chain
        _mode = OperationalMode.EOA; // Set to EOA mode so we can deploy the impls in the EOA script
        DeploySourceChain._runAsEOA();

        // Deploy destination chain
        if (!_areProxiesDeployed()) {
            DeployDestinationChainProxies._runAsMultisig();
            _unsafeResetHasPranked(); // reset hasPranked so we can use it in the execute()
        } else {
            // Since the proxies are already deployed, we need to update the env with the proper addresses
            _addContractsToEnv();
        }

        // Configure the registry
        execute();

        // Validate the registry
        _validateRegistry();
    }

    function _validateRegistry() internal view {
        (uint256[] memory chainIDs, address[] memory operatorTableUpdaters) =
            Env.proxy.crossChainRegistry().getSupportedChains();

        if (Env._strEq(Env.env(), "preprod")) {
            // Preprod should have 1 chain: current chain
            require(chainIDs.length == 1, "Invalid number of chains for preprod");
            require(chainIDs[0] == block.chainid, "Invalid chain ID for preprod");
            require(
                operatorTableUpdaters[0] == address(Env.proxy.operatorTableUpdater()),
                "Invalid operator table updater for preprod"
            );
        } else if (Env._strEq(Env.env(), "testnet-sepolia")) {
            // Testnet should have 2 chains: sepolia and base-sepolia
            require(chainIDs.length == 2, "Invalid number of chains for testnet");

            // Check both chains are present (order may vary)
            bool hasSepoliaChain = false;
            bool hasBaseSepolia = false;

            for (uint256 i = 0; i < chainIDs.length; i++) {
                if (chainIDs[i] == block.chainid) {
                    hasSepoliaChain = true;
                    require(
                        operatorTableUpdaters[i] == address(Env.proxy.operatorTableUpdater()),
                        "Invalid operator table updater for sepolia"
                    );
                } else if (chainIDs[i] == 84_532) {
                    hasBaseSepolia = true;
                    require(
                        operatorTableUpdaters[i] == address(Env.proxy.operatorTableUpdater()),
                        "Invalid operator table updater for base-sepolia"
                    );
                }
            }

            require(hasSepoliaChain && hasBaseSepolia, "Missing expected chains for testnet");
        } else if (Env._strEq(Env.env(), "mainnet")) {
            // Mainnet should have 2 chains: mainnet and base
            require(chainIDs.length == 2, "Invalid number of chains for mainnet");

            // Check both chains are present (order may vary)
            bool hasMainnetChain = false;
            bool hasBase = false;

            for (uint256 i = 0; i < chainIDs.length; i++) {
                if (chainIDs[i] == block.chainid) {
                    hasMainnetChain = true;
                    require(
                        operatorTableUpdaters[i] == address(Env.proxy.operatorTableUpdater()),
                        "Invalid operator table updater for mainnet"
                    );
                } else if (chainIDs[i] == 8453) {
                    hasBase = true;
                    require(
                        operatorTableUpdaters[i] == address(Env.proxy.operatorTableUpdater()),
                        "Invalid operator table updater for base"
                    );
                }
            }

            require(hasMainnetChain && hasBase, "Missing expected chains for mainnet");
        }
    }

    /// @dev On preprod, the source and destination chains are the same
    function _callPreprod(
        address operatorTableUpdater
    ) internal {
        uint256[] memory chainIDs = new uint256[](1);
        address[] memory operatorTableUpdaters = new address[](1);

        chainIDs[0] = block.chainid;
        operatorTableUpdaters[0] = operatorTableUpdater;

        Env.proxy.crossChainRegistry().addChainIDsToWhitelist(chainIDs, operatorTableUpdaters);
    }

    function _callTestnet(
        address operatorTableUpdater
    ) internal {
        uint256[] memory chainIDs = new uint256[](2);
        address[] memory operatorTableUpdaters = new address[](2);

        chainIDs[0] = block.chainid; // sepolia
        operatorTableUpdaters[0] = operatorTableUpdater;
        chainIDs[1] = 84_532; // base-sepolia
        operatorTableUpdaters[1] = operatorTableUpdater;

        Env.proxy.crossChainRegistry().addChainIDsToWhitelist(chainIDs, operatorTableUpdaters);
    }

    function _callMainnet(
        address operatorTableUpdater
    ) internal {
        uint256[] memory chainIDs = new uint256[](2);
        address[] memory operatorTableUpdaters = new address[](2);

        chainIDs[0] = block.chainid; // mainnet
        operatorTableUpdaters[0] = operatorTableUpdater;
        chainIDs[1] = 8453; // base
        operatorTableUpdaters[1] = operatorTableUpdater;

        Env.proxy.crossChainRegistry().addChainIDsToWhitelist(chainIDs, operatorTableUpdaters);
    }
}
