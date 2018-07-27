package clouddataflow

import "testing"
import "fmt"



func TestListStream(t *testing.T) {
	var clouddataflow Clouddataflow

	liststream := map[string]string{
		"Project": "gocloud-206919",
	}

	resp, err := clouddataflow.ListStream(liststream)

	if err != nil {
		t.Errorf("Test Fail")
	}
  response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}





func TestDescribeStream(t *testing.T) {
	var clouddataflow Clouddataflow

	liststream := map[string]string{
		"Project": "gocloud-206919",
    "JobId" : "2018-07-27_08_37_46-11774589915372519551",
	}

	resp, err := clouddataflow.DescribeStream(liststream)

	if err != nil {
		t.Errorf("Test Fail")
	}
  response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}
