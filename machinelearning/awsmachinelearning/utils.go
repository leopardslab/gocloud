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
		createMLModeljsonmap["RecipeUri"] = createMLModel.RecipeURI
	}

	if createMLModel.TrainingDataSourceID != "" {
		createMLModeljsonmap["TrainingDataSourceId"] = createMLModel.TrainingDataSourceID
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

func prepareupdatemodelparamsdict(updatemodeljsonmap map[string]interface{}, MLModelId string, ScoreThreshold int, MLModelName string) {

	if MLModelId != "" {
		updatemodeljsonmap["MLModelId"] = MLModelId
	}

	if MLModelName != "" {
		updatemodeljsonmap["MLModelName"] = MLModelName
	}

	if ScoreThreshold != 0 {
		updatemodeljsonmap["ScoreThreshold"] = ScoreThreshold
	}
}

func prepareupdatemodel(params map[string]string, Region string) {

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "AmazonML_20141212.UpdateMLModel"
}
