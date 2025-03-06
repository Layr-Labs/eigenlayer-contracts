#!/bin/bash
#
# This script should not be run directly but instead through a link
#

usage="$0 [-hd] [-f '<fun1> <fun2> ...'] [-l <bytes>] [<msg>] [<rule>]"

usage_long="\
Script for verifying contracts.\n\
\n\
 -h|--help        Print this message and exit\n\
 -d|--dev         Use dev mode: send to staging and run on master using
                  certoraRun.py\n\
 -f|--functions   Verify the listed functions only (for parametric rules)\n\
 -l|--hash-length hash bound length in bytes\n\
"

script=$0

scriptDir=$(cd $(dirname $script); pwd)
scriptName=$(basename $0 .sh)
contractName=$(echo $scriptName |sed 's/^verify//g')
module=$(basename $scriptDir)
projectBase=$scriptDir/../../..
confDir=$projectBase/certora/confs/$module

devFlags=""
certoraRun="certoraRun"

while [ $# -gt 0 ]; do
    case $1 in
        -h|--help)
            printf -- "${usage_long}\nIn summary:\n\n  $usage\n\n"
            exit 1
            ;;
        -d|--dev)
            devFlags="--server staging --commit_sha1 d6b1d13a2e01dda0b070d7c12a94f3d4bf27885c" # 7.25.2
            certoraRun=certoraRun.py
            ;;
        -f|--functions)
            methods=$2;
            shift
            ;;
        -l|--hash-length)
            hashLength=$2;
            shift
            ;;
        -*)
            echo "Error: invalid option '$1'"
            exit 1
            ;;
        *)
            break
    esac
    shift;
done

ARGS=

if [[ "$2" ]]; then
    ARGS+="--rule $2"
fi

if [[ $methods ]]; then
    ARGS+=" --method $methods"
fi

if [[ $hashLength ]]; then
    ARGS+=" --hashing_length_bound $hashLength"
fi

(
  cd $projectBase
  $certoraRun $confDir/$contractName.conf $ARGS --msg "$contractName $1 $2" $devFlags
)
