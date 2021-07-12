#!/usr/bin/env bash

set -xv

wget --ca-certificate ./certs/tls.crt https://localhost:443/test
