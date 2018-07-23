package vultrcompute

import "encoding/json"

type CreateNodeResp struct {
	SUBID string
}

func ParseCreateNodeResp(body interface{}) (createNodeResp CreateNodeResp, err error) {
	err = json.Unmarshal([]byte(body.(string)), &createNodeResp)
	return
}
