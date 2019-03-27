package rackspacecompute

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudlibz/gocloud/rackspaceauth"
)

var cloudserverEndpoints = rackspaceauth.Token.Endpoints["cloudServer"]

// CreateNode function creates a new rackspacecompute instance.
func (rackspacecompute *Rackspacecompute) CreateNode(request interface{}) (resp interface{}, err error) {
	var cloudserverInstance Cloudserver // Initialize Cloudserver struct
	var region string
	rackSpaceAuthToken := rackspaceauth.Token.RackSpaceAuthToken // Fetch the DigiOceanAccessToken

	param := request.(map[string]interface{})

	for key, value := range param {

		switch key {

		case "Name":
			name, _ := value.(string)
			cloudserverInstance.Name = name

		case "ImageRef":
			imageparam, _ := value.(string)
			cloudserverInstance.ImageRef = imageparam

		case "FlavourRef":
			flavourparam, _ := value.(string)
			cloudserverInstance.FlavorRef = flavourparam

		case "Region":
			region, _ = value.(string)
		} // Closes switch-case

	} // Closes for loop

	rackspacecompute.Server = cloudserverInstance
	rackspacecomputeJSON, _ := json.Marshal(rackspacecompute)

	var cloudserverEndpoint string
	flag := 0
	for _, endpoint := range cloudserverEndpoints {
		if endpoint.Region == region {
			cloudserverEndpoint = endpoint.URL
			flag = 1
			break
		}
	}
	if flag == 0 {
		return nil, fmt.Errorf("Could not find a endpoint for the giver region : %s", region)
	}

	createNodereq, err := http.NewRequest("POST", cloudserverEndpoint, bytes.NewBuffer(rackspacecomputeJSON))
	if err != nil {
		fmt.Println(err)
	}
	createNodereq.Header.Set("Content-Type", "application/json")
	createNodereq.Header.Set("X-Auth-Token", rackSpaceAuthToken)

	createNoderesp, err := http.DefaultClient.Do(createNodereq)
	if err != nil {
		fmt.Println(err)
	}

	defer createNoderesp.Body.Close()

	responseBody, err := ioutil.ReadAll(createNoderesp.Body)
	createNoderesponse := make(map[string]interface{})
	createNoderesponse["status"] = createNoderesp.StatusCode
	createNoderesponse["body"] = string(responseBody)
	resp = createNoderesponse

	return resp, err
}

// StartNode function starts a rackspacecompute instance.
func (rackspacecompute *Rackspacecompute) StartNode(request interface{}) (resp interface{}, err error) {
	return resp, err
}

// StopNode function stops a rackspacecompute instance.
func (rackspacecompute *Rackspacecompute) StopNode(request interface{}) (resp interface{}, err error) {
	return resp, err
}

// RebootNode function reboots a rackspacecompute instance.
func (rackspacecompute *Rackspacecompute) RebootNode(request interface{}) (resp interface{}, err error) {
	return resp, err
}

// DeleteNode function deletes a rackspacecompute instance.
func (rackspacecompute *Rackspacecompute) DeleteNode(request interface{}) (resp interface{}, err error) {
	return resp, err
}
