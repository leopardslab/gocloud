```
package alidns
    import "github.com/cloudlibz/gocloud/dns/alidns"
```

### TYPES

```
type Alidns struct {
}
    Alidns represents Alidns attribute and method associates with it.

func (alidns *Alidns) CreateDns(request interface{}) (resp interface{}, err error)
    CreateDns add DNS record accept map[string]interface{}

func (alidns *Alidns) DeleteDns(request interface{}) (resp interface{}, err error)
    DeleteDns delete DNS record accept map[string]interface{}

func (alidns *Alidns) ListDns(request interface{}) (resp interface{}, err error)
    ListDns list DNS record accept map[string]interface{}

func (alidns *Alidns) ListResourceDnsRecordSets(request interface{}) (resp interface{}, err error)
    ListResourceDnsRecordSets list resource DNS record sets accept
    map[string]interface{}

type CreateDNS struct {
    DomainName string
    RR         string
    Type       string
    Value      string
    TTL        int
    Priority   int
    Line       string
}
    CreateDNS to store all attribute to create Ali-cloud DNS

type CreateDNSBuilder struct {
    // contains filtered or unexported fields
}
    CreateDNS builder pattern code

func NewCreateDNSBuilder() *CreateDNSBuilder

func (b *CreateDNSBuilder) Build() (map[string]interface{}, error)

func (b *CreateDNSBuilder) DomainName(domainName string) *CreateDNSBuilder

func (b *CreateDNSBuilder) Line(line string) *CreateDNSBuilder

func (b *CreateDNSBuilder) Priority(priority int) *CreateDNSBuilder

func (b *CreateDNSBuilder) RR(rR string) *CreateDNSBuilder

func (b *CreateDNSBuilder) TTL(tTL int) *CreateDNSBuilder

func (b *CreateDNSBuilder) Type(typeV string) *CreateDNSBuilder

func (b *CreateDNSBuilder) Value(value string) *CreateDNSBuilder

type DeleteDNS struct {
    RecordId string
}
    DeleteDNS to store all attribute to delete Ali-cloud DNS

type DeleteDNSBuilder struct {
    // contains filtered or unexported fields
}
    DeleteDNS builder pattern code

func NewDeleteDNSBuilder() *DeleteDNSBuilder

func (b *DeleteDNSBuilder) Build() (map[string]interface{}, error)

func (b *DeleteDNSBuilder) RecordId(recordId string) *DeleteDNSBuilder

type ListDNS struct {
    DomainName   string
    PageNumber   int
    PageSize     int
    RRKeyWord    string
    TypeKeyWord  string
    ValueKeyWord string
}
    ListResourceDNSRecordSets to store all attribute to list resource
    Ali-cloud DNS record sets

type ListDNSBuilder struct {
    // contains filtered or unexported fields
}
    ListDNS builder pattern code

func NewListDNSBuilder() *ListDNSBuilder

func (b *ListDNSBuilder) Build() (map[string]interface{}, error)

func (b *ListDNSBuilder) DomainName(domainName string) *ListDNSBuilder

func (b *ListDNSBuilder) PageNumber(pageNumber int) *ListDNSBuilder

func (b *ListDNSBuilder) PageSize(pageSize int) *ListDNSBuilder

func (b *ListDNSBuilder) RRKeyWord(rRKeyWord string) *ListDNSBuilder

func (b *ListDNSBuilder) TypeKeyWord(typeKeyWord string) *ListDNSBuilder

func (b *ListDNSBuilder) ValueKeyWord(valueKeyWord string) *ListDNSBuilder

type ListDnsResp struct {
    StatusCode    int
    DomainRecords struct {
        Record []RecordInfo
    }
}

func ParseListDnsResp(resp interface{}) (listDnsResp ListDnsResp, err error)

type RecordInfo struct {
    DomainName string
    RecordId   string
    RR         string
    Type       string
    Value      string
    Line       string
    Priority   int
    TTL        int
    Status     string
    Locked     bool
}
```
