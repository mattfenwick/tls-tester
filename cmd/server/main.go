package main

import (
	"context"
	"os"

	"github.com/mattfenwick/tls-tester/pkg"
	"github.com/sirupsen/logrus"
)

func main() {
	stop := make(chan struct{})

	logrus.Infof("args: %+v", os.Args)

	if len(os.Args) > 1 {
		certFile, keyFile := os.Args[1], os.Args[2]

		pkg.RunServerWithTlsTermination([]int{80, 443}, certFile, keyFile, stop)
	} else {
		pkg.RunServer(context.TODO(), 8081)
	}
}
