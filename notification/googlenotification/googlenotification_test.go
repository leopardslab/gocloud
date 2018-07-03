package googlenotification

import "testing"


func TestListTopic(t *testing.T) {
	var googlenotification Googlenotification

	listtopic := map[string]string{
		"Project": "gocloud-206919",
    "PageSize" : "",
    "PageToken": "",
	}
	_, err := googlenotification.ListTopic(listtopic)

	if err != nil {
		t.Errorf("Test Fail")
	}
}


func TestGetTopic(t *testing.T) {
	var googlenotification Googlenotification

	gettopic := map[string]string{
		"Project": "gocloud-206919",
    "Topic" : "gocloud",
	}

	_, err := googlenotification.GetTopic(gettopic)

	if err != nil {
		t.Errorf("Test Fail")
	}
}



func TestDeleteTopic(t *testing.T) {
	var googlenotification Googlenotification

	deletetopic := map[string]string{
		"Project": "gocloud-206919",
    "Topic" : "gocloud",
	}

	_, err := googlenotification.DeleteTopic(deletetopic)

	if err != nil {
		t.Errorf("Test Fail")
	}
}



func TestCreateTopic(t *testing.T) {
	var googlenotification Googlenotification

	createtopic := map[string]string{
		"Project": "gocloud-206919",
    "Topic" : "gocloud3",
	}

	_, err := googlenotification.CreateTopic(createtopic)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
