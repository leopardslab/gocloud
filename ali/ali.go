package ali

import (
	"github.com/cloudlibz/gocloud/compute/ecs"
	"github.com/cloudlibz/gocloud/container/alicontainer"
	"github.com/cloudlibz/gocloud/database/alinosql"
	"github.com/cloudlibz/gocloud/dns/alidns"
	"github.com/cloudlibz/gocloud/gocloudinterface"
	"github.com/cloudlibz/gocloud/loadbalancer/aliloadbalancer"
	alimachinelearning "github.com/cloudlibz/gocloud/machinelearning/alimachinelearning"
	"github.com/cloudlibz/gocloud/serverless/aliserverless"
	"github.com/cloudlibz/gocloud/storage/alistorage"
)

//Ali struct represents Ali-cloud provider.
type Ali struct {
	ecs.ECS
	alistorage.Alistorage
	aliloadbalancer.Aliloadbalancer
	alicontainer.Alicontainer
	alidns.Alidns
	aliserverless.Aliserverless
	alinosql.Alinosql
	alimachinelearning.Alimachinelearning
}

func (*Ali) Compute() gocloudinterface.Compute {
	return &ecs.ECS{}
}

func (*Ali) Storage() gocloudinterface.Storage {
	return &alistorage.Alistorage{}
}

func (*Ali) LoadBalancer() gocloudinterface.LoadBalancer {
	return &aliloadbalancer.Aliloadbalancer{}
}

func (*Ali) Container() gocloudinterface.Container {
	return &alicontainer.Alicontainer{}
}

func (*Ali) DNS() gocloudinterface.DNS {
	return &alidns.Alidns{}
}
