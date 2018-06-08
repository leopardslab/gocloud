```
package aliloadbalancer
    import "github.com/cloudlibz/gocloud/loadbalancer/aliloadbalancer"
```

### TYPES

```
type Aliloadbalancer struct {
}
```
Aliloadbalancer represents Aliloadbalancer struct.

```
func (aliloadbalancer *Aliloadbalancer) Attachnodewithloadbalancer(request interface{}) (resp interface{}, err error)
```
Attachnodewithloadbalancer attach ali ecs instance to ali loadbalancer

```
func (aliloadbalancer *Aliloadbalancer) Createloadbalancer(request interface{}) (resp interface{}, err error)
```
Createloadbalancer creates ali loadbalancer

```
func (aliloadbalancer *Aliloadbalancer) Deleteloadbalancer(request interface{}) (resp interface{}, err error)
```
Deleteloadbalancer deletes ali loadbalancer

```
func (aliloadbalancer *Aliloadbalancer) Detachnodewithloadbalancer(request interface{}) (resp interface{}, err error)
```
Detachnodewithloadbalancer detach ali ecs instance from ali loadbalancer

```
func (aliloadbalancer *Aliloadbalancer) Listloadbalancer(request interface{}) (resp interface{}, err error)
```
Listloadbalancer lists ali loadbalancer

```
type AttachLoadBalancer struct {
    LoadBalancerID string
    BackendServers string
}
```
AttachLoadBalancer represents Attach node with loadbalancer attribute.

```
type AttachLoadBalancerBuilder struct {
    // contains filtered or unexported fields
}

func NewAttachLoadBalancerBuilder() *AttachLoadBalancerBuilder

func (b *AttachLoadBalancerBuilder) BackendServers(backendServers string) *AttachLoadBalancerBuilder

func (b *AttachLoadBalancerBuilder) Build() (map[string]interface{}, error)

func (b *AttachLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *AttachLoadBalancerBuilder
```
AttachLoadBalancer builder pattern code

```
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
```
CreateLoadBalancer struct represents attribute of create LoadBalancer.

```
type CreateLoadBalancerBuilder struct {
    // contains filtered or unexported fields
}

func NewCreateLoadBalancerBuilder() *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) AddressType(addressType string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) Autopay(autopay bool) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) Bandwidth(bandwidth int) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) Build() (map[string]interface{}, error)

func (b *CreateLoadBalancerBuilder) ClientToken(clientToken string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) Duration(duration string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) InternetChargeType(internetChargeType string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) LoadBalancerName(loadBalancerName string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) LoadBalancerSpec(loadBalancerSpec string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) MasterZoneID(masterZoneID string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) PayType(payType string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) PricingCycle(pricingCycle string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) RegionID(regionID string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) ResourceGroupID(resourceGroupID string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) SlaveZoneID(slaveZoneID string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) VSwitchID(vSwitchID string) *CreateLoadBalancerBuilder
```
CreateLoadBalancer builder pattern code

```
type DeleteLoadBalancer struct {
    RegionID       string
    LoadBalancerID string
}
```
DeleteLoadBalancer struct represents attribute of delete LoadBalancer.

```
type DeleteLoadBalancerBuilder struct {
    // contains filtered or unexported fields
}

func NewDeleteLoadBalancerBuilder() *DeleteLoadBalancerBuilder

func (b *DeleteLoadBalancerBuilder) Build() (map[string]interface{}, error)

func (b *DeleteLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *DeleteLoadBalancerBuilder

func (b *DeleteLoadBalancerBuilder) RegionID(regionID string) *DeleteLoadBalancerBuilder
```
DeleteLoadBalancer builder pattern code

```
type DetachLoadBalancer struct {
    RegionID       string
    LoadBalancerID string
    BackendServers string
}
```
DetachLoadBalancer represents Detach node with loadbalancer attribute.

```
type DetachLoadBalancerBuilder struct {
    // contains filtered or unexported fields
}

func NewDetachLoadBalancerBuilder() *DetachLoadBalancerBuilder

func (b *DetachLoadBalancerBuilder) BackendServers(backendServers string) *DetachLoadBalancerBuilder

func (b *DetachLoadBalancerBuilder) Build() (map[string]interface{}, error)

func (b *DetachLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *DetachLoadBalancerBuilder

func (b *DetachLoadBalancerBuilder) RegionID(regionID string) *DetachLoadBalancerBuilder
```
DetachLoadBalancer builder pattern code

```
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
```
ListLoadBalancer struct represents attribute of list LoadBalancer.

```
type ListLoadBalancerBuilder struct {
    // contains filtered or unexported fields
}

func NewListLoadBalancerBuilder() *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) Address(address string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) AddressType(addressType string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) Build() (map[string]interface{}, error)

func (b *ListLoadBalancerBuilder) InternetChargeType(internetChargeType string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) LoadBalancerName(loadBalancerName string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) MasterZoneID(masterZoneID string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) NetworkType(networkType string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) RegionID(regionID string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) ServerID(serverID string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) ServerIntranetAddress(serverIntranetAddress int) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) SlaveZoneID(slaveZoneID string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) VpcID(vpcID string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) VswitchID(vswitchID string) *ListLoadBalancerBuilder
```
ListLoadBalancer builder pattern code
