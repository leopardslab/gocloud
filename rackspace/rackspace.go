package rackspace

import (
	rackspacecompute "github.com/cloudlibz/gocloud/compute/rackspacecompute"
	rackspacestorage "github.com/cloudlibz/gocloud/storage/rackspacestorage"
	rackspaceloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/rackspaceloadbalancer"
	rackspacecontainer "github.com/cloudlibz/gocloud/container/rackspacecontainer"
	rackspacedns "github.com/cloudlibz/gocloud/dns/rackspacedns"
	rackspaceserverless "github.com/cloudlibz/gocloud/serverless/rackspaceserverless"
	rackspacenosql "github.com/cloudlibz/gocloud/database/rackspacenosql"
)

// Rackspace  struct represents Rackspace cloud provider.
type Rackspace struct {
	rackspacecompute.Rackspacecompute
	rackspacestorage.Rackspacestorage
	rackspaceloadbalancer.Rackspaceloadbalancer
	rackspacecontainer.Rackspacecontainer
  rackspacedns.Rackspacedns
	rackspaceserverless.Rackspaceserverless
	rackspacenosql.Rackspacenosql
}
