package aws

import (
	ec2 "github.com/scorelab/gocloud-v2/compute/ec2"
	awsloadbalancer "github.com/scorelab/gocloud-v2/loadbalancer/awsloadbalancer"
	amazonstorage "github.com/scorelab/gocloud-v2/storage/amazonstorage"
)

//AWS struct reperents amazon cloud provider.
type AWS struct {
	ec2.EC2
	amazonstorage.Amazonstorage
	awsloadbalancer.Awsloadbalancer
}
