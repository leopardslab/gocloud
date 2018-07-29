package aliloadbalancer

import "encoding/json"

type CreateLoadBalancerResp struct {
	StatusCode       int
	LoadBalancerId   string
	Address          string
	NetworkType      string
	MasterZoneId     string
	SlaveZoneId      string
	LoadBalancerName string
}

type AttachLoadBalancerResp struct {
	StatusCode     int
	LoadBalancerId string
	BackendServers struct {
		BackendServer []BackendServerInfo
	}
}

type DetachLoadBalancerResp struct {
	StatusCode     int
	LoadBalancerId string
	BackendServers struct {
		BackendServer []BackendServerInfo
	}
}

type BackendServerInfo struct {
	ServerId string
	Weight   int
	Type     string
}

func ParseCreateLoadBalancerResp(resp interface{}) (createLoadBalancerResp CreateLoadBalancerResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &createLoadBalancerResp)
	createLoadBalancerResp.StatusCode = response["status"].(int)
	return
}

func ParseAttachLoadBalancerResp(resp interface{}) (attachLoadBalancerResp AttachLoadBalancerResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &attachLoadBalancerResp)
	attachLoadBalancerResp.StatusCode = response["status"].(int)
	return
}

func ParseDetachLoadBalancerResp(resp interface{}) (detachLoadBalancerResp DetachLoadBalancerResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &detachLoadBalancerResp)
	detachLoadBalancerResp.StatusCode = response["status"].(int)
	return
}
