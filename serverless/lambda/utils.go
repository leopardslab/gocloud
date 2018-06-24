package lambda

func preparedeleteserviceparams(params map[string]interface{}, deletefunction DeleteFunction, Region string) {
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
