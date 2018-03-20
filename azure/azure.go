package azure

import (
	azurecompute "github.com/cloudlibz/gocloud/compute/azurecompute"
	azurestorage "github.com/cloudlibz/gocloud/container/azurestorage"
	azureloadbalancer "github.com/cloudlibz/gocloud/dns/azureloadbalancer"
	azurecontainer "github.com/cloudlibz/gocloud/loadbalancer/azurecontainer"
	azuredns "github.com/cloudlibz/gocloud/storage/azuredns"
)

// Azure  struct represents Microsoft Azure cloud provider.
type Azure struct {
	azurecompute.Azurecompute
	azurestorage.Azurestorage
	azureloadbalancer.Azureloadbalancer
	azurecontainer.Azurecontainer
  azuredns.Azuredns
}
