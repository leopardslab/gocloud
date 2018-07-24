package lambda

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"
import "fmt"

func init() {
	awsAuth.LoadConfig()
}


func TestDeleteFunction(t *testing.T) {

	var lambda Lambda

	deletedunction := map[string]interface{}{
		"Region":     "us-east-1",
		"FunctionName": "gocloud",
    "Qualifier" : "Qualifier",
	}

	resp, err := lambda.DeleteFunction(deletedunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}
