package alicontainer

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"io"
	"bytes"
	"crypto/x509"
	"crypto/tls"
)

func clusterProjectSignAndDoRequest(method string, path string, clusterID string, options interface{}) (resp interface{}, err error) {

	clusterCerts, err := getClusterCerts(clusterID)
	if err != nil {
		return
	}
	fmt.Println(clusterCerts)

	masterURL, err := getClusterMasterURL(clusterID)
	if err != nil {
		return
	}
	fmt.Println(masterURL)

	certs, err := tls.X509KeyPair([]byte(clusterCerts.Cert), []byte(clusterCerts.Key))
	if err != nil {
		return
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(clusterCerts.CA))

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Certificates:       []tls.Certificate{certs},
				ClientCAs:          caCertPool,
				ClientAuth:         tls.RequireAndVerifyClientCert,
			},
		},
	}

	var reqBody []byte

	// generate request body
	if options != nil {
		reqBody, err = json.Marshal(options)
		if err != nil {
			return
		}
	}
	fmt.Println(string(reqBody))

	// generate request url
	requestURL := masterURL + path
	fmt.Println(requestURL)

	var bodyReader io.Reader
	if reqBody != nil {
		bodyReader = bytes.NewReader(reqBody)
	}
	httpReq, err := http.NewRequest(method, requestURL, bodyReader)
	if err != nil {
		return
	}

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	response := make(map[string]interface{})

	response["body"] = string(body)
	response["status"] = httpResp.StatusCode

	resp = response
	return
}

func getClusterCerts(clusterID string) (certs clusterCerts, err error) {
	response := make(map[string]interface{})
	err = aliauth.ContainerSignAndDoRequest("", http.MethodGet, "/clusters/"+clusterID+"/certs", nil, nil, response)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response["body"].(string)), &certs)
	return
}

func getClusterMasterURL(clusterID string) (masterURL string, err error) {
	response := make(map[string]interface{})
	err = aliauth.ContainerSignAndDoRequest("", http.MethodGet, "/clusters/"+clusterID, nil, nil, response)
	if err != nil {
		return
	}
	var clusterInfo cluster
	err = json.Unmarshal([]byte(response["body"].(string)), &clusterInfo)
	if err != nil {
		return
	}
	masterURL = clusterInfo.MasterURL
	return
}
