package rackspace

import (
	rackspacecompute "github.com/cloudlibz/gocloud/compute/rackspacecompute"
	rackspacestorage "github.com/cloudlibz/gocloud/storage/rackspacestorage"
	rackspaceloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/rackspaceloadbalancer"
	rackspacecontainer "github.com/cloudlibz/gocloud/container/rackspacecontainer"
	rackspacedns "github.com/cloudlibz/gocloud/dns/rackspacedns"
)

// rackspace  struct represents rackspace cloud provider.
type Rackspace struct {
	rackspacecompute.Rackspacecompute
	rackspacestorage.Rackspacestorage
	rackspaceloadbalancer.Rackspaceloadbalancer
	rackspacecontainer.Rackspacecontainer
  rackspacedns.Rackspacedns
}
