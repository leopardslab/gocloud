package vultrdns

import (
	"encoding/json"
)

type ListDnsResp struct {
	StatusCode int
	DnsSlice   []DnsInfo
}

type DnsInfo struct {
	Tpye     string `json:"tpye"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	Priority int    `json:"priority"`
	RecordID int    `json:"RECORDID"`
	Ttl      int    `json:"ttl"`
}

func ParseListDnsResp(resp interface{}) (listDnsResp ListDnsResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &listDnsResp.DnsSlice)
	listDnsResp.StatusCode = response["status"].(int)
	return
}
