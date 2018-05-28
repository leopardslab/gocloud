package googlecloudfunctions

import (
	"bytes"
	"encoding/json"
	googleauth "github.com/cloudlibz/gocloud/googleauth"
	"io/ioutil"
	"net/http"
	"time"
	"fmt"
)

const (
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	RFC3339  = "2006-01-02T15:04:05Z07:00"
)

type Googlecloudfunctions struct {
}

type CreateGooglecloudfunction struct {
	Name                string       `json:"name"`
	Status              string       `json:"status"`
	EntryPoint          string       `json:"entryPoint"`
	Timeout             string       `json:"timeout"`
	AvailableMemoryMb   int          `json:"availableMemoryMb"`
	ServiceAccountEmail string       `json:"serviceAccountEmail"`
	UpdateTime          string       `json:"updateTime"`
	VersionID           string       `json:"versionId"`
	SourceUploadURL     string       `json:"sourceUploadUrl"`
	Labels              labels       `json:"labels"`
	HTTPSTrigger        httpstrigger `json:"httpsTrigger"`
}

type labels struct {
	DeploymentTool string `json:"deployment-tool"`
}

type httpstrigger struct {
	URL string `json:"url"`
}

func (googlecloudfunctions *Googlecloudfunctions) Createfunction(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var Location string
	var option CreateGooglecloudfunction

	for key, value := range param {
		switch key {

		case "Location":
			LocationV, _ := value.(string)
			Location = LocationV

		case "Timeout":
			TimeoutV, _ := value.(string)
			option.Timeout = TimeoutV

		case "UpdateTime":
			UpdateTimeV, _ := value.(string)
			option.UpdateTime = UpdateTimeV
			option.UpdateTime = time.Now().UTC().Format(time.RFC3339)

		case "Name":
			NameV, _ := value.(string)
			option.Name = NameV

		case "Status":
			StatusV, _ := value.(string)
			option.Status = StatusV

		case "EntryPoint":
			EntryPointV, _ := value.(string)
			option.EntryPoint = EntryPointV

		case "ServiceAccountEmail":
			ServiceAccountEmailV, _ := value.(string)
			option.ServiceAccountEmail = ServiceAccountEmailV

		case "VersionID":
			VersionIDV, _ := value.(string)
			option.VersionID = VersionIDV

		case "SourceUploadURL":
			SourceUploadURLV, _ := value.(string)
			option.SourceUploadURL = SourceUploadURLV

		case "AvailableMemoryMb":
			AvailableMemoryMbV, _ := value.(int)
			option.AvailableMemoryMb = AvailableMemoryMbV

		case "Labels":
			LabelsV, _ := value.(map[string]string)
			option.Labels.DeploymentTool = LabelsV["DeploymentTool"]

		case "HTTPSTrigger":
			HTTPSTriggerV, _ := value.(map[string]string)
			option.HTTPSTrigger.URL = HTTPSTriggerV["URL"]
		}
	}

	CreateGooglecloudfunctionjsonmap := make(map[string]interface{})

	fmt.Println(option.SourceUploadURL)

	CreateGooglecloudfunctionedictnoaryconvert(option, CreateGooglecloudfunctionjsonmap)

	CreateGooglecloudfunctionjson, _ := json.Marshal(CreateGooglecloudfunctionjsonmap)

	CreateGooglecloudfunctionjsonstring := string(CreateGooglecloudfunctionjson)

	var CreateGooglecloudfunctionjsonstringbyte = []byte(CreateGooglecloudfunctionjsonstring)

	url := "https://cloudfunctions.googleapis.com/v1/" + Location + "/functions"

	client := googleauth.SignJWT()

	CreateGooglecloudfunctionrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(CreateGooglecloudfunctionjsonstringbyte))

	CreateGooglecloudfunctionrequest.Header.Set("Content-Type", "application/json")

	CreateGooglecloudfunctionresp, err := client.Do(CreateGooglecloudfunctionrequest)

	defer CreateGooglecloudfunctionresp.Body.Close()

	body, err := ioutil.ReadAll(CreateGooglecloudfunctionresp.Body)

	CreateGooglecloudfunctionresponse := make(map[string]interface{})
	CreateGooglecloudfunctionresponse["status"] = CreateGooglecloudfunctionresp.StatusCode
	CreateGooglecloudfunctionresponse["body"] = string(body)
	resp = CreateGooglecloudfunctionresponse
	return resp, err
}

/*
{
  "name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
  "httpsTrigger": {
    "url": "https://us-central1-adept-comfort-202709.cloudfunctions.net/function-1"
  },
  "status": "ACTIVE",
  "entryPoint": "helloWorld",
  "timeout": "60s",
  "availableMemoryMb": 256,
  "serviceAccountEmail": "adept-comfort-202709@appspot.gserviceaccount.com",
  "updateTime": "2018-05-11T18:20:33Z",
  "versionId": "1",
  "labels": {
    "deployment-tool": "console-cloud"
  },
  "sourceUploadUrl": "https://storage.googleapis.com/gcf-upload-us-central1-f24bda97-6cd1-46cc-b37d-1f60eac4210a/8548b011-9626-42c1-86ed-6190892b328e.zip?GoogleAccessId=126778294088@cloudservices.gserviceaccount.com&Expires=1526064618&Signature=nB%2FI6cwIap0DF5T0Uo9eYCnlmi3HLqvoRW4MfodzVI%2FXuC7HU%2BE9SwduVQKYeTRddo5iFNdm4VDmBu4A4fGQvZ5PaCuoKG4i7jZXRJgq1B4NIpocaFnHmY6ZWaCS0Av%2Bus29FHs2nTYIqp9zHWHHORSQC%2BPF8GP2mRToDOShpodkQFkxP6wsXUnkk8tDUf5mvTRkeqtgf0rX0huidbEVl7ZtGkcQiusDcS9Nhe3dwqOdsJ7xs2khl2D%2FOmch6jgrZ11MtXum3G5XnFLqMYupS0pvB%2BQiy7g7eLfIw%2BdvtRTEuEFyxWP49lCHUG8wWad0hNEVf29oAHS2x%2B4Q%2FaIGbA%3D%3D",
  "runtime": "nodejs6"
}

*/

func CreateGooglecloudfunctionedictnoaryconvert(option CreateGooglecloudfunction, CreateGooglecloudfunctionjsonmap map[string]interface{}) {

	if option.AvailableMemoryMb != 0 {
		CreateGooglecloudfunctionjsonmap["availableMemoryMb"] = option.AvailableMemoryMb
	}

	if option.Name != "" {
		CreateGooglecloudfunctionjsonmap["name"] = option.Name
	}

	if option.Status != "" {
		CreateGooglecloudfunctionjsonmap["status"] = option.Status
	}

	if option.EntryPoint != "" {
		CreateGooglecloudfunctionjsonmap["entryPoint"] = option.EntryPoint
	}

	if option.Timeout != "" {
		CreateGooglecloudfunctionjsonmap["timeout"] = option.Timeout
	}

	if option.ServiceAccountEmail != "" {
		CreateGooglecloudfunctionjsonmap["serviceAccountEmail"] = option.ServiceAccountEmail
	}

	if option.UpdateTime != "" {
		CreateGooglecloudfunctionjsonmap["updateTime"] = option.UpdateTime
	}

	if option.VersionID != "" {
		CreateGooglecloudfunctionjsonmap["versionId"] = option.VersionID
	}

	if option.ServiceAccountEmail != "" {
		CreateGooglecloudfunctionjsonmap["sourceUploadUrl"] = option.SourceUploadURL
	}

	if option.Labels.DeploymentTool != "" {
		labels := make(map[string]string)
		labels["deployment-tool"] = option.Labels.DeploymentTool
		CreateGooglecloudfunctionjsonmap["labels"] = labels
	}

	if option.HTTPSTrigger.URL != "" {
		labels := make(map[string]string)
		labels["url"] = option.HTTPSTrigger.URL
		CreateGooglecloudfunctionjsonmap["httpsTrigger"] = labels
	}

}

func (googlecloudfunctions *Googlecloudfunctions) Deletefunction(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://cloudfunctions.googleapis.com/v1/" + options["name"]

	client := googleauth.SignJWT()

	DeleteGooglecloudfunctionrequest, err := http.NewRequest("DELETE", url, nil)

	DeleteGooglecloudfunctionrequest.Header.Set("Content-Type", "application/json")

	DeleteGooglecloudfunctionresp, err := client.Do(DeleteGooglecloudfunctionrequest)

	defer DeleteGooglecloudfunctionresp.Body.Close()

	body, err := ioutil.ReadAll(DeleteGooglecloudfunctionresp.Body)

	DeleteGooglecloudfunctionresponse := make(map[string]interface{})
	DeleteGooglecloudfunctionresponse["status"] = DeleteGooglecloudfunctionresp.StatusCode
	DeleteGooglecloudfunctionresponse["body"] = string(body)
	resp = DeleteGooglecloudfunctionresponse
	return resp, err
}

func (googlecloudfunctions *Googlecloudfunctions) Getfunction(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://cloudfunctions.googleapis.com/v1/" + options["name"]

	client := googleauth.SignJWT()

	GetGooglecloudfunctionrequest, err := http.NewRequest("GET", url, nil)

	GetGooglecloudfunctionrequest.Header.Set("Content-Type", "application/json")

	GetGooglecloudfunctionsresp, err := client.Do(GetGooglecloudfunctionrequest)

	defer GetGooglecloudfunctionsresp.Body.Close()

	body, err := ioutil.ReadAll(GetGooglecloudfunctionsresp.Body)

	GetGooglecloudfunctionresponse := make(map[string]interface{})
	GetGooglecloudfunctionresponse["status"] = GetGooglecloudfunctionsresp.StatusCode
	GetGooglecloudfunctionresponse["body"] = string(body)
	resp = GetGooglecloudfunctionresponse
	return resp, err

}

func (googlecloudfunctions *Googlecloudfunctions) Listfunction(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://cloudfunctions.googleapis.com/v1/" + options["name"] + "/functions"

	client := googleauth.SignJWT()

	listgooglecloudfunctionrequest, err := http.NewRequest("GET", url, nil)

	listgooglecloudfunctionrequestparam := listgooglecloudfunctionrequest.URL.Query()

	if options["pageToken"] != "" {
		listgooglecloudfunctionrequestparam.Add("pageToken", options["pageToken"])
	}

	if options["pageSize"] != "0" {
		listgooglecloudfunctionrequestparam.Add("pageSize", options["pageSize"])
	}

	listgooglecloudfunctionrequest.URL.RawQuery = listgooglecloudfunctionrequestparam.Encode()

	listgooglecloudfunctionrequest.Header.Set("Content-Type", "application/json")

	listgooglecloudfunctionresp, err := client.Do(listgooglecloudfunctionrequest)

	defer listgooglecloudfunctionresp.Body.Close()

	body, err := ioutil.ReadAll(listgooglecloudfunctionresp.Body)

	listgooglecloudfunctionresponse := make(map[string]interface{})
	listgooglecloudfunctionresponse["status"] = listgooglecloudfunctionresp.StatusCode
	listgooglecloudfunctionresponse["body"] = string(body)
	resp = listgooglecloudfunctionresponse
	return resp, err
}

func (googlecloudfunctions *Googlecloudfunctions) Callfunction(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://cloudfunctions.googleapis.com/v1/" + options["name"] + ":call"

	callGooglecloudfunctionjsonmap := make(map[string]interface{})

	callGooglecloudfunctionjsonmap["data"] = options["data"]

	callGooglecloudfunctionjson, _ := json.Marshal(callGooglecloudfunctionjsonmap)

	callGooglecloudfunctionjsonstring := string(callGooglecloudfunctionjson)

	var callGooglecloudfunctionjsonstringbyte = []byte(callGooglecloudfunctionjsonstring)

	client := googleauth.SignJWT()

	callGooglecloudfunctionrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(callGooglecloudfunctionjsonstringbyte))

	callGooglecloudfunctionrequest.Header.Set("Content-Type", "application/json")

	callGooglecloudfunctionresp, err := client.Do(callGooglecloudfunctionrequest)

	defer callGooglecloudfunctionresp.Body.Close()

	body, err := ioutil.ReadAll(callGooglecloudfunctionresp.Body)

	callGooglecloudfunctionresponse := make(map[string]interface{})
	callGooglecloudfunctionresponse["status"] = callGooglecloudfunctionresp.StatusCode
	callGooglecloudfunctionresponse["body"] = string(body)
	resp = callGooglecloudfunctionresponse
	return resp, err
}

//projects/adept-comfort-202709/locations/us-central1/functions/function-1

//projects/adept-comfort-202709/locations/us-central1/functions/function-1
