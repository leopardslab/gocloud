package googlecloudfunctions


type Googlecloudfunctions struct{

}

type CreateGooglecloudfunction struct {
	Name         string `json:"name"`
	Status              string    `json:"status"`
	EntryPoint          string    `json:"entryPoint"`
	Timeout             string    `json:"timeout"`
	AvailableMemoryMb   int       `json:"availableMemoryMb"`
	ServiceAccountEmail string    `json:"serviceAccountEmail"`
	UpdateTime          time.Time `json:"updateTime"`
	VersionID           string    `json:"versionId"`
	SourceUploadURL string `json:"sourceUploadUrl"`
  Labels   labels `json:"labels"`
  HTTPSTrigger httpstrigger `json:"httpsTrigger"`
}

type labels  struct {
  DeploymentTool string `json:"deployment-tool"`
}

type httpstrigger struct {
  URL string `json:"url"`
}


func (googlecloudfunctions *Googlecloudfunctions) Create(request interface{}) (resp interface{}, err error) {

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
			LabelsV, _ := value.(map[sting]string)
			option.Labels.DeploymentTool = LabelsV["DeploymentTool"]

    case "HTTPSTrigger":
			HTTPSTriggerV, _ := value.(map[sting]string)
			option.HTTPSTrigger.URL = HTTPSTriggerV["URL"]
		}
	}

	CreateGooglecloudfunctionjsonmap := make(map[string]interface{})

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

func (googlecloudfunctions *Googlecloudfunctions) delete(request interface{}) (resp interface{}, err error) {

  options := request.(map[string]string)

	url := "https://cloudfunctions.googleapis.com/v1/" + options["name"]

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

func (googlecloudfunctions *Googlecloudfunctions) get(request interface{}) (resp interface{}, err error) {

  options := request.(map[string]string)

  url := "https://cloudfunctions.googleapis.com/v1/" + options["name"]

  client := googleauth.SignJWT()

  GetGooglecloudfunctionrequest, err := http.NewRequest("GET", url, nil)

  GetGooglecloudfunctionrequest.Header.Set("Content-Type", "application/json")

  GetGooglecloudfunctionsresp, err := client.Do(GetGooglecloudfunctionrequest)

  defer GetGooglecloudfunctionresp.Body.Close()

  body, err := ioutil.ReadAll(GetGooglecloudfunctionresp.Body)

  GetGooglecloudfunctionresponse := make(map[string]interface{})
  GetGooglecloudfunctionresponse["status"] = GetGooglecloudfunctionresp.StatusCode
  GetGooglecloudfunctionresponse["body"] = string(body)
  resp = GetGooglecloudfunctionresponse
  return resp, err

}

func (googlecloudfunctions *Googlecloudfunctions) list(request interface{}) (resp interface{}, err error) {

  options := request.(map[string]string)

  url := "https://cloudfunctions.googleapis.com/v1/" + options["name"] + ":call"

  client := googleauth.SignJWT()

  listgooglecloudfunctionrequest, err := http.NewRequest("GET", url, nil)

  listgooglecloudfunctionrequestparam := Listdnsrequest.URL.Query()

  if options["pageToken"] != "" {
    listgooglecloudfunctionrequestparam.Add("pageToken", options["pageToken"])
  }

  if options["pageSize"] != "0" {
    listgooglecloudfunctionrequestparam.Add("pageSize", options["pageSize"])
  }

  listgooglecloudfunctionrequest.URL.RawQuery = Listdnsrequestparam.Encode()

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


func (googlecloudfunctions *Googlecloudfunctions) call(request interface{}) (resp interface{}, err error) {

  options := request.(map[string]string)

	url := "https://cloudfunctions.googleapis.com/v1/ + options["name] + ":call"

  callGooglecloudfunctionjsonmap := make(map[string]interface{})

  callGooglecloudfunctionjsonmap = options["data"]

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
