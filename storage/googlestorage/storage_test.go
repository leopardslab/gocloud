package googlestorage

import "testing"

func TestCreatedisk(t *testing.T) {
	var googlestorage GoogleStorage

	createdisk := map[string]interface{}{
		"projectid": "sheltermap-1493101612061",
		"Name":      "disk-11",
		"Type":      "pd-standard",
		"Zone":      "us-east4-c",
	}

	_, err := googlestorage.Createdisk(createdisk)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeletedisk(t *testing.T) {
	var googlestorage GoogleStorage

	deletedisk := map[string]string{
		"VolumeId": "vol-05371175f10d2e6a4",
	}

	_, err := googlestorage.Deletedisk(deletedisk)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestAttachdisk(t *testing.T) {
	var googlestorage GoogleStorage

	attachdisk := map[string]interface{}{
		"projectid":  "sheltermap-1493101612061",
		"instance":   "testing-scorelab",
		"Zone":       "us-east4-c",
		"deviceName": "projects/sheltermap-1493101612061/zones/us-east4-c/disks/disk-12",
	}

	_, err := googlestorage.Attachdisk(attachdisk)

	if err != nil {
		t.Errorf("Test Fail")
	}
}


func TestCreatesnapshot(t *testing.T) {
	var googlestorage GoogleStorage

	createsnapshot := map[string]interface{}{
		"projectid": "sheltermap-1493101612061",
		"disk":      "disk-11",
		"Zone":      "us-east4-c",
		"Name":      "disk-12",
	}
	_, err := googlestorage.Createsnapshot(createsnapshot)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeletesnapshot(t *testing.T) {
	var googlestorage GoogleStorage

	deletesnapshot := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"snapshot":  "disk-12",
	}

	_, err := googlestorage.Deletesnapshot(deletesnapshot)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
