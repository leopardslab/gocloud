package aws

import (
	ec2 "github.com/shlokgilda/gocloud/compute/ec2"
	awscontainer "github.com/shlokgilda/gocloud/container/awscontainer"
	awsdns "github.com/shlokgilda/gocloud/dns/awsdns"
	awsloadbalancer "github.com/shlokgilda/gocloud/loadbalancer/awsloadbalancer"
	amazonstorage "github.com/shlokgilda/gocloud/storage/amazonstorage"
)

//AWS struct reperents amazon cloud provider.
type AWS struct {
	ec2.EC2
	amazonstorage.Amazonstorage
	awsloadbalancer.Awsloadbalancer
	awscontainer.Ecscontainer
	awsdns.Awsdns
}
