#!/bin/bash

# Define the RPC URL
RPC_URL="https://rpc.flashbots.net" # Replace with your actual RPC URL

# List of addresses
declare -a addresses=(
    "0xf1C9acDc66974dFB6dEcB12aA385b9cD01190E38"
    "0xf951E335afb289353dc249e82926178EaC7DEd78"
    "0x856c4Efb76C1D1AE02e20CEB03A2A6a08b0b8dC3"
    "0xA35b1B31Ce002FBF2058D22F30f95D405200A15b"
    "0xa2E3356610840701BDf5611a53974510Ae27E2e1"
    "0xE95A203B1a91a908F9B9CE46459d101078c2c3cb"
    "0xac3E018457B222d93114458476f3E3416Abbe38F"
    "0x8c1bed5b9a0928467c9b1341da1d7bd5e10b6549"
)

# Loop through each address and call the cast command
for address in "${addresses[@]}"
do
    cast call --rpc-url $RPC_URL "$address" "decimals()"
done
