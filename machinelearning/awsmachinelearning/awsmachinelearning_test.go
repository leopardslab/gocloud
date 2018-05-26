package awsmachinelearning

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"
import "fmt"

func init() {
	awsAuth.LoadConfig()
}

    func TestGetMLModel(t *testing.T) {

      var awsmachinelearning Awsmachinelearning

      getMLModel := map[string]interface{}{
        "Region":    "us-east-1",
        "MLModelId": "ml-EL5FRUNlk7p",
      }

    	resp, err := awsmachinelearning.GetMLModel(getMLModel)

    	if err != nil {
    		t.Errorf("Test Fail")
    	}

    	response := resp.(map[string]interface{})

    	fmt.Println(response["body"])
    }




func TestDeleteMLModel(t *testing.T) {

  var awsmachinelearning Awsmachinelearning

  deleteMLModel := map[string]interface{}{
    "Region":    "us-east-1",
    "MLModelId": "ds-Lf3D4KaPukx",
  }

  resp, err := awsmachinelearning.DeleteMLModel(deleteMLModel)

  if err != nil {
    t.Errorf("Test Fail")
  }

  response := resp.(map[string]interface{})

  fmt.Println(response["body"])
}



func TestUpdateMLModel(t *testing.T) {

  var awsmachinelearning Awsmachinelearning

  updateMLModel := map[string]interface{}{
    "Region":    "us-east-1",
    "MLModelId": "ds-Lf3D4KaPukx",
  }

  resp, err := awsmachinelearning.UpdateMLModel(updateMLModel)

  if err != nil {
    t.Errorf("Test Fail")
  }

  response := resp.(map[string]interface{})

  fmt.Println(response["body"])
}
