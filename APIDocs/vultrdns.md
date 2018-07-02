```
package vultrdns
    import "github.com/cloudlibz/gocloud/dns/vultrdns"
```

TYPES

```
type CreateDNS struct {
    Type string // Type (A, AAAA, MX, etc) of record
    // contains filtered or unexported fields
}

type DeleteDNS struct {
    RECORDID int // ID of record to delete (see /dns/records)
    // contains filtered or unexported fields
}

type VultrDNS struct {
}

func (vultrDNS *VultrDNS) CreateDns(request interface{}) (resp interface{}, err error)
    CreateDns function creates a new DNS record.

func (vultrDNS *VultrDNS) DeleteDns(request interface{}) (resp interface{}, err error)
    DeleteDns function deletes a DNS record.

func (vultrDNS *VultrDNS) ListResourceDnsRecordSets(request interface{}) (resp interface{}, err error)
    ListResourceDnsRecordSets function lists DNS record sets.

func (vultrDNS *VultrDNS) ListDns(request interface{}) (resp interface{}, err error)
    ListDns function lists DNS records.
```
