package ec2

import (
	"encoding/xml"
)

//RunInstances to store all attribute to create EC2 instance.
type RunInstances struct {
	ImageID               string
	MinCount              int
	MaxCount              int
	KeyName               string
	InstanceType          string
	SecurityGroups        []SecurityGroup
	KernelID              string
	RamdiskID             string
	UserData              []byte
	AvailZone             string
	PlacementGroupName    string
	Monitoring            bool
	SubnetID              string
	DisableAPITermination bool
	ShutdownBehavior      string
	PrivateIPAddress      string
	BlockDeviceMappings   []BlockDeviceMapping
	NetworkInterfaces     []RunNetworkInterface
}

//SecurityGroup struct.
type SecurityGroup struct {
	ID   string `xml:"groupID"`
	Name string `xml:"groupName"`
}

//BlockDeviceMapping struct to attach device
type BlockDeviceMapping struct {
	DeviceName          string `xml:"deviceName"`
	VirtualName         string `xml:"virtualName"`
	SnapshotID          string `xml:"ebs>snapshotID"`
	VolumeType          string `xml:"ebs>volumeType"`
	VolumeSize          int64  `xml:"ebs>volumeSize"`
	DeleteOnTermination bool   `xml:"ebs>deleteOnTermination"`
	IOPS                int64  `xml:"ebs>iops"`
}

//RunNetworkInterface struct for Ec2.
type RunNetworkInterface struct {
	ID                     string
	DeviceIndex             int
	SubnetID               string
	Description             string
	PrivateIPs              []PrivateIP
	SecurityGroupIds        []string
	DeleteOnTermination     bool
	SecondaryPrivateIPCount int
}

//PrivateIP to assign PrivateIP.
type PrivateIP struct {
	Address   string `xml:"privateIpAddress"`
	DNSName   string `xml:"privateDnsName"`
	IsPrimary bool   `xml:"primary"`
}

//RunInstancesResp response.
type RunInstancesResp struct {
	RequestID      string          `xml:"RequestID"`
	ReservationID string          `xml:"ReservationID"`
	OwnerID        string          `xml:"OwnerID"`
	SecurityGroups []SecurityGroup `xml:"groupSet>item"`
	Instances      []Instance      `xml:"instancesSet>item"`
}

//Instance struct represents running instance.
type Instance struct {
	InstanceID         string             `xml:"instanceID"`
	InstanceType       string             `xml:"instanceType"`
	ImageID            string             `xml:"imageId"`
	PrivateDNSName     string             `xml:"privateDnsName"`
	DNSName            string             `xml:"dnsName"`
	IPAddress          string             `xml:"ipAddress"`
	PrivateIPAddress   string             `xml:"privateIpAddress"`
	SubnetID           string             `xml:"subnetId"`
	VPCID              string             `xml:"vpcID"`
	SourceDestCheck    bool               `xml:"sourceDestCheck"`
	KeyName            string             `xml:"keyName"`
	AMILaunchIndex     int                `xml:"amiLaunchIndex"`
	Hypervisor         string             `xml:"hypervisor"`
	VirtType           string             `xml:"virtualizationType"`
	Monitoring         string             `xml:"monitoring>state"`
	AvailZone          string             `xml:"placement>availabilityZone"`
	PlacementGroupName string             `xml:"placement>groupName"`
	State              InstanceState      `xml:"instanceState"`
	Tags               []Tag              `xml:"tagSet>item"`
	SecurityGroups     []SecurityGroup    `xml:"groupSet>item"`
	NetworkInterfaces  []NetworkInterface `xml:"networkInterfaceSet>item"`
}

//InstanceStateChange stuct represents instance state change.
type InstanceStateChange struct {
	InstanceID    string        `xml:"instanceId"`
	CurrentState  InstanceState `xml:"currentState"`
	PreviousState InstanceState `xml:"previousState"`
}

//SimpleResp stuct represents SimpleResp.
type SimpleResp struct {
	XMLName   xml.Name
	RequestID string `xml:"requestId"`
}

//TerminateInstance struct represents TerminateInstance response.
type TerminateInstancesResp struct {
	RequestID    string                `xml:"requestID"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}

//InstanceState struct represents InstanceState.
type InstanceState struct {
	Code int    `xml:"code"`
	Name string `xml:"name"`
}

//NetworkInterface reperents running instance NetworkInterface.
type NetworkInterface struct {
	ID               string                     `xml:"networkInterfaceID"`
	SubnetId         string                     `xml:"subnetId"`
	VPCId            string                     `xml:"vpcId"`
	AvailZone        string                     `xml:"availabilityZone"`
	Description      string                     `xml:"description"`
	OwnerID          string                     `xml:"OwnerID"`
	RequesterID      string                     `xml:"requesterID"`
	RequesterManaged bool                       `xml:"requesterManaged"`
	Status           string                     `xml:"status"`
	MACAddress       string                     `xml:"macAddress"`
	PrivateIPAddress string                     `xml:"privateIpAddress"`
	PrivateDNSName   string                     `xml:"privateDnsName"`
	SourceDestCheck  bool                       `xml:"sourceDestCheck"`
	Groups           []SecurityGroup            `xml:"groupSet>item"`
	Attachment       NetworkInterfaceAttachment `xml:"attachment"`
	Tags             []Tag                      `xml:"tagSet>item"`
	PrivateIPs       []PrivateIP                `xml:"privateIpAddressesSet>item"`
}

//NetworkInterfaceAttachment represents running instance
type NetworkInterfaceAttachment struct {
	ID                  string `xml:"attachmentIDs"`
	InstanceID          string `xml:"instanceID"`
	InstanceOwnerId     string `xml:"instanceOwnerId"`
	DeviceIndex         int    `xml:"deviceIndex"`
	Status              string `xml:"status"`
	AttachTime          string `xml:"attachTime"`
	DeleteOnTermination bool   `xml:"deleteOnTermination"`
}

//Tag reperent tag assgin to instance
type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

//StartInstanceResp response.
type StartInstanceResp struct {
	RequestID    string                `xml:"requestID"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}

//StopInstanceResp response.
type StopInstanceResp struct {
	RequestID    string                `xml:"requestID"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}
