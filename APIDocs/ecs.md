```
package ecs
    import "/home/oddcn/Code/go/src/github.com/cloudlibz/gocloud/compute/ecs"
```

### TYPES

```
type CreateInstance struct {
    RegionID                string
    ZoneID                  string
    ImageID                 string
    InstanceType            string
    SecurityGroupID         string
    InstanceName            string
    Description             string
    InternetChargeType      string
    InternetMaxBandwidthIn  int
    InternetMaxBandwidthOut int
    HostName                string
    Password                string
    IoOptimized             string
    SystemDiskCategory      string
    SystemDiskSize          string
    SystemDiskName          string
    SystemDiskDescription   string
}
```

CreateInstance to store all attribute to create Ali-cloud ECS instance

```
type CreateNodeBuilder struct {
    // contains filtered or unexported fields
}
```

```
func NewCreateNodeBuilder() *CreateNodeBuilder

func (b *CreateNodeBuilder) Build() (map[string]interface{}, error)

func (b *CreateNodeBuilder) Description(description string) *CreateNodeBuilder

func (b *CreateNodeBuilder) HostName(hostName string) *CreateNodeBuilder

func (b *CreateNodeBuilder) ImageID(imageID string) *CreateNodeBuilder

func (b *CreateNodeBuilder) InstanceName(instanceName string) *CreateNodeBuilder

func (b *CreateNodeBuilder) InstanceType(instanceType string) *CreateNodeBuilder

func (b *CreateNodeBuilder) InternetChargeType(internetChargeType string) *CreateNodeBuilder

func (b *CreateNodeBuilder) InternetMaxBandwidthIn(internetMaxBandwidthIn int) *CreateNodeBuilder

func (b *CreateNodeBuilder) InternetMaxBandwidthOut(internetMaxBandwidthOut int) *CreateNodeBuilder

func (b *CreateNodeBuilder) IoOptimized(ioOptimized string) *CreateNodeBuilder

func (b *CreateNodeBuilder) Password(password string) *CreateNodeBuilder

func (b *CreateNodeBuilder) RegionID(regionID string) *CreateNodeBuilder

func (b *CreateNodeBuilder) SecurityGroupID(securityGroupID string) *CreateNodeBuilder

func (b *CreateNodeBuilder) SystemDiskCategory(systemDiskCategory string) *CreateNodeBuilder

func (b *CreateNodeBuilder) SystemDiskDescription(systemDiskDescription string) *CreateNodeBuilder

func (b *CreateNodeBuilder) SystemDiskName(systemDiskName string) *CreateNodeBuilder

func (b *CreateNodeBuilder) SystemDiskSize(systemDiskSize string) *CreateNodeBuilder

func (b *CreateNodeBuilder) ZoneID(zoneID string) *CreateNodeBuilder
```

```
type DeleteInstance struct {
    InstanceID string
}
```
    DeleteInstance to store all attribute to Delete Ali-cloud ECS instance

```
type ECS struct {
}
```
    Ali ECS struct

```
func (ecs *ECS) Createnode(request interface{}) (resp interface{}, err error)
```
    Createnode create ECS instances accept map[string]interface{}

```
func (ecs *ECS) Deletenode(request interface{}) (resp interface{}, err error)
```
    Deletenode delete ECS instances accept map[string]interface{}

```
func (ecs *ECS) Rebootnode(request interface{}) (resp interface{}, err error)
```
    Rebootnode reboot ECS instances accept map[string]interface{}

```
func (ecs *ECS) Startnode(request interface{}) (resp interface{}, err error)
```
    Startnode start ECS instances accept map[string]interface{}

```
func (ecs *ECS) Stopnode(request interface{}) (resp interface{}, err error)
```
    Stopnode stop ECS instances accept map[string]interface{}

```
type RebootInstance struct {
    InstanceID string
    ForceStop  bool
}
```
    RebootInstance to store all attribute to Reboot Ali-cloud ECS instance

```
type StartInstance struct {
    InstanceID    string
    InitLocalDisk bool
}
```
    StartInstance to store all attribute to start Ali-cloud ECS instance

```
type StopInstance struct {
    InstanceID  string
    ForceStop   bool
    ConfirmStop bool
    StoppedMode string
}
```
    StopInstance to store all attribute to Stop Ali-cloud ECS instance

