package vultrcompute

import (
	"fmt"
	"github.com/cloudlibz/gocloud/vultrauth"
	"net/http"
)

// CreateNode function creates a new VultrCompute instance.
func (vultrCompute *VultrCompute) CreateNode(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/server/create", param, response)

	resp = response
	return resp, err
}

// StartNode function starts a VultrCompute instance.
func (vultrCompute *VultrCompute) StartNode(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/server/start", param, response)

	resp = response
	return resp, err
}

// StopNode function stops a VultrCompute instance.
func (vultrCompute *VultrCompute) StopNode(request interface{}) (resp interface{}, err error) {
	fmt.Println("\nThis API is not provided by Vultr cloud")
	return resp, err
}

// RebootNode function reboots a VultrCompute instance.
func (vultrCompute *VultrCompute) RebootNode(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/server/reboot", param, response)

	resp = response
	return resp, err
}

// DeleteNode function deletes a VultrCompute instance.
func (vultrCompute *VultrCompute) DeleteNode(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/server/destroy", param, response)

	resp = response
	return resp, err
}

// ListNode function lists VultrCompute instances.
func (vultrCompute *VultrCompute) ListNode(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	if request != nil {
		param = request.(map[string]interface{})
	}

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodGet, "/v1/server/list", param, response)

	resp = response
	return resp, err
}
