package vultrstorage

import (
	"github.com/cloudlibz/gocloud/vultrauth"
	"net/http"
)

// Createdisk function creates a new disk.
func (vultrStorage *VultrStorage) Createdisk(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/block/create", param, response)

	resp = response
	return resp, err
}

// Deletedisk function deletes a disk.
func (vultrStorage *VultrStorage) Deletedisk(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/block/delete", param, response)

	resp = response
	return resp, err
}

// Createsnapshot function creates a new snapshot.
func (vultrStorage *VultrStorage) Createsnapshot(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/snapshot/create", param, response)

	resp = response
	return resp, err
}

// Deletesnapshot function deletes a snapshot.
func (vultrStorage *VultrStorage) Deletesnapshot(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/snapshot/destroy", param, response)

	resp = response
	return resp, err
}

// Attachdisk function function attaches a disk to a Vultr server.
func (vultrStorage *VultrStorage) Attachdisk(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/block/attach", param, response)

	resp = response
	return resp, err
}

// Detachdisk function function detaches a disk from a Vultr server.
func (vultrStorage *VultrStorage) Detachdisk(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/block/detach", param, response)

	resp = response
	return resp, err
}
