// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import "@openzeppelin/contracts/proxy/beacon/IBeacon.sol";

import "../interfaces/IStrategy.sol";
import "../interfaces/IEigenPodManager.sol";
import "../interfaces/IStrategyManager.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IETHPOSDeposit.sol";
import "../interfaces/IEigenPod.sol";

abstract contract EigenPodManagerStorage is IEigenPodManager {
    /**
     *
     *                            CONSTANTS / IMMUTABLES
     *
     */

    /// @notice The ETH2 Deposit Contract
    IETHPOSDeposit public immutable ethPOS;

    /// @notice Beacon proxy to which the EigenPods point
    IBeacon public immutable eigenPodBeacon;

    /// @notice EigenLayer's DelegationManager contract
    IDelegationManager public immutable delegationManager;

    /**
     * @notice Stored code of type(BeaconProxy).creationCode
     * @dev Maintained as a constant to solve an edge case - changes to OpenZeppelin's BeaconProxy code should not cause
     * addresses of EigenPods that are pre-computed with Create2 to change, even upon upgrading this contract, changing compiler version, etc.
     */
    bytes internal constant beaconProxyBytecode =
        hex"608060405260405161090e38038061090e83398101604081905261002291610460565b61002e82826000610035565b505061058a565b61003e83610100565b6040516001600160a01b038416907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e90600090a260008251118061007f5750805b156100fb576100f9836001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e99190610520565b836102a360201b6100291760201c565b505b505050565b610113816102cf60201b6100551760201c565b6101725760405162461bcd60e51b815260206004820152602560248201527f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e6044820152641d1c9858dd60da1b60648201526084015b60405180910390fd5b6101e6816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d79190610520565b6102cf60201b6100551760201c565b61024b5760405162461bcd60e51b815260206004820152603060248201527f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960448201526f1cc81b9bdd08184818dbdb9d1c9858dd60821b6064820152608401610169565b806102827fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5060001b6102de60201b6100641760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606102c883836040518060600160405280602781526020016108e7602791396102e1565b9392505050565b6001600160a01b03163b151590565b90565b6060600080856001600160a01b0316856040516102fe919061053b565b600060405180830381855af49150503d8060008114610339576040519150601f19603f3d011682016040523d82523d6000602084013e61033e565b606091505b5090925090506103508683838761035a565b9695505050505050565b606083156103c65782516103bf576001600160a01b0385163b6103bf5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610169565b50816103d0565b6103d083836103d8565b949350505050565b8151156103e85781518083602001fd5b8060405162461bcd60e51b81526004016101699190610557565b80516001600160a01b038116811461041957600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561044f578181015183820152602001610437565b838111156100f95750506000910152565b6000806040838503121561047357600080fd5b61047c83610402565b60208401519092506001600160401b038082111561049957600080fd5b818501915085601f8301126104ad57600080fd5b8151818111156104bf576104bf61041e565b604051601f8201601f19908116603f011681019083821181831017156104e7576104e761041e565b8160405282815288602084870101111561050057600080fd5b610511836020830160208801610434565b80955050505050509250929050565b60006020828403121561053257600080fd5b6102c882610402565b6000825161054d818460208701610434565b9190910192915050565b6020815260008251806020840152610576816040850160208701610434565b601f01601f19169190910160400192915050565b61034e806105996000396000f3fe60806040523661001357610011610017565b005b6100115b610027610022610067565b610100565b565b606061004e83836040518060600160405280602781526020016102f260279139610124565b9392505050565b6001600160a01b03163b151590565b90565b600061009a7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50546001600160a01b031690565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100fb9190610249565b905090565b3660008037600080366000845af43d6000803e80801561011f573d6000f35b3d6000fd5b6060600080856001600160a01b03168560405161014191906102a2565b600060405180830381855af49150503d806000811461017c576040519150601f19603f3d011682016040523d82523d6000602084013e610181565b606091505b50915091506101928683838761019c565b9695505050505050565b6060831561020d578251610206576001600160a01b0385163b6102065760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064015b60405180910390fd5b5081610217565b610217838361021f565b949350505050565b81511561022f5781518083602001fd5b8060405162461bcd60e51b81526004016101fd91906102be565b60006020828403121561025b57600080fd5b81516001600160a01b038116811461004e57600080fd5b60005b8381101561028d578181015183820152602001610275565b8381111561029c576000848401525b50505050565b600082516102b4818460208701610272565b9190910192915050565b60208152600082518060208401526102dd816040850160208701610272565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220d51e81d3bc5ed20a26aeb05dce7e825c503b2061aa78628027300c8d65b9d89a64736f6c634300080c0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564";

    // @notice Internal constant used in calculations, since the beacon chain stores balances in Gwei rather than wei
    uint256 internal constant GWEI_TO_WEI = 1e9;

    /// @notice Canonical, virtual beacon chain ETH strategy
    IStrategy public constant beaconChainETHStrategy = IStrategy(0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0);

    /**
     *
     *                                STATE VARIABLES
     *
     */

    /// @notice [DEPRECATED] Previously used to query beacon block roots. We now use eip-4788 directly
    address internal __deprecated_beaconChainOracle;

    /// @notice Pod owner to deployed EigenPod address
    mapping(address podOwner => IEigenPod) public ownerToPod;

    // BEGIN STORAGE VARIABLES ADDED AFTER FIRST TESTNET DEPLOYMENT -- DO NOT SUGGEST REORDERING TO CONVENTIONAL ORDER
    /// @notice The number of EigenPods that have been deployed
    uint256 public numPods;

    /// @notice [DEPRECATED] Was initially used to limit growth early on but there is no longer
    /// a maximum number of EigenPods that can be deployed.
    uint256 private __deprecated_maxPods;

    // BEGIN STORAGE VARIABLES ADDED AFTER MAINNET DEPLOYMENT -- DO NOT SUGGEST REORDERING TO CONVENTIONAL ORDER
    /**
     * @notice mapping from pod owner to the deposit shares they have in the virtual beacon chain ETH strategy
     *
     * @dev When an EigenPod registers a balance increase, deposit shares are increased. When registering a balance
     * decrease, however, deposit shares are NOT decreased. Instead, the pod owner's beacon chain slashing factor
     * is decreased proportional to the balance decrease. This impacts the number of shares that will be withdrawn
     * when the deposit shares are queued for withdrawal in the DelegationManager.
     *
     * Note that prior to the slashing release, deposit shares were decreased when balance decreases occurred.
     * In certain cases, a combination of queueing a withdrawal plus registering a balance decrease could result
     * in a staker having negative deposit shares in this mapping. This negative value would be corrected when the
     * staker completes a withdrawal (as tokens or as shares).
     *
     * With the slashing release, negative shares are no longer possible. However, a staker can still have negative
     * shares if they met the conditions for them before the slashing release. If this is the case, that staker
     * should complete any outstanding queued withdrawal in the DelegationManager ("as shares"). This will correct
     * the negative share count and allow the staker to continue using their pod as normal.
     */
    mapping(address podOwner => int256 shares) public podOwnerDepositShares;

    uint64 internal __deprecated_denebForkTimestamp;

    /// @notice Returns the slashing factor applied to the `staker` for the `beaconChainETHStrategy`
    /// Note: this value starts at 1 WAD (1e18) for all stakers, and is updated when a staker's pod registers
    /// a balance decrease.
    mapping(address staker => BeaconChainSlashingFactor) internal _beaconChainSlashingFactor;

    /// @notice Returns the amount of `shares` that have been slashed on EigenLayer but not burned yet.
    uint256 public burnableETHShares;

    constructor(IETHPOSDeposit _ethPOS, IBeacon _eigenPodBeacon, IDelegationManager _delegationManager) {
        ethPOS = _ethPOS;
        eigenPodBeacon = _eigenPodBeacon;
        delegationManager = _delegationManager;
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[42] private __gap;
}
