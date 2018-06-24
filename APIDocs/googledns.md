package googledns
    import "github.com/cloudlibz/gocloud/dns/googledns"


CONSTANTS

const (
    UnixDate = "Mon Jan _2 15:04:05 MST 2006"
    RFC3339  = "2006-01-02T15:04:05Z07:00"
)

FUNCTIONS

func CreateDnsedictnoaryconvert(option CreateDns, CreateDnsjsonmap map[string]interface{})
    CreateDnsedictnoaryconvert convert CreateDns parameters to CreateDns
    dictnoary.

TYPES

type CreateDns struct {
    CreationTime  string   `json:"creationTime"`
    Description   string   `json:"description"`
    DNSName       string   `json:"dnsName"`
    ID            string   `json:"id"`
    Kind          string   `json:"kind"`
    Name          string   `json:"name"`
    NameServerSet string   `json:"nameServerSet"`
    NameServers   []string `json:"nameServers"`
}
    CreateDns struct represents CreateDns attribute.

type Googledns struct {
}
    Googledns struct represents Googledns attribute and methods associates
    with it.

func (googledns *Googledns) CreateDns(request interface{}) (resp interface{}, err error)
    CreateDns creates DNS.

func (googledns *Googledns) DeleteDns(request interface{}) (resp interface{}, err error)
    DeleteDns deletes DNS.

func (googledns *Googledns) ListResourceDnsRecordSets(request interface{}) (resp interface{}, err error)
    ListResourceDnsRecordSets list ListResourceDnsRecordSets.

func (googledns *Googledns) ListDns(request interface{}) (resp interface{}, err error)
    ListDns lists DNS.


