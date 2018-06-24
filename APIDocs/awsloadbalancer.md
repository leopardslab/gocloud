package awsloadbalancer
    import "github.com/cloudlibz/gocloud/loadbalancer/awsloadbalancer"


TYPES

type AttachNodeWithLoadBalancer struct {
    Instances        []string
    LoadBalancerName string
}
    AttachNodeWithLoadBalancer represents AttachNodeWithLoadBalancer
    attribute.

type Awsloadbalancer struct {
}
    Awsloadbalancer represents Awsloadbalancer struct.

func (awsloadbalancer *Awsloadbalancer) AttachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error)
    AttachNodeWithLoadBalancer method Attach node with loadbalancer.

func (awsloadbalancer *Awsloadbalancer) CreateLoadBalancer(request interface{}) (resp interface{}, err error)
    CreateLoadBalancer creates classic loadbalancer.

func (awsloadbalancer *Awsloadbalancer) DeleteLoadBalancer(request interface{}) (resp interface{}, err error)
    DeleteLoadBalancer Delete loadbalancer accept LoadBalancerName.

func (awsloadbalancer *Awsloadbalancer) DetachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error)
    DetachNodeWithLoadBalancer Detach node with loadbalancer.

func (awsloadbalancer *Awsloadbalancer) ListLoadBalancer(request interface{}) (resp interface{}, err error)
    ListLoadBalancer List running loadbalancer.

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

type DetachNodeWithLoadBalancer struct {
    Instances        []string
    LoadBalancerName string
}
    DetachNodeWithLoadBalancer represents AttachNodeWithLoadBalancer
    attribute.

type Listener struct {
    InstancePort     string
    InstanceProtocol string
    LoadBalancerPort string
    Protocol         string
    SSLCertificateId string
}
    Listener represents Listener attribute of CreateLoadBalancer.

type Tag struct {
    Key   string `xml:"key"`
    Value string `xml:"value"`
}
    Tag represents attribute of Tag fields.


