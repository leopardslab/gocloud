package googlecontainer

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/scorelab/gocloud-v2/googleauth"
	"io/ioutil"
	"net/http"
)

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

type upgradeOptions struct {
}

type management struct {
	AutoUpgrade    bool           `json:"autoUpgrade"`
	AutoRepair     bool           `json:"autoRepair"`
	UpgradeOptions upgradeOptions `json:"upgradeOptions"`
}

type NodePool struct {
	Name             string      `json:"name"`
	InitialNodeCount int         `json:"initialNodeCount"`
	Config           config      `json:"config"`
	Autoscaling      autoscaling `json:"autoscaling"`
	Management       management  `json:"management"`
}

func (googlecontainer *Googlecontainer) Deletecontainer(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://container.googleapis.com/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/clusters/" + options["clusterId"]

	client := googleauth.SignJWT()

	Deletecontainerrequest, err := http.NewRequest("DELETE", url, nil)

	Deletecontainerrequest.Header.Set("Content-Type", "application/json")

	Deletecontainerresp, err := client.Do(Deletecontainerrequest)

	defer Deletecontainerresp.Body.Close()

	body, err := ioutil.ReadAll(Deletecontainerresp.Body)

	fmt.Println(string(body))

	return
}

func (googlecontainer *Googlecontainer) Createcontainer(request interface{}) (resp interface{}, err error) {

	var option Createcluster

	var Projectid string

	var Zone string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "projectid":
			Projectid, _ = value.(string)

		case "name":
			name, _ := value.(string)
			option.Name = name

		case "Zone":
			ZoneV, _ := value.(string)
			Zone = ZoneV

		case "network":
			NetworkV, _ := value.(string)
			option.Network = NetworkV

		case "loggingService":
			LoggingServiceV, _ := value.(string)
			option.LoggingService = LoggingServiceV

		case "monitoringService":
			MonitoringServiceV, _ := value.(string)
			option.MonitoringService = MonitoringServiceV

		case "initialClusterVersion":
			InitialClusterVersionV, _ := value.(string)
			option.InitialClusterVersion = InitialClusterVersionV

		case "subnetwork":
			SubnetworkV, _ := value.(string)
			option.Subnetwork = SubnetworkV

		case "masterAuth":
			masterAuthV, _ := value.(map[string]interface{})
			for key, value := range masterAuthV {
				switch key {
				case "username":
					usernameV, _ := value.(string)
					option.MasterAuth.Username = usernameV
				case "clientCertificateConfig":
					clientCertificateConfigV, _ := value.(map[string]interface{})
					for key, value := range clientCertificateConfigV {
						switch key {
						case "issueClientCertificate":
							issueClientCertificateV, _ := value.(bool)
							option.MasterAuth.ClientCertificateConfig.IssueClientCertificate = issueClientCertificateV
						}
					}
				}
			}

		case "nodePools":
			nodePoolsV, _ := value.([]map[string]interface{})
			for i := 0; i < len(nodePoolsV); i++ {
				var nodePool NodePool
				for key, value := range nodePoolsV[i] {
					switch key {
					case "name":
						nameV, _ := value.(string)
						nodePool.Name = nameV

					case "initialNodeCount":
						InitialNodeCountV, _ := value.(int)
						nodePool.InitialNodeCount = InitialNodeCountV

					case "config":
						configV, _ := value.(map[string]interface{})
						for key, value := range configV {
							switch key {
							case "machineType":
								machineTypeV, _ := value.(string)
								nodePool.Config.MachineType = machineTypeV

							case "imageType":
								imageTypeV, _ := value.(string)
								nodePool.Config.ImageType = imageTypeV

							case "diskSizeGb":
								DiskSizeGbV, _ := value.(int)
								nodePool.Config.DiskSizeGb = DiskSizeGbV

							case "preemptible":
								preemptibleV, _ := value.(bool)
								nodePool.Config.Preemptible = preemptibleV

							case "oauthScopes":
								oauthScopesV, _ := value.([]string)
								nodePool.Config.OauthScopes = oauthScopesV
							}
						}

					case "autoscaling":
						autoscalingV, _ := value.(map[string]interface{})
						for key, value := range autoscalingV {
							switch key {
							case "enabled":
								enabledV, _ := value.(bool)
								nodePool.Autoscaling.Enabled = enabledV
							}
						}

					case "management":
						managementV, _ := value.(map[string]interface{})
						for key, value := range managementV {
							switch key {
							case "autoUpgrade":
								autoUpgradeV, _ := value.(bool)
								nodePool.Management.AutoUpgrade = autoUpgradeV

							case "AutoRepair":
								autoRepairV, _ := value.(bool)
								nodePool.Management.AutoRepair = autoRepairV

							case "UpgradeOptions":
							}
						}

					}
				}
				option.NodePools = append(option.NodePools, nodePool)
			}
		}
	}

	zonevalue := "projects/" + Projectid + "/zones/" + Zone
	option.Zone = zonevalue

	Createclusterjsonmap := make(map[string]interface{})

	Createclusterdictnoaryconvert(option, Createclusterjsonmap)

	Createclusterjson, _ := json.Marshal(Createclusterjsonmap)

	Createclusterjsonstring := string(Createclusterjson)

	fmt.Println(Createclusterjsonstring)

	var Createclusterjsonstringbyte = []byte(Createclusterjsonstring)

	url := "https://container.googleapis.com/v1/projects/" + Projectid + "/zones/" + Zone + "/clusters"

	client := googleauth.SignJWT()

	Createclusterrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Createclusterjsonstringbyte))

	Createclusterrequest.Header.Set("Content-Type", "application/json")

	Createclusterresp, err := client.Do(Createclusterrequest)

	defer Createclusterresp.Body.Close()

	body, err := ioutil.ReadAll(Createclusterresp.Body)

	fmt.Println(string(body))

	return
}

func Createclusterdictnoaryconvert(option Createcluster, Createclusterjsonmap map[string]interface{}) {
	if option.Name != "" {
		Createclusterjsonmap["name"] = option.Name
	}
	if option.Network != "" {
		Createclusterjsonmap["network"] = option.Network
	}

	if option.LoggingService != "" {
		Createclusterjsonmap["loggingService"] = option.LoggingService
	}

	if option.MonitoringService != "" {
		Createclusterjsonmap["monitoringService"] = option.MonitoringService
	}

	if option.InitialClusterVersion != "" {
		Createclusterjsonmap["initialClusterVersion"] = option.InitialClusterVersion
	}

	if option.Subnetwork != "" {
		Createclusterjsonmap["subnetwork"] = option.Subnetwork
	}

	if option.Zone != "" {
		Createclusterjsonmap["zone"] = option.Zone
	}
}

func (googlecontainer *Googlecontainer) Createservice(request interface{}) (resp interface{}, err error) {

	var option nodepool

	var Projectid string

	var ClusterId string

	var Zone string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "projectid":
			Projectid, _ = value.(string)

		case "name":
			name, _ := value.(string)
			option.Name = name

		case "Zone":
			ZoneV, _ := value.(string)
			Zone = ZoneV

		case "clusterId":
			ClusterIdV, _ := value.(string)
			ClusterId = ClusterIdV

		case "statusMessage":
			StatusMessageV, _ := value.(string)
			option.StatusMessage = StatusMessageV

		case "initialNodeCount":
			InitialNodeCountV, _ := value.(int)
			option.InitialNodeCount = InitialNodeCountV

		case "selfLink":
			SelfLinkV, _ := value.(string)
			option.SelfLink = SelfLinkV

		case "version":
			VersionV, _ := value.(string)
			option.Version = VersionV

		case "status":
			StatusV, _ := value.(string)
			option.Status = StatusV

		}
	}

	option.Config.MachineType = "Hello"
	fmt.Println(option.Config.MachineType)

	//zonevalue := "projects/" + Projectid + "/zones/" + Zone
	//option.Zone = zonevalue

	Createservicejsonmap := make(map[string]interface{})

	Createservicedictnoaryconvert(option, Createservicejsonmap)

	Createservicejson, _ := json.Marshal(Createservicejsonmap)

	Createservicejsonstring := string(Createservicejson)

	fmt.Println(Createservicejsonstring)

	var Createservicejsonstringbyte = []byte(Createservicejsonstring)

	url := "https://container.googleapis.com/v1/projects/" + Projectid + "/zones/" + Zone + "/clusters" + ClusterId + "/nodePools"

	client := googleauth.SignJWT()

	Createservicerequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Createservicejsonstringbyte))

	Createservicerequest.Header.Set("Content-Type", "application/json")

	Createservicerresp, err := client.Do(Createservicerequest)

	defer Createservicerresp.Body.Close()

	body, err := ioutil.ReadAll(Createservicerresp.Body)

	fmt.Println(string(body))

	return
}

func Createservicedictnoaryconvert(option nodepool, Createservicejsonmap map[string]interface{}) {

	if option.Name != "" {
		Createclusterjsonmap["name"] = option.Name
	}

	if option.StatusMessage != "" {
		Createclusterjsonmap["statusMessage"] = option.StatusMessage
	}

	if option.InitialNodeCount != 0 {
		Createclusterjsonmap["initialNodeCount"] = option.InitialNodeCount
	}

	if option.SelfLink != "" {
		Createclusterjsonmap["selfLink"] = option.SelfLink
	}

	if option.Version != "" {
		Createclusterjsonmap["version"] = option.Version
	}

	if option.Status != "" {
		Createclusterjsonmap["status"] = option.Status
	}

}

func (googlecontainer *Googlecontainer) Runtask(request interface{}) (resp interface{}, err error) {
	fmt.Println("\nThis API is not provided by Google cloud")
	return
}

func (googlecontainer *Googlecontainer) Starttask(request interface{}) (resp interface{}, err error) {
	fmt.Println("\nThis API is not provided by Google cloud")
	return
}
func (googlecontainer *Googlecontainer) Deleteservice(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://container.googleapis.com/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/clusters/" + options["clusterId"] + "/nodePools/" + options["nodePoolId"]

	client := googleauth.SignJWT()

	Deleteservicerequest, err := http.NewRequest("POST", url, nil)
	Deleteservicerequest.Header.Set("Content-Type", "application/json")

	Deleteserviceresp, err := client.Do(Deleteservicerequest)

	defer Deleteserviceresp.Body.Close()

	body, err := ioutil.ReadAll(Deleteserviceresp.Body)

	fmt.Println(string(body))

	return
}

func (googlecontainer *Googlecontainer) Stoptask(request interface{}) (resp interface{}, err error) {
	options := request.(map[string]string)

	url := "https://container.googleapis.com/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/operations/" + options["operationId"] + ":cancel"

	client := googleauth.SignJWT()

	Stoptaskrequest, err := http.NewRequest("POST", url, nil)
	Stoptaskrequest.Header.Set("Content-Type", "application/json")

	Stoptaskresp, err := client.Do(Stoptaskrequest)

	defer Stoptaskresp.Body.Close()

	body, err := ioutil.ReadAll(Stoptaskresp.Body)

	fmt.Println(string(body))

	return
}
