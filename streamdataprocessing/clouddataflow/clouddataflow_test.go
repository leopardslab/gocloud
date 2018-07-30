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

	describedtream := map[string]string{
		"Project": "gocloud-206919",
		"JobId":   "2018-07-27_08_37_46-11774589915372519551",
	}

	resp, err := clouddataflow.DescribeStream(describedtream)

	if err != nil {
		t.Errorf("Test Fail")
	}
	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}




func TestCreateStream(t *testing.T) {
	var clouddataflow Clouddataflow


		userAgent :=make(map[string]interface{})

		support := make(map[string]interface{})

		support["Status"]= "STALE"
    support["URL"] = "https://cloud.google.com/dataflow/support"


		userAgent["Name"] = "Google Cloud Dataflow SDK for Java"

		userAgent["BuildDate"] = "2017-09-01 05:54"

 		userAgent["Version"] = "2.1.0"

		userAgent["Support"] = support

		version := make(map[string]interface{})

		version["Major"] =  "6"

		version["JobType"] =  "JAVA_BATCH_AUTOSCALING"

		environment := make(map[string]interface{})

		environment["UserAgent"] = userAgent
		environment["Version"] = version

		stageStates :=  []map[string]interface{}{
			{
		"ExecutionStageName": "F19",
		"ExecutionStageState": "JOB_STATE_DONE",
		"CurrentStateTime": "2018-07-27T15:39:00.225Z",
	},
	{
		"ExecutionStageName": "s10",
		"ExecutionStageState": "JOB_STATE_DONE",
		"CurrentStateTime": "2018-07-27T15:39:00.430Z",
	},
	{
		"ExecutionStageName": "s44-close-shuffle15",
		"ExecutionStageState": "JOB_STATE_DONE",
		"CurrentStateTime": "2018-07-27T15:38:51.161Z",
	},
	{
		"ExecutionStageName": "F20",
		"ExecutionStageState": "JOB_STATE_DONE",
		"CurrentStateTime": "2018-07-27T15:38:51.072Z",
	},
	{
		"ExecutionStageName": "F18",
		"ExecutionStageState": "JOB_STATE_DONE",
		"CurrentStateTime": "2018-07-27T15:39:05.296Z",
	},
	{
		"ExecutionStageName": "s44-open-shuffle13",
		"ExecutionStageState": "JOB_STATE_DONE",
		"CurrentStateTime": "2018-07-27T15:37:56.465Z",
	},
}

	createstream :=  map[string]interface{}{

		"ProjectID": "gocloud-206919",
		"View":   "JOB_VIEW_ALL",
		"Location" : "us-central1",
		"CurrentState" : "JOB_STATE_DONE",
		"ClientRequestID"  :"20180727153742037_7574",
		"ID" : "2018-07-27_08_37_46-11774589915372519551",
		"Name" : "dataflow-intro",
		"Type" :  "JOB_TYPE_BATCH",
		"StageStates" : stageStates,
		"Environment" : environment,
	}

	resp, err := clouddataflow.CreateStream(createstream)

	if err != nil {
		t.Errorf("Test Fail")
	}
	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}



func TestUpdateStream(t *testing.T) {

		var clouddataflow Clouddataflow

		userAgent :=make(map[string]interface{})

		support := make(map[string]interface{})

		support["Status"]= "STALE"
    support["URL"] = "https://cloud.google.com/dataflow/support"


		userAgent["Name"] = "Google Cloud Dataflow SDK for Java"

		userAgent["BuildDate"] = "2017-09-01 05:54"

 		userAgent["Version"] = "2.1.0"

		userAgent["Support"] = support

		version := make(map[string]interface{})

		version["Major"] =  "6"

		version["JobType"] =  "JAVA_BATCH_AUTOSCALING"

		environment := make(map[string]interface{})

		environment["UserAgent"] = userAgent
		environment["Version"] = version

		stageStates :=  []map[string]interface{}{
			{
		"ExecutionStageName": "F19",
		"ExecutionStageState": "JOB_STATE_DONE",
		"CurrentStateTime": "2018-07-27T15:39:00.225Z",
	},
	{
		"ExecutionStageName": "s10",
		"ExecutionStageState": "JOB_STATE_DONE",
		"CurrentStateTime": "2018-07-27T15:39:00.430Z",
	},
	{
		"ExecutionStageName": "s44-close-shuffle15",
		"ExecutionStageState": "JOB_STATE_DONE",
		"CurrentStateTime": "2018-07-27T15:38:51.161Z",
	},
	{
		"ExecutionStageName": "F20",
		"ExecutionStageState": "JOB_STATE_DONE",
		"CurrentStateTime": "2018-07-27T15:38:51.072Z",
	},
	{
		"ExecutionStageName": "F18",
		"ExecutionStageState": "JOB_STATE_DONE",
		"CurrentStateTime": "2018-07-27T15:39:05.296Z",
	},
	{
		"ExecutionStageName": "s44-open-shuffle13",
		"ExecutionStageState": "JOB_STATE_DONE",
		"CurrentStateTime": "2018-07-27T15:37:56.465Z",
	},
}

	updatestream :=  map[string]interface{}{

		"ProjectID": "gocloud-206919",
		"View":   "JOB_VIEW_ALL",
		"Location" : "us-central1",
		"CurrentState" : "JOB_STATE_DONE",
		"ClientRequestID"  :"20180727153742037_7574",
		"ID" : "2018-07-27_08_37_46-11774589915372519551",
		"Name" : "dataflow-intro",
		"Type" :  "JOB_TYPE_BATCH",
		"StageStates" : stageStates,
		"Environment" : environment,
	}

	resp, err := clouddataflow.UpdateStream(updatestream)

	if err != nil {
		t.Errorf("Test Fail")
	}
	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}
