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

//CreateStream Create Stream
func (clouddataflow *Clouddataflow) CreateStream(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var Project string
	var option Createstream

	for key, value := range param {
		switch key {

		case "ProjectID":
			projectIdv, _ := value.(string)
			option.ProjectID = projectIdv

		case "ClientRequestID":
			clientRequestIDv, _ := value.(string)
			option.ClientRequestID = clientRequestIDv

		case "ID":
			idv, _ := value.(string)
			option.ID = idv

		case "Name":
			nameV, _ := value.(string)
			option.Name = nameV

		case "Type":
			typev, _ := value.(string)
			option.Type = typeV

		case "CurrentState":
			currentStatev, _ := value.(string)
			option.CreateTime = currentStatev

		case "CurrentStateTime":
			currentStateTimev, _ := value.(string)
			option.CurrentStateTime = currentStateTimev

		case "RequestedState":
			requestedStatev, _ := value.(string)
			option.RequestedState = requestedStatev

		case "CreateTime":
			createTimev, _ := value.(string)
			option.CreateTime = createTimev
			option.CreateTime = createTimev

		case "ReplaceJobId":
			replaceJobIdv, _ := value.(string)
			option.ReplaceJobId = replaceJobIdv

		case "ReplacedByJobID":
			replacedByJobIDv, _ := value.(string)
			option.ReplacedByJobID = replacedByJobIDv

		case "Location":
			locationv, _ := value.(string)
			option.Location = locationv

		case "TempFiles":
			tempFilesv, _ := value.([]string)
			option.TempFiles = tempFilesv

		case "StageStates":

			stageStatesparam, _ := value.([]map[string]interface{})

			for i := 0; i < len(stageStatesparam); i++ {
				var stageState StageStates
				for stageStatesparamkey, stageStatesparamvalue := range stageStatesparam[i] {
					switch key {

					case "CurrentStateTime":
						currentStateTimev, _ := stageStatesparamvalue.(string)
						stageState.CurrentStateTime = currentStateTimev

					case "ExecutionStageName":
						executionStageNamev, _ := stageStatesparamvalue.(string)
						stageState.CurrentStateTime = executionStageNamev

					case "ExecutionStageState":
						executionStageStatev, _ := stageStatesparamvalue.(string)
						stageState.ExecutionStageState = executionStageStatev

					}

				}
				option.stageStates = append(option.stageStates, stageState)
			}

		case "Environment":

			environmentparam, _ := value.([]map[string]interface{})

			for environmentparamkey, environmentparamvalue := range environmentparam {

				switch environmentparamkey {

				case "Version":

					versionparam, _ := environmentparamvalue.([]map[string]interface{})

					for versionparamkey, versionparamvalue := range versionparam {
						switch versionparamkey {

						case "Major":
							majorv, _ := versionparamvalue.(string)
							option.environment.version.Major = majorv

						case "JobType":
							jobTypev, _ := versionparamvalue.(string)
							option.environment.version.JobType = jobTypev
						}
					}

				case "UserAgent":

					useragentparam, _ := environmentparamvalue.(map[string]interface{})

					for useragentparamkey, useragentparamvalue := range useragentparam {

						switch useragentparamkey {

						case "Name":
							namev, _ := useragentparamvalue.(string)
							option.environment.userAgent.Name = namev

						case "BuildDate":
							buildDatev, _ := useragentparamvalue.(string)
							option.environment.userAgent.BuildDate = buildDatev

						case "Version":
							versionv, _ := useragentparamvalue.(string)
							option.environment.userAgent.Version = versionv

						case "Support":
							supportparam, _ := useragentparamvalue.(map[string]interface{})

							for supportparamkey, supportparamvalue := range supportparam {

								switch useragentparamkey {

								case "Status":
									statusv, _ := supportparamvalue.(string)
									option.environment.userAgent.support.Status = statusv

								case "URL":
									urlv, _ := supportparamvalue.(string)
									option.environment.userAgent.support.URL = urlv

								}
							}

						}
					}

				}
			}

			//end of switch key
		}
	} //end of parse

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

func createstreamdictnoaryconvert(option Createstream, createstreamjsonmap map[string]interface{}) {

	if option.Id != "" {
		createstreamjsonmap["id"] = option.Id
	}

	if option.ProjectId != "" {
		createstreamjsonmap["projectId"] = option.ProjectId
	}

	if option.Name != "" {
		createstreamjsonmap["name"] = option.Name
	}

	if option.Type != "" {
		createstreamjsonmap["type"] = option.Type
	}

	if option.CurrentState != "" {
		createstreamjsonmap["currentState"] = option.CurrentState
	}

	if option.CurrentStateTime != "" {
		createstreamjsonmap["currentStateTime"] = option.CurrentStateTime
	}

	if option.RequestedState != "" {
		createstreamjsonmap["requestedState"] = option.RequestedState
	}

	if option.CreateTime != "" {
		createstreamjsonmap["createTime"] = option.CreateTime
	}

	if option.ReplaceJobId != "" {
		createstreamjsonmap["replaceJobId"] = option.ReplaceJobId
	}

	if option.Location != "" {
		createstreamjsonmap["location"] = option.Location
	}

}
