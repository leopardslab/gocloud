package ecs

//CreateInstance to store all attribute to create Ali-cloud ECS instance
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

//StartInstance to store all attribute to start Ali-cloud ECS instance
type StartInstance struct {
	InstanceID    string
	InitLocalDisk bool
}
