package gce

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

type Disk struct {
	Type             string          `json:"type"`
	Boot             bool            `json:"boot"`
	Mode             string          `json:"mode"`
	AutoDelete       bool            `json:"autoDelete"`
	DeviceName       string          `json:"deviceName"`
	InitializeParams InitializeParam `json:"initializeParams"`
}

type InitializeParam struct {
	SourceImage string `json:"sourceImage"`
	DiskType    string `json:"diskType"`
	DiskSizeGb  string `json:"diskSizeGb"`
}

type NetworkInterface struct {
	Network       string         `json:"network"`
	Subnetwork    string         `json:"subnetwork"`
	AccessConfigs []accessConfig `json:"accessConfigs"`
}

type accessConfig struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Scheduling struct {
	Preemptible       bool   `json:"preemptible"`
	OnHostMaintenance string `json:"onHostMaintenance"`
	AutomaticRestart  bool   `json:"automaticRestart"`
}
