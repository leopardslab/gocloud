package aws

import (
	ec2 "github.com/scorelab/gocloud-v2/compute/ec2"
	awscontainer "github.com/scorelab/gocloud-v2/container/awscontainer"
	awsloadbalancer "github.com/scorelab/gocloud-v2/loadbalancer/awsloadbalancer"
	amazonstorage "github.com/scorelab/gocloud-v2/storage/amazonstorage"
	awsdns "github.com/scorelab/gocloud-v2/dns/awsdns"
)

//AWS struct reperents amazon cloud provider.
type AWS struct {
	ec2.EC2
	amazonstorage.Amazonstorage
	awsloadbalancer.Awsloadbalancer
	awscontainer.Ecscontainer
	awsdns.Awsdns
}
