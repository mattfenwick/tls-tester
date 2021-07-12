package pkg

import "github.com/sirupsen/logrus"

func DoOrDie(err error) {
	if err != nil {
		logrus.Fatalf("%+v", err)
	}
}
