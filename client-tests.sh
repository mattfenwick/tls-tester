#!/usr/bin/env bash

set -xv

# use this to see certificate:
# openssl s_client -showcerts -connect localhost:443

CACERT=${CACERT:-"./certs/ca.crt"}
HOST=${HOST:-"localhost"}
PORT=${PORT:-"443"}
PATH=${PATH:-"test"}
URL="https://$HOST:$PORT/$PATH"

# start server somewhere else

# TODO there's a problem with this:
#   we get 'Verify return code: 21 (unable to verify the first certificate)'
echo '' | openssl s_client -CApath "$CACERT" -showcerts -connect "$HOST:$PORT"

# run clients
printf "golang client:\n"
go run cmd/client/main.go "$CACERT" "$URL"
printf "\n\ncurl client:\n"
curl --cacert "$CACERT" "$URL"
printf "\n\nwget client:\n"
wget --ca-certificate "$CACERT" "$URL"
