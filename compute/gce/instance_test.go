package gce

import "testing"

import "fmt"

func TestStartnode(t *testing.T) {
	var gce GCE
	start := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"instance":  "sumesh-110",
		"Zone":      "us-east4-c",
	}
	resp, _ := gce.Startnode(start)

	response := resp.(map[string]interface{})

	if response["status"] == "200" {
		fmt.Println(" Test pass")
	} else {
		fmt.Println(" Test fail")
	}
}



func TestStopnode(t *testing.T) {
	var gce GCE

	stop := map[string]string{
  "projectid":"sheltermap-1493101612061",
	"instance":"instance-10",
	"Zone": "us-west1-c",
 }

 resp, _ := gce.Stopnode(stop)

 response := resp.(map[string]interface{})

 if response["status"] == "200" {
 	fmt.Println(" Test pass")
 } else {
 	fmt.Println(" Test fail")
 }

 }
