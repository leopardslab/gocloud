package ecs

import (
	"errors"
	"github.com/cloudlibz/gocloud/aliauth"
)

const errCommon = "miss required parameter: "

// CreateInstance to store all attributes to create Ali-cloud ECS instance.
type CreateInstance struct {
	RegionID                string
	ZoneID                  string
	ImageID                 string
	InstanceType            string
	SecurityGroupID         string
	InstanceName            string
	Description             string
	InternetChargeType      string
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	HostName                string
	Password                string
	IoOptimized             string
	SystemDiskCategory      string
	SystemDiskSize          string
	SystemDiskName          string
	SystemDiskDescription   string
}

// StartInstance to store all attributes to start Ali-cloud ECS instance.
type StartInstance struct {
	InstanceID    string
	InitLocalDisk bool
}

// StopInstance to store all attributes to stop Ali-cloud ECS instance.
type StopInstance struct {
	InstanceID  string
	ForceStop   bool
	ConfirmStop bool
	StoppedMode string
}

// RebootInstance to store all attributes to reboot Ali-cloud ECS instance.
type RebootInstance struct {
	InstanceID string
	ForceStop  bool
}

// DeleteInstance to store all attributes to Delete Ali-cloud ECS instance.
type DeleteInstance struct {
	InstanceID string
}

type CreateNodeBuilder struct {
	params *CreateInstance
}

func NewCreateNodeBuilder() *CreateNodeBuilder {
	createInstance := &CreateInstance{}
	createNodeBuilder := &CreateNodeBuilder{params: createInstance}
	return createNodeBuilder
}
func (b *CreateNodeBuilder) RegionID(regionID string) *CreateNodeBuilder {
	b.params.RegionID = regionID
	return b
}
func (b *CreateNodeBuilder) ImageID(imageID string) *CreateNodeBuilder {
	b.params.ImageID = imageID
	return b
}
func (b *CreateNodeBuilder) InstanceType(instanceType string) *CreateNodeBuilder {
	b.params.InstanceType = instanceType
	return b
}
func (b *CreateNodeBuilder) SecurityGroupID(securityGroupID string) *CreateNodeBuilder {
	b.params.SecurityGroupID = securityGroupID
	return b
}
func (b *CreateNodeBuilder) ZoneID(zoneID string) *CreateNodeBuilder {
	b.params.ZoneID = zoneID
	return b
}
func (b *CreateNodeBuilder) InstanceName(instanceName string) *CreateNodeBuilder {
	b.params.InstanceName = instanceName
	return b
}
func (b *CreateNodeBuilder) Description(description string) *CreateNodeBuilder {
	b.params.Description = description
	return b
}
func (b *CreateNodeBuilder) InternetChargeType(internetChargeType string) *CreateNodeBuilder {
	b.params.InternetChargeType = internetChargeType
	return b
}
func (b *CreateNodeBuilder) InternetMaxBandwidthIn(internetMaxBandwidthIn int) *CreateNodeBuilder {
	b.params.InternetMaxBandwidthIn = internetMaxBandwidthIn
	return b
}
func (b *CreateNodeBuilder) InternetMaxBandwidthOut(internetMaxBandwidthOut int) *CreateNodeBuilder {
	b.params.InternetMaxBandwidthOut = internetMaxBandwidthOut
	return b
}
func (b *CreateNodeBuilder) HostName(hostName string) *CreateNodeBuilder {
	b.params.HostName = hostName
	return b
}
func (b *CreateNodeBuilder) Password(password string) *CreateNodeBuilder {
	b.params.Password = password
	return b
}
func (b *CreateNodeBuilder) IoOptimized(ioOptimized string) *CreateNodeBuilder {
	b.params.IoOptimized = ioOptimized
	return b
}
func (b *CreateNodeBuilder) SystemDiskCategory(systemDiskCategory string) *CreateNodeBuilder {
	b.params.SystemDiskCategory = systemDiskCategory
	return b
}
func (b *CreateNodeBuilder) SystemDiskSize(systemDiskSize string) *CreateNodeBuilder {
	b.params.SystemDiskSize = systemDiskSize
	return b
}
func (b *CreateNodeBuilder) SystemDiskName(systemDiskName string) *CreateNodeBuilder {
	b.params.SystemDiskName = systemDiskName
	return b
}
func (b *CreateNodeBuilder) SystemDiskDescription(systemDiskDescription string) *CreateNodeBuilder {
	b.params.SystemDiskDescription = systemDiskDescription
	return b
}

func (b *CreateNodeBuilder) Build() (map[string]interface{}, error) {
	if b.params.RegionID == "" {
		return nil, errors.New(errCommon + "RegionID")
	}
	if b.params.ImageID == "" {
		return nil, errors.New(errCommon + "ImageID")
	}
	if b.params.InstanceType == "" {
		return nil, errors.New(errCommon + "InstanceType")
	}
	if b.params.SecurityGroupID == "" {
		return nil, errors.New(errCommon + "SecurityGroupID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.params)
	return params, nil
}
