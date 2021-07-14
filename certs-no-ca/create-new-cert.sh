#!/usr/bin/env bash

set -xv

TLS_CERT_PATH=tls.crt
TLS_KEY_PATH=tls.key
KEYSTORE_PATH=keystore.jks
TLS_P12_PATH=tls.p12
KEYSTORE_PASSWORD=${KEYSTORE_PASSWORD:-"tls-tester"}

openssl req -x509 -nodes -days 730 \
  -newkey rsa:2048 \
  -keyout $TLS_KEY_PATH \
  -out $TLS_CERT_PATH \
  -config openssl-config.conf \
  -sha256

openssl pkcs12 -export \
  -in $TLS_CERT_PATH \
  -inkey $TLS_KEY_PATH \
  -name privatekey \
  -out $TLS_P12_PATH \
  -password "pass:$KEYSTORE_PASSWORD"

keytool -importkeystore \
  -srckeystore $TLS_P12_PATH \
  -srcstoretype PKCS12 \
  -srcstorepass "$KEYSTORE_PASSWORD" \
  -srcalias privatekey \
  -destkeystore $KEYSTORE_PATH \
  -deststoretype JKS \
  -deststorepass "$KEYSTORE_PASSWORD" \
  -destalias privatekey \
  -noprompt
