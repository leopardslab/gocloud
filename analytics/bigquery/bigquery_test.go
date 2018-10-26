package bigquery

import "testing"
import "fmt"

func TestListDatasets(t *testing.T) {
	var bigquery Bigquery

	listDatasets := map[string]interface{}{
		"ProjectId":  "gocloud-206919",
		"All":        true,
		"Filter":     "",
		"MaxResults": 1,
		"PageToken":  "",
	}

	_, err := bigquery.ListDatasets(listDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteDatasets(t *testing.T) {
	var bigquery Bigquery

	deleteDatasets := map[string]string{
		"ProjectId": "gocloud-206919",
		"DatasetId": "gocloudv2",
	}

	_, err := bigquery.DeleteDatasets(deleteDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestCreateDatasets(t *testing.T) {
	var bigquery Bigquery

	datasetReference := map[string]string{
		"DatasetId": "gocloudv4",
		"ProjectId": "gocloud-206919",
	}

	createDatasets := map[string]interface{}{
		"ProjectId":        "gocloud-206919",
		"Description":      "gocloudv4 created",
		"Kind":             "bigquery#dataset",
		"Etag":             "wJ6J76UJduYf9EzfNz0gJw==",
		"SelfLink":         "https://www.googleapis.com/bigquery/v2/projects/gocloud-206919/datasets/gocloudv3",
		"Id":               "gocloud-206919:gocloudv4",
		"Location":         "US",
		"DatasetReference": datasetReference,
	}

	_, err := bigquery.CreateDatasets(createDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestUpdateDatasets(t *testing.T) {
	var bigquery Bigquery

	datasetReference := map[string]string{
		"DatasetId": "gocloudv3",
		"ProjectId": "gocloud-206919",
	}

	updateDatasets := map[string]interface{}{
		"ProjectId":        "gocloud-206919",
		"Description":      "gocloudv3 created",
		"Kind":             "bigquery#dataset",
		"Etag":             "wJ6J76UJduYf9EzfNz0gJw==",
		"SelfLink":         "https://www.googleapis.com/bigquery/v2/projects/gocloud-206919/datasets/gocloudv3",
		"Id":               "gocloud-206919:gocloudv3",
		"Location":         "US",
		"DatasetReference": datasetReference,
		"DatasetId":        "gocloudv3",
	}

	_, err := bigquery.UpdateDatasets(updateDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

}

func TestGetDatasets(t *testing.T) {
	var bigquery Bigquery

	getDatasets := map[string]string{
		"ProjectId": "gocloud-206919",
		"DatasetId": "gocloudv3",
	}

	_, err := bigquery.GetDatasets(getDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
