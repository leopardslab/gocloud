package awsloadbalancer

const (
	loadbalancerversion = "2012-06-01"
)

//Awsloadbalancer represents Awsloadbalancer struct.
type Awsloadbalancer struct {
}

//CreateLoadBalancer struct represents attribute of CreateLoadBalancer.
type CreateLoadBalancer struct {
	LoadBalancerName  string
	AvailabilityZones []string
	Scheme            string
	Tags              []Tag
	SecurityGroups    []string
	Subnets           []string
	Listeners         []Listener
}

//Tag represents attribute of Tag fields.
type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

//Listener represents Listener attribute of Creatloadbalancer.
type Listener struct {
	InstancePort     string
	InstanceProtocol string
	LoadBalancerPort string
	Protocol         string
	SSLCertificateId string
}

//Attachnodewithloadbalancer represents Attachnodewithloadbalancer attribute.
type Attachnodewithloadbalancer struct {
	Instances        []string
	LoadBalancerName string
}

//Detachnodewithloadbalancer represents Attachnodewithloadbalancer attribute.
type Detachnodewithloadbalancer struct {
	Instances        []string
	LoadBalancerName string
}
