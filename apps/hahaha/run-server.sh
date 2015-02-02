#!/bin/bash

TOP_DIR=$(git rev-parse --show-toplevel)

twistd -n web -p 8080 --path ${TOP_DIR}/apps/hahaha/
