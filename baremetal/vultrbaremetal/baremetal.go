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
	return
}

func (*VultrBareMetal) ReinstallBareMetal(request interface{}) (resp interface{}, err error) {
	return
}

func (*VultrBareMetal) HaltBareMetal(request interface{}) (resp interface{}, err error) {
	return
}

func (*VultrBareMetal) ListBareMetal(request interface{}) (resp interface{}, err error) {
	return
}
