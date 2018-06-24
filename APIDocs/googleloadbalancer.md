package googleloadbalancer
    import "github.com/cloudlibz/gocloud/loadbalancer/googleloadbalancer"


CONSTANTS

const (
    //UnixDate reperesnts unix date-time format.
    UnixDate = "Mon Jan _2 15:04:05 MST 2006"
    //RFC3339 reperesnts RFC3339 date-time format.
    RFC3339 = "2006-01-02T15:04:05Z07:00"
)

FUNCTIONS

func CreateLoadBalancerdictnoaryconvert(option TargetPools, CreateLoadBalancerjsonmap map[string]interface{})
    CreateLoadBalancerdictnoaryconvert creates a dictnoary for
    CreateLoadBalancer api.

TYPES

type Googleloadbalancer struct {
}
    Googleloadbalancer reperents google loadbalancer methods and attributes.

func (googleloadbalancer *Googleloadbalancer) AttachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error)
    AttachNodeWithLoadBalancer attach new google compute instance to google
    loadbalancer pool.

func (googleloadbalancer *Googleloadbalancer) CreateLoadBalancer(request interface{}) (resp interface{}, err error)
    CreateLoadBalancer creates google loadbalancer pool.

func (googleloadbalancer *Googleloadbalancer) DeleteLoadBalancer(request interface{}) (resp interface{}, err error)
    DeleteLoadBalancer deletes google loadbalancer pool.

func (googleloadbalancer *Googleloadbalancer) DetachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error)
    DetachNodeWithLoadBalancer Detach google compute instance from google
    loadbalancer pool.

func (googleloadbalancer *Googleloadbalancer) ListLoadBalancer(request interface{}) (resp interface{}, err error)
    ListLoadBalancer lists google loadbalancer pool.

type TargetPools struct {
    BackupPool        string   `json:"backupPool"`
    CreationTimestamp string   `json:"creationTimestamp"`
    Description       string   `json:"description"`
    HealthChecks      []string `json:"healthChecks"`
    FailoverRatio     int      `json:"failoverRatio"`
    ID                string   `json:"id"`
    Instances         []string `json:"instances"`
    Kind              string   `json:"kind"`
    Name              string   `json:"name"`
    Region            string   `json:"region"`
    SelfLink          string   `json:"selfLink"`
    SessionAffinity   string   `json:"sessionAffinity"`
}
    TargetPools reperents google loadbalancer.


