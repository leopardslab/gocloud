package vultrdns

import (
	"github.com/cloudlibz/gocloud/vultrauth"
	"net/http"
)

// ListResourcednsRecordSets function lists DNS record sets.
func (vultrDNS *VultrDNS) ListResourcednsRecordSets(request interface{}) (resp interface{}, err error) {
	return resp, err
}

// Listdns function lists DNS records.
func (vultrDNS *VultrDNS) Listdns(request interface{}) (resp interface{}, err error) {
	return resp, err
}

// Deletedns function deletes a DNS record.
func (vultrDNS *VultrDNS) Deletedns(request interface{}) (resp interface{}, err error) {
	return resp, err
}

// Createdns function creates a new DNS record.
func (vultrDNS *VultrDNS) Createdns(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/dns/create_record", param, response)

	resp = response
	return resp, err
}
