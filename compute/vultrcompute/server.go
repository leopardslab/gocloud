package vultrcompute

import (
	"github.com/cloudlibz/gocloud/vultrauth"
	"net/http"
	"fmt"
)

// Createnode function creates a new VultrCompute instance.
func (vultrCompute *VultrCompute) Createnode(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/server/create", param, response)

	resp = response
	return resp, err
}

// Startnode function starts a VultrCompute instance.
func (vultrCompute *VultrCompute) Startnode(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/server/start", param, response)

	resp = response
	return resp, err
}

// Stopnode function stops a VultrCompute instance.
func (vultrCompute *VultrCompute) Stopnode(request interface{}) (resp interface{}, err error) {
	fmt.Println("\nThis API is not provided by Vultr cloud")
	return resp, err
}

// Rebootnode function reboots a VultrCompute instance.
func (vultrCompute *VultrCompute) Rebootnode(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/server/reboot", param, response)

	resp = response
	return resp, err
}

// Deletenode function deletes a VultrCompute instance.
func (vultrCompute *VultrCompute) Deletenode(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/server/destroy", param, response)

	resp = response
	return resp, err
}

// ListNode function lists VultrCompute instances.
func (vultrCompute *VultrCompute) ListNode() (resp interface{}, err error) {
	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodGet, "/v1/server/list", nil, response)

	resp = response
	return resp, err
}
