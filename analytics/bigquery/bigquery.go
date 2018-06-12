package bigquery

//CreateDatasets Create Datasets.
func (bigquery *Bigquery) CreateDatasets(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var projectId string
	var option CreateDatasets

	for key, value := range param {
		switch key {

		case "projectId":
			ProjectIdV, _ := value.(string)
			ProjectId = ProjectIdV

		case "creationTime":
			CreationTimeV, _ := value.(string)
			option.CreationTime = CreationTimeV
			option.CreationTime = time.Now().UTC().Format(time.RFC3339)

		case "defaultTableExpirationMs":
			defaultTableExpirationMsV, _ := value.(string)
			option.defaultTableExpirationMs = defaultTableExpirationMsV

		case "description":
			descriptionV, _ := value.(string)
			option.description = descriptionV

		case "etag":
			etagV, _ := value.(string)
			option.etag = etagV

		case "friendlyName":
			friendlyNameV, _ := value.(string)
			option.friendlyName = friendlyNameV

		case "id":
			idV, _ := value.(string)
			option.id = idV

		case "kind":
			kindV, _ := value.(string)
			option.kind = kindV

		case "lastModifiedTime":
			lastModifiedTimeV, _ := value.(string)
			option.lastModifiedTime = lastModifiedTimeV

		case "location":
			locationV, _ := value.(string)
			option.location = locationV

		case "selfLink":
			selfLinkV, _ := value.(string)
			option.selfLink = selfLinkV

		case "datasetReference":
			datasetReferenceV, _ := value.(map[string]string)
			option.datasetReference.DatasetID = datasetReferenceV["DatasetID"]
			option.datasetReference.ProjectID = datasetReferenceV["ProjectID"]

		case "datasetReference":
			datasetReferenceV, _ := value.(map[string]string)
			option.datasetReference.DatasetID = datasetReferenceV["DatasetID"]
			option.datasetReference.ProjectID = datasetReferenceV["ProjectID"]

		case "access":
			accessparam, _ := value.(map[string]interface{})

			for accessparamkey, accessparamvalue := range accessparam {
				switch accessparamkey {
					case "MLModelId":
						MLModelIdV, _ := value.(string)
						MLModelId = MLModelIdV

				case "Region":
						RegionV, _ := value.(string)
						Region = RegionV
				}
		}
	}

	createdatasetsjsonmap := make(map[string]interface{})

	createdatasetsdictnoaryconvert(option, createdatasetsjsonmap)

	createdatasetsjson, _ := json.Marshal(createdatasetsjsonmap)

	createdatasetsjsonstring := string(createdatasetsjson)

	var createdatasetsjsonstringbyte = []byte(createdatasetsjsonstring)

	url := "https://www.googleapis.com/dns/v1/projects/" + Project + "/managedZones"

	client := googleauth.SignJWT()

	createdatasetsrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(createdatasetsjsonstringbyte))

	createdatasetsrequest.Header.Set("Content-Type", "application/json")

	createdatasetsrresp, err := client.Do(createdatasetsrequest)

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
