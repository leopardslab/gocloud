package vultrbaremetal

import (
	"encoding/json"
)

type V6Network struct {
	V6Network     string `json:"v6_network"`
	V6MainIP      string `json:"v6_main_ip"`
	V6NetworkSize int    `json:"v6_network_size"`
}

type BareMetalInfo struct {
	SUBID           string
	OS              string      `json:"os"`
	RAM             string      `json:"ram"`
	Disk            string      `json:"disk"`
	MainIP          string      `json:"main_ip"`
	CPUCount        float64     `json:"cpu_count"`
	Location        string      `json:"location"`
	DCID            string
	DefaultPassword string      `json:"default_password"`
	DateCreated     string      `json:"date_created"`
	Status          string      `json:"status"`
	NetmaskV4       string      `json:"netmask_v4"`
	GatewayV4       string      `json:"gateway_v4"`
	METALPLANID     float64
	V6Networks      []V6Network `json:"v6_networks"`
	Label           string      `json:"label"`
	Tag             string      `json:"tag"`
	OSID            string
	APPID           string
}

func ParseListBareMetalResp(body interface{}) (listBareMetalResp []BareMetalInfo, err error) {
	respMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(body.(string)), &respMap)
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
			}
		}
		listBareMetalResp = append(listBareMetalResp, *bareMetalInfo)
	}
	return
}
