package main

import (
	"fmt"
	"github.com/cloudlibz/gocloud/gocloud"
)

func main() {

	digioceancloud := gocloud.DigitalOceanProvider()

	/*
	   	image := map[string]interface{}{
	       "Slug": "ubuntu-16-04-x64",
	     }

	     create := map[string]interface{}{
	       "Name":   "example.com",
	       "Region": "nyc3",
	       "Size":   "s-1vcpu-1gb",
	       "Image": image,
	       "SSHKeys": nil,
	       "Backups": false,
	       "IPv6": true,
	       "UserData": nil,
	       "PrivateNetworking": nil,
	       "Volumes": nil,
	       "Monitoring": false,
	       "Tags": []string{"web"},
	     }

	    resp, err := digioceancloud.CreateNode(create)
	    if err != nil{
	   	 fmt.Println(err)
	    }
	    response := resp.(map[string]interface{})
	    fmt.Println(response["body"])
	*/

	stop := map[string]string{
		"ID": "102694881",
	}

	resp, _ := digioceancloud.StopNode(stop)
	response := resp.(map[string]interface{})
	fmt.Println(response["body"])

	start := map[string]string{
		"ID": "102694881",
	}

	resp, _ = digioceancloud.StartNode(start)
	response = resp.(map[string]interface{})
	fmt.Println(response["body"])

	delete := map[string]string{
		"ID": "102694881",
	}

	resp, _ = digioceancloud.DeleteNode(delete)
	response = resp.(map[string]interface{})
	fmt.Println(response["body"])
}
