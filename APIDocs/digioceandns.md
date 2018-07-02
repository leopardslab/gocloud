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

DeleteDns function deletes a DNS record.
```
func (digioceandns *Digioceandns) DeleteDns(request interface{}) (resp interface{}, err error)
```

CreateDns function creates a new DNS record.
```
func (digioceandns *Digioceandns) CreateDns(request interface{}) (resp interface{}, err error)
```

ListDns function lists DNS records.
```
func (digioceandns *Digioceandns) ListDns(request interface{}) (resp interface{}, err error)
```
