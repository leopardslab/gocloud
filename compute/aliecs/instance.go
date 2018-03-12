package aliecs

import (
	"github.com/cloudlibz/gocloud/aliauth"
)

func (ecs *ECS) Startnode(request interface{}) (resp interface{}, err error) {
	return resp, err
}
func (ecs *ECS) Stopnode(request interface{}) (resp interface{}, err error) {
	return resp, err
}
func (ecs *ECS) Rebootnode(request interface{}) (resp interface{}, err error) {
	return resp, err
}
func (ecs *ECS) Deletenode(request interface{}) (resp interface{}, err error) {
	return resp, err
}
func (ecs *ECS) Createnode(request interface{}) (resp interface{}, err error) {
	var options CreateInstance

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "RegionId":
			options.RegionId = value.(string)
		case "ImageId":
			options.ImageId = value.(string)
		case "InstanceType":
			options.InstanceType = value.(string)
		case "SecurityGroupId":
			options.SecurityGroupId = value.(string)
		case "ZoneId":
			options.ZoneId = value.(string)
		case "InstanceName":
			options.InstanceName = value.(string)
		case "Description":
		case "InternetChargeType":
		case "InternetMaxBandwidthIn":
		case "InternetMaxBandwidthOut":
		case "HostName":
		case "Password":
		case "IoOptimized":
		case "SystemDisk.Category":
		case "SystemDisk.Size":
		case "SystemDisk.DiskName":
		case "SystemDisk.Description":
		}
	}

	params := make(map[string]interface{})

	if options.RegionId != "" {
		params["RegionId"] = options.RegionId
	}
	if options.ImageId != "" {
		params["ImageId"] = options.ImageId
	}
	if options.InstanceType != "" {
		params["InstanceType"] = options.InstanceType
	}
	if options.SecurityGroupId != "" {
		params["SecurityGroupId"] = options.SecurityGroupId
	}
	if options.ZoneId != "" {
		params["ZoneId"] = options.ZoneId
	}
	if options.InstanceName != "" {
		params["InstanceName"] = options.InstanceName
	}

	response := make(map[string]interface{})
	err = aliauth.SignAndDoRequest("CreateInstance", params, response)
	resp = response
	return resp, err
}
