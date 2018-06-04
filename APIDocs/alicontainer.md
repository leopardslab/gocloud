```
package alicontainer
    import "github.com/cloudlibz/gocloud/container/alicontainer"
```

### TYPES

```
type Alicontainer struct {
}
```
Alicontainer struct represents Alicontainer attribute and methods associates with it.

```
func (alicontainer *Alicontainer) Createcluster(request interface{}) (resp interface{}, err error)
```
Createcluster creates container cluster

```
func (alicontainer *Alicontainer) Createservice(request interface{}) (resp interface{}, err error)
```
Createservice Create service is not provided by Alibaba cloud

```
func (alicontainer *Alicontainer) Deletecluster(request interface{}) (resp interface{}, err error)
```
Deletecluster deletes container cluster

```
func (alicontainer *Alicontainer) Deleteservice(request interface{}) (resp interface{}, err error)
```
Deleteservice delete service is not provided by Alibaba cloud

```
func (alicontainer *Alicontainer) Runtask(request interface{}) (resp interface{}, err error)
```
Runtask creates project of container cluster

```
func (alicontainer *Alicontainer) Starttask(request interface{}) (resp interface{}, err error)
```
Runtask starts project of container cluster

```
func (alicontainer *Alicontainer) Stoptask(request interface{}) (resp interface{}, err error)
```
Runtask stops project of container cluster

```
type CreateCluster struct {
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
```
CreateCluster struct represents create cluster attributes.

```
type RunTask struct {
    Name        string            `json:"name"`
    Description string            `json:"description"`
    Template    string            `json:"template"`
    Version     string            `json:"version"`
    Environment map[string]string `json:"environment"`
    LatestImage bool              `json:"latest_image"`
}
```
RunTask struct represents create project of cluster attributes.
