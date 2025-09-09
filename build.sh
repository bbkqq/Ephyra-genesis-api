#!/bin/bash
set -e

RUN_NAME=Ephyra-genesis-api
mkdir -p output/bin
cp script/bootstrap.sh output 2>/dev/null
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}