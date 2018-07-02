package vultrdns

import (
	"fmt"
	"github.com/cloudlibz/gocloud/vultrauth"
	"net/http"
)

// ListResourceDnsRecordSets function lists DNS record sets.
func (vultrDNS *VultrDNS) ListResourceDnsRecordSets(request interface{}) (resp interface{}, err error) {
	fmt.Println("\nThis API is not provided by Vultr cloud")
	return resp, err
}

// ListDns function lists DNS records.
func (vultrDNS *VultrDNS) ListDns(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	var domain string

	for key, value := range param {
		switch key {
		case "domain":
			domain = value.(string)
		}
	}

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodGet, "/v1/dns/records?domain="+domain, param, response)

	resp = response
	return resp, err
}

// DeleteDns function deletes a DNS record.
func (vultrDNS *VultrDNS) DeleteDns(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/dns/delete_record", param, response)

	resp = response
	return resp, err
}

// CreateDns function creates a new DNS record.
func (vultrDNS *VultrDNS) CreateDns(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/dns/create_record", param, response)

	resp = response
	return resp, err
}
