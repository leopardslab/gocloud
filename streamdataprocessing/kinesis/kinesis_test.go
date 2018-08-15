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
		"Region": "us-east-1",
	}

	_, err := kinesis.ListStream(liststream)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestCreateStream(t *testing.T) {

	var kinesis Kinesis

	createstream := map[string]interface{}{
		"Region":     "us-east-1",
		"StreamName": "gocloud",
		"ShardCount": 1,
	}

	_, err := kinesis.CreateStream(createstream)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDescribeStream(t *testing.T) {

	var kinesis Kinesis

	describestream := map[string]interface{}{
		"Region":     "us-east-1",
		"StreamName": "gocloud",
	}

	_, err := kinesis.DescribeStream(describestream)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteStream(t *testing.T) {

	var kinesis Kinesis

	deletestream := map[string]interface{}{
		"Region":     "us-east-1",
		"StreamName": "gocloud",
	}

	_, err := kinesis.DeleteStream(deletestream)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
