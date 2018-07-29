package alistorage

import "encoding/json"

type CreateDiskResp struct {
	StatusCode int
	DiskId     string
}

func ParseCreateDiskResp(resp interface{}) (createDiskResp CreateDiskResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &createDiskResp)
	createDiskResp.StatusCode = response["status"].(int)
	return
}
