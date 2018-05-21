package alicontainer

//Alicontainer struct represents Alicontainer attribute and methods associates with it.
type Alicontainer struct {
}

type CreateCluster struct {
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

type DeleteCluster struct {
}

type RunTask struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Template    string            `json:"template"`
	Version     string            `json:"version"`
	Environment map[string]string `json:"environment"`
	LatestImage bool              `json:"latest_image"`
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
