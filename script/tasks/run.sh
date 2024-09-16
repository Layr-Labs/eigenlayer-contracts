#!/bin/bash

# These tasks deploy the `slashing-magnitudes` contracts and set up the sender (`address(PRIVATE_KEY)`) as an `AVS`, `Operator` and `Staker`. 
# We then register the `Operator` to an `OperatorSet`, allocate the `strategy` in that `OperatorSet`, deposit `TestTokens` and perform a `slashing`.

# Environment Configuration
RPC_URL="127.0.0.1:8545"
PRIVATE_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
SENDER=$(cast wallet address --private-key $PRIVATE_KEY)
OUTPUT_DIR="../output/local"

# Define amount of shares to deposit/withdraw
DEPOSIT_SHARES=1000

# Ensure output directory exists
mkdir -p $OUTPUT_DIR

# Deploy contracts
forge script -C src/contracts --via-ir ../deploy/local/deploy_from_scratch.slashing.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile)" \
    -- local/deploy_from_scratch.slashing.anvil.config.json

# Compile task contracts
forge build -C script/tasks

# Extract contract addresses from deployment output
DELEGATION_MANAGER=$(jq -r '.addresses.delegationManager' "$OUTPUT_DIR/slashing_output.json")
STRATEGY=$(jq -r '.addresses.strategy' "$OUTPUT_DIR/slashing_output.json")
TOKEN=$(jq -r '.addresses.TestToken' "$OUTPUT_DIR/slashing_output.json")
BALANCE=$(cast call $TOKEN "balanceOf(address)(uint256)" $SENDER --rpc-url $RPC_URL)

# Unpause the AVS Directory
forge script ../tasks/unpause_avsDirectory.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile)" \
    -- local/slashing_output.json

# Deposit shares into strategy
forge script ../tasks/deposit_into_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $DEPOSIT_SHARES

# Register as Operator
forge script ../tasks/register_as_operator.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address operator,string metadataURI)" \
    -- local/slashing_output.json $SENDER "metadataURI"

# Register Operator to OperatorSet
forge script ../tasks/register_operator_to_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --tc RegisterOperatorToOperatorSets \
    --sig "run(string memory configFile)" \
    -- local/slashing_output.json

# Advance the blockchain by 600 blocks to pass pending delay
cast rpc anvil_mine 600 --rpc-url $RPC_URL

# Allocate OperatorSet (50% allocation)
forge script ../tasks/allocate_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address avs,uint32 operatorSetId,uint64 magnitude)" \
    -- local/slashing_output.json $STRATEGY $SENDER 00000001 0500000000000000000

# Slash the OperatorSet (50%)
forge script ../tasks/slash_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address operator,uint32 operatorSetId,uint256 wadToSlash)" \
    -- local/slashing_output.json $SENDER 00000001 0500000000000000000

# Deposit more shares into strategy
forge script ../tasks/deposit_into_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $DEPOSIT_SHARES

# Fetch current withdrawable shares and nonce
DEPOSITS=$(cast call $DELEGATION_MANAGER "getDepositedShares(address)(address[],uint256[])" $SENDER "[$STRATEGY]" --rpc-url $RPC_URL | sed -n '2p' | tr -d '[]')
SHARES=$(cast call $DELEGATION_MANAGER "getWithdrawableShares(address,address[])(uint256[],uint256[])" $SENDER "[$STRATEGY]" --rpc-url $RPC_URL | sed -n '1p' | tr -d '[]')
NONCE=$(cast call $DELEGATION_MANAGER "cumulativeWithdrawalsQueued(address)(uint256)" $SENDER --rpc-url $RPC_URL)

# Withdraw slashed shares from Delegation Manager
forge script ../tasks/withdraw_from_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $DEPOSITS

# Capture block number after initiating withdrawal
WITHDRAWAL_START_BLOCK_NUMBER=$(cast block-number --rpc-url $RPC_URL)

# Advance the blockchain by 5 blocks to meet `MIN_WITHDRAWAL_DELAY_BLOCKS`
cast rpc anvil_mine 5 --rpc-url $RPC_URL

# Slash the OperatorSet (50%)
forge script ../tasks/slash_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address operator,uint32 operatorSetId,uint256 wadToSlash)" \
    -- local/slashing_output.json $SENDER 00000001 0500000000000000000

# Complete the withdrawal process
forge script ../tasks/complete_withdrawal_from_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount,uint256 nonce,uint32 startBlock)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $SHARES $NONCE $WITHDRAWAL_START_BLOCK_NUMBER

# Verification
FINAL_SHARES=$(cast call $DELEGATION_MANAGER "getWithdrawableShares(address,address[])(uint256[],uint256[])" $SENDER "[$STRATEGY]" --rpc-url $RPC_URL | sed -n '1p' | tr -d '[]')
FINAL_BALANCE=$(cast call $TOKEN "balanceOf(address)(uint256)" $SENDER --rpc-url $RPC_URL)
BALANCE_AS_DEC=$(echo "$BALANCE" | awk '{print $1}')
FINAL_BALANCE_AS_DEC=$(echo "$FINAL_BALANCE" | awk '{print $1}')
SLASHED_BY_BALANCE=$(bc <<< "$BALANCE_AS_DEC - $FINAL_BALANCE_AS_DEC")

# Print details
echo -e "==========================\n"
echo -e "Addresses saved to: $(realpath $(pwd)/../../script/output/local/slashing_output.json)"
echo -e "\n==========================\n"
echo "Number of tokens held initially: $BALANCE"
echo "Number of shares deposited: $DEPOSITS"
echo "Number of shares after slashing: $SHARES"
echo "Number of shares remaining: $FINAL_SHARES"
echo "Number of tokens held: $FINAL_BALANCE"
echo "Number of tokens slashed: $SLASHED_BY_BALANCE"
echo -e "\n=========================="
