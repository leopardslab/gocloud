```
package digioceanloadbalancer

import "github.com/cloudlibz/gocloud/loadbalancer/digioceanloadbalancer"
```

### CONSTANTS

dnsBasePath is the endpoint URL for digitalocean API.
```
const dnsBasePath
```

### TYPES

Digioceandns struct represents a DigitalOcean DNS service.
```
type Digioceandns struct {
	Type       string `json:"type"`
	Name       string `json:"name"`
	Data       string `json:"data"`
	Priority   int    `json:priority,omitempty`
	Port       int    `json:port,omitempty`
	TimeToLive int    `json:ttl`
	Weight     int    `json:weight,omitempty`
	Flags      int    `json:flags,omitempty`
	Tag        string `json:"tag"`
}
```

### FUNCTIONS

Deletedns function deletes a DNS record.
```
func (digioceandns *Digioceandns) Deletedns(request interface{}) (resp interface{}, err error)
```

Createdns function creates a new DNS record.
```
func (digioceandns *Digioceandns) Createdns(request interface{}) (resp interface{}, err error)
```
