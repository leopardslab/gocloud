package digioceanstorage

import (
  "testing"
  digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
)

func init() {
	digioceanAuth.LoadConfig()
}

func TestCreatedisk(t *testing.T) {

  var digioceancloud LoadBalancer

	create := map[string]interface{}{
    "Name": "example-01",
    "Region": "nyc3",
    "Description":  "Block store for examples",
    "SizeGigaBytes":  1,
    "SnapshotID": nil,
	}

	_, err := digioceancloud.Createdisk(create)

	if err != nil {
		t.Errorf("Test to create DigitalOcean Storage volume failed.")
	}
}

func TestDeletedisk(t *testing.T) {

  var digioceancloud LoadBalancer

  delete := map[string]string{
    "VolumeID": "7724db7c-e098-11e5-b522-000f53304e51",
   }

	_, err := digioceancloud.Deletedisk(delete)

	if err != nil {
		t.Errorf("Test to delete DigitalOcean Storage volume failed.")
	}
}

func TestCreatesnapshot(t *testing.T) {

  var digioceancloud LoadBalancer

  create := map[string]interface{}{
    "VolumeID": "7724db7c-e098-11e5-b522-000f53304e5",
    "Name": "big-data-snapshot1475261774",
	}

	_, err := digioceancloud.Createsnapshot(create)

	if err != nil {
		t.Errorf("Test to create DigitalOcean Snapshot failed.")
	}
}

func TestDeletesnapshot(t *testing.T) {

  var digioceancloud LoadBalancer

  delete := map[string]string{
    "SnapshotID": "7724db7c-e098-11e5-b522-000f53304e51",
   }

	_, err := digioceancloud.Deletesnapshot(delete)

	if err != nil {
		t.Errorf("Test to delete DigitalOcean Snapshot failed.")
	}
}

func TestAttachdisk(t *testing.T) {

  var digioceancloud LoadBalancer

  create := map[string]interface{}{
    "VolumeID": "7724db7c-e098-11e5-b522-000f53304e51",
    "DropletID": "9978454",
    "Region":  "nyc3",
	}

	_, err := digioceancloud.Attachdisk(create)

	if err != nil {
		t.Errorf("Test to attach disk to DigitalOcean Droplet failed.")
	}
}

func TestDetachdisk(t *testing.T) {

  var digioceancloud LoadBalancer

  create := map[string]interface{}{
    "VolumeID": "7724db7c-e098-11e5-b522-000f53304e51",
    "DropletID": "9978454",
    "Region":  "nyc3",
	}

	_, err := digioceancloud.Detachdisk(create)

	if err != nil {
		t.Errorf("Test to detach disk from DigitalOcean Droplet failed.")
	}
}
