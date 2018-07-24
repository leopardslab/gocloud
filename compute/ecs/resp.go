package ecs

import (
	"encoding/json"
	"time"
)

type CreateNodeResp struct {
	InstanceId string
}

type ListNodeResp struct {
	Instances struct {
		Instance []NodeInfo
	}
}

type NodeInfo struct {
	InstanceId         string
	InstanceName       string
	Description        string
	ImageId            string
	RegionId           string
	ZoneId             string
	Cpu                int
	Memory             int
	InstanceType       string
	InstanceTypeFamily string
	HostName           string
	SerialNumber       string
	Status             string
	SecurityGroupIds struct {
		SecurityGroupId []string
	}
	PublicIpAddress         IpAddressSetType
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	InternetChargeType      string
	CreationTime            time.Time
	InstanceNetworkType     string
	VpcAttributes           VpcAttributesType
	EipAddress              EipAddressAssociateType
	InnerIpAddress          IpAddressSetType
	OperationLocks          OperationLocksType
	InstanceChargeType      string
	SpotStrategy            string
	DeviceAvailable         StringBool
	StoppedMode             string
	DeploymentSetId         string
	NetworkInterfaces       NetworkInterfaceType
	IoOptimized             StringBool
	ExpiredTime             time.Time
	KeyPairName             string
}
type IpAddressSetType struct {
	IpAddress []string
}

type VpcAttributesType struct {
	VpcId            string
	VSwitchId        string
	PrivateIpAddress IpAddressSetType
	NatIpAddress     string
}
type LockReasonType struct {
	LockReason string
}
type OperationLocksType struct {
	LockReason []LockReasonType //enum for financial, security
}

type EipAddressAssociateType struct {
	AllocationId         string
	IpAddress            string
	Bandwidth            int
	InternetChargeType   string
	IsSupportUnassociate StringBool
}

type NetworkInterfaceType struct {
	NetworkInterfaceId string
	PrimaryIpAddress   string
	MacAddress         string
}

type StringBool struct {
	Value bool
}

func (io *StringBool) UnmarshalJSON(value []byte) error {
	if value[0] == '"' {
		var str string
		err := json.Unmarshal(value, &str)
		if err == nil {
			io.Value = str == "true"
		}
		return err
	} else {
		var boolV bool
		err := json.Unmarshal(value, &boolV)
		if err == nil {
			io.Value = boolV
		}
		return err
	}
}

func ParseCreateNodeResp(body interface{}) (createNodeResp CreateNodeResp, err error) {
	err = json.Unmarshal([]byte(body.(string)), &createNodeResp)
	return
}

func ParseListNodeResp(body interface{}) (nodes []NodeInfo, err error) {
	listNodeResp := &ListNodeResp{}
	err = json.Unmarshal([]byte(body.(string)), &listNodeResp)
	nodes = listNodeResp.Instances.Instance
	return
}
