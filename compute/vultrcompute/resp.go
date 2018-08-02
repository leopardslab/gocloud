package vultrcompute

import "encoding/json"

type CreateNodeResp struct {
	StatusCode int
	SUBID      string
}

func ParseCreateNodeResp(resp interface{}) (createNodeResp CreateNodeResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &createNodeResp)
	createNodeResp.StatusCode = response["status"].(int)
	return
}
