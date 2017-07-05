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

type Awsloadbalancer struct {
}

type CreateLoadBalancer struct {
	Name           string
	IpAddressType  string
	Scheme         string
	Tags           []Tag
	SecurityGroups []string
	Subnets        []string
}

type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

func (awsloadbalancer *Awsloadbalancer) Creatloadbalancer(request interface{}) (resp interface{}, err error) {

	var options CreateLoadBalancer

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {

		case "Name":
			NameV, _ := value.(string)
			options.Name = NameV

		case "Subnets":
			SubnetsV, _ := value.([]string)
			options.Subnets = SubnetsV
		}
	}

	fmt.Println("options :", options)

	params := makeParamsWithVersion("CreateLoadBalancer")

	prepareSubnets(params, options.Subnets)

	if options.Name != "" {
		params["Name"] = options.Name
	}

	if options.IpAddressType != "" {
		params["IpAddressType"] = options.IpAddressType
	}

	if options.Scheme != "" {
		params["Scheme"] = options.Scheme
	}

	fmt.Println(params)

	awsloadbalancer.PrepareSignatureV2query(params)

	return
}

func prepareSubnets(params map[string]string, Subnets []string) {

	for i := range Subnets {
		n := strconv.Itoa(i + 1)
		prefix := "Subnets.member." + n
		params[prefix] = Subnets[i]
	}

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
