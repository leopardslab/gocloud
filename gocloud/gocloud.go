package gocloud

import (
	"errors"
	"fmt"
	"github.com/scorelab/gocloud/aws"
	"github.com/scorelab/gocloud/google"
)

//Gocloud is a interface which hide differece between different cloud providers.
type Gocloud interface {
	Createnode(request interface{}) (resp interface{}, err error)
	Startnode(request interface{}) (resp interface{}, err error)
	Stopnode(request interface{}) (resp interface{}, err error)
	Deletenode(request interface{}) (resp interface{}, err error)
	Rebootnode(request interface{}) (resp interface{}, err error)
}

const (
	Amazonprovider = "aws"
	Googleprovider = "google"
	Secretkey      = "SECRET_KEY"
	Secretid       = "SECRET_ID"
)

//cloud provider return the instance of respepted cloud and map to the Gocloud so we can call the method like createnote on CloudProvider instance
//this is a delegation of cloud provider.
func CloudProvider(provider, Secretkey, secretid string) (Gocloud, error) {

	switch provider {
	case Amazonprovider:
		return new(aws.AWS), nil

	case Googleprovider:
		return new(google.Google), nil

	default:
		return nil, errors.New(fmt.Sprintf("provider %s not recognized\n", provider))
	}
}
