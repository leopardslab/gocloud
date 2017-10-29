package gce

//GCE struct reperesnts Google Compute Engine.
type GCE struct {
	Name string `json:"name"`
	Zone string `json:"zone"`
	MachineType string `json:"machineType"`
	Disks []Disk `json:"disks"`
	CanIPForward bool `json:"canIpForward"`
	NetworkInterfaces []NetworkInterface `json:"networkInterfaces"`
	Description string `json:"description"`
	selfLink string `json:"selfLink"`
	Scheduling `json:"scheduling"`
}

//Disk represents GCE disk.
type Disk struct {
	Type             string          `json:"type"`
	Boot             bool            `json:"boot"`
	Mode             string          `json:"mode"`
	AutoDelete       bool            `json:"autoDelete"`
	DeviceName       string          `json:"deviceName"`
	InitializeParams InitializeParam `json:"initializeParams"`
}

//InitializeParam represents GCE disk InitializeParam.
type InitializeParam struct {
	SourceImage string `json:"sourceImage"`
	DiskType    string `json:"diskType"`
	DiskSizeGb  string `json:"diskSizeGb"`
}

//NetworkInterface represents GCE NetworkInterface.
type NetworkInterface struct {
	Network       string         `json:"network"`
	Subnetwork    string         `json:"subnetwork"`
	AccessConfigs []accessConfig `json:"accessConfigs"`
}

//accessConfig represents GCE NetworkInterface accessConfig.
type accessConfig struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

//Scheduling represents GCE instance Scheduling.
type Scheduling struct {
	Preemptible       bool   `json:"preemptible"`
	OnHostMaintenance string `json:"onHostMaintenance"`
	AutomaticRestart  bool   `json:"automaticRestart"`
}

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
