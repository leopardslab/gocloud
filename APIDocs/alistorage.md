```
package alistorage
    import "/home/oddcn/Code/go/src/github.com/cloudlibz/gocloud/storage/alistorage/"
```

### TYPES

```
type Alistorage struct {
}
```
    Alistorage struct represents Alistorage attribute and method associates with it.

```
func (aliStorage *Alistorage) Attachdisk(request interface{}) (resp interface{}, err error)
```
    Attachdisk attach ECS-Disk to ECS, accept map[string]interface{}

```
func (aliStorage *Alistorage) Createdisk(request interface{}) (resp interface{}, err error)
```
    Createdisk create ECS-Disk accept map[string]interface{}

```
func (aliStorage *Alistorage) Createsnapshot(request interface{}) (resp interface{}, err error)
```
    Createsnapshot create snapshot accept map[string]interface{}

```
func (aliStorage *Alistorage) Deletedisk(request interface{}) (resp interface{}, err error)
```
    Deletedisk delete ECS-Disk accept map[string]interface{}

```
func (aliStorage *Alistorage) Deletesnapshot(request interface{}) (resp interface{}, err error)
```
    Deletesnapshot delete snapshot accept map[string]interface{}

```
func (aliStorage *Alistorage) Detachdisk(request interface{}) (resp interface{}, err error)
```
    Detachdisk detach ECS-Disk from ECS, accept map[string]interface{}

```
type AttachDisk struct {
    InstanceID         string
    DiskID             string
    DeleteWithInstance bool
}
```
    AttachDisk to store all attribute of attach Ali-cloud ECS-Disk to ECS

```
type CreateDisk struct {
    RegionID     string
    ZoneID       string
    DiskName     string
    Description  string
    Encrypted    bool
    DiskCategory string
    Size         int
    SnapshotID   string
    ClientToken  string
}
```
    CreateDisk to store all attribute to create Ali-cloud ECS-Disk

```
type DeleteDisk struct {
    DiskID string
}
```
    DeleteDisk to store all attribute to delete Ali-cloud ECS-Disk

```
type DetachDisk struct {
    InstanceID string
    DiskID     string
}
```
    DetachDisk to store all attribute of detach Ali-cloud ECS-Disk from ECS


