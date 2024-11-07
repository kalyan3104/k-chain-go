#!/usr/bin/env bash

export KALYANTESTNETSCRIPTSDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

$KALYANTESTNETSCRIPTSDIR/stop.sh
$KALYANTESTNETSCRIPTSDIR/clean.sh
$KALYANTESTNETSCRIPTSDIR/config.sh
$KALYANTESTNETSCRIPTSDIR/start.sh $1
