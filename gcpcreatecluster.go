package main

import (
  "fmt"
  "github.com/cloudlibz/gocloud/gocloud"
)

func main() {

  googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider)

  InitializeParams := map[string]string{
     "SourceImage": "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301",
     "DiskType":    "projects/sheltermap-1493101612061/zones/us-east4-c/diskTypes/pd-standard",
     "DiskSizeGb":  "10",
   }

   disk := []map[string]interface{}{
      {
 	"Boot":  true,
 	"AutoDelete":   false,
 	"DeviceName":   "DeviceName",
 	"Type":         "PERSISTENT",
 	"Mode":         "READ_WRITE",
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
 	"Network": "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
 	"Subnetwork":    "projects/sheltermap-1493101612061/regions/us-east4/subnetworks/default",
         "AccessConfigs": AccessConfigs,
     },
    }

    createnode := map[string]interface{}{
       "projectid":   "sheltermap-1493101612061",
 	  "Name" :   "testing-scorelab",
     "MachineType":    "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/machineTypes/n1-standard-1",
     "Zone":        "us-central1-b",
     "disk":        disk,
     "NetworkInterfaces": NetworkInterfaces,
     }

   resp, err := googlecloud.Createnode(createnode)
   fmt.Println(err)
   response := resp.(map[string]interface{})
   fmt.Println(response["body"])

}
