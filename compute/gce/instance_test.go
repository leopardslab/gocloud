package gce

import "testing"

import "fmt"

func TestDeletenode(t *testing.T) {
	var gce GCE
	deletenode := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"instance":  "testing-scorelab",
		"Zone":      "us-west1-c",
	}

	_, err := gce.Deletenode(deletenode)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestStopnode(t *testing.T) {
	var gce GCE
	stopnode := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"instance":  "testing-scorelab",
		"Zone":      "us-west1-c",
	}

	_, err := gce.Stopnode(stopnode)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestRebootnode(t *testing.T) {
	var gce GCE
	reboot := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"instance":  "testing-scorelab",
		"Zone":      "us-west1-c",
	}
	_, err := gce.Rebootnode(reboot)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}

}

func TestCreatenode(t *testing.T) {

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
		"Name":              "testing-scorelab",
		"MachineType":       "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/machineTypes/n1-standard-1",
		"Zone":              "us-east4-c",
		"disk":              disk,
		"NetworkInterfaces": NetworkInterfaces,
	}

	_, err := gce.Createnode(createnode)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}
