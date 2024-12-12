// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import {EOADeployer} from "zeus-templates/templates/EOADeployer.sol";
import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "zeus-templates/utils/Encode.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
// import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "script/releases/Env.sol";

// import "src/contracts/strategies/EigenStrategy.sol";
import "src/contracts/strategies/StrategyFactory.sol";
import "forge-std/Script.sol";

/**
    anvil --fork-url $RPC_HOLESKY
    forge script script/deploy/devnet/preprod_strategy_hotfix/1-deployNewEigenStrategy.s.sol:EigenStrategyPreprodHotfix \
        --broadcast \
        --rpc-url http://127.0.0.1:8545 \
        --private-key $PRIVATE_KEY -vvvv

    forge script script/deploy/devnet/preprod_strategy_hotfix/1-deployNewEigenStrategy.s.sol:EigenStrategyPreprodHotfix \
        --broadcast \
        --rpc-url $RPC_HOLESKY \
        --private-key $PRIVATE_KEY -vvvv \
        --verify
 */
contract EigenStrategyPreprodHotfix is Script {
    using Env for *;
    using Encode for *;

    /// These addresses not found in zeus config so hard coding them here
    address bEIGEN = 0xA72942289a043874249E60469F68f08B8c6ECCe8;
    address EIGEN = 0xD58f6844f79eB1fbd9f7091d05f7cb30d3363926;

    function run() public {
        // EigenStrategy to remove from whitelist
        EigenStrategy existingEigenStrategyToRemove = Env.proxy.eigenStrategy();
        StrategyFactory factory = Env.proxy.strategyFactory();

        // 1. deploy new EigenStrategy TUPS and implementation
        vm.startBroadcast();
        (
            EigenStrategy eigenStrategy,
            EigenStrategy eigenStrategyImpl
        ) = deployEigenStrategy();
        vm.stopBroadcast();
        // 2. test view functions
        _sanityTest(eigenStrategy, eigenStrategyImpl);
    }

    /**
     * @notice Deploy new strategy with same EigenToken
     */
    function deployEigenStrategy() public returns (
        EigenStrategy eigenStrategy,
        EigenStrategy eigenStrategyImpl
    ) {
        eigenStrategyImpl = new EigenStrategy(
            Env.proxy.strategyManager(),
            Env.impl.pauserRegistry()
        );

        eigenStrategy = EigenStrategy(
            address(
                new TransparentUpgradeableProxy(
                    address(eigenStrategyImpl),
                    address(Env.proxyAdmin()),
                    abi.encodeWithSelector(
                        EigenStrategy.initialize.selector,
                        IEigen(EIGEN),
                        IERC20(bEIGEN)
                    )
                )
            )
        );
    }

    function _sanityTest(
        EigenStrategy eigenStrategy,
        EigenStrategy eigenStrategyImpl
    ) internal {
        require(
            eigenStrategy.underlyingToken() == IERC20(bEIGEN),
            "underlyingToken not set"
        );
        require(
            eigenStrategy.EIGEN() == IEigen(EIGEN),
            "eigenToken not set"
        );
        // proxy points to implementation address
        require(
            IProxyAdmin(Env.proxyAdmin()).getProxyImplementation(
                ITransparentUpgradeableProxy(address(eigenStrategy))
            ) == address(eigenStrategyImpl),
            "proxy admin not set"
        );
    }
}