package ec2

import "testing"
import "fmt"



func TestCreatenode(t *testing.T) {
  var amazoncloud EC2
  create := map[string]interface{}{
 "ImageId"     : "ami-ccf405a5",
 "InstanceType": "t1.micro",
 "Region" :"us-east-1",
}

resp,_:= amazoncloud.Createnode(create)
fmt.Println(resp)
}


func TestStartnode(t *testing.T) {
  var amazoncloud EC2
  start := map[string]string{
		"instance-id":	"i-0174bd6f54178e89b",
		"Region"     :	"us-east-1",
	}
	 resp,_:= amazoncloud.Startnode(start)
  fmt.Println(resp)
}


func TestStopnode(t *testing.T) {
  var amazoncloud EC2
  stop := map[string]string{
		"instance-id":	"i-0174bd6f54178e89b",
		"Region"     :	"us-east-1",
	}
	 resp,_:= amazoncloud.Stopnode(stop)
  fmt.Println(resp)
}


func TestRebootnode(t *testing.T) {
  var amazoncloud EC2
  Reboot := map[string]string{
		"instance-id":	"i-0174bd6f54178e89b",
		"Region"     :	"us-east-1",
	}
	 resp,_:= amazoncloud.Rebootnode(Reboot)
  fmt.Println(resp)
}


func TestDeletnode(t *testing.T) {
  var amazoncloud EC2
  delete := map[string]string{
		"instance-id":	"i-0174bd6f54178e89b",
		"Region"     :	"us-east-1",
	}
	 resp,_:= amazoncloud.Deletenode(delete)
  fmt.Println(resp)
}
