package droplet

import (
  digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
  "bytes"
  "fmt"
  "io/ioutil"
	"net/http"
	"strconv"
)

// dropletBasePath is the endpoint URL for digitalocean API.
const dropletBasePath = "https://api.digitalocean.com/v2/droplets"

// Createnode function creates a new droplet.
func (droplet *Droplet) Createnode(request interface{}) (resp interface{}, err error) {

  var DropletInstance Droplet // Initialize Droplet struct
  DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

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

  Createnodereq, err := http.NewRequest("POST", dropletBasePath, bytes.NewBuffer(dropletInstanceJSONStringbyte))
  if err != nil {
    fmt.Println(err)
  }
  Createnodereq.Header.Set("Content-Type", "application/json")
  Createnodereq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

	Createnoderesp, err := http.DefaultClient.Do(Createnodereq)
	if err != nil {
		fmt.Println(err)
	}

	defer Createnoderesp.Body.Close()

	responseBody, err := ioutil.ReadAll(Createnoderesp.Body)
	Createnoderesponse := make(map[string]interface{})
	Createnoderesponse["status"] = Createnoderesp.StatusCode
	Createnoderesponse["body"] = string(responseBody)
	resp = Createnoderesponse

	return resp, err

}

// Startnode function starts a droplet.
func (droplet *Droplet) Startnode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)
	inputID, err := strconv.Atoi(options["ID"])
	if err != nil {
		fmt.Println(err)
	}
	if inputID < 1 {
		return nil, errors.New("dropletID cannot be less than 1")
	}
	url := dropletBasePath + "/" + options["ID"] + "/actions"
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

	startRequest := &ActionRequest{"type": "power_on"}
	startRequestJSON, _ := json.Marshal(startRequest)
	startRequestJSONString := string(startRequestJSON)
	var startRequestJSONStringbyte = []byte(startRequestJSONString)

	Startnodereq, err := http.NewRequest("POST", url, bytes.NewBuffer(startRequestJSONStringbyte))
	if err != nil {
		fmt.Println(err)
	}
	Startnodereq.Header.Set("Content-Type", "application/json")
	Startnodereq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

	Startnoderesp, err := http.DefaultClient.Do(Startnodereq)
	if err != nil {
		fmt.Println(err)
	}
	defer Startnoderesp.Body.Close()

	responseBody, err := ioutil.ReadAll(Startnoderesp.Body)
	Startnoderesponse := make(map[string]interface{})
	Startnoderesponse["body"] = string(responseBody)
	resp = Startnoderesponse

	return resp, err

}

// Stopnode function stops a droplet.
func (droplet *Droplet) Stopnode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)
	inputID, err := strconv.Atoi(options["ID"])
	if err != nil {
		fmt.Println(err)
	}
	if inputID < 1 {
		return nil, errors.New("dropletID cannot be less than 1")
	}
	url := dropletBasePath + "/" + options["ID"] + "/actions"
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

	stopRequest := &ActionRequest{"type": "power_off"}
	stopRequestJSON, _ := json.Marshal(stopRequest)
	stopRequestJSONString := string(stopRequestJSON)
	var stopRequestJSONStringbyte = []byte(stopRequestJSONString)

	Stopnodereq, err := http.NewRequest("POST", url, bytes.NewBuffer(stopRequestJSONStringbyte))
	if err != nil {
		fmt.Println(err)
	}
	Stopnodereq.Header.Set("Content-Type", "application/json")
	Stopnodereq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

	Stopnoderesp, err := http.DefaultClient.Do(Stopnodereq)
	if err != nil {
		fmt.Println(err)
	}
	defer Stopnoderesp.Body.Close()

	responseBody, err := ioutil.ReadAll(Stopnoderesp.Body)
	Stopnoderesponse := make(map[string]interface{})
	Stopnoderesponse["body"] = string(responseBody)
	resp = Stopnoderesponse

	return resp, err

}

// Rebootnode function reboots a droplet.
func (droplet *Droplet) Rebootnode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)
	inputID, err := strconv.Atoi(options["ID"])
	if err != nil {
		fmt.Println(err)
	}
	if inputID < 1 {
		return nil, errors.New("dropletID cannot be less than 1")
	}
	url := dropletBasePath + "/" + options["ID"] + "/actions"
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

	rebootRequest := &ActionRequest{"type": "reboot"}
	rebootRequestJSON, _ := json.Marshal(rebootRequest)
	rebootRequestJSONString := string(rebootRequestJSON)
	var rebootRequestJSONStringbyte = []byte(rebootRequestJSONString)

	Rebootnodereq, err := http.NewRequest("POST", url, bytes.NewBuffer(rebootRequestJSONStringbyte))
	if err != nil {
		fmt.Println(err)
	}
	Rebootnodereq.Header.Set("Content-Type", "application/json")
	Rebootnodereq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

	Rebootnoderesp, err := http.DefaultClient.Do(Rebootnodereq)
	if err != nil {
		fmt.Println(err)
	}
	defer Rebootnoderesp.Body.Close()

	responseBody, err := ioutil.ReadAll(Rebootnoderesp.Body)
	Rebootnoderesponse := make(map[string]interface{})
	Rebootnoderesponse["status"] = Rebootnoderesp.StatusCode
	Rebootnoderesponse["body"] = string(responseBody)
	resp = Rebootnoderesponse

	return resp, err

}

// Deletenode function deletes a droplet.
func (droplet *Droplet) Deletenode(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)
	inputID, err := strconv.Atoi(options["ID"])
	if err != nil {
		fmt.Println(err)
	}
	if inputID < 1 {
		return nil, errors.New("dropletID cannot be less than 1")
	}

	url := dropletBasePath + "/" + options["ID"]
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

	Deletenoderequest, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	Deletenoderequest.Header.Set("Content-Type", "application/json")
	Deletenoderequest.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

	Deletenoderesp, err := http.DefaultClient.Do(Deletenoderequest)
	if err != nil {
		fmt.Println(err)
	}

	defer Deletenoderesp.Body.Close()

	responseBody, err := ioutil.ReadAll(Deletenoderesp.Body)
	Deletenoderesponse := make(map[string]interface{})
	Deletenoderesponse["status"] = Deletenoderesp.StatusCode
	Deletenoderesponse["body"] = string(responseBody)
	resp = Deletenoderesponse

	return resp, err
}
