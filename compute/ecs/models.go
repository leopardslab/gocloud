package ecs

import (
	"errors"
	"github.com/cloudlibz/gocloud/aliauth"
)

const errCommon = "miss required parameter: "

// CreateNode to store all attribute to create Ali-cloud ECS instance
type CreateNode struct {
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

// StartNode to store all attribute to start Ali-cloud ECS instance
type StartNode struct {
	InstanceID    string
	InitLocalDisk bool
}

// StopNode to store all attribute to Stop Ali-cloud ECS instance
type StopNode struct {
	InstanceID  string
	ForceStop   bool
	ConfirmStop bool
	StoppedMode string
}

// RebootNode to store all attribute to Reboot Ali-cloud ECS instance
type RebootNode struct {
	InstanceID string
	ForceStop  bool
}

// DeleteNode to store all attribute to Delete Ali-cloud ECS instance
type DeleteNode struct {
	InstanceID string
}

// CreateNode builder pattern code
type CreateNodeBuilder struct {
	createNode *CreateNode
}

func NewCreateNodeBuilder() *CreateNodeBuilder {
	createNode := &CreateNode{}
	b := &CreateNodeBuilder{createNode: createNode}
	return b
}

func (b *CreateNodeBuilder) RegionID(regionID string) *CreateNodeBuilder {
	b.createNode.RegionID = regionID
	return b
}

func (b *CreateNodeBuilder) ZoneID(zoneID string) *CreateNodeBuilder {
	b.createNode.ZoneID = zoneID
	return b
}

func (b *CreateNodeBuilder) ImageID(imageID string) *CreateNodeBuilder {
	b.createNode.ImageID = imageID
	return b
}

func (b *CreateNodeBuilder) InstanceType(instanceType string) *CreateNodeBuilder {
	b.createNode.InstanceType = instanceType
	return b
}

func (b *CreateNodeBuilder) SecurityGroupID(securityGroupID string) *CreateNodeBuilder {
	b.createNode.SecurityGroupID = securityGroupID
	return b
}

func (b *CreateNodeBuilder) InstanceName(instanceName string) *CreateNodeBuilder {
	b.createNode.InstanceName = instanceName
	return b
}

func (b *CreateNodeBuilder) Description(description string) *CreateNodeBuilder {
	b.createNode.Description = description
	return b
}

func (b *CreateNodeBuilder) InternetChargeType(internetChargeType string) *CreateNodeBuilder {
	b.createNode.InternetChargeType = internetChargeType
	return b
}

func (b *CreateNodeBuilder) InternetMaxBandwidthIn(internetMaxBandwidthIn int) *CreateNodeBuilder {
	b.createNode.InternetMaxBandwidthIn = internetMaxBandwidthIn
	return b
}

func (b *CreateNodeBuilder) InternetMaxBandwidthOut(internetMaxBandwidthOut int) *CreateNodeBuilder {
	b.createNode.InternetMaxBandwidthOut = internetMaxBandwidthOut
	return b
}

func (b *CreateNodeBuilder) HostName(hostName string) *CreateNodeBuilder {
	b.createNode.HostName = hostName
	return b
}

func (b *CreateNodeBuilder) Password(password string) *CreateNodeBuilder {
	b.createNode.Password = password
	return b
}

func (b *CreateNodeBuilder) IoOptimized(ioOptimized string) *CreateNodeBuilder {
	b.createNode.IoOptimized = ioOptimized
	return b
}

func (b *CreateNodeBuilder) SystemDiskCategory(systemDiskCategory string) *CreateNodeBuilder {
	b.createNode.SystemDiskCategory = systemDiskCategory
	return b
}

func (b *CreateNodeBuilder) SystemDiskSize(systemDiskSize string) *CreateNodeBuilder {
	b.createNode.SystemDiskSize = systemDiskSize
	return b
}

func (b *CreateNodeBuilder) SystemDiskName(systemDiskName string) *CreateNodeBuilder {
	b.createNode.SystemDiskName = systemDiskName
	return b
}

func (b *CreateNodeBuilder) SystemDiskDescription(systemDiskDescription string) *CreateNodeBuilder {
	b.createNode.SystemDiskDescription = systemDiskDescription
	return b
}

func (b *CreateNodeBuilder) Build() (map[string]interface{}, error) {
	if b.createNode.RegionID == "" {
		return nil, errors.New(errCommon + "RegionID")
	}
	if b.createNode.ImageID == "" {
		return nil, errors.New(errCommon + "ImageID")
	}
	if b.createNode.InstanceType == "" {
		return nil, errors.New(errCommon + "InstanceType")
	}
	if b.createNode.SecurityGroupID == "" {
		return nil, errors.New(errCommon + "SecurityGroupID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.createNode)
	return params, nil
}

// StartNode builder pattern code
type StartNodeBuilder struct {
	startNode *StartNode
}

func NewStartNodeBuilder() *StartNodeBuilder {
	startNode := &StartNode{}
	b := &StartNodeBuilder{startNode: startNode}
	return b
}

func (b *StartNodeBuilder) InstanceID(instanceID string) *StartNodeBuilder {
	b.startNode.InstanceID = instanceID
	return b
}

func (b *StartNodeBuilder) InitLocalDisk(initLocalDisk bool) *StartNodeBuilder {
	b.startNode.InitLocalDisk = initLocalDisk
	return b
}

func (b *StartNodeBuilder) Build() (map[string]interface{}, error) {
	if b.startNode.InstanceID == "" {
		return nil, errors.New(errCommon + "InstanceID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.startNode)
	return params, nil
}

// StopNode builder pattern code
type StopNodeBuilder struct {
	stopNode *StopNode
}

func NewStopNodeBuilder() *StopNodeBuilder {
	stopNode := &StopNode{}
	b := &StopNodeBuilder{stopNode: stopNode}
	return b
}

func (b *StopNodeBuilder) InstanceID(instanceID string) *StopNodeBuilder {
	b.stopNode.InstanceID = instanceID
	return b
}

func (b *StopNodeBuilder) ForceStop(forceStop bool) *StopNodeBuilder {
	b.stopNode.ForceStop = forceStop
	return b
}

func (b *StopNodeBuilder) ConfirmStop(confirmStop bool) *StopNodeBuilder {
	b.stopNode.ConfirmStop = confirmStop
	return b
}

func (b *StopNodeBuilder) StoppedMode(stoppedMode string) *StopNodeBuilder {
	b.stopNode.StoppedMode = stoppedMode
	return b
}

func (b *StopNodeBuilder) Build() (map[string]interface{}, error) {
	if b.stopNode.InstanceID == "" {
		return nil, errors.New(errCommon + "InstanceID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.stopNode)
	return params, nil
}

// RebootNode builder pattern code
type RebootNodeBuilder struct {
	rebootNode *RebootNode
}

func NewRebootNodeBuilder() *RebootNodeBuilder {
	rebootNode := &RebootNode{}
	b := &RebootNodeBuilder{rebootNode: rebootNode}
	return b
}

func (b *RebootNodeBuilder) InstanceID(instanceID string) *RebootNodeBuilder {
	b.rebootNode.InstanceID = instanceID
	return b
}

func (b *RebootNodeBuilder) ForceStop(forceStop bool) *RebootNodeBuilder {
	b.rebootNode.ForceStop = forceStop
	return b
}

func (b *RebootNodeBuilder) Build() (map[string]interface{}, error) {
	if b.rebootNode.InstanceID == "" {
		return nil, errors.New(errCommon + "InstanceID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.rebootNode)
	return params, nil
}

// DeleteNode builder pattern code
type DeleteNodeBuilder struct {
	deleteNode *DeleteNode
}

func NewDeleteNodeBuilder() *DeleteNodeBuilder {
	deleteNode := &DeleteNode{}
	b := &DeleteNodeBuilder{deleteNode: deleteNode}
	return b
}

func (b *DeleteNodeBuilder) InstanceID(instanceID string) *DeleteNodeBuilder {
	b.deleteNode.InstanceID = instanceID
	return b
}

func (b *DeleteNodeBuilder) Build() (map[string]interface{}, error) {
	if b.deleteNode.InstanceID == "" {
		return nil, errors.New(errCommon + "InstanceID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.deleteNode)
	return params, nil
}
