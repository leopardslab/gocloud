package digiocean

import (
	droplet "github.com/cloudlibz/gocloud/compute/droplet"
	digioceandns "github.com/cloudlibz/gocloud/dns/digioceandns"
	digioceanloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/digioceanloadbalancer"
	digioceanstorage "github.com/cloudlibz/gocloud/storage/digioceanstorage"
	digioceancontainer "github.com/cloudlibz/gocloud/container/digioceancontainer"
	digioceanserverless "github.com/cloudlibz/gocloud/serverless/digioceanserverless"
	digioceannosql "github.com/cloudlibz/gocloud/database/digioceannosql"
)

// DigitalOcean struct represents DigitalOcean cloud provider.
type DigitalOcean struct {
	droplet.Droplet
	digioceandns.Digioceandns
	digioceanloadbalancer.LoadBalancer
	digioceanstorage.Digioceanstorage
	digioceancontainer.Digioceancontainer
	digioceanserverless.Digioceanserverless
	digioceannosql.Digioceannosql
}
