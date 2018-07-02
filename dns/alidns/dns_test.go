package alidns

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"testing"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreateDns(t *testing.T) {
	var aliDNS Alidns
	createDNS := map[string]interface{}{
		"DomainName": "oddcn.cn",
		"RR":         "gocloud.test",
		"Type":       "A",
		"Value":      "202.106.0.20",
		"TTL":        600,
	}
	_, err := aliDNS.CreateDns(createDNS)
	if err != nil {
		t.Errorf("CreateDns Test Fail: %s", err)
		return
	}
	t.Logf("Ali DNS is created successfully.")
}

func TestDeleteDns(t *testing.T) {
	var aliDNS Alidns
	deleteDNS := map[string]interface{}{
		"RecordId": "9999985",
	}
	_, err := aliDNS.DeleteDns(deleteDNS)
	if err != nil {
		t.Errorf("DeleteDns Test Fail: %s", err)
		return
	}
	t.Logf("Ali DNS is deleted successfully.")
}

func TestListDns(t *testing.T) {
	var aliDNS Alidns
	listDomain := map[string]interface{}{
		"PageNumber": 1,
		"PageSize":   20,
		"KeyWord":    "com",
		"GroupId":    "2223",
	}
	_, err := aliDNS.ListDns(listDomain)
	if err != nil {
		t.Errorf("ListDns Test Fail: %s", err)
		return
	}
	t.Logf("Ali DNS is listed successfully.")
}
func TestListResourceDnsRecordSets(t *testing.T) {
	var aliDNS Alidns
	listResourcednsRecordSets := map[string]interface{}{
		"DomainName":   "oddcn.cn",
		"PageNumber":   1,
		"PageSize":     20,
		"RRKeyWord":    "www",
		"TypeKeyWord":  "MX",
		"ValueKeyWord": "com",
	}
	_, err := aliDNS.ListResourceDnsRecordSets(listResourcednsRecordSets)
	if err != nil {
		t.Errorf("ListResourceDnsRecordSets Test Fail: %s", err)
		return
	}
	t.Logf("Ali resource DNS record sets is listed successfully.")
}
