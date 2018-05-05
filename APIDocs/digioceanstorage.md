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

Createdisk function creates a new disk.
```
func (digioceanstorage *Digioceanstorage) Createdisk(request interface{}) (resp interface{}, err error)
```

Deletedisk function deletes a disk.
```
func (digioceanstorage *Digioceanstorage) Deletedisk(request interface{}) (resp interface{}, err error)
```

Createsnapshot function creates a new snapshot.
```
func (digioceanstorage *Digioceanstorage) Createsnapshot(request interface{}) (resp interface{}, err error)
```

Deletesnapshot function deletes a snapshot.
```
func (digioceanstorage *Digioceanstorage) Deletesnapshot(request interface{}) (resp interface{}, err error)
```

Attachdisk function function attaches a disk to a droplet.
```
func (digioceanstorage *Digioceanstorage) Attachdisk(request interface{}) (resp interface{}, err error)
```

Detachdisk function function detaches a disk from a droplet.
```
func (digioceanstorage *Digioceanstorage) Detachdisk(request interface{}) (resp interface{}, err error)
```
