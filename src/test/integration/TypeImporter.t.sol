// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "src/contracts/interfaces/IAllocationManager.sol";
import "src/contracts/interfaces/IAVSDirectory.sol";
import "src/contracts/interfaces/IDelegationManager.sol";
import "src/contracts/interfaces/IEigenPod.sol";
import "src/contracts/interfaces/IEigenPodManager.sol";
import "src/contracts/interfaces/IStrategyManager.sol";

/// @dev A master interface contract that imports types defined in our
/// contract interfaces so they can be used without needing to refer to
/// the interface, e.g:
///
/// `AllocateParams memory params;`
/// vs
/// `IAllocationManagerTypes.AllocateParams memory params;`
interface TypeImporter is IAllocationManagerTypes, IAVSDirectoryTypes, IDelegationManagerTypes, IEigenPodManagerTypes, IEigenPodTypes {}
