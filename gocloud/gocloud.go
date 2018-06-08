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
	Createnode(request interface{}) (resp interface{}, err error)
	Startnode(request interface{}) (resp interface{}, err error)
	Stopnode(request interface{}) (resp interface{}, err error)
	Deletenode(request interface{}) (resp interface{}, err error)
	Rebootnode(request interface{}) (resp interface{}, err error)
	Createdisk(request interface{}) (resp interface{}, err error)
	Deletedisk(request interface{}) (resp interface{}, err error)
	Createsnapshot(request interface{}) (resp interface{}, err error)
	Deletesnapshot(request interface{}) (resp interface{}, err error)
	Attachdisk(request interface{}) (resp interface{}, err error)
	Detachdisk(request interface{}) (resp interface{}, err error)
	Createloadbalancer(request interface{}) (resp interface{}, err error)
	Deleteloadbalancer(request interface{}) (resp interface{}, err error)
	Listloadbalancer(request interface{}) (resp interface{}, err error)
	Attachnodewithloadbalancer(request interface{}) (resp interface{}, err error)
	Detachnodewithloadbalancer(request interface{}) (resp interface{}, err error)
	Createcluster(request interface{}) (resp interface{}, err error)
	Deletecluster(request interface{}) (resp interface{}, err error)
	Createservice(request interface{}) (resp interface{}, err error)
	Runtask(request interface{}) (resp interface{}, err error)
	Deleteservice(request interface{}) (resp interface{}, err error)
	Stoptask(request interface{}) (resp interface{}, err error)
	Starttask(request interface{}) (resp interface{}, err error)
	Listdns(request interface{}) (resp interface{}, err error)
	Createdns(request interface{}) (resp interface{}, err error)
	Deletedns(request interface{}) (resp interface{}, err error)
	ListResourcednsRecordSets(request interface{}) (resp interface{}, err error)
	Getfunction(request interface{}) (resp interface{}, err error)
	Createfunction(request interface{}) (resp interface{}, err error)
	Callfunction(request interface{}) (resp interface{}, err error)
	Listfunction(request interface{}) (resp interface{}, err error)
	Deletefunction(request interface{}) (resp interface{}, err error)
	Listtables(request interface{}) (resp interface{}, err error)
	Deletetables(request interface{}) (resp interface{}, err error)
	Describetables(request interface{}) (resp interface{}, err error)
	Createtables(request interface{}) (resp interface{}, err error)
	GetMLModel(request interface{}) (resp interface{}, err error)
	DeleteMLModel(request interface{}) (resp interface{}, err error)
	CreateMLModel(request interface{}) (resp interface{}, err error)
	UpdateMLModel(request interface{}) (resp interface{}, err error)
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
// Gocloud so that we can call the method like Createnode on CloudProvider instance.
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
