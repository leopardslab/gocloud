package googlecontainer
    import "github.com/cloudlibz/gocloud/container/googlecontainer"


FUNCTIONS

func CreateclusterLegacyAbacdictnoaryconvert(option Createcluster, Createclusterjsonmap map[string]interface{})
    CreateclusterLegacyAbacdictnoaryconvert create dictnoary for
    Createcluster.

func CreateclusterMasterAuthdictnoaryconvert(option Createcluster, Createclusterjsonmap map[string]interface{})
    CreateclusterMasterAuthdictnoaryconvert create dictnoary for
    Createcluster.

func CreateclusterNodePoolsdictnoaryconvert(option Createcluster, Createclusterjsonmap map[string]interface{})
    CreateclusterNodePoolsdictnoaryconvert create dictnoary for
    Createcluster.

func Createclusterdictnoaryconvert(option Createcluster, Createclusterjsonmap map[string]interface{})
    Createclusterdictnoaryconvert create dictnoary for Createcluster.

func CreateserviceAutoscalingdictnoaryconvert(option nodepool, Createservicejsonmap map[string]interface{})
    CreateserviceAutoscalingdictnoaryconvert create dictnoary for
    Createservice.

func CreateserviceConfigdictnoaryconvert(option nodepool, Createservicejsonmap map[string]interface{})
    CreateserviceConfigdictnoaryconvert create dictnoary for Createservice.

func CreateserviceManagementdictnoaryconvert(option nodepool, Createservicejsonmap map[string]interface{})
    CreateserviceManagementdictnoaryconvert create dictnoary for
    Createservice.

func Createservicedictnoaryconvert(option nodepool, Createservicejsonmap map[string]interface{})
    Createservicedictnoaryconvert create dictnoary for Createservice.

TYPES

type ClientCertificateConfigs struct {
    IssueClientCertificate bool `json:"issueClientCertificate"`
}
    ClientCertificateConfigs struct represents ClientCertificateConfigs
    attributes of masterAuth.

type Createcluster struct {
    Name                  string     `json:"name"`
    Zone                  string     `json:"zone"`
    Network               string     `json:"network"`
    LoggingService        string     `json:"loggingService"`
    MonitoringService     string     `json:"monitoringService"`
    InitialClusterVersion string     `json:"initialClusterVersion"`
    Subnetwork            string     `json:"subnetwork"`
    LegacyAbac            legacyAbac `json:"legacyAbac"`
    MasterAuth            masterAuth `json:"masterAuth"`
    NodePools             []NodePool `json:"nodePools"`
}
    Createcluster struct represents Createcluster attributes.

type Googlecontainer struct {
}
    Googlecontainer struct represents Googlecontainer attributes and
    methods.

func (googlecontainer *Googlecontainer) Createcluster(request interface{}) (resp interface{}, err error)
    Createcluster creates cluster.

func (googlecontainer *Googlecontainer) Createservice(request interface{}) (resp interface{}, err error)
    Createservice crestes loadbalancer service.

func (googlecontainer *Googlecontainer) Deletecluster(request interface{}) (resp interface{}, err error)
    Deletecluster deletes cluster.

func (googlecontainer *Googlecontainer) Deleteservice(request interface{}) (resp interface{}, err error)
    Deleteservice crestes loadbalancer service.

func (googlecontainer *Googlecontainer) Runtask(request interface{}) (resp interface{}, err error)
    runtask runs container.

func (googlecontainer *Googlecontainer) Starttask(request interface{}) (resp interface{}, err error)
    Starttask start container.

func (googlecontainer *Googlecontainer) Stoptask(request interface{}) (resp interface{}, err error)
    Stoptask stops container.

type LegacyAbac struct {
    Enabled bool `json:"enabled"`
}
    LegacyAbac struct represents LegacyAbac attributes of Createcluster.

type NodePool struct {
    Name             string      `json:"name"`
    InitialNodeCount int         `json:"initialNodeCount"`
    Config           config      `json:"config"`
    Autoscaling      autoscaling `json:"autoscaling"`
    Management       management  `json:"management"`
}
    NodePool struct represents NodePool attributes of Createcluster.


