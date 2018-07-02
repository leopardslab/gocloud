package googlecontainer

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/cloudlibz/gocloud/googleauth"
	"io/ioutil"
	"net/http"
)

//CreateCluster creates cluster.
func (googlecontainer *Googlecontainer) CreateCluster(request interface{}) (resp interface{}, err error) {

	var option CreateCluster

	var Projectid string

	var Zone string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "Project":
			Projectid, _ = value.(string)

		case "Name":
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

	//zonevalue := "projects/" + Projectid + "/zones/" + Zone
	option.Zone = Zone

	CreateClusterjsonmap := make(map[string]interface{})

	CreateClusterdictnoaryconvert(option, CreateClusterjsonmap)

	CreateClusterdict := make(map[string]interface{})

	CreateClusterdict["cluster"] = CreateClusterjsonmap

	CreateClusterjson, _ := json.Marshal(CreateClusterdict)

	CreateClusterjsonstring := string(CreateClusterjson)

	var CreateClusterjsonstringbyte = []byte(CreateClusterjsonstring)

	url := "https://container.googleapis.com/v1/projects/" + Projectid + "/zones/" + Zone + "/clusters"

	client := googleauth.SignJWT()

	CreateClusterrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(CreateClusterjsonstringbyte))

	CreateClusterrequest.Header.Set("Content-Type", "application/json")

	CreateClusterresp, err := client.Do(CreateClusterrequest)

	fmt.Println(err)
	defer CreateClusterresp.Body.Close()

	body, err := ioutil.ReadAll(CreateClusterresp.Body)

	CreateClusteresponse := make(map[string]interface{})
	CreateClusteresponse["status"] = CreateClusterresp.StatusCode
	CreateClusteresponse["body"] = string(body)
	resp = CreateClusteresponse
	return resp, err
}

//DeleteCluster deletes cluster.
func (googlecontainer *Googlecontainer) DeleteCluster(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://container.googleapis.com/v1/projects/" + options["Project"] + "/zones/" + options["Zone"] + "/clusters/" + options["clusterId"]

	client := googleauth.SignJWT()

	DeleteClusterrequest, err := http.NewRequest("DELETE", url, nil)

	DeleteClusterrequest.Header.Set("Content-Type", "application/json")

	DeleteClusterresp, err := client.Do(DeleteClusterrequest)

	defer DeleteClusterresp.Body.Close()

	body, err := ioutil.ReadAll(DeleteClusterresp.Body)

	DeleteClusterresponse := make(map[string]interface{})
	DeleteClusterresponse["status"] = DeleteClusterresp.StatusCode
	DeleteClusterresponse["body"] = string(body)
	resp = DeleteClusterresponse
	return resp, err
}

//CreateService crestes loadbalancer service.
func (googlecontainer *Googlecontainer) CreateService(request interface{}) (resp interface{}, err error) {

	var option nodepool

	var Projectid string

	var ClusterId string

	var Zone string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "Project":
			Projectid, _ = value.(string)

		case "Name":
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

	CreateServicejsonmap := make(map[string]interface{})

	CreateServicedictnoaryconvert(option, CreateServicejsonmap)

	CreateServicedict := make(map[string]interface{})

	CreateServicedict["nodePool"] = CreateServicejsonmap

	CreateServicejson, _ := json.Marshal(CreateServicedict)

	CreateServicejsonstring := string(CreateServicejson)

	var CreateServicejsonstringbyte = []byte(CreateServicejsonstring)

	url := "https://container.googleapis.com/v1/projects/" + Projectid + "/zones/" + Zone + "/clusters/" + ClusterId + "/nodePools"

	client := googleauth.SignJWT()

	CreateServicerequest, err := http.NewRequest("POST", url, bytes.NewBuffer(CreateServicejsonstringbyte))

	CreateServicerequest.Header.Set("Content-Type", "application/json")

	CreateServicerresp, err := client.Do(CreateServicerequest)

	defer CreateServicerresp.Body.Close()

	body, err := ioutil.ReadAll(CreateServicerresp.Body)

	CreateServiceresponse := make(map[string]interface{})
	CreateServiceresponse["status"] = CreateServicerresp.StatusCode
	CreateServiceresponse["body"] = string(body)
	resp = CreateServiceresponse
	return resp, err
}

//runtask runs container.
func (googlecontainer *Googlecontainer) RunTask(request interface{}) (resp interface{}, err error) {
	fmt.Println("\nThis API is not provided by Google cloud")
	return
}

//StartTask start container.
func (googlecontainer *Googlecontainer) StartTask(request interface{}) (resp interface{}, err error) {
	fmt.Println("\nThis API is not provided by Google cloud")
	return
}

//DeleteService crestes loadbalancer service.
func (googlecontainer *Googlecontainer) DeleteService(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://container.googleapis.com/v1/projects/" + options["Project"] + "/zones/" + options["Zone"] + "/clusters/" + options["clusterId"] + "/nodePools/" + options["nodePoolId"]

	client := googleauth.SignJWT()

	DeleteServicerequest, err := http.NewRequest("DELETE", url, nil)
	DeleteServicerequest.Header.Set("Content-Type", "application/json")

	DeleteServiceresp, err := client.Do(DeleteServicerequest)

	defer DeleteServiceresp.Body.Close()

	body, err := ioutil.ReadAll(DeleteServiceresp.Body)

	DeleteServiceresponse := make(map[string]interface{})
	DeleteServiceresponse["status"] = DeleteServiceresp.StatusCode
	DeleteServiceresponse["body"] = string(body)
	resp = DeleteServiceresponse
	return resp, err
}

//StopTask stops container.
func (googlecontainer *Googlecontainer) StopTask(request interface{}) (resp interface{}, err error) {
	options := request.(map[string]string)

	url := "https://container.googleapis.com/v1/projects/" + options["Project"] + "/zones/" + options["Zone"] + "/operations/" + options["OperationId"] + ":cancel"

	client := googleauth.SignJWT()

	StopTaskrequest, err := http.NewRequest("POST", url, nil)
	StopTaskrequest.Header.Set("Content-Type", "application/json")

	StopTaskresp, err := client.Do(StopTaskrequest)

	defer StopTaskresp.Body.Close()

	body, err := ioutil.ReadAll(StopTaskresp.Body)

	StopTaskrespresponse := make(map[string]interface{})
	StopTaskrespresponse["status"] = StopTaskresp.StatusCode
	StopTaskrespresponse["body"] = string(body)
	resp = StopTaskrespresponse
	return resp, err
}
