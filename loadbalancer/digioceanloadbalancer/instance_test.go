package digioceanloadbalancer

import (
	digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
	"testing"
)

func init() {
	digioceanAuth.LoadConfig()
}

func TestCreateLoadBalancer(t *testing.T) {

	var digioceancloud DigioceanLoadBalancer

	forwardingrules := []map[string]interface{}{
		{
			"EntryProtocol":  "https",
			"EntryPort":      444,
			"TargetProtocol": "https",
			"TargetPort":     443,
			"TLSPassthrough": true,
		},
	}

	healthcheck := map[string]interface{}{
		"Protocol":               "http",
		"Port":                   80,
		"Path":                   "/",
		"CheckIntervalSeconds":   10,
		"ResponseTimeoutSeconds": 5,
		"HealthyThreshold":       5,
		"UnhealthyThreshold":     3,
	}

	stickysessions := map[string]interface{}{
		"Type": "none",
	}

	create := map[string]interface{}{
		"Name":            "example-01",
		"Algorithm":       "round_robin",
		"Region":          "nyc3",
		"ForwardingRules": forwardingrules,
		"HealthCheck":     healthcheck,
		"StickySessions":  stickysessions,
		"DropletIDs":      []int{3164444, 3164445},
		"Tag":             nil,
		"RedirectHTTPToHTTPS": false,
	}

	_, err := digioceancloud.CreateLoadBalancer(create)

	if err != nil {
		t.Errorf("Test to create DigitalOcean LoadBalancer failed.")
	}
}

func TestDeleteLoadBalancer(t *testing.T) {

	var digioceancloud DigioceanLoadBalancer

	delete := map[string]string{
		"ID": "86407564",
	}

	_, err := digioceancloud.DeleteLoadBalancer(delete)

	if err != nil {
		t.Errorf("Test to delete DigitalOcean LoadBalancer failed")
	}
}

func TestListLoadBalancer(t *testing.T) {

	var digioceancloud DigioceanLoadBalancer

	_, err := digioceancloud.ListLoadBalancer(nil)

	if err != nil {
		t.Errorf("Test to list DigitalOcean LoadBalancer failed")
	}
}

func TestAttachNodeWithLoadBalancer(t *testing.T) {

	var digioceancloud DigioceanLoadBalancer

	attachnodewithloadbalancer := map[string]interface{}{
		"LoadBalancerID": "my-load-balancer",
		"DropletIDs":     []int{31331, 31313},
	}

	_, err := digioceancloud.AttachNodeWithLoadBalancer(attachnodewithloadbalancer)

	if err != nil {
		t.Errorf("Test to attach Droplets to DigitalOcean LoadBalancer failed")
	}
}

func TestDetachNodeWithLoadBalancer(t *testing.T) {

	var digioceancloud DigioceanLoadBalancer

	detachnodewithloadbalancer := map[string]interface{}{
		"LoadBalancerID": "my-load-balancer",
		"DropletIDs":     []int{31331, 31313},
	}

	_, err := digioceancloud.DetachNodeWithLoadBalancer(detachnodewithloadbalancer)

	if err != nil {
		t.Errorf("Test to detach Droplets from DigitalOcean LoadBalancer failed")
	}
}
