package awscontainer

//Ecscontainer struct represents Ecscontainer attribute and methods associates with it.
type Ecscontainer struct {
}

//CreateService struct represents Ecscontainer CreateService methods.
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

//RunTask struct represents RunTask attributes.
type RunTask struct {
	PlacementConstraints []Placementconstraint
	PlacementStrategys   []Placementstrategy
	Cluster              string
	Count                int
	Group                string
	StartedBy            string
	TaskDefinition       string
	overrides            override
}

//StartTask struct represents StartTask attributes.
type StartTask struct {
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

//DeleteService  struct represents DeleteService attributes.
type DeleteService struct {
	Cluster string
	Service string
}

//StopTask  struct represents StopTask attributes.
type StopTask struct {
	Cluster string
	Reason  string
	Task    string
}
