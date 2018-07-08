package vultrdns

import (
	"errors"
	"github.com/cloudlibz/gocloud/vultrauth"
)

type VultrDNS struct {
}

type CreateDNS struct {
	domain   string // Domain name to add record to
	name     string // Name (subdomain) of record
	Type     string // Type (A, AAAA, MX, etc) of record
	data     string // Data for this record
	ttl      int    // (optional) TTL of this record
	priority int    // (only required for MX and SRV) Priority of this record (omit the priority from the data)
}

type ListDNS struct {
	domain string // Domain to list records for
}

type DeleteDNS struct {
	domain   string // Domain name to delete record from
	RECORDID int    // ID of record to delete (see /dns/records)
}

// CreateDNS builder pattern code
type CreateDNSBuilder struct {
	createDNS *CreateDNS
}

func NewCreateDNSBuilder() *CreateDNSBuilder {
	createDNS := &CreateDNS{}
	b := &CreateDNSBuilder{createDNS: createDNS}
	return b
}

func (b *CreateDNSBuilder) Domain(domain string) *CreateDNSBuilder {
	b.createDNS.domain = domain
	return b
}

func (b *CreateDNSBuilder) Name(name string) *CreateDNSBuilder {
	b.createDNS.name = name
	return b
}

func (b *CreateDNSBuilder) Type(typeV string) *CreateDNSBuilder {
	b.createDNS.Type = typeV
	return b
}

func (b *CreateDNSBuilder) Data(data string) *CreateDNSBuilder {
	b.createDNS.data = data
	return b
}

func (b *CreateDNSBuilder) Ttl(ttl int) *CreateDNSBuilder {
	b.createDNS.ttl = ttl
	return b
}

func (b *CreateDNSBuilder) Priority(priority int) *CreateDNSBuilder {
	b.createDNS.priority = priority
	return b
}

func (b *CreateDNSBuilder) Build() (map[string]interface{}, error) {
	if b.createDNS.domain == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "domain")
	}
	if b.createDNS.name == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "name")
	}
	if b.createDNS.Type == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "Type")
	}
	if b.createDNS.data == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "data")
	}

	params := make(map[string]interface{})
	params["domain"] = b.createDNS.domain
	params["name"] = b.createDNS.name
	params["Type"] = b.createDNS.Type
	params["data"] = b.createDNS.data

	if b.createDNS.ttl != 0 {
		params["ttl"] = b.createDNS.ttl
	}
	if b.createDNS.priority != 0 {
		params["priority"] = b.createDNS.priority
	}

	return params, nil
}

// ListDNS builder pattern code
type ListDNSBuilder struct {
	listDNS *ListDNS
}

func NewListDNSBuilder() *ListDNSBuilder {
	listDNS := &ListDNS{}
	b := &ListDNSBuilder{listDNS: listDNS}
	return b
}

func (b *ListDNSBuilder) Domain(domain string) *ListDNSBuilder {
	b.listDNS.domain = domain
	return b
}

func (b *ListDNSBuilder) Build() (map[string]interface{}, error) {
	if b.listDNS.domain == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "domain")
	}
	params := map[string]interface{}{
		"domain": b.listDNS.domain,
	}
	return params, nil
}

// DeleteDNS builder pattern code
type DeleteDNSBuilder struct {
	deleteDNS *DeleteDNS
}

func NewDeleteDNSBuilder() *DeleteDNSBuilder {
	deleteDNS := &DeleteDNS{}
	b := &DeleteDNSBuilder{deleteDNS: deleteDNS}
	return b
}

func (b *DeleteDNSBuilder) Domain(domain string) *DeleteDNSBuilder {
	b.deleteDNS.domain = domain
	return b
}

func (b *DeleteDNSBuilder) RECORDID(rECORDID int) *DeleteDNSBuilder {
	b.deleteDNS.RECORDID = rECORDID
	return b
}

func (b *DeleteDNSBuilder) Build() (map[string]interface{}, error) {
	if b.deleteDNS.domain == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "domain")
	}
	if b.deleteDNS.RECORDID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "RECORDID")
	}
	params := map[string]interface{}{
		"domain":   b.deleteDNS.domain,
		"RECORDID": b.deleteDNS.RECORDID,
	}
	return params, nil
}
