package main

import (
	"fmt"
	"github.com/cloudlibz/gocloud/gocloud"
)

func main() {

	digioceancloud := gocloud.DigitalOceanProvider()

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
