```
package digiocean

import "github.com/cloudlibz/gocloud/digiocean"
```

### TYPES

DigitalOcean struct represents DigitalOcean cloud provider.
```
type DigitalOcean struct {
	droplet.Droplet
	digioceandns.Digioceandns
	digioceanloadbalancer.Digioceanloadbalancer
	digioceanstorage.Digioceanstorage
	digioceancontainer.Digioceancontainer
}
```
