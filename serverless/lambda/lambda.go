package lambda

//GetFunction  describe lambda function.
func (lambda *Lambda) GetFunction(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var Region string
	var getfunction Getfunction

	for key, value := range param {
		switch key {
			case "FunctionName":
				FunctionNameV, _ := value.(string)
				getfunction.FunctionName = FunctionNameV

			case "Region":
				RegionV, _ := value.(string)
				Region = RegionV

			case "Qualifier":
				QualifierV, _ := value.(string)
				getfunction.Qualifier = QualifierV
			}
	}

	params := make(map[string]string)
	preparegetfunctionqueryString(params, getfunction)

	response := make(map[string]interface{})
	err = Preparerequest(params,Region,response)

	resp = response
	return resp, err
}

//CreateFunction  Create lambda function.
func (lambda *Lambda) CreateFunction(request interface{}) (resp interface{}, err error) {

	return resp, err
}

//CallFunction  call lambda function.
func (lambda *Lambda) CallFunction(request interface{}) (resp interface{}, err error) {

	return resp, err
}



//ListFunction  list lambda function.
func (lambda *Lambda) ListFunction(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var Region string
	var listfunction Listfunction

	for key, value := range param {
		switch key {

		case "FunctionVersion":
			functionVersionv, _ := value.(string)
			listfunction.functionVersion = functionVersionv

		case "Marker":
			markerv, _ := value.(string)
			listfunction.marker = markerv

		case "MasterRegion":
			masterRegionv, _ := value.(string)
			listfunction.masterRegion = masterRegionv


		case "MaxItems":
			maxItemsv, _ := value.(string)
			listfunction.maxItems = maxItemsv


		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)
	preparelistfunctionqueryString(params, listfunction)

	response := make(map[string]interface{})
		err = Preparerequest(params,Region,response)

	resp = response
	return resp, err

}

//DeleteFunction  delete lambda function.
func (lambda *Lambda) DeleteFunction(request interface{}) (resp interface{}, err error) {
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

	params := make(map[string]string)
	preparedeleteservicequeryString(params, deletefunction)

	response := make(map[string]interface{})
	err = Preparerequest(params,Region,response)

	resp = response
	return resp, err
}
