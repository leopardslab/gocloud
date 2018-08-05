package alidns

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"errors"
)

//Alidns represents Alidns attribute and method associates with it.
type Alidns struct {
}

// CreateDNS to store all attribute to create Ali-cloud DNS
type CreateDNS struct {
	DomainName string
	RR         string
	Type       string
	Value      string
	TTL        int
	Priority   int
	Line       string
}

// DeleteDNS to store all attribute to delete Ali-cloud DNS
type DeleteDNS struct {
	RecordId string
}

// ListResourceDNSRecordSets to store all attribute to list resource Ali-cloud DNS record sets
type ListDNS struct {
	DomainName   string
	PageNumber   int
	PageSize     int
	RRKeyWord    string
	TypeKeyWord  string
	ValueKeyWord string
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

func (b *CreateDNSBuilder) DomainName(domainName string) *CreateDNSBuilder {
	b.createDNS.DomainName = domainName
	return b
}

func (b *CreateDNSBuilder) RR(rR string) *CreateDNSBuilder {
	b.createDNS.RR = rR
	return b
}

func (b *CreateDNSBuilder) Type(typeV string) *CreateDNSBuilder {
	b.createDNS.Type = typeV
	return b
}

func (b *CreateDNSBuilder) Value(value string) *CreateDNSBuilder {
	b.createDNS.Value = value
	return b
}

func (b *CreateDNSBuilder) TTL(tTL int) *CreateDNSBuilder {
	b.createDNS.TTL = tTL
	return b
}

func (b *CreateDNSBuilder) Priority(priority int) *CreateDNSBuilder {
	b.createDNS.Priority = priority
	return b
}

func (b *CreateDNSBuilder) Line(line string) *CreateDNSBuilder {
	b.createDNS.Line = line
	return b
}

func (b *CreateDNSBuilder) Build() (map[string]interface{}, error) {
	if b.createDNS.DomainName == "" {
		return nil, errors.New(aliauth.StrMissRequired + "DomainName")
	}
	if b.createDNS.RR == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RR")
	}
	if b.createDNS.Type == "" {
		return nil, errors.New(aliauth.StrMissRequired + "Type")
	}
	if b.createDNS.Value == "" {
		return nil, errors.New(aliauth.StrMissRequired + "Value")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.createDNS)
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

func (b *DeleteDNSBuilder) RecordId(recordId string) *DeleteDNSBuilder {
	b.deleteDNS.RecordId = recordId
	return b
}

func (b *DeleteDNSBuilder) Build() (map[string]interface{}, error) {
	if b.deleteDNS.RecordId == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RecordId")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.deleteDNS)
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

func (b *ListDNSBuilder) DomainName(domainName string) *ListDNSBuilder {
	b.listDNS.DomainName = domainName
	return b
}

func (b *ListDNSBuilder) PageNumber(pageNumber int) *ListDNSBuilder {
	b.listDNS.PageNumber = pageNumber
	return b
}

func (b *ListDNSBuilder) PageSize(pageSize int) *ListDNSBuilder {
	b.listDNS.PageSize = pageSize
	return b
}

func (b *ListDNSBuilder) RRKeyWord(rRKeyWord string) *ListDNSBuilder {
	b.listDNS.RRKeyWord = rRKeyWord
	return b
}

func (b *ListDNSBuilder) TypeKeyWord(typeKeyWord string) *ListDNSBuilder {
	b.listDNS.TypeKeyWord = typeKeyWord
	return b
}

func (b *ListDNSBuilder) ValueKeyWord(valueKeyWord string) *ListDNSBuilder {
	b.listDNS.ValueKeyWord = valueKeyWord
	return b
}

func (b *ListDNSBuilder) Build() (map[string]interface{}, error) {
	if b.listDNS.DomainName == "" {
		return nil, errors.New(aliauth.StrMissRequired + "DomainName")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.listDNS)
	return params, nil
}
