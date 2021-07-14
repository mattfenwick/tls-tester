#!/usr/bin/env bash

set -xv

# use this to see certificate:
# openssl s_client -showcerts -connect localhost:443

CACERT=${CACERT:-"./certs/ca.crt"}
URL=${URL:-"https://localhost:443/test"}

# start server somewhere else

# run clients
printf "golang client:\n"
go run cmd/client/main.go "$CACERT" "$URL"
printf "\n\ncurl client:\n"
curl --cacert "$CACERT" "$URL"
printf "\n\nwget client:\n"
wget --ca-certificate "$CACERT" "$URL"
