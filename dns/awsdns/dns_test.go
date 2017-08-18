package awsdns

import "testing"
import "github.com/scorelab/gocloud-v2/auth"

func init() {
	auth.LoadConfig()
}

func TestCreatedns(t *testing.T) {
	var awsdns Awsdns
	createdns := map[string]interface{}{
		"name":             "rootmonk.me",
		"hostedZoneConfig": "hostedZoneConfig",
	}

	awsdns.Createdns(createdns)
}

func TestDeletedns(t *testing.T) {
	var awsdns Awsdns
	deletedns := map[string]string{
		"ID": "ZOD7SUP0ZGGQQ",
	}

	_, err := awsdns.Deletedns(deletedns)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListdns(t *testing.T) {
	var awsdns Awsdns

	listdns := map[string]interface{}{
		"marker":   "",
		"maxItems": 2,
	}

	_, err := awsdns.Listdns(listdns)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListResourcednsRecordSets(t *testing.T) {
	var awsdns Awsdns
	listResourcednsRecordSets := map[string]interface{}{
		"zone": "ZBNX5TIW033J2",
	}

	_, err := awsdns.ListResourcednsRecordSets(listResourcednsRecordSets)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
