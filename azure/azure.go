package azure

import (
	azurecompute "github.com/cloudlibz/gocloud/compute/azurecompute"
	azurecontainer "github.com/cloudlibz/gocloud/container/azurecontainer"
	azurenosql "github.com/cloudlibz/gocloud/database/azurenosql"
	azuredns "github.com/cloudlibz/gocloud/dns/azuredns"
	azureloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/azureloadbalancer"
	azureserverless "github.com/cloudlibz/gocloud/serverless/azureserverless"
	azurestorage "github.com/cloudlibz/gocloud/storage/azurestorage"
	azuremachinelearning "github.com/cloudlibz/gocloud/machinelearning/azuremachinelearning"
)

// Azure  struct represents Microsoft Azure cloud provider.
type Azure struct {
	azurecompute.Azurecompute
	azurestorage.Azurestorage
	azureloadbalancer.Azureloadbalancer
	azurecontainer.Azurecontainer
	azuredns.Azuredns
	azureserverless.Azureserverless
	azurenosql.Azurenosql
	azuremachinelearning.Azuremachinelearning
}
