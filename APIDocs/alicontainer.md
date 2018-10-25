```
package alicontainer
    import "github.com/cloudlibz/gocloud/container/alicontainer"
```

### TYPES

```
type Alicontainer struct {
}
    Alicontainer struct represents Alicontainer attribute and methods
    associates with it.

func (alicontainer *Alicontainer) CreateCluster(request interface{}) (resp interface{}, err error)
    CreateCluster creates container cluster

func (alicontainer *Alicontainer) CreateService(request interface{}) (resp interface{}, err error)
    CreateService Create service is not provided by Alibaba cloud

func (alicontainer *Alicontainer) DeleteCluster(request interface{}) (resp interface{}, err error)
    DeleteCluster deletes container cluster

func (alicontainer *Alicontainer) DeleteService(request interface{}) (resp interface{}, err error)
    DeleteService delete service is not provided by Alibaba cloud

func (alicontainer *Alicontainer) RunTask(request interface{}) (resp interface{}, err error)
    RunTask creates project of container cluster

func (alicontainer *Alicontainer) StartTask(request interface{}) (resp interface{}, err error)
    RunTask starts project of container cluster

func (alicontainer *Alicontainer) StopTask(request interface{}) (resp interface{}, err error)
    RunTask stops project of container cluster

type CreateCluster struct {
    RegionID         string `json:"region_id"`
    Name             string `json:"name"`
    Size             int64  `json:"size"`
    InstanceType     string `json:"instance_type"`
    NetworkMode      string `json:"network_mode"`
    SubnetCIDR       string `json:"subnet_cidr,omitempty"`
    VPCID            string `json:"vpc_id,omitempty"`
    VSwitchID        string `json:"vswitch_id,omitempty"`
    Password         string `json:"password"`
    DataDiskCategory string `json:"data_disk_category"`
    DataDiskSize     int64  `json:"data_disk_size"`
    ECSImageID       string `json:"ecs_image_id,omitempty"`
    IOOptimized      string `json:"io_optimized"`
    NeedSLB          bool   `json:"need_slb"`
    ReleaseEipFlag   bool   `json:"release_eip_flag"`
}
    CreateCluster struct represents create cluster attributes.

type CreateClusterBuilder struct {
    // contains filtered or unexported fields
}
    CreateCluster builder pattern code

func NewCreateClusterBuilder() *CreateClusterBuilder

func (b *CreateClusterBuilder) Build() (map[string]interface{}, error)

func (b *CreateClusterBuilder) DataDiskCategory(dataDiskCategory string) *CreateClusterBuilder

func (b *CreateClusterBuilder) DataDiskSize(dataDiskSize int64) *CreateClusterBuilder

func (b *CreateClusterBuilder) ECSImageID(eCSImageID string) *CreateClusterBuilder

func (b *CreateClusterBuilder) IOOptimized(iOOptimized string) *CreateClusterBuilder

func (b *CreateClusterBuilder) InstanceType(instanceType string) *CreateClusterBuilder

func (b *CreateClusterBuilder) Name(name string) *CreateClusterBuilder

func (b *CreateClusterBuilder) NeedSLB(needSLB bool) *CreateClusterBuilder

func (b *CreateClusterBuilder) NetworkMode(networkMode string) *CreateClusterBuilder

func (b *CreateClusterBuilder) Password(password string) *CreateClusterBuilder

func (b *CreateClusterBuilder) RegionID(regionID string) *CreateClusterBuilder

func (b *CreateClusterBuilder) ReleaseEipFlag(releaseEipFlag bool) *CreateClusterBuilder

func (b *CreateClusterBuilder) Size(size int64) *CreateClusterBuilder

func (b *CreateClusterBuilder) SubnetCIDR(subnetCIDR string) *CreateClusterBuilder

func (b *CreateClusterBuilder) VPCID(vPCID string) *CreateClusterBuilder

func (b *CreateClusterBuilder) VSwitchID(vSwitchID string) *CreateClusterBuilder

type CreateClusterResp struct {
    StatusCode int
    ClusterID  string `json:"cluster_id"`
    TaskID     string `json:"task_id"`
}

func ParseCreateClusterResp(resp interface{}) (createClusterResp CreateClusterResp, err error)

type DeleteCluster struct {
    RegionID  string `json:"region_id"`
    ClusterID string `json:"cluster_id"`
}
    DeleteCluster struct represents create cluster attributes.

type DeleteClusterBuilder struct {
    // contains filtered or unexported fields
}
    DeleteCluster builder pattern code

func NewDeleteClusterBuilder() *DeleteClusterBuilder

func (b *DeleteClusterBuilder) Build() (map[string]interface{}, error)

func (b *DeleteClusterBuilder) ClusterID(clusterID string) *DeleteClusterBuilder

func (b *DeleteClusterBuilder) RegionID(regionID string) *DeleteClusterBuilder

type RunTask struct {
    ClusterID   string            `json:"cluster_id"`
    Name        string            `json:"name"`
    Description string            `json:"description"`
    Template    string            `json:"template"`
    Version     string            `json:"version"`
    Environment map[string]string `json:"environment"`
    LatestImage bool              `json:"latest_image"`
}
    RunTask struct represents create project of cluster attributes.

type RunTaskBuilder struct {
    // contains filtered or unexported fields
}
    RunTask builder pattern code

func NewRunTaskBuilder() *RunTaskBuilder

func (b *RunTaskBuilder) Build() (map[string]interface{}, error)

func (b *RunTaskBuilder) ClusterID(clusterID string) *RunTaskBuilder

func (b *RunTaskBuilder) Description(description string) *RunTaskBuilder

func (b *RunTaskBuilder) Environment(environment map[string]string) *RunTaskBuilder

func (b *RunTaskBuilder) LatestImage(latestImage bool) *RunTaskBuilder

func (b *RunTaskBuilder) Name(name string) *RunTaskBuilder

func (b *RunTaskBuilder) Template(template string) *RunTaskBuilder

func (b *RunTaskBuilder) Version(version string) *RunTaskBuilder

type StartTask struct {
    ClusterID string `json:"cluster_id"`
    Name      string `json:"name"`
}
    StartTask struct represents create project of cluster attributes.

type StartTaskBuilder struct {
    // contains filtered or unexported fields
}
    StartTask builder pattern code

func NewStartTaskBuilder() *StartTaskBuilder

func (b *StartTaskBuilder) Build() (map[string]interface{}, error)

func (b *StartTaskBuilder) ClusterID(clusterID string) *StartTaskBuilder

func (b *StartTaskBuilder) Name(name string) *StartTaskBuilder

type StopTask struct {
    ClusterID string `json:"cluster_id"`
    Name      string `json:"name"`
    Timeout   int    `json:"timeout"`
}
    StopTask struct represents create project of cluster attributes.

type StopTaskBuilder struct {
    // contains filtered or unexported fields
}
    StopTask builder pattern code

func NewStopTaskBuilder() *StopTaskBuilder

func (b *StopTaskBuilder) Build() (map[string]interface{}, error)

func (b *StopTaskBuilder) ClusterID(clusterID string) *StopTaskBuilder

func (b *StopTaskBuilder) Name(name string) *StopTaskBuilder

func (b *StopTaskBuilder) Timeout(timeout int) *StopTaskBuilder
```
