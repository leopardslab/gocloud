package gce

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/cloudlibz/gocloud/googleauth"
	"io/ioutil"
	"net/http"
)

//create gce instance

func (gce *GCE) CreateNode(request interface{}) (resp interface{}, err error) {
	var gceinstance GCE
	var projectid string
	var Zone string
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "projectid":
			projectid, _ = value.(string)

		case "Zone":
			zoneV, _ := value.(string)
			gceinstance.Zone = zoneV
			Zone = zoneV

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
	urlv := "https://www.googleapis.com/compute/v1/projects/" + projectid + "/zones/" + Zone + "/instances"
	CreateNoderequest, err := http.NewRequest("POST", urlv, bytes.NewBuffer(gceinstancejsonstringbyte))
	CreateNoderequest.Header.Set("Content-Type", "application/json")

	CreateNoderesp, err := client.Do(CreateNoderequest)
	fmt.Println(err)
	defer CreateNoderesp.Body.Close()

	body, err := ioutil.ReadAll(CreateNoderesp.Body)

	CreateNoderesponse := make(map[string]interface{})
	CreateNoderesponse["status"] = CreateNoderesp.StatusCode
	CreateNoderesponse["body"] = string(body)
	resp = CreateNoderesponse
	return resp, err
}

func (gce *GCE) StartNode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)
	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/instances/" + options["instance"] + "/start"
	client := googleauth.SignJWT()

	StartNoderequest, err := http.NewRequest("POST", url, nil)
	StartNoderequest.Header.Set("Content-Type", "application/json")

	StartNoderesp, err := client.Do(StartNoderequest)

	defer StartNoderesp.Body.Close()

	body, err := ioutil.ReadAll(StartNoderesp.Body)
	StartNoderesponse := make(map[string]interface{})
	StartNoderesponse["status"] = StartNoderesp.StatusCode
	StartNoderesponse["body"] = string(body)
	resp = StartNoderesponse

	return resp, err
}

//stop gce instance currentnly running
//accept projectid, zone, instance

func (gce *GCE) StopNode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)
	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/instances/" + options["instance"] + "/stop"
	client := googleauth.SignJWT()

	StopNoderequest, err := http.NewRequest("POST", url, nil)
	StopNoderequest.Header.Set("Content-Type", "application/json")

	StopNoderesp, err := client.Do(StopNoderequest)
	defer StopNoderesp.Body.Close()

	body, err := ioutil.ReadAll(StopNoderesp.Body)

	StopNoderesponse := make(map[string]interface{})
	StopNoderesponse["status"] = StopNoderesp.StatusCode
	StopNoderesponse["body"] = string(body)
	resp = StopNoderesponse

	return resp, err
}

//delete gce instance currentnly running
//accept projectid, zone, instance

func (gce *GCE) DeleteNode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)
	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/instances/" + options["instance"]
	client := googleauth.SignJWT()

	DeleteNoderequest, err := http.NewRequest("DELETE", url, nil)
	DeleteNoderequest.Header.Set("Content-Type", "application/json")

	DeleteNoderesp, err := client.Do(DeleteNoderequest)

	defer DeleteNoderesp.Body.Close()
	body, err := ioutil.ReadAll(DeleteNoderesp.Body)

	DeleteNoderesponse := make(map[string]interface{})
	DeleteNoderesponse["status"] = DeleteNoderesp.StatusCode
	DeleteNoderesponse["body"] = string(body)
	resp = DeleteNoderesponse

	return resp, err
}

//reboot/reset gce instance currentnly ***running***
//accept projectid, zone, instance

func (gce *GCE) RebootNode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)
	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/instances/" + options["instance"] + "/reset"
	client := googleauth.SignJWT()

	RebootNoderequest, err := http.NewRequest("POST", url, nil)
	RebootNoderequest.Header.Set("Content-Type", "application/json")

	RebootNoderesp, err := client.Do(RebootNoderequest)

	defer RebootNoderesp.Body.Close()

	body, err := ioutil.ReadAll(RebootNoderesp.Body)

	RebootNoderesponse := make(map[string]interface{})
	RebootNoderesponse["status"] = RebootNoderesp.StatusCode
	RebootNoderesponse["body"] = string(body)
	resp = RebootNoderesponse

	return resp, err
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

	listnoderesponse := make(map[string]interface{})
	listnoderesponse["status"] = listnoderesp.StatusCode
	listnoderesponse["body"] = string(body)
	resp = listnoderesponse

	return resp, err
}
