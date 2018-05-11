package lambda


func preparedeleteserviceparams(params map[string]interface{}, deletefunction Deletefunction, Region string) {
	if Region != "" {
		params["Region"] = Region
	}

	if deletefunction.FunctionName != "" {
		params["FunctionName"] = deletefunction.FunctionName
	}

	if deletefunction.Qualifier != "" {
		params["Qualifier"] = deletefunction.Qualifier
	}

}
