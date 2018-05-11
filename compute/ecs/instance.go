package ecs

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"strconv"
	"reflect"
)

// Startnode start ECS instances accept map[string]interface{}
func (ecs *ECS) Startnode(request interface{}) (resp interface{}, err error) {
	var options StartInstance

	param = make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "InstanceId":
			instanceID, _ := value.(string)
			options.InstanceID = instanceID
		case "InitLocalDisk":
			switch value.(type) {
			case bool:
				options.InitLocalDisk = value.(bool)
			case string:
				options.InitLocalDisk = value.(string) == "true" || value.(string) == "True"
			}
		}
	}

	params := make(map[string]interface{})

	// Put all of options into params
	params = aliauth.PutStructIntoMap(&options)
	response := make(map[string]interface{})
	err = aliauth.ECSSignAndDoRequest("StartInstance", params, response)
	resp = response
	return resp, err

}

// Stopnode stop ECS instances accept map[string]interface{}
func (ecs *ECS) Stopnode(request interface{}) (resp interface{}, err error) {
	var options StopInstance

	param = make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "InstanceId":
			instanceID, _ := value.(string)
			options.InstanceID = instanceID
		case "ForceStop":
			switch value.(type) {
			case bool:
				options.ForceStop = value.(bool)
			case string:
				options.ForceStop = value.(string) == "true" || value.(string) == "True"
			}
		case "ConfirmStop":
			switch value.(type) {
			case bool:
				options.ConfirmStop = value.(bool)
			case string:
				options.ConfirmStop = value.(string) == "true" || value.(string) == "True"
			}
		case "StoppedMode":
			stoppedMode, _ := value.(string)
			options.StoppedMode = stoppedMode
		}
	}

	params := make(map[string]interface{})

	// Put all of options into params
	params = aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.ECSSignAndDoRequest("StopInstance", params, response)
	resp = response
	return resp, err
}

// Rebootnode reboot ECS instances accept map[string]interface{}
func (ecs *ECS) Rebootnode(request interface{}) (resp interface{}, err error) {
	var options RebootInstance

	param = make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "InstanceId":
			instanceID, _ := value.(string)
			options.InstanceID = instanceID
		case "ForceStop":
			switch value.(type) {
			case bool:
				options.ForceStop = value.(bool)
			case string:
				options.ForceStop = value.(string) == "true" || value.(string) == "True"
			}
		}
	}

	params := make(map[string]interface{})

	// Put all of options into params
	params = aliauth.PutStructIntoMap(&options)

	response = make(map[string]interface{})
	err = aliauth.SignAndDoRequest("RebootInstance", params, response)
	resp = response
	return resp, err
}

// Deletenode delete ECS instances accept map[string]interface{}
func (ecs *ECS) Deletenode(request interface{}) (resp interface{}, err error) {
	var options DeleteInstance

	param = make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "InstanceId":
			instanceID, _ := value.(string)
			options.InstanceID = instanceID
		}
	}

	params := make(map[string]interface{})

	// Put all of options into params
	params = aliauth.PutStructIntoMap(&options)
	response := make(map[string]interface{})
	err = aliauth.ECSSignAndDoRequest("DeleteInstance", params, response)
	resp = response
	return resp, err
}

// Createnode create ECS instances accept map[string]interface{}
func (ecs *ECS) Createnode(request interface{}) (resp interface{}, err error) {
	var options CreateInstance

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	params := make(map[string]interface{})

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
			switch value.(type) {
			case int:
				options.InternetMaxBandwidthIn = value.(int)
			case string:
				options.InternetMaxBandwidthIn, _ = strconv.Atoi(value.(string))
			}
		case "InternetMaxBandwidthOut":
			switch value.(type) {
			case int:
				options.InternetMaxBandwidthOut = value.(int)
			case string:
				options.InternetMaxBandwidthOut, _ = strconv.Atoi(value.(string))
			}
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
		default:
			switch value.(type) {
			case string:
				params[key] = value.(string)
			case int:
				params[key] = strconv.Itoa(value.(int))
			case bool:
				params[key] = strconv.FormatBool(value.(bool))
			}
		}
	}

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
	err = aliauth.ECSSignAndDoRequest("CreateInstance", params, response)
	resp = response
	return resp, err
}

func (ecs *ECS) ListNodeType(request interface{}) (resp interface{}, err error)  {
	params := make(map[string]interface{})
	response := make(map[string]interface{})
	err = aliauth.ECSSignAndDoRequest("DescribeInstanceTypes", params, response)
	resp = response
	return resp, err
}
