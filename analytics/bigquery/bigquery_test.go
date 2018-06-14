package bigquery

import "testing"
import "fmt"


/*
func TestListDatasets(t *testing.T) {

	var bigquery Bigquery

	listDatasets := map[string]interface{}{
		"ProjectId": "gocloud-206919",
		"All" : true,
		"Filter" : "",
		"MaxResults" : 1,
		"PageToken" : "",
	}

	resp, err := bigquery.ListDatasets(listDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}

*/


/*

/*
func TestDeleteDatasets(t *testing.T) {

	var bigquery Bigquery

	deleteDatasets := map[string]string{
		"ProjectId": "gocloud-206919",
		"DatasetId" : "gocloudv2",
	}

	resp, err := bigquery.DeleteDatasets(deleteDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}

*/



/*
{
 "kind": "bigquery#dataset",
 "etag": "wJ6J76UJduYf9EzfNz0gJw==",
 "id": "gocloud-206919:gocloudv3",
 "selfLink": "https://www.googleapis.com/bigquery/v2/projects/gocloud-206919/datasets/gocloudv3",

 "access": [
  {
   "role": "WRITER",
   "specialGroup": "projectWriters"
  },
  {
   "role": "OWNER",
   "specialGroup": "projectOwners"
  },
  {
   "role": "OWNER",
   "userByEmail": "i.divenire@gmail.com"
  },
  {
   "role": "READER",
   "specialGroup": "projectReaders"
  }
 ],
 }
*/

/*
func TestCreateDatasets(t *testing.T) {

	var bigquery Bigquery

	datasetReference := map[string]string {
	 "DatasetId": "gocloudv4",
	 "ProjectId": "gocloud-206919",
	}

	createDatasets := map[string]interface{}{
		"ProjectId": "gocloud-206919",
		"Description" : "gocloudv4 created",
		"Kind": "bigquery#dataset",
		"Etag": "wJ6J76UJduYf9EzfNz0gJw==",
		"SelfLink": "https://www.googleapis.com/bigquery/v2/projects/gocloud-206919/datasets/gocloudv3",
		 "Id": "gocloud-206919:gocloudv4",
		"Location": "US",
		"DatasetReference" : datasetReference,
	}

	resp, err := bigquery.CreateDatasets(createDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}

*/
func TestUpdateDatasets(t *testing.T) {

	var bigquery Bigquery

	datasetReference := map[string]string {
	 "DatasetId": "gocloudv3",
	 "ProjectId": "gocloud-206919",
	}

	updateDatasets := map[string]interface{}{
		"ProjectId": "gocloud-206919",
		"Description" : "gocloudv3 created",
		"Kind": "bigquery#dataset",
		"Etag": "wJ6J76UJduYf9EzfNz0gJw==",
		"SelfLink": "https://www.googleapis.com/bigquery/v2/projects/gocloud-206919/datasets/gocloudv3",
		 "Id": "gocloud-206919:gocloudv3",
		"Location": "US",
		"DatasetReference" : datasetReference,
		 "DatasetId": "gocloudv3",
	}

	resp, err := bigquery.UpdateDatasets(updateDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}

func TestGetDatasets(t *testing.T) {

	var bigquery Bigquery

	getDatasets := map[string]string{
		"ProjectId": "gocloud-206919",
		"DatasetId" : "gocloudv3",
	}

	resp, err := bigquery.GetDatasets(getDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}
