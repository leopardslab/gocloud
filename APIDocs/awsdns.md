package awsdns
    import "github.com/shlokgilda/gocloud/dns/awsdns"


FUNCTIONS

func CleanZoneID(ID string) string
    CleanZoneID cleans zone ID.

TYPES

type AliasTarget struct {
    HostedZoneId         string
    DNSName              string
    EvaluateTargetHealth bool
}
    AliasTarget represents AliasTarget.

type Awsdns struct {
}
    Awsdns represents Awsdns attribute and method associates with it.

func (awsdns *Awsdns) Createdns(request interface{}) (resp interface{}, err error)
    Createdns creates awsdns.

func (awsdns *Awsdns) Deletedns(request interface{}) (resp interface{}, err error)

func (awsdns *Awsdns) ListResourcednsRecordSets(request interface{}) (resp interface{}, err error)
    ListResourcednsRecordSets represents awsdns ListResourcednsRecordSets.

func (awsdns *Awsdns) Listdns(request interface{}) (resp interface{}, err error)
    Listdns Listdns awsdns.

func (awsdns *Awsdns) PrepareSignatureV4query(method, path string, req, resp interface{}, response map[string]interface{}) error
    PrepareSignatureV4query prepare signatue for awsdns.

type Change struct {
    Action string            `xml:"Action"`
    Record ResourceRecordSet `xml:"ResourceRecordSet"`
}
    Change represents Change.

type ChangeInfo struct {
    ID          string `xml:"Id"`
    Status      string `xml:"Status"`
    SubmittedAt string `xml:"SubmittedAt"`
}
    ChangeInfo represents ChangeInfo.

type ChangeResourceRecordSetsRequest struct {
    Comment string   `xml:"ChangeBatch>Comment,omitempty"`
    Changes []Change `xml:"ChangeBatch>Changes>Change"`
}
    ChangeResourceRecordSetsRequest respeptes
    ChangeResourceRecordSetsRequest.

type ChangeResourceRecordSetsResponse struct {
    ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}
    ChangeResourceRecordSetsResponse represents
    ChangeResourceRecordSetsResponse

type CreateHostedZoneRequest struct {
    Name            string `xml:"Name"`
    CallerReference string `xml:"CallerReference"`
    Comment         string `xml:"HostedZoneConfig>Comment"`
}
    CreateHostedZoneRequest represents CreateHostedZoneRequest.

type CreateHostedZoneResponse struct {
    HostedZone    HostedZone    `xml:"HostedZone"`
    ChangeInfo    ChangeInfo    `xml:"ChangeInfo"`
    DelegationSet DelegationSet `xml:"DelegationSet"`
}
    CreateHostedZoneResponse represents CreateHostedZoneResponse.

type DelegationSet struct {
    NameServers []string `xml:"NameServers>NameServer"`
}
    DelegationSet represents DelegationSet.

type DeleteHostedZoneResponse struct {
    ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}
    DeleteHostedZoneResponse represents DeleteHostedZoneResponse.

type HostedZone struct {
    ID              string `xml:"Id"`
    Name            string `xml:"Name"`
    CallerReference string `xml:"CallerReference"`
    Comment         string `xml:"Config>Comment"`
    ResourceCount   int    `xml:"ResourceRecordSetCount"`
}
    HostedZone represents HostedZone.

type ListHostedZonesResponse struct {
    HostedZones []HostedZone `xml:"HostedZones>HostedZone"`
    Marker      string       `xml:"Marker"`
    IsTruncated bool         `xml:"IsTruncated"`
    NextMarker  string       `xml:"NextMarker"`
    MaxItems    int          `xml:"MaxItems"`
}
    ListHostedZonesResponse represents ListHostedZonesResponse.

type ListResourceRecordSetsResponse struct {
    Records              []ResourceRecordSet `xml:"ResourceRecordSets>ResourceRecordSet"`
    IsTruncated          bool                `xml:"IsTruncated"`
    MaxItems             int                 `xml:"MaxItems"`
    NextRecordName       string              `xml:"NextRecordName"`
    NextRecordType       string              `xml:"NextRecordType"`
    NextRecordIdentifier string              `xml:"NextRecordIdentifier"`
}
    ListResourceRecordSetsResponse represents
    ListResourceRecordSetsResponse.

type ListResourcednsRecordSets struct {
    Name       string
    Type       string
    Identifier string
    MaxItems   int
}
    ListResourcednsRecordSets represents ListResourcednsRecordSets.

type ResourceRecordSet struct {
    Name          string       `xml:"Name"`
    Type          string       `xml:"Type"`
    TTL           int          `xml:"TTL"`
    Records       []string     `xml:"ResourceRecords>ResourceRecord>Value,omitempty"`
    SetIdentifier string       `xml:"SetIdentifier,omitempty"`
    Weight        int          `xml:"Weight,omitempty"`
    HealthCheckId string       `xml:"HealthCheckId,omitempty"`
    Region        string       `xml:"Region,omitempty"`
    Failover      string       `xml:"Failover,omitempty"`
    AliasTarget   *AliasTarget `xml:"AliasTarget,omitempty"`
    RecordsXML    string       `xml:",innerxml"`
}
    ResourceRecordSet represents ResourceRecordSet.


