package awsmachinelearning

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}

    func TestDeleteMLModel(t *testing.T) {

      var awsmachinelearning Awsmachinelearning

      getMLModel := map[string]interface{}{
        "Region":    "us-east-1",
        "MLModelId": "ml-EL5FRUNlk7p",
      }

    	_, err := awsmachinelearning.GetMLModel(getMLModel)

    	if err != nil {
    		t.Errorf("Test Fail")
    	}
    }
