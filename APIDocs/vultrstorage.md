```
package vultrstorage
    import "github.com/cloudlibz/gocloud/storage/vultrstorage"
```

TYPES

```
type AttachDisk struct {
    SUBID int // ID of the block storage subscription to attach
    // contains filtered or unexported fields
}

type CreateDisk struct {
    DCID int // DCID of the location to create this subscription in.  See /v1/regions/list
    // contains filtered or unexported fields
}

type CreateSnapshot struct {
    SUBID int
}

type DeleteDisk struct {
    SUBID int // ID of the block storage subscription to delete
}

type DeleteSnapshot struct {
    SNAPSHOTID string
}

type DetachDisk struct {
    SUBID int // ID of the block storage subscription to detach
}

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
