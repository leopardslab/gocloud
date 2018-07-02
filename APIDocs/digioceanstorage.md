```
package digioceanloadbalancer

import "github.com/cloudlibz/gocloud/loadbalancer/digioceanloadbalancer"
```

### CONSTANTS

storageBasePath is the endpoint URL for digitalocean API.
```
const storageBasePath
```

### TYPES

Digioceanstorage struct represents a DigitalOcean storage object.
```
type Digioceanstorage struct {
	Name          string `json:"name"`
  Region        string `json:"region,omitempty"`
	Description   string `json:"description,omitempty"`
	SizeGigaBytes int64  `json:"size_gigabytes"`
	SnapshotID    string `json:"snapshot_id,omitempty"`
}
```

### FUNCTIONS

CreateDisk function creates a new disk.
```
func (digioceanstorage *Digioceanstorage) CreateDisk(request interface{}) (resp interface{}, err error)
```

DeleteDisk function deletes a disk.
```
func (digioceanstorage *Digioceanstorage) DeleteDisk(request interface{}) (resp interface{}, err error)
```

CreateSnapshot function creates a new snapshot.
```
func (digioceanstorage *Digioceanstorage) CreateSnapshot(request interface{}) (resp interface{}, err error)
```

DeleteSnapshot function deletes a snapshot.
```
func (digioceanstorage *Digioceanstorage) DeleteSnapshot(request interface{}) (resp interface{}, err error)
```

AttachDisk function function attaches a disk to a droplet.
```
func (digioceanstorage *Digioceanstorage) AttachDisk(request interface{}) (resp interface{}, err error)
```

DetachDisk function function detaches a disk from a droplet.
```
func (digioceanstorage *Digioceanstorage) DetachDisk(request interface{}) (resp interface{}, err error)
```
