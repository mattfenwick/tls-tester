#!/usr/bin/env bash

set -xv

# use this to see certificate:
# openssl s_client -showcerts -connect localhost:443

curl --cacert ./certs/tls.crt https://localhost:443/test
