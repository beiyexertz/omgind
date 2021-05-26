#!/bin/bash

BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

echo "sourcing: ${BASH_SOURCE[0]}"

export DEPLOY_DIR="$BASE_DIR"
export PROJECT_DIR=$(dirname "${DEPLOY_DIR}")

export VOLUME_DIR=${PROJECT_DIR}/data
export CONFIG_DIR=${PROJECT_DIR}/data/config

#export CONFIG_FILE=${CONFIG_DIR}/config.dev.txt
export CONFIG_FILE=${DEPLOY_DIR}/config.${CMODE}.txt

