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

func (vultrDNS *VultrDNS) Createdns(request interface{}) (resp interface{}, err error)
    Createdns function creates a new DNS record.

func (vultrDNS *VultrDNS) Deletedns(request interface{}) (resp interface{}, err error)
    Deletedns function deletes a DNS record.

func (vultrDNS *VultrDNS) ListResourcednsRecordSets(request interface{}) (resp interface{}, err error)
    ListResourcednsRecordSets function lists DNS record sets.

func (vultrDNS *VultrDNS) Listdns(request interface{}) (resp interface{}, err error)
    Listdns function lists DNS records.
```
