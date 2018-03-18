package droplet

import (
  digioceanauth "github.com/cloudlibz/gocloud/digioceanauth"
  "bytes"
  "fmt"
  "io/ioutil"
	"net/http"
  "encoding/json"
)

const dropletBasePath = "https://api.digitalocean.com/v2/droplets"

func (droplet *Droplet) Createnode(request interface{}) (resp interface{}, err error) {

  var DropletInstance DropletCreateRequest // Initialize DropletCreateRequest struct
  AccessToken := digioceanauth.Token.AccessToken  // Fetch the AccessToken

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
        var dropletCreateSSHKey DropletCreateSSHKey
        dropletCreateSSHKey.ID = sshkeyparam[i]["ID"]
        dropletCreateSSHKey.Fingerprint = sshkeyparam[i]["Fingerprint"]
        DropletInstance.SSHKeys = append(DropletInstance.SSHKeys, dropletCreateSSHKey)
      }

    case "Volumes":
      volumeparam, _ := value.([]map[string]string)
      for i := 0; i < len(volumeparam); i++ {
        var dropletCreateVolume DropletCreateVolume
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
      private_networking, _ := value.(bool)
      DropletInstance.PrivateNetworking = private_networking

    case "Monitoring":
      monitoring, _ := value.(bool)
      DropletInstance.Monitoring = monitoring

    case "UserData":
      user_data, _ := value.(string)
      DropletInstance.UserData = user_data

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
  Createnodereq.Header.Set("Authorization", "Bearer " + AccessToken)

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

} // Createnode function closes.

//TODO
func (droplet *Droplet) Startnode(request interface{}) (resp interface{}, err error) {
	return resp, err
}
//TODO
func (droplet *Droplet) Stopnode(request interface{}) (resp interface{}, err error) {
	return resp, err
}
//TODO
func (droplet *Droplet) Rebootnode(request interface{}) (resp interface{}, err error) {
	return resp, err
}
//TODO
func (droplet *Droplet) Deletenode(request interface{}) (resp interface{}, err error) {
	return resp, err
}
