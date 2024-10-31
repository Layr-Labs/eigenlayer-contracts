# Setup slashing locally

These tasks deploy the `slashing-magnitudes` contracts and set up the sender (`address(PRIVATE_KEY)`) as an `AVS`, `Operator` and `Staker`. We then register the `Operator` to an `OperatorSet`, allocate the `strategy` in that `OperatorSet` and perform a `slashing`.

1. Start anvil in one terminal
```sh
$ anvil
```

2. Deploy contracts in another terminal
```sh
$ export RPC_URL=127.0.0.1:8545
$ export PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
$
$ mkdir ./script/output/local
$ forge script script/deploy/local/deploy_from_scratch.slashing.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(string memory configFile)" -- local/deploy_from_scratch.slashing.anvil.config.json
```

3. Unpause the avsDirectory
```sh
$ forge script script/tasks/unpause_avsDirectory.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address avsDir)" -- 0x5FC8d32690cc91D4c39d9d3abcBD16989F875707
```

4. Deposit into strategy
```sh
$ forge script script/tasks/deposit_into_strategy.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address strategyManager,address strategy,address token,uint256 amount)" -- 0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9 0x8aCd85898458400f7Db866d53FCFF6f0D49741FF 0x67d269191c92Caf3cD7723F116c85e6E9bf55933 1000
```

5. Register as Operator
```sh
$ forge script script/tasks/register_as_operator.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address delegationManager,address operator,string metadataURI)" -- 0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 "metadatURI"
```

6. Register Operator to OperatorSet
```sh
$ forge script script/tasks/register_operator_to_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address avsDir)" -- 0x5FC8d32690cc91D4c39d9d3abcBD16989F875707
```

7. Move the chain by 600 blocks (to move beyond `pendingDelay`)
```
$ cast rpc anvil_mine 600 --rpc-url $RPC_URL
```

8. Allocate the OperatorSet (50%)
```sh
$ forge script script/tasks/allocate_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address allocationManager,address strategy,address operator,uint32 operatorSetId,uint64 magnitude)" -- 0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6 0x8aCd85898458400f7Db866d53FCFF6f0D49741FF 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 00000001 0500000000000000000
```

9. Slash the OperatorSet (50%) - we expect that 25% of our shares will be slashed when we withdraw them
```sh
$ forge script script/tasks/slash_operatorSet.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast --sig "run(address allocationManager,address strategy,address operator,uint32 operatorSetId,uint256 wadToSlash)" -- 0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6 0x8aCd85898458400f7Db866d53FCFF6f0D49741FF 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 00000001 0500000000000000000
```
