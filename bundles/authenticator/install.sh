#!/bin/sh

go get -u github.com/edwinvautier/go-cli
go mod tidy

printf "## authenticator-bundle\n# token duration (minutes)\nTOKEN_VALID_DURATION=20\nRSA_PUBLIC_PATH=public.pem\nRSA_PRIVATE_PATH=private.pem\nRSA_PASSWORD=YOURPASSWORDHERE\nDOMAIN=localhost\n## authenticator-bundle end" >> .env.dist
