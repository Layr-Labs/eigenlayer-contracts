// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.12;

import "./Env.sol";
import {Encode, MultisigCall} from "zeus-templates/utils/Encode.sol";
import {IPausable} from "src/contracts/interfaces/IPausable.sol";
import {console} from "forge-std/console.sol";

/**
 * @title CoreUpgradeQueueBuilder
 * @notice Provides reusable helpers for constructing multisig upgrade calls.
 * Usage:
 * ```solidity
 * MultisigCall[] storage executorCalls = Encode.newMultisigCalls();
 * executorCalls.upgradeAVSDirectory();
 * ```
 */
library CoreUpgradeQueueBuilder {
    using Env for *;
    using Encode for *;
    using CoreUpgradeQueueBuilder for MultisigCall[];

    /**
     * permissions/
     */
    function upgradePermissionController(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        console.log("Upgrading permission controller");
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.permissionController()),
                impl: address(Env.impl.permissionController())
            })
        });
    }

    function upgradeKeyRegistrar(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.keyRegistrar()),
                impl: address(Env.impl.keyRegistrar())
            })
        });
    }

    /**
     * core/
     */
    function upgradeAllocationManager(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.allocationManager()),
                impl: address(Env.impl.allocationManager())
            })
        });
    }

    function upgradeAVSDirectory(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.avsDirectory()),
                impl: address(Env.impl.avsDirectory())
            })
        });
    }

    function upgradeDelegationManager(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.delegationManager()),
                impl: address(Env.impl.delegationManager())
            })
        });
    }

    function upgradeProtocolRegistry(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.protocolRegistry()),
                impl: address(Env.impl.protocolRegistry())
            })
        });
    }

    function upgradeReleaseManager(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.releaseManager()),
                impl: address(Env.impl.releaseManager())
            })
        });
    }

    function upgradeRewardsCoordinator(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.rewardsCoordinator()),
                impl: address(Env.impl.rewardsCoordinator())
            })
        });
    }

    function upgradeStrategyManager(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.strategyManager()),
                impl: address(Env.impl.strategyManager())
            })
        });
    }

    /**
     * pods/
     */
    function upgradeEigenPodManager(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.eigenPodManager()),
                impl: address(Env.impl.eigenPodManager())
            })
        });
    }

    function upgradeEigenPod(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: address(Env.beacon.eigenPod()),
            data: Encode.upgradeableBeacon.upgradeTo({newImpl: address(Env.impl.eigenPod())})
        });
    }

    /**
     * strategies/
     */
    function upgradeEigenStrategy(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.eigenStrategy()),
                impl: address(Env.impl.eigenStrategy())
            })
        });
    }

    function upgradeStrategyBase(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: address(Env.beacon.strategyBase()),
            data: Encode.upgradeableBeacon.upgradeTo({newImpl: address(Env.impl.strategyBase())})
        });
    }

    function upgradeStrategyBaseTVLLimits(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        uint256 count = Env.instance.strategyBaseTVLLimits_Count();
        for (uint256 i = 0; i < count; i++) {
            address proxyInstance = address(Env.instance.strategyBaseTVLLimits(i));
            calls.append({
                to: Env.proxyAdmin(),
                data: Encode.proxyAdmin.upgrade({proxy: proxyInstance, impl: address(Env.impl.strategyBaseTVLLimits())})
            });
        }
        return calls;
    }

    function upgradeStrategyFactory(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.strategyFactory()),
                impl: address(Env.impl.strategyFactory())
            })
        });
    }

    /**
     * multichain/
     */
    function upgradeBN254CertificateVerifier(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.bn254CertificateVerifier()),
                impl: address(Env.impl.bn254CertificateVerifier())
            })
        });
    }

    function upgradeCrossChainRegistry(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.crossChainRegistry()),
                impl: address(Env.impl.crossChainRegistry())
            })
        });
    }

    function upgradeECDSACertificateVerifier(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.ecdsaCertificateVerifier()),
                impl: address(Env.impl.ecdsaCertificateVerifier())
            })
        });
    }

    function upgradeOperatorTableUpdater(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({
                proxy: address(Env.proxy.operatorTableUpdater()),
                impl: address(Env.impl.operatorTableUpdater())
            })
        });
    }

    /**
     * avs/
     */
    function upgradeTaskMailbox(
        MultisigCall[] storage calls
    ) internal returns (MultisigCall[] storage) {
        return calls.append({
            to: Env.proxyAdmin(),
            data: Encode.proxyAdmin.upgrade({proxy: address(Env.proxy.taskMailbox()), impl: address(Env.impl.taskMailbox())})
        });
    }
}
