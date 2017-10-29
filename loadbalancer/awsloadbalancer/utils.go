package awsloadbalancer

import (
	"github.com/cloudlibz/gocloud/auth"
	awsauth "github.com/cloudlibz/gocloud/awsauth"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

//timeNow represents time variable.
var timeNow = time.Now

//PrepareSignatureV2query for elasticloadbalancing.
func (awsloadbalancer *Awsloadbalancer) PrepareSignatureV2query(params map[string]string, response map[string]interface{}) error {

	ElasticloadbalancingEndpoint := "https://elasticloadbalancing.amazonaws.com"

	req, err := http.NewRequest("GET", ElasticloadbalancingEndpoint, nil)
	if err != nil {
		return err
	}

	// Add the params passed in to the query string.
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

	response["body"] = string(body)
	response["status"] = r.StatusCode
	return err
}

//prepareAvailabilityZones prepareAvailabilityZones attribute for CreateLoadBalancer.
func prepareAvailabilityZones(params map[string]string, AvailabilityZones []string) {

	for i := range AvailabilityZones {
		n := strconv.Itoa(i + 1)
		prefix := "AvailabilityZones.member." + n
		params[prefix] = AvailabilityZones[i]
	}

}

//prepareSubnets prepare prepareSubnets attribute for CreateLoadBalancer.
func prepareSubnets(params map[string]string, Subnets []string) {

	for i := range Subnets {
		n := strconv.Itoa(i + 1)
		prefix := "Subnets.member." + n
		params[prefix] = Subnets[i]
	}

}

//prepareSecurityGroups prepare prepareSecurityGroups attribute for CreateLoadBalancer.
func prepareSecurityGroups(params map[string]string, SecurityGroups []string) {

	for i := range SecurityGroups {
		n := strconv.Itoa(i + 1)
		prefix := "SecurityGroups.member." + n
		params[prefix] = SecurityGroups[i]
	}

}

//prepareListeners prepare prepareListeners attribute for CreateLoadBalancer.
func prepareListeners(params map[string]string, Listeners []Listener) {
	for i := range Listeners {
		n := strconv.Itoa(i + 1)
		prefix := "Listeners.member." + n
		params[prefix+".LoadBalancerPort"] = Listeners[i].LoadBalancerPort
		params[prefix+".InstancePort"] = Listeners[i].InstancePort
		params[prefix+".Protocol"] = Listeners[i].Protocol
		params[prefix+".InstanceProtocol"] = Listeners[i].InstanceProtocol
		params[prefix+".SSLCertificateId"] = Listeners[i].SSLCertificateId
	}
}

//prepareInstances prepare Instances attribute for Attachnodewithloadbalancer.
func prepareInstances(params map[string]string, Instances []string) {

	for i := range Instances {
		n := strconv.Itoa(i + 1)
		prefix := "Instances.member." + n
		params[prefix+".InstanceId"] = Instances[i]
	}
}

//makeParamsWithVersion add loadbalancerversion to queryString.
func makeParamsWithVersion(action string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	params["Version"] = loadbalancerversion
	return params
}
