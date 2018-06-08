package googlemachinelearning

//Googlemachinelearning struct reperesnts google machine learning service.
type Googlemachinelearning struct {
}

//CreateMLModel struct reperesnts CreateMLModel service api.
type CreateMLModel struct {
	name                    string
	description             string
	onlinePredictionLogging bool
	regions                 []string
	defaultVersion          DefaultVersion
}

//DefaultVersion struct reperesnts CreateMLModel sub-struct.
type DefaultVersion struct {
	name           string
	description    string
	isDefault      bool
	deploymentUri  string
	createTime     string
	lastUseTime    string
	runtimeVersion string
	state          string
	framework      string
	pythonVersion  string
	errorMessage   string
	autoScaling    AutoScaling
	manualScaling  ManualScaling
}

//AutoScaling struct reperesnts DefaultVersion sub-struct.
type AutoScaling struct {
	minNodes string
}

//ManualScaling struct reperesnts DefaultVersion sub-struct.
type ManualScaling struct {
	nodes string
}
