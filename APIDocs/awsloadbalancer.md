package awsloadbalancer
    import "github.com/cloudlibz/gocloud/loadbalancer/awsloadbalancer"


TYPES

type Attachnodewithloadbalancer struct {
    Instances        []string
    LoadBalancerName string
}
    Attachnodewithloadbalancer represents Attachnodewithloadbalancer
    attribute.

type Awsloadbalancer struct {
}
    Awsloadbalancer represents Awsloadbalancer struct.

func (awsloadbalancer *Awsloadbalancer) Attachnodewithloadbalancer(request interface{}) (resp interface{}, err error)
    Attachnodewithloadbalancer method Attach node with loadbalancer.

func (awsloadbalancer *Awsloadbalancer) Createloadbalancer(request interface{}) (resp interface{}, err error)
    Createloadbalancer creates classic loadbalancer.

func (awsloadbalancer *Awsloadbalancer) Deleteloadbalancer(request interface{}) (resp interface{}, err error)
    Deleteloadbalancer Delete loadbalancer accept LoadBalancerName.

func (awsloadbalancer *Awsloadbalancer) Detachnodewithloadbalancer(request interface{}) (resp interface{}, err error)
    Detachnodewithloadbalancer Detach node with loadbalancer.

func (awsloadbalancer *Awsloadbalancer) Listloadbalancer(request interface{}) (resp interface{}, err error)
    Listloadbalancer List running loadbalancer.

func (awsloadbalancer *Awsloadbalancer) PrepareSignatureV2query(params map[string]string, response map[string]interface{}) error
    PrepareSignatureV2query for elasticloadbalancing.

type CreateLoadBalancer struct {
    LoadBalancerName  string
    AvailabilityZones []string
    Scheme            string
    Tags              []Tag
    SecurityGroups    []string
    Subnets           []string
    Listeners         []Listener
}
    CreateLoadBalancer struct represents attribute of CreateLoadBalancer.

type Detachnodewithloadbalancer struct {
    Instances        []string
    LoadBalancerName string
}
    Detachnodewithloadbalancer represents Attachnodewithloadbalancer
    attribute.

type Listener struct {
    InstancePort     string
    InstanceProtocol string
    LoadBalancerPort string
    Protocol         string
    SSLCertificateId string
}
    Listener represents Listener attribute of Createloadbalancer.

type Tag struct {
    Key   string `xml:"key"`
    Value string `xml:"value"`
}
    Tag represents attribute of Tag fields.


