package main

import "github.com/cloudlibz/gocloud/gocloud"
import "fmt"

func main() {

	googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

	deletetables := map[string]string{
		"name": "projects/adept-comfort-202709/instances/helloo/tables/bokkkya",
	}

	resp, _ := googlecloud.Deletetables(deletetables)

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}
