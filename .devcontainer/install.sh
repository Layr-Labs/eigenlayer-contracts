#!/usr/bin/env bash

set -e

NVM_DIR=${HOME}/.nvm
NODE_VERSION=v22.3.0

# Install node
function npm_install {
    mkdir -p ${NVM_DIR}/etc
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash
    bash -c ". ${NVM_DIR}/nvm.sh && nvm install ${NODE_VERSION} && nvm alias default ${NODE_VERSION} && nvm use default"

    NVM_NODE_PATH=${NVM_DIR}/versions/node/${NODE_VERSION}
    NODE_PATH=${NVM_NODE_PATH}/lib/node_modules
    PATH=${NVM_NODE_PATH}/bin:$PATH

    npm install npm -g
    npm install yarn -g
}

# Install foundry
function foundry_install {
    curl -L https://foundry.paradigm.xyz | bash
    ~/.foundry/bin/foundryup
}

npm_install
foundry_install

export NVM_NODE_PATH=${NVM_DIR}/versions/node/${NODE_VERSION}
export NODE_PATH=${NVM_NODE_PATH}/lib/node_modules
export PATH=${PATH}:${NVM_NODE_PATH}/bin: