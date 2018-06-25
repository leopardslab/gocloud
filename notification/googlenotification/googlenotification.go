package googlenotification


func (googledns *Googledns) ListTopic(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://pubsub.googleapis.com/v1/" + options["Project"] + "/subscriptions"

	client := googleauth.SignJWT()

	listtopicrequest, err := http.NewRequest("GET", url, nil)

	listtopicrequestparam := listtopicrequest.URL.Query()

	if options["pageSize"] != "" {
		listtopicrequestparam.Add("pageSize", options["pageSize"])
	}

	if options["pageToken"] != "0" {
		listtopicrequestparam.Add("pageToken", options["pageToken"])
	}

	listtopicrequest.URL.RawQuery = listtopicrequestparam.Encode()

	listtopicrequest.Header.Set("Content-Type", "application/json")

	listtopicresp, err := client.Do(Listdnsrequest)

	defer listtopicresp.Body.Close()

	body, err := ioutil.ReadAll(listtopicresp.Body)

	listtopicresponse := make(map[string]interface{})
	listtopicresponse["status"] = listtopicresp.StatusCode
	listtopicresponse["body"] = string(body)
	resp = listtopicresponse
	return resp, err
}


func (googledns *Googledns) GetTopic(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://pubsub.googleapis.com/" + options["Project"] + "/subscriptions/" + options["Subscription"]

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



func (googledns *Googledns) DeleteTopic(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://pubsub.googleapis.com/" + options["Project"] + "/subscriptions/" + options["Subscription"]

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
