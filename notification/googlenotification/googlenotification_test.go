package googlenotification

import "testing"
import "fmt"

/*

func TestListTopic(t *testing.T) {
	var googlenotification Googlenotification

	listtopic := map[string]string{
		"Project": "gocloud-206919",
    "PageSize" : "",
    "PageToken": "",
	}
	resp, err := googlenotification.ListTopic(listtopic)

	if err != nil {
		t.Errorf("Test Fail")
	}


  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
}


func TestGetTopic(t *testing.T) {
	var googlenotification Googlenotification

	gettopic := map[string]string{
		"Project": "gocloud-206919",
    "Topic" : "gocloud",
	}

	resp, err := googlenotification.GetTopic(gettopic)

	if err != nil {
		t.Errorf("Test Fail")
	}


  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
}





func TestDeleteTopic(t *testing.T) {
	var googlenotification Googlenotification

	deletetopic := map[string]string{
		"Project": "gocloud-206919",
    "Topic" : "gocloud",
	}

	resp, err := googlenotification.DeleteTopic(deletetopic)

	if err != nil {
		t.Errorf("Test Fail")
	}


  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
}


*/


func TestCreateTopic(t *testing.T) {
	var googlenotification Googlenotification

	createtopic := map[string]string{
		"Project": "gocloud-206919",
    "Topic" : "gocloud3",
	}

	resp, err := googlenotification.CreateTopic(createtopic)

  fmt.Println("Hi")

	if err != nil {
		t.Errorf("Test Fail")
	}


  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
}
