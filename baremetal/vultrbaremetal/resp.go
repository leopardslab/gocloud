package vultrbaremetal

import (
	"encoding/json"
)

type V6Network struct {
	V6Network     string  `json:"v6_network"`
	V6MainIP      string  `json:"v6_main_ip"`
	V6NetworkSize float64 `json:"v6_network_size"`
}

type BareMetalInfo struct {
	SUBID           string
	OS              string  `json:"os"`
	RAM             string  `json:"ram"`
	Disk            string  `json:"disk"`
	MainIP          string  `json:"main_ip"`
	CPUCount        float64 `json:"cpu_count"`
	Location        string  `json:"location"`
	DCID            string
	DefaultPassword string `json:"default_password"`
	DateCreated     string `json:"date_created"`
	Status          string `json:"status"`
	NetmaskV4       string `json:"netmask_v4"`
	GatewayV4       string `json:"gateway_v4"`
	METALPLANID     float64
	V6Networks      []V6Network `json:"v6_networks"`
	Label           string      `json:"label"`
	Tag             string      `json:"tag"`
	OSID            string
	APPID           string
}

type CreateBareMetalResp struct {
	StatusCode int
	SUBID      string
}

type ListBareMetalResp struct {
	StatusCode     int
	BareMetalSlice []BareMetalInfo
}

func ParseCreateBareMetalResp(resp interface{}) (createBareMetalResp CreateBareMetalResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &createBareMetalResp)
	createBareMetalResp.StatusCode = response["status"].(int)
	return
}

func ParseListBareMetalResp(resp interface{}) (listBareMetalResp ListBareMetalResp, err error) {
	response := resp.(map[string]interface{})
	respMap := make(map[string]interface{})
	listBareMetalResp.StatusCode = response["status"].(int)
	err = json.Unmarshal([]byte(response["body"].(string)), &respMap)
	for _, bareMetal := range respMap {
		bareMetalInfo := &BareMetalInfo{}
		for key, value := range bareMetal.(map[string]interface{}) {
			switch key {
			case "SUBID":
				bareMetalInfo.SUBID = value.(string)
				break
			case "os":
				bareMetalInfo.OS = value.(string)
				break
			case "ram":
				bareMetalInfo.RAM = value.(string)
				break
			case "disk":
				bareMetalInfo.Disk = value.(string)
				break
			case "main_ip":
				bareMetalInfo.MainIP = value.(string)
				break
			case "cpu_count":
				bareMetalInfo.CPUCount = value.(float64)
				break
			case "location":
				bareMetalInfo.Location = value.(string)
				break
			case "DCID":
				bareMetalInfo.DCID = value.(string)
				break
			case "default_password":
				bareMetalInfo.DefaultPassword = value.(string)
				break
			case "date_created":
				bareMetalInfo.DateCreated = value.(string)
				break
			case "status":
				bareMetalInfo.Status = value.(string)
				break
			case "netmask_v4":
				bareMetalInfo.NetmaskV4 = value.(string)
				break
			case "gateway_v4":
				bareMetalInfo.GatewayV4 = value.(string)
				break
			case "METALPLANID":
				bareMetalInfo.METALPLANID = value.(float64)
				break
			case "label":
				bareMetalInfo.Label = value.(string)
				break
			case "tag":
				bareMetalInfo.Tag = value.(string)
				break
			case "OSID":
				bareMetalInfo.OSID = value.(string)
				break
			case "APPID":
				bareMetalInfo.APPID = value.(string)
				break
			case "v6_networks":
				_, ok := value.([]map[string]interface{})
				if !ok {
					break
				}
				for _, v6NetworksMap := range value.([]map[string]interface{}) {
					v6Network := &V6Network{}
					for key, value := range v6NetworksMap {
						switch key {
						case "v6_network":
							v6Network.V6Network = value.(string)
							break
						case "v6_main_ip":
							v6Network.V6MainIP = value.(string)
							break
						case "v6_network_size":
							v6Network.V6NetworkSize = value.(float64)
							break
						}
					}
					bareMetalInfo.V6Networks = append(bareMetalInfo.V6Networks, *v6Network)
				}
				break
			}
		}
		listBareMetalResp.BareMetalSlice = append(listBareMetalResp.BareMetalSlice, *bareMetalInfo)
	}
	return
}
