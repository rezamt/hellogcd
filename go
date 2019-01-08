#!/bin/bash

source ./go.vars


if [ -z "$GO_PIPELINE_LABEL" ]; then
    GO_PIPELINE_LABEL="dev"
fi


function check_credentials {
    : "${SAML2AWS_USERNAME:?Need to set SAML2AWS_USERNAME to \\username}"
    : "${SAML2AWS_PASSWORD:?Need to set SAML2AWS_PASSWORD to \\password}"
}


check_credentials


echo GO_PIPELINE_LABEL is: $GO_PIPELINE_LABEL

echo ISO Path is: $ISO_PATH

echo "All validation passed"
