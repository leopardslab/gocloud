package googlecontainer

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/scorelab/gocloud-v2/googleauth"
	"io/ioutil"
	"net/http"
)

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
