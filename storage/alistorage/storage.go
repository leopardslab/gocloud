package alistorage

import (
	"strconv"
	"reflect"
)

//TODO
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

	params := make(map[string]interface{})

	// Put all of options into params
	e := reflect.ValueOf(&options).Elem()
	typeOfOptions := e.Type()
	for i := 0; i < e.NumField(); i++ {
		switch e.Field(i).Type().String() {
		case "string":
			if e.Field(i).Interface() != "" {
				params[typeOfOptions.Field(i).Name] = e.Field(i).Interface()
			}
		case "bool":
			params[typeOfOptions.Field(i).Name] = e.Field(i).Interface()
		case "int":
			if e.Field(i).Interface() != 0 {
				params[typeOfOptions.Field(i).Name] = e.Field(i).Interface()
			}
		}
	}

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
