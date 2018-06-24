package droplet

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
	"io/ioutil"
	"net/http"
	"strconv"
)

// dropletBasePath is the endpoint URL for digitalocean API.
const dropletBasePath = "https://api.digitalocean.com/v2/droplets"

// CreateNode function creates a new droplet.
func (droplet *Droplet) CreateNode(request interface{}) (resp interface{}, err error) {

	var DropletInstance Droplet                                      // Initialize Droplet struct
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken // Fetch the DigiOceanAccessToken

	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	for key, value := range param {

		switch key {

		case "Name":
			name, _ := value.(string)
			DropletInstance.Name = name

		case "Region":
			region, _ := value.(string)
			DropletInstance.Region = region

		case "Size":
			size, _ := value.(string)
			DropletInstance.Size = size

		case "Image":
			imageparam, _ := value.(map[string]interface{})
			for key, value := range imageparam {
				switch key {
				case "ID":
					id, _ := value.(int)
					DropletInstance.Image.ID = id
					break
				case "Slug":
					slug, _ := value.(string)
					DropletInstance.Image.Slug = slug
					break
				}
			}

		case "SSHKeys":
			sshkeyparam, _ := value.([]map[string]string)
			for i := 0; i < len(sshkeyparam); i++ {
				var dropletCreateSSHKey CreateSSHKey
				dropletCreateSSHKey.ID = sshkeyparam[i]["ID"]
				dropletCreateSSHKey.Fingerprint = sshkeyparam[i]["Fingerprint"]
				DropletInstance.SSHKeys = append(DropletInstance.SSHKeys, dropletCreateSSHKey)
			}

		case "Volumes":
			volumeparam, _ := value.([]map[string]string)
			for i := 0; i < len(volumeparam); i++ {
				var dropletCreateVolume CreateVolume
				dropletCreateVolume.ID = volumeparam[i]["ID"]
				dropletCreateVolume.Name = volumeparam[i]["Name"]
				DropletInstance.Volumes = append(DropletInstance.Volumes, dropletCreateVolume)
			}

		case "Backups":
			backups, _ := value.(bool)
			DropletInstance.Backups = backups

		case "IPv6":
			ipv6, _ := value.(bool)
			DropletInstance.IPv6 = ipv6

		case "PrivateNetworking":
			privateNetworking, _ := value.(bool)
			DropletInstance.PrivateNetworking = privateNetworking

		case "Monitoring":
			monitoring, _ := value.(bool)
			DropletInstance.Monitoring = monitoring

		case "UserData":
			userData, _ := value.(string)
			DropletInstance.UserData = userData

		case "Tags":
			tags, _ := value.([]string)
			DropletInstance.Tags = tags

		} // Closes switch-case

	} // Closes for loop

	dropletInstanceJSON, _ := json.Marshal(DropletInstance)
	dropletInstanceJSONString := string(dropletInstanceJSON)
	var dropletInstanceJSONStringbyte = []byte(dropletInstanceJSONString)

	CreateNodereq, err := http.NewRequest("POST", dropletBasePath, bytes.NewBuffer(dropletInstanceJSONStringbyte))
	if err != nil {
		fmt.Println(err)
	}
	CreateNodereq.Header.Set("Content-Type", "application/json")
	CreateNodereq.Header.Set("Authorization", "Bearer "+DigiOceanAccessToken)

	CreateNoderesp, err := http.DefaultClient.Do(CreateNodereq)
	if err != nil {
		fmt.Println(err)
	}

	defer CreateNoderesp.Body.Close()

	responseBody, err := ioutil.ReadAll(CreateNoderesp.Body)
	CreateNoderesponse := make(map[string]interface{})
	CreateNoderesponse["status"] = CreateNoderesp.StatusCode
	CreateNoderesponse["body"] = string(responseBody)
	resp = CreateNoderesponse

	return resp, err

}

// StartNode function starts a droplet.
func (droplet *Droplet) StartNode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)
	inputID, err := strconv.Atoi(options["ID"])
	if err != nil {
		fmt.Println(err)
	}
	if inputID < 1 {
		return nil, errors.New("dropletID cannot be less than 1")
	}
	url := dropletBasePath + "/" + options["ID"] + "/actions"
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken // Fetch the DigiOceanAccessToken

	startRequest := &ActionRequest{"type": "power_on"}
	startRequestJSON, _ := json.Marshal(startRequest)
	startRequestJSONString := string(startRequestJSON)
	var startRequestJSONStringbyte = []byte(startRequestJSONString)

	StartNodereq, err := http.NewRequest("POST", url, bytes.NewBuffer(startRequestJSONStringbyte))
	if err != nil {
		fmt.Println(err)
	}
	StartNodereq.Header.Set("Content-Type", "application/json")
	StartNodereq.Header.Set("Authorization", "Bearer "+DigiOceanAccessToken)

	StartNoderesp, err := http.DefaultClient.Do(StartNodereq)
	if err != nil {
		fmt.Println(err)
	}
	defer StartNoderesp.Body.Close()

	responseBody, err := ioutil.ReadAll(StartNoderesp.Body)
	StartNoderesponse := make(map[string]interface{})
	StartNoderesponse["body"] = string(responseBody)
	resp = StartNoderesponse

	return resp, err

}

// StopNode function stops a droplet.
func (droplet *Droplet) StopNode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)
	inputID, err := strconv.Atoi(options["ID"])
	if err != nil {
		fmt.Println(err)
	}
	if inputID < 1 {
		return nil, errors.New("dropletID cannot be less than 1")
	}
	url := dropletBasePath + "/" + options["ID"] + "/actions"
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken // Fetch the DigiOceanAccessToken

	stopRequest := &ActionRequest{"type": "power_off"}
	stopRequestJSON, _ := json.Marshal(stopRequest)
	stopRequestJSONString := string(stopRequestJSON)
	var stopRequestJSONStringbyte = []byte(stopRequestJSONString)

	StopNodereq, err := http.NewRequest("POST", url, bytes.NewBuffer(stopRequestJSONStringbyte))
	if err != nil {
		fmt.Println(err)
	}
	StopNodereq.Header.Set("Content-Type", "application/json")
	StopNodereq.Header.Set("Authorization", "Bearer "+DigiOceanAccessToken)

	StopNoderesp, err := http.DefaultClient.Do(StopNodereq)
	if err != nil {
		fmt.Println(err)
	}
	defer StopNoderesp.Body.Close()

	responseBody, err := ioutil.ReadAll(StopNoderesp.Body)
	StopNoderesponse := make(map[string]interface{})
	StopNoderesponse["body"] = string(responseBody)
	resp = StopNoderesponse

	return resp, err

}

// RebootNode function reboots a droplet.
func (droplet *Droplet) RebootNode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)
	inputID, err := strconv.Atoi(options["ID"])
	if err != nil {
		fmt.Println(err)
	}
	if inputID < 1 {
		return nil, errors.New("dropletID cannot be less than 1")
	}
	url := dropletBasePath + "/" + options["ID"] + "/actions"
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken // Fetch the DigiOceanAccessToken

	rebootRequest := &ActionRequest{"type": "reboot"}
	rebootRequestJSON, _ := json.Marshal(rebootRequest)
	rebootRequestJSONString := string(rebootRequestJSON)
	var rebootRequestJSONStringbyte = []byte(rebootRequestJSONString)

	RebootNodereq, err := http.NewRequest("POST", url, bytes.NewBuffer(rebootRequestJSONStringbyte))
	if err != nil {
		fmt.Println(err)
	}
	RebootNodereq.Header.Set("Content-Type", "application/json")
	RebootNodereq.Header.Set("Authorization", "Bearer "+DigiOceanAccessToken)

	RebootNoderesp, err := http.DefaultClient.Do(RebootNodereq)
	if err != nil {
		fmt.Println(err)
	}
	defer RebootNoderesp.Body.Close()

	responseBody, err := ioutil.ReadAll(RebootNoderesp.Body)
	RebootNoderesponse := make(map[string]interface{})
	RebootNoderesponse["status"] = RebootNoderesp.StatusCode
	RebootNoderesponse["body"] = string(responseBody)
	resp = RebootNoderesponse

	return resp, err

}

// DeleteNode function deletes a droplet.
func (droplet *Droplet) DeleteNode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)
	inputID, err := strconv.Atoi(options["ID"])
	if err != nil {
		fmt.Println(err)
	}
	if inputID < 1 {
		return nil, errors.New("dropletID cannot be less than 1")
	}

	url := dropletBasePath + "/" + options["ID"]
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken // Fetch the DigiOceanAccessToken

	DeleteNoderequest, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	DeleteNoderequest.Header.Set("Content-Type", "application/json")
	DeleteNoderequest.Header.Set("Authorization", "Bearer "+DigiOceanAccessToken)

	DeleteNoderesp, err := http.DefaultClient.Do(DeleteNoderequest)
	if err != nil {
		fmt.Println(err)
	}

	defer DeleteNoderesp.Body.Close()

	responseBody, err := ioutil.ReadAll(DeleteNoderesp.Body)
	DeleteNoderesponse := make(map[string]interface{})
	DeleteNoderesponse["status"] = DeleteNoderesp.StatusCode
	DeleteNoderesponse["body"] = string(responseBody)
	resp = DeleteNoderesponse

	return resp, err
}
