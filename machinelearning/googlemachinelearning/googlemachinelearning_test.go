package googlemachinelearning

import "testing"
import "fmt"



func TestUpdateMLModel(t *testing.T) {

	var googlemachinelearning Googlemachinelearning

	defaultVersion := make(map[string]interface{})

	defaultVersion["Name"] = "bokkya"

	updateMLModel := map[string]interface{}{
		"Parent": "projects/adept-comfort-202709/models/pratik",
		"Name"  :"pratik",
		"DefaultVersion" :defaultVersion,
		"UpdateMask" : "description",
	}

	resp, err := googlemachinelearning.UpdateMLModel(updateMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}


/*
func TestCreateMLModel(t *testing.T) {

	var googlemachinelearning Googlemachinelearning

	defaultVersion := make(map[string]interface{})

	defaultVersion["Name"] = "bokkya"
	defaultVersion["DeploymentUri"] = "gs:bokkya"

	createMLModel := map[string]interface{}{
		"Parent": "projects/adept-comfort-202709",
		"Name"  :"pratik",
		"DefaultVersion" :defaultVersion,
	}

	resp, err := googlemachinelearning.CreateMLModel(createMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}


func TestGetMLModel(t *testing.T) {

	var googlemachinelearning Googlemachinelearning

	getMLModel := map[string]string{
		"name": "projects/adept-comfort-202709/models/hello",
	}

	resp, err := googlemachinelearning.GetMLModel(getMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}



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
