# BeaconProxyNEW
[Git Source](https://github.com/bowenli86/eigenlayer-contracts/blob/0800603ae0e71de6487dd628cace5380fa364f74/src/test/SigP/BeaconProxy.sol)

**Inherits:**
Proxy, ERC1967Upgrade

*This contract implements a proxy that gets the implementation address for each call from an {UpgradeableBeacon}.
The beacon address is stored in storage slot `uint256(keccak256('eip1967.proxy.beacon')) - 1`, so that it doesn't
conflict with the storage layout of the implementation behind the proxy.
_Available since v3.4._*


## Functions
### constructor

*Initializes the proxy with `beacon`.
If `data` is nonempty, it's used as data in a delegate call to the implementation returned by the beacon. This
will typically be an encoded function call, and allows initializing the storage of the proxy like a Solidity
constructor.
Requirements:
- `beacon` must be a contract with the interface {IBeacon}.*


```solidity
constructor(address beacon, bytes memory data) payable;
```

### _beacon

*Returns the current beacon address.*


```solidity
function _beacon() internal view virtual returns (address);
```

### _implementation

*Returns the current implementation address of the associated beacon.*


```solidity
function _implementation() internal view virtual override returns (address);
```

### _setBeacon

*Changes the proxy to use a new beacon. Deprecated: see {_upgradeBeaconToAndCall}.
If `data` is nonempty, it's used as data in a delegate call to the implementation returned by the beacon.
Requirements:
- `beacon` must be a contract.
- The implementation returned by `beacon` must be a contract.*


```solidity
function _setBeacon(address beacon, bytes memory data) internal virtual;
```

