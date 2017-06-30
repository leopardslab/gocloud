package gce

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/scorelab/gocloud-v2/googleauth"
	"io/ioutil"
	"net/http"
)

//create gce instance

func (gce *GCE) Createnode(request interface{}) (resp interface{}, err error) {

	var gceinstance GCE

	var projectid string

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "projectid":
			projectid, _ = value.(string)
			fmt.Println(projectid)

		case "Zone":
			Zone, _ := value.(string)
			gceinstance.Zone = Zone

		case "selfLink":
			selfLink, _ := value.(string)
			gceinstance.selfLink = selfLink

		case "Description":
			Description, _ := value.(string)
			gceinstance.Description = Description

		case "CanIPForward":
			CanIPForward, _ := value.(bool)
			gceinstance.CanIPForward = CanIPForward

		case "Name":
			Name, _ := value.(string)
			gceinstance.Name = Name

		case "MachineType":
			MachineType, _ := value.(string)
			gceinstance.MachineType = MachineType

		case "disk":
			diskparam, _ := value.([]map[string]interface{})
			var disk Disk
			var initializeParam InitializeParam
			for i := 0; i < len(diskparam); i++ {
				for diskparamkey, diskparamvalue := range diskparam[i] {
					switch diskparamkey {
					case "Type":
						disk.Type = diskparamvalue.(string)
					case "Boot":
						disk.Boot = diskparamvalue.(bool)
					case "Mode":
						disk.Mode = diskparamvalue.(string)
					case "AutoDelete":
						disk.AutoDelete = diskparamvalue.(bool)
					case "DeviceName":
						disk.DeviceName = diskparamvalue.(string)
					case "InitializeParams":
						InitializeParams, _ := diskparamvalue.(map[string]string)
						initializeParam.SourceImage = InitializeParams["SourceImage"]
						initializeParam.DiskType = InitializeParams["DiskType"]
						initializeParam.DiskSizeGb = InitializeParams["DiskSizeGb"]

					}
				}
				gceinstance.Disks = append(gceinstance.Disks, Disk{Type: disk.Type,
					Boot:       disk.Boot,
					Mode:       disk.Mode,
					AutoDelete: disk.AutoDelete,
					DeviceName: disk.DeviceName,
					InitializeParams: InitializeParam{
						SourceImage: initializeParam.SourceImage,
						DiskType:    initializeParam.DiskType,
						DiskSizeGb:  initializeParam.DiskSizeGb,
					}})

			}
		case "NetworkInterfaces":
			NetworkInterfacesparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(NetworkInterfacesparam); i++ {
				var networkInterfaceParam NetworkInterface
				for NetworkInterfaceparamkey, NetworkInterfaceparamvalue := range NetworkInterfacesparam[i] {
					switch NetworkInterfaceparamkey {
					case "Network":
						networkInterfaceParam.Network = NetworkInterfaceparamvalue.(string)

					case "Subnetwork":
						networkInterfaceParam.Subnetwork = NetworkInterfaceparamvalue.(string)

					case "AccessConfigs":
						AccessConfigsparam, _ := NetworkInterfaceparamvalue.([]map[string]string)
						for i := 0; i < len(AccessConfigsparam); i++ {
							var accessConfigParam accessConfig
							accessConfigParam.Name = AccessConfigsparam[i]["Name"]
							accessConfigParam.Type = AccessConfigsparam[i]["Type"]
							networkInterfaceParam.AccessConfigs = append(networkInterfaceParam.AccessConfigs, accessConfigParam)
						}
					}
				}
				gceinstance.NetworkInterfaces = append(gceinstance.NetworkInterfaces, networkInterfaceParam)
			}

		case "scheduling":
			schedulingparam, _ := value.(map[string]interface{})
			for key, value := range schedulingparam {
				switch key {
				case "Preemptible":
					Preemptible, _ := value.(bool)
					gceinstance.Scheduling.Preemptible = Preemptible

				case "onHostMaintenance":
					onHostMaintenance, _ := value.(string)
					gceinstance.Scheduling.OnHostMaintenance = onHostMaintenance

				case "automaticRestart":
					automaticRestart, _ := value.(bool)
					gceinstance.Scheduling.AutomaticRestart = automaticRestart
				}
			}

		}

	}

	gceinstancejson, _ := json.Marshal(gceinstance)
	gceinstancejsonstring := string(gceinstancejson)
	var gceinstancejsonstringbyte = []byte(gceinstancejsonstring)

	client := googleauth.SignJWT()

	Createnoderequest, err := http.NewRequest("POST", "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/instances", bytes.NewBuffer(gceinstancejsonstringbyte))
	Createnoderequest.Header.Set("Content-Type", "application/json")

	Createnoderesp, err := client.Do(Createnoderequest)
	defer Createnoderesp.Body.Close()

	body, err := ioutil.ReadAll(Createnoderesp.Body)
	fmt.Println(string(body))

	return
}

func (gce *GCE) Startnode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/instances/" + options["instance"] + "/start"

	client := googleauth.SignJWT()

	Startnoderequest, err := http.NewRequest("POST", url, nil)

	Startnoderequest.Header.Set("Content-Type", "application/json")

	Startnoderesp, err := client.Do(Startnoderequest)

	defer Startnoderesp.Body.Close()

	body, err := ioutil.ReadAll(Startnoderesp.Body)

	fmt.Println(string(body))

	return
}

//stop gce instance currentnly running
//accept projectid, zone, instance

func (gce *GCE) Stopnode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/instances/" + options["instance"] + "/stop"

	client := googleauth.SignJWT()

	Stopnoderequest, err := http.NewRequest("POST", url, nil)

	Stopnoderequest.Header.Set("Content-Type", "application/json")

	Stopnoderesp, err := client.Do(Stopnoderequest)
	defer Stopnoderesp.Body.Close()

	body, err := ioutil.ReadAll(Stopnoderesp.Body)
	fmt.Println(string(body))

	return
}

//delete gce instance currentnly running
//accept projectid, zone, instance

func (gce *GCE) Deletenode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/instances/" + options["instance"]

	client := googleauth.SignJWT()

	Deletenoderequest, err := http.NewRequest("DELETE", url, nil)
	Deletenoderequest.Header.Set("Content-Type", "application/json")

	Deletenoderesp, err := client.Do(Deletenoderequest)

	defer Deletenoderesp.Body.Close()
	body, err := ioutil.ReadAll(Deletenoderesp.Body)

	fmt.Println(string(body))

	return
}

//reboot/reset gce instance currentnly ***running***
//accept projectid, zone, instance

func (gce *GCE) Rebootnode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/instances/" + options["instance"] + "/reset"

	client := googleauth.SignJWT()

	Rebootnoderequest, err := http.NewRequest("POST", url, nil)
	Rebootnoderequest.Header.Set("Content-Type", "application/json")

	Rebootnoderesp, err := client.Do(Rebootnoderequest)

	defer Rebootnoderesp.Body.Close()
	body, err := ioutil.ReadAll(Rebootnoderesp.Body)
	fmt.Println(string(body))

	return
}

//list gce instance currentnly created
//accept projectid, zone

func (gce *GCE) listnode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/instances/"

	client := googleauth.SignJWT()

	listnoderequest, err := http.NewRequest("POST", url, nil)
	listnoderequest.Header.Set("Content-Type", "application/json")

	listnoderesp, err := client.Do(listnoderequest)

	defer listnoderesp.Body.Close()
	body, err := ioutil.ReadAll(listnoderesp.Body)

	fmt.Println(string(body))

	return
}
