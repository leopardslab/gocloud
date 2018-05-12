package aws

import (
	ec2 "github.com/cloudlibz/gocloud/compute/ec2"
	awscontainer "github.com/cloudlibz/gocloud/container/awscontainer"
	awsdns "github.com/cloudlibz/gocloud/dns/awsdns"
	awsloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/awsloadbalancer"
	amazonstorage "github.com/cloudlibz/gocloud/storage/amazonstorage"
	lambda "github.com/cloudlibz/gocloud/serverless/lambda"
	dynamodb "github.com/cloudlibz/gocloud/database/dynamodb"

)

//AWS struct reperents amazon cloud provider.
type AWS struct {
	ec2.EC2
	amazonstorage.Amazonstorage
	awsloadbalancer.Awsloadbalancer
	awscontainer.Ecscontainer
	awsdns.Awsdns
	lambda.Lambda
	dynamodb.Dynamodb
}
