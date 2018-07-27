package clouddataflow

import (
	//"bytes"
	//"encoding/json"
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

	describestreamrequestparam := describestreamrequest.URL.Query()

	if options["View"] != "" {
		describestreamrequestparam.Add("view", options["View"])
	}

	if options["Location"] != "" {
		describestreamrequestparam.Add("location", options["Location"])
	}

	describestreamrequest.URL.RawQuery = describestreamrequestparam.Encode()
	describestreamrequest.Header.Set("Content-Type", "application/json")

	describestreamresp, err := client.Do(describestreamrequest)
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

	liststreamrequest.URL.RawQuery = liststreamrequestparam.Encode()
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

/*

//CreateStream Create Stream
func (clouddataflow *Clouddataflow) CreateStream(request interface{}) (resp interface{}, err error) {

 param := request.(map[string]interface{})
 var Project string
 var option Createstream

 for key, value := range param {
	 switch key {

	 case "ProjectId":
		 projectIdv, _ := value.(string)
		 option.projectId = projectIdv

	 case "Id":
		 idv, _ := value.(string)
		 option.id = idv

	 case "Name":
		 nameV, _ := value.(string)
		 option.name = nameV

	 case "Type":
		 typev, _ := value.(string)
		 option.type = typeV

	 case "CurrentState":
		 currentStatev, _ := value.([]string)
		 option.currentState = currentStatev

	 case "CurrentStateTime":
		 currentStateTimev, _ := value.(string)
		 option.currentStateTime = currentStateTimev

	 case "RequestedState":
		 requestedStatev, _ := value.(string)
		 option.requestedState = requestedStatev

	 case "CreateTime":
		 createTimev, _ := value.(string)
		 option.createTimev = createTimev
		 option.createTimev = createTimev

	 case "ReplaceJobId":
		 replaceJobIdv, _ := value.(string)
		 option.replaceJobId = replaceJobIdv

	 case "Location":
		 locationv, _ := value.(string)
		 option.location = locationv

	 }
 }

 createstreamjsonmap := make(map[string]interface{})

 createstreamdictnoaryconvert(option, createstreamjsonmap)

 createstreamjson, _ := json.Marshal(createstreamjsonmap)

 createstreamjsonstring := string(createstreamjson)

 var createstreamjsonstringbyte = []byte(createstreamjsonstring)

 url := "https://www.googleapis.com/dns/v1/projects/" + Project + "/managedZones"

 client := googleauth.SignJWT()

 createstreamrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(createstreamjsonstringbyte))

 createstreamrequest.Header.Set("Content-Type", "application/json")

 createstreamresp, err := client.Do(createstreamrequest)

 defer createstreamresp.Body.Close()

 body, err := ioutil.ReadAll(createstreamresp.Body)

 createstreamresponse := make(map[string]interface{})
 createstreamresponse["status"] = createstreamrresp.StatusCode
 createstreamresponse["body"] = string(body)
 resp = createstreamresponse
 return resp, err
}

func createstreamdictnoaryconvert(option Createstream ,createstreamjsonmap map[string]interface{} ){

	if option.id != "" {
		createstreamjsonmap["id"] = option.id
	}

	if option.projectId != "" {
		createstreamjsonmap["projectId"] = option.projectId
	}

	if option.name != "" {
		createstreamjsonmap["name"] = option.name
	}

	if option.Type != "" {
		createstreamjsonmap["type"] = option.Type
	}

	if option.currentState != "" {
		createstreamjsonmap["currentState"] = option.currentState
	}

	if option.currentStateTime != "" {
		createstreamjsonmap["currentStateTime"] = option.currentStateTime
	}

	if option.currentStateTime != "" {
		createstreamjsonmap["currentStateTime"] = option.currentStateTime
	}

	if option.requestedState != "" {
		createstreamjsonmap["requestedState"] = option.requestedState
	}

	if option.requestedState != "" {
		createstreamjsonmap["requestedState"] = option.requestedState
	}

	if option.createTime != "" {
		createstreamjsonmap["createTime"] = option.createTime
	}

	if option.replaceJobId != "" {
		createstreamjsonmap["replaceJobId"] = option.replaceJobId
	}

	if option.location != "" {
		createstreamjsonmap["location"] = option.location
	}

}

*/
