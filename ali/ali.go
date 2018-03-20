package ali

import (
	"github.com/cloudlibz/gocloud/compute/ecs"
	"github.com/cloudlibz/gocloud/storage/alistorage"
	"github.com/cloudlibz/gocloud/loadbalancer/aliloadbalancer"
	"github.com/cloudlibz/gocloud/container/alicontainer"
	"github.com/cloudlibz/gocloud/dns/alidns"
)

//Ali struct represents Ali-cloud provider.
type Ali struct {
	ecs.ECS
	alistorage.Alistorage
	aliloadbalancer.Aliloadbalancer
	alicontainer.Alicontainer
	alidns.Alidns
}