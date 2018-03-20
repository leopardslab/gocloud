package azure

import (
	azurecompute "github.com/shlokgilda/gocloud/compute/azurecompute"
	azurestorage "github.com/shlokgilda/gocloud/storage/azurestorage"
	azureloadbalancer "github.com/shlokgilda/gocloud/loadbalancer/azureloadbalancer"
	azurecontainer "github.com/shlokgilda/gocloud/container/azurecontainer"
	azuredns "github.com/shlokgilda/gocloud/dns/azuredns"
)

// Azure  struct represents Microsoft Azure cloud provider.
type Azure struct {
	azurecompute.Azurecompute
	azurestorage.Azurestorage
	azureloadbalancer.Azureloadbalancer
	azurecontainer.Azurecontainer
  azuredns.Azuredns
}
