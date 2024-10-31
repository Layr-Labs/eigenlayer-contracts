// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "forge-std/Test.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";

import "src/contracts/permissions/PauserRegistry.sol";
import "src/contracts/strategies/StrategyBase.sol";

import "src/test/mocks/AVSDirectoryMock.sol";
import "src/test/mocks/AllocationManagerMock.sol";
import "src/test/mocks/StrategyManagerMock.sol";
import "src/test/mocks/DelegationManagerMock.sol";
import "src/test/mocks/EigenPodManagerMock.sol";

import "src/test/utils/SingleItemArrayLib.sol";

abstract contract EigenLayerUnitTestSetup is Test {
    using SingleItemArrayLib for *;

    Vm cheats = Vm(VM_ADDRESS);

    address constant pauser = address(555);
    address constant unpauser = address(556);

    PauserRegistry pauserRegistry;
    ProxyAdmin eigenLayerProxyAdmin;

    AVSDirectoryMock avsDirectoryMock;
    AllocationManagerMock allocationManagerMock;
    StrategyManagerMock strategyManagerMock;
    DelegationManagerMock delegationManagerMock;
    EigenPodManagerMock eigenPodManagerMock;

    mapping(address => bool) public isExcludedFuzzAddress;

    modifier filterFuzzedAddressInputs(address addr) {
        cheats.assume(!isExcludedFuzzAddress[addr]);
        _;
    }

    function setUp() public virtual {
        address[] memory pausers = new address[](2);
        pausers[0] = pauser;
        pausers[1] = address(this);

        pauserRegistry = new PauserRegistry(pausers, unpauser);
        eigenLayerProxyAdmin = new ProxyAdmin();

        avsDirectoryMock = AVSDirectoryMock(payable(address(new AVSDirectoryMock())));
        allocationManagerMock = AllocationManagerMock(payable(address(new AllocationManagerMock())));
        strategyManagerMock = StrategyManagerMock(payable(address(new StrategyManagerMock())));
        delegationManagerMock = DelegationManagerMock(payable(address(new DelegationManagerMock())));
        eigenPodManagerMock = EigenPodManagerMock(payable(address(new EigenPodManagerMock(pauserRegistry))));

        isExcludedFuzzAddress[address(0)] = true;
        isExcludedFuzzAddress[address(pauserRegistry)] = true;
        isExcludedFuzzAddress[address(eigenLayerProxyAdmin)] = true;
        isExcludedFuzzAddress[address(avsDirectoryMock)] = true;
        isExcludedFuzzAddress[address(allocationManagerMock)] = true;
        isExcludedFuzzAddress[address(strategyManagerMock)] = true;
        isExcludedFuzzAddress[address(delegationManagerMock)] = true;
        isExcludedFuzzAddress[address(eigenPodManagerMock)] = true;
    }
}