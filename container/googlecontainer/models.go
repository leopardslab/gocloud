package googlecontainer

type Googlecontainer struct {
}

type Createcluster struct {
	Name                  string     `json:"name"`
	Zone                  string     `json:"zone"`
	Network               string     `json:"network"`
	LoggingService        string     `json:"loggingService"`
	MonitoringService     string     `json:"monitoringService"`
	InitialClusterVersion string     `json:"initialClusterVersion"`
	Subnetwork            string     `json:"subnetwork"`
	LegacyAbac            legacyAbac `json:"legacyAbac"`
	MasterAuth            masterAuth `json:"masterAuth"`
	NodePools             []NodePool `json:"nodePools"`
}

type legacyAbac struct {
	Enabled bool `json:"enabled"`
}

type masterAuth struct {
	Username                string                   `json:"username"`
	ClientCertificateConfig ClientCertificateConfigs `json:"clientCertificateConfig"`
}

type ClientCertificateConfigs struct {
	IssueClientCertificate bool `json:"issueClientCertificate"`
}

type LegacyAbac struct {
	Enabled bool `json:"enabled"`
}

type config struct {
	MachineType string   `json:"machineType"`
	ImageType   string   `json:"imageType"`
	DiskSizeGb  int      `json:"diskSizeGb"`
	Preemptible bool     `json:"preemptible"`
	OauthScopes []string `json:"oauthScopes"`
}

type autoscaling struct {
	Enabled bool `json:"enabled"`
}

type management struct {
	AutoUpgrade bool `json:"autoUpgrade"`
	AutoRepair  bool `json:"autoRepair"`
}

type NodePool struct {
	Name             string      `json:"name"`
	InitialNodeCount int         `json:"initialNodeCount"`
	Config           config      `json:"config"`
	Autoscaling      autoscaling `json:"autoscaling"`
	Management       management  `json:"management"`
}

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
