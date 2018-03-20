package ecs

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"strconv"
	"reflect"
)

//TODO
func (ecs *ECS) Startnode(request interface{}) (resp interface{}, err error) {
	return resp, err
}
//TODO
func (ecs *ECS) Stopnode(request interface{}) (resp interface{}, err error) {
	return resp, err
}
//TODO
func (ecs *ECS) Rebootnode(request interface{}) (resp interface{}, err error) {
	return resp, err
}
//TODO
func (ecs *ECS) Deletenode(request interface{}) (resp interface{}, err error) {
	return resp, err
}

//Create Ec2 instances accept map[string]interface{}
func (ecs *ECS) Createnode(request interface{}) (resp interface{}, err error) {
	var options CreateInstance

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "RegionId":
			regionID, _ := value.(string)
			options.RegionID = regionID
		case "ImageId":
			imageID, _ := value.(string)
			options.ImageID = imageID
		case "InstanceType":
			instanceType, _ := value.(string)
			options.InstanceType = instanceType
		case "SecurityGroupId":
			securityGroupID, _ := value.(string)
			options.SecurityGroupID = securityGroupID
		case "ZoneId":
			zoneID, _ := value.(string)
			options.ZoneID = zoneID
		case "InstanceName":
			instanceName, _ := value.(string)
			options.InstanceName = instanceName
		case "Description":
			description, _ := value.(string)
			options.Description = description
		case "InternetChargeType":
			internetChargeType, _ := value.(string)
			options.InternetChargeType = internetChargeType
		case "InternetMaxBandwidthIn":
			internetMaxBandwidthIn, ok := value.(int)
			if !ok {
				internetMaxBandwidthIn, _ = strconv.Atoi(value.(string))
			}
			options.InternetMaxBandwidthIn = internetMaxBandwidthIn
		case "InternetMaxBandwidthOut":
			internetMaxBandwidthOut, ok := value.(int)
			if !ok {
				internetMaxBandwidthOut, _ = strconv.Atoi(value.(string))
			}
			options.InternetMaxBandwidthOut = internetMaxBandwidthOut
		case "HostName":
			hostName, _ := value.(string)
			options.HostName = hostName
		case "Password":
			password, _ := value.(string)
			options.Password = password
		case "IoOptimized":
			ioOptimized, _ := value.(string)
			options.IoOptimized = ioOptimized
		case "SystemDisk.Category":
			category, _ := value.(string)
			options.SystemDiskCategory = category
		case "SystemDisk.Size":
			systemDiskSize, _ := value.(string)
			options.SystemDiskSize = systemDiskSize
		case "SystemDisk.DiskName":
			systemDiskName, _ := value.(string)
			options.SystemDiskName = systemDiskName
		case "SystemDisk.Description":
			systemDiskDescription, _ := value.(string)
			options.SystemDiskDescription = systemDiskDescription
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
		case "int":
			if e.Field(i).Interface() != 0 {
				params[typeOfOptions.Field(i).Name] = e.Field(i).Interface()
			}
		}
	}

	response := make(map[string]interface{})
	err = aliauth.SignAndDoRequest("CreateInstance", params, response)
	resp = response
	return resp, err
}
