package awscontainer
    import "github.com/cloudlibz/gocloud/container/awscontainer"


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

type CreateService struct {
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
    CreateService struct represents Ecscontainer CreateService methods.

type DeleteService struct {
    Cluster string
    Service string
}
    DeleteService struct represents DeleteService attributes.

type DeploymentConfiguration struct {
    MaximumPercent        int
    MinimumHealthyPercent int
}
    DeploymentConfiguration struct represents DeploymentConfiguration.

type Ecscontainer struct {
}
    Ecscontainer struct represents Ecscontainer attribute and methods
    associates with it.

func (ecscontainer *Ecscontainer) CreateCluster(request interface{}) (resp interface{}, err error)
    CreateCluster creates cluster.

func (ecscontainer *Ecscontainer) CreateService(request interface{}) (resp interface{}, err error)
    CreateService creates container service.

func (ecscontainer *Ecscontainer) DeleteCluster(request interface{}) (resp interface{}, err error)
    DeleteCluster delete cluster.

func (ecscontainer *Ecscontainer) DeleteService(request interface{}) (resp interface{}, err error)
    DeleteService Delete container service.

func (ecscontainer *Ecscontainer) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}, response map[string]interface{}) error
    PrepareSignatureV4query creates PrepareSignatureV4 for request.

func (ecscontainer *Ecscontainer) RunTask(request interface{}) (resp interface{}, err error)
    RunTask runs container.

func (ecscontainer *Ecscontainer) StartTask(request interface{}) (resp interface{}, err error)
    StartTask start container service.

func (ecscontainer *Ecscontainer) StopTask(request interface{}) (resp interface{}, err error)
    StopTask stops container.

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

type RunTask struct {
    PlacementConstraints []Placementconstraint
    PlacementStrategys   []Placementstrategy
    Cluster              string
    Count                int
    Group                string
    StartedBy            string
    TaskDefinition       string
    // contains filtered or unexported fields
}
    RunTask struct represents RunTask attributes.

type StartTask struct {
    Cluster            string
    ContainerInstances []string
    Group              string
    StartedBy          string
    TaskDefinition     string
    // contains filtered or unexported fields
}
    StartTask struct represents StartTask attributes.

type StopTask struct {
    Cluster string
    Reason  string
    Task    string
}
    StopTask struct represents StopTask attributes.


