```
package vultrdns
    import "github.com/cloudlibz/gocloud/dns/vultrdns"
```

### TYPES

```
type CreateDNS struct {
    Type string // Type (A, AAAA, MX, etc) of record
    // contains filtered or unexported fields
}

type CreateDNSBuilder struct {
    // contains filtered or unexported fields
}
    CreateDNS builder pattern code

func NewCreateDNSBuilder() *CreateDNSBuilder

func (b *CreateDNSBuilder) Build() (map[string]interface{}, error)

func (b *CreateDNSBuilder) Data(data string) *CreateDNSBuilder

func (b *CreateDNSBuilder) Domain(domain string) *CreateDNSBuilder

func (b *CreateDNSBuilder) Name(name string) *CreateDNSBuilder

func (b *CreateDNSBuilder) Priority(priority int) *CreateDNSBuilder

func (b *CreateDNSBuilder) Ttl(ttl int) *CreateDNSBuilder

func (b *CreateDNSBuilder) Type(typeV string) *CreateDNSBuilder

type DeleteDNS struct {
    RECORDID int // ID of record to delete (see /dns/records)
    // contains filtered or unexported fields
}

type DeleteDNSBuilder struct {
    // contains filtered or unexported fields
}
    DeleteDNS builder pattern code

func NewDeleteDNSBuilder() *DeleteDNSBuilder

func (b *DeleteDNSBuilder) Build() (map[string]interface{}, error)

func (b *DeleteDNSBuilder) Domain(domain string) *DeleteDNSBuilder

func (b *DeleteDNSBuilder) RECORDID(rECORDID int) *DeleteDNSBuilder

type DnsInfo struct {
    Tpye     string `json:"tpye"`
    Name     string `json:"name"`
    Data     string `json:"data"`
    Priority int    `json:"priority"`
    RecordID int    `json:"RECORDID"`
    Ttl      int    `json:"ttl"`
}

type ListDNS struct {
    // contains filtered or unexported fields
}

type ListDNSBuilder struct {
    // contains filtered or unexported fields
}
    ListDNS builder pattern code

func NewListDNSBuilder() *ListDNSBuilder

func (b *ListDNSBuilder) Build() (map[string]interface{}, error)

func (b *ListDNSBuilder) Domain(domain string) *ListDNSBuilder

type ListDnsResp struct {
    StatusCode int
    DnsSlice   []DnsInfo
}

func ParseListDnsResp(resp interface{}) (listDnsResp ListDnsResp, err error)

type VultrDNS struct {
}

func (vultrDNS *VultrDNS) CreateDns(request interface{}) (resp interface{}, err error)
    CreateDns function creates a new DNS record.

func (vultrDNS *VultrDNS) DeleteDns(request interface{}) (resp interface{}, err error)
    DeleteDns function deletes a DNS record.

func (vultrDNS *VultrDNS) ListDns(request interface{}) (resp interface{}, err error)
    ListDns function lists DNS records.

func (vultrDNS *VultrDNS) ListResourceDnsRecordSets(request interface{}) (resp interface{}, err error)
    ListResourceDnsRecordSets function lists DNS record sets.
```

