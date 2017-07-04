package awsloadbalancer

import (
	"encoding/xml"
	"fmt"
	"github.com/scorelab/gocloud-v2/auth"
	awsauth "github.com/scorelab/gocloud-v2/awsauth"
	"io/ioutil"
	"net/http"
	"time"
)

type Awsloadbalancer struct {
}

func (awsloadbalancer *Awsloadbalancer) Creatloadbalancer(request interface{}) (resp interface{}, err error) {

	params := makeParamsWithVersion("CreateLoadBalancer")
	fmt.Println(params)

	return
}

func (awsloadbalancer *Awsloadbalancer) Deleteloadbalancer(request interface{}) (resp interface{}, err error) {
	return
}

func makeParamsWithVersion(action string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	params["Version"] = loadbalancerversion
	return params
}

var timeNow = time.Now

func (awsloadbalancer *Awsloadbalancer) PrepareSignatureV2query(params map[string]string, Region string, resp interface{}) error {

	ElasticloadbalancingEndpoint := "https://elasticloadbalancing.amazonaws.com"

	req, err := http.NewRequest("GET", ElasticloadbalancingEndpoint, nil)
	if err != nil {
		return err
	}

	// Add the params passed in to the query string
	query := req.URL.Query()
	for varName, varVal := range params {
		query.Add(varName, varVal)
	}
	query.Add("Timestamp", timeNow().In(time.UTC).Format(time.RFC3339))

	req.URL.RawQuery = query.Encode()

	auth := map[string]string{"AccessKey": auth.Config.AWSAccessKeyID, "SecretKey": auth.Config.AWSSecretKey}

	awsauth.SignatureV2(req, auth)

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	fmt.Println(string(body))

	return xml.NewDecoder(r.Body).Decode(resp)
}
