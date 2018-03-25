package ec2

import (
	"encoding/base64"
	"strconv"
)

// start ec2 instance accept array of instance-id

func (ec2 *EC2) Startnode(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]string)
	ids := []string{}
	ids = append(ids, param["instance-id"])
	Region := param["Region"]

	params := makeParams("StartInstances")

	addParamsList(params, "InstanceID", ids)

	response := make(map[string]interface{})

	err = ec2.PrepareSignatureV2query(params, Region, response)
	resp = response
	return resp, err
}

// stop ec2 instance accept array of instance-id

func (ec2 *EC2) Stopnode(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]string)
	ids := []string{}
	ids = append(ids, param["instance-id"])
	Region := param["Region"]

	params := makeParams("StopInstances")
	addParamsList(params, "InstanceID", ids)
	resp = &StopInstanceResp{}

	response := make(map[string]interface{})

	err = ec2.PrepareSignatureV2query(params, Region, response)

	if err != nil {
		return nil, err
	}

	resp = response
	return resp, nil
}

// reboot ec2 instance accept array of instance-id

func (ec2 *EC2) Rebootnode(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]string)
	ids := []string{}
	ids = append(ids, param["instance-id"])
	Region := param["Region"]

	params := makeParams("RebootInstances")
	addParamsList(params, "InstanceID", ids)

	response := make(map[string]interface{})

	err = ec2.PrepareSignatureV2query(params, Region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, err
}

// delete ec2 instance accept array of instance-id

func (ec2 *EC2) Deletenode(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]string)
	instIds := []string{}
	instIds = append(instIds, param["instance-id"])
	Region := param["Region"]

	params := makeParams("TerminateInstances")
	addParamsList(params, "InstanceID", instIds)
	response := make(map[string]interface{})

	err = ec2.PrepareSignatureV2query(params, Region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, err
}

//create Ec2 instances accept map[string]interface{} with attribute Define in EC2 documentation

func (ec2 *EC2) Createnode(request interface{}) (resp interface{}, err error) {

	var options RunInstances
	var Region string

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {

		case "Region":
			regionV, _ := value.(string)
			Region = regionV

		case "ImageId":
			ImageId, _ := value.(string)
			options.ImageId = ImageId

		case "MinCount":
			MinCount, _ := value.(int)
			options.MinCount = MinCount

		case "MaxCount":
			MaxCount, _ := value.(int)
			options.MaxCount = MaxCount

		case "KeyName":
			KeyName, _ := value.(string)
			options.KeyName = KeyName

		case "KernelId":
			KernelId, _ := value.(string)
			options.KernelId = KernelId

		case "InstanceType":
			InstanceType, _ := value.(string)
			options.InstanceType = InstanceType

		case "RamdiskId":
			RamdiskId, _ := value.(string)
			options.RamdiskId = RamdiskId

		case "AvailZone":
			AvailZone, _ := value.(string)
			options.AvailZone = AvailZone

		case "PlacementGroupName":
			PlacementGroupName, _ := value.(string)
			options.PlacementGroupName = PlacementGroupName

		case "Monitoring":
			Monitoring, _ := value.(bool)
			options.Monitoring = Monitoring

		case "SubnetId":
			SubnetId, _ := value.(string)
			options.SubnetId = SubnetId

		case "DisableAPITermination":
			DisableAPITermination, _ := value.(bool)
			options.DisableAPITermination = DisableAPITermination

		case "ShutdownBehavior":
			ShutdownBehavior, _ := value.(string)
			options.ShutdownBehavior = ShutdownBehavior

		case "PrivateIPAddress":
			PrivateIPAddress, _ := value.(string)
			options.PrivateIPAddress = PrivateIPAddress

		case "SecurityGroup":
			SecurityGroupparam, _ := value.([]map[string]string)
			for i := 0; i < len(SecurityGroupparam); i++ {
				var securityGroup SecurityGroup
				securityGroup.Id = SecurityGroupparam[i]["Id"]
				securityGroup.Name = SecurityGroupparam[i]["Name"]
				options.SecurityGroups = append(options.SecurityGroups, securityGroup)
			}

		case "BlockDevice":
			BlockDeviceparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(BlockDeviceparam); i++ {
				var BlockDeviceMappingParam BlockDeviceMapping
				for BlockDeviceparamkey, BlockDeviceparamvalue := range BlockDeviceparam[i] {
					switch BlockDeviceparamkey {
					case "DeviceName":
						BlockDeviceMappingParam.DeviceName = BlockDeviceparamvalue.(string)

					case "VirtualName":
						BlockDeviceMappingParam.VirtualName = BlockDeviceparamvalue.(string)

					case "SnapshotId":
						BlockDeviceMappingParam.SnapshotId = BlockDeviceparamvalue.(string)

					case "VolumeType":
						BlockDeviceMappingParam.VolumeType = BlockDeviceparamvalue.(string)

					case "VolumeSize":
						BlockDeviceMappingParam.VolumeSize = BlockDeviceparamvalue.(int64)

					case "DeleteOnTermination":
						BlockDeviceMappingParam.DeleteOnTermination = BlockDeviceparamvalue.(bool)

					case "IOPS":
						BlockDeviceMappingParam.IOPS = BlockDeviceparamvalue.(int64)

					}
				}
				options.BlockDeviceMappings = append(options.BlockDeviceMappings, BlockDeviceMappingParam)

			}

		case "RunNetworkInterface":
			RunNetworkInterfaceparam, _ := value.([]map[string]interface{})
			var runNetworkInterface RunNetworkInterface
			for i := 0; i < len(RunNetworkInterfaceparam); i++ {
				for RunNetworkInterfaceparamkey, RunNetworkInterfaceparamvalue := range RunNetworkInterfaceparam[i] {
					switch RunNetworkInterfaceparamkey {
					case "Id":
						runNetworkInterface.Id = RunNetworkInterfaceparamvalue.(string)
					case "DeviceIndex":
						runNetworkInterface.DeviceIndex = RunNetworkInterfaceparamvalue.(int)
					case "SubnetId":
						runNetworkInterface.Id = RunNetworkInterfaceparamvalue.(string)
					case "Description":
						runNetworkInterface.Description = RunNetworkInterfaceparamvalue.(string)
					case "DeleteOnTermination":
						runNetworkInterface.DeleteOnTermination = RunNetworkInterfaceparamvalue.(bool)
					case "SecondaryPrivateIPCount":
						runNetworkInterface.SecondaryPrivateIPCount = RunNetworkInterfaceparamvalue.(int)
					case "SecurityGroupIds":
						securityGroupIds, _ := RunNetworkInterfaceparamvalue.([]string)
						runNetworkInterface.SecurityGroupIds = securityGroupIds

					case "PrivateIPs":
						privateIPsParam, _ := RunNetworkInterfaceparamvalue.([]map[string]interface{})
						for i := 0; i < len(privateIPsParam); i++ {
							var privateIP PrivateIP
							for privateIPsParamkey, privateIPsParamvalue := range privateIPsParam[i] {
								switch privateIPsParamkey {
								case "Address":
									privateIP.Address = privateIPsParamvalue.(string)
								case "DNSName":
									privateIP.DNSName = privateIPsParamvalue.(string)
								case "IsPrimary":
									privateIP.IsPrimary = privateIPsParamvalue.(bool)
								}

							}
							runNetworkInterface.PrivateIPs = append(runNetworkInterface.PrivateIPs, privateIP)
						}

					}

				}
				options.NetworkInterfaces = append(options.NetworkInterfaces, runNetworkInterface)
			}
		case "UserData":
			options.UserData = value.([]byte)
		}
	}

	params := makeParams("RunInstances")

	params["ImageId"] = options.ImageId

	params["InstanceType"] = options.InstanceType
	var min, max int
	if options.MinCount == 0 && options.MaxCount == 0 {
		min = 1
		max = 1
	} else if options.MaxCount == 0 {
		min = options.MinCount
		max = min
	} else {
		min = options.MinCount
		max = options.MaxCount
	}
	params["MinCount"] = strconv.Itoa(min)
	params["MaxCount"] = strconv.Itoa(max)
	i, j := 1, 1
	for _, g := range options.SecurityGroups {
		if g.Id != "" {
			params["SecurityGroupId."+strconv.Itoa(i)] = g.Id
			i++
		} else {
			params["SecurityGroup."+strconv.Itoa(j)] = g.Name
			j++
		}
	}
	prepareBlockDevices(params, options.BlockDeviceMappings)
	prepareNetworkInterfaces(params, options.NetworkInterfaces)
	token, err := clientToken()
	if err != nil {
		return nil, err
	}
	params["ClientToken"] = token

	if options.KeyName != "" {
		params["KeyName"] = options.KeyName
	}
	if options.KernelId != "" {
		params["KernelId"] = options.KernelId
	}
	if options.RamdiskId != "" {
		params["RamdiskId"] = options.RamdiskId
	}
	if options.UserData != nil {
		userData := make([]byte, base64.StdEncoding.EncodedLen(len(options.UserData)))
		base64.StdEncoding.Encode(userData, options.UserData)
		params["UserData"] = string(userData)
	}
	if options.AvailZone != "" {
		params["Placement.AvailabilityZone"] = options.AvailZone
	}
	if options.PlacementGroupName != "" {
		params["Placement.GroupName"] = options.PlacementGroupName
	}
	if options.Monitoring {
		params["Monitoring.Enabled"] = "true"
	}
	if options.SubnetId != "" {
		params["SubnetId"] = options.SubnetId
	}
	if options.DisableAPITermination {
		params["DisableApiTermination"] = "true"
	}
	if options.ShutdownBehavior != "" {
		params["InstanceInitiatedShutdownBehavior"] = options.ShutdownBehavior
	}
	if options.PrivateIPAddress != "" {
		params["PrivateIpAddress"] = options.PrivateIPAddress
	}

	response := make(map[string]interface{})
	err = ec2.PrepareSignatureV2query(params, Region, response)
	resp = response

	return resp, err
}
