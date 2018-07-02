```
package alistorage
    import "github.com/cloudlibz/gocloud/storage/alistorage"
```

### TYPES

```
type Alistorage struct {
}
```
Alistorage struct represents Alistorage attribute and method associates with it.

```
func (aliStorage *Alistorage) AttachDisk(request interface{}) (resp interface{}, err error)
```
AttachDisk attach ECS-Disk to ECS, accept map[string]interface{}

```
func (aliStorage *Alistorage) CreateDisk(request interface{}) (resp interface{}, err error)
```
CreateDisk create ECS-Disk accept map[string]interface{}

```
func (aliStorage *Alistorage) CreateSnapshot(request interface{}) (resp interface{}, err error)
```
CreateSnapshot create snapshot accept map[string]interface{}

```
func (aliStorage *Alistorage) DeleteDisk(request interface{}) (resp interface{}, err error)
```
DeleteDisk delete ECS-Disk accept map[string]interface{}

```
func (aliStorage *Alistorage) DeleteSnapshot(request interface{}) (resp interface{}, err error)
```
DeleteSnapshot delete snapshot accept map[string]interface{}

```
func (aliStorage *Alistorage) DetachDisk(request interface{}) (resp interface{}, err error)
```
DetachDisk detach ECS-Disk from ECS, accept map[string]interface{}

```
type AttachDisk struct {
    InstanceID         string
    DiskID             string
    DeleteWithInstance bool
}
```
AttachDisk to store all attribute of attach Ali-cloud ECS-Disk to ECS

```
type AttachDiskBuilder struct {
    // contains filtered or unexported fields
}

func NewAttachDiskBuilder() *AttachDiskBuilder

func (b *AttachDiskBuilder) Build() (map[string]interface{}, error)

func (b *AttachDiskBuilder) DeleteWithInstance(deleteWithInstance bool) *AttachDiskBuilder

func (b *AttachDiskBuilder) DiskID(diskID string) *AttachDiskBuilder

func (b *AttachDiskBuilder) InstanceID(instanceID string) *AttachDiskBuilder
```
AttachDisk builder pattern code

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
type CreateDiskBuilder struct {
    // contains filtered or unexported fields
}

func NewCreateDiskBuilder() *CreateDiskBuilder

func (b *CreateDiskBuilder) Build() (map[string]interface{}, error)

func (b *CreateDiskBuilder) ClientToken(clientToken string) *CreateDiskBuilder

func (b *CreateDiskBuilder) Description(description string) *CreateDiskBuilder

func (b *CreateDiskBuilder) DiskCategory(diskCategory string) *CreateDiskBuilder

func (b *CreateDiskBuilder) DiskName(diskName string) *CreateDiskBuilder

func (b *CreateDiskBuilder) Encrypted(encrypted bool) *CreateDiskBuilder

func (b *CreateDiskBuilder) RegionID(regionID string) *CreateDiskBuilder

func (b *CreateDiskBuilder) Size(size int) *CreateDiskBuilder

func (b *CreateDiskBuilder) SnapshotID(snapshotID string) *CreateDiskBuilder

func (b *CreateDiskBuilder) ZoneID(zoneID string) *CreateDiskBuilder
```
CreateDisk builder pattern code

```
type CreateSnapshot struct {
    DiskID       string
    SnapshotName string
    Description  string
    ClientToken  string
}
```
CreateSnapshot to store all attribute of create Ali-cloud ECS-Disk 's Snapshot

```
type CreateSnapshotBuilder struct {
    // contains filtered or unexported fields
}

func NewCreateSnapshotBuilder() *CreateSnapshotBuilder

func (b *CreateSnapshotBuilder) Build() (map[string]interface{}, error)

func (b *CreateSnapshotBuilder) ClientToken(clientToken string) *CreateSnapshotBuilder

func (b *CreateSnapshotBuilder) Description(description string) *CreateSnapshotBuilder

func (b *CreateSnapshotBuilder) DiskID(diskID string) *CreateSnapshotBuilder

func (b *CreateSnapshotBuilder) SnapshotName(snapshotName string) *CreateSnapshotBuilder
```
CreateSnapshot builder pattern code

```
type DeleteDisk struct {
    DiskID string
}
```
DeleteDisk to store all attribute to delete Ali-cloud ECS-Disk

```
type DeleteDiskBuilder struct {
    // contains filtered or unexported fields
}

func NewDeleteDiskBuilder() *DeleteDiskBuilder

func (b *DeleteDiskBuilder) Build() (map[string]interface{}, error)

func (b *DeleteDiskBuilder) DiskID(diskID string) *DeleteDiskBuilder
```
DeleteDisk builder pattern code

```
type DeleteSnapshot struct {
    SnapshotID string
}
```
DeleteSnapshot to store all attribute of delete Ali-cloud ECS-Disk 's Snapshot

```
type DeleteSnapshotBuilder struct {
    // contains filtered or unexported fields
}

func NewDeleteSnapshotBuilder() *DeleteSnapshotBuilder

func (b *DeleteSnapshotBuilder) Build() (map[string]interface{}, error)

func (b *DeleteSnapshotBuilder) SnapshotID(snapshotID string) *DeleteSnapshotBuilder
```
DeleteSnapshot builder pattern code

```
type DetachDisk struct {
    InstanceID string
    DiskID     string
}
```
DetachDisk to store all attribute of detach Ali-cloud ECS-Disk from ECS

```
type DetachDiskBuilder struct {
    // contains filtered or unexported fields
}

func NewDetachDiskBuilder() *DetachDiskBuilder

func (b *DetachDiskBuilder) Build() (map[string]interface{}, error)

func (b *DetachDiskBuilder) DiskID(diskID string) *DetachDiskBuilder

func (b *DetachDiskBuilder) InstanceID(instanceID string) *DetachDiskBuilder
```
DetachDisk builder pattern code

