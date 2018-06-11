package vultr

import (
	"github.com/cloudlibz/gocloud/baremetal/vultrbaremetal"
	"github.com/cloudlibz/gocloud/compute/vultrcompute"
	"github.com/cloudlibz/gocloud/container/vultrcontainer"
	"github.com/cloudlibz/gocloud/database/vultrnosql"
	"github.com/cloudlibz/gocloud/dns/vultrdns"
	"github.com/cloudlibz/gocloud/gocloudinterface"
	"github.com/cloudlibz/gocloud/loadbalancer/vultrloadbalancer"
	"github.com/cloudlibz/gocloud/machinelearning/vultrmachinelearning"
	"github.com/cloudlibz/gocloud/serverless/vultrserverless"
	"github.com/cloudlibz/gocloud/storage/vultrstorage"
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

func (*Vultr) Compute() gocloudinterface.Compute {
	return &vultrcompute.VultrCompute{}
}

func (*Vultr) Storage() gocloudinterface.Storage {
	return &vultrstorage.VultrStorage{}
}

func (*Vultr) DNS() gocloudinterface.DNS {
	return &vultrdns.VultrDNS{}
}

func (*Vultr) BareMetal() gocloudinterface.BareMetal {
	return &vultrbaremetal.VultrBareMetal{}
}
