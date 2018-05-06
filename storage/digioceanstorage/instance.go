package digioceanstorage

import (
  digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
  "bytes"
  "fmt"
  "io/ioutil"
	"net/http"
  "encoding/json"
)

// storageBasePath is the endpoint URL for digitalocean API.
const storageBasePath = "https://api.digitalocean.com/v2/volumes"

// Createdisk function creates a new disk.
func (digioceanstorage *Digioceanstorage) Createdisk(request interface{}) (resp interface{}, err error) {

	var storageInstance Digioceanstorage	// Initialize LoadBalancer struct
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

	param := make(map[string]interface{})
  param = request.(map[string]interface{})

	for key, value := range param {

    switch key {

      case "Name":
        name, _ := value.(string)
        storageInstance.Name = name

      case "Region":
        region, _ := value.(string)
        storageInstance.Region = region

      case "Description":
        description, _ := value.(string)
        storageInstance.Description = description

      case "SizeGigaBytes":
        sizeGigaBytes, _ := value.(int64)
        storageInstance.SizeGigaBytes = sizeGigaBytes

      case "SnapshotID":
        snapshotID, _ := value.(string)
        storageInstance.SnapshotID = snapshotID

    } // Closes switch-case

  } // Closes for loop

	storageInstanceJSON, _ := json.Marshal(storageInstance)
  storageInstanceJSONString := string(storageInstanceJSON)
  var storageInstanceJSONStringbyte = []byte(storageInstanceJSONString)

  createStorageReq, err := http.NewRequest("POST", storageBasePath, bytes.NewBuffer(storageInstanceJSONStringbyte))
  if err != nil {
    fmt.Println(err)
  }
  createStorageReq.Header.Set("Content-Type", "application/json")
  createStorageReq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

	createStorageResp, err := http.DefaultClient.Do(createStorageReq)
	if err != nil {
		fmt.Println(err)
	}

	defer createStorageResp.Body.Close()

	responseBody, err := ioutil.ReadAll(createStorageResp.Body)
  createStorageResponse := make(map[string]interface{})
	createStorageResponse["status"] = createStorageResp.StatusCode
	createStorageResponse["body"] = string(responseBody)
	resp = createStorageResponse

	return resp, err
}

// Deletedisk function deletes a disk.
func (digioceanstorage *Digioceanstorage) Deletedisk(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := storageBasePath + "/" + options["VolumeID"]
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

	deleteSnapshotReq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	deleteSnapshotReq.Header.Set("Content-Type", "application/json")
	deleteSnapshotReq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

	deleteStorageResp, err := http.DefaultClient.Do(deleteSnapshotReq)
	if err != nil {
		fmt.Println(err)
	}

	defer deleteStorageResp.Body.Close()

	responseBody, err := ioutil.ReadAll(deleteStorageResp.Body)
	deleteStorageResponse := make(map[string]interface{})
	deleteStorageResponse["status"] = deleteStorageResp.StatusCode
	deleteStorageResponse["body"] = string(responseBody)
	resp = deleteStorageResponse

	return resp, err
}

// Createsnapshot function creates a new snapshot.
func (digioceanstorage *Digioceanstorage) Createsnapshot(request interface{}) (resp interface{}, err error) {

	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

	var volumeID string	// To store the volumeID
	var snapshotName string	// To store the name of the snapshot

	param := make(map[string]interface{})
  param = request.(map[string]interface{})

	for key, value := range param {

    switch key {
			case "VolumeID":
        volumeIDvalue, _ := value.(string)
        volumeID = volumeIDvalue

      case "SnapshotName":
        snapshotNameValue, _ := value.(string)
        snapshotName = snapshotNameValue
    } // Closes switch-case

  } // Closes for loop

	url := storageBasePath + "/" + volumeID + "/snapshots"

	snapshotname := map[string]string{
    "name": snapshotName,
  }

	storageInstanceJSON, _ := json.Marshal(snapshotname)
	storageInstanceJSONString := string(storageInstanceJSON)
	var storageInstanceJSONStringbyte = []byte(storageInstanceJSONString)

	createSnapshotReq, err := http.NewRequest("POST", url, bytes.NewBuffer(storageInstanceJSONStringbyte))
	if err != nil {
		fmt.Println(err)
	}
	createSnapshotReq.Header.Set("Content-Type", "application/json")
	createSnapshotReq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

	createSnapshotResp, err := http.DefaultClient.Do(createSnapshotReq)
	if err != nil {
		fmt.Println(err)
	}

	defer createSnapshotResp.Body.Close()

	responseBody, err := ioutil.ReadAll(createSnapshotResp.Body)
	createSnapshotResponse := make(map[string]interface{})
	createSnapshotResponse["status"] = createSnapshotResp.StatusCode
	createSnapshotResponse["body"] = string(responseBody)
	resp = createSnapshotResponse

	return resp, err
}

// Deletesnapshot function deletes a snapshot.
func (digioceanstorage *Digioceanstorage) Deletesnapshot(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://api.digitalocean.com/v2/snapshots/" + options["SnapshotID"]
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

	deleteSnapshotReq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	deleteSnapshotReq.Header.Set("Content-Type", "application/json")
	deleteSnapshotReq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

	deleteSnapshotResp, err := http.DefaultClient.Do(deleteSnapshotReq)
	if err != nil {
		fmt.Println(err)
	}

	defer deleteSnapshotResp.Body.Close()

	responseBody, err := ioutil.ReadAll(deleteSnapshotResp.Body)
	deleteSnapshotResponse := make(map[string]interface{})
	deleteSnapshotResponse["status"] = deleteSnapshotResp.StatusCode
	deleteSnapshotResponse["body"] = string(responseBody)
	resp = deleteSnapshotResponse

	return resp, err
}

// Attachdisk function function attaches a disk to a droplet.
func (digioceanstorage *Digioceanstorage) Attachdisk(request interface{}) (resp interface{}, err error) {

	var volumeID string	// To store the volumeID
	var dropletID int	// To store the Droplet the volume will be attached to.
	var region string	// To store the region

	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

	param := make(map[string]interface{})
  param = request.(map[string]interface{})

	for key, value := range param {

    switch key {
			case "VolumeID":
        volumeIDvalue, _ := value.(string)
        volumeID = volumeIDvalue

      case "DropletID":
        dropletIDvalue, _ := value.(int)
        dropletID = dropletIDvalue

			case "Region":
        regionvalue, _ := value.(string)
        region = regionvalue
    } // Closes switch-case

  } // Closes for loop

	url := storageBasePath + "/" + volumeID + "/actions"

	dropletinstances := map[string]interface{}{
		"type":	"attach",
    "droplet_id": dropletID,
		"region":	region,
  }

	storageInstanceJSON, _ := json.Marshal(dropletinstances)
  storageInstanceJSONString := string(storageInstanceJSON)
  var storageInstanceJSONStringbyte = []byte(storageInstanceJSONString)

  attachVolumetoDropletReq, err := http.NewRequest("POST", url, bytes.NewBuffer(storageInstanceJSONStringbyte))
  if err != nil {
    fmt.Println(err)
  }
  attachVolumetoDropletReq.Header.Set("Content-Type", "application/json")
  attachVolumetoDropletReq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

  attachVolumetoDropletResp, err := http.DefaultClient.Do(attachVolumetoDropletReq)
  if err != nil {
    fmt.Println(err)
  }

  defer attachVolumetoDropletResp.Body.Close()

  responseBody, err := ioutil.ReadAll(attachVolumetoDropletResp.Body)
  attachVolumetoDropletResponse := make(map[string]interface{})
  attachVolumetoDropletResponse["status"] = attachVolumetoDropletResp.StatusCode
  attachVolumetoDropletResponse["body"] = string(responseBody)
  resp = attachVolumetoDropletResponse

  return resp, err
}

// Detachdisk function function detaches a disk from a droplet.
func (digioceanstorage *Digioceanstorage) Detachdisk(request interface{}) (resp interface{}, err error) {

	var volumeID string	// To store the volumeID
	var dropletID int	// To store the Droplet the volume will be attached to.
	var region string	// To store the region

	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken  // Fetch the DigiOceanAccessToken

	param := make(map[string]interface{})
  param = request.(map[string]interface{})

	for key, value := range param {

    switch key {
			case "VolumeID":
        volumeIDvalue, _ := value.(string)
        volumeID = volumeIDvalue

      case "DropletID":
        dropletIDvalue, _ := value.(int)
        dropletID = dropletIDvalue

			case "Region":
        regionvalue, _ := value.(string)
        region = regionvalue
    } // Closes switch-case

  } // Closes for loop

	url := storageBasePath + "/" + volumeID + "/actions"

	dropletinstances := map[string]interface{}{
		"type":	"detach",
    "droplet_id": dropletID,
		"region":	region,
  }

	storageInstanceJSON, _ := json.Marshal(dropletinstances)
  storageInstanceJSONString := string(storageInstanceJSON)
  var storageInstanceJSONStringbyte = []byte(storageInstanceJSONString)

  detachVolumetoDropletReq, err := http.NewRequest("POST", url, bytes.NewBuffer(storageInstanceJSONStringbyte))
  if err != nil {
    fmt.Println(err)
  }
  detachVolumetoDropletReq.Header.Set("Content-Type", "application/json")
  detachVolumetoDropletReq.Header.Set("Authorization", "Bearer " + DigiOceanAccessToken)

  detachVolumetoDropletResp, err := http.DefaultClient.Do(detachVolumetoDropletReq)
  if err != nil {
    fmt.Println(err)
  }

  defer detachVolumetoDropletResp.Body.Close()

  responseBody, err := ioutil.ReadAll(detachVolumetoDropletResp.Body)
  detachVolumetoDropletResponse := make(map[string]interface{})
  detachVolumetoDropletResponse["status"] = detachVolumetoDropletResp.StatusCode
  detachVolumetoDropletResponse["body"] = string(responseBody)
  resp = detachVolumetoDropletResponse

  return resp, err
}
