# Setting up a test environment

If you already have a state with strategies deployed, continue to (Step 2)[#step-2].

## Step 1: Deploy the strategies

In the first terminal

```
anvil --dump-state state.json --gas-limit 18446744073709551615
```

In the second terminal, deploy the core contracts

```
forge script script/deploy/devnet/Deploy_From_Scratch.s.sol  --rpc-url localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast  --skip-simulation --sig "run(string memory configFile)" -- deploy_from_scratch.anvil.config.json
```

After configuring `DeployStrategies.s.sol` with the number of strategies you want to deploy,
```
forge script script/deploy/devnet/operatorSets/DeployStrategies.s.sol:DeployStrategies --rpc-url localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast  --skip-simulation --gas-limit 18446744073709551615
```

Stop the anvil chain and retain the state

## Step 2: Populate the AVSDirectory and StakeRootCompendium

Start the anvil chain with the state

```
anvil --load-state state.json --gas-limit 18446744073709551615
```

First, in order to run a new script you need to add a block to the chain. Simply do that with at transfer:

```
cast send --rpc-url localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --value=0ether 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb9226
```

Then run the script, after configuring `PopulateSRC.sol` with the number of opsets and the number of operators per opset,

```
rm script/output/devnet/populate_src/*

forge script script/deploy/devnet/operatorSets/PopulateSRC.sol:PopulateSRC --rpc-url localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --skip-simulation --broadcast --gas-limit 18446744073709551615
```

Then move on to proving