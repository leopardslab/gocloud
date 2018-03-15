package aliecs


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