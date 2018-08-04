package alidns

import "encoding/json"

type ListDnsResp struct {
	StatusCode int
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

func ParseListDnsResp(resp interface{}) (listDnsResp ListDnsResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &listDnsResp)
	listDnsResp.StatusCode = response["status"].(int)
	return
}
