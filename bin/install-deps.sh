#!/usr/bin/env bash

OS=$(uname -s | tr '[:upper:]' '[:lower:]')

if [[ "$OS" == "linux" ]]; then
    sudo apt-get update
    sudo apt-get install -y make curl git software-properties-common jq
    sudo add-apt-repository -y ppa:ethereum/ethereum
    sudo apt-get update
    sudo apt-get install ethereum=1.14.3+build29843+noble
    curl -L https://foundry.paradigm.xyz | bash
elif [[ "$OS" == "darwin" ]]; then
    brew tap ethereum/ethereum
    brew install libusb ethereum@1.14.3
    curl -L https://foundry.paradigm.xyz | bash
else
    echo "Unsupported OS: $OS"
    exit 1
fi
