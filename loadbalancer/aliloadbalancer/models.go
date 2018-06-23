package aliloadbalancer

import (
	"errors"
	"github.com/cloudlibz/gocloud/aliauth"
)

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

// CreateLoadBalancer builder pattern code
type CreateLoadBalancerBuilder struct {
	createLoadBalancer *CreateLoadBalancer
}

func NewCreateLoadBalancerBuilder() *CreateLoadBalancerBuilder {
	createLoadBalancer := &CreateLoadBalancer{}
	b := &CreateLoadBalancerBuilder{createLoadBalancer: createLoadBalancer}
	return b
}

func (b *CreateLoadBalancerBuilder) RegionID(regionID string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.RegionID = regionID
	return b
}

func (b *CreateLoadBalancerBuilder) MasterZoneID(masterZoneID string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.MasterZoneID = masterZoneID
	return b
}

func (b *CreateLoadBalancerBuilder) SlaveZoneID(slaveZoneID string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.SlaveZoneID = slaveZoneID
	return b
}

func (b *CreateLoadBalancerBuilder) LoadBalancerSpec(loadBalancerSpec string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.LoadBalancerSpec = loadBalancerSpec
	return b
}

func (b *CreateLoadBalancerBuilder) LoadBalancerName(loadBalancerName string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.LoadBalancerName = loadBalancerName
	return b
}

func (b *CreateLoadBalancerBuilder) AddressType(addressType string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.AddressType = addressType
	return b
}

func (b *CreateLoadBalancerBuilder) VSwitchID(vSwitchID string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.VSwitchID = vSwitchID
	return b
}

func (b *CreateLoadBalancerBuilder) PayType(payType string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.PayType = payType
	return b
}

func (b *CreateLoadBalancerBuilder) PricingCycle(pricingCycle string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.PricingCycle = pricingCycle
	return b
}

func (b *CreateLoadBalancerBuilder) Duration(duration string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.Duration = duration
	return b
}

func (b *CreateLoadBalancerBuilder) Autopay(autopay bool) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.Autopay = autopay
	return b
}

func (b *CreateLoadBalancerBuilder) InternetChargeType(internetChargeType string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.InternetChargeType = internetChargeType
	return b
}

func (b *CreateLoadBalancerBuilder) Bandwidth(bandwidth int) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.Bandwidth = bandwidth
	return b
}

func (b *CreateLoadBalancerBuilder) ClientToken(clientToken string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.ClientToken = clientToken
	return b
}

func (b *CreateLoadBalancerBuilder) ResourceGroupID(resourceGroupID string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.ResourceGroupID = resourceGroupID
	return b
}

func (b *CreateLoadBalancerBuilder) Build() (map[string]interface{}, error) {
	if b.createLoadBalancer.RegionID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RegionID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.createLoadBalancer)
	return params, nil
}

// DeleteLoadBalancer builder pattern code
type DeleteLoadBalancerBuilder struct {
	deleteLoadBalancer *DeleteLoadBalancer
}

func NewDeleteLoadBalancerBuilder() *DeleteLoadBalancerBuilder {
	deleteLoadBalancer := &DeleteLoadBalancer{}
	b := &DeleteLoadBalancerBuilder{deleteLoadBalancer: deleteLoadBalancer}
	return b
}

func (b *DeleteLoadBalancerBuilder) RegionID(regionID string) *DeleteLoadBalancerBuilder {
	b.deleteLoadBalancer.RegionID = regionID
	return b
}

func (b *DeleteLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *DeleteLoadBalancerBuilder {
	b.deleteLoadBalancer.LoadBalancerID = loadBalancerID
	return b
}

func (b *DeleteLoadBalancerBuilder) Build() (map[string]interface{}, error) {
	if b.deleteLoadBalancer.RegionID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RegionID")
	}
	if b.deleteLoadBalancer.LoadBalancerID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "LoadBalancerID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.deleteLoadBalancer)
	return params, nil
}

// ListLoadBalancer builder pattern code
type ListLoadBalancerBuilder struct {
	listLoadBalancer *ListLoadBalancer
}

func NewListLoadBalancerBuilder() *ListLoadBalancerBuilder {
	listLoadBalancer := &ListLoadBalancer{}
	b := &ListLoadBalancerBuilder{listLoadBalancer: listLoadBalancer}
	return b
}

func (b *ListLoadBalancerBuilder) RegionID(regionID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.RegionID = regionID
	return b
}

func (b *ListLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.LoadBalancerID = loadBalancerID
	return b
}

func (b *ListLoadBalancerBuilder) LoadBalancerName(loadBalancerName string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.LoadBalancerName = loadBalancerName
	return b
}

func (b *ListLoadBalancerBuilder) AddressType(addressType string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.AddressType = addressType
	return b
}

func (b *ListLoadBalancerBuilder) NetworkType(networkType string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.NetworkType = networkType
	return b
}

func (b *ListLoadBalancerBuilder) VpcID(vpcID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.VpcID = vpcID
	return b
}

func (b *ListLoadBalancerBuilder) VswitchID(vswitchID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.VswitchID = vswitchID
	return b
}

func (b *ListLoadBalancerBuilder) Address(address string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.Address = address
	return b
}

func (b *ListLoadBalancerBuilder) ServerIntranetAddress(serverIntranetAddress int) *ListLoadBalancerBuilder {
	b.listLoadBalancer.ServerIntranetAddress = serverIntranetAddress
	return b
}

func (b *ListLoadBalancerBuilder) InternetChargeType(internetChargeType string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.InternetChargeType = internetChargeType
	return b
}

func (b *ListLoadBalancerBuilder) ServerID(serverID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.ServerID = serverID
	return b
}

func (b *ListLoadBalancerBuilder) MasterZoneID(masterZoneID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.MasterZoneID = masterZoneID
	return b
}

func (b *ListLoadBalancerBuilder) SlaveZoneID(slaveZoneID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.SlaveZoneID = slaveZoneID
	return b
}

func (b *ListLoadBalancerBuilder) Build() (map[string]interface{}, error) {
	if b.listLoadBalancer.RegionID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RegionID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.listLoadBalancer)
	return params, nil
}

// AttachLoadBalancer builder pattern code
type AttachLoadBalancerBuilder struct {
	attachLoadBalancer *AttachLoadBalancer
}

func NewAttachLoadBalancerBuilder() *AttachLoadBalancerBuilder {
	attachLoadBalancer := &AttachLoadBalancer{}
	b := &AttachLoadBalancerBuilder{attachLoadBalancer: attachLoadBalancer}
	return b
}

func (b *AttachLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *AttachLoadBalancerBuilder {
	b.attachLoadBalancer.LoadBalancerID = loadBalancerID
	return b
}

func (b *AttachLoadBalancerBuilder) BackendServers(backendServers string) *AttachLoadBalancerBuilder {
	b.attachLoadBalancer.BackendServers = backendServers
	return b
}

func (b *AttachLoadBalancerBuilder) Build() (map[string]interface{}, error) {
	if b.attachLoadBalancer.LoadBalancerID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "LoadBalancerID")
	}
	if b.attachLoadBalancer.BackendServers == "" {
		return nil, errors.New(aliauth.StrMissRequired + "BackendServers")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.attachLoadBalancer)
	return params, nil
}

// DetachLoadBalancer builder pattern code
type DetachLoadBalancerBuilder struct {
	detachLoadBalancer *DetachLoadBalancer
}

func NewDetachLoadBalancerBuilder() *DetachLoadBalancerBuilder {
	detachLoadBalancer := &DetachLoadBalancer{}
	b := &DetachLoadBalancerBuilder{detachLoadBalancer: detachLoadBalancer}
	return b
}

func (b *DetachLoadBalancerBuilder) RegionID(regionID string) *DetachLoadBalancerBuilder {
	b.detachLoadBalancer.RegionID = regionID
	return b
}

func (b *DetachLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *DetachLoadBalancerBuilder {
	b.detachLoadBalancer.LoadBalancerID = loadBalancerID
	return b
}

func (b *DetachLoadBalancerBuilder) BackendServers(backendServers string) *DetachLoadBalancerBuilder {
	b.detachLoadBalancer.BackendServers = backendServers
	return b
}

func (b *DetachLoadBalancerBuilder) Build() (map[string]interface{}, error) {
	if b.detachLoadBalancer.RegionID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RegionID")
	}
	if b.detachLoadBalancer.LoadBalancerID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "LoadBalancerID")
	}
	if b.detachLoadBalancer.BackendServers == "" {
		return nil, errors.New(aliauth.StrMissRequired + "BackendServers")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.detachLoadBalancer)
	return params, nil
}
