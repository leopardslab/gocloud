package bigquery

import "testing"
import "fmt"

/*
func TestUpdateMLModel(t *testing.T) {

	var googlemachinelearning Googlemachinelearning

	defaultVersion := make(map[string]interface{})

	defaultVersion["Name"] = "bokkya"

	updateMLModel := map[string]interface{}{
		"Parent":         "projects/adept-comfort-202709/models/pratik",
		"Name":           "pratik",
		"DefaultVersion": defaultVersion,
		"UpdateMask":     "description",
	}

	resp, err := googlemachinelearning.UpdateMLModel(updateMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}

func TestCreateMLModel(t *testing.T) {

	var googlemachinelearning Googlemachinelearning

	defaultVersion := make(map[string]interface{})

	defaultVersion["Name"] = "bokkya"
	defaultVersion["DeploymentUri"] = "gs:bokkya"

	createMLModel := map[string]interface{}{
		"Parent":         "projects/adept-comfort-202709",
		"Name":           "pratik",
		"DefaultVersion": defaultVersion,
	}

	resp, err := googlemachinelearning.CreateMLModel(createMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}

*/

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

func TestGetDatasets(t *testing.T) {

	var bigquery Bigquery

	getDatasets := map[string]string{
		"ProjectId": "gocloud-206919",
		"DatasetId" : "getd",
	}

	resp, err := bigquery.GetDatasets(getDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}

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



/*
func TestDeleteMLModel(t *testing.T) {

	var googlemachinelearning Googlemachinelearning

	deleteMLModel := map[string]string{
		"name": "projects/adept-comfort-202709/models/hi",
	}

	resp, err := googlemachinelearning.DeleteMLModel(deleteMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}
*/
