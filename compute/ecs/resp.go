package ecs

import (
	"encoding/json"
)

type CreateNodeResp struct {
	InstanceId string
}

func ParseCreateNodeResp(body interface{}) (createNodeResp CreateNodeResp, err error) {
	err = json.Unmarshal([]byte(body.(string)), &createNodeResp)
	return
}