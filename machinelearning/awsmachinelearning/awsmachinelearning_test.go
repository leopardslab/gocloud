package awsmachinelearning

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}

func TestGetMLModel(t *testing.T) {

	var awsmachinelearning Awsmachinelearning

	getMLModel := map[string]interface{}{
		"Region":    "us-east-1",
		"MLModelId": "ml-EL5FRUNlk72",
	}

	_, err := awsmachinelearning.GetMLModel(getMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteMLModel(t *testing.T) {

	var awsmachinelearning Awsmachinelearning

	deleteMLModel := map[string]interface{}{
		"Region":    "us-east-1",
		"MLModelId": "ml-EL5FRUNlk73",
	}

	_, err := awsmachinelearning.DeleteMLModel(deleteMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}

}

func TestUpdateMLModel(t *testing.T) {

	var awsmachinelearning Awsmachinelearning

	updateMLModel := map[string]interface{}{
		"Region":    "us-east-1",
		"MLModelId": "ml-EL5FRUNlk72",
	}

	_, err := awsmachinelearning.UpdateMLModel(updateMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}

}

func TestCreateMLModel(t *testing.T) {

	var awsmachinelearning Awsmachinelearning

	createMLModel := map[string]interface{}{
		"Region":               "us-east-1",
		"MLModelName":          "EXAMPLE2",
		"MLModelId":            "ml-EL5FRUNlk73",
		"MLModelType":          "REGRESSION",
		"RecipeUri":            "s3://bokya/census.csv",
		"TrainingDataSourceId": "ds-ydIch00SVNu",
	}

	_, err := awsmachinelearning.CreateMLModel(createMLModel)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
