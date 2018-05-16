package aliloadbalancer

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
