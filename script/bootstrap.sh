#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=Ephyra-genesis-api
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}