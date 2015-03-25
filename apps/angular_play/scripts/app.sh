#!/bin/bash

TOP_DIR=$(git rev-parse --show-toplevel)

#twistd -n web -p 8080 --path ${TOP_DIR}/apps/angular_play/
${TOP_DIR}/bin/goapp serve --host=0.0.0.0 --port=5000 app.yaml
