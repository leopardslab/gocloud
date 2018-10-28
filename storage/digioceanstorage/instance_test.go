package digioceanstorage

import (
	digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
	"testing"
)

func init() {
	digioceanAuth.LoadConfig()
}

func TestCreateDisk(t *testing.T) {
	var digioceancloud Digioceanstorage

	create := map[string]interface{}{
		"Name":          "example-01",
		"Region":        "nyc3",
		"Description":   "Block store for examples",
		"SizeGigaBytes": 1,
		"SnapshotID":    nil,
	}

	_, err := digioceancloud.CreateDisk(create)

	if err != nil {
		t.Errorf("Test to create DigitalOcean Storage volume failed.")
	}
}

func TestDeleteDisk(t *testing.T) {
	var digioceancloud Digioceanstorage

	delete := map[string]string{
		"VolumeID": "7724db7c-e098-11e5-b522-000f53304e51",
	}

	_, err := digioceancloud.DeleteDisk(delete)

	if err != nil {
		t.Errorf("Test to delete DigitalOcean Storage volume failed.")
	}
}

func TestCreateSnapshot(t *testing.T) {

	var digioceancloud Digioceanstorage

	create := map[string]interface{}{
		"VolumeID":     "7724db7c-e098-11e5-b522-000f53304e5",
		"SnapshotName": "big-data-snapshot1475261774",
	}

	_, err := digioceancloud.CreateSnapshot(create)

	if err != nil {
		t.Errorf("Test to create DigitalOcean Snapshot failed.")
	}
}

func TestDeleteSnapshot(t *testing.T) {
	var digioceancloud Digioceanstorage

	delete := map[string]string{
		"SnapshotID": "7724db7c-e098-11e5-b522-000f53304e51",
	}

	_, err := digioceancloud.DeleteSnapshot(delete)

	if err != nil {
		t.Errorf("Test to delete DigitalOcean Snapshot failed.")
	}
}

func TestAttachDisk(t *testing.T) {

	var digioceancloud Digioceanstorage

	create := map[string]interface{}{
		"VolumeID":  "7724db7c-e098-11e5-b522-000f53304e51",
		"DropletID": 9978454,
		"Region":    "nyc3",
	}

	_, err := digioceancloud.AttachDisk(create)

	if err != nil {
		t.Errorf("Test to attach disk to DigitalOcean Droplet failed.")
	}
}

func TestDetachDisk(t *testing.T) {
	var digioceancloud Digioceanstorage

	delete := map[string]interface{}{
		"VolumeID":  "7724db7c-e098-11e5-b522-000f53304e51",
		"DropletID": 9978454,
		"Region":    "nyc3",
	}

	_, err := digioceancloud.DetachDisk(delete)

	if err != nil {
		t.Errorf("Test to detach disk from DigitalOcean Droplet failed.")
	}
}
