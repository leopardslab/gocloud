package gocloud

import (
	"fmt"
	"github.com/cloudlibz/gocloud/ali"
	aliAuth "github.com/cloudlibz/gocloud/aliauth"
	awsAuth "github.com/cloudlibz/gocloud/auth"
	"github.com/cloudlibz/gocloud/aws"
	"github.com/cloudlibz/gocloud/azure"
	"github.com/cloudlibz/gocloud/digiocean"
	digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
	"github.com/cloudlibz/gocloud/google"
	"github.com/cloudlibz/gocloud/openstack"
	"github.com/cloudlibz/gocloud/rackspace"
	"github.com/cloudlibz/gocloud/vultr"
	"github.com/cloudlibz/gocloud/vultrauth"
)
// Gocloud is a interface which hides the difference between different cloud providers.
type Gocloud interface {
	CreateNode(request interface{}) (resp interface{}, err error)
	StartNode(request interface{}) (resp interface{}, err error)
	StopNode(request interface{}) (resp interface{}, err error)
	DeleteNode(request interface{}) (resp interface{}, err error)
	RebootNode(request interface{}) (resp interface{}, err error)
	CreateDisk(request interface{}) (resp interface{}, err error)
	DeleteDisk(request interface{}) (resp interface{}, err error)
	CreateSnapshot(request interface{}) (resp interface{}, err error)
	DeleteSnapshot(request interface{}) (resp interface{}, err error)
	AttachDisk(request interface{}) (resp interface{}, err error)
	DetachDisk(request interface{}) (resp interface{}, err error)
	CreateLoadBalancer(request interface{}) (resp interface{}, err error)
	DeleteLoadBalancer(request interface{}) (resp interface{}, err error)
	ListLoadBalancer(request interface{}) (resp interface{}, err error)
	AttachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error)
	DetachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error)
	CreateCluster(request interface{}) (resp interface{}, err error)
	DeleteCluster(request interface{}) (resp interface{}, err error)
	CreateService(request interface{}) (resp interface{}, err error)
	RunTask(request interface{}) (resp interface{}, err error)
	DeleteService(request interface{}) (resp interface{}, err error)
	StopTask(request interface{}) (resp interface{}, err error)
	StartTask(request interface{}) (resp interface{}, err error)
	ListDns(request interface{}) (resp interface{}, err error)
	CreateDns(request interface{}) (resp interface{}, err error)
	DeleteDns(request interface{}) (resp interface{}, err error)
	ListResourceDnsRecordSets(request interface{}) (resp interface{}, err error)
	GetFunction(request interface{}) (resp interface{}, err error)
	CreateFunction(request interface{}) (resp interface{}, err error)
	CallFunction(request interface{}) (resp interface{}, err error)
	ListFunction(request interface{}) (resp interface{}, err error)
	DeleteFunction(request interface{}) (resp interface{}, err error)
	ListTables(request interface{}) (resp interface{}, err error)
	DeleteTables(request interface{}) (resp interface{}, err error)
	DescribeTables(request interface{}) (resp interface{}, err error)
	CreateTables(request interface{}) (resp interface{}, err error)
}
const (
	// Amazonprovider represents Amazon cloud.
	Amazonprovider = "aws"

	// Googleprovider represents Google cloud.
	Googleprovider = "google"

	// Openstackprovider represents Openstack cloud.
	Openstackprovider = "openstack"

	// Azureprovider represents Openstack cloud.
	Azureprovider = "azure"

	// Digioceanprovider represents Digital Ocean cloud.
	Digioceanprovider = "digiocean"

	// Aliprovider reperents Google cloud.
	Aliprovider = "ali"

	// Rackspaceprovider reperents rackspace cloud.
	Rackspaceprovider = "rackspace"

	// Rackspaceprovider reperents rackspace cloud.
	Vultrprovider = "vultr"
)
// CloudProvider returns the instance of respective cloud and maps it to
// Gocloud so that we can call the method like CreateNode on CloudProvider instance.
// This is a delegation of CloudProvider.
func CloudProvider(provider string) (Gocloud, error) {

	switch provider {
	case Amazonprovider:
		// Calls authentication procedure for AWS
		awsAuth.LoadConfig()
		return new(aws.AWS), nil

	case Googleprovider:
		return new(google.Google), nil

	case Openstackprovider:
		return new(openstack.Openstack), nil

	case Digioceanprovider:
		// Calls authentication procedure for Digital Ocean.
		digioceanAuth.LoadConfig()
		return new(digiocean.DigitalOcean), nil

	case Azureprovider:
		return new(azure.Azure), nil

	case Aliprovider:
		aliAuth.LoadConfig()
		return new(ali.Ali), nil

	case Rackspaceprovider:
		return new(rackspace.Rackspace), nil

	case Vultrprovider:
		vultrauth.LoadConfig()
		return new(vultr.Vultr), nil

	default:
		return nil, fmt.Errorf("provider %s not recognized", provider)
	}
}
