package amazonsimplenotification

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}

func TestCreateTopic(t *testing.T) {
	var amazonsimplenotification Amazonsimplenotification

	createtopic := map[string]string{
		"Region": "us-east-1",
		"Name":   "Bokya",
	}

	_, err := amazonsimplenotification.CreateTopic(createtopic)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteTopic(t *testing.T) {
	var amazonsimplenotification Amazonsimplenotification

	deletetopic := map[string]string{
		"Region":   "us-east-1",
		"TopicArn": "arn:aws:sns:us-east-1:478991680879:Bokya",
	}

	_, err := amazonsimplenotification.DeleteTopic(deletetopic)

	if err != nil {
		t.Errorf("Test Fail")
	}

}

func TestListTopic(t *testing.T) {
	var amazonsimplenotification Amazonsimplenotification

	listtopic := map[string]string{
		"Region": "us-east-1",
	}

	_, err := amazonsimplenotification.ListTopic(listtopic)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
