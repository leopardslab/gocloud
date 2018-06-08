package rackspace

import (
	rackspacecompute "github.com/cloudlibz/gocloud/compute/rackspacecompute"
	rackspacecontainer "github.com/cloudlibz/gocloud/container/rackspacecontainer"
	rackspacenosql "github.com/cloudlibz/gocloud/database/rackspacenosql"
	rackspacedns "github.com/cloudlibz/gocloud/dns/rackspacedns"
	rackspaceloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/rackspaceloadbalancer"
	rackspacemachinelearning "github.com/cloudlibz/gocloud/machinelearning/rackspacemachinelearning"
	rackspaceserverless "github.com/cloudlibz/gocloud/serverless/rackspaceserverless"
	rackspacestorage "github.com/cloudlibz/gocloud/storage/rackspacestorage"
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
	rackspacemachinelearning.Rackspacemachinelearning
}
