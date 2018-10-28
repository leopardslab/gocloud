package bigquery

import (
	"bytes"
	"encoding/json"
	googleauth "github.com/cloudlibz/gocloud/googleauth"
	"io/ioutil"
	"net/http"
	"time"
)

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

		case "CreationTime":
			creationTime := time.Now().UTC().Format(time.RFC3339)
			option.creationTime = creationTime

		case "DefaultTableExpirationMs":
			defaultTableExpirationMsV, _ := value.(string)
			option.defaultTableExpirationMs = defaultTableExpirationMsV

		case "Description":
			descriptionV, _ := value.(string)
			option.description = descriptionV

		case "Etag":
			etagV, _ := value.(string)
			option.etag = etagV

		case "FriendlyName":
			friendlyNameV, _ := value.(string)
			option.friendlyName = friendlyNameV

		case "Id":
			idV, _ := value.(string)
			option.id = idV

		case "Kind":
			kindV, _ := value.(string)
			option.kind = kindV

		case "LastModifiedTime":
			lastModifiedTimeV, _ := value.(string)
			option.lastModifiedTime = lastModifiedTimeV

		case "Location":
			locationV, _ := value.(string)
			option.location = locationV

		case "SelfLink":
			selfLinkV, _ := value.(string)
			option.selfLink = selfLinkV

		case "DatasetReference":
			datasetReferenceV, _ := value.(map[string]string)
			option.datasetReference.datasetID = datasetReferenceV["DatasetID"]
			option.datasetReference.projectID = datasetReferenceV["ProjectID"]

		case "Access":
			accessparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(accessparam); i++ {
				var access Access
				for accessparamkey, accessparamvalue := range accessparam[i] {
					switch accessparamkey {
					case "Domain":
						DomainV, _ := accessparamvalue.(string)
						access.domain = DomainV

					case "GroupByEmail":
						GroupByEmailV, _ := accessparamvalue.(string)
						access.groupByEmail = GroupByEmailV

					case "Role":
						RoleV, _ := accessparamvalue.(string)
						access.role = RoleV

					case "SpecialGroup":
						SpecialGroupV, _ := accessparamvalue.(string)
						access.specialGroup = SpecialGroupV

					case "UserByEmail":
						UserByEmailV, _ := accessparamvalue.(string)
						access.userByEmail = UserByEmailV

					case "View":
						viewparam, _ := value.(map[string]interface{})
						for viewparamkey, viewparamvalue := range viewparam {
							switch viewparamkey {
							case "ProjectID":
								projectIDV, _ := viewparamvalue.(string)
								access.view.projectID = projectIDV

							case "DatasetID":
								datasetIDV, _ := viewparamvalue.(string)
								access.view.datasetID = datasetIDV

							case "TableID":
								tableIDV, _ := viewparamvalue.(string)
								access.view.tableID = tableIDV
							}
						}
					}
					option.access = append(option.access, access)
				}
			}
		}
	}

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

	url := "https://www.googleapis.com/bigquery/v2/projects/" + options["ProjectId"] + "/datasets/" + options["DatasetId"]

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

	url := "https://www.googleapis.com/bigquery/v2/projects/" + options["ProjectId"] + "/datasets/" + options["DatasetId"]

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

		case "DatasetId":
			datasetIdV, _ := value.(string)
			datasetId = datasetIdV

		case "CreationTime":
			creationTime := time.Now().UTC().Format(time.RFC3339)
			option.creationTime = creationTime

		case "DefaultTableExpirationMs":
			defaultTableExpirationMsV, _ := value.(string)
			option.defaultTableExpirationMs = defaultTableExpirationMsV

		case "Description":
			descriptionV, _ := value.(string)
			option.description = descriptionV

		case "Etag":
			etagV, _ := value.(string)
			option.etag = etagV

		case "FriendlyName":
			friendlyNameV, _ := value.(string)
			option.friendlyName = friendlyNameV

		case "Id":
			idV, _ := value.(string)
			option.id = idV

		case "Kind":
			kindV, _ := value.(string)
			option.kind = kindV

		case "LastModifiedTime":
			lastModifiedTimeV, _ := value.(string)
			option.lastModifiedTime = lastModifiedTimeV

		case "Location":
			locationV, _ := value.(string)
			option.location = locationV

		case "SelfLink":
			selfLinkV, _ := value.(string)
			option.selfLink = selfLinkV

		case "DatasetReference":
			datasetReferenceV, _ := value.(map[string]string)
			option.datasetReference.datasetID = datasetReferenceV["DatasetID"]
			option.datasetReference.projectID = datasetReferenceV["ProjectID"]

		case "Access":
			accessparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(accessparam); i++ {
				var access Access
				for accessparamkey, accessparamvalue := range accessparam[i] {
					switch accessparamkey {
					case "Domain":
						DomainV, _ := accessparamvalue.(string)
						access.domain = DomainV

					case "GroupByEmail":
						GroupByEmailV, _ := accessparamvalue.(string)
						access.groupByEmail = GroupByEmailV

					case "Role":
						RoleV, _ := accessparamvalue.(string)
						access.role = RoleV

					case "SpecialGroup":
						SpecialGroupV, _ := accessparamvalue.(string)
						access.specialGroup = SpecialGroupV

					case "UserByEmail":
						UserByEmailV, _ := accessparamvalue.(string)
						access.userByEmail = UserByEmailV

					case "View":
						viewparam, _ := value.(map[string]interface{})
						for viewparamkey, viewparamvalue := range viewparam {
							switch viewparamkey {
							case "ProjectID":
								projectIDV, _ := viewparamvalue.(string)
								access.view.projectID = projectIDV

							case "DatasetID":
								datasetIDV, _ := viewparamvalue.(string)
								access.view.datasetID = datasetIDV

							case "TableID":
								tableIDV, _ := viewparamvalue.(string)
								access.view.tableID = tableIDV
							}
						}
					}
					option.access = append(option.access, access)
				}
			}
		}
	}

	createdatasetsjsonmap := make(map[string]interface{})

	createdatasetsdictnoaryconvert(option, createdatasetsjsonmap)

	updatedatasetsjson, _ := json.Marshal(createdatasetsjsonmap)

	updatedatasetsjsonstring := string(updatedatasetsjson)

	var updatedatasetsjsonstringbyte = []byte(updatedatasetsjsonstring)

	url := "https://www.googleapis.com/bigquery/v2/projects/" + projectId + "/datasets/" + datasetId

	client := googleauth.SignJWT()

	updatedatasetsrequest, err := http.NewRequest("PUT", url, bytes.NewBuffer(updatedatasetsjsonstringbyte))

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
	options := request.(map[string]interface{})

	var projectId, pageToken, filter string
	var all bool
	var maxResults int

	for key, value := range options {
		switch key {
		case "ProjectId":
			projectIdV, _ := value.(string)
			projectId = projectIdV

		case "All":
			allV, _ := value.(bool)
			all = allV

		case "Filter":
			filterV, _ := value.(string)
			filter = filterV

		case "PageToken":
			pageTokenV, _ := value.(string)
			pageToken = pageTokenV

		case "MaxResults":
			maxResultsV, _ := value.(int)
			maxResults = maxResultsV
		}
	}

	url := "https://www.googleapis.com/bigquery/v2/projects/" + projectId + "/datasets"

	client := googleauth.SignJWT()

	listdatasetsrequest, err := http.NewRequest("GET", url, nil)

	listdatasetsrequestparam := listdatasetsrequest.URL.Query()

	if all != false {
		listdatasetsrequestparam.Add("all", "true")
	}

	if filter != "" {
		listdatasetsrequestparam.Add("filter", filter)
	}

	if maxResults != 0 {
		listdatasetsrequestparam.Add("maxResults", "1")
	}

	if pageToken != "" {
		listdatasetsrequestparam.Add("pageToken", pageToken)
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
