package awscontainer
    import "github.com/scorelab/gocloud-v2/container/awscontainer"


TYPES

type ContainerOverride struct {
    Name              string
    MemoryReservation string
    Memory            int
    Cpu               int
    Command           []string
    Environments      []Environment
}
    ContainerOverride struct represents ContainerOverride attributes.

type Createservice struct {
    ServiceName              string
    TaskDefinition           string
    DesiredCount             int
    ClientToken              string
    Cluster                  string
    Role                     string
    DeploymentConfigurations DeploymentConfiguration
    LoadBalancers            []LoadBalancer
    PlacementConstraints     []Placementconstraint
    PlacementStrategys       []Placementstrategy
}
    Createservice struct represents Ecscontainer Createservice methods.

type Deleteservice struct {
    Cluster string
    Service string
}
    Deleteservice struct represents Deleteservice attributes.

type DeploymentConfiguration struct {
    MaximumPercent        int
    MinimumHealthyPercent int
}
    DeploymentConfiguration struct represents DeploymentConfiguration.

type Ecscontainer struct {
}
    Ecscontainer struct represents Ecscontainer attribute and methods
    associates with it.

func (ecscontainer *Ecscontainer) Createcluster(request interface{}) (resp interface{}, err error)
    Createcluster creates cluster.

func (ecscontainer *Ecscontainer) Createservice(request interface{}) (resp interface{}, err error)
    Createservice creates container service.

func (ecscontainer *Ecscontainer) Deletecluster(request interface{}) (resp interface{}, err error)
    Deletecluster delete cluster.

func (ecscontainer *Ecscontainer) Deleteservice(request interface{}) (resp interface{}, err error)
    Deleteservice Delete container service.

func (ecscontainer *Ecscontainer) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}, response map[string]interface{}) error
    PrepareSignatureV4query creates PrepareSignatureV4 for request.

func (ecscontainer *Ecscontainer) Runtask(request interface{}) (resp interface{}, err error)
    Runtask runs container.

func (ecscontainer *Ecscontainer) Starttask(request interface{}) (resp interface{}, err error)
    Starttask start container service.

func (ecscontainer *Ecscontainer) Stoptask(request interface{}) (resp interface{}, err error)
    Stoptask stops container.

type Environment struct {
    Name  string
    Value string
}
    Environment struct represents Environment attributes.

type LoadBalancer struct {
    ContainerName    string
    ContainerPort    int
    LoadBalancerName string
    TargetGroupArn   string
}
    LoadBalancer struct represents LoadBalancer.

type Placementconstraint struct {
    Expression string
    Type       string
}
    Placementconstraint struct represents Ecscontainer Placement constraint.

type Placementstrategy struct {
    Field string
    Type  string
}
    Placementstrategy struct represents Ecscontainer Placement strategy.

type Runtask struct {
    PlacementConstraints []Placementconstraint
    PlacementStrategys   []Placementstrategy
    Cluster              string
    Count                int
    Group                string
    StartedBy            string
    TaskDefinition       string
    // contains filtered or unexported fields
}
    Runtask struct represents Runtask attributes.

type Starttask struct {
    Cluster            string
    ContainerInstances []string
    Group              string
    StartedBy          string
    TaskDefinition     string
    // contains filtered or unexported fields
}
    Starttask struct represents Starttask attributes.

type Stoptask struct {
    Cluster string
    Reason  string
    Task    string
}
    Stoptask struct represents Stoptask attributes.


