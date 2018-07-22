package clouddataflow

import (
	"bytes"
	"encoding/json"
	googleauth "github.com/cloudlibz/gocloud/googleauth"
	"io/ioutil"
	"net/http"
)

//DescribeStream Describe Stream
func (clouddataflow *Clouddataflow) DescribeStream(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://dataflow.googleapis.com/v1b3/projects/" + options["Project"] + "/jobs/" + options["JobId"]

	client := googleauth.SignJWT()

	describestreamrequest, err := http.NewRequest("GET", url, nil)

	describestreamrequestparam := liststreamrequest.URL.Query()

	if options["View"] != "" {
		describestreamrequestparam.Add("view", options["View"])
	}

	if options["Location"] != "" {
		describestreamrequestparam.Add("location", options["Location"])
	}

	describestreamrequest.URL.RawQuery = describestreamrequestparam.Encode()
	describestreamrequest.Header.Set("Content-Type", "application/json")

	liststreamresp, err := client.Do(describestreamrequest)
	defer describestreamresp.Body.Close()

	body, err := ioutil.ReadAll(describestreamresp.Body)

	describestreamresponse := make(map[string]interface{})
	describestreamresponse["status"] = describestreamresp.StatusCode
	describestreamresponse["body"] = string(body)
	resp = describestreamresponse
	return resp, err
}

//ListStream ListStream
func (clouddataflow *Clouddataflow) ListStream(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://dataflow.googleapis.com/v1b3/projects/" + options["Project"] + "/jobs"

	client := googleauth.SignJWT()

	liststreamrequest, err := http.NewRequest("GET", url, nil)

	liststreamrequestparam := liststreamrequest.URL.Query()

	if options["PageSize"] != "" {
		liststreamrequestparam.Add("pageSize", options["pageSize"])
	}

	if options["PageToken"] != "" {
		liststreamrequestparam.Add("pageToken", options["PageToken"])
	}

	if options["Filter"] != "" {
		liststreamrequestparam.Add("filter", options["Filter"])
	}

	if options["View"] != "" {
		liststreamrequestparam.Add("view", options["View"])
	}

	if options["Location"] != "" {
		liststreamrequestparam.Add("location", options["Location"])
	}

	lliststreamrequest.URL.RawQuery = liststreamrequestparam.Encode()
	liststreamrequest.Header.Set("Content-Type", "application/json")

	liststreamresp, err := client.Do(liststreamrequest)
	defer liststreamresp.Body.Close()

	body, err := ioutil.ReadAll(liststreamresp.Body)

	liststreamresponse := make(map[string]interface{})
	liststreamresponse["status"] = liststreamresp.StatusCode
	liststreamresponse["body"] = string(body)
	resp = liststreamresponse
	return resp, err
}

//DeleteStream Delete Stream
func (clouddataflow *Clouddataflow) DeleteStream(request interface{}) (resp interface{}, err error) {

	return resp, err
}

//CreateStream Create Stream
func (clouddataflow *Clouddataflow) CreateStream(request interface{}) (resp interface{}, err error) {

	return resp, err
}
