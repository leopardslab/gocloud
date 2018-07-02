package amazonstorage

import (
	"strconv"
)

func (amazonstorage *Amazonstorage) CreateDisk(request interface{}) (resp interface{}, err error) {

	param, _ := request.(map[string]interface{})

	var createvolume CreateVolume
	var Region string
	for key, value := range param {
		switch key {

		case "Region":
			regionV, _ := value.(string)
			Region = regionV

		case "AvailZone":
			AvailZoneV, _ := value.(string)
			createvolume.AvailZone = AvailZoneV

		case "VolumeType":
			VolumeTypeV, _ := value.(string)
			createvolume.VolumeType = VolumeTypeV

		case "VolumeSize":
			VolumeSizeV, _ := value.(int)
			createvolume.VolumeSize = VolumeSizeV

		case "IOPS":
			IOPSV, _ := value.(int64)
			createvolume.IOPS = IOPSV

		case "Encrypted":
			EncryptedV, _ := value.(bool)
			createvolume.Encrypted = EncryptedV

		case "SnapshotId":
			SnapshotIdV, _ := value.(string)
			createvolume.SnapshotId = SnapshotIdV
		}
	}

	params := makeParams("CreateVolume")
	prepareVolume(params, createvolume)

	response := make(map[string]interface{})

	err = amazonstorage.PrepareSignatureV2query(params, Region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, nil

}

func (amazonstorage *Amazonstorage) DeleteDisk(request interface{}) (resp interface{}, err error) {
	param, _ := request.(map[string]string)
	params := makeParams("DeleteVolume")
	params["VolumeId"] = param["VolumeId"]
	Region := param["Region"]

	response := make(map[string]interface{})

	err = amazonstorage.PrepareSignatureV2query(params, Region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, nil

}

//volumeId, description string
func (amazonstorage *Amazonstorage) CreateSnapshot(request interface{}) (resp interface{}, err error) {

	param, _ := request.(map[string]string)
	params := makeParams("CreateSnapshot")

	params["VolumeId"] = param["VolumeId"]
	params["Description"] = param["Description"]
	Region := param["Region"]
	response := make(map[string]interface{})

	err = amazonstorage.PrepareSignatureV2query(params, Region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, nil
}

func (amazonstorage *Amazonstorage) DeleteSnapshot(request interface{}) (resp interface{}, err error) {
	ids := []string{}
	param, _ := request.(map[string]string)
	ids = append(ids, param["SnapshotId"])
	params := makeParams("DeleteSnapshot")
	Region := param["Region"]
	for i, id := range ids {
		params["SnapshotId."+strconv.Itoa(i+1)] = id
	}

	response := make(map[string]interface{})

	err = amazonstorage.PrepareSignatureV2query(params, Region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, nil

}

func (amazonstorage *Amazonstorage) AttachDisk(request interface{}) (resp interface{}, err error) {
	param, _ := request.(map[string]string)
	params := makeParams("AttachVolume")
	params["VolumeId"] = param["VolumeId"]
	params["InstanceId"] = param["InstanceId"]
	params["Device"] = param["Device"]
	Region := param["Region"]
	response := make(map[string]interface{})

	err = amazonstorage.PrepareSignatureV2query(params, Region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, nil
}

func (amazonstorage *Amazonstorage) DetachDisk(request interface{}) (resp interface{}, err error) {
	param, _ := request.(map[string]string)
	params := makeParams("DetachVolume")
	params["VolumeId"] = param["VolumeId"]
	params["InstanceId"] = param["InstanceId"]
	params["Device"] = param["Device"]
	if param["Force"] == "true" {
		params["Force"] = "true"
	}
	Region := param["Region"]
	response := make(map[string]interface{})

	err = amazonstorage.PrepareSignatureV2query(params, Region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, nil
}
