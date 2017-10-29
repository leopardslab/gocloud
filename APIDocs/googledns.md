package googledns
    import "github.com/cloudlibz/gocloud/dns/googledns"


CONSTANTS

const (
    UnixDate = "Mon Jan _2 15:04:05 MST 2006"
    RFC3339  = "2006-01-02T15:04:05Z07:00"
)

FUNCTIONS

func Creatednsedictnoaryconvert(option Createdns, Creatednsjsonmap map[string]interface{})
    Creatednsedictnoaryconvert convert Createdns parameters to Createdns
    dictnoary.

TYPES

type Createdns struct {
    CreationTime  string   `json:"creationTime"`
    Description   string   `json:"description"`
    DNSName       string   `json:"dnsName"`
    ID            string   `json:"id"`
    Kind          string   `json:"kind"`
    Name          string   `json:"name"`
    NameServerSet string   `json:"nameServerSet"`
    NameServers   []string `json:"nameServers"`
}
    Createdns struct represents Createdns attribute.

type Googledns struct {
}
    Googledns struct represents Googledns attribute and methods associates
    with it.

func (googledns *Googledns) Createdns(request interface{}) (resp interface{}, err error)
    Createdns creates DNS.

func (googledns *Googledns) Deletedns(request interface{}) (resp interface{}, err error)
    Deletedns deletes DNS.

func (googledns *Googledns) ListResourcednsRecordSets(request interface{}) (resp interface{}, err error)
    ListResourcednsRecordSets list ListResourcednsRecordSets.

func (googledns *Googledns) Listdns(request interface{}) (resp interface{}, err error)
    Listdns lists DNS.


