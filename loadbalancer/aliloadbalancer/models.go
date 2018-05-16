package aliloadbalancer

//Aliloadbalancer represents Aliloadbalancer struct.
type Aliloadbalancer struct {
}

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
type DeleteLoadBalancer struct {
	RegionID       string
	LoadBalancerID string
}

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
