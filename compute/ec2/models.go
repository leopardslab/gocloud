package ec2

import (
	"encoding/xml"
	"time"
)

//runinstance to store all attribute to create EC2 instance

type RunInstances struct {
	ImageId               string
	MinCount              int
	MaxCount              int
	KeyName               string
	InstanceType          string
	SecurityGroups        []SecurityGroup
	KernelId              string
	RamdiskId             string
	UserData              []byte
	AvailZone             string
	PlacementGroupName    string
	Monitoring            bool
	SubnetId              string
	DisableAPITermination bool
	ShutdownBehavior      string
	PrivateIPAddress      string
	BlockDeviceMappings   []BlockDeviceMapping
	NetworkInterfaces     []RunNetworkInterface
}

//SecurityGroup struct

type SecurityGroup struct {
	Id   string `xml:"groupId"`
	Name string `xml:"groupName"`
}

// BlockDevice struct to attach device

type BlockDeviceMapping struct {
	DeviceName          string `xml:"deviceName"`
	VirtualName         string `xml:"virtualName"`
	SnapshotId          string `xml:"ebs>snapshotId"`
	VolumeType          string `xml:"ebs>volumeType"`
	VolumeSize          int64  `xml:"ebs>volumeSize"`
	DeleteOnTermination bool   `xml:"ebs>deleteOnTermination"`
	IOPS                int64  `xml:"ebs>iops"`
}

//NetworkInterface struct for Ec2

type RunNetworkInterface struct {
	Id                      string
	DeviceIndex             int
	SubnetId                string
	Description             string
	PrivateIPs              []PrivateIP
	SecurityGroupIds        []string
	DeleteOnTermination     bool
	SecondaryPrivateIPCount int
}

//Private ip to assign PrivateIP

type PrivateIP struct {
	Address   string `xml:"privateIpAddress"`
	DNSName   string `xml:"privateDnsName"`
	IsPrimary bool   `xml:"primary"`
}

// run instance response

type RunInstancesResp struct {
	RequestID      string          `xml:"requestId"`
	ReservationId  string          `xml:"reservationId"`
	OwnerId        string          `xml:"ownerId"`
	SecurityGroups []SecurityGroup `xml:"groupSet>item"`
	Instances      []Instance      `xml:"instancesSet>item"`
}

// this struct represents running instance

type Instance struct {
	InstanceID         string             `xml:"InstanceID"`
	InstanceType       string             `xml:"instanceType"`
	ImageId            string             `xml:"imageId"`
	PrivateDNSName     string             `xml:"privateDnsName"`
	DNSName            string             `xml:"dnsName"`
	IPAddress          string             `xml:"ipAddress"`
	PrivateIPAddress   string             `xml:"privateIpAddress"`
	SubnetId           string             `xml:"subnetId"`
	VPCId              string             `xml:"vpcId"`
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

//This stuct represents instance state change

type InstanceStateChange struct {
	InstanceID    string        `xml:"InstanceID"`
	CurrentState  InstanceState `xml:"currentState"`
	PreviousState InstanceState `xml:"previousState"`
}

type SimpleResp struct {
	XMLName   xml.Name
	RequestID string `xml:"requestId"`
}

//struct to TerminateInstance

type TerminateInstancesResp struct {
	RequestID    string                `xml:"requestId"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}

// InstanceState struct

type InstanceState struct {
	Code int    `xml:"code"`
	Name string `xml:"name"`
}

//reperent ruuing instance NetworkInterface

type NetworkInterface struct {
	Id               string                     `xml:"networkInterfaceId"`
	SubnetId         string                     `xml:"subnetId"`
	VPCId            string                     `xml:"vpcId"`
	AvailZone        string                     `xml:"availabilityZone"`
	Description      string                     `xml:"description"`
	OwnerId          string                     `xml:"ownerId"`
	RequesterId      string                     `xml:"requesterId"`
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

type NetworkInterfaceAttachment struct {
	Id                  string `xml:"attachmentId"`
	InstanceID          string `xml:"InstanceID"`
	InstanceOwnerId     string `xml:"instanceOwnerId"`
	DeviceIndex         int    `xml:"deviceIndex"`
	Status              string `xml:"status"`
	AttachTime          string `xml:"attachTime"`
	DeleteOnTermination bool   `xml:"deleteOnTermination"`
}

// reperent tag assgin to instance
type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

//start instance response

type StartInstanceResp struct {
	RequestID    string                `xml:"requestId"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}

//stop instances response

type StopInstanceResp struct {
	RequestID    string                `xml:"requestId"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}

type Responsestruct struct {
	RunInstancesResponse struct {
		Xmlns         string `json:"-xmlns"`
		RequestID     string `json:"requestId"`
		ReservationID string `json:"reservationId"`
		OwnerID       string `json:"ownerId"`
		InstancesSet  struct {
			Item struct {
				InstanceID    string `json:"InstanceID"`
				ImageID       string `json:"imageId"`
				InstanceState struct {
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"instanceState"`
				PrivateDNSName string    `json:"privateDnsName"`
				AmiLaunchIndex string    `json:"amiLaunchIndex"`
				InstanceType   string    `json:"instanceType"`
				LaunchTime     time.Time `json:"launchTime"`
				Placement      struct {
					AvailabilityZone string `json:"availabilityZone"`
					Tenancy          string `json:"tenancy"`
				} `json:"placement"`
				KernelID   string `json:"kernelId"`
				Monitoring struct {
					State string `json:"state"`
				} `json:"monitoring"`
				SubnetID         string `json:"subnetId"`
				VpcID            string `json:"vpcId"`
				PrivateIPAddress string `json:"privateIpAddress"`
				SourceDestCheck  string `json:"sourceDestCheck"`
				GroupSet         struct {
					Item struct {
						GroupID   string `json:"groupId"`
						GroupName string `json:"groupName"`
					} `json:"item"`
				} `json:"groupSet"`
				StateReason struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				} `json:"stateReason"`
				Architecture        string `json:"architecture"`
				RootDeviceType      string `json:"rootDeviceType"`
				RootDeviceName      string `json:"rootDeviceName"`
				VirtualizationType  string `json:"virtualizationType"`
				ClientToken         string `json:"clientToken"`
				Hypervisor          string `json:"hypervisor"`
				NetworkInterfaceSet struct {
					Item struct {
						NetworkInterfaceID string `json:"networkInterfaceId"`
						SubnetID           string `json:"subnetId"`
						VpcID              string `json:"vpcId"`
						OwnerID            string `json:"ownerId"`
						Status             string `json:"status"`
						PrivateIPAddress   string `json:"privateIpAddress"`
						PrivateDNSName     string `json:"privateDnsName"`
						SourceDestCheck    string `json:"sourceDestCheck"`
						GroupSet           struct {
							Item struct {
								GroupID   string `json:"groupId"`
								GroupName string `json:"groupName"`
							} `json:"item"`
						} `json:"groupSet"`
						Attachment struct {
							AttachmentID        string    `json:"attachmentId"`
							DeviceIndex         string    `json:"deviceIndex"`
							Status              string    `json:"status"`
							AttachTime          time.Time `json:"attachTime"`
							DeleteOnTermination string    `json:"deleteOnTermination"`
						} `json:"attachment"`
					} `json:"item"`
				} `json:"networkInterfaceSet"`
			} `json:"item"`
		} `json:"instancesSet"`
	} `json:"RunInstancesResponse"`
}
