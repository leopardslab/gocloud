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
func (aliloadbalancer *Aliloadbalancer) Creatloadbalancer(request interface{}) (resp interface{}, err error)
```
Creatloadbalancer creates ali loadbalancer

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
type DeleteLoadBalancer struct {
    RegionID       string
    LoadBalancerID string
}
```
DeleteLoadBalancer struct represents attribute of delete LoadBalancer.

```
type DetachLoadBalancer struct {
    RegionID       string
    LoadBalancerID string
    BackendServers string
}
```
DetachLoadBalancer represents Detach node with loadbalancer attribute.

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

