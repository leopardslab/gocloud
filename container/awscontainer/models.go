package awscontainer

//Ecscontainer struct represents Ecscontainer attribute and methods associates with it.
type Ecscontainer struct {
}

//Createservice struct represents Ecscontainer Createservice methods.
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

//Placementconstraint struct represents Ecscontainer Placement constraint.
type Placementconstraint struct {
	Expression string
	Type       string
}

//Placementstrategy struct represents Ecscontainer Placement strategy.
type Placementstrategy struct {
	Field string
	Type  string
}

//LoadBalancer struct represents LoadBalancer.
type LoadBalancer struct {
	ContainerName    string
	ContainerPort    int
	LoadBalancerName string
	TargetGroupArn   string
}

//DeploymentConfiguration struct represents DeploymentConfiguration.
type DeploymentConfiguration struct {
	MaximumPercent        int
	MinimumHealthyPercent int
}

//Runtask struct represents Runtask attributes.
type Runtask struct {
	PlacementConstraints []Placementconstraint
	PlacementStrategys   []Placementstrategy
	Cluster              string
	Count                int
	Group                string
	StartedBy            string
	TaskDefinition       string
	overrides            override
}

//Starttask struct represents Starttask attributes.
type Starttask struct {
	Cluster            string
	ContainerInstances []string
	Group              string
	StartedBy          string
	TaskDefinition     string
	overrides          override
}

//override struct represents override attributes.
type override struct {
	ContainerOverrides []ContainerOverride
	TaskRoleArn        string
}

//ContainerOverride struct represents ContainerOverride attributes.
type ContainerOverride struct {
	Name              string
	MemoryReservation string
	Memory            int
	Cpu               int
	Command           []string
	Environments      []Environment
}

//Environment struct represents Environment attributes.
type Environment struct {
	Name  string
	Value string
}

//Deleteservice  struct represents Deleteservice attributes.
type Deleteservice struct {
	Cluster string
	Service string
}

//Stoptask  struct represents Stoptask attributes.
type Stoptask struct {
	Cluster string
	Reason  string
	Task    string
}
