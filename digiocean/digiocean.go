package digiocean

import (
	droplet "github.com/shlokgilda/gocloud/compute/droplet"
	digioceandns "github.com/shlokgilda/gocloud/dns/digioceandns"
	digioceanloadbalancer "github.com/shlokgilda/gocloud/loadbalancer/digioceanloadbalancer"
	digioceanstorage "github.com/shlokgilda/gocloud/storage/digioceanstorage"
	digioceancontainer "github.com/shlokgilda/gocloud/container/digioceancontainer"
)

// DigitalOcean  struct represents Digital Ocean cloud provider.
type DigitalOcean struct {
	droplet.Droplet
	digioceandns.Digioceandns
	digioceanloadbalancer.Digioceanloadbalancer
	digioceanstorage.Digioceanstorage
	digioceancontainer.Digioceancontainer
}
