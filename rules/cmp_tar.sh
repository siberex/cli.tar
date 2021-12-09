#!/usr/bin/env bash

set -euo pipefail

ls1=$(mktemp)
ls2=$(mktemp)

tar -tvf "$1" > "$ls1"
tar -tvf "$2" > "$ls2"

if cmp --silent "$ls1" "$ls2"; then
    echo "Success: Compared tar listings are equal"
    rm "$ls1" "$ls2"
else
    echo "Error: Compared tar listings are different"
    rm "$ls1" "$ls2"
    exit 1
fi
