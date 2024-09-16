# Setup slashing locally

These tasks deploy the `slashing-magnitudes` contracts and set up the sender (`address(PRIVATE_KEY)`) as an `AVS`, `Operator` and `Staker`. 

We then register the `Operator` to an `OperatorSet`, allocate the `Strategy` in that `OperatorSet` and perform a `slashing`.

---

1. Start `anvil` in one terminal
```sh
anvil
```

2. Run the full setup in another terminal
```sh
./run.sh
```

OR

2. Deploy contracts in another terminal (this will save addresses to [../output/local/slashing_output.json](../output/local/slashing_output.json))
```sh
export RPC_URL=127.0.0.1:8545
export PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
export SENDER=$(cast wallet address --private-key $PRIVATE_KEY)

mkdir ./script/output/local
forge script -C src/contracts --via-ir ../deploy/local/deploy_from_scratch.slashing.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile)" \
    -- local/deploy_from_scratch.slashing.anvil.config.json
```

3. Build the task scripts
```sh
forge build -C script/tasks
```

4. Extract `DELEGATION_MANAGER`, `STRATEGY_MANAGER`, `STRATEGY` and `TOKEN` addresses from deployment output
```sh
export DELEGATION_MANAGER=$(jq -r '.addresses.delegationManager' "../output/local/slashing_output.json")
export STRATEGY_MANAGER=$(jq -r '.addresses.strategyManager' "../output/local/slashing_output.json")
export STRATEGY=$(jq -r '.addresses.strategy' "../output/local/slashing_output.json")
export TOKEN=$(jq -r '.addresses.TestToken' "../output/local/slashing_output.json")
```

5. Unpause the `avsDirectory`
```sh
forge script ../tasks/unpause_avsDirectory.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile)" \
    -- local/slashing_output.json
```

6. Deposit into `Strategy`
```sh
forge script ../tasks/deposit_into_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN 1000
```

7. Register as `Operator`
```sh
forge script ../tasks/register_as_operator.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address operator,string metadataURI)" \
    -- local/slashing_output.json $SENDER "metadataURI"
```

8. Register `Operator` to `OperatorSet`
```sh
forge script ../tasks/register_operator_to_operatorSet.s.sol \
    --tc RegisterOperatorToOperatorSets \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile)" \
    -- local/slashing_output.json
```

9. Move the chain by **600** blocks (to move beyond `pendingDelay`)
```
cast rpc anvil_mine 600 --rpc-url $RPC_URL
```

10. Allocate the `OperatorSet` **(50%)**
```sh
forge script ../tasks/allocate_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address avs,uint32 operatorSetId,uint64 magnitude)" \
    -- local/slashing_output.json $STRATEGY $SENDER 00000001 0500000000000000000
```

11. Slash the `OperatorSet` **(50%)** - we expect that 25% of our shares will be slashed when we withdraw them
```sh
forge script ../tasks/slash_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address operator,uint32 operatorSetId,uint256 wadToSlash)" \
    -- local/slashing_output.json $SENDER 00000001 0500000000000000000
```

12. Verify that the sender holds **1000** Deposited `TOKEN` shares:
```sh
cast call $STRATEGY_MANAGER "getDeposits(address)(address[],uint256[])" $SENDER  --rpc-url $RPC_URL
```

13. Verify that the sender holds **750** Withdrawable `TOKEN` shares:
```sh
cast call $DELEGATION_MANAGER "getWithdrawableShares(address,address[])(uint256[])" $SENDER "[$STRATEGY]" --rpc-url $RPC_URL
```

14. Withdraw slashed shares from `DelegationManager`

- Extract Nonce and available shares from $DELEGATION_MANAGER
```sh
export DEPOSITS=$(cast call $DELEGATION_MANAGER "getDepositedShares(address)(address[],uint256[])" $SENDER "[$STRATEGY]" --rpc-url $RPC_URL | sed -n '2p' | tr -d '[]')
export SHARES=$(cast call $DELEGATION_MANAGER "getWithdrawableShares(address,address[])(uint256[],uint256[])" $SENDER "[$STRATEGY]" --rpc-url $RPC_URL | sed -n '1p' | tr -d '[]')
export NONCE=$(cast call $DELEGATION_MANAGER "cumulativeWithdrawalsQueued(address)(uint256)" $SENDER --rpc-url $RPC_URL)
```

- Queue withdrawal
```sh
forge script ../tasks/withdraw_from_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $DEPOSITS
```

- Record the withdrawal `START_BLOCK`
```sh
export WITHDRAWAL_START_BLOCK_NUMBER=$(cast block-number --rpc-url $RPC_URL)
```

- Move the chain by 5 blocks (to move beyond `MIN_WITHDRAWAL_DELAY_BLOCKS` (5))
```sh
cast rpc anvil_mine 5 --rpc-url $RPC_URL
```

- Complete withdrawal

```sh
forge script ../tasks/complete_withdrawal_from_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount,uint256 nonce,uint32 startBlock)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $SHARES $NONCE $WITHDRAWAL_START_BLOCK_NUMBER
```

15. Verify that the `SHARES` we're withdrawn back to the sender
```sh
cast call $TOKEN "balanceOf(address)(uint256)" $SENDER --rpc-url $RPC_URL
```

16. Verify that the sender holds 0 withdrawable `TOKEN` shares:
```sh
cast call $DELEGATION_MANAGER "getWithdrawableShares(address,address[])(uint256[])" $SENDER "[$STRATEGY]" --rpc-url $RPC_URL
```
