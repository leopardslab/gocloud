package kinesis

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"
import "fmt"

func init() {
	awsAuth.LoadConfig()
}

func TestListStream(t *testing.T) {

	var kinesis Kinesis

	liststream := map[string]interface{}{
		"Region":    "us-east-1",
	}

	resp, err := kinesis.ListStream(liststream)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}
