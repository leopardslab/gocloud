package bigquery

//CreateDatasets Create Datasets.
func (bigquery *Bigquery) CreateDatasets(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var projectId string
	var option CreateDatasets

	for key, value := range param {
		switch key {

		case "ProjectId":
			ProjectIdV, _ := value.(string)
			projectId = ProjectIdV
		}
	}

	createdatasetsstruct(option, param)

	createdatasetsjsonmap := make(map[string]interface{})

	createdatasetsdictnoaryconvert(option, createdatasetsjsonmap)

	createdatasetsjson, _ := json.Marshal(createdatasetsjsonmap)

	createdatasetsjsonstring := string(createdatasetsjson)

	var createdatasetsjsonstringbyte = []byte(createdatasetsjsonstring)

	url := "https://www.googleapis.com/bigquery/v2/projects/" + projectId + "/datasets"

	client := googleauth.SignJWT()

	createdatasetsrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(createdatasetsjsonstringbyte))

	createdatasetsrequest.Header.Set("Content-Type", "application/json")

	createdatasetsresp, err := client.Do(createdatasetsrequest)

	defer createdatasetsresp.Body.Close()

	body, err := ioutil.ReadAll(createdatasetsresp.Body)

	createdatasetsresponse := make(map[string]interface{})
	createdatasetsresponse["status"] = createdatasetsresp.StatusCode
	createdatasetsresponse["body"] = string(body)
	resp = createdatasetsresponse
	return resp, err
}

//DeleteDatasets delete Datasets.
func (bigquery *Bigquery) DeleteDatasets(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/bigquery/v2/projects/" + options["projectId"] + "/datasets" + options["datasetId"]

	client := googleauth.SignJWT()

	deletedatasetsrequest, err := http.NewRequest("DELETE", url, nil)

	deletedatasetsrequestparam := deletedatasetsrequest.URL.Query()

	if options["deleteContents"] != "" {
		deletedatasetsrequestparam.Add("deleteContents", options["deleteContents"])
	}

	deletedatasetsrequest.URL.RawQuery = deletedatasetsrequestparam.Encode()

	deletedatasetsrequest.Header.Set("Content-Type", "application/json")

	deletedatasetsrequestresp, err := client.Do(deletedatasetsrequest)

	defer deletedatasetsrequestresp.Body.Close()

	body, err := ioutil.ReadAll(deletedatasetsrequestresp.Body)

	deletedatasetsrequestrespresponse := make(map[string]interface{})
	deletedatasetsrequestrespresponse["status"] = deletedatasetsrequestresp.StatusCode
	deletedatasetsrequestrespresponse["body"] = string(body)
	resp = deletedatasetsrequestrespresponse
	return resp, err
}

//GetDatasets get Datasets.
func (bigquery *Bigquery) GetDatasets(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/bigquery/v2/projects/" + options["projectId"] + "/datasets" + options["datasetId"]

	client := googleauth.SignJWT()

	getdatasetsrequest, err := http.NewRequest("GET", url, nil)

	getdatasetsrequest.Header.Set("Content-Type", "application/json")

	getdatasetsrequestresp, err := client.Do(getdatasetsrequest)

	defer getdatasetsrequestresp.Body.Close()

	body, err := ioutil.ReadAll(getdatasetsrequestresp.Body)

	getdatasetsrequestrespresponse := make(map[string]interface{})
	getdatasetsrequestrespresponse["status"] = getdatasetsrequestresp.StatusCode
	getdatasetsrequestrespresponse["body"] = string(body)
	resp = getdatasetsrequestrespresponse
	return resp, err
}

//UpdateDatasets  Update Datasets.
func (bigquery *Bigquery) UpdateDatasets(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var projectId, datasetId string
	var option CreateDatasets

	for key, value := range param {
		switch key {
		case "ProjectId":
			projectIdV, _ := value.(string)
			projectId = projectIdV

		case "datasetId":
			datasetIdV, _ := value.(string)
			datasetId = datasetIdV
		}
	}

	createdatasetsstruct(option, param)

	createdatasetsjsonmap := make(map[string]interface{})

	createdatasetsdictnoaryconvert(option, createdatasetsjsonmap)

	updatedatasetsjson, _ := json.Marshal(createdatasetsjsonmap)

	updatedatasetsjsonstring := string(updatedatasetsjson)

	var updatedatasetsjsonstringbyte = []byte(updatedatasetsjsonstring)

	url := "https://www.googleapis.com/bigquery/v2/projects/" + projectId + "/datasets/" + datasetId

	client := googleauth.SignJWT()

	updatedatasetsrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(updatedatasetsjsonstringbyte))

	updatedatasetsrequest.Header.Set("Content-Type", "application/json")

	updatedatasetsresp, err := client.Do(updatedatasetsrequest)

	defer updatedatasetsresp.Body.Close()

	body, err := ioutil.ReadAll(updatedatasetsresp.Body)

	updatedatasetsresponse := make(map[string]interface{})
	updatedatasetsresponse["status"] = updatedatasetsresp.StatusCode
	updatedatasetsresponse["body"] = string(body)
	resp = updatedatasetsresponse
	return resp, err
}

//ListDatasets  list Datasets.
func (bigquery *Bigquery) ListDatasets(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/bigquery/v2/projects/" + options["projectId"] + "/datasets"

	client := googleauth.SignJWT()

	listdatasetsrequest, err := http.NewRequest("GET", url, nil)

	listdatasetsrequestparam := listdatasetsrequest.URL.Query()

	if options["filter"] != "" {
		listdatasetsrequestparam.Add("filter", options["filter"])
	}

	if options["maxResults"] != "" {
		listdatasetsrequestparam.Add("maxResults", options["maxResults"])
	}

	if options["pageToken"] != "" {
		listdatasetsrequestparam.Add("pageToken", options["pageToken"])
	}

	listdatasetsrequest.URL.RawQuery = listdatasetsrequestparam.Encode()

	listdatasetsrequest.Header.Set("Content-Type", "application/json")

	listdatasetsrequestresp, err := client.Do(listdatasetsrequest)

	defer listdatasetsrequestresp.Body.Close()

	body, err := ioutil.ReadAll(listdatasetsrequestresp.Body)

	listdatasetsrequestrespresponse := make(map[string]interface{})
	listdatasetsrequestrespresponse["status"] = listdatasetsrequestresp.StatusCode
	listdatasetsrequestrespresponse["body"] = string(body)
	resp = listdatasetsrequestrespresponse
	return resp, err
}
