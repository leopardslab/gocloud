package googleloadbalancer
    import "github.com/shlokgilda/gocloud/loadbalancer/googleloadbalancer"


CONSTANTS

const (
    //UnixDate reperesnts unix date-time format.
    UnixDate = "Mon Jan _2 15:04:05 MST 2006"
    //RFC3339 reperesnts RFC3339 date-time format.
    RFC3339 = "2006-01-02T15:04:05Z07:00"
)

FUNCTIONS

func Creatloadbalancerdictnoaryconvert(option TargetPools, Creatloadbalancerjsonmap map[string]interface{})
    Creatloadbalancerdictnoaryconvert creates a dictnoary for
    Creatloadbalancer api.

TYPES

type Googleloadbalancer struct {
}
    Googleloadbalancer reperents google loadbalancer methods and attributes.

func (googleloadbalancer *Googleloadbalancer) Attachnodewithloadbalancer(request interface{}) (resp interface{}, err error)
    Attachnodewithloadbalancer attach new google compute instance to google
    loadbalancer pool.

func (googleloadbalancer *Googleloadbalancer) Creatloadbalancer(request interface{}) (resp interface{}, err error)
    Creatloadbalancer creates google loadbalancer pool.

func (googleloadbalancer *Googleloadbalancer) Deleteloadbalancer(request interface{}) (resp interface{}, err error)
    Deleteloadbalancer deletes google loadbalancer pool.

func (googleloadbalancer *Googleloadbalancer) Detachnodewithloadbalancer(request interface{}) (resp interface{}, err error)
    Detachnodewithloadbalancer Detach google compute instance from google
    loadbalancer pool.

func (googleloadbalancer *Googleloadbalancer) Listloadbalancer(request interface{}) (resp interface{}, err error)
    Listloadbalancer lists google loadbalancer pool.

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


