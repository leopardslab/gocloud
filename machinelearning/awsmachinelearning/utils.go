package awsmachinelearning

func preparecreateMLModelparamsdict(createMLModeljsonmap map[string]interface{}, createMLModel CreateMLModel) {

	if createMLModel.MLModelID != "" {
		createMLModeljsonmap["MLModelId"] = createMLModel.MLModelID
	}
	if createMLModel.MLModelName != "" {
		createMLModeljsonmap["MLModelName"] = createMLModel.MLModelName
	}

	if createMLModel.MLModelType != "" {
		createMLModeljsonmap["MLModelType"] = createMLModel.MLModelType
	}
	if createMLModel.Recipe != "" {
		createMLModeljsonmap["Recipe"] = createMLModel.Recipe
	}

	if createMLModel.RecipeURI != "" {
		createMLModeljsonmap["RecipeURI"] = createMLModel.RecipeURI
	}

	if createMLModel.TrainingDataSourceID != "" {
		createMLModeljsonmap["TrainingDataSourceID"] = createMLModel.TrainingDataSourceID
	}

	if createMLModel.parameters.String != "" {
		parameters := make(map[string]string)
		parameters["string"] = createMLModel.parameters.String
		createMLModeljsonmap["parameters"] = parameters
	}

}

func preparecreateMLModel(params map[string]string, createMLModel CreateMLModel, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonML_20141212.CreateMLModel"
}



func preparedeletemodel(params map[string]string, MLModelId string, Region string) {

	if MLModelId != "" {
		params["MLModelId"] = MLModelId
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "AmazonML_20141212.DeleteMLModel"
}
func preparegetmodel(params map[string]string, MLModelId string, Verbose string, Region string) {

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

func prepareupdatemodel(params map[string]string, MLModelId string, ScoreThreshold string, MLModelName string, Region string) {

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
