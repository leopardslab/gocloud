package awsmachinelearning

//Awsmachinelearning struct reperesnts aws machine learning service.
type Awsmachinelearning struct {
}

//CreateMLModel struct reperesnts aws machine learning service CreateMLModel.
type CreateMLModel struct {
	MLModelID            string     `json:"MLModelId"`
	MLModelName          string     `json:"MLModelName"`
	MLModelType          string     `json:"MLModelType"`
	parameters           Parameters `json:"Parameters"`
	Recipe               string     `json:"Recipe"`
	RecipeURI            string     `json:"RecipeUri"`
	TrainingDataSourceID string     `json:"TrainingDataSourceId"`
}

//Parameters struct reperesnts CreateMLModel parameters.
type Parameters struct {
	String string `json:"string"`
}
