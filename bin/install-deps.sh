#!/usr/bin/env bash

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -a | tr '[:upper:]' '[:lower:]')

linuxAmd64="https://gethstore.blob.core.windows.net/builds/geth-alltools-linux-amd64-1.14.6-aadddf3a.tar.gz"
linuxArm64="https://gethstore.blob.core.windows.net/builds/geth-alltools-linux-arm64-1.14.5-0dd173a7.tar.gz"


if [[ "$OS" == "linux" ]]; then
    sudo apt-get update
    sudo apt-get install -y make curl git software-properties-common jq

    if [[ $ARCH == *"x86_64"* ]]; then
        curl -L $linuxAmd64 | tar -xz
    elif [[ $ARCH == *"aarch64"* ]]; then
        curl -L $linuxArm64 | tar -xz
    else
        echo "Unsupported architecture: $ARCH"
        exit 1
    fi
elif [[ "$OS" == "darwin" ]]; then
    brew tap ethereum/ethereum
    brew install libusb ethereum@1.14.5
else
    echo "Unsupported OS: $OS"
    exit 1
fi

curl -L https://foundry.paradigm.xyz | bash

cp -R /root/.foundry/bin/* /usr/local/bin/

foundryup

cp -R /root/.foundry/bin/* /usr/local/bin/
