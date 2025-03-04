#!/bin/bash

# These tasks deploy the `slashing-magnitudes` contracts and set up the sender (`address(PRIVATE_KEY)`) as an `AVS`, `Operator` and `Staker`.
# We then register the `Operator` to an `OperatorSet`, allocate the `strategy` in that `OperatorSet`, deposit `TestTokens` and perform a `slashing`.

# Environment Configuration
set -a
source .env
export $(grep -v '^#' .env | xargs)
set +a

# Get sender from provided private key
SENDER=$(cast wallet address --private-key $PRIVATE_KEY)     # deposits, delegates, is slashed and withdraws
SENDER_1=$(cast wallet address --private-key $PRIVATE_KEY_1) # deposits, delegates and is slashed
SENDER_2=$(cast wallet address --private-key $PRIVATE_KEY_2) # deposits, delegates, is slashed and undelegates (doesn't complete)
SENDER_3=$(cast wallet address --private-key $PRIVATE_KEY_3) # deposits, is slashed

# Define amount of shares to deposit/withdraw
DEPOSIT_SHARES=1000

# Ensure output directory exists
mkdir -p $OUTPUT_DIR

# Deploy contracts
forge script -C src/contracts --via-ir ../deploy/local/deploy_from_scratch.slashing.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile)" \
    -- local/deploy_from_scratch.slashing.anvil.config.json

# Extract contract addresses from deployment output
DELEGATION_MANAGER=$(jq -r '.addresses.delegationManager' "$OUTPUT_DIR/slashing_output.json")
STRATEGY=$(jq -r '.addresses.strategy' "$OUTPUT_DIR/slashing_output.json")
TOKEN=$(jq -r '.addresses.TestToken' "$OUTPUT_DIR/slashing_output.json")

# Start deploying multicall3 contract
echo -e "==========================\n"
echo -e "## Deploying multicall3.\n"
echo "##### deploy mulitcall3"

# Fund Multicall3 deployer
FUND_OUTPUT=$(cast send $MULTICALL3_DEPLOYER_ADDRESS --value 0.1ether --rpc-url $RPC_URL --private-key $PRIVATE_KEY 2>&1)
FUND_STATUS=$(echo "$FUND_OUTPUT" | grep "status" | awk '{print $2}')
if [[ "$FUND_OUTPUT" == *"error"* ]]; then
    echo "❌  Multicall3 deployer funding failed: $(echo "$FUND_OUTPUT" | grep -o 'error code [^:]*: .*')"
else
    FUND_MESSAGE=$([ "$FUND_STATUS" = "1" ] && echo "success" || echo "failure")
    echo "✅  Multicall3 deployer funded: $FUND_MESSAGE"
fi

# Deploy Multicall3
DEPLOY_OUTPUT=$(cast publish $MULTICALL3_DEPLOYMENT_TX --rpc-url $RPC_URL 2>&1)
DEPLOY_STATUS=$(echo "$DEPLOY_OUTPUT" | jq -r '.status' 2>/dev/null)
if [[ "$DEPLOY_OUTPUT" == *"error"* ]]; then
    echo "❌  Multicall3 deployment failed: $(echo "$DEPLOY_OUTPUT" | grep -o 'error code [^:]*: .*')"
else
    DEPLOY_MESSAGE=$([ "$DEPLOY_STATUS" = "0x1" ] && echo "success" || echo "failure")
    echo "✅  Multicall3 contract deployed: $DEPLOY_MESSAGE"
fi

# End of multicall3 deployment
echo -e "\n==========================\n"

# Fund other senders
FUND_OUTPUT_SENDER_1=$(cast send $TOKEN "transfer(address,uint256)" $SENDER_1 5000 --private-key $PRIVATE_KEY --rpc-url $RPC_URL 2>&1)
FUND_STATUS_SENDER_1=$(echo "$FUND_OUTPUT_SENDER_1" | grep "status" | awk '{print $2}')
FUND_OUTPUT_SENDER_2=$(cast send $TOKEN "transfer(address,uint256)" $SENDER_2 5000 --private-key $PRIVATE_KEY --rpc-url $RPC_URL 2>&1)
FUND_STATUS_SENDER_2=$(echo "$FUND_OUTPUT_SENDER_2" | grep "status" | awk '{print $2}')
FUND_OUTPUT_SENDER_3=$(cast send $TOKEN "transfer(address,uint256)" $SENDER_3 5000 --private-key $PRIVATE_KEY --rpc-url $RPC_URL 2>&1)
FUND_STATUS_SENDER_3=$(echo "$FUND_OUTPUT_SENDER_3" | grep "status" | awk '{print $2}')

# Record initial balance of $SENDER after funding other $SENDER_s
BALANCE=$(cast call $TOKEN "balanceOf(address)(uint256)" $SENDER --rpc-url $RPC_URL)

# Unpause the AVS Directory
forge script -C script/tasks ../tasks/unpause_avsDirectory.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile)" \
    -- local/slashing_output.json

# Deposit shares into strategy
forge script -C script/tasks ../tasks/deposit_into_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $DEPOSIT_SHARES

# Deposit shares into strategy
forge script -C script/tasks ../tasks/deposit_into_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY_1 --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $DEPOSIT_SHARES

# Deposit shares into strategy
forge script -C script/tasks ../tasks/deposit_into_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY_2 --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $DEPOSIT_SHARES

# Deposit shares into strategy
forge script -C script/tasks ../tasks/deposit_into_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY_3 --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $DEPOSIT_SHARES

# Register as Operator
forge script -C script/tasks ../tasks/register_as_operator.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address operator,string metadataURI)" \
    -- local/slashing_output.json $SENDER "metadataURI"

# Register Operator to OperatorSet
forge script -C script/tasks ../tasks/register_operator_to_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --tc RegisterOperatorToOperatorSets \
    --sig "run(string memory configFile)" \
    -- local/slashing_output.json

# Delegate to $SENDERs operator
forge script -C script/tasks ../tasks/delegate_to_operator.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY_1 --broadcast \
    --sig "run(string memory configFile,address operator,uint256 operatorPrivateKey,uint256 saltInt)" \
    -- local/slashing_output.json $SENDER $PRIVATE_KEY 1

# Advance the blockchain by 600 blocks to pass pending delay
cast rpc anvil_mine 600 --rpc-url $RPC_URL

# Allocate OperatorSet (50% allocation)
forge script -C script/tasks ../tasks/allocate_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address avs,uint32 operatorSetId,uint64 magnitude)" \
    -- local/slashing_output.json $STRATEGY $SENDER 00000001 0500000000000000000

# Slash the OperatorSet (50%)
forge script -C script/tasks ../tasks/slash_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address operator,uint32 operatorSetId,address[] strategies,uint256[] wadToSlash)" \
    -- local/slashing_output.json $SENDER 00000001 "[$STRATEGY]" "[0500000000000000000]"

# Deposit more shares into strategy
forge script -C script/tasks ../tasks/deposit_into_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $DEPOSIT_SHARES

# Delegate to $SENDERs operator
forge script -C script/tasks ../tasks/delegate_to_operator.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY_2 --broadcast \
    --sig "run(string memory configFile,address operator,uint256 operatorPrivateKey,uint256 saltInt)" \
    -- local/slashing_output.json $SENDER $PRIVATE_KEY 2

# Undelegate from $SENDERs operator (withdraw as tokens - this will need completing)
forge script -C script/tasks ../tasks/undelegate_from_operator.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY_2 --broadcast \
    --sig "run(string memory configFile)" \
    -- local/slashing_output.json

# Fetch current withdrawable shares and nonce (to complete withdrawal)
DEPOSITS=$(cast call $DELEGATION_MANAGER "getDepositedShares(address)(address[],uint256[])" $SENDER "[$STRATEGY]" --rpc-url $RPC_URL | sed -n '2p' | tr -d '[]')
SHARES=$(cast call $DELEGATION_MANAGER "getWithdrawableShares(address,address[])(uint256[],uint256[])" $SENDER "[$STRATEGY]" --rpc-url $RPC_URL | sed -n '1p' | tr -d '[]')
NONCE=$(cast call $DELEGATION_MANAGER "cumulativeWithdrawalsQueued(address)(uint256)" $SENDER --rpc-url $RPC_URL)
SF=$(cast call $DELEGATION_MANAGER "depositScalingFactor(address,address)(uint256)" $SENDER $STRATEGY --rpc-url $RPC_URL | awk '{print $1}')

# Withdraw slashed shares from Delegation Manager
forge script -C script/tasks ../tasks/withdraw_from_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 amount)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $DEPOSITS

# Capture block number after initiating withdrawal
WITHDRAWAL_START_BLOCK_NUMBER=$(cast block-number --rpc-url $RPC_URL)

# Advance the blockchain by 5 blocks to meet `MIN_WITHDRAWAL_DELAY_BLOCKS`
cast rpc anvil_mine 6 --rpc-url $RPC_URL

# Slash the OperatorSet (50%) - this won't apply to $SENDER who has already successfully undelegated
forge script -C script/tasks ../tasks/slash_operatorSet.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address operator,uint32 operatorSetId,address[] strategies,uint256[] wadToSlash)" \
    -- local/slashing_output.json $SENDER 00000001 "[$STRATEGY]" "[0500000000000000000]"

# Complete the withdrawal process
forge script -C script/tasks ../tasks/complete_withdrawal_from_strategy.s.sol \
    --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast \
    --sig "run(string memory configFile,address strategy,address token,uint256 SF,uint256 amount,uint256 nonce,uint32 startBlock)" \
    -- local/slashing_output.json $STRATEGY $TOKEN $SF $DEPOSITS $NONCE $WITHDRAWAL_START_BLOCK_NUMBER

# Verification
FINAL_SHARES=$(cast call $DELEGATION_MANAGER "getWithdrawableShares(address,address[])(uint256[],uint256[])" $SENDER "[$STRATEGY]" --rpc-url $RPC_URL | sed -n '1p' | tr -d '[]')
FINAL_BALANCE=$(cast call $TOKEN "balanceOf(address)(uint256)" $SENDER --rpc-url $RPC_URL)
BALANCE_AS_DEC=$(echo "$BALANCE" | awk '{print $1}')
FINAL_BALANCE_AS_DEC=$(echo "$FINAL_BALANCE" | awk '{print $1}')
SLASHED_BY_BALANCE=$(bc <<<"$BALANCE_AS_DEC - $FINAL_BALANCE_AS_DEC")

# Print details
echo -e "==========================\n"
echo -e "Addresses saved to: $(realpath $(pwd)/../../script/output/local/slashing_output.json)"
echo -e "\n==========================\n"
echo -e "Spot check sender: $SENDER\n"
echo "Number of shares deposited: $DEPOSITS"
echo "Number of shares after slashing: $SHARES"
echo "Number of shares remaining: $FINAL_SHARES"
echo "Number of shares slashed: $SLASHED_BY_BALANCE"
echo -e "\n=========================="
