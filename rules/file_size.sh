#!/usr/bin/env bash

wc -c < {INPUT} | awk '{print $1}'
