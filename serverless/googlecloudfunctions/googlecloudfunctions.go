package googlecloudfunctions

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/cloudlibz/gocloud/googleauth"
	"io/ioutil"
	"net/http"
	"time"
)

//Create  Create Google cloud function.
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

//Delete delete function.
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

//Getfunction get function.
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

//Listfunction list function.
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

//Callfunction call function.
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
