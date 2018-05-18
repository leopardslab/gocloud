package awsmachinelearning

func(awsmachinelearning *Awsmachinelearning) CreateMLModel(request interface{}) (resp interface{}, err error) {

	return resp, err
}

func(awsmachinelearning *Awsmachinelearning) DeleteMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var MLModelId, Region string

	for key, value := range param {
		switch key {
			case "MLModelId":
				TableNameV, _ := value.(string)
				TableName = TableNameV

			case "Region":
				RegionV, _ := value.(string)
				Region = RegionV
		}
	}

	params := make(map[string]string)


	preparedeletemodel(params, MLModelId, Region)

	deletemodeljsonmap := map[string]interface{}{
		"MLModelId": MLModelId ,
	}

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, deletemodeljsonmap, response)
	resp = response
	return resp, err
}

func preparedeletemodel(params map[string]string, MLModelId string,Region string) {

	if MLModelId != "" {
		params["MLModelId"] = MLModelId
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "AmazonML_20141212.DeleteMLModel"
}

func preparegetmodel(params map[string]string, MLModelId string ,Verbose string,Region string) {

	if MLModelId != "" {
		params["MLModelId"] = MLModelId
	}

	if Verbose != "" {
		params["Verbose"] = Verbose
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "AmazonML_20141212.GetMLModel"
}

func prepareupdatemodel(params map[string]string, MLModelId string ,ScoreThreshold string, MLModelName string,Region string) {

	if MLModelId != "" {
		params["MLModelId"] = MLModelId
	}

	if MLModelName != "" {
		params["MLModelName"] = MLModelName
	}

	if ScoreThreshold != "" {
		params["ScoreThreshold"] = ScoreThreshold
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "AmazonML_20141212.UpdateMLModel"
}



//PrepareSignatureV4query creates PrepareSignatureV4 for request.
func (dynamodb *Dynamodb) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}, response map[string]interface{}) error {
	ECSEndpoint := "https://machinelearning." + params["Region"] + ".amazonaws.com"
	fmt.Println("ECSEndpoint : ",ECSEndpoint)
	service := "machinelearning"
	method := "POST"
	host := service + "." + params["Region"] + ".amazonaws.com"
	fmt.Println("host : ",host)
	ContentType := "application/x-amz-json-1.1"
	signedheaders := "content-type;host;x-amz-date;x-amz-target"
	amztarget := params["amztarget"]

	requestparametersjson, _ := json.Marshal(paramsmap)
	requestparametersjsonstring := string(requestparametersjson)
	requestparametersjsonstringbyte := []byte(requestparametersjsonstring)
	client := new(http.Client)
	request, _ := http.NewRequest("POST", ECSEndpoint, bytes.NewBuffer(requestparametersjsonstringbyte))
	request = awsauth.SignatureV4(request, requestparametersjsonstringbyte, amztarget, method, params["Region"], service, host, ContentType, signedheaders)
	resp, err := client.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	response["body"] = string(body)
	response["status"] = resp.StatusCode
	return err
}

func(awsmachinelearning *Awsmachinelearning) GetMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var MLModelId, Verbose, Region string

	for key, value := range param {
		switch key {
			case "MLModelId":
				TableNameV, _ := value.(string)
				TableName = TableNameV

			case "Verbose":
				VerboseV, _ := value.(string)
				Verbose = VerboseV

			case "Region":
				RegionV, _ := value.(string)
				Region = RegionV
		}
	}

	params := make(map[string]string)


	preparedeletemodel(params, MLModelId, Region)

	deletemodeljsonmap := map[string]interface{}{
		"MLModelId": MLModelId
		"Verbose" : Verbose,
	}

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, deletemodeljsonmap, response)
	resp = response
	return resp, err
}

func (awsmachinelearning *Awsmachinelearning) UpdateMLModel(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var MLModelId, MLModelName, ScoreThreshold Region string

	for key, value := range param {
		switch key {
			case "MLModelId":
				TableNameV, _ := value.(string)
				TableName = TableNameV

			case "MLModelName":
				MLModelNameV, _ := value.(string)
				MLModelName = MLModelNameV

			case "ScoreThreshold":
				ScoreThresholdV, _ := value.(string)
				ScoreThreshold = ScoreThresholdV

			case "Region":
				RegionV, _ := value.(string)
				Region = RegionV
		}
	}

	params := make(map[string]string)


	prepareupdatemodel(params, MLModelId, Region)

	updatemodeljsonmap := map[string]interface{}{
		"MLModelId": MLModelId ,
		"MLModelName": MLModelName,
		"ScoreThreshold" : ScoreThreshold,
	}

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, updatemodeljsonmap, response)
	resp = response
	return resp, err
}
