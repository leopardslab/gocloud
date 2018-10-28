package clouddataflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/cloudlibz/gocloud/googleauth"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	RFC3339  = "2006-01-02T15:04:05Z07:00"
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

//DescribeStream Describe Stream
func (clouddataflow *Clouddataflow) UpdateStream(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var jobId string
	var option Createstream

	for key, value := range param {
		switch key {

		case "JobId":
			jobIdv, _ := value.(string)
			jobId = jobIdv

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
			option.Type = typev

		case "CurrentState":
			currentStatev, _ := value.(string)
			option.CurrentState = currentStatev

		case "CurrentStateTime":
			currentStateTimev, _ := value.(string)
			option.CurrentStateTime = currentStateTimev
			option.CurrentStateTime = time.Now().UTC().Format(time.RFC3339)

		case "RequestedState":
			requestedStatev, _ := value.(string)
			option.RequestedState = requestedStatev

		case "CreateTime":
			createTimev, _ := value.(string)
			option.CreateTime = createTimev
			option.CreateTime = time.Now().UTC().Format(time.RFC3339)

		case "ReplaceJobId":
			replaceJobIdv, _ := value.(string)
			option.ReplaceJobID = replaceJobIdv

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
					switch stageStatesparamkey {

					case "CurrentStateTime":
						currentStateTimev, _ := stageStatesparamvalue.(string)
						stageState.CurrentStateTime = currentStateTimev

					case "ExecutionStageName":
						executionStageNamev, _ := stageStatesparamvalue.(string)
						stageState.ExecutionStageName = executionStageNamev

					case "ExecutionStageState":
						executionStageStatev, _ := stageStatesparamvalue.(string)
						stageState.ExecutionStageState = executionStageStatev

					}

				}
				option.stageStates = append(option.stageStates, stageState)
			}

		case "Environment":
			environmentparam, _ := value.(map[string]interface{})

			for environmentparamkey, environmentparamvalue := range environmentparam {

				switch environmentparamkey {

				case "Version":
					versionparam, _ := environmentparamvalue.(map[string]interface{})

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

								switch supportparamkey {

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
		}
	}

	createstreamjsonmap := make(map[string]interface{})
	createstreamdictnoaryconvert(option, createstreamjsonmap)
	updatestreamjson, _ := json.Marshal(createstreamjsonmap)
	updatestreamjsonstring := string(updatestreamjson)
	var updatestreamjsonstringbyte = []byte(updatestreamjsonstring)

	url := "https://dataflow.googleapis.com/v1b3/projects/" + option.ProjectID + "/jobs/" + jobId
	client := googleauth.SignJWT()

	updatestreamrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(updatestreamjsonstringbyte))
	updatestreamrequestparam := updatestreamrequest.URL.Query()

	if option.Location != "" {
		updatestreamrequestparam.Add("location", option.Location)
	}

	updatestreamrequest.URL.RawQuery = updatestreamrequestparam.Encode()
	updatestreamrequest.Header.Set("Content-Type", "application/json")

	updatestreamresp, err := client.Do(updatestreamrequest)
	defer updatestreamresp.Body.Close()

	body, err := ioutil.ReadAll(updatestreamresp.Body)

	updatestreamresponse := make(map[string]interface{})
	updatestreamresponse["status"] = updatestreamresp.StatusCode
	updatestreamresponse["body"] = string(body)
	resp = updatestreamresponse
	return resp, err

}

//CreateStream Create Stream
func (clouddataflow *Clouddataflow) CreateStream(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var View string
	var option Createstream

	for key, value := range param {
		switch key {

		case "ProjectID":
			projectIdv, _ := value.(string)
			option.ProjectID = projectIdv

		case "View":
			viewv, _ := value.(string)
			View = viewv

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
			option.Type = typev

		case "CurrentState":
			currentStatev, _ := value.(string)
			option.CurrentState = currentStatev

		case "CurrentStateTime":
			currentStateTimev, _ := value.(string)
			option.CurrentStateTime = currentStateTimev
			option.CurrentStateTime = time.Now().UTC().Format(time.RFC3339)

		case "RequestedState":
			requestedStatev, _ := value.(string)
			option.RequestedState = requestedStatev

		case "CreateTime":
			createTimev, _ := value.(string)
			option.CreateTime = createTimev
			option.CreateTime = time.Now().UTC().Format(time.RFC3339)

		case "ReplaceJobId":
			replaceJobIdv, _ := value.(string)
			option.ReplaceJobID = replaceJobIdv

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
					switch stageStatesparamkey {

					case "CurrentStateTime":
						currentStateTimev, _ := stageStatesparamvalue.(string)
						stageState.CurrentStateTime = currentStateTimev

					case "ExecutionStageName":
						executionStageNamev, _ := stageStatesparamvalue.(string)
						stageState.ExecutionStageName = executionStageNamev

					case "ExecutionStageState":
						executionStageStatev, _ := stageStatesparamvalue.(string)
						stageState.ExecutionStageState = executionStageStatev

					}

				}
				option.stageStates = append(option.stageStates, stageState)
			}

		case "Environment":
			environmentparam, _ := value.(map[string]interface{})

			for environmentparamkey, environmentparamvalue := range environmentparam {

				switch environmentparamkey {

				case "Version":

					versionparam, _ := environmentparamvalue.(map[string]interface{})

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

								switch supportparamkey {

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
		}
	}

	createstreamjsonmap := make(map[string]interface{})
	createstreamdictnoaryconvert(option, createstreamjsonmap)
	createstreamjson, _ := json.Marshal(createstreamjsonmap)
	createstreamjsonstring := string(createstreamjson)
	var createstreamjsonstringbyte = []byte(createstreamjsonstring)

	url := "https://dataflow.googleapis.com/v1b3/projects/" + option.ProjectID + "/jobs"
	client := googleauth.SignJWT()

	createstreamrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(createstreamjsonstringbyte))
	createstreamrequestparam := createstreamrequest.URL.Query()

	if View != "" {
		createstreamrequestparam.Add("view", View)
	}

	if option.Location != "" {
		createstreamrequestparam.Add("location", option.Location)
	}

	if option.ReplaceJobID != "" {
		createstreamrequestparam.Add("replaceJobId", option.ReplaceJobID)
	}

	createstreamrequest.URL.RawQuery = createstreamrequestparam.Encode()
	createstreamrequest.Header.Set("Content-Type", "application/json")

	createstreamresp, err := client.Do(createstreamrequest)
	defer createstreamresp.Body.Close()

	body, err := ioutil.ReadAll(createstreamresp.Body)

	createstreamresponse := make(map[string]interface{})
	createstreamresponse["status"] = createstreamresp.StatusCode
	createstreamresponse["body"] = string(body)
	resp = createstreamresponse
	return resp, err
}
