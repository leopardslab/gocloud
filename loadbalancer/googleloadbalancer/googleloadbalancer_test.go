package googleloadbalancer

import "testing"

func TestCreateLoadBalancer(t *testing.T) {

	var googleloadbalancer Googleloadbalancer

	creatloadbalancer := map[string]interface{}{
		"Project":   "sheltermap-1493101612061",
		"Name":      "google-loadbalancer",
		"Region":    "us-central1",
		"Instances": []string{"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-central1-b/instances/instance-1"},
	}

	_, err := googleloadbalancer.CreateLoadBalancer(creatloadbalancer)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListLoadBalancer(t *testing.T) {
	var googleloadbalancer Googleloadbalancer

	listloadbalancer := map[string]string{
		"Project": "sheltermap-1493101612061",
		"Region":  "us-central1",
	}

	_, err := googleloadbalancer.ListLoadBalancer(listloadbalancer)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestAttachNodeWithLoadBalancer(t *testing.T) {
	var googleloadbalancer Googleloadbalancer

	attachnodewithloadbalancer := map[string]interface{}{
		"Project":    "sheltermap-1493101612061",
		"Region":     "us-central1",
		"TargetPool": "google-loadbalancer",
		"Instances":  []string{"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-central1-b/instances/instance-2"},
	}

	_, err := googleloadbalancer.AttachNodeWithLoadBalancer(attachnodewithloadbalancer)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDetachNodeWithLoadBalancer(t *testing.T) {
	var googleloadbalancer Googleloadbalancer

	detachnodewithloadbalancer := map[string]interface{}{
		"Project":    "sheltermap-1493101612061",
		"Region":     "us-central1",
		"TargetPool": "google-loadbalancer",
		"Instances":  []string{"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-central1-b/instances/instance-2"},
	}

	_, err := googleloadbalancer.DetachNodeWithLoadBalancer(detachnodewithloadbalancer)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteLoadBalancer(t *testing.T) {

	var googleloadbalancer Googleloadbalancer

	deleteloadbalancer := map[string]string{
		"Project":    "sheltermap-1493101612061",
		"Region":     "us-central1",
		"TargetPool": "google-loadbalancer",
	}

	_, err := googleloadbalancer.DeleteLoadBalancer(deleteloadbalancer)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
