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
	resp,_ := gce.Startnode(start)

  response := resp.(map[string]interface{})

	if(response["status"]=="200 OK"){
		fmt.Println(" Test pass")
	}else{
			fmt.Println("fail")
	}
}
