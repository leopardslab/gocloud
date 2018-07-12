package vultrstorage

import (
	"github.com/cloudlibz/gocloud/vultrauth"
	"testing"
)

func init() {
	vultrauth.LoadConfig()
}

func TestVultrStorage_CreateSnapshot(t *testing.T) {
	var vultrStorage VultrStorage
	create := map[string]interface{}{
		"SUBID": 1312965,
	}
	resp, err := vultrStorage.CreateSnapshot(create)
	if err != nil {
		t.Errorf("CreateSnapshot Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr snapshot is created successfully.")
}

func TestCreateSnapshotBuilder(t *testing.T) {
	var vultrStorage VultrStorage
	create, err := NewCreateSnapshotBuilder().
		SUBID(1312965).
		Description("desc").
		Build()
	if err != nil {
		t.Errorf("CreateSnapshot Test Fail: %s", err)
		return
	}
	resp, err := vultrStorage.CreateSnapshot(create)
	if err != nil {
		t.Errorf("CreateSnapshot Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr snapshot is created successfully.")
}

func TestVultrStorage_DeleteSnapshot(t *testing.T) {
	var vultrStorage VultrStorage
	deleteSnapshot := map[string]interface{}{
		"SNAPSHOTID": "5359435d28b9a",
	}
	resp, err := vultrStorage.DeleteSnapshot(deleteSnapshot)
	if err != nil {
		t.Errorf("DeleteSnapshot Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr snapshot is deleted successfully.")
}

func TestDeleteSnapshotBuilder(t *testing.T) {
	var vultrStorage VultrStorage
	deleteSnapshot, err := NewDeleteSnapshotBuilder().
		SNAPSHOTID("5359435d28b9a").
		Build()
	if err != nil {
		t.Errorf("DeleteSnapshot Test Fail: %s", err)
		return
	}
	resp, err := vultrStorage.DeleteSnapshot(deleteSnapshot)
	if err != nil {
		t.Errorf("DeleteSnapshot Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr snapshot is deleted successfully.")
}

func TestVultrStorage_CreateDisk(t *testing.T) {
	var vultrStorage VultrStorage
	createDisk := map[string]interface{}{
		"DCID":    1,
		"size_gb": 50,
		"label":   "test",
	}
	resp, err := vultrStorage.CreateDisk(createDisk)
	if err != nil {
		t.Errorf("CreateDisk Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr disk is created successfully.")
}

func TestCreateDiskBuilder(t *testing.T) {
	var vultrStorage VultrStorage
	createDisk, err := NewCreateDiskBuilder().
		DCID(1).
		SizeGb(50).
		Label("test").
		Build()
	if err != nil {
		t.Errorf("CreateDisk Test Fail: %s", err)
		return
	}
	resp, err := vultrStorage.CreateDisk(createDisk)
	if err != nil {
		t.Errorf("CreateDisk Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr disk is created successfully.")
}

func TestVultrStorage_DeleteDisk(t *testing.T) {
	var vultrStorage VultrStorage
	deleteDisk := map[string]interface{}{
		"SUBID": 1313217,
	}
	resp, err := vultrStorage.DeleteDisk(deleteDisk)
	if err != nil {
		t.Errorf("DeleteDisk Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr disk is deleted successfully.")
}

func TestDeleteDiskBuilder(t *testing.T) {
	var vultrStorage VultrStorage
	deleteDisk, err := NewDeleteDiskBuilder().
		SUBID(1313217).
		Build()
	if err != nil {
		t.Errorf("DeleteDisk Test Fail: %s", err)
		return
	}
	resp, err := vultrStorage.DeleteDisk(deleteDisk)
	if err != nil {
		t.Errorf("DeleteDisk Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr disk is deleted successfully.")
}

func TestVultrStorage_AttachDisk(t *testing.T) {
	var vultrStorage VultrStorage
	attachDisk := map[string]interface{}{
		"SUBID":           1313217,
		"attach_to_SUBID": 1313207,
	}
	resp, err := vultrStorage.AttachDisk(attachDisk)
	if err != nil {
		t.Errorf("AttachDisk Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr disk is attached successfully.")
}

func TestAttachDiskBuilder(t *testing.T) {
	var vultrStorage VultrStorage
	attachDisk, err := NewAttachDiskBuilder().
		SUBID(1313217).
		AttachToSUBID(1313207).
		Build()
	if err != nil {
		t.Errorf("AttachDisk Test Fail: %s", err)
		return
	}
	resp, err := vultrStorage.AttachDisk(attachDisk)
	if err != nil {
		t.Errorf("AttachDisk Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr disk is attached successfully.")
}

func TestVultrStorage_DetachDisk(t *testing.T) {
	var vultrStorage VultrStorage
	detachDisk := map[string]interface{}{
		"SUBID": 1313217,
	}
	resp, err := vultrStorage.DetachDisk(detachDisk)
	if err != nil {
		t.Errorf("DetachDisk Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr disk is detached successfully.")
}

func TestNewDetachDiskBuilder(t *testing.T) {
	var vultrStorage VultrStorage
	detachDisk, err := NewDetachDiskBuilder().
		SUBID(1313217).
		Build()
	if err != nil {
		t.Errorf("DetachDisk Test Fail: %s", err)
		return
	}
	resp, err := vultrStorage.DetachDisk(detachDisk)
	if err != nil {
		t.Errorf("DetachDisk Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr disk is detached successfully.")
}
