package clouddataflow

type Clouddataflow struct {
}

/*
type Createstream struct{
  id string,
 projectId string,
 name string,
 Type string,
 currentState string,
 currentStateTime string,
 requestedState string,
 createTime string,
 replaceJobId string,
 location string,
}
*/

type Createstream struct {
	ID               string   `json:"id"`
	ClientRequestID  string   `json:"clientRequestId"`
	CreateTime       string   `json:"createTime"`
	CurrentState     string   `json:"currentState"`
	CurrentStateTime string   `json:"currentStateTime"`
	Location         string   `json:"location"`
	Name             string   `json:"name"`
	ProjectID        string   `json:"projectId"`
	ReplaceJobID     string   `json:"replaceJobId"`
	ReplacedByJobID  string   `json:"replacedByJobId"`
	RequestedState   string   `json:"requestedState"`
	TempFiles        []string `json:"tempFiles"`
	Type             string   `json:"type"`

	environment Environment `json:"environment"`

	jobMetadata JobMetadata `json:"jobMetadata"`

	labels Labels `json:"labels"`

	pipelineDescription PipelineDescription `json:"pipelineDescription"`

	stageStates []StageStates `json:"stageStates"`

	steps []Steps `json:"steps"`

	transformNameMapping TransformNameMapping `json:"transformNameMapping"`

	executionInfo ExecutionInfo `json:"executionInfo"`
}

type StageStates struct {
	CurrentStateTime    string `json:"currentStateTime"`
	ExecutionStageName  string `json:"executionStageName"`
	ExecutionStageState string `json:"executionStageState"`
}

type Properties struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type ExecutionInfo struct {
	stages Stages `json:"stages"`
}

type Stages struct {
}

type TransformNameMapping struct {
}

type Steps struct {
	Kind       string     `json:"kind"`
	Name       string     `json:"name"`
	properties Properties `json:"properties"`
}

type DisplayData struct {
}

type OriginalPipelineTransform struct {
}

type ExecutionPipelineStage struct {
}

type PipelineDescription struct {
	displayData []DisplayData `json:"displayData"`

	executionPipelineStage []ExecutionPipelineStage `json:"executionPipelineStage"`

	originalPipelineTransform []OriginalPipelineTransform `json:"originalPipelineTransform"`
}

type JobMetadata struct {
}

type Labels struct {
}

type InternalExperiments struct {
}

type SdkPipelineOptions struct {
}

type Environment struct {
	ClusterManagerAPIService string              `json:"clusterManagerApiService"`
	Dataset                  string              `json:"dataset"`
	Experiments              []string            `json:"experiments"`
	internalExperiments      InternalExperiments `json:"internalExperiments"`
	sdkPipelineOptions       SdkPipelineOptions  `json:"sdkPipelineOptions"`
	ServiceAccountEmail      string              `json:"serviceAccountEmail"`
	TempStoragePrefix        string              `json:"tempStoragePrefix"`
	userAgent                UserAgent           `json:"userAgent"`
	version                  Version             `json:"version"`
	workerPools              WorkerPools         `json:"workerPools"`
}

type UserAgent struct {
	Name string `json:"name"`
	support Support `json:"support"`
	BuildDate string `json:"build.date"`
	Version   string `json:"version"`
}

type Support struct {
	Status string `json:"status"`
	URL    string `json:"url"`
}

type Version struct {
	Major   string `json:"major"`
	JobType string `json:"job_type"`
}

type Metadata struct {
}

type TaskrunnerSettings struct {
	Alsologtostderr          bool                   `json:"alsologtostderr"`
	BaseTaskDir              string                 `json:"baseTaskDir"`
	BaseURL                  string                 `json:"baseUrl"`
	CommandlinesFileName     string                 `json:"commandlinesFileName"`
	ContinueOnException      bool                   `json:"continueOnException"`
	DataflowAPIVersion       string                 `json:"dataflowApiVersion"`
	HarnessCommand           string                 `json:"harnessCommand"`
	LanguageHint             string                 `json:"languageHint"`
	LogDir                   string                 `json:"logDir"`
	LogToSerialconsole       bool                   `json:"logToSerialconsole"`
	LogUploadLocation        string                 `json:"logUploadLocation"`
	OauthScopes              []string               `json:"oauthScopes"`
	parallelWorkerSettings   ParallelWorkerSettings `json:"parallelWorkerSettings"`
	StreamingWorkerMainClass string                 `json:"streamingWorkerMainClass"`
	TaskGroup                string                 `json:"taskGroup"`
	TaskUser                 string                 `json:"taskUser"`
	TempStoragePrefix        string                 `json:"tempStoragePrefix"`
	VMID                     string                 `json:"vmId"`
	WorkflowFileName         string                 `json:"workflowFileName"`
}

type ParallelWorkerSettings struct {
	BaseURL            string `json:"baseUrl"`
	ReportingEnabled   bool   `json:"reportingEnabled"`
	ServicePath        string `json:"servicePath"`
	ShuffleServicePath string `json:"shuffleServicePath"`
	TempStoragePrefix  string `json:"tempStoragePrefix"`
	WorkerID           string `json:"workerId"`
}

type PoolArgs struct {
}

type Packages struct {
	Location string `json:"location"`
	Name     string `json:"name"`
}

type DataDisks struct {
	DiskType   string `json:"diskType"`
	MountPoint string `json:"mountPoint"`
	SizeGb     int    `json:"sizeGb"`
}

type AutoscalingSettings struct {
	Algorithm     string `json:"algorithm"`
	MaxNumWorkers int    `json:"maxNumWorkers"`
}

type WorkerPools struct {
	autoscalingSettings         AutoscalingSettings `json:"autoscalingSettings"`
	dataDisks                   []DataDisks         `json:"dataDisks"`
	DefaultPackageSet           string              `json:"defaultPackageSet"`
	DiskSizeGb                  int                 `json:"diskSizeGb"`
	DiskSourceImage             string              `json:"diskSourceImage"`
	DiskType                    string              `json:"diskType"`
	IPConfiguration             string              `json:"ipConfiguration"`
	Kind                        string              `json:"kind"`
	MachineType                 string              `json:"machineType"`
	metadata                    Metadata            `json:"metadata"`
	Network                     string              `json:"network"`
	NumThreadsPerWorker         int                 `json:"numThreadsPerWorker"`
	NumWorkers                  int                 `json:"numWorkers"`
	OnHostMaintenance           string              `json:"onHostMaintenance"`
	packages                    []Packages          `json:"packages"`
	poolArgs                    PoolArgs            `json:"poolArgs"`
	Subnetwork                  string              `json:"subnetwork"`
	tsaskrunnerSettings         TaskrunnerSettings  `json:"taskrunnerSettings"`
	TeardownPolicy              string              `json:"teardownPolicy"`
	WorkerHarnessContainerImage string              `json:"workerHarnessContainerImage"`
	Zone                        string              `json:"zone"`
}
