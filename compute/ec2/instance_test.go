package ec2

import "testing"
import "fmt"

import "github.com/scorelab/gocloud-v2/auth"

func init() {
	auth.LoadConfig()
}

func TestCreatenode(t *testing.T) {
	var amazoncloud EC2
	create := map[string]interface{}{
		"ImageId":      "ami-ccf405a5",
		"InstanceType": "t1.micro",
		"Region":       "us-east-1",
	}

	_, err := amazoncloud.Createnode(create)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestStopnode(t *testing.T) {
	var amazoncloud EC2
	stop := map[string]string{
		"instance-id": "i-06d518ba15b68685c",
		"Region":      "us-east-1",
	}
	_, err := amazoncloud.Stopnode(stop)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestStartnode(t *testing.T) {
	var amazoncloud EC2
	start := map[string]string{
		"instance-id": "i-0174bd6f54178e89b",
		"Region":      "us-east-1",
	}
	_, err := amazoncloud.Startnode(start)
	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestRebootnode(t *testing.T) {
	var amazoncloud EC2
	Reboot := map[string]string{
		"instance-id": "i-0174bd6f54178e89b",
		"Region":      "us-east-1",
	}
	_, err := amazoncloud.Rebootnode(Reboot)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestDeletnode(t *testing.T) {
	var amazoncloud EC2
	delete := map[string]string{
		"instance-id": "i-0174bd6f54178e89b",
		"Region":      "us-east-1",
	}
	_, err := amazoncloud.Deletenode(delete)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}
