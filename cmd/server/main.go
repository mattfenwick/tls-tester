package main

import "github.com/mattfenwick/tls-tester/pkg"

func main() {
	stop := make(chan struct{})
	pkg.RunServer([]int{80, 443}, "certs/tls.crt", "certs/tls.key", stop)
}
