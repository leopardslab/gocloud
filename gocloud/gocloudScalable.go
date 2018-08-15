package gocloud

import (
	"github.com/cloudlibz/gocloud/ali"
	"github.com/cloudlibz/gocloud/aliauth"
	awsAuth "github.com/cloudlibz/gocloud/auth"
	"github.com/cloudlibz/gocloud/aws"
	"github.com/cloudlibz/gocloud/digiocean"
	"github.com/cloudlibz/gocloud/digioceanauth"
	"github.com/cloudlibz/gocloud/gocloudinterface"
	"github.com/cloudlibz/gocloud/google"
	"github.com/cloudlibz/gocloud/vultr"
	"github.com/cloudlibz/gocloud/vultrauth"
)

/*
goCloudCommon contains unified API of compute, storage, load balancer, container, DNS module
goCloudCommon also contains interface returning each module 's API. So users can do that for example awsProvider.Compute().CreateNode()
*/
type goCloudCommon interface {
	gocloudinterface.Compute
	Compute() gocloudinterface.Compute

	gocloudinterface.Storage
	Storage() gocloudinterface.Storage

	gocloudinterface.LoadBalancer
	LoadBalancer() gocloudinterface.LoadBalancer

	gocloudinterface.Container
	Container() gocloudinterface.Container

	gocloudinterface.DNS
	DNS() gocloudinterface.DNS
}

/*
awsProvider provides users with AWS API
AWS API contains common module API and particular module API
*/
type awsProvider interface {
	goCloudCommon

	gocloudinterface.Serverless
	Serverless() gocloudinterface.Serverless

	gocloudinterface.Database
	Database() gocloudinterface.Database

	gocloudinterface.MachineLearning
	MachineLearning() gocloudinterface.MachineLearning

	gocloudinterface.Analytics
	Analytics() gocloudinterface.Analytics

	gocloudinterface.Notification
	Notification() gocloudinterface.Notification

	gocloudinterface.Streamdataprocessing
	Streamdataprocessing() gocloudinterface.Streamdataprocessing

}

// AmazonProvider return AWS API to users
func AmazonProvider() awsProvider {
	awsAuth.LoadConfig()
	return new(aws.AWS)
}

/*
googleProvider provides users with Google cloud API
Google cloud API contains common module API and particular module API
*/
type googleProvider interface {
	goCloudCommon

	gocloudinterface.Serverless
	Serverless() gocloudinterface.Serverless

	gocloudinterface.Database
	Database() gocloudinterface.Database

	gocloudinterface.MachineLearning
	MachineLearning() gocloudinterface.MachineLearning

	gocloudinterface.Analytics
	Analytics() gocloudinterface.Analytics

	gocloudinterface.Notification
	Notification() gocloudinterface.Notification

	gocloudinterface.Streamdataprocessing
	Streamdataprocessing() gocloudinterface.Streamdataprocessing

}

// GoogleProvider return Google cloud API to users
func GoogleProvider() googleProvider {
	return new(google.Google)
}

// digitalOceanProvider contains compute, storage, load balancer, DNS module etc. which Digital Ocean supports
type digitalOceanProvider interface {
	gocloudinterface.Compute
	Compute() gocloudinterface.Compute

	gocloudinterface.Storage
	Storage() gocloudinterface.Storage

	gocloudinterface.LoadBalancer
	LoadBalancer() gocloudinterface.LoadBalancer

	gocloudinterface.DNS
	DNS() gocloudinterface.DNS

	//no container service
}

// DigitalOceanProvider return Digital Ocean API to users
func DigitalOceanProvider() digitalOceanProvider {
	digioceanauth.LoadConfig()
	return new(digiocean.DigitalOcean)
}

/*
aliProvider provides users with Alibaba cloud API
Ali-cloud API contains common module API
*/
type aliProvider interface {
	goCloudCommon
}

// AlibabaCloudProvider return Ali-cloud API to users
func AlibabaCloudProvider() aliProvider {
	aliauth.LoadConfig()
	return new(ali.Ali)
}

/**
vultrProvider provides users with Vultr API
Vultr API contains API which GoCloud implemented now
*/
type vultrProvider interface {
	gocloudinterface.Compute
	Compute() gocloudinterface.Compute

	gocloudinterface.Storage
	Storage() gocloudinterface.Storage

	gocloudinterface.DNS
	DNS() gocloudinterface.DNS

	BareMetal() gocloudinterface.BareMetal
}

// VultrProvider return Vultr API to users
func VultrProvider() vultrProvider {
	vultrauth.LoadConfig()
	return new(vultr.Vultr)
}

type azureProvider interface {
	// not yet
}

type openStackProvider interface {
	//not yet
}

type rackSpaceProvider interface {
	//not yet
}
