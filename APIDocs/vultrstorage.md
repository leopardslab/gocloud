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

func (vultrStorage *VultrStorage) Attachdisk(request interface{}) (resp interface{}, err error)
    Attachdisk function function attaches a disk to a Vultr server.

func (vultrStorage *VultrStorage) Createdisk(request interface{}) (resp interface{}, err error)
    Createdisk function creates a new disk.

func (vultrStorage *VultrStorage) Createsnapshot(request interface{}) (resp interface{}, err error)
    Createsnapshot function creates a new snapshot.

func (vultrStorage *VultrStorage) Deletedisk(request interface{}) (resp interface{}, err error)
    Deletedisk function deletes a disk.

func (vultrStorage *VultrStorage) Deletesnapshot(request interface{}) (resp interface{}, err error)
    Deletesnapshot function deletes a snapshot.

func (vultrStorage *VultrStorage) Detachdisk(request interface{}) (resp interface{}, err error)
    Detachdisk function function detaches a disk from a Vultr server.

```
