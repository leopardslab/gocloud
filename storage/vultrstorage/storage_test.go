package vultrstorage

import (
	"github.com/cloudlibz/gocloud/vultrauth"
	"testing"
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

func TestVultrStorage_Createdisk(t *testing.T) {
	var vultrStorage VultrStorage
	createDisk := map[string]interface{}{
		"DCID":    1,
		"size_gb": 50,
		"label":   "test",
	}
	resp, err := vultrStorage.Createdisk(createDisk)
	if err != nil {
		t.Errorf("Createdisk Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr disk is created successfully.")
}

func TestVultrStorage_Deletedisk(t *testing.T) {
	var vultrStorage VultrStorage
	deleteDisk := map[string]interface{}{
		"SUBID": 1313217,
	}
	resp, err := vultrStorage.Deletedisk(deleteDisk)
	if err != nil {
		t.Errorf("Deletedisk Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr disk is deleted successfully.")
}

func TestVultrStorage_Attachdisk(t *testing.T) {
	var vultrStorage VultrStorage
	attachDisk := map[string]interface{}{
		"SUBID":           1313217,
		"attach_to_SUBID": 1313207,
	}
	resp, err := vultrStorage.Attachdisk(attachDisk)
	if err != nil {
		t.Errorf("Attachdisk Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr disk is attached successfully.")
}

func TestVultrStorage_Detachdisk(t *testing.T) {
	var vultrStorage VultrStorage
	detachDisk := map[string]interface{}{
		"SUBID": 1313217,
	}
	resp, err := vultrStorage.Detachdisk(detachDisk)
	if err != nil {
		t.Errorf("Detachdisk Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr disk is detached successfully.")
}
