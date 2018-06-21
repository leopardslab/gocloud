package redshift

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"
import "fmt"

func init() {
	awsAuth.LoadConfig()
}

func TestGetDatasets(t *testing.T) {

	var redshift Redshift

	getDatasets := map[string]interface{}{
		"Region": "us-east-1",
	}

	_, err := redshift.GetDatasets(getDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

	//response := resp.(map[string]interface{})

	//fmt.Println(response["body"])

	fmt.Println("hi")
}
