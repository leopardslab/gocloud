package digioceanloadbalancer

import (
  "testing"
  digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
)

func init() {
	digioceanAuth.LoadConfig()
}

func TestCreateloadbalancer(t *testing.T) {

  forwardingrules := []map[string]interface{}{
      {
        "EntryProtocol":  "https",
        "EntryPort":  443,
        "TargetProtocol": "http",
        "TargetPort": 80,
        "CertificateID":  "a-b-c",
      },
      {
        "EntryProtocol":  "https",
        "EntryPort":  444,
        "TargetProtocol": "https",
        "TargetPort": 443,
        "TLSPassthrough": true,
      },
  }

  healthcheck := map[string]interface{}{
    "Protocol": "http",
    "Port": 80,
    "Path": "/",
    "CheckIntervalSeconds": 10,
    "ResponseTimeoutSeconds": 5,
    "HealthyThreshold": 5,
    "UnhealthyThreshold": 3,
	}

  stickysessions := map[string]interface{}{
    "Type": "none",
  }

	create := map[string]interface{}{
    "Name": "example-01",
    "Algorithm":  "round_robin",
    "Region": "nyc3",
    "ForwardingRules":  forwardingrules,
    "HealthCheck":  healthcheck,
    "StickySessions": stickysessions,
    "DropletIDs":  []int{3164444, 3164445},
    "Tag":  nil,
    "RedirectHTTPToHTTPS": false,
	}

	_, err := digioceancloud.Creatloadbalancer(create)

	if err != nil {
		t.Errorf("Test to create DigitalOcean LoadBalancer failed.")
	}
}

func TestDeleteloadbalancer(t *testing.T) {

  delete := map[string]string{
    "ID": "86407564",
   }

	_, err := digioceancloud.Deleteloadbalancer(delete)

	if err != nil {
		t.Errorf("Test to delete DigitalOcean LoadBalancer failed")
	}
}

func TestListloadbalancer(t *testing.T) {

	_, err := digioceancloud.Listloadbalancer(nil)

	if err != nil {
		t.Errorf("Test to list DigitalOcean LoadBalancer failed")
	}
}

func TestAttachnodewithloadbalancer(t *testing.T) {

  attachnodewithloadbalancer := map[string]interface{}{
    "LoadBalancerID":   "my-load-balancer",
    "DropletIDs":        []int{31331, 31313},
	}

  _, err := digioceancloud.Attachnodewithloadbalancer(attachnodewithloadbalancer)

	if err != nil {
		t.Errorf("Test to attach Droplets to DigitalOcean LoadBalancer failed")
	}
}

func TestDetachnodewithloadbalancer(t *testing.T) {

  detachnodewithloadbalancer := map[string]interface{}{
    "LoadBalancerID":   "my-load-balancer",
    "DropletIDs":        []int{31331, 31313},
	}

  _, err := digioceancloud.Detachnodewithloadbalancer(detachnodewithloadbalancer)

	if err != nil {
		t.Errorf("Test to detach Droplets to DigitalOcean LoadBalancer failed")
	}
}
