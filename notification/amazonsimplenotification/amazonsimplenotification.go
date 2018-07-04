package amazonsimplenotification

//ListTopic list topic
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

//CreateTopic create topic
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

//DeleteTopic deletetopic
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


//GetTopic gettopic
func (amazonsimplenotification *Amazonsimplenotification) GetTopic(request interface{}) (resp interface{}, err error) {

	return resp, err
}
