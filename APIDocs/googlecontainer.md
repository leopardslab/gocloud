package googlecontainer
    import "github.com/cloudlibz/gocloud/container/googlecontainer"


FUNCTIONS

func CreateClusterLegacyAbacdictnoaryconvert(option CreateCluster, CreateClusterjsonmap map[string]interface{})
    CreateClusterLegacyAbacdictnoaryconvert create dictnoary for
    CreateCluster.

func CreateClusterMasterAuthdictnoaryconvert(option CreateCluster, CreateClusterjsonmap map[string]interface{})
    CreateClusterMasterAuthdictnoaryconvert create dictnoary for
    CreateCluster.

func CreateClusterNodePoolsdictnoaryconvert(option CreateCluster, CreateClusterjsonmap map[string]interface{})
    CreateClusterNodePoolsdictnoaryconvert create dictnoary for
    CreateCluster.

func CreateClusterdictnoaryconvert(option CreateCluster, CreateClusterjsonmap map[string]interface{})
    CreateClusterdictnoaryconvert create dictnoary for CreateCluster.

func CreateServiceAutoscalingdictnoaryconvert(option nodepool, CreateServicejsonmap map[string]interface{})
    CreateServiceAutoscalingdictnoaryconvert create dictnoary for
    CreateService.

func CreateServiceConfigdictnoaryconvert(option nodepool, CreateServicejsonmap map[string]interface{})
    CreateServiceConfigdictnoaryconvert create dictnoary for CreateService.

func CreateServiceManagementdictnoaryconvert(option nodepool, CreateServicejsonmap map[string]interface{})
    CreateServiceManagementdictnoaryconvert create dictnoary for
    CreateService.

func CreateServicedictnoaryconvert(option nodepool, CreateServicejsonmap map[string]interface{})
    CreateServicedictnoaryconvert create dictnoary for CreateService.

TYPES

type ClientCertificateConfigs struct {
    IssueClientCertificate bool `json:"issueClientCertificate"`
}
    ClientCertificateConfigs struct represents ClientCertificateConfigs
    attributes of masterAuth.

type CreateCluster struct {
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
    CreateCluster struct represents CreateCluster attributes.

type Googlecontainer struct {
}
    Googlecontainer struct represents Googlecontainer attributes and
    methods.

func (googlecontainer *Googlecontainer) CreateCluster(request interface{}) (resp interface{}, err error)
    CreateCluster creates cluster.

func (googlecontainer *Googlecontainer) CreateService(request interface{}) (resp interface{}, err error)
    CreateService crestes loadbalancer service.

func (googlecontainer *Googlecontainer) DeleteCluster(request interface{}) (resp interface{}, err error)
    DeleteCluster deletes cluster.

func (googlecontainer *Googlecontainer) DeleteService(request interface{}) (resp interface{}, err error)
    DeleteService crestes loadbalancer service.

func (googlecontainer *Googlecontainer) RunTask(request interface{}) (resp interface{}, err error)
    runtask runs container.

func (googlecontainer *Googlecontainer) StartTask(request interface{}) (resp interface{}, err error)
    StartTask start container.

func (googlecontainer *Googlecontainer) StopTask(request interface{}) (resp interface{}, err error)
    StopTask stops container.

type LegacyAbac struct {
    Enabled bool `json:"enabled"`
}
    LegacyAbac struct represents LegacyAbac attributes of CreateCluster.

type NodePool struct {
    Name             string      `json:"name"`
    InitialNodeCount int         `json:"initialNodeCount"`
    Config           config      `json:"config"`
    Autoscaling      autoscaling `json:"autoscaling"`
    Management       management  `json:"management"`
}
    NodePool struct represents NodePool attributes of CreateCluster.


