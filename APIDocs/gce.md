package gce
    import "github.com/scorelab/gocloud-v2/compute/gce"


TYPES

type Disk struct {
    Type             string          `json:"type"`
    Boot             bool            `json:"boot"`
    Mode             string          `json:"mode"`
    AutoDelete       bool            `json:"autoDelete"`
    DeviceName       string          `json:"deviceName"`
    InitializeParams InitializeParam `json:"initializeParams"`
}
    Disk represents GCE disk.

type GCE struct {
    Name string `json:"name"`

    Zone string `json:"zone"`

    MachineType string `json:"machineType"`

    Disks []Disk `json:"disks"`

    CanIPForward bool `json:"canIpForward"`

    NetworkInterfaces []NetworkInterface `json:"networkInterfaces"`

    Description string `json:"description"`

    Scheduling `json:"scheduling"`
    // contains filtered or unexported fields
}
    GCE struct reperesnts Google Compute Engine.

func (gce *GCE) Createnode(request interface{}) (resp interface{}, err error)

func (gce *GCE) Deletenode(request interface{}) (resp interface{}, err error)

func (gce *GCE) Rebootnode(request interface{}) (resp interface{}, err error)

func (gce *GCE) Startnode(request interface{}) (resp interface{}, err error)

func (gce *GCE) Stopnode(request interface{}) (resp interface{}, err error)

type GCEResponse struct {
    Kind          string `json:"kind"`
    ID            string `json:"id"`
    Name          string `json:"name"`
    Zone          string `json:"zone"`
    OperationType string `json:"operationType"`
    TargetLink    string `json:"targetLink"`
    TargetID      string `json:"targetId"`
    Status        string `json:"status"`
    User          string `json:"user"`
    Progress      int    `json:"progress"`
    InsertTime    string `json:"insertTime"`
    StartTime     string `json:"startTime"`
    EndTime       string `json:"endTime"`
    SelfLink      string `json:"selfLink"`
}

type InitializeParam struct {
    SourceImage string `json:"sourceImage"`
    DiskType    string `json:"diskType"`
    DiskSizeGb  string `json:"diskSizeGb"`
}
    InitializeParam represents GCE disk InitializeParam.

type NetworkInterface struct {
    Network       string         `json:"network"`
    Subnetwork    string         `json:"subnetwork"`
    AccessConfigs []accessConfig `json:"accessConfigs"`
}
    NetworkInterface represents GCE NetworkInterface.

type Scheduling struct {
    Preemptible       bool   `json:"preemptible"`
    OnHostMaintenance string `json:"onHostMaintenance"`
    AutomaticRestart  bool   `json:"automaticRestart"`
}
    Scheduling represents GCE instance Scheduling.


