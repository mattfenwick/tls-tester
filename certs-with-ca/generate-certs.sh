#!/usr/bin/env bash

set -xv
set -euo pipefail

readonly DAYS=3650

key () {
  openssl genrsa 2048
}

keyid () {
  openssl rsa -pubout | \
    openssl asn1parse -strparse 19 -noout -out - | \
    openssl dgst -sha1 -binary | \
    od -An -tx1 | \
    tr -d ' \n'
}

readonly CA_CONF='
  basicConstraints=critical,CA:TRUE,pathlen:0
  keyUsage=critical,keyCertSign
  subjectKeyIdentifier=hash
'
readonly CA_KEY="$( key )"
readonly CA_KEYID="$( keyid <<< "$CA_KEY" )"

echo "$CA_KEY" > ca.key
echo "$CA_KEYID" > ca_keyid.txt

# make a certificate signing request
openssl req -subj "/CN=$CA_KEYID" -new -key <( echo "$CA_KEY" ) -out "ca.csr"
# make a cert
openssl x509 \
  -req \
  -in "ca.csr" \
  -signkey <( echo "$CA_KEY" ) \
  -out "ca.crt" \
  -days "$DAYS" \
  -set_serial "0x$CA_KEYID" \
  -extfile <( echo "$CA_CONF" )

endEntity () {
  local eeConf="$1"
  local eeKey="$( key )"
  local eeKeyid="$( keyid <<< "$eeKey" )"

  cat <<< "$eeKey" > "end-entity.key"
  chmod 600 "end-entity.key"

  # make an end entity csr
  openssl req -subj "/CN=$eeKeyid" -new -key <( echo "$eeKey" ) -out "end-entity.csr"
  # make an end entity cert
  openssl x509 \
    -req \
    -in "end-entity.csr" \
    -CA "ca.crt" \
    -CAkey <( echo "$CA_KEY" ) \
    -out "end-entity.crt" \
    -days "$DAYS" -set_serial "0x$eeKeyid" -extfile <( echo "$eeConf" )
}

endEntity '
  basicConstraints=critical,CA:FALSE
  keyUsage=critical,digitalSignature,keyEncipherment,keyAgreement
  extendedKeyUsage=serverAuth,clientAuth
  subjectKeyIdentifier=hash
  authorityKeyIdentifier=keyid:always
  subjectAltName=DNS:localhost,IP:127.0.0.1,IP:::1,IP:172.17.0.1
'
