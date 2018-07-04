package googlenotification

import (
	"bytes"
	"encoding/json"
	googleauth "github.com/cloudlibz/gocloud/googleauth"
	"io/ioutil"
	"net/http"
)

//ListTopic list topic
func (googlenotification *Googlenotification) ListTopic(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://pubsub.googleapis.com/v1/projects/" + options["Project"] + "/topics"

	client := googleauth.SignJWT()

	listtopicrequest, err := http.NewRequest("GET", url, nil)

	listtopicrequestparam := listtopicrequest.URL.Query()

	if options["PageSize"] != "" {
		listtopicrequestparam.Add("pageSize", options["pageSize"])
	}

	if options["PageToken"] != "" {
		listtopicrequestparam.Add("pageToken", options["PageToken"])
	}

	listtopicrequest.URL.RawQuery = listtopicrequestparam.Encode()
	listtopicrequest.Header.Set("Content-Type", "application/json")

	listtopicresp, err := client.Do(listtopicrequest)
	defer listtopicresp.Body.Close()

	body, err := ioutil.ReadAll(listtopicresp.Body)

	listtopicresponse := make(map[string]interface{})
	listtopicresponse["status"] = listtopicresp.StatusCode
	listtopicresponse["body"] = string(body)
	resp = listtopicresponse
	return resp, err
}

//GetTopic gets topic
func (googlenotification *Googlenotification) GetTopic(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://pubsub.googleapis.com/v1/projects/" + options["Project"] + "/topics/" + options["Topic"]

	client := googleauth.SignJWT()

	gettopicrequest, err := http.NewRequest("GET", url, nil)

	gettopicresp, err := client.Do(gettopicrequest)
	defer gettopicresp.Body.Close()

	body, err := ioutil.ReadAll(gettopicresp.Body)

	gettopicresponse := make(map[string]interface{})
	gettopicresponse["status"] = gettopicresp.StatusCode
	gettopicresponse["body"] = string(body)
	resp = gettopicresponse
	return resp, err
}

//DeleteTopic delete topic
func (googlenotification *Googlenotification) DeleteTopic(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://pubsub.googleapis.com/v1/projects/" + options["Project"] + "/topics/" + options["Topic"]

	client := googleauth.SignJWT()

	gettopicrequest, err := http.NewRequest("DELETE", url, nil)

	gettopicresp, err := client.Do(gettopicrequest)

	defer gettopicresp.Body.Close()

	body, err := ioutil.ReadAll(gettopicresp.Body)

	gettopicresponse := make(map[string]interface{})
	gettopicresponse["status"] = gettopicresp.StatusCode
	gettopicresponse["body"] = string(body)
	resp = gettopicresponse
	return resp, err
}

//CreateTopic creates tpoic
func (googlenotification *Googlenotification) CreateTopic(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://pubsub.googleapis.com/v1/projects/" + options["Project"] + "/topics/" + options["Topic"]

	client := googleauth.SignJWT()

	createtopicjsonmap := make(map[string]interface{})

	createtopicjson, _ := json.Marshal(createtopicjsonmap)

	createtopicjsonstring := string(createtopicjson)

	var createtopicjsonstringbyte = []byte(createtopicjsonstring)

	gettopicrequest, err := http.NewRequest("PUT", url, bytes.NewBuffer(createtopicjsonstringbyte))

	gettopicresp, err := client.Do(gettopicrequest)

	defer gettopicresp.Body.Close()

	body, err := ioutil.ReadAll(gettopicresp.Body)

	gettopicresponse := make(map[string]interface{})
	gettopicresponse["status"] = gettopicresp.StatusCode
	gettopicresponse["body"] = string(body)
	resp = gettopicresponse
	return resp, err
}
