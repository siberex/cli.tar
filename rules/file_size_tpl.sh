#!/usr/bin/env bash

set -euo pipefail

INPUT="{INPUT}"

if test -f "$INPUT"; then
    SIZE=$(wc -c < "$INPUT" | awk '{print $1}')
    echo "$SIZE"
else
    echo "Input file does not exist: $INPUT"
    exit 1
fi
