package awsloadbalancer

import (
	"encoding/xml"
	"fmt"
	"github.com/scorelab/gocloud-v2/auth"
	awsauth "github.com/scorelab/gocloud-v2/awsauth"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var timeNow = time.Now

func (awsloadbalancer *Awsloadbalancer) PrepareSignatureV2query(params map[string]string) error {

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

	resp, err := ioutil.ReadAll(r.Body)

	fmt.Println(string(resp))

	return xml.NewDecoder(r.Body).Decode(resp)
}

func prepareAvailabilityZones(params map[string]string, AvailabilityZones []string) {

	for i := range AvailabilityZones {
		n := strconv.Itoa(i + 1)
		prefix := "AvailabilityZones.member." + n
		params[prefix] = AvailabilityZones[i]
	}

}

func prepareSubnets(params map[string]string, Subnets []string) {

	for i := range Subnets {
		n := strconv.Itoa(i + 1)
		prefix := "Subnets.member." + n
		params[prefix] = Subnets[i]
	}

}

func prepareSecurityGroups(params map[string]string, SecurityGroups []string) {

	for i := range SecurityGroups {
		n := strconv.Itoa(i + 1)
		prefix := "SecurityGroups.member." + n
		params[prefix] = SecurityGroups[i]
	}

}

func prepareListeners(params map[string]string, Listeners []Listener) {

	for i := range Listeners {
		n := strconv.Itoa(i + 1)
		prefix := "Listeners.member." + n
		params[prefix+"LoadBalancerPort"] = Listeners[i].LoadBalancerPort
		params[prefix+"InstancePort"] = Listeners[i].InstancePort
		params[prefix+"Protocol"] = Listeners[i].Protocol
		params[prefix+"InstanceProtocol"] = Listeners[i].InstanceProtocol
		params[prefix+"SSLCertificateId"] = Listeners[i].SSLCertificateId

	}

}

func makeParamsWithVersion(action string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	params["Version"] = loadbalancerversion
	return params
}
