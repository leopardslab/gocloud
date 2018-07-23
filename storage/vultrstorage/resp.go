package vultrstorage

import "encoding/json"

type CreateDiskResp struct {
	SUBID string `json:"SUBID"`
}

func ParseCreateDiskResp(body interface{}) (createDiskResp CreateDiskResp, err error) {
	err = json.Unmarshal([]byte(body.(string)), &createDiskResp)
	return
}
