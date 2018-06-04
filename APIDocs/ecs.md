```
package ecs
    import "github.com/cloudlibz/gocloud/compute/ecs"
```

### CONSTANTS

```
const (
    DefaultRegion = "ecs.aliyuncs.com"
    Zhangjiakou   = "ecs.cn-zhangjiakou.aliyuncs.com"
    Hohhot        = "ecs.cn-huhehaote.aliyuncs.com"
    Tokyo         = "ecs.ap-northeast-1.aliyuncs.com"
    Sydney        = "ecs.ap-southeast-2.aliyuncs.com"
    KualaLumpur   = "ecs.ap-southeast-3.aliyuncs.com"
    Jakarta       = "ecs.ap-southeast-5.aliyuncs.com"
    Mumbai        = "ecs.ap-south-1.aliyuncs.com"
    Dubai         = "ecs.me-east-1.aliyuncs.com"
    Frankfurt     = "ecs.eu-central-1.aliyuncs.com"
)
```
Endpoint of Ali ecs

### TYPES

```
type CreateNode struct {
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
CreateNode to store all attribute to create Ali-cloud ECS instance

```
type CreateNodeBuilder struct {
    // contains filtered or unexported fields
}

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
CreateNode builder pattern code

```
type DeleteNode struct {
    InstanceID string
}
```
DeleteNode to store all attribute to Delete Ali-cloud ECS instance

```
type DeleteNodeBuilder struct {
    // contains filtered or unexported fields
}

func NewDeleteNodeBuilder() *DeleteNodeBuilder

func (b *DeleteNodeBuilder) Build() (map[string]interface{}, error)

func (b *DeleteNodeBuilder) InstanceID(instanceID string) *DeleteNodeBuilder
```
DeleteNode builder pattern code

```
type ECS struct {
}
```
ECS struct

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
type RebootNode struct {
    InstanceID string
    ForceStop  bool
}
```
RebootNode to store all attribute to Reboot Ali-cloud ECS instance

```
type RebootNodeBuilder struct {
    // contains filtered or unexported fields
}

func NewRebootNodeBuilder() *RebootNodeBuilder

func (b *RebootNodeBuilder) Build() (map[string]interface{}, error)

func (b *RebootNodeBuilder) ForceStop(forceStop bool) *RebootNodeBuilder

func (b *RebootNodeBuilder) InstanceID(instanceID string) *RebootNodeBuilder
```
RebootNode builder pattern code

```
type StartNode struct {
    InstanceID    string
    InitLocalDisk bool
}
```
StartNode to store all attribute to start Ali-cloud ECS instance

```
type StartNodeBuilder struct {
    // contains filtered or unexported fields
}

func NewStartNodeBuilder() *StartNodeBuilder

func (b *StartNodeBuilder) Build() (map[string]interface{}, error)

func (b *StartNodeBuilder) InitLocalDisk(initLocalDisk bool) *StartNodeBuilder

func (b *StartNodeBuilder) InstanceID(instanceID string) *StartNodeBuilder
```
StartNode builder pattern code

```
type StopNode struct {
    InstanceID  string
    ForceStop   bool
    ConfirmStop bool
    StoppedMode string
}
```
StopNode to store all attribute to Stop Ali-cloud ECS instance

```
type StopNodeBuilder struct {
    // contains filtered or unexported fields
}

func NewStopNodeBuilder() *StopNodeBuilder

func (b *StopNodeBuilder) Build() (map[string]interface{}, error)

func (b *StopNodeBuilder) ConfirmStop(confirmStop bool) *StopNodeBuilder

func (b *StopNodeBuilder) ForceStop(forceStop bool) *StopNodeBuilder

func (b *StopNodeBuilder) InstanceID(instanceID string) *StopNodeBuilder

func (b *StopNodeBuilder) StoppedMode(stoppedMode string) *StopNodeBuilder
```
StopNode builder pattern code

