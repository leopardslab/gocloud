package droplet

import (
	digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
	"testing"
)

func init() {
	digioceanAuth.LoadConfig()
}

func TestCreatenode(t *testing.T) {

	var digioceancloud Droplet

  image := map[string]interface{}{
    "Slug": "ubuntu-16-04-x64",
	}

	create := map[string]interface{}{
		"Name":              "example.com",
		"Region":            "nyc3",
		"Size":              "s-1vcpu-1gb",
		"Image":             image,
		"SSHKeys":           nil,
		"Backups":           false,
		"IPv6":              true,
		"UserData":          nil,
		"PrivateNetworking": nil,
		"Volumes":           nil,
		"Monitoring":        false,
		"Tags":              []string{"web"},
	}

	_, err := digioceancloud.Createnode(create)

	if err != nil {
		t.Errorf("Test failed.")
	}
}

func TestStopnode(t *testing.T) {

	var digioceancloud Droplet

  stop := map[string]string{
    "ID": "86407564",
   }

	_, err := digioceancloud.Stopnode(stop)

	if err != nil {
		t.Errorf("Test failed.")
	}
}

func TestStartnode(t *testing.T) {

  var digioceancloud Droplet

  start := map[string]string{
    "ID": "86407564",
   }

	_, err := digioceancloud.Startnode(start)
	if err != nil {
		t.Errorf("Test failed.")
	}
}

func TestRebootnode(t *testing.T) {

  var digioceancloud Droplet

  reboot := map[string]string{
    "ID": "86407564",
   }

	_, err := digioceancloud.Rebootnode(reboot)

	if err != nil {
		t.Errorf("Test Pass")
	}
}

func TestDeletnode(t *testing.T) {

  var digioceancloud Droplet

  delete := map[string]string{
    "ID": "86407564",
   }

	digioceancloud.Deletenode(delete)
}
