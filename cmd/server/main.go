package main

import (
	"github.com/mattfenwick/tls-tester/pkg"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	stop := make(chan struct{})

	logrus.Infof("args: %+v", os.Args)

	certFile, keyFile := os.Args[1], os.Args[2]

	pkg.RunServer([]int{80, 443}, certFile, keyFile, stop)
}
