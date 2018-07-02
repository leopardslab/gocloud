package droplet

import (
	digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
	"testing"
)

func init() {
	digioceanAuth.LoadConfig()
}

func TestCreateNode(t *testing.T) {

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

	_, err := digioceancloud.CreateNode(create)

	if err != nil {
		t.Errorf("Test failed.")
	}
}

func TestStopNode(t *testing.T) {

	var digioceancloud Droplet

	stop := map[string]string{
		"ID": "86407564",
	}

	_, err := digioceancloud.StopNode(stop)

	if err != nil {
		t.Errorf("Test failed.")
	}
}

func TestStartNode(t *testing.T) {

	var digioceancloud Droplet

	start := map[string]string{
		"ID": "86407564",
	}

	_, err := digioceancloud.StartNode(start)
	if err != nil {
		t.Errorf("Test failed.")
	}
}

func TestRebootNode(t *testing.T) {

	var digioceancloud Droplet

	reboot := map[string]string{
		"ID": "86407564",
	}

	_, err := digioceancloud.RebootNode(reboot)

	if err != nil {
		t.Errorf("Test Pass")
	}
}

func TestDeletnode(t *testing.T) {

	var digioceancloud Droplet

	delete := map[string]string{
		"ID": "86407564",
	}

	digioceancloud.DeleteNode(delete)
}
