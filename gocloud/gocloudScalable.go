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
)

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

// AWS
type awsProvider interface {
	goCloudCommon

	gocloudinterface.Serverless
	Serverless() gocloudinterface.Serverless

	gocloudinterface.Database
	Database() gocloudinterface.Database

	//gocloudinterface.MachineLearning
	MachineLearning() gocloudinterface.MachineLearning
}

func AWSProvider() awsProvider {
	awsAuth.LoadConfig()
	return new(aws.AWS)
}

// Google
type googleProvider interface {
	goCloudCommon

	gocloudinterface.Serverless
	Serverless() gocloudinterface.Serverless

	gocloudinterface.Database
	Database() gocloudinterface.Database
}

func GoogleProvider() googleProvider {
	return new(google.Google)
}

// Digital Ocean
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

func DigitalOceanProvider() digitalOceanProvider {
	digioceanauth.LoadConfig()
	return new(digiocean.DigitalOcean)
}

// Alibaba
type aliProvider interface {
	goCloudCommon
}

func AliCloudProvider() aliProvider {
	aliauth.LoadConfig()
	return new(ali.Ali)
}

//Vultr
type vultrProvider interface {
	gocloudinterface.Compute
	Compute() gocloudinterface.Compute
}

func VultrProvider() vultrProvider {
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
