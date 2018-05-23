package awsmachinelearning


type Awsmachinelearning struct{

}

type CreateMLModel struct {
	MLModelID   string `json:"MLModelId"`
	MLModelName string `json:"MLModelName"`
	MLModelType string `json:"MLModelType"`
	parameters Parameters`json:"Parameters"`
	Recipe               string `json:"Recipe"`
	RecipeURI            string `json:"RecipeUri"`
	TrainingDataSourceID string `json:"TrainingDataSourceId"`
}


type Parameters  struct {
  String string `json:"string"`
}
