package amazonsimplenotification

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"
import "fmt"

func init() {
	awsAuth.LoadConfig()
}

/*

func TestCreateTopic(t *testing.T) {
	var amazonsimplenotification Amazonsimplenotification

	create := map[string]string{
		"Region": "us-east-1",
		"Name": "Bokya",
	}

	resp, err := amazonsimplenotification.CreateTopic(create)

	if err != nil {
		t.Errorf("Test Fail")
	}


  response := resp.(map[string]interface{})

	fmt.Println(response["body"])

	fmt.Println("hi")
}

*/




func TestDeleteTopic(t *testing.T) {
	var amazonsimplenotification Amazonsimplenotification

	create := map[string]string{
		"Region": "us-east-1",
		"TopicArn" :"arn:aws:sns:us-east-1:478991680879:Bokya",
	}

	resp, err := amazonsimplenotification.DeleteTopic(create)

	if err != nil {
		t.Errorf("Test Fail")
	}


  response := resp.(map[string]interface{})

	fmt.Println(response["body"])

	fmt.Println("hi")
}



func TestListTopic(t *testing.T) {
	var amazonsimplenotification Amazonsimplenotification

	create := map[string]string{
		"Region": "us-east-1",

	}

	resp, err := amazonsimplenotification.ListTopic(create)

	if err != nil {
		t.Errorf("Test Fail")
	}


  response := resp.(map[string]interface{})

	fmt.Println(response["body"])

	fmt.Println("hi\n\n\n")
}
