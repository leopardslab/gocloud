package main

import (
	"fmt"
	"github.com/cloudlibz/gocloud/gocloud"
)

func main() {

	googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

	createtables := map[string]interface{}{
		"parent": "projects/adept-comfort-202709/instances/helloo",
	}

	resp, _ := googlecloud.Createtables(createtables)

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])

	/*
		deletetables := map[string]string{
			"name": "projects/adept-comfort-202709/instances/helloo/tables/hello",
		}
	*/
	//	projects/adept-comfort-202709/instances//tables/<table>.

	describetables := map[string]string{
		"name": "projects/adept-comfort-202709/instances/helloo/tables/bokkkya",
		//	"view" : "FULL",
		//"pageToken" : "",
	}
	resp, _ = googlecloud.Describetables(describetables)

	response = resp.(map[string]interface{})

	fmt.Println(response["body"])

	listtables := map[string]string{
		"parent":    "projects/adept-comfort-202709/instances/helloo",
		"view":      "NAME_ONLY",
		"pageToken": "",
	}
	resp, _ = googlecloud.Listtables(listtables)

	response = resp.(map[string]interface{})

	fmt.Println(response["body"])

}

/*
func main() {

	amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)

	listtables := map[string]interface{}{
		"Region":    "us-east-2",
		"TableName": "hello",
	}

	resp, _ := amazoncloud.Describetables(listtables)
	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
	fmt.Println(response["status"])


	    deletetables := map[string]interface{}{
	    "Region": "us-east-2",
	    "TableName" : "hello",
	    }

	   resp, _ := amazoncloud.Deletetables(deletetables)
	   response := resp.(map[string]interface{})
	   fmt.Println(response["body"])
	   fmt.Println(response["status"])


}

*/

/*
func main(){


	googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)
//projects/adept-comfort-202709/locations/us-central1/functions/function-1


	deletefunction := map[string]string{
		"name": "projects/adept-comfort-202709/locations/us-central1",
		"pageSize": "1",
	}


	deletefunction := map[string]string{
		"name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
	}

  resp, _ := googlecloud.Callfunction(deletefunction)

 	response := resp.(map[string]interface{})

 	fmt.Println(response["body"])
}
*/
