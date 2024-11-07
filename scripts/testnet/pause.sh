#!/usr/bin/env bash

# Pause the testnet, by sending SIGSTOP to all the processes of the testnet
# (seednode, observers, validators, proxy, txgen)

export KALYANTESTNETSCRIPTSDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

if [ "$1" == "keep" ]; then
  KEEPOPEN=1
else
  KEEPOPEN=0
fi

source "$KALYANTESTNETSCRIPTSDIR/variables.sh"
source "$KALYANTESTNETSCRIPTSDIR/include/validators.sh"
source "$KALYANTESTNETSCRIPTSDIR/include/observers.sh"
source "$KALYANTESTNETSCRIPTSDIR/include/tools.sh"

if [ $USE_PROXY -eq 1 ]; then
  pauseProxy
fi

if [ $USE_TXGEN -eq 1 ]; then
  pauseTxGen
fi

pauseValidators
pauseObservers
pauseSeednode
