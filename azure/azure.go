package azure

import (
	azurecompute "github.com/shlokgilda/gocloud/compute/azurecompute"
	azurestorage "github.com/shlokgilda/gocloud/container/azurestorage"
	azureloadbalancer "github.com/shlokgilda/gocloud/dns/azureloadbalancer"
	azurecontainer "github.com/shlokgilda/gocloud/loadbalancer/azurecontainer"
	azuredns "github.com/shlokgilda/gocloud/storage/azuredns"
)

// Azure  struct represents Microsoft Azure cloud provider.
type Azure struct {
	azurecompute.Azurecompute
	azurestorage.Azurestorage
	azureloadbalancer.Azureloadbalancer
	azurecontainer.Azurecontainer
  azuredns.Azuredns
}
