package main

import (
	"github.com/mattfenwick/tls-tester/pkg"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.Infof("args: %+v", os.Args)

	certPath, url := os.Args[1], os.Args[2]

	//pkg.ListInstalledCerts()

	logrus.Infof("running resty client")
	pkg.RunRestyClient(certPath, url)

	logrus.Infof("running vanilla client")
	pkg.RunVanillaClient(certPath, url)
}
