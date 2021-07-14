package main

import "github.com/mattfenwick/tls-tester/pkg"

func main() {
	stop := make(chan struct{})

	certFile := "certs-with-ca/end-entity.crt"
	keyFile := "certs-with-ca/end-entity.key"

	pkg.RunServer([]int{80, 443}, certFile, keyFile, stop)
}
