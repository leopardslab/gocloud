package ali

import (
	"github.com/cloudlibz/gocloud/compute/aliecs"
	"github.com/cloudlibz/gocloud/storage/alistorage"
	"github.com/cloudlibz/gocloud/loadbalancer/aliloadbalancer"
	"github.com/cloudlibz/gocloud/container/alicontainer"
	"github.com/cloudlibz/gocloud/dns/alidns"
)

type Ali struct {
	aliecs.ECS
	alistorage.AliStorage
	aliloadbalancer.Aliloadbalancer
	alicontainer.Alicontainer
	alidns.Alidns
}