package vultrdns

import (
	"encoding/json"
)

type ListDnsResp []DnsInfo

type DnsInfo struct {
	Tpye     string `json:"tpye"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	Priority int    `json:"priority"`
	RecordID int    `json:"RECORDID"`
	Ttl      int    `json:"ttl"`
}

func ParseListDnsResp(body interface{}) (listDnsResp ListDnsResp, err error) {
	err = json.Unmarshal([]byte(body.(string)), &listDnsResp)
	return
}
