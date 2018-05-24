package digiocean

import (
	droplet "github.com/cloudlibz/gocloud/compute/droplet"
	digioceancontainer "github.com/cloudlibz/gocloud/container/digioceancontainer"
	digioceannosql "github.com/cloudlibz/gocloud/database/digioceannosql"
	digioceandns "github.com/cloudlibz/gocloud/dns/digioceandns"
	digioceanloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/digioceanloadbalancer"
	digioceanserverless "github.com/cloudlibz/gocloud/serverless/digioceanserverless"
	digioceanstorage "github.com/cloudlibz/gocloud/storage/digioceanstorage"
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
