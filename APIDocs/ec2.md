package ec2
    import "github.com/scorelab/gocloud-v2/compute/ec2"


VARIABLES

var USEast = Region{
    "us-east-1",
    "https://ec2.us-east-1.amazonaws.com",
}

TYPES

type Auth struct {
    AccessKey string
    SecretKey string
}
    Authentication struct to store AccessKey and SecretKey

type BlockDeviceMapping struct {
    DeviceName          string `xml:"deviceName"`
    VirtualName         string `xml:"virtualName"`
    SnapshotId          string `xml:"ebs>snapshotId"`
    VolumeType          string `xml:"ebs>volumeType"`
    VolumeSize          int64  `xml:"ebs>volumeSize"`
    DeleteOnTermination bool   `xml:"ebs>deleteOnTermination"`
    IOPS                int64  `xml:"ebs>iops"`
}

type EC2 struct {
    Auth
    Region
    Private byte
}

func New2(auth Auth, region Region) *EC2
    Function return EC2 instance

func (ec2 *EC2) Createnode(request interface{}) (resp interface{}, err error)

func (ec2 *EC2) Deletenode(request interface{}) (resp interface{}, err error)

func (ec2 *EC2) PrepareSignatureV2query(params map[string]string, Region string, response map[string]interface{}) error

func (ec2 *EC2) Rebootnode(request interface{}) (resp interface{}, err error)

func (ec2 *EC2) Startnode(request interface{}) (resp interface{}, err error)

func (ec2 *EC2) Stopnode(request interface{}) (resp interface{}, err error)

type Error struct {
    StatusCode int

    Code string

    Message string

    RequestId string `xml:"RequestID"`
}
    represents error

func (err *Error) Error() string

type Instance struct {
    InstanceId         string             `xml:"instanceId"`
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

type InstanceState struct {
    Code int    `xml:"code"`
    Name string `xml:"name"`
}

type InstanceStateChange struct {
    InstanceId    string        `xml:"instanceId"`
    CurrentState  InstanceState `xml:"currentState"`
    PreviousState InstanceState `xml:"previousState"`
}

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
    InstanceId          string `xml:"instanceId"`
    InstanceOwnerId     string `xml:"instanceOwnerId"`
    DeviceIndex         int    `xml:"deviceIndex"`
    Status              string `xml:"status"`
    AttachTime          string `xml:"attachTime"`
    DeleteOnTermination bool   `xml:"deleteOnTermination"`
}

type PrivateIP struct {
    Address   string `xml:"privateIpAddress"`
    DNSName   string `xml:"privateDnsName"`
    IsPrimary bool   `xml:"primary"`
}

type Region struct {
    Name        string
    EC2Endpoint string
}
    repersents Region


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

type RunInstancesResp struct {
    RequestId      string          `xml:"requestId"`
    ReservationId  string          `xml:"reservationId"`
    OwnerId        string          `xml:"ownerId"`
    SecurityGroups []SecurityGroup `xml:"groupSet>item"`
    Instances      []Instance      `xml:"instancesSet>item"`
}

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

type SecurityGroup struct {
    Id   string `xml:"groupId"`
    Name string `xml:"groupName"`
}

type SimpleResp struct {
    XMLName   xml.Name
    RequestId string `xml:"requestId"`
}

type StartInstanceResp struct {
    RequestId    string                `xml:"requestId"`
    StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}

type StopInstanceResp struct {
    RequestId    string                `xml:"requestId"`
    StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}

type Tag struct {
    Key   string `xml:"key"`
    Value string `xml:"value"`
}
    reperent tag assgin to instance

type TerminateInstancesResp struct {
    RequestId    string                `xml:"requestId"`
    StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}


