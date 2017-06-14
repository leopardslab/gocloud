package gce

import(
  "fmt"
  "net/http"
  "io/ioutil"
)


func (gce *GCE)Createnode(request interface{})(resp interface{},err error){

 return
}


func (gce *GCE)Startnode(request interface{}) (resp interface{}, err error){

    options := request.(map[string]string)

    url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/"+ options["Zone"] + "/instances/" + options["instance"] + "/start"


    token := sign()

    client := &http.Client{}

    Startnoderequest, err := http.NewRequest("POST",url, nil)

    Startnoderequest.Header.Set("Content-Type", "application/json")

    token.SetAuthHeader(Startnoderequest)

    Startnoderesp, err := client.Do(Startnoderequest)

    defer Startnoderesp.Body.Close()

    body, err := ioutil.ReadAll(Startnoderesp.Body)

    fmt.Println(string(body))

  return
}


func (gce *GCE)Stopnode(request interface{}) (resp interface{}, err error){

  options := request.(map[string]string)

  url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/"+ options["Zone"] + "/instances/" + options["instance"] + "/stop"

  token := sign()

  client := &http.Client{}

  Stopnoderequest, err := http.NewRequest("POST",url, nil)

  Stopnoderequest.Header.Set("Content-Type", "application/json")

  token.SetAuthHeader(Stopnoderequest)

  Stopnoderesp, err := client.Do(Stopnoderequest)

  defer Stopnoderesp.Body.Close()

  body, err := ioutil.ReadAll(Stopnoderesp.Body)

  fmt.Println(string(body))

  return
}

func (gce *GCE)Deletenode(request interface{}) (resp interface{}, err error){

  options := request.(map[string]string)

  url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/"+ options["Zone"] + "/instances/" + options["instance"]

  token := sign()

  client := &http.Client{}

  Deletenoderequest, err := http.NewRequest("DELETE",url, nil)

  Deletenoderequest.Header.Set("Content-Type", "application/json")

  token.SetAuthHeader(Deletenoderequest)

  Deletenoderesp, err := client.Do(Deletenoderequest)

  defer Deletenoderesp.Body.Close()

  body, err := ioutil.ReadAll(Deletenoderesp.Body)

  fmt.Println(string(body))

  return
}


func (gce *GCE)Rebootnode(request interface{}) (resp interface{}, err error){

  options := request.(map[string]string)

  url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/"+ options["Zone"] + "/instances/" + options["instance"] + "/reset"

  fmt.Println(url)

  token := sign()

  client := &http.Client{}

  Rebootnoderequest, err := http.NewRequest("POST",url, nil)

  Rebootnoderequest.Header.Set("Content-Type", "application/json")

  token.SetAuthHeader(Rebootnoderequest)

  Rebootnoderesp, err := client.Do(Rebootnoderequest)

  defer Rebootnoderesp.Body.Close()

  body, err := ioutil.ReadAll(Rebootnoderesp.Body)

  fmt.Println(string(body))


  return
}

func (gce *GCE)listnode(request interface{})(resp interface{},err error){

  options := request.(map[string]string)

  url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/"+ options["Zone"] + "/instances/"

  fmt.Println(url)

  token := sign()

  client := &http.Client{}

  listnoderequest, err := http.NewRequest("POST",url, nil)

  listnoderequest.Header.Set("Content-Type", "application/json")

  token.SetAuthHeader(listnoderequest)

  listnoderesp, err := client.Do(listnoderequest)

  defer listnoderesp.Body.Close()

  body, err := ioutil.ReadAll(listnoderesp.Body)

  fmt.Println(string(body))

  return
}


/*
import(
  "fmt"
  "net"
  "net/http"
  "io/ioutil"
  "bytes"
  "reflect"
 "encoding/json"
)




func (gce *GCE)listnode(options interface{})(resp interface{},err error){


  token := sign()

  client := &http.Client{}

  req, _ := http.NewRequest("GET", "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/instances", nil)

  token.SetAuthHeader(req)

  res,_ := client.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
  return
}


func (gce *GCE)Createnode(options interface{})(resp interface{},err error){


	 token := sign()

  client := &http.Client{}

	var jsonStr2 = []byte(`{
  "disks":[
    {
      "boot":true,
      "initializeParams":{
        "sourceImage":"https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301"
      }
    }
  ],
  "machineType":"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/machineTypes/n1-standard-1",
  "name":"gocloud-100",
  "networkInterfaces":[
    {
      "accessConfigs":[
        {
          "name":"external-nat",
          "type":"ONE_TO_ONE_NAT"
        }
      ],
      "network":"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default"
    }
  ],
  "zone":"https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c"
}`)

	req2, err := http.NewRequest("POST", "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/instances", bytes.NewBuffer(jsonStr2))

	req2.Header.Set("Content-Type", "application/json")

	//req2.Header.Set("Content-length", "743")

	token.SetAuthHeader(req2)

	resp2, err := client.Do(req2)

	defer resp2.Body.Close()

	body2, err := ioutil.ReadAll(resp2.Body)

	fmt.Println(string(body2))


	return

}




func (gce *GCE)deletenode(options interface{})(resp interface{},err error){

req5, err := http.NewRequest("DELETE", "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/instances/instance-10", nil)


	req5.Header.Set("Content-Type", "application/json")

	//req2.Header.Set("Content-length", "743")

	tok.SetAuthHeader(req5)

	resp5, err := client.Do(req5)

	defer resp4.Body.Close()

	body5, err := ioutil.ReadAll(resp5.Body)

	fmt.Println("\n\n\n\n\n\n\n\n\n\n")

	fmt.Println(string(body5))



}





func (gce *GCE)stopnode(options interface{})(resp interface{},err error){

	req4, err := http.NewRequest("POST", "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/instances/instance-10/stop", nil)


	req3.Header.Set("Content-Type", "application/json")

	//req2.Header.Set("Content-length", "743")

	tok.SetAuthHeader(req4)

	resp4, err := client.Do(req4)

	defer resp4.Body.Close()

	body4, err := ioutil.ReadAll(resp4.Body)

	fmt.Println("\n\n\n\n\n\n\n\n\n\n")

	fmt.Println(string(body4))


}




func (gce *GCE)startnode(options interface{})(resp interface{},err error){

	req3, err := http.NewRequest("POST", "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/instances/instance-11/start", nil)


	req3.Header.Set("Content-Type", "application/json")

	//req2.Header.Set("Content-length", "743")

	tok.SetAuthHeader(req3)

	resp3, err := client.Do(req3)

	defer resp3.Body.Close()

	body3, err := ioutil.ReadAll(resp3.Body)


	fmt.Println(string(body3))


	fmt.Println("\n\n\n\n\n\n\n\n\n\n")

}






func (gce *GCE)Createnode(options interface{})(resp interface{},err error){

    var gceinstance GCE

    var projectid string

    param := make(map[string]interface{})

    param = options.(map[string]interface{})

    for key,value := range param{
        switch key {
             case "projectid":
                projectid,_ = value.(string)

              case "Zone":
                 Zone,_ := value.(string)
                 gceinstance.Zone = Zone

              case "selfLink":
                  selfLink,_ := value.(string)
                  gceinstance.selfLink = selfLink

              case "Description":
                    Description,_ := value.(string)
                    gceinstance.Description = Description

              case "CanIPForward":
                      CanIPForward,_ := value.(bool)
                      gceinstance.CanIPForward = CanIPForward

              case "Name":
                    Name,_ := value.(string)
                    gceinstance.Name = Name

              case "MachineType":
                    MachineType,_ := value.(string)
                    gceinstance.MachineType = "https://www.googleapis.com/compute/v1/projects/" + projectid + "zones/" +  gceinstance.Zone + "machineTypes/" +  MachineType

              case "disk":
                     diskparam,_ := value.([]map[string]interface{})
                     var disk Disk
                     var initializeParam InitializeParam
                     for i := 0; i < len(diskparam); i++{
                        for diskparamkey,diskparamvalue := range diskparam[i]{
                                switch diskparamkey{
                                   case "Type":
                                        disk.Type =diskparamvalue.(string)
                                    case "Boot":
                                        disk.Boot =diskparamvalue.(bool)
                                    case "Mode":
                                        disk.Mode =diskparamvalue.(string)
                                      case "AutoDelete":
                                          disk.AutoDelete =diskparamvalue.(bool)
                                      case "DeviceName":
                                          disk.DeviceName =diskparamvalue.(string)
                                     case "InitializeParams":
                                         InitializeParams,_ := diskparamvalue.(map[string]string)
                                         initializeParam.SourceImage=InitializeParams["SourceImage"]
                                         initializeParam.DiskType=InitializeParams["DiskType"]
                                         initializeParam.DiskSizeGb=InitializeParams["DiskSizeGb"]

                               }
                     }
                     gceinstance.Disks = append(gceinstance.Disks,Disk{ Type:disk.Type,
                                                                        Boot:disk.Boot,
                                                                        Mode:disk.Mode,
                                                                        AutoDelete:disk.AutoDelete,
                                                                        DeviceName:disk.DeviceName,
                                                                        InitializeParams:InitializeParam{
                                                                                  SourceImage:initializeParam.SourceImage,
                                                                                  DiskType:initializeParam.DiskType,
                                                                                  DiskSizeGb : initializeParam.DiskSizeGb,
                                                                                  }})

                   }
                   case "NetworkInterfaces":
                       NetworkInterfacesparam,_ := value.([]map[string]interface{})
                       for i := 0; i < len(NetworkInterfacesparam); i++{
                         var networkInterfaceParam NetworkInterface
                         for NetworkInterfaceparamkey,NetworkInterfaceparamvalue := range NetworkInterfacesparam[i]{
                            switch NetworkInterfaceparamkey{
                            case "Network":
                              networkInterfaceParam.Network = NetworkInterfaceparamvalue.(string)

                            case "Subnetwork":
                              networkInterfaceParam.Subnetwork = NetworkInterfaceparamvalue.(string)

                            case "AccessConfigs":
                                 AccessConfigsparam,_ := NetworkInterfaceparamvalue.([]map[string]string)
                                  for i := 0; i < len(AccessConfigsparam); i++{
                                         var accessConfigParam accessConfig
                                                  accessConfigParam.Name =AccessConfigsparam[i]["Name"]
                                                  accessConfigParam.Type =AccessConfigsparam[i]["Type"]
                                  networkInterfaceParam.AccessConfigs = append(networkInterfaceParam.AccessConfigs,accessConfigParam)
                                  }
                            }
                          }
                          gceinstance.NetworkInterfaces = append(gceinstance.NetworkInterfaces,networkInterfaceParam)
                        }

                    case "scheduling":
                    schedulingparam,_ := value.(map[string]interface{})
                    for key,value := range schedulingparam {
                        switch key {
                            case "Preemptible":
                                Preemptible,_ := value.(bool)
                                gceinstance.Scheduling.Preemptible = Preemptible

                              case "onHostMaintenance":
                                 onHostMaintenance,_ := value.(string)
                                 gceinstance.Scheduling.OnHostMaintenance = onHostMaintenance

                              case "automaticRestart":
                                  automaticRestart,_ := value.(bool)
                                  gceinstance.Scheduling.AutomaticRestart = automaticRestart
                        }
                    }

                 }

}

    fmt.Println(gceinstance)
}

*/
