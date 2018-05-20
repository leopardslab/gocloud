package aliloadbalancer

import "github.com/cloudlibz/gocloud/aliauth"

//Aliloadbalancer represents Aliloadbalancer struct.
type Aliloadbalancer struct {
}

// CreateLoadBalancer struct represents attribute of create LoadBalancer.
type CreateLoadBalancer struct {
	RegionID           string
	MasterZoneID       string
	SlaveZoneID        string
	LoadBalancerSpec   string
	LoadBalancerName   string
	AddressType        string
	VSwitchID          string
	PayType            string
	PricingCycle       string
	Duration           string
	Autopay            bool
	InternetChargeType string
	Bandwidth          int
	ClientToken        string
	ResourceGroupID    string
}

// DeleteLoadBalancer struct represents attribute of delete LoadBalancer.
type DeleteLoadBalancer struct {
	RegionID       string
	LoadBalancerID string
}

// ListLoadBalancer struct represents attribute of list LoadBalancer.
type ListLoadBalancer struct {
	RegionID              string
	LoadBalancerID        string
	LoadBalancerName      string
	AddressType           string
	NetworkType           string
	VpcID                 string
	VswitchID             string
	Address               string
	ServerIntranetAddress int
	InternetChargeType    string
	ServerID              string
	MasterZoneID          string
	SlaveZoneID           string
}

// AttachLoadBalancer represents Attach node with loadbalancer attribute.
type AttachLoadBalancer struct {
	LoadBalancerID string
	BackendServers string
}

// DetachLoadBalancer represents Detach node with loadbalancer attribute.
type DetachLoadBalancer struct {
	RegionID       string
	LoadBalancerID string
	BackendServers string
}

func init() {
	SetEndpoint(DefaultRegion)
}

const (
	DefaultRegion = "slb.aliyuncs.com"
	Zhangjiakou   = "slb.cn-zhangjiakou.aliyuncs.com"
	Hohhot        = "slb.cn-huhehaote.aliyuncs.com"
	Tokyo         = "slb.ap-northeast-1.aliyuncs.com"
	Sydney        = "slb.ap-southeast-2.aliyuncs.com"
	KualaLumpur   = "slb.ap-southeast-3.aliyuncs.com"
	Jakarta       = "slb.ap-southeast-5.aliyuncs.com"
	Mumbai        = "slb.ap-south-1.aliyuncs.com"
	Dubai         = "slb.me-east-1.aliyuncs.com"
	Frankfurt     = "slb.eu-central-1.aliyuncs.com"
)

func SetEndpoint(region string) {
	aliauth.LoadBalancerRegion = region
}
