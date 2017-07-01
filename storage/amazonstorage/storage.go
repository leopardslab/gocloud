package amazonstorage

import (
	"fmt"
	"strconv"
)

func (amazonstorage *Amazonstorage) Createdisk(request interface{}) (resp interface{}, err error) {

	param, _ := request.(map[string]interface{})

	fmt.Println(param)

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

	fmt.Println(createvolume)

	volume1 := CreateVolume{
		AvailZone:  "us-east-1a",
		VolumeType: "gp2",
		VolumeSize: 100,
		IOPS:       3000,
		Encrypted:  true,
	}
	fmt.Println(volume1)

	params := makeParams("CreateVolume")
	prepareVolume(params, createvolume)
	resp = &CreateVolumeResp{}
	err = amazonstorage.PrepareSignatureV2query(params, Region, resp)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	return resp, nil
}

func (amazonstorage *Amazonstorage) Deletedisk(request interface{}) (resp interface{}, err error) {
	param, _ := request.(map[string]string)
	params := makeParams("DeleteVolume")
	params["VolumeId"] = param["VolumeId"]
	Region := param["Region"]
	resp = &SimpleResp{}
	err = amazonstorage.PrepareSignatureV2query(params, Region, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//volumeId, description string
func (amazonstorage *Amazonstorage) Createsnapshot(request interface{}) (resp interface{}, err error) {

	param, _ := request.(map[string]string)
	params := makeParams("CreateSnapshot")

	params["VolumeId"] = param["VolumeId"]
	params["Description"] = param["Description"]
	Region := param["Region"]
	resp = &CreateSnapshotResp{}
	err = amazonstorage.PrepareSignatureV2query(params, Region, resp)
	if err != nil {
		return nil, err
	}
	return
}

func (amazonstorage *Amazonstorage) Deletesnapshot(request interface{}) (resp interface{}, err error) {
	ids := []string{}
	param, _ := request.(map[string]string)
	ids = append(ids, param["SnapshotId"])
	params := makeParams("DeleteSnapshot")
	Region := param["Region"]
	for i, id := range ids {
		params["SnapshotId."+strconv.Itoa(i+1)] = id
	}

	resp = &SimpleResp{}
	err = amazonstorage.PrepareSignatureV2query(params, Region, resp)
	if err != nil {
		return nil, err
	}
	return
}

func (amazonstorage *Amazonstorage) Attachdisk(request interface{}) (resp interface{}, err error) {
	param, _ := request.(map[string]string)
	params := makeParams("AttachVolume")
	params["VolumeId"] = param["VolumeId"]
	params["InstanceId"] = param["InstanceId"]
	params["Device"] = param["Device"]
	Region := param["Region"]
	resp = &VolumeAttachmentResp{}
	err = amazonstorage.PrepareSignatureV2query(params, Region, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (amazonstorage *Amazonstorage) Detachdisk(request interface{}) (resp interface{}, err error) {
	param, _ := request.(map[string]string)
	params := makeParams("DetachVolume")
	params["VolumeId"] = param["VolumeId"]
	params["InstanceId"] = param["InstanceId"]
	params["Device"] = param["Device"]
	if param["Force"] == "true" {
		params["Force"] = "true"
	}
	Region := param["Region"]
	resp = &VolumeAttachmentResp{}
	err = amazonstorage.PrepareSignatureV2query(params, Region, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
