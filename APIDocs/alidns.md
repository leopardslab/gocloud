```
package alidns
    import "github.com/cloudlibz/gocloud/dns/alidns"
```

### TYPES

```
type Alidns struct {
}
```
Alidns represents Alidns attribute and method associates with it.

```
func (alidns *Alidns) Createdns(request interface{}) (resp interface{}, err error)
```
Createdns add DNS record accept map[string]interface{}

```
func (alidns *Alidns) Deletedns(request interface{}) (resp interface{}, err error)
```
Deletedns delete DNS record accept map[string]interface{}

```
func (alidns *Alidns) ListResourcednsRecordSets(request interface{}) (resp interface{}, err error)
```
ListResourcednsRecordSets list resource DNS record sets accept map[string]interface{}

```
func (alidns *Alidns) Listdns(request interface{}) (resp interface{}, err error)
```
Listdns list DNS record accept map[string]interface{}

```
type CreateDNS struct {
    DomainName string
    RR         string
    Type       string
    Value      string
    TTL        int
    Priority   int
    Line       string
}
```
CreateDNS to store all attribute to create Ali-cloud DNS

```
type DeleteDNS struct {
    RecordId string
}
```
DeleteDNS to store all attribute to delete Ali-cloud DNS

```
type ListDNS struct {
    PageNumber int
    PageSize   int
    KeyWord    string
    GroupId    string
}
```
ListDNS to store all attribute to list Ali-cloud DNS

```
type ListResourceDNSRecordSets struct {
    DomainName   string
    PageNumber   int
    PageSize     int
    RRKeyWord    string
    TypeKeyWord  string
    ValueKeyWord string
}
```
ListResourceDNSRecordSets to store all attribute to list resource Ali-cloud DNS record sets