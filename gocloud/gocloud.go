package gocloud

import (
	"errors"
	"fmt"
	"github.com/scorelab/gocloud-v2/auth"
	"github.com/scorelab/gocloud-v2/aws"
	"github.com/scorelab/gocloud-v2/google"
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
}

const (
	//Amazonprovider reperents Amazon cloud.
	Amazonprovider = "aws"
	//Googleprovider reperents Google cloud.
	Googleprovider = "google"
)

func init() {
	auth.LoadConfig()
}

//CloudProvider return the instance of respepted cloud and map to the Gocloud so we can call the method like createnote on CloudProvider instance
//this is a delegation of cloud provider.
func CloudProvider(provider string) (Gocloud, error) {

	switch provider {
	case Amazonprovider:
		return new(aws.AWS), nil

	case Googleprovider:
		return new(google.Google), nil

	default:
		return nil, errors.New(fmt.Sprintf("provider %s not recognized\n", provider))
	}

}
