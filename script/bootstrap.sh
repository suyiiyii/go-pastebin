#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=pastebin
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}