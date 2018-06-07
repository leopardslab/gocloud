package main

import "github.com/cloudlibz/gocloud/gocloud"
import "fmt"

func main() {

	googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)


		defaultVersion := make(map[string]interface{})

		defaultVersion["Name"] = "bokkya"

	updateMLModel := map[string]interface{}{
		"Parent": "projects/adept-comfort-202709/models/pratik",
		"Name"  :"pratik",
		"DefaultVersion" :defaultVersion,
		"UpdateMask" : "description",
	}

	resp, err := googlecloud.UpdateMLModel(updateMLModel)

	if err != nil {

	}

	response := resp.(map[string]interface{})
	fmt.Println(response["body"])

}
