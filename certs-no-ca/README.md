# Certs

Usage:

```bash
./create-new-cert.sh
```

This will create a self-signed tls.crt and tls.key file.

## Reading cert metadata

The cert metadata can be read using:

```bash
openssl x509 -in tls.crt -noout -text
```
