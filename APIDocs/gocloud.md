package gocloud
    import "github.com/shlokgilda/gocloud/gocloud"


CONSTANTS

const (
    //Amazonprovider reperents Amazon cloud.
    Amazonprovider = "aws"
    //Googleprovider reperents Google cloud.
    Googleprovider = "google"
)

TYPES

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
    Gocloud is a interface which hide differece between different cloud
    providers.

func CloudProvider(provider string) (Gocloud, error)
    CloudProvider return the instance of respepted cloud and map to the
    Gocloud so we can call the method like createnote on CloudProvider
    instance this is a delegation of cloud provider.


