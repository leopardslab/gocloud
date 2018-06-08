package vultr

import (
	"github.com/cloudlibz/gocloud/compute/vultrcompute"
	"github.com/cloudlibz/gocloud/container/vultrcontainer"
	"github.com/cloudlibz/gocloud/database/vultrnosql"
	"github.com/cloudlibz/gocloud/dns/vultrdns"
	"github.com/cloudlibz/gocloud/loadbalancer/vultrloadbalancer"
	"github.com/cloudlibz/gocloud/serverless/vultrserverless"
	"github.com/cloudlibz/gocloud/storage/vultrstorage"
	"github.com/cloudlibz/gocloud/machinelearning/vultrmachinelearning"
)

// Vultr struct represents Vultr cloud provider.
type Vultr struct {
	vultrcompute.VultrCompute
	vultrstorage.VultrStorage
	vultrloadbalancer.VultrLoadBalancer
	vultrcontainer.VultrContainer
	vultrdns.VultrDNS
	vultrserverless.Vultrserverless
	vultrnosql.Vultrnosql
	vultrmachinelearning.Vultrmachinelearning
}
