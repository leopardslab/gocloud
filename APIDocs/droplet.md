```
package digiocean

import "github.com/cloudlibz/gocloud/compute/droplet"
```

### CONSTANTS

dropletBasePath is the endpoint URL for digitalocean API.
```
const dropletBasePath
```

### TYPES

Droplet represents a request to create a droplet.
```
type Droplet struct {
	Name              string                `json:"name"`
	Region            string                `json:"region"`
	Size              string                `json:"size"`
	Image             CreateImage           `json:"image"`
	SSHKeys           []CreateSSHKey        `json:"ssh_keys"`
	Backups           bool                  `json:"backups"`
	IPv6              bool                  `json:"ipv6"`
	PrivateNetworking bool                  `json:"privateNetworking"`
	Monitoring        bool                  `json:"monitoring"`
	UserData          string                `json:"userData,omitempty"`
	Volumes           []CreateVolume        `json:"volumes,omitempty"`
	Tags              []string              `json:"tags"`
}
```

CreateImage identifies an image for the create request. It prefers slug over ID.
```
type CreateImage struct {
	ID   int	`json:"id,omitempty"`
	Slug string	`json:"name,omitempty"`
}
```

CreateSSHKey identifies a SSH Key for the create request. It prefers fingerprint over ID.
```
type CreateSSHKey struct {
	ID          string
	Fingerprint string
}
```

CreateVolume identifies a volume to attach for the create request. It prefers Name over ID,
```
type CreateVolume struct {
	ID   string
	Name string
}
```

ActionRequest reprents DigitalOcean Action Request
```
type ActionRequest map[string]interface{}
```

### FUNCTIONS

MarshalJSON returns either the slug or id of the image. It returns the id if the slug is empty.
```
func (d CreateImage) MarshalJSON() ([]byte, error)
```

MarshalJSON returns either the fingerprint or id of the ssh key. It returns the id if the fingerprint is empty.
```
func (d CreateSSHKey) MarshalJSON() ([]byte, error)
```

MarshalJSON returns an object with either the name or id of the volume. It returns the id if the name is empty.
```
func (d CreateVolume) MarshalJSON() ([]byte, error)
```

Createnode function creates a new droplet.
```
func (droplet *Droplet) Createnode(request interface{}) (resp interface{}, err error)
```

Startnode function starts a droplet.
```
func (droplet *Droplet) Startnode(request interface{}) (resp interface{}, err error)
```

Stopnode function stops a droplet.
```
func (droplet *Droplet) Stopnode(request interface{}) (resp interface{}, err error)
```

Rebootnode function reboots a droplet.
```
func (droplet *Droplet) Rebootnode(request interface{}) (resp interface{}, err error)
```

Deletenode function deletes a droplet.
```
func (droplet *Droplet) Deletenode(request interface{}) (resp interface{}, err error)
```
