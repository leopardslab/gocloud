package ec2

import (
	"strconv"
)

// prepareRunParams base on vps or legacy params
func prepareRunParams(options RunInstances) map[string]string {
	if options.SubnetId != "" || len(options.NetworkInterfaces) > 0 {
		return makeParamsVPC("RunInstances")
	} else {
		return makeParams("RunInstances")
	}
}

func makeParams(action string) map[string]string {
	return makeParamsWithVersion(action, legacyAPIVersion)
}

func makeParamsVPC(action string) map[string]string {
	return makeParamsWithVersion(action, vpcAPIVersion)
}

//add version to params
func makeParamsWithVersion(action, version string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	params["Version"] = version
	return params
}

// prepare Block Devices add into request param
func prepareBlockDevices(params map[string]string, blockDevs []BlockDeviceMapping) {
	for i, b := range blockDevs {
		n := strconv.Itoa(i + 1)
		prefix := "BlockDeviceMapping." + n
		if b.DeviceName != "" {
			params[prefix+".DeviceName"] = b.DeviceName
		}
		if b.VirtualName != "" {
			params[prefix+".VirtualName"] = b.VirtualName
		}
		if b.SnapshotId != "" {
			params[prefix+".Ebs.SnapshotId"] = b.SnapshotId
		}
		if b.VolumeType != "" {
			params[prefix+".Ebs.VolumeType"] = b.VolumeType
		}
		if b.VolumeSize > 0 {
			params[prefix+".Ebs.VolumeSize"] = strconv.FormatInt(b.VolumeSize, 10)
		}
		if b.IOPS > 0 {
			params[prefix+".Ebs.Iops"] = strconv.FormatInt(b.IOPS, 10)
		}
		if b.DeleteOnTermination {
			params[prefix+".Ebs.DeleteOnTermination"] = "true"
		}
	}
}

// prepareNetworkInterfaces add to request param

func prepareNetworkInterfaces(params map[string]string, nics []RunNetworkInterface) {
	for i, ni := range nics {
		n := strconv.Itoa(i)
		prefix := "NetworkInterface." + n
		if ni.Id != "" {
			params[prefix+".NetworkInterfaceId"] = ni.Id
		}
		params[prefix+".DeviceIndex"] = strconv.Itoa(ni.DeviceIndex)
		if ni.SubnetId != "" {
			params[prefix+".SubnetId"] = ni.SubnetId
		}
		if ni.Description != "" {
			params[prefix+".Description"] = ni.Description
		}
		for j, gid := range ni.SecurityGroupIds {
			k := strconv.Itoa(j + 1)
			params[prefix+".SecurityGroupId."+k] = gid
		}
		if ni.DeleteOnTermination {
			params[prefix+".DeleteOnTermination"] = "true"
		}
		if ni.SecondaryPrivateIPCount > 0 {
			val := strconv.Itoa(ni.SecondaryPrivateIPCount)
			params[prefix+".SecondaryPrivateIpAddressCount"] = val
		}
		for j, ip := range ni.PrivateIPs {
			k := strconv.Itoa(j)
			subprefix := prefix + ".PrivateIpAddresses." + k
			params[subprefix+".PrivateIpAddress"] = ip.Address
			params[subprefix+".Primary"] = strconv.FormatBool(ip.IsPrimary)
		}
	}
}

// use to create params to start stop and reboot, terminate instance
func addParamsList(params map[string]string, label string, ids []string) {
	for i, id := range ids {
		params[label+"."+strconv.Itoa(i+1)] = id
	}
}
