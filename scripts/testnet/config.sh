#!/usr/bin/env bash

# Prepare the configuration files for a testnet:
#		Build the k-chain-deploy-go binary and run it
#		Copy the generated files to the testnet folder
#		Copy configuration for the seednode, nodes, proxy and txgen into the testnet folder

export KALYANTESTNETSCRIPTSDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

source "$KALYANTESTNETSCRIPTSDIR/variables.sh"
source "$KALYANTESTNETSCRIPTSDIR/include/config.sh"
source "$KALYANTESTNETSCRIPTSDIR/include/build.sh"

prepareFolders

buildConfigGenerator

generateConfig

copyConfig

copySeednodeConfig
updateSeednodeConfig

copyNodeConfig
updateNodeConfig

if [ $USE_PROXY -eq 1 ]; then
  prepareFolders_Proxy
  copyProxyConfig
  updateProxyConfig
fi

if [ $USE_TXGEN -eq 1 ]; then
  prepareFolders_TxGen
  copyTxGenConfig
  updateTxGenConfig
fi

if [ $USE_HARDFORK -eq 1 ]; then
  changeConfigForHardfork
fi

if [ $COPY_BACK_CONFIGS -eq 1 ]; then
  copyBackConfigs
fi
