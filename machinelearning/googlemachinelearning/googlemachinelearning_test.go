package googlemachinelearning

import "testing"

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

	_, err := googlemachinelearning.UpdateMLModel(updateMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}

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

	_, err := googlemachinelearning.CreateMLModel(createMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}

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

}

func TestDeleteMLModel(t *testing.T) {

	var googlemachinelearning Googlemachinelearning

	deleteMLModel := map[string]string{
		"name": "projects/adept-comfort-202709/models/hi",
	}

	_, err := googlemachinelearning.DeleteMLModel(deleteMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}

}
