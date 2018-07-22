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


func TestCreateStream(t *testing.T) {

	var kinesis Kinesis

	createstream := map[string]interface{}{
		"Region":    "us-east-1",
		"StreamName" : "gocloud",
		"ShardCount" : 1,
	}

	resp, err := kinesis.CreateStream(createstream)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}


func TestDescribeStream(t *testing.T) {

	var kinesis Kinesis

	describestream := map[string]interface{}{
		"Region":    "us-east-1",
		"StreamName" : "gocloud",
	}

	resp, err := kinesis.DescribeStream(describestream)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}


func TestDeleteStream(t *testing.T) {

	var kinesis Kinesis

	deletestream := map[string]interface{}{
		"Region":    "us-east-1",
		"StreamName" : "gocloud",
	}

	resp, err := kinesis.DeleteStream(deletestream)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}
