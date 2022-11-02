#!/usr/bin/env bash

set -xv
set -euo pipefail

pushd cmd/server
    ./build.sh
popd


kubectl create ns tls-hacking || true

pushd certs-no-ca
  kubectl create secret tls \
    server-tls-hacking-ingress \
    --namespace tls-hacking \
    --cert=tls.crt \
    --key=tls.key \
    -o yaml --dry-run=client | kubectl apply -f -

  kubectl create secret tls \
    reverse-proxy-tls-hacking-ingress \
    --namespace tls-hacking \
    --cert=tls.crt \
    --key=tls.key \
    -o yaml --dry-run=client | kubectl apply -f -
popd

helm upgrade --install my-tls \
  ./charts/tls-hacking \
  --namespace tls-hacking
