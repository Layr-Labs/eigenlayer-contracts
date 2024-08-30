// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "../../utils/ExistingDeploymentParser.sol";

/**
 * anvil --fork-url $RPC_HOLESKY
 * forge script script/deploy/holesky/longtail-preprod-upgrade.s.sol:Longtail_Upgrade_Preprod --rpc-url http://127.0.0.1:8545 --private-key $PRIVATE_KEY --broadcast -vvvv
 * forge script script/deploy/holesky/longtail-preprod-upgrade.s.sol:Longtail_Upgrade_Preprod --rpc-url $RPC_HOLESKY --private-key $PRIVATE_KEY --verify --broadcast -vvvv
 */
contract Longtail_Upgrade_Preprod is ExistingDeploymentParser {

    address testAddress = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;
    address initOwner = 0xDA29BB71669f46F2a779b4b62f03644A84eE3479;

    function run() external virtual {
        _parseInitialDeploymentParams(
            "script/configs/holesky/eigenlayer_preprod.config.json"
        );
        _parseDeployedContracts(
            "script/configs/holesky/eigenlayer_addresses_preprod.config.json"
        );

        emit log_named_address("Deployer Address", msg.sender);

        // START RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.startBroadcast();

        _deployLongtail();
        _upgradeLongtail();

        // STOP RECORDING TRANSACTIONS FOR DEPLOYMENT
        vm.stopBroadcast();

        _sanityChecks();

        logAndOutputContractAddresses("script/output/holesky/longtail-preprod.output.json");
    }

    function _deployLongtail() internal {
        // Deploy implementations
        strategyFactoryImplementation = new StrategyFactory(strategyManager);
        strategyFactoryBeaconImplementation = new StrategyBase(strategyManager);

        // Deploy and initialize proxies
        strategyBeacon = new UpgradeableBeacon(address(strategyFactoryBeaconImplementation));
        strategyFactory = StrategyFactory(address(new TransparentUpgradeableProxy(
            address(strategyFactoryImplementation),
            address(eigenLayerProxyAdmin),
            abi.encodeWithSelector(
                StrategyFactory.initialize.selector,
                initOwner,
                eigenLayerPauserReg,
                0,
                strategyBeacon
            )
        )));
    }

    function _tokensToBlacklist() internal pure returns (IERC20[] memory) {
        // pulled from preprod strategy manager events, filtering by topic:
        // 0x0c35b17d91c96eb2751cd456e1252f42a386e524ef9ff26ecc9950859fdc04fe
        // (https://holesky.etherscan.io/address/0xF9fbF2e35D8803273E214c99BF15174139f4E67a#events)
        IERC20[] memory t = new IERC20[](12);
        t[0] = IERC20(0x1aea86558d3FF59176Fe7D5BE48B59B09c96bbf7); // WETHQ2
        t[1] = IERC20(0xa63f56985F9C7F3bc9fFc5685535649e0C1a55f3); // sfrxETH
        t[2] = IERC20(0x8783C9C904e1bdC87d9168AE703c8481E8a477Fd); // ankrETH
        t[3] = IERC20(0xe3C063B1BEe9de02eb28352b55D49D85514C67FF); // mETH
        t[4] = IERC20(0xF603c5A3F774F05d4D848A9bB139809790890864); // osETH 
        t[5] = IERC20(0x1d8b30cC38Dba8aBce1ac29Ea27d9cFd05379A09); // lsETH
        t[6] = IERC20(0x17845EA6a9BfD2caF1b9E558948BB4999dF2656e); // frxETH
        t[7] = IERC20(0x8720095Fa5739Ab051799211B146a2EEE4Dd8B37); // cbETH
        t[8] = IERC20(0xB4F5fc289a778B80392b86fa70A7111E5bE0F859); // ETHx
        t[9] = IERC20(0x7322c24752f79c05FFD1E2a6FCB97020C1C264F1); // rETH
        t[10] = IERC20(0x3F1c547b21f65e10480dE3ad8E19fAAC46C95034); // stETH
        t[11] = IERC20(0x94373a4919B3240D86eA41593D5eBa789FEF3848); // WETH

        return t;
    }

    function _upgradeLongtail() internal {
        IERC20[] memory tokensToBlacklist = _tokensToBlacklist();

        strategyFactory.blacklistTokens(tokensToBlacklist);
        strategyManager.setStrategyWhitelister(address(strategyFactory));
    }

    function _sanityChecks() internal {
        // Sanity checks

        require(eigenLayerProxyAdmin.getProxyAdmin(TransparentUpgradeableProxy(payable(address(strategyFactory)))) == address(eigenLayerProxyAdmin), "proxy admin not set correctly");
        require(eigenLayerProxyAdmin.getProxyImplementation(TransparentUpgradeableProxy(payable(address(strategyFactory)))) == address(strategyFactoryImplementation), "proxy impl not set correctly");

        require(strategyFactory.owner() == initOwner, "owner not set correctly");
        require(strategyFactory.pauserRegistry() == eigenLayerPauserReg, "pauser not set correctly");
        require(strategyFactory.strategyBeacon() == strategyBeacon, "beacon not set correctly");
        require(strategyFactory.strategyManager() == strategyManager, "strategy manager not set correctly");
        
        require(strategyManager.strategyWhitelister() == address(strategyFactory), "whitelist role not set correctly");
        
        IERC20[] memory tokensToBlacklist = _tokensToBlacklist();
        for (uint i = 0; i < tokensToBlacklist.length; i++) {
            require(strategyFactory.isBlacklisted(tokensToBlacklist[i]), "token not blacklisted");
        }

        // Deploy a strategy and check that it's whitelisted
        IERC20 dummy = IERC20(0x89Aa2dE0beC1b85c1A73111aee7E9A3EE3CBb593);
        IStrategy newStrategy = strategyFactory.deployNewStrategy(dummy);

        require(newStrategy.underlyingToken() == dummy, "underlying not set correctly");
        require(strategyManager.strategyIsWhitelistedForDeposit(newStrategy), "did not whitelist in strategymanager");
    }
}
