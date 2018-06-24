package gce

import "testing"

func TestDeleteNode(t *testing.T) {
	var gce GCE
	deletenode := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"instance":  "testing-scorelab2",
		"Zone":      "us-east4-c",
	}
	gce.DeleteNode(deletenode)
}

func TestStopNode(t *testing.T) {
	var gce GCE
	stopnode := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"instance":  "testing-scorelab2",
		"Zone":      "us-east4-c",
	}

	_, err := gce.StopNode(stopnode)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestRebootNode(t *testing.T) {
	var gce GCE
	reboot := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"instance":  "testing-scorelab",
		"Zone":      "us-east4-c",
	}
	_, err := gce.RebootNode(reboot)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestCreateNode(t *testing.T) {

	var gce GCE

	InitializeParams := map[string]string{
		"SourceImage": "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301",
		"DiskType":    "projects/sheltermap-1493101612061/zones/us-east4-c/diskTypes/pd-standard",
		"DiskSizeGb":  "10",
	}

	disk := []map[string]interface{}{
		{
			"Boot":             true,
			"AutoDelete":       false,
			"DeviceName":       "bokya",
			"Type":             "PERSISTENT",
			"Mode":             "READ_WRITE",
			"InitializeParams": InitializeParams,
		},
	}

	AccessConfigs := []map[string]string{{
		"Name": "external-nat",
		"Type": "ONE_TO_ONE_NAT",
	},
	}

	NetworkInterfaces := []map[string]interface{}{
		{
			"Network":       "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
			"Subnetwork":    "projects/sheltermap-1493101612061/regions/us-east4/subnetworks/default",
			"AccessConfigs": AccessConfigs,
		},
	}

	createnode := map[string]interface{}{
		"projectid":         "sheltermap-1493101612061",
		"Name":              "testing-scorelab2",
		"MachineType":       "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/machineTypes/n1-standard-1",
		"Zone":              "us-east4-c",
		"disk":              disk,
		"NetworkInterfaces": NetworkInterfaces,
	}

	_, err := gce.CreateNode(createnode)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
