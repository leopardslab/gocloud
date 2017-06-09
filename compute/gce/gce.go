package gce


type GCE struct {

	Name string `json:"name"`

	Zone string `json:"zone"`

	MachineType string `json:"machineType"`

	Metadata struct {
		Items []interface{} `json:"items"`
	} `json:"metadata"`

	Tags struct {Items []interface{} `json:"items"`} `json:"tags"`

	Disks []struct {
		Type string `json:"type"`
		Boot bool `json:"boot"`
		Mode string `json:"mode"`
		AutoDelete bool `json:"autoDelete"`
		DeviceName string `json:"deviceName"`

		InitializeParams struct {
			SourceImage string `json:"sourceImage"`
			DiskType string `json:"diskType"`
			DiskSizeGb string `json:"diskSizeGb"`
		} `json:"initializeParams"`

	} `json:"disks"`

	CanIPForward bool `json:"canIpForward"`

	NetworkInterfaces []struct {
		Network string `json:"network"`

		Subnetwork string `json:"subnetwork"`

		AccessConfigs []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"accessConfigs"`
	} `json:"networkInterfaces"`

	Description string `json:"description"`

	Labels struct {
	} `json:"labels"`

	Scheduling struct {
		Preemptible bool `json:"preemptible"`
		OnHostMaintenance string `json:"onHostMaintenance"`
		AutomaticRestart bool `json:"automaticRestart"`
	} `json:"scheduling"`
}
