package vultrstorage

import (
	"testing"
	"github.com/cloudlibz/gocloud/vultrauth"
)

func init() {
	vultrauth.LoadConfig()
}

func TestVultrStorage_Createsnapshot(t *testing.T) {
	var vultrStorage VultrStorage
	create := map[string]interface{}{
		"SUBID": 1312965,
	}
	resp, err := vultrStorage.Createsnapshot(create)
	if err != nil {
		t.Errorf("Createsnapshot Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr snapshot is created successfully.")
}

func TestVultrStorage_Deletesnapshot(t *testing.T) {
	var vultrStorage VultrStorage
	deleteSnapshot := map[string]interface{}{
		"SNAPSHOTID": "5359435d28b9a",
	}
	resp, err := vultrStorage.Deletesnapshot(deleteSnapshot)
	if err != nil {
		t.Errorf("Deletesnapshot Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr snapshot is deleted successfully.")
}
