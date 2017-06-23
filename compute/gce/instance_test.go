package gce

import "testing"

import "fmt"

func TestStartnode(t *testing.T) {
	var gce GCE
	start := map[string]string{
		"projectid": "sheltermap-1493101612061",
		"instance":  "sumesh-10",
		"Zone":      "us-east4-c",
	}
	resp, _ := gce.Startnode(start)
	fmt.Println(resp)
}
