PACKAGE DOCUMENTATION

package clouddataflow
    import "github.com/cloudlibz/gocloud/streamdataprocessing/clouddataflow


CONSTANTS

const (
    UnixDate = "Mon Jan _2 15:04:05 MST 2006"
    RFC3339  = "2006-01-02T15:04:05Z07:00"
)

TYPES

type AutoscalingSettings struct {
    Algorithm     string `json:"algorithm"`
    MaxNumWorkers int    `json:"maxNumWorkers"`
}
    AutoscalingSettings struct reperesnts Clouddataflow internal params.

type Clouddataflow struct {
}
    Clouddataflow struct reperesnts google streaming service.

func (clouddataflow *Clouddataflow) CreateStream(request interface{}) (resp interface{}, err error)
    CreateStream Create Stream

func (clouddataflow *Clouddataflow) DeleteStream(request interface{}) (resp interface{}, err error)
    DeleteStream Delete Stream

func (clouddataflow *Clouddataflow) DescribeStream(request interface{}) (resp interface{}, err error)
    DescribeStream Describe Stream

func (clouddataflow *Clouddataflow) ListStream(request interface{}) (resp interface{}, err error)
    ListStream ListStream

func (clouddataflow *Clouddataflow) UpdateStream(request interface{}) (resp interface{}, err error)
    DescribeStream Describe Stream

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
    // contains filtered or unexported fields
}
    Clouddataflow struct reperesnts Cloud dataflow service.

type DataDisks struct {
    DiskType   string `json:"diskType"`
    MountPoint string `json:"mountPoint"`
    SizeGb     int    `json:"sizeGb"`
}
    DataDisks struct reperesnts Clouddataflow internal params.

type DisplayData struct {
}
    DisplayData struct reperesnts Clouddataflow internal params.

type Environment struct {
    ClusterManagerAPIService string   `json:"clusterManagerApiService"`
    Dataset                  string   `json:"dataset"`
    Experiments              []string `json:"experiments"`

    ServiceAccountEmail string `json:"serviceAccountEmail"`
    TempStoragePrefix   string `json:"tempStoragePrefix"`
    // contains filtered or unexported fields
}
    Environment struct reperesnts Clouddataflow internal params.

type ExecutionInfo struct {
    // contains filtered or unexported fields
}
    ExecutionInfo struct reperesnts Clouddataflow internal params.

type ExecutionPipelineStage struct {
}
    OriginalPipelineTransform struct reperesnts Clouddataflow internal
    params.

type InternalExperiments struct {
}
    InternalExperiments struct reperesnts Clouddataflow internal params.

type JobMetadata struct {
}
    JobMetadata struct reperesnts Clouddataflow internal params.

type Labels struct {
}
    Labels struct reperesnts Clouddataflow internal params.

type Metadata struct {
}
    Metadata struct reperesnts Clouddataflow internal params.

type OriginalPipelineTransform struct {
}
    OriginalPipelineTransform struct reperesnts Clouddataflow internal
    params.

type Packages struct {
    Location string `json:"location"`
    Name     string `json:"name"`
}
    Packages struct reperesnts Clouddataflow internal params.

type ParallelWorkerSettings struct {
    BaseURL            string `json:"baseUrl"`
    ReportingEnabled   bool   `json:"reportingEnabled"`
    ServicePath        string `json:"servicePath"`
    ShuffleServicePath string `json:"shuffleServicePath"`
    TempStoragePrefix  string `json:"tempStoragePrefix"`
    WorkerID           string `json:"workerId"`
}
    ParallelWorkerSettings struct reperesnts Clouddataflow internal params.

type PipelineDescription struct {
    // contains filtered or unexported fields
}
    PipelineDescription struct reperesnts Clouddataflow internal params.

type PoolArgs struct {
}
    PoolArgs struct reperesnts Clouddataflow internal params.

type Properties struct {
    Kind string `json:"kind"`
    Name string `json:"name"`
}
    Properties struct reperesnts Clouddataflow internal params.

type SdkPipelineOptions struct {
}
    SdkPipelineOptions struct reperesnts Clouddataflow internal params.

type StageStates struct {
    CurrentStateTime    string `json:"currentStateTime"`
    ExecutionStageName  string `json:"executionStageName"`
    ExecutionStageState string `json:"executionStageState"`
}
    StageStates struct reperesnts Clouddataflow internal params.

type Stages struct {
}
    Stages struct reperesnts Clouddataflow internal params.

type Steps struct {
    Kind string `json:"kind"`
    Name string `json:"name"`
    // contains filtered or unexported fields
}
    Steps struct reperesnts Clouddataflow internal params.

type Support struct {
    Status string `json:"status"`
    URL    string `json:"url"`
}
    Support struct reperesnts Clouddataflow internal params.

type TaskrunnerSettings struct {
    Alsologtostderr      bool     `json:"alsologtostderr"`
    BaseTaskDir          string   `json:"baseTaskDir"`
    BaseURL              string   `json:"baseUrl"`
    CommandlinesFileName string   `json:"commandlinesFileName"`
    ContinueOnException  bool     `json:"continueOnException"`
    DataflowAPIVersion   string   `json:"dataflowApiVersion"`
    HarnessCommand       string   `json:"harnessCommand"`
    LanguageHint         string   `json:"languageHint"`
    LogDir               string   `json:"logDir"`
    LogToSerialconsole   bool     `json:"logToSerialconsole"`
    LogUploadLocation    string   `json:"logUploadLocation"`
    OauthScopes          []string `json:"oauthScopes"`

    StreamingWorkerMainClass string `json:"streamingWorkerMainClass"`
    TaskGroup                string `json:"taskGroup"`
    TaskUser                 string `json:"taskUser"`
    TempStoragePrefix        string `json:"tempStoragePrefix"`
    VMID                     string `json:"vmId"`
    WorkflowFileName         string `json:"workflowFileName"`
    // contains filtered or unexported fields
}
    TaskrunnerSettings struct reperesnts Clouddataflow internal params.

type TransformNameMapping struct {
}
    TransformNameMapping struct reperesnts Clouddataflow internal params.

type UserAgent struct {
    Name string `json:"name"`

    BuildDate string `json:"build.date"`
    Version   string `json:"version"`
    // contains filtered or unexported fields
}
    UserAgent struct reperesnts Clouddataflow internal params.

type Version struct {
    Major   string `json:"major"`
    JobType string `json:"job_type"`
}
    Version struct reperesnts Clouddataflow internal params.

type WorkerPools struct {
    DefaultPackageSet string `json:"defaultPackageSet"`
    DiskSizeGb        int    `json:"diskSizeGb"`
    DiskSourceImage   string `json:"diskSourceImage"`
    DiskType          string `json:"diskType"`
    IPConfiguration   string `json:"ipConfiguration"`
    Kind              string `json:"kind"`
    MachineType       string `json:"machineType"`

    Network             string `json:"network"`
    NumThreadsPerWorker int    `json:"numThreadsPerWorker"`
    NumWorkers          int    `json:"numWorkers"`
    OnHostMaintenance   string `json:"onHostMaintenance"`

    Subnetwork string `json:"subnetwork"`

    TeardownPolicy              string `json:"teardownPolicy"`
    WorkerHarnessContainerImage string `json:"workerHarnessContainerImage"`
    Zone                        string `json:"zone"`
    // contains filtered or unexported fields
}
    WorkerPools struct reperesnts Clouddataflow internal params.
