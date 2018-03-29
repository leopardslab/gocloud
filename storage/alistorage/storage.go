package alistorage

import (
	"strconv"
	"github.com/cloudlibz/gocloud/aliauth"
)

// Createdisk create ECS-Disk accept map[string]interface{}
func (aliStorage *Alistorage) Createdisk(request interface{}) (resp interface{}, err error) {
	var options CreateDisk

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "RegionId":
			regionID, _ := value.(string)
			options.RegionID = regionID
		case "ZoneId":
			zoneID, _ := value.(string)
			options.ZoneID = zoneID
		case "DiskName":
			diskName, _ := value.(string)
			options.DiskName = diskName
		case "Description":
			description, _ := value.(string)
			options.Description = description
		case "Encrypted":
			switch value.(type) {
			case bool:
				options.Encrypted = value.(bool)
			case string:
				options.Encrypted = value.(string) == "true" || value.(string) == "True"
			}
		case "DiskCategory":
			diskCategory, _ := value.(string)
			options.DiskCategory = diskCategory
		case "Size":
			switch value.(type) {
			case int:
				options.Size = value.(int)
			case string:
				options.Size, _ = strconv.Atoi(value.(string))
			}
		case "SnapshotId":
			snapshotId, _ := value.(string)
			options.SnapshotId = snapshotId
		case "ClientToken":
			clientToken, _ := value.(string)
			options.ClientToken = clientToken
		}
	}

	// Put all of options into params
	params := aliauth.PutStructToMap(&options)

	response := make(map[string]interface{})
	err = aliauth.SignAndDoRequest("CreateDisk", params, response)
	resp = response
	return resp, err
}

//TODO
func (aliStorage *Alistorage) Deletedisk(request interface{}) (resp interface{}, err error) {
	return resp, err
}

//TODO
func (aliStorage *Alistorage) Createsnapshot(request interface{}) (resp interface{}, err error) {
	return resp, err
}

//TODO
func (aliStorage *Alistorage) Deletesnapshot(request interface{}) (resp interface{}, err error) {
	return resp, err
}

//TODO
func (aliStorage *Alistorage) Attachdisk(request interface{}) (resp interface{}, err error) {
	return resp, err
}

//TODO
func (aliStorage *Alistorage) Detachdisk(request interface{}) (resp interface{}, err error) {
	return resp, err
}
