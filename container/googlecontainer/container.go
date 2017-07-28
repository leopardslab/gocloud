package googlecontainer

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/scorelab/gocloud-v2/googleauth"
	"io/ioutil"
	"net/http"
	 "reflect"
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

	CreateserviceLegacyAbacdictnoaryconvert(option, Createclusterjsonmap)
	CreateserviceMasterAuthdictnoaryconvert(option, Createclusterjsonmap)
	CreateserviceNodePoolsdictnoaryconvert(option, Createclusterjsonmap)
}

func CreateserviceLegacyAbacdictnoaryconvert(option Createcluster, Createclusterjsonmap map[string]interface{}) {
	legacyAbac := make(map[string]interface{})

	if option.LegacyAbac.Enabled {
		legacyAbac["Enabled"] = option.LegacyAbac.Enabled
	}

	Createclusterjsonmap["legacyAbac"] = legacyAbac
}

func CreateserviceMasterAuthdictnoaryconvert(option Createcluster, Createclusterjsonmap map[string]interface{}) {
	masterAuth := make(map[string]interface{})

	if option.MasterAuth.Username != "" {
		masterAuth["username"] = option.MasterAuth.Username
	}

	if option.MasterAuth.ClientCertificateConfig != (ClientCertificateConfigs{}) {
		clientCertificateConfigs := make(map[string]interface{})
		if option.MasterAuth.ClientCertificateConfig.IssueClientCertificate {
			clientCertificateConfigs["clientCertificateConfig"] = option.MasterAuth.ClientCertificateConfig
		}
		masterAuth["clientCertificateConfig"] = clientCertificateConfigs
	}

	Createclusterjsonmap["legacyAbac"] = masterAuth

}

func (c config) IsEmpty() bool {
  return reflect.DeepEqual(c,config{})
}

func CreateserviceNodePoolsdictnoaryconvert(option Createcluster, Createclusterjsonmap map[string]interface{}) {
	nodePools := make([]map[string]interface{}, 0)
	if len(option.NodePools) != 0 {
		for i := 0; i < len(option.NodePools); i++ {
			nodePool := make(map[string]interface{})
			if option.NodePools[i].Name != "" {
				nodePool["name"] = option.NodePools[i].Name
			}
			if option.NodePools[i].InitialNodeCount != 0 {
				nodePool["initialNodeCount"] = option.NodePools[i].InitialNodeCount
			}
			if (!option.NodePools[i].Config.IsEmpty()){
				config := make(map[string]interface{})
				if(option.NodePools[i].Config.MachineType) != ""{
					config["MachineType"]= option.NodePools[i].Config.MachineType
				}

				if(option.NodePools[i].Config.DiskSizeGb) != 0 {
					config["diskSizeGb"]= option.NodePools[i].Config.DiskSizeGb
				}

				if(option.NodePools[i].Config.Preemptible){
					config["preemptible"]= option.NodePools[i].Config.Preemptible
				}

				if(len(option.NodePools[i].Config.OauthScopes)) != 0{
					config["oauthScopes"]= option.NodePools[i].Config.OauthScopes
				}

				if(option.NodePools[i].Config.ImageType) != ""{
					config["oauthScopes"]= option.NodePools[i].Config.ImageType
				}

				nodePool["config"] = config
			}
			if (option.NodePools[i].Autoscaling) != (autoscaling{}) {
				autoscaling := make(map[string]interface{})
				if option.NodePools[i].Autoscaling.Enabled {
					autoscaling["Enabled"] = option.NodePools[i].Autoscaling.Enabled
				}
				nodePool["autoscaling"] = autoscaling
			}
			if (option.NodePools[i].Management) != (management{}) {
				management := make(map[string]interface{})
				if option.NodePools[i].Management.AutoRepair {
					management["autoRepair"] = option.NodePools[i].Management.AutoRepair
				}

				if option.NodePools[i].Management.AutoUpgrade {
					management["autoUpgrade"] = option.NodePools[i].Management.AutoRepair
				}

				nodePool["management"] = management
			}
			nodePools = append(nodePools, nodePool)
		}
	}
	Createclusterjsonmap["nodePools"] = nodePools
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

		case "config":
			configV, _ := value.(map[string]interface{})
			for key, value := range configV {
				switch key {
				case "machineType":
					machineTypeV, _ := value.(string)
					option.Config.MachineType = machineTypeV

				case "imageType":
					imageTypeV, _ := value.(string)
					option.Config.ImageType = imageTypeV

				case "diskSizeGb":
					DiskSizeGbV, _ := value.(int)
					option.Config.DiskSizeGb = DiskSizeGbV

				case "preemptible":
					preemptibleV, _ := value.(bool)
					option.Config.Preemptible = preemptibleV

				case "oauthScopes":
					oauthScopesV, _ := value.([]string)
					option.Config.OauthScopes = oauthScopesV

				case "ServiceAccount":
					ServiceAccountV, _ := value.(string)
					option.Config.ServiceAccount = ServiceAccountV

				case "localSsdCount":
					LocalSsdCountV, _ := value.(int)
					option.Config.LocalSsdCount = LocalSsdCountV
				}
			}

		case "autoscaling":
			autoscalingV, _ := value.(map[string]interface{})
			for key, value := range autoscalingV {
				switch key {
				case "enabled":
					enabledV, _ := value.(bool)
					option.Autoscaling.Enabled = enabledV

				case "minNodeCount":
					minNodeCountV, _ := value.(int)
					option.Autoscaling.MinNodeCount = minNodeCountV

				case "maxNodeCount":
					maxNodeCountV, _ := value.(int)
					option.Autoscaling.MaxNodeCount = maxNodeCountV

				}
			}

		case "instanceGroupUrls":
			autoscalingV, _ := value.([]string)
			option.InstanceGroupUrls = autoscalingV

		case "management":
			managementV, _ := value.(map[string]interface{})
			for key, value := range managementV {
				switch key {
				case "autoUpgrade":
					autoUpgradeV, _ := value.(bool)
					option.Management.AutoUpgrade = autoUpgradeV

				case "AutoRepair":
					autoRepairV, _ := value.(bool)
					option.Management.AutoRepair = autoRepairV

				}
			}
		}
	}

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
		Createservicejsonmap["name"] = option.Name
	}

	if option.StatusMessage != "" {
		Createservicejsonmap["statusMessage"] = option.StatusMessage
	}

	if option.InitialNodeCount != 0 {
		Createservicejsonmap["initialNodeCount"] = option.InitialNodeCount
	}

	if option.SelfLink != "" {
		Createservicejsonmap["selfLink"] = option.SelfLink
	}

	if option.Version != "" {
		Createservicejsonmap["version"] = option.Version
	}

	if option.Status != "" {
		Createservicejsonmap["status"] = option.Status
	}
	CreateserviceConfigdictnoaryconvert(option, Createservicejsonmap)
	CreateserviceAutoscalingdictnoaryconvert(option, Createservicejsonmap)
}

func CreateserviceConfigdictnoaryconvert(option nodepool, Createservicejsonmap map[string]interface{}) {

	config := make(map[string]interface{})

	if option.Config.MachineType != "" {
		config["machineType"] = option.Config.MachineType
	}

	if option.Config.ImageType != "" {
		config["imageType"] = option.Config.ImageType
	}

	if len(option.Config.OauthScopes) != 0 {
		config["oauthScopes"] = option.Config.OauthScopes
	}

	if option.Config.Preemptible {
		config["preemptible"] = option.Config.Preemptible
	}

	if option.Config.ServiceAccount != "" {
		config["serviceAccount"] = option.Config.ServiceAccount
	}

	if option.Config.DiskSizeGb != 0 {
		config["diskSizeGb"] = option.Config.DiskSizeGb
	}

	if option.Config.LocalSsdCount != 0 {
		config["localSsdCount"] = option.Config.LocalSsdCount
	}

	Createservicejsonmap["config"] = config
}

func CreateserviceManagementdictnoaryconvert(option nodepool, Createservicejsonmap map[string]interface{}) {

	management := make(map[string]interface{})

	if option.Management.AutoRepair {
		management["autoRepair"] = option.Management.AutoRepair
	}

	if option.Management.AutoUpgrade {
		management["autoUpgrade"] = option.Management.AutoUpgrade
	}

	Createservicejsonmap["management"] = management

}

func CreateserviceAutoscalingdictnoaryconvert(option nodepool, Createservicejsonmap map[string]interface{}) {
	autoscaling := make(map[string]interface{})

	if option.Autoscaling.MaxNodeCount != 0 {
		autoscaling["maxNodeCount"] = option.Autoscaling.MaxNodeCount
	}

	if option.Autoscaling.MinNodeCount != 0 {
		autoscaling["minNodeCount"] = option.Autoscaling.MinNodeCount
	}

	if option.Autoscaling.Enabled {
		autoscaling["enabled"] = option.Autoscaling.Enabled
	}

	Createservicejsonmap["autoscaling"] = autoscaling
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
