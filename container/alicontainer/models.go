package alicontainer

import (
	"errors"
	"github.com/cloudlibz/gocloud/aliauth"
)

//Alicontainer struct represents Alicontainer attribute and methods associates with it.
type Alicontainer struct {
}

// CreateCluster struct represents create cluster attributes.
type CreateCluster struct {
	RegionID         string `json:"region_id"`
	Name             string `json:"name"`
	Size             int64  `json:"size"`
	InstanceType     string `json:"instance_type"`
	NetworkMode      string `json:"network_mode"`
	SubnetCIDR       string `json:"subnet_cidr,omitempty"`
	VPCID            string `json:"vpc_id,omitempty"`
	VSwitchID        string `json:"vswitch_id,omitempty"`
	Password         string `json:"password"`
	DataDiskCategory string `json:"data_disk_category"`
	DataDiskSize     int64  `json:"data_disk_size"`
	ECSImageID       string `json:"ecs_image_id,omitempty"`
	IOOptimized      string `json:"io_optimized"`
	NeedSLB          bool   `json:"need_slb"`
	ReleaseEipFlag   bool   `json:"release_eip_flag"`
}

// DeleteCluster struct represents create cluster attributes.
type DeleteCluster struct {
	RegionID  string `json:"region_id"`
	ClusterID string `json:"cluster_id"`
}

// RunTask struct represents create project of cluster attributes.
type RunTask struct {
	ClusterID   string            `json:"cluster_id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Template    string            `json:"template"`
	Version     string            `json:"version"`
	Environment map[string]string `json:"environment"`
	LatestImage bool              `json:"latest_image"`
}

// StartTask struct represents create project of cluster attributes.
type StartTask struct {
	ClusterID string `json:"cluster_id"`
	Name      string `json:"name"`
}

// StopTask struct represents create project of cluster attributes.
type StopTask struct {
	ClusterID string `json:"cluster_id"`
	Name      string `json:"name"`
	Timeout   int    `json:"timeout"`
}

type clusterCerts struct {
	CA   string `json:"ca,omitempty"`
	Key  string `json:"key,omitempty"`
	Cert string `json:"cert,omitempty"`
}
type cluster struct {
	ClusterID string `json:"cluster_id"`
	Name      string `json:"name"`
	MasterURL string `json:"master_url"`
}

// CreateCluster builder pattern code
type CreateClusterBuilder struct {
	createCluster *CreateCluster
}

func NewCreateClusterBuilder() *CreateClusterBuilder {
	createCluster := &CreateCluster{}
	b := &CreateClusterBuilder{createCluster: createCluster}
	return b
}

func (b *CreateClusterBuilder) RegionID(regionID string) *CreateClusterBuilder {
	b.createCluster.RegionID = regionID
	return b
}

func (b *CreateClusterBuilder) Name(name string) *CreateClusterBuilder {
	b.createCluster.Name = name
	return b
}

func (b *CreateClusterBuilder) Size(size int64) *CreateClusterBuilder {
	b.createCluster.Size = size
	return b
}

func (b *CreateClusterBuilder) InstanceType(instanceType string) *CreateClusterBuilder {
	b.createCluster.InstanceType = instanceType
	return b
}

func (b *CreateClusterBuilder) NetworkMode(networkMode string) *CreateClusterBuilder {
	b.createCluster.NetworkMode = networkMode
	return b
}

func (b *CreateClusterBuilder) SubnetCIDR(subnetCIDR string) *CreateClusterBuilder {
	b.createCluster.SubnetCIDR = subnetCIDR
	return b
}

func (b *CreateClusterBuilder) VPCID(vPCID string) *CreateClusterBuilder {
	b.createCluster.VPCID = vPCID
	return b
}

func (b *CreateClusterBuilder) VSwitchID(vSwitchID string) *CreateClusterBuilder {
	b.createCluster.VSwitchID = vSwitchID
	return b
}

func (b *CreateClusterBuilder) Password(password string) *CreateClusterBuilder {
	b.createCluster.Password = password
	return b
}

func (b *CreateClusterBuilder) DataDiskCategory(dataDiskCategory string) *CreateClusterBuilder {
	b.createCluster.DataDiskCategory = dataDiskCategory
	return b
}

func (b *CreateClusterBuilder) DataDiskSize(dataDiskSize int64) *CreateClusterBuilder {
	b.createCluster.DataDiskSize = dataDiskSize
	return b
}

func (b *CreateClusterBuilder) ECSImageID(eCSImageID string) *CreateClusterBuilder {
	b.createCluster.ECSImageID = eCSImageID
	return b
}

func (b *CreateClusterBuilder) IOOptimized(iOOptimized string) *CreateClusterBuilder {
	b.createCluster.IOOptimized = iOOptimized
	return b
}

var isNeedSLBSet bool

func (b *CreateClusterBuilder) NeedSLB(needSLB bool) *CreateClusterBuilder {
	isNeedSLBSet = true
	b.createCluster.NeedSLB = needSLB
	return b
}

func (b *CreateClusterBuilder) ReleaseEipFlag(releaseEipFlag bool) *CreateClusterBuilder {
	b.createCluster.ReleaseEipFlag = releaseEipFlag
	return b
}

func (b *CreateClusterBuilder) Build() (map[string]interface{}, error) {
	if b.createCluster.RegionID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RegionID")
	}
	if b.createCluster.Name == "" {
		return nil, errors.New(aliauth.StrMissRequired + "Name")
	}
	if b.createCluster.Size == 0 {
		return nil, errors.New(aliauth.StrMissRequired + "Size")
	}
	if b.createCluster.InstanceType == "" {
		return nil, errors.New(aliauth.StrMissRequired + "InstanceType")
	}
	if b.createCluster.NetworkMode == "" {
		return nil, errors.New(aliauth.StrMissRequired + "NetworkMode")
	}
	if b.createCluster.SubnetCIDR == "" {
		return nil, errors.New(aliauth.StrMissRequired + "SubnetCIDR")
	}
	if b.createCluster.VPCID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "VPCID")
	}
	if b.createCluster.VSwitchID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "VSwitchID")
	}
	if b.createCluster.Password == "" {
		return nil, errors.New(aliauth.StrMissRequired + "Password")
	}
	if b.createCluster.DataDiskCategory == "" {
		return nil, errors.New(aliauth.StrMissRequired + "DataDiskCategory")
	}
	if b.createCluster.DataDiskSize == 0 {
		return nil, errors.New(aliauth.StrMissRequired + "DataDiskSize")
	}

	params := make(map[string]interface{})

	params["region_id"] = b.createCluster.RegionID
	params["name"] = b.createCluster.Name
	params["size"] = b.createCluster.Size
	params["instance_type"] = b.createCluster.InstanceType
	params["network_mode"] = b.createCluster.NetworkMode
	params["subnet_cidr"] = b.createCluster.SubnetCIDR
	params["vpc_id"] = b.createCluster.VPCID
	params["vswitch_id"] = b.createCluster.VSwitchID
	params["password"] = b.createCluster.Password
	params["data_disk_category"] = b.createCluster.DataDiskCategory
	params["data_disk_size"] = b.createCluster.DataDiskSize

	if b.createCluster.ECSImageID != "" {
		params["ecs_image_id"] = b.createCluster.ECSImageID
	}
	if b.createCluster.IOOptimized != "" {
		params["io_optimized"] = b.createCluster.IOOptimized
	}
	if isNeedSLBSet {
		params["need_slb"] = b.createCluster.NeedSLB
	} else {
		params["need_slb"] = true
	}
	params["release_eip_flag"] = b.createCluster.ReleaseEipFlag

	return params, nil
}

// DeleteCluster builder pattern code
type DeleteClusterBuilder struct {
	deleteCluster *DeleteCluster
}

func NewDeleteClusterBuilder() *DeleteClusterBuilder {
	deleteCluster := &DeleteCluster{}
	b := &DeleteClusterBuilder{deleteCluster: deleteCluster}
	return b
}

func (b *DeleteClusterBuilder) RegionID(regionID string) *DeleteClusterBuilder {
	b.deleteCluster.RegionID = regionID
	return b
}

func (b *DeleteClusterBuilder) ClusterID(clusterID string) *DeleteClusterBuilder {
	b.deleteCluster.ClusterID = clusterID
	return b
}

func (b *DeleteClusterBuilder) Build() (map[string]interface{}, error) {
	if b.deleteCluster.RegionID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RegionID")
	}
	if b.deleteCluster.ClusterID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "ClusterID")
	}
	params := make(map[string]interface{})
	params["region_id"] = b.deleteCluster.RegionID
	params["cluster_id"] = b.deleteCluster.ClusterID
	return params, nil
}

// RunTask builder pattern code
type RunTaskBuilder struct {
	runTask *RunTask
}

func NewRunTaskBuilder() *RunTaskBuilder {
	runTask := &RunTask{}
	b := &RunTaskBuilder{runTask: runTask}
	return b
}

func (b *RunTaskBuilder) ClusterID(clusterID string) *RunTaskBuilder {
	b.runTask.ClusterID = clusterID
	return b
}

func (b *RunTaskBuilder) Name(name string) *RunTaskBuilder {
	b.runTask.Name = name
	return b
}

func (b *RunTaskBuilder) Description(description string) *RunTaskBuilder {
	b.runTask.Description = description
	return b
}

func (b *RunTaskBuilder) Template(template string) *RunTaskBuilder {
	b.runTask.Template = template
	return b
}

func (b *RunTaskBuilder) Version(version string) *RunTaskBuilder {
	b.runTask.Version = version
	return b
}

func (b *RunTaskBuilder) Environment(environment map[string]string) *RunTaskBuilder {
	b.runTask.Environment = environment
	return b
}

func (b *RunTaskBuilder) LatestImage(latestImage bool) *RunTaskBuilder {
	b.runTask.LatestImage = latestImage
	return b
}

func (b *RunTaskBuilder) Build() (map[string]interface{}, error) {
	if b.runTask.ClusterID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "ClusterID")
	}
	if b.runTask.Name == "" {
		return nil, errors.New(aliauth.StrMissRequired + "Name")
	}
	if b.runTask.Template == "" {
		return nil, errors.New(aliauth.StrMissRequired + "Template")
	}
	params := make(map[string]interface{})
	params["cluster_id"] = b.runTask.ClusterID
	params["name"] = b.runTask.Name
	params["template"] = b.runTask.Template
	params["description"] = b.runTask.Description
	params["version"] = b.runTask.Version
	params["environment"] = b.runTask.Environment
	params["latest_image"] = b.runTask.LatestImage

	return params, nil
}

// StartTask builder pattern code
type StartTaskBuilder struct {
	startTask *StartTask
}

func NewStartTaskBuilder() *StartTaskBuilder {
	startTask := &StartTask{}
	b := &StartTaskBuilder{startTask: startTask}
	return b
}

func (b *StartTaskBuilder) ClusterID(clusterID string) *StartTaskBuilder {
	b.startTask.ClusterID = clusterID
	return b
}

func (b *StartTaskBuilder) Name(name string) *StartTaskBuilder {
	b.startTask.Name = name
	return b
}

func (b *StartTaskBuilder) Build() (map[string]interface{}, error) {
	if b.startTask.ClusterID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "ClusterID")
	}
	if b.startTask.Name == "" {
		return nil, errors.New(aliauth.StrMissRequired + "Name")
	}
	params := make(map[string]interface{})
	params["cluster_id"] = b.startTask.ClusterID
	params["name"] = b.startTask.Name
	return params, nil
}

// StopTask builder pattern code
type StopTaskBuilder struct {
	stopTask *StopTask
}

func NewStopTaskBuilder() *StopTaskBuilder {
	stopTask := &StopTask{}
	b := &StopTaskBuilder{stopTask: stopTask}
	return b
}

func (b *StopTaskBuilder) ClusterID(clusterID string) *StopTaskBuilder {
	b.stopTask.ClusterID = clusterID
	return b
}

func (b *StopTaskBuilder) Name(name string) *StopTaskBuilder {
	b.stopTask.Name = name
	return b
}

func (b *StopTaskBuilder) Timeout(timeout int) *StopTaskBuilder {
	b.stopTask.Timeout = timeout
	return b
}

func (b *StopTaskBuilder) Build() (map[string]interface{}, error) {
	if b.stopTask.ClusterID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "ClusterID")
	}
	if b.stopTask.Name == "" {
		return nil, errors.New(aliauth.StrMissRequired + "Name")
	}
	params := make(map[string]interface{})
	params["cluster_id"] = b.stopTask.ClusterID
	params["name"] = b.stopTask.Name
	params["timeout"] = b.stopTask.Timeout
	return params, nil
}
