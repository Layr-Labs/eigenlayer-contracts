#/bin/bash

source .env

json=$(cat ./script/output/M1_deployment_mainnet_2023_6_9.json)

eigenLayerProxyAdmin=$(echo $json | jq -r '.addresses.eigenLayerProxyAdmin')
eigenLayerPauserReg=$(echo $json | jq -r '.addresses.eigenLayerPauserReg')
slasher=$(echo $json | jq -r '.addresses.slasher')
slasherImplementation=$(echo $json | jq -r '.addresses.slasherImplementation')
strategyManager=$(echo $json | jq -r '.addresses.strategyManager')
strategyManagerImplementation=$(echo $json | jq -r '.addresses.strategyManagerImplementation')
delegation=$(echo $json | jq -r '.addresses.delegation')
delegationImplementation=$(echo $json | jq -r '.addresses.delegationImplementation')
eigenPodManager=$(echo $json | jq -r '.addresses.eigenPodManager')
eigenPodManagerImplementation=$(echo $json | jq -r '.addresses.eigenPodManagerImplementation')
delayedWithdrawalRouter=$(echo $json | jq -r '.addresses.delayedWithdrawalRouter')
delayedWithdrawalRouterImplementation=$(echo $json | jq -r '.addresses.delayedWithdrawalRouterImplementation')
eigenPodBeacon=$(echo $json | jq -r '.addresses.eigenPodBeacon')
eigenPodImplementation=$(echo $json | jq -r '.addresses.eigenPodImplementation')
baseStrategyImplementation=$(echo $json | jq -r '.addresses.baseStrategyImplementation')
cbETHStrategy=$(echo $json | jq -r '.addresses.strategies.cbETH')
stETHStrategy=$(echo $json | jq -r '.addresses.strategies.stETH')
rETHStrategy=$(echo $json | jq -r '.addresses.strategies.rETH')

operationsMultisig=$(echo $json | jq -r '.parameters.operationsMultisig')
executorMultisig=$(echo $json | jq -r '.parameters.executorMultisig')
pauserMultisig=$(echo $json | jq -r '.parameters.pauserMultisig')

# print all addresses
echo "EigenLayer Proxy Admin: $eigenLayerProxyAdmin"
echo "EigenLayer Pauser Registry: $eigenLayerPauserReg"
echo "Slasher: $slasher"
echo "Slasher Implementation: $slasherImplementation"
echo "Strategy Manager: $strategyManager"
echo "Strategy Manager Implementation: $strategyManagerImplementation"
echo "Delegation: $delegation"
echo "Delegation Implementation: $delegationImplementation"
echo "EigenPod Manager: $eigenPodManager"
echo "EigenPod Manager Implementation: $eigenPodManagerImplementation"
echo "Delayed Withdrawal Router: $delayedWithdrawalRouter"
echo "EigenPod Beacon: $eigenPodBeacon"
echo "EigenPod Implementation: $eigenPodImplementation"
echo "Base Strategy Implementation: $baseStrategyImplementation"
echo "cbETH Strategy: $cbETHStrategy"
echo "stETH Strategy: $stETHStrategy"
echo "rETH Strategy: $rETHStrategy"

export ETH_RPC_URL=https://rpc.flashbots.net

echo "EigenLayer Proxy Admin"
cast call $eigenLayerProxyAdmin "owner()" --rpc-url $ETH_RPC_URL

echo "EigenLayer Pauser Registry"
cast call $eigenLayerPauserReg "unpauser()" --rpc-url $ETH_RPC_URL
cast call $eigenLayerPauserReg "isPauser(address)" --rpc-url $ETH_RPC_URL $operationsMultisig
cast call $eigenLayerPauserReg "isPauser(address)" --rpc-url $ETH_RPC_URL $executorMultisig
cast call $eigenLayerPauserReg "isPauser(address)" --rpc-url $ETH_RPC_URL $pauserMultisig

echo "Slasher"
cast call $eigenLayerProxyAdmin "getProxyImplementation(address)" --rpc-url $ETH_RPC_URL $slasher
cast call $slasher "owner()" --rpc-url $ETH_RPC_URL
cast call $slasher "pauserRegistry()" --rpc-url $ETH_RPC_URL
cast call $slasher "paused()" --rpc-url $ETH_RPC_URL

echo "Slasher Implementation"
cast call $slasherImplementation "strategyManager()" --rpc-url $ETH_RPC_URL
cast call $slasherImplementation "delegation()" --rpc-url $ETH_RPC_URL

echo "Delegation"
cast call $eigenLayerProxyAdmin "getProxyImplementation(address)" --rpc-url $ETH_RPC_URL $delegation
cast call $delegation "owner()" --rpc-url $ETH_RPC_URL
cast call $delegation "pauserRegistry()" --rpc-url $ETH_RPC_URL
cast call $delegation "paused()" --rpc-url $ETH_RPC_URL

echo "Delegation Implementation"
cast call $delegationImplementation "strategyManager()" --rpc-url $ETH_RPC_URL
cast call $delegationImplementation "slasher()" --rpc-url $ETH_RPC_URL

echo "Strategy Manager"
cast call $eigenLayerProxyAdmin "getProxyImplementation(address)" --rpc-url $ETH_RPC_URL $strategyManager
cast call $strategyManager "owner()" --rpc-url $ETH_RPC_URL
cast call $strategyManager "pauserRegistry()" --rpc-url $ETH_RPC_URL
cast call $strategyManager "strategyWhitelister()" --rpc-url $ETH_RPC_URL
cast call $strategyManager "withdrawalDelayBlocks()" --rpc-url $ETH_RPC_URL
cast call $strategyManager "paused()" --rpc-url $ETH_RPC_URL

echo "Strategy Manager Implementation"
cast call $strategyManagerImplementation "delegation()" --rpc-url $ETH_RPC_URL
cast call $strategyManagerImplementation "slasher()" --rpc-url $ETH_RPC_URL
cast call $strategyManagerImplementation "eigenPodManager()" --rpc-url $ETH_RPC_URL

echo "EigenPod Manager"
cast call $eigenLayerProxyAdmin "getProxyImplementation(address)" --rpc-url $ETH_RPC_URL $eigenPodManager
cast call $eigenPodManager "owner()" --rpc-url $ETH_RPC_URL
cast call $eigenPodManager "pauserRegistry()" --rpc-url $ETH_RPC_URL
cast call $eigenPodManager "beaconChainOracle()" --rpc-url $ETH_RPC_URL
cast call $eigenPodManager "paused()" --rpc-url $ETH_RPC_URL

echo "EigenPod Manager Implementation"
cast call $eigenPodManagerImplementation "strategyManager()" --rpc-url $ETH_RPC_URL
cast call $eigenPodManagerImplementation "slasher()" --rpc-url $ETH_RPC_URL
cast call $eigenPodManagerImplementation "eigenPodBeacon()" --rpc-url $ETH_RPC_URL
cast call $eigenPodManagerImplementation "ethPOS()" --rpc-url $ETH_RPC_URL

echo "Delayed Withdrawal Router"
cast call $eigenLayerProxyAdmin "getProxyImplementation(address)" --rpc-url $ETH_RPC_URL $delayedWithdrawalRouter
cast call $delayedWithdrawalRouter "owner()" --rpc-url $ETH_RPC_URL
cast call $delayedWithdrawalRouter "pauserRegistry()" --rpc-url $ETH_RPC_URL
cast call $delayedWithdrawalRouter "withdrawalDelayBlocks()" --rpc-url $ETH_RPC_URL
cast call $delayedWithdrawalRouter "paused()" --rpc-url $ETH_RPC_URL

echo "Delayed Withdrawal Router Implementation"
cast call $delayedWithdrawalRouterImplementation "eigenPodManager()" --rpc-url $ETH_RPC_URL

echo "EigenPod Beacon"
cast call $eigenPodBeacon "owner()" --rpc-url $ETH_RPC_URL
cast call $eigenPodBeacon "implementation()" --rpc-url $ETH_RPC_URL

echo "EigenPod Implementation"
cast call $eigenPodImplementation "eigenPodManager()" --rpc-url $ETH_RPC_URL
cast call $eigenPodImplementation "delayedWithdrawalRouter()" --rpc-url $ETH_RPC_URL
cast call $eigenPodImplementation "ethPOS()" --rpc-url $ETH_RPC_URL
cast call $eigenPodImplementation "REQUIRED_BALANCE_WEI()" --rpc-url $ETH_RPC_URL

echo "Base Strategy Implementation"
cast call $baseStrategyImplementation "strategyManager()" --rpc-url $ETH_RPC_URL

echo "rETH Strategy"
cast call $eigenLayerProxyAdmin "getProxyImplementation(address)" --rpc-url $ETH_RPC_URL $rETHStrategy
cast call $rETHStrategy "underlyingToken()" --rpc-url $ETH_RPC_URL
cast call $rETHStrategy "pauserRegistry()" --rpc-url $ETH_RPC_URL
cast call $rETHStrategy "paused()" --rpc-url $ETH_RPC_URL
cast call $rETHStrategy "maxPerDeposit()" --rpc-url $ETH_RPC_URL
cast call $rETHStrategy "maxTotalDeposits()" --rpc-url $ETH_RPC_URL

echo "stETH Strategy"
cast call $eigenLayerProxyAdmin "getProxyImplementation(address)" --rpc-url $ETH_RPC_URL $stETHStrategy
cast call $stETHStrategy "underlyingToken()" --rpc-url $ETH_RPC_URL
cast call $stETHStrategy "pauserRegistry()" --rpc-url $ETH_RPC_URL
cast call $stETHStrategy "paused()" --rpc-url $ETH_RPC_URL
cast call $stETHStrategy "maxPerDeposit()" --rpc-url $ETH_RPC_URL
cast call $stETHStrategy "maxTotalDeposits()" --rpc-url $ETH_RPC_URL

echo "cbETH Strategy"
cast call $eigenLayerProxyAdmin "getProxyImplementation(address)" --rpc-url $ETH_RPC_URL $cbETHStrategy
cast call $cbETHStrategy "underlyingToken()" --rpc-url $ETH_RPC_URL
cast call $cbETHStrategy "pauserRegistry()" --rpc-url $ETH_RPC_URL
cast call $cbETHStrategy "paused()" --rpc-url $ETH_RPC_URL
cast call $cbETHStrategy "maxPerDeposit()" --rpc-url $ETH_RPC_URL
cast call $cbETHStrategy "maxTotalDeposits()" --rpc-url $ETH_RPC_URL
