```
package vultrstorage
    import "github.com/cloudlibz/gocloud/storage/vultrstorage"
```

### TYPES

```
type AttachDisk struct {
    SUBID int // ID of the block storage subscription to attach
    // contains filtered or unexported fields
}

type AttachDiskBuilder struct {
    // contains filtered or unexported fields
}
    AttachDisk builder pattern code

func NewAttachDiskBuilder() *AttachDiskBuilder

func (b *AttachDiskBuilder) AttachToSUBID(attach_to_SUBID int) *AttachDiskBuilder

func (b *AttachDiskBuilder) Build() (map[string]interface{}, error)

func (b *AttachDiskBuilder) SUBID(sUBID int) *AttachDiskBuilder

type CreateDisk struct {
    DCID int // DCID of the location to create this subscription in.  See /v1/regions/list
    // contains filtered or unexported fields
}

type CreateDiskBuilder struct {
    // contains filtered or unexported fields
}
    CreateDisk builder pattern code

func NewCreateDiskBuilder() *CreateDiskBuilder

func (b *CreateDiskBuilder) Build() (map[string]interface{}, error)

func (b *CreateDiskBuilder) DCID(dCID int) *CreateDiskBuilder

func (b *CreateDiskBuilder) Label(label string) *CreateDiskBuilder

func (b *CreateDiskBuilder) SizeGb(size_gb int) *CreateDiskBuilder

type CreateSnapshot struct {
    SUBID int
    // contains filtered or unexported fields
}

type CreateSnapshotBuilder struct {
    // contains filtered or unexported fields
}
    CreateSnapshot builder pattern code

func NewCreateSnapshotBuilder() *CreateSnapshotBuilder

func (b *CreateSnapshotBuilder) Build() (map[string]interface{}, error)

func (b *CreateSnapshotBuilder) Description(description string) *CreateSnapshotBuilder

func (b *CreateSnapshotBuilder) SUBID(sUBID int) *CreateSnapshotBuilder

type DeleteDisk struct {
    SUBID int // ID of the block storage subscription to delete
}

type DeleteDiskBuilder struct {
    // contains filtered or unexported fields
}
    DeleteDisk builder pattern code

func NewDeleteDiskBuilder() *DeleteDiskBuilder

func (b *DeleteDiskBuilder) Build() (map[string]interface{}, error)

func (b *DeleteDiskBuilder) SUBID(sUBID int) *DeleteDiskBuilder

type DeleteSnapshot struct {
    SNAPSHOTID string
}

type DeleteSnapshotBuilder struct {
    // contains filtered or unexported fields
}
    DeleteSnapshot builder pattern code

func NewDeleteSnapshotBuilder() *DeleteSnapshotBuilder

func (b *DeleteSnapshotBuilder) Build() (map[string]interface{}, error)

func (b *DeleteSnapshotBuilder) SNAPSHOTID(sNAPSHOTID string) *DeleteSnapshotBuilder

type DetachDisk struct {
    SUBID int // ID of the block storage subscription to detach
}

type DetachDiskBuilder struct {
    // contains filtered or unexported fields
}
    DetachDisk builder pattern code

func NewDetachDiskBuilder() *DetachDiskBuilder

func (b *DetachDiskBuilder) Build() (map[string]interface{}, error)

func (b *DetachDiskBuilder) SUBID(sUBID int) *DetachDiskBuilder

type VultrStorage struct {
}

func (vultrStorage *VultrStorage) AttachDisk(request interface{}) (resp interface{}, err error)
    AttachDisk function function attaches a disk to a Vultr server.

func (vultrStorage *VultrStorage) CreateDisk(request interface{}) (resp interface{}, err error)
    CreateDisk function creates a new disk.

func (vultrStorage *VultrStorage) CreateSnapshot(request interface{}) (resp interface{}, err error)
    CreateSnapshot function creates a new snapshot.

func (vultrStorage *VultrStorage) DeleteDisk(request interface{}) (resp interface{}, err error)
    DeleteDisk function deletes a disk.

func (vultrStorage *VultrStorage) DeleteSnapshot(request interface{}) (resp interface{}, err error)
    DeleteSnapshot function deletes a snapshot.

func (vultrStorage *VultrStorage) DetachDisk(request interface{}) (resp interface{}, err error)
    DetachDisk function function detaches a disk from a Vultr server.
```
