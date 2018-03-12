package main

import "github.com/cloudlibz/gocloud/gocloud"
import "fmt"


func main(){

	InitializeParams := map[string]string{
		"SourceImage": "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301",
		"DiskType":    "projects/sheltermap-1493101612061/zones/us-east4-c/diskTypes/pd-standard",
		"DiskSizeGb":  "10",
	}

	disk := []map[string]interface{}{
		{
		"Boot":             true,
		"AutoDelete":       false,
		"DeviceName":       "bokya",
		"Type":             "PERSISTENT",
		"Mode":             "READ_WRITE",
		"InitializeParams": InitializeParams,
		},
	}

	AccessConfigs := []map[string]string{{
		"Name": "external-nat",
		"Type": "ONE_TO_ONE_NAT",
	},
	}

NetworkInterfaces := []map[string]interface{}{
	{
		"Network":       "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
		"Subnetwork":    "projects/sheltermap-1493101612061/regions/us-east4/subnetworks/default",
		"AccessConfigs": AccessConfigs,
	},
	}

	createnode := map[string]interface{}{
	"projectid":         "sheltermap-1493101612061",
	"Name":              "testing-scorelab10",
	"MachineType":       "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/machineTypes/n1-standard-1",
	"Zone":              "us-east4-c",
	"disk":              disk,
	"NetworkInterfaces": NetworkInterfaces,
}

googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

 resp, err := googlecloud.Createnode(createnode)

 fmt.Println(err)
 	if err != nil {
 		fmt.Println(err)
 	}
 response := resp.(map[string]interface{})
 fmt.Println(response["body"])


	amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)

	create := map[string]interface{}{
		"ImageId":      "ami-ccf405a5",
		"InstanceType": "t1.micro",
		"Region":       "us-east-1",
	  }

	 resp, err = amazoncloud.Createnode(create)
	 response = resp.(map[string]interface{})
	 fmt.Println(response["body"])
}
