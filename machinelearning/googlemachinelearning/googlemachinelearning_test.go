package googlemachinelearning

import "testing"
import "fmt"

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
		"name": "projects/adept-comfort-202709/models/hello",
	}

	resp, err := googlemachinelearning.DeleteMLModel(deleteMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}
