package azure

import (
	azurecompute "github.com/cloudlibz/gocloud/compute/azurecompute"
	azurestorage "github.com/cloudlibz/gocloud/storage/azurestorage"
	azureloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/azureloadbalancer"
	azurecontainer "github.com/cloudlibz/gocloud/container/azurecontainer"
	azuredns "github.com/cloudlibz/gocloud/dns/azuredns"
	azureserverless "github.com/cloudlibz/gocloud/serverless/azureserverless"
	azurenosql "github.com/cloudlibz/gocloud/database/azurenosql"
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
}
