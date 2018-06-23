package digiocean

import (
	droplet "github.com/cloudlibz/gocloud/compute/droplet"
	digioceancontainer "github.com/cloudlibz/gocloud/container/digioceancontainer"
	digioceannosql "github.com/cloudlibz/gocloud/database/digioceannosql"
	digioceandns "github.com/cloudlibz/gocloud/dns/digioceandns"
	"github.com/cloudlibz/gocloud/gocloudinterface"
	digioceanloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/digioceanloadbalancer"
	digioceanserverless "github.com/cloudlibz/gocloud/serverless/digioceanserverless"
	digioceanstorage "github.com/cloudlibz/gocloud/storage/digioceanstorage"
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
