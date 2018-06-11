package bigquery

//CreateDatasets Create Datasets.
func (bigquery *Bigquery) CreateDatasets(request interface{}) (resp interface{}, err error) {

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
