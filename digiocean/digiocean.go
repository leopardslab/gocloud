package digiocean

import (
	"github.com/cloudlibz/gocloud/compute/droplet"
	"github.com/cloudlibz/gocloud/dns/digioceandns"
	"github.com/cloudlibz/gocloud/loadbalancer/digioceanloadbalancer"
	"github.com/cloudlibz/gocloud/storage/digioceanstorage"
	"github.com/cloudlibz/gocloud/container/digioceancontainer"
	"github.com/cloudlibz/gocloud/serverless/digioceanserverless"
	"github.com/cloudlibz/gocloud/database/digioceannosql"
	"github.com/cloudlibz/gocloud/gocloudinterface"
)

// DigitalOcean struct represents DigitalOcean cloud provider.
type DigitalOcean struct {
	droplet.Droplet
	digioceandns.Digioceandns
	digioceanloadbalancer.DigioceanLoadBalancer
	digioceanstorage.Digioceanstorage
	digioceancontainer.Digioceancontainer
	digioceanserverless.Digioceanserverless
	digioceannosql.Digioceannosql
}

func (*DigitalOcean) Compute() gocloudinterface.Compute {
	return &droplet.Droplet{}
}

func (*DigitalOcean) Storage() gocloudinterface.Storage {
	return &digioceanstorage.Digioceanstorage{}
}

func (*DigitalOcean) LoadBalancer() gocloudinterface.LoadBalancer {
	return &digioceanloadbalancer.DigioceanLoadBalancer{}
}

func (*DigitalOcean) DNS() gocloudinterface.DNS {
	return &digioceandns.Digioceandns{}
}
