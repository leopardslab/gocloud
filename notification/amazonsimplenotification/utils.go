package amazonsimplenotification

import (
	"github.com/cloudlibz/gocloud/auth"
	awsAuth "github.com/cloudlibz/gocloud/awsauth"
	"io/ioutil"
	"net/http"
	"time"
)

var timeNow = time.Now

func (amazonsimplenotification *Amazonsimplenotification) PrepareSignatureV2query(params map[string]string, Region string, response map[string]interface{}) error {

	EC2Endpoint := "https://sns." + Region + ".amazonaws.com"

	req, err := http.NewRequest("GET", EC2Endpoint, nil)
	if err != nil {
		return err
	}

	// Add the params passed in to the query string
	query := req.URL.Query()
	for varName, varVal := range params {
		query.Add(varName, varVal)
	}
	query.Add("Timestamp", timeNow().In(time.UTC).Format(time.RFC3339))

	req.URL.RawQuery = query.Encode()
	auth := map[string]string{"AccessKey": auth.Config.AWSAccessKeyID, "SecretKey": auth.Config.AWSSecretKey}
	awsAuth.SignatureV2(req, auth)
	r, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	response["body"] = string(body)
	response["status"] = r.StatusCode

	return err
}

func preparedefaultcreatetopicpram(createtopicpram map[string]string) {

	createtopicpram["Action"] = "CreateTopic"
}

func preparecreatetopicpram(createtopicpram map[string]string, name string) {

	if name != "" {
		createtopicpram["Name"] = name
	}
}

func preparedefaultlisttopicpram(listtopicpram map[string]string) {

	listtopicpram["Action"] = "ListTopics"
}

func preparelisttopicpram(listtopicpram map[string]string, nextToken string) {

	if nextToken != "" {
		listtopicpram["NextToken"] = nextToken
	}
}


func preparedefaultdeletetopicpram(deletetopicpram map[string]string) {

	deletetopicpram["Action"] = "DeleteTopic"
}

func preparedeletetopicpram(deletetopicpram map[string]string, topicArn string) {

	if topicArn != "" {
		deletetopicpram["TopicArn"] = topicArn
	}
}
