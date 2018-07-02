package gocloudinterface

// Compute module unified API
type Compute interface {
	CreateNode(request interface{}) (resp interface{}, err error)
	StartNode(request interface{}) (resp interface{}, err error)
	StopNode(request interface{}) (resp interface{}, err error)
	DeleteNode(request interface{}) (resp interface{}, err error)
	RebootNode(request interface{}) (resp interface{}, err error)
}

// Storage module unified API
type Storage interface {
	CreateDisk(request interface{}) (resp interface{}, err error)
	DeleteDisk(request interface{}) (resp interface{}, err error)
	CreateSnapshot(request interface{}) (resp interface{}, err error)
	DeleteSnapshot(request interface{}) (resp interface{}, err error)
	AttachDisk(request interface{}) (resp interface{}, err error)
	DetachDisk(request interface{}) (resp interface{}, err error)
}

// LoadBalancer module unified API
type LoadBalancer interface {
	CreateLoadBalancer(request interface{}) (resp interface{}, err error)
	DeleteLoadBalancer(request interface{}) (resp interface{}, err error)
	ListLoadBalancer(request interface{}) (resp interface{}, err error)
	AttachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error)
	DetachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error)
}

// Container module unified API
type Container interface {
	CreateCluster(request interface{}) (resp interface{}, err error)
	DeleteCluster(request interface{}) (resp interface{}, err error)
	CreateService(request interface{}) (resp interface{}, err error)
	RunTask(request interface{}) (resp interface{}, err error)
	DeleteService(request interface{}) (resp interface{}, err error)
	StopTask(request interface{}) (resp interface{}, err error)
	StartTask(request interface{}) (resp interface{}, err error)
}

// DNS module unified API
type DNS interface {
	ListDns(request interface{}) (resp interface{}, err error)
	CreateDns(request interface{}) (resp interface{}, err error)
	DeleteDns(request interface{}) (resp interface{}, err error)
	ListResourceDnsRecordSets(request interface{}) (resp interface{}, err error)
}
