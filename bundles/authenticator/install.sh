#!/bin/sh

go get -u github.com/edwinvautier/go-cli
go mod tidy

echo `
## authenticator-bundle
  # token duration (minutes)
  TOKEN_VALID_DURATION=20
  RSA_PUBLIC_PATH=public.pem
  RSA_PRIVATE_PATH=private.pem
  RSA_PASSWORD=YOURPASSWORDHERE
  DOMAIN=localhost
## authenticator-bundle end
` >> .env.dist
