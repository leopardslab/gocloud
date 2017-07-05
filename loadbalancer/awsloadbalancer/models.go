package awsloadbalancer

const (
	loadbalancerversion = "2015-12-01"
)

type Awsloadbalancer struct {
}

type CreateLoadBalancer struct {
	Name           string
	IpAddressType  string
	Scheme         string
	Tags           []Tag
	SecurityGroups []string
	Subnets        []string
}

type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}
