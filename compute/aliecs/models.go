package aliecs


type CreateInstance struct {
	RegionId                    string
	ZoneId                      string
	ImageId                     string
	InstanceType                string
	SecurityGroupId             string
	InstanceName                string
	Description                 string
	InternetChargeType          string
	InternetMaxBandwidthIn      int
	InternetMaxBandwidthOut     int
	HostName                    string
	Password                    string
	IoOptimized                 string
	SystemDiskCategory          string
	SystemDiskSize              string
	SystemDiskName              string
	SystemDiskDescription       string
}