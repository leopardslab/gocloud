package alidns

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"testing"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreatedns(t *testing.T) {
	var aliDNS Alidns
	createDNS := map[string]interface{}{
		"DomainName": "oddcn.cn",
		"RR":         "gocloud.test",
		"Type":       "A",
		"Value":      "202.106.0.20",
		"TTL":        600,
	}
	_, err := aliDNS.Createdns(createDNS)
	if err != nil {
		t.Errorf("Createdns Test Fail: %s", err)
		return
	}
	t.Logf("Ali DNS is created successfully.")
}

func TestDeletedns(t *testing.T) {
	var aliDNS Alidns
	deleteDNS := map[string]interface{}{
		"RecordId": "9999985",
	}
	_, err := aliDNS.Deletedns(deleteDNS)
	if err != nil {
		t.Errorf("Deletedns Test Fail: %s", err)
		return
	}
	t.Logf("Ali DNS is deleted successfully.")
}

func TestListdns(t *testing.T) {
	var aliDNS Alidns
	listDomain := map[string]interface{}{
		"PageNumber": 1,
		"PageSize":   20,
		"KeyWord":    "com",
		"GroupId":    "2223",
	}
	_, err := aliDNS.Listdns(listDomain)
	if err != nil {
		t.Errorf("Listdns Test Fail: %s", err)
		return
	}
	t.Logf("Ali DNS is listed successfully.")
}
func TestListResourcednsRecordSets(t *testing.T) {
	var aliDNS Alidns
	listResourcednsRecordSets := map[string]interface{}{
		"DomainName":   "oddcn.cn",
		"PageNumber":   1,
		"PageSize":     20,
		"RRKeyWord":    "www",
		"TypeKeyWord":  "MX",
		"ValueKeyWord": "com",
	}
	_, err := aliDNS.ListResourcednsRecordSets(listResourcednsRecordSets)
	if err != nil {
		t.Errorf("ListResourcednsRecordSets Test Fail: %s", err)
		return
	}
	t.Logf("Ali resource DNS record sets is listed successfully.")
}
