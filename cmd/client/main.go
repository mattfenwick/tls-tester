package main

import "github.com/mattfenwick/tls-tester/pkg"

func main() {
	rootCaPemPath := "certs/tls.crt"
	clientPemCertPath := "certs/tls.crt"
	clientPemKeyPath := "certs/tls.key"
	pkg.RunClient(rootCaPemPath, clientPemCertPath, clientPemKeyPath)
}
