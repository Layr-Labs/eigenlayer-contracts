#!/bin/bash

source .env 

# Parse command-line arguments using getopt
while getopts ":n:c:a:k:" opt; do
  case $opt in
    n) NETWORK="$OPTARG";;
    c) CONTRACT="$OPTARG";;
    a) ADDRESS="$OPTARG";;
    \?) echo "Invalid option -$OPTARG" >&2; exit 1;;
  esac
done

# Validate that network and contract inputs are provided
if [ -z "$NETWORK" ] || [ -z "$CONTRACT" ] || [ -z "$ADDRESS" ]; then
  echo "Usage: $0 -n <network> -c <contract> -a <address> -k"
  exit 1
fi

# Validate the network input
if [ "$NETWORK" != "mainnet" ] && [ "$NETWORK" != "goerli" ]; then
  echo "Invalid network. Use 'mainnet' or 'goerli'."
  exit 1
fi

# Get local path for contract & validate contract input
case $CONTRACT in
    "strategyManager") CONTRACT_PATH="src/contracts/core/StrategyManager.sol:StrategyManager";;
    "delegation") CONTRACT_PATH="src/contracts/core/DelegationManager.sol:DelegationManager";;
    "eigenPodManager") CONTRACT_PATH="src/contracts/pods/EigenPodManager.sol:EigenPodManager";;
    "eigenPod") CONTRACT_PATH="src/contracts/pods/EigenPod.sol:EigenPod";;
    "slasher") CONTRACT_PATH="src/contracts/core/Slasher.sol:Slasher";;
    *)
    echo "Invalid contract name."
    exit 1
    ;;
esac

# Set RPC
if [ "$NETWORK" == "goerli" ]; then
    RPC_URL="$RPC_GOERLI"
else
    RPC_URL="$RPC_MAINNET"
fi

# Print the selected network and contract
echo "Checking storage layouts for contract on: $NETWORK"
echo "Contract to validate upgrade: $CONTRACT"

# Get storage layout for on-chain contract
echo "Retrieving on-chain storage layout for $ADDRESS"
command="cast storage $ADDRESS --rpc-url $RPC_URL --etherscan-api-key $ETHERSCAN_API_KEY"
eval "$command > /dev/null 2>&1" # precompile contracts so onChainLayout.csv isn't filled with warnings
output=$(eval $command)
echo "$output" | tail -n +2 > onChainLayout.csv
echo "On-chain storage saved to onChainLayout.csv"

# Get storage layout for local contract
echo "Retrieving local storage layout for $CONTRACT at $CONTRACT_PATH"
command="forge inspect $CONTRACT_PATH storage --pretty"
output=$(eval $command)
echo "$output" | tail -n +1 > localLayout.csv
echo "Local storage saved to localLayout.csv"

# Compare the two storage layouts via typescript script
echo "Comparing storage layouts..."

# Add -k operator if present
if [ ! -k "$1" ]; then
  echo "Keeping old storage layout files"
  eval "npx ts-node script/utils/validateStorage/validateStorage.ts --old onChainLayout.csv --new localLayout.csv --keep"
else
  eval "npx ts-node script/utils/validateStorage/validateStorage.ts --old onChainLayout.csv --new localLayout.csv"
fi