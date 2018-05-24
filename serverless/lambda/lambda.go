package lambda

func (lambda *Lambda) Getfunction(request interface{}) (resp interface{}, err error) {

	return resp, err
}

func (lambda *Lambda) Createfunction(request interface{}) (resp interface{}, err error) {

	return resp, err
}

func (lambda *Lambda) Callfunction(request interface{}) (resp interface{}, err error) {

	return resp, err
}

func (lambda *Lambda) Listfunction(request interface{}) (resp interface{}, err error) {

	return resp, err
}

func (lambda *Lambda) Deletefunction(request interface{}) (resp interface{}, err error) {

	return resp, err
}

/*
func (lambda *Lambda) Deletefunction(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var Region string
	var deletefunction Deletefunction

	for key, value := range param {
		switch key {
		case "FunctionName":
			FunctionNameV, _ := value.(string)
			deletefunction.FunctionName = FunctionNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "Qualifier":
			QualifierV, _ := value.(string)
			deletefunction.Qualifier = QualifierV
		}
	}
	params := make(map[string]interface{})
	preparedeletefunctionparams(params, deletefunction, Region)
	response := make(map[string]interface{})
	err = lambda.PrepareSignatureV4query(params,response)
	resp = response
	return resp, err
}


func (lambda *Lambda) PrepareSignatureV4query(params map[string]interface{},response map[string]interface{}) error {
	ECSEndpoint := "https://lambda." + params["Region"] + ".amazonaws.com"
	service := "lambda"
	method := "POST"
	host := service + "." + params["Region"] + ".amazonaws.com"
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
*/
