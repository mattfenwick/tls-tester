package pkg

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func RunVanillaClient(rootCaPemPath string, url string) {
	caCert, err := ioutil.ReadFile(rootCaPemPath)
	DoOrDie(err)

	//caCertPool := x509.NewCertPool()

	caCertPool, err := x509.SystemCertPool()
	DoOrDie(err)
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	resp, err := client.Get(url)
	DoOrDie(err)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		DoOrDie(err)
		logrus.Infof("body: %s", string(bodyBytes))
	} else {
		DoOrDie(errors.Errorf("not ok: status code %d", resp.StatusCode))
	}
}

func RunRestyClient(rootCaPemPath string, url string) {
	client := resty.New()

	if rootCaPemPath != "" {
		client.SetRootCertificate(rootCaPemPath)
		logrus.Infof("set root cert to %s", rootCaPemPath)
		//ListInstalledCerts()
	} else {
		logrus.Warnf("skipping setting root certificate")
	}

	resp, err := client.R().Get(url)
	DoOrDie(err)

	respBody, statusCode := resp.String(), resp.StatusCode()
	logrus.Debugf("response code %d; body %s", statusCode, respBody)

	if !resp.IsSuccess() {
		DoOrDie(errors.Errorf("unsuccessful status code: %d", statusCode))
	}
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
