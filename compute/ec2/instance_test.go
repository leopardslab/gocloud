package ec2

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}

func TestCreateNode(t *testing.T) {
	var amazoncloud EC2
	create := map[string]interface{}{
		"ImageId":      "ami-ccf405a5",
		"InstanceType": "t1.micro",
		"Region":       "us-east-1",
	}

	_, err := amazoncloud.CreateNode(create)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestStopNode(t *testing.T) {
	var amazoncloud EC2
	stop := map[string]string{
		"instance-id": "i-06d518ba15b68685c",
		"Region":      "us-east-1",
	}
	_, err := amazoncloud.StopNode(stop)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestStartNode(t *testing.T) {
	var amazoncloud EC2
	start := map[string]string{
		"instance-id": "i-0174bd6f54178e89b",
		"Region":      "us-east-1",
	}
	_, err := amazoncloud.StartNode(start)
	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestRebootNode(t *testing.T) {
	var amazoncloud EC2
	Reboot := map[string]string{
		"instance-id": "i-037a9fae81c33ac30",
		"Region":      "us-east-1",
	}
	_, err := amazoncloud.RebootNode(Reboot)

	if err != nil {
		t.Errorf("Test Pass")
	}
}

func TestDeletnode(t *testing.T) {
	var amazoncloud EC2
	delete := map[string]string{
		"instance-id": "i-0076d8bd2fbefcb61",
		"Region":      "us-east-1",
	}

	amazoncloud.DeleteNode(delete)
}
