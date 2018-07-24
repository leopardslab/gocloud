package alidns

import "encoding/json"

type ListDnsResp struct {
	DomainRecords struct {
		Record []RecordInfo
	}
}

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

func ParseListDnsResp(body interface{}) (records []RecordInfo, err error) {
	listDnsResp := &ListDnsResp{}
	err = json.Unmarshal([]byte(body.(string)), &listDnsResp)
	records = listDnsResp.DomainRecords.Record
	return
}
