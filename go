#!/bin/bash

#!/bin/bash

function check_credentials {
    : "${SAML2AWS_USERNAME:?Need to set SAML2AWS_USERNAME to \\username}"
    : "${SAML2AWS_PASSWORD:?Need to set SAML2AWS_PASSWORD to \\password}"
}


check_credentials
echo "All validation passed"
