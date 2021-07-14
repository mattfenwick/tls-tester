# tls-tester

Example project to test certs from client and server sides.

## Usage

Generate certs:

```bash
pushd certs
  ./generate-certs.sh
popd
```

Run server:

```bash
go run cmd/server/main.go \
  certs/end-entity.crt \
  certs/end-entity.key
```

Run client:

```bash
./client-tests.sh
```