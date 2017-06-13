package main

import(
  . "github.com/scorelab/gocloud-v2/gocloud"
)


func main(){

    InitializeParams := map[string]string{
    "SourceImage" : "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301",
    "DiskType" :  "1",
    "DiskSizeGb" :"1",
  }

    disk := []map[string]interface{}{
    {
    "Type": "usb",
    "Boot":  true,
    "Mode": "",
    "AutoDelete": false,
    "deviceName": "",
    "InitializeParams": InitializeParams,
  },

    {
    "Type": "uss",
    "Boot":  true,
    "Mode": "",
    "AutoDelete": false,
    "deviceName": "",
    "InitializeParams": InitializeParams,
  },
  }

    AccessConfigs := []map[string]string{{
      "Name" : "external-nat",
      "Type" :  "ONE_TO_ONE_NAT",
    },{
      "Name" : "external-nat",
      "Type" :  "ONE_TO_ONE_NAT",
    }}

  AccessConfigs2 := []map[string]string{{
      "Name" : "external-nat",
      "Type" :  "ONE_TO_ONE_NAT",
      },}


    NetworkInterfaces := []map[string]interface{}{{
    "Network": "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
    "Subnetwork": "",
    "AccessConfigs": AccessConfigs,
  },

    {
    "Network": "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
    "Subnetwork": "hello",
    "AccessConfigs": AccessConfigs2,
  },
  }

  scheduling := map[string]interface{}{
    "Preemptible": false,
    "OnHostMaintenance": "MIGRATE",
    "AutomaticRestart": true,
  }


    gce := map[string]interface{}{
    "projectid":"sheltermap-1493101612061",
    "Name": "scorelab",
    "MachineType":"n1-standard-1",
    "Zone":"us-east4-c",
    "disk": disk,
    "NetworkInterfaces": NetworkInterfaces,
    "description":""
    "scheduling" : scheduling,
    }





  googlecloud, _ := CloudProvider(Googleprovider,Secretkey,Secretid)
  googlecloud.Createnode(gce)


}
