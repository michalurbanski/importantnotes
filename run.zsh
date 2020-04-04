#!/bin/bash
if [[ -z ${IMPORTANTNOTES_FILE} ]]; then
    echo "Please set IMPORTANTNOTES_FILE env variable"
    exit 1
fi

go run main.go -file "${IMPORTANTNOTES_FILE}"
