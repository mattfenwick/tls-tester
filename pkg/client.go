package pkg

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func RunClient(rootCaPemPath string, clientPemCertPath string, clientPemKeyPath string) {
	ListInstalledCerts()

	client := resty.New()

	if rootCaPemPath != "" {
		client.SetRootCertificate(rootCaPemPath)
		logrus.Infof("set root cert to %s", rootCaPemPath)
		ListInstalledCerts()
	} else {
		logrus.Warnf("skipping setting root certificate")
	}

	//cert, err := tls.LoadX509KeyPair(clientPemCertPath, clientPemKeyPath)
	//DoOrDie(err)
	//client.SetCertificates(cert)

	client.HostURL = "https://localhost"
	resp, err := IssueRequest(client, "GET", "test", nil, nil)
	DoOrDie(err)
	fmt.Printf("response: %s\n", resp)
}

func IssueRequest(restyClient *resty.Client, verb string, path string, body interface{}, result interface{}) (string, error) {
	var err error
	request := restyClient.R()
	if body != nil {
		reqBody, err := json.MarshalIndent(body, "", "  ")
		if err != nil {
			return "", errors.Wrapf(err, "unable to marshal json")
		}
		logrus.Tracef("request body: %s", string(reqBody))
		request = request.SetBody(body)
	}
	if result != nil {
		request = request.SetResult(result)
	}

	urlPath := fmt.Sprintf("%s/%s", restyClient.HostURL, path)
	logrus.Debugf("issuing %s to %s", verb, urlPath)

	var resp *resty.Response
	switch verb {
	case "GET":
		resp, err = request.Get(path)
	case "POST":
		resp, err = request.Post(path)
	case "PUT":
		resp, err = request.Put(path)
	case "DELETE":
		resp, err = request.Delete(path)
	default:
		return "", errors.Errorf("unrecognized http verb %s to %s", verb, path)
	}
	if err != nil {
		return "", errors.Wrapf(err, "unable to issue %s to %s", verb, path)
	}

	respBody, statusCode := resp.String(), resp.StatusCode()
	logrus.Debugf("response code %d from %s to %s", statusCode, verb, urlPath)
	logrus.Tracef("response body: %s", respBody)

	if !resp.IsSuccess() {
		return respBody, errors.Errorf("bad status code for %s to path %s: %d, response %s", verb, path, statusCode, respBody)
	}
	return respBody, nil
}

func ListInstalledCerts() {
	pool, err := x509.SystemCertPool()
	DoOrDie(err)
	for _, subject := range pool.Subjects() {
		fmt.Printf("subject found: %s\n", subject)
	}
}
