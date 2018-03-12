package gocloud

import (
	"errors"
	"fmt"
	"github.com/cloudlibz/gocloud/auth"
	"github.com/cloudlibz/gocloud/aws"
	"github.com/cloudlibz/gocloud/google"
	"github.com/cloudlibz/gocloud/aliauth"
	"github.com/cloudlibz/gocloud/ali"
)

//Gocloud is a interface which hide differece between different cloud providers.
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
	Creatloadbalancer(request interface{}) (resp interface{}, err error)
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
}

const (
	//Amazonprovider reperents Amazon cloud.
	Amazonprovider = "aws"
	//Googleprovider reperents Google cloud.
	Googleprovider = "google"
	//Aliprovider reperents Google cloud.
	Aliprovider = "ali"
)

func init() {
	auth.LoadConfig()
	aliauth.LoadConfig()
}

//CloudProvider return the instance of respepted cloud and map to the Gocloud so we can call the method like createnote on CloudProvider instance
//this is a delegation of cloud provider.
func CloudProvider(provider string) (Gocloud, error) {

	switch provider {
	case Amazonprovider:
		return new(aws.AWS), nil

	case Googleprovider:
		return new(google.Google), nil

	case Aliprovider:
		return new(ali.Ali), nil

	default:
		return nil, errors.New(fmt.Sprintf("provider %s not recognized\n", provider))
	}

}
