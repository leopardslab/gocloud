package awscontainer


type Ecscontainer struct {
}

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

type Placementconstraint struct {
	Expression string
	Type       string
}

type Placementstrategy struct {
	Field string
	Type  string
}

type LoadBalancer struct {
	ContainerName    string
	ContainerPort    int
	LoadBalancerName string
	TargetGroupArn   string
}

type DeploymentConfiguration struct {
	MaximumPercent        int
	MinimumHealthyPercent int
}

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

type Starttask struct {
	Cluster            string
	ContainerInstances []string
	Group              string
	StartedBy          string
	TaskDefinition     string
	overrides          override
}

type override struct {
	ContainerOverrides []ContainerOverride
	TaskRoleArn        string
}

type ContainerOverride struct {
	Name              string
	MemoryReservation string
	Memory            int
	Cpu               int
	Command           []string
	Environments      []Environment
}

type Environment struct {
	Name  string
	Value string
}

type Deleteservice struct {
	Cluster string
	Service string
}

type Stoptask struct {
	Cluster string
	Reason  string
	Task    string
}
