package googlemachinelearning

//Googlemachinelearning struct reperesnts google machine learning service.
type Googlemachinelearning struct {
}

type CreateMLModel struct{
  name string
  description string
  onlinePredictionLogging string
  regions []string
  defaultVersion DefaultVersion
}

type DefaultVersion struct{
  name string
  description string
  isDefault bool
  deploymentUri string
  createTime string
  lastUseTime string
  runtimeVersion string
  state string
  framework string
  pythonVersion string
  errorMessage string
  autoScaling AutoScaling
  manualScaling ManualScaling
}

type AutoScaling struct{
  minNodes string
}

type ManualScaling struct{
  nodes string
}
