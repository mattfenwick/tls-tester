package main

import (
	"github.com/mattfenwick/tls-tester/pkg"
	"github.com/sirupsen/logrus"
)

func main() {
	certPath := "certs/tls.crt"
	//keyPath := "certs/tls.key"

	pkg.ListInstalledCerts()

	logrus.Infof("running resty client")
	pkg.RunRestyClient(certPath)

	logrus.Infof("running vanilla client")
	pkg.RunVanillaClient(certPath)
}
