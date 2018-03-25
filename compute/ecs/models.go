package ecs

// CreateInstance to store all attribute to create Ali-cloud ECS instance
type CreateInstance struct {
	RegionID                string
	ZoneID                  string
	ImageID                 string
	InstanceType            string
	SecurityGroupID         string
	InstanceName            string
	Description             string
	InternetChargeType      string
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	HostName                string
	Password                string
	IoOptimized             string
	SystemDiskCategory      string
	SystemDiskSize          string
	SystemDiskName          string
	SystemDiskDescription   string
}

// StartInstance to store all attribute to start Ali-cloud ECS instance
type StartInstance struct {
	InstanceID    string
	InitLocalDisk bool
}

// StopInstance to store all attribute to Stop Ali-cloud ECS instance
type StopInstance struct {
	InstanceID string
	ForceStop bool
	ConfirmStop bool
	StoppedMode string
}

// RebootInstance to store all attribute to Reboot Ali-cloud ECS instance
type RebootInstance struct {
	InstanceID string
	ForceStop  bool
}

// DeleteInstance to store all attribute to Delete Ali-cloud ECS instance
type DeleteInstance struct {
	InstanceID string
}
