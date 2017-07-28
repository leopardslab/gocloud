package googlecontainer

type nodepool struct {
	Config struct {
		MachineType string        `json:"machineType"`
		ImageType   string        `json:"imageType"`
		OauthScopes []interface{} `json:"oauthScopes"`
		Preemptible bool          `json:"preemptible"`
		Labels      struct {
		} `json:"labels"`
		LocalSsdCount int `json:"localSsdCount"`
		Metadata      struct {
		} `json:"metadata"`
		DiskSizeGb     int           `json:"diskSizeGb"`
		Tags           []interface{} `json:"tags"`
		ServiceAccount string        `json:"serviceAccount"`
	} `json:"config"`
	Name          string `json:"name"`
	StatusMessage string `json:"statusMessage"`
	Autoscaling   struct {
		MaxNodeCount int  `json:"maxNodeCount"`
		MinNodeCount int  `json:"minNodeCount"`
		Enabled      bool `json:"enabled"`
	} `json:"autoscaling"`
	InitialNodeCount int `json:"initialNodeCount"`
	Management       struct {
		AutoRepair     bool `json:"autoRepair"`
		AutoUpgrade    bool `json:"autoUpgrade"`
		UpgradeOptions struct {
		} `json:"upgradeOptions"`
	} `json:"management"`
	SelfLink          string        `json:"selfLink"`
	Version           string        `json:"version"`
	InstanceGroupUrls []interface{} `json:"instanceGroupUrls"`
	Status            string        `json:"status"`
}
