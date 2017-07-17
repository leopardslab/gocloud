package googlecontainer

type Googlecontainer struct {
}



type Createcluster struct {
	Name string `json:"name"`
	Zone string `json:"zone"`
	Network string `json:"network"`
	LoggingService string `json:"loggingService"`
	MonitoringService string `json:"monitoringService"`
	NodePools []struct {
		Name string `json:"name"`
		InitialNodeCount int `json:"initialNodeCount"`
		Config struct {
			MachineType string `json:"machineType"`
			ImageType string `json:"imageType"`
			DiskSizeGb int `json:"diskSizeGb"`
			Preemptible bool `json:"preemptible"`
			OauthScopes []string `json:"oauthScopes"`
		} `json:"config"`
		Autoscaling struct {
			Enabled bool `json:"enabled"`
		} `json:"autoscaling"`
		Management struct {
			AutoUpgrade bool `json:"autoUpgrade"`
			AutoRepair bool `json:"autoRepair"`
			UpgradeOptions struct {
			} `json:"upgradeOptions"`
		} `json:"management"`
	} `json:"nodePools"`
	InitialClusterVersion string `json:"initialClusterVersion"`
	MasterAuth struct {
		Username string `json:"username"`
		ClientCertificateConfig struct {
			IssueClientCertificate bool `json:"issueClientCertificate"`
		} `json:"clientCertificateConfig"`
	} `json:"masterAuth"`
	Subnetwork string `json:"subnetwork"`
	LegacyAbac struct {
		Enabled bool `json:"enabled"`
	} `json:"legacyAbac"`
}


func (googlecontainer *Googlecontainer) Createcontainer(request interface{}) (resp interface{}, err error) {

	return
}

func (googlecontainer *Googlecontainer) Deletecontainer(request interface{}) (resp interface{}, err error) {

	return
}

func (googlecontainer *Googlecontainer) Createservice(request interface{}) (resp interface{}, err error) {
	return
}

func (googlecontainer *Googlecontainer) Runtask(request interface{}) (resp interface{}, err error) {
	return
}

func (googlecontainer *Googlecontainer) Deleteservice(request interface{}) (resp interface{}, err error) {
	return
}


func (googlecontainer *Googlecontainer)Stoptask(request interface{}) (resp interface{}, err error){
	return
}
