#!/bin/bash

# Authorization Endpoint Request Example
# OAuth 2.0 / OpenID Connect Authorization Code Flow

BASE_URL="http://localhost:8080"

# Required parameters
CLIENT_ID="my-client"
REDIRECT_URI="http://localhost:3000/callback"
RESPONSE_TYPE="code"
SCOPE="openid"
STATE="random-state-for-csrf-protection"

# Optional parameters
NONCE="random-nonce-for-replay-protection"
CODE_CHALLENGE="E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM"  # Base64URL encoded SHA256 hash
CODE_CHALLENGE_METHOD="S256"

# Authenticated user (normally set by reverse proxy/gateway)
USER_ID="user123"

echo "=== Authorization Request ==="
echo ""

# Build the URL
AUTH_URL="${BASE_URL}/authorize?\
client_id=${CLIENT_ID}&\
redirect_uri=${REDIRECT_URI}&\
response_type=${RESPONSE_TYPE}&\
scope=${SCOPE}&\
state=${STATE}&\
nonce=${NONCE}&\
code_challenge=${CODE_CHALLENGE}&\
code_challenge_method=${CODE_CHALLENGE_METHOD}"

echo "URL:"
echo "${AUTH_URL}"
echo ""

# Make the request with curl
echo "=== Request ==="
curl -v -X GET "${AUTH_URL}" \
  -H "X-Authenticated-User: ${USER_ID}"

echo ""
echo ""
echo "=== Expected Response ==="
echo "Success: HTTP 302 Redirect to ${REDIRECT_URI}?code=XXX&state=${STATE}"
echo "Error:   HTTP 400 with JSON { \"error\": \"error_code\" }"
