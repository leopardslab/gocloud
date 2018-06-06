package vultr

import (
	"github.com/cloudlibz/gocloud/compute/vultrcompute"
	"github.com/cloudlibz/gocloud/container/vultrcontainer"
	"github.com/cloudlibz/gocloud/dns/vultrdns"
	"github.com/cloudlibz/gocloud/loadbalancer/vultrloadbalancer"
	"github.com/cloudlibz/gocloud/storage/vultrstorage"
)

// Vultr struct represents Vultr cloud provider.
type Vultr struct {
	vultrcompute.VultrCompute
	vultrstorage.VultrStorage
	vultrloadbalancer.VultrLoadBalancer
	vultrcontainer.VultrContainer
	vultrdns.VultrDNS
}
