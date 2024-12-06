| Name          | Type                                                                      | Slot | Offset | Bytes | Contract                                                                |
|---------------|---------------------------------------------------------------------------|------|--------|-------|-------------------------------------------------------------------------|
| _initialized  | uint8                                                                     | 0    | 0      | 1     | src/contracts/permissions/PermissionController.sol:PermissionController |
| _initializing | bool                                                                      | 0    | 1      | 1     | src/contracts/permissions/PermissionController.sol:PermissionController |
| _permissions  | mapping(address => struct PermissionControllerStorage.AccountPermissions) | 1    | 0      | 32    | src/contracts/permissions/PermissionController.sol:PermissionController |
| __gap         | uint256[49]                                                               | 2    | 0      | 1568  | src/contracts/permissions/PermissionController.sol:PermissionController |
