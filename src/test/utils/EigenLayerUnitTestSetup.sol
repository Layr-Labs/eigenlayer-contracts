// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "src/contracts/permissions/PauserRegistry.sol";
import "src/contracts/strategies/StrategyBase.sol";
import "src/test/mocks/StrategyManagerMock.sol";
import "src/test/mocks/DelegationManagerMock.sol";
import "src/test/mocks/SlasherMock.sol";
import "src/test/mocks/EigenPodManagerMock.sol";
import "src/test/mocks/StakeRegistryMock.sol";
import "src/test/mocks/Reenterer.sol";
import "src/test/utils/Utils.sol";


import "forge-std/Test.sol";

abstract contract EigenLayerUnitTestSetup is Test, Utils {
    Vm cheats = Vm(HEVM_ADDRESS);

    // Declare Mocks
    StrategyManagerMock strategyManagerMock;
    DelegationManagerMock public delegationManagerMock;
    SlasherMock public slasherMock;
    EigenPodManagerMock public eigenPodManagerMock;
    StakeRegistryMock stakeRegistryMock;
    // strategies
    IERC20 public weth;
    IERC20 public eigenToken;
    StrategyBase public wethStrat;
    StrategyBase public eigenStrat;
    StrategyBase public baseStrategyImplementation;

    PauserRegistry public pauserRegistry;
    ProxyAdmin public eigenLayerProxyAdmin;
    Reenterer public reenterer;

    mapping(address => bool) public addressIsExcludedFromFuzzedInputs;

    address public constant pauser = address(555);
    address public constant unpauser = address(556);
    uint256 wethInitialSupply = 10e50;
    uint256 public constant eigenTotalSupply = 1e50;

    // Helper Functions/Modifiers
    modifier filterFuzzedAddressInputs(address fuzzedAddress) {
        cheats.assume(!addressIsExcludedFromFuzzedInputs[fuzzedAddress]);
        _;
    }

    function setUp() public virtual {
        strategyManagerMock = new StrategyManagerMock();
        delegationManagerMock = new DelegationManagerMock();
        slasherMock = new SlasherMock();
        eigenPodManagerMock = new EigenPodManagerMock();
        eigenLayerProxyAdmin = new ProxyAdmin();
        //deploy pauser registry
        address[] memory pausers = new address[](1);
        pausers[0] = pauser;
        pauserRegistry = new PauserRegistry(pausers, unpauser);

        // deploy upgradeable proxy that points to StrategyBase implementation and initialize it
        baseStrategyImplementation = new StrategyBase(strategyManagerMock);
        weth = new ERC20PresetFixedSupply("weth", "WETH", wethInitialSupply, address(this));
        wethStrat = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(baseStrategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, weth, pauserRegistry)
                )
            )
        );
        eigenToken = new ERC20PresetFixedSupply("eigen", "EIGEN", eigenTotalSupply, address(this));
        eigenStrat = StrategyBase(
            address(
                new TransparentUpgradeableProxy(
                    address(baseStrategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(StrategyBase.initialize.selector, eigenToken, pauserRegistry)
                )
            )
        );
        

        addressIsExcludedFromFuzzedInputs[address(0)] = true;
        addressIsExcludedFromFuzzedInputs[address(strategyManagerMock)] = true;
        addressIsExcludedFromFuzzedInputs[address(delegationManagerMock)] = true;
        addressIsExcludedFromFuzzedInputs[address(slasherMock)] = true;
        addressIsExcludedFromFuzzedInputs[address(eigenPodManagerMock)] = true;
        addressIsExcludedFromFuzzedInputs[address(eigenLayerProxyAdmin)] = true;
    }
}