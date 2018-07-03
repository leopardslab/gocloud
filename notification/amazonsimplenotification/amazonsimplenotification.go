package amazonsimplenotification

func (amazonsimplenotification *Amazonsimplenotification) GetTopic(request interface{}) (resp interface{}, err error) {
	return resp, err
}

func (amazonsimplenotification *Amazonsimplenotification) ListTopic(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]string)
	nextToken := param["NextToken"]
	region := param["Region"]

	listtopicpram := make(map[string]string)

	preparedefaultlisttopicpram(listtopicpram)

	preparelisttopicpram(listtopicpram, nextToken)

	response := make(map[string]interface{})

	err = amazonsimplenotification.PrepareSignatureV2query(listtopicpram, region, response)

	if err != nil {
		return nil, err
	}

	resp = response

	return resp, err
}

func (amazonsimplenotification *Amazonsimplenotification) CreateTopic(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]string)
	name := param["Name"]
	region := param["Region"]

	createtopicpram := make(map[string]string)

	preparedefaultcreatetopicpram(createtopicpram)

	preparecreatetopicpram(createtopicpram, name)

	response := make(map[string]interface{})

	err = amazonsimplenotification.PrepareSignatureV2query(createtopicpram, region, response)

	if err != nil {
		return nil, err
	}

	resp = response

	return resp, err
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

func (amazonsimplenotification *Amazonsimplenotification) DeleteTopic(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]string)
	topicArn := param["TopicArn"]
	region := param["Region"]

	deletetopicpram := make(map[string]string)

	preparedefaultdeletetopicpram(deletetopicpram)

	preparedeletetopicpram(deletetopicpram, topicArn)

	response := make(map[string]interface{})

	err = amazonsimplenotification.PrepareSignatureV2query(deletetopicpram, region, response)

	if err != nil {
		return nil, err
	}

	resp = response

	return resp, err
}

func preparedefaultdeletetopicpram(deletetopicpram map[string]string) {

	deletetopicpram["Action"] = "DeleteTopic"
}

func preparedeletetopicpram(deletetopicpram map[string]string, topicArn string) {

	if topicArn != "" {
		deletetopicpram["TopicArn"] = topicArn
	}
}
