#!/bin/sh

HOST="heat.csc.uvic.ca"
PORT="443"
CERT_FILE="/app/server_cert.pem"

# Fetch and save the server certificate
openssl s_client -showcerts -connect "${HOST}:${PORT}" </dev/null 2>/dev/null \
  | openssl x509 -outform PEM > "${CERT_FILE}"

if [[ $? -eq 0 ]]; then
  echo "Certificate saved to ${CERT_FILE}"
else
  echo "Failed to fetch certificate"
  exit 1
fi