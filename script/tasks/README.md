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
forge script ../deploy/local/deploy_from_scratch.slashing.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile)" \
    -- local/deploy_from_scratch.slashing.anvil.config.json
```

3. Extract `DELEGATION_MANAGER`, `STRATEGY` and `TOKEN` addresses from deployment output
```sh
export DELEGATION_MANAGER=$(jq -r '.addresses.delegationManager' "../output/local/slashing_output.json")
export STRATEGY=$(jq -r '.addresses.strategy' "../output/local/slashing_output.json")
export TOKEN=$(jq -r '.addresses.TestToken' "../output/local/slashing_output.json")
```

4. Unpause the `avsDirectory`
```sh
forge script ../tasks/unpause_avsDirectory.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile)" \
    -- local/slashing_output.json
```

5. Deposit into `Strategy`
```sh
forge script ../tasks/deposit_into_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN 1000
```

6. Register as `Operator`
```sh
forge script ../tasks/register_as_operator.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address operator,string metadataURI)" \
    -- local/slashing_output.json $SENDER "metadataURI"
```

7. Register `Operator` to `OperatorSet`
```sh
forge script ../tasks/register_operator_to_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile)" \
    -- local/slashing_output.json
```

8. Move the chain by **600** blocks (to move beyond `pendingDelay`)
```
cast rpc anvil_mine 600 --rpc-url $RPC_URL
```

9. Allocate the `OperatorSet` **(50%)**
```sh
forge script ../tasks/allocate_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address avs,uint32 operatorSetId,uint64 magnitude)" \
    -- local/slashing_output.json $STRATEGY $SENDER 00000001 0500000000000000000
```

10. Slash the `OperatorSet` **(50%)** - we expect that 25% of our shares will be slashed when we withdraw them
```sh
forge script ../tasks/slash_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address operator,uint32 operatorSetId,uint256 wadToSlash)" \
    -- local/slashing_output.json $STRATEGY $SENDER 00000001 0500000000000000000
```

11. Verify that the sender holds **1000** Deposited `TOKEN` shares:
```sh
cast call 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9 "getDeposits(address)(address[],uint256[])" $SENDER  --rpc-url $RPC_URL
```

12. Verify that the sender holds **750** Withdrawable `TOKEN` shares:
```sh
cast call 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9 "getWithdrawableShares(address,address[])(uint256[])" 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 "[0x8aCd85898458400f7Db866d53FCFF6f0D49741FF]" --rpc-url $RPC_URL
```

13. Withdraw slashed shares from `DelegationManager`

- Extract Nonce and available shares from $DELEGATION_MANAGER
```sh
export SHARES=$(cast call $DELEGATION_MANAGER "getWithdrawableShares(address,address[])(uint256[])" $SENDER "[$STRATEGY]" --rpc-url $RPC_URL | sed 's/[][]//g')
export NONCE=$(cast call $DELEGATION_MANAGER "stakerNonce(address)(uint256)" $SENDER --rpc-url $RPC_URL)
```

- Queue withdrawal
```sh
forge script ../tasks/withdraw_from_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $SHARES
```

- Record the withdrawal `START_BLOCK`
```sh
export WITHDRAWAL_START_BLOCK_NUMBER=$(cast block-number --rpc-url $RPC_URL)
```

- Move the chain by 6 blocks (to move beyond `MIN_WITHDRAWAL_DELAY_BLOCKS` (5))
```sh
cast rpc anvil_mine 6 --rpc-url $RPC_URL
```

- Complete withdrawal

```sh
forge script ../tasks/complete_withdrawal_from_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount,uint256 nonce,uint32 startBlock)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $SHARES $NONCE $WITHDRAWAL_START_BLOCK_NUMBER
```

14. Verify that the `SHARES` we're withdrawn back to the sender
```sh
cast call $TOKEN "balanceOf(address)(uint256)" $SENDER --rpc-url $RPC_URL
```

15. Verify that the sender holds 0 withdrawable `TOKEN` shares:
```sh
cast call $DELEGATION_MANAGER "getWithdrawableShares(address,address[])(uint256[])" $SENDER "[$STRATEGY]" --rpc-url $RPC_URL
```
