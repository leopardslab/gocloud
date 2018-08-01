package alicontainer

import "encoding/json"

type CreateClusterResp struct {
	StatusCode int
	ClusterID  string `json:"cluster_id"`
	TaskID     string `json:"task_id"`
}

func ParseCreateClusterResp(resp interface{}) (createClusterResp CreateClusterResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &createClusterResp)
	createClusterResp.StatusCode = response["status"].(int)
	return
}
