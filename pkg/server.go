package pkg

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mattfenwick/collections/pkg/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func TestServer(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("request: %s", json.MustMarshalToString([]string{r.Host, r.Method, r.Proto}))
	logrus.Infof("headers: %s", json.MustMarshalToString(r.Header))
	fmt.Printf("\n\nheaders: %s\n\n", json.MustMarshalToString(r.Header))

	w.Header().Set("Content-Type", "text/plain")
	_, err := fmt.Fprintf(w, "Success? %s to %s\n", r.Method, r.URL.String())
	DoOrDie(errors.Wrapf(err, "unable to write response"))
}

func RunServerWithTlsTermination(ports []int, certFile string, keyFile string, stop <-chan struct{}) {
	if len(ports) == 0 {
		DoOrDie(errors.Errorf("found 0 ports to run server on"))
	}
	for _, port := range ports {
		logrus.Infof("setting up server on port %d", port)
		serveMux := http.NewServeMux()
		serveMux.HandleFunc("/", TestServer)
		serveMux.HandleFunc("/test", TestServer)
		address := fmt.Sprintf(":%d", port)

		useSimple := false
		if useSimple {
			go func() {
				DoOrDie(errors.Wrapf(http.ListenAndServeTLS(address, certFile, keyFile, serveMux), "unable to ListenAndServeTLS"))
			}()
		} else {
			//cert, err := tls.LoadX509KeyPair(certFile, keyFile)
			//DoOrDie(err)
			server := http.Server{
				Addr:    address,
				Handler: serveMux,
				//TLSConfig: &tls.Config{
				//InsecureSkipVerify: false,
				//Certificates:       []tls.Certificate{cert},
				//RootCAs: rootCAs,
				//},
			}
			go func() {
				DoOrDie(errors.Wrapf(server.ListenAndServeTLS(certFile, keyFile), "unable to ListenAndServeTLS"))
			}()
		}
	}
	<-stop
}

func RunServer(ctx context.Context, port int) {
	logrus.Infof("setting up server on port %d", port)
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", TestServer)
	serveMux.HandleFunc("/test", TestServer)
	address := fmt.Sprintf(":%d", port)

	DoOrDie(errors.Wrapf(http.ListenAndServe(address, serveMux), "unable to ListenAndServe"))
}
