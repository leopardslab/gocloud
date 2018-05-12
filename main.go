package main

import("fmt"
		"github.com/cloudlibz/gocloud/gocloud"
)

func main(){

	fmt.Println("Hello World")

	googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)
//projects/adept-comfort-202709/locations/us-central1/functions/function-1

	/*
	deletefunction := map[string]string{
		"name": "projects/adept-comfort-202709/locations/us-central1",
		"pageSize": "1",
	}
	*/
	deletefunction := map[string]string{
		"name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
	}

  resp, _ := googlecloud.Callfunction(deletefunction)

 response := resp.(map[string]interface{})

 	fmt.Println(response["body"])
}
