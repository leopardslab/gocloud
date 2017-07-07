package awsloadbalancer

const (
	loadbalancerversion = "2012-06-01"
)

type Awsloadbalancer struct {
}

type CreateLoadBalancer struct {
	LoadBalancerName   string
	AvailabilityZones []string
	Scheme            string
	Tags              []Tag
	SecurityGroups    []string
	Subnets           []string
	Listeners         []Listener
}

type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

type Listener struct {
	InstancePort     string
	InstanceProtocol string
	LoadBalancerPort string
	Protocol         string
	SSLCertificateId string
}
