package httpUtil

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	. "github.com/gyf841010/github.com/gyf841010/pz-infra/logging"
)

func PostJson(url string, header map[string]string, body interface{}) ([]byte, error) {
	bodyByte, err := json.Marshal(body)
	if err != nil {
		Log.Error("occur error when marshal object to json", WithError(err))
		return nil, err
	}

	if bodyByte != nil {
		Log.Debug("request body", With("requestBody", string(bodyByte)))
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyByte))
	if err != nil {
		Log.Error("occur error when new http request, ", WithError(err))
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		Log.Error("occur error when get response, ", WithError(err))
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, Log.Error("Failed to read response body from HTTP Request", With("url", url), WithError(err))
	}

	return respBody, nil
}

func PostXmlWithCert(url string, body string, cacrtFile, crtFile, keyFile string) ([]byte, error) {
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile(cacrtFile)
	if err != nil {
		Log.Error("Failed to Read Cert File, ", WithError(err))
		return nil, err
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		Log.Error("Failed to Load x509 Key Pair ", WithError(err))
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		Log.Error("occur error when new http request, ", WithError(err))
		return nil, err
	}
	req.Header.Set("Content-Type", "text/xml:charset=UTF-8")

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		Log.Error("occur error when get response, ", WithError(err))
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, Log.Error("Failed to read response body from HTTP Request", With("url", url), WithError(err))
	}

	return respBody, nil
}

func GetJson(url string, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		Log.Error("occur error when new http request, ", WithError(err))
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		Log.Error("occur error when get response, ", WithError(err))
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, Log.Error("Failed to read response body from HTTP Request", With("url", url), WithError(err))
	}

	return respBody, nil
}
