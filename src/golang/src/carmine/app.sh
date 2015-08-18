#!/bin/bash

TOP_DIR=$(git rev-parse --show-toplevel)

PROJECT_DIR=${TOP_DIR}/src/golang/src/carmine


# NOTE: The default application YAML configuration must
# be the first in the following list immediately after
# dispatch.yaml.
${TOP_DIR}/bin/goapp serve \
  --host=0.0.0.0 \
  --port=8080 \
  --admin_port=8000 \
  ${PROJECT_DIR}/dispatch.yaml \
  ${PROJECT_DIR}/default/default.yaml
