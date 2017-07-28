package googlecontainer

type nodepool struct {
	Config struct {
		MachineType    string   `json:"machineType"`
		ImageType      string   `json:"imageType"`
		OauthScopes    []string `json:"oauthScopes"`
		Preemptible    bool     `json:"preemptible"`
		LocalSsdCount  int      `json:"localSsdCount"`
		DiskSizeGb     int      `json:"diskSizeGb"`
		ServiceAccount string   `json:"serviceAccount"`
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
		AutoRepair  bool `json:"autoRepair"`
		AutoUpgrade bool `json:"autoUpgrade"`
	} `json:"management"`
	SelfLink          string   `json:"selfLink"`
	Version           string   `json:"version"`
	InstanceGroupUrls []string `json:"instanceGroupUrls"`
	Status            string   `json:"status"`
}
