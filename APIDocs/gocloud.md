package gocloud
    import "github.com/cloudlibz/gocloud/gocloud"


CONSTANTS

const (
    //Amazonprovider reperents Amazon cloud.
    Amazonprovider = "aws"
    //Googleprovider reperents Google cloud.
    Googleprovider = "google"
)

TYPES

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
}
    Gocloud is a interface which hide differece between different cloud
    providers.

func CloudProvider(provider string) (Gocloud, error)
    CloudProvider return the instance of respepted cloud and map to the
    Gocloud so we can call the method like createnote on CloudProvider
    instance this is a delegation of cloud provider.


