package vultrbaremetal

import (
	"github.com/cloudlibz/gocloud/vultrauth"
	"net/http"
)

func (*VultrBareMetal) CreateBareMetal(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/baremetal/create", param, response)

	resp = response
	return resp, err
}

func (*VultrBareMetal) DeleteBareMetal(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/baremetal/destroy", param, response)

	resp = response
	return resp, err
}

func (*VultrBareMetal) RebootBareMetal(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/baremetal/reboot", param, response)

	resp = response
	return resp, err
}

func (*VultrBareMetal) ReinstallBareMetal(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/baremetal/reinstall", param, response)

	resp = response
	return resp, err
}

func (*VultrBareMetal) HaltBareMetal(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/baremetal/halt", param, response)

	resp = response
	return resp, err
}

func (*VultrBareMetal) ListBareMetal(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})

	if request != nil {
		param = request.(map[string]interface{})
	}

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodGet, "/v1/baremetal/list", param, response)

	resp = response
	return resp, err
}
