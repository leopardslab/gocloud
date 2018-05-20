package alistorage

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"testing"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreatedisk(t *testing.T) {
	var alistorage Alistorage
	createDisk := map[string]interface{}{
		"RegionId":    "cn-qingdao",
		"ZoneId":      "cn-qingdao-b",
		"Size":        20,
		"DiskName":    "ThisIsDiskName",
		"Description": "ThisIsDescription",
	}
	_, err := alistorage.Createdisk(createDisk)
	if err != nil {
		t.Errorf("CreateDisk Test Fail")
		return
	}
	t.Logf("Ali disk is created successfully.")
}

func TestDeletedisk(t *testing.T) {
	var alistorage Alistorage
	deleteDisk := map[string]interface{}{
		"DiskId": "d-m5e2g66ws9m7r00ux87h",
	}
	_, err := alistorage.Deletedisk(deleteDisk)
	if err != nil {
		t.Errorf("DeleteDisk Test Fail")
		return
	}
	t.Logf("Ali disk is deleted successfully.")
}

func TestAttachdisk(t *testing.T) {
	var alistorage Alistorage
	attachDisk := map[string]interface{}{
		"InstanceId":         "i-m5e1lploaf58bddf0gah",
		"DiskId":             "d-m5edwiwlyyn7bwz6cdd4",
		"DeleteWithInstance": false,
	}
	_, err := alistorage.Attachdisk(attachDisk)
	if err != nil {
		t.Errorf("AttachDisk Test Fail")
		return
	}
	t.Logf("Ali disk is attached successfully.")
}

func TestDetachdisk(t *testing.T) {
	var alistorage Alistorage
	detachDisk := map[string]interface{}{
		"InstanceId": "i-m5e1lploaf58bddf0gah",
		"DiskId":     "d-m5edwiwlyyn7bwz6cdd4",
	}
	_, err := alistorage.Detachdisk(detachDisk)
	if err != nil {
		t.Errorf("DetachDisk Test Fail")
		return
	}
	t.Logf("Ali disk is detached successfully.")
}

func TestCreatesnapshot(t *testing.T) {
	var alistorage Alistorage
	createsnapshot := map[string]interface{}{
		"DiskId":       "d-m5e7k6ycnx8b5zzsm0yp",
		"SnapshotName": "ThisIsSnapshotName",
	}
	_, err := alistorage.Createsnapshot(createsnapshot)
	if err != nil {
		t.Errorf("CreateSnapshot Test Fail: %s", err)
		return
	}
	t.Logf("Ali disk snapshot is created successfully.")
}

func TestDeletesnapshot(t *testing.T) {
	var alistorage Alistorage
	deletesnapshot := map[string]interface{}{
		"SnapshotId": "s-m5eave3s6oufpctcxynu",
	}
	_, err := alistorage.Deletesnapshot(deletesnapshot)
	if err != nil {
		t.Errorf("DeleteSnapshot Test Fail: %s", err)
		return
	}
	t.Logf("Ali disk snapshot is deleted successfully.")
}
