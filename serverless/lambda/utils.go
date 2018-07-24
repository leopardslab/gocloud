package lambda

func preparedeleteservicequeryString(params map[string]string, deletefunction Deletefunction) {

	if deletefunction.FunctionName != "" {
		params["FunctionName"] = deletefunction.FunctionName
	}

	if deletefunction.Qualifier != "" {
		params["Qualifier"] = deletefunction.Qualifier
	}

}

func preparegetfunctionqueryString(params map[string]string, getfunction Getfunction) {

	if getfunction.FunctionName != "" {
		params["FunctionName"] = getfunction.FunctionName
	}

	if getfunction.Qualifier != "" {
		params["Qualifier"] = getfunction.Qualifier
	}
}

func preparelistfunctionqueryString(params map[string]string, listfunction Listfunction) {

	if listfunction.functionVersion != "" {
		params["FunctionVersion"] = listfunction.functionVersion
	}

	if listfunction.marker != "" {
		params["Marker"] = listfunction.marker
	}

	if listfunction.masterRegion != "" {
		params["MasterRegion"] = listfunction.masterRegion
	}

	if listfunction.maxItems != "" {
		params["MaxItems"] = listfunction.maxItems
	}
}
