package vultrstorage

import (
	"github.com/cloudlibz/gocloud/vultrauth"
	"net/http"
)

// CreateDisk function creates a new disk.
func (vultrStorage *VultrStorage) CreateDisk(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/block/create", param, response)

	resp = response
	return resp, err
}

// DeleteDisk function deletes a disk.
func (vultrStorage *VultrStorage) DeleteDisk(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/block/delete", param, response)

	resp = response
	return resp, err
}

// CreateSnapshot function creates a new snapshot.
func (vultrStorage *VultrStorage) CreateSnapshot(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/snapshot/create", param, response)

	resp = response
	return resp, err
}

// DeleteSnapshot function deletes a snapshot.
func (vultrStorage *VultrStorage) DeleteSnapshot(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/snapshot/destroy", param, response)

	resp = response
	return resp, err
}

// AttachDisk function function attaches a disk to a Vultr server.
func (vultrStorage *VultrStorage) AttachDisk(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/block/attach", param, response)

	resp = response
	return resp, err
}

// DetachDisk function function detaches a disk from a Vultr server.
func (vultrStorage *VultrStorage) DetachDisk(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/block/detach", param, response)

	resp = response
	return resp, err
}
