package googlecontainer

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/shlokgilda/gocloud/googleauth"
	"io/ioutil"
	"net/http"
)

//Createcluster creates cluster.
func (googlecontainer *Googlecontainer) Createcluster(request interface{}) (resp interface{}, err error) {

	var option Createcluster

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

	Createclusterjsonmap := make(map[string]interface{})

	Createclusterdictnoaryconvert(option, Createclusterjsonmap)

	Createclusterdict := make(map[string]interface{})

	Createclusterdict["cluster"] = Createclusterjsonmap

	Createclusterjson, _ := json.Marshal(Createclusterdict)

	Createclusterjsonstring := string(Createclusterjson)

	var Createclusterjsonstringbyte = []byte(Createclusterjsonstring)

	url := "https://container.googleapis.com/v1/projects/" + Projectid + "/zones/" + Zone + "/clusters"

	client := googleauth.SignJWT()

	Createclusterrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Createclusterjsonstringbyte))

	Createclusterrequest.Header.Set("Content-Type", "application/json")

	Createclusterresp, err := client.Do(Createclusterrequest)

	fmt.Println(err)
	defer Createclusterresp.Body.Close()

	body, err := ioutil.ReadAll(Createclusterresp.Body)

	Createclusteresponse := make(map[string]interface{})
	Createclusteresponse["status"] = Createclusterresp.StatusCode
	Createclusteresponse["body"] = string(body)
	resp = Createclusteresponse
	return resp, err
}

//Deletecluster deletes cluster.
func (googlecontainer *Googlecontainer) Deletecluster(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://container.googleapis.com/v1/projects/" + options["Project"] + "/zones/" + options["Zone"] + "/clusters/" + options["clusterId"]

	client := googleauth.SignJWT()

	Deleteclusterrequest, err := http.NewRequest("DELETE", url, nil)

	Deleteclusterrequest.Header.Set("Content-Type", "application/json")

	Deleteclusterresp, err := client.Do(Deleteclusterrequest)

	defer Deleteclusterresp.Body.Close()

	body, err := ioutil.ReadAll(Deleteclusterresp.Body)

	Deleteclusterresponse := make(map[string]interface{})
	Deleteclusterresponse["status"] = Deleteclusterresp.StatusCode
	Deleteclusterresponse["body"] = string(body)
	resp = Deleteclusterresponse
	return resp, err
}

//Createservice crestes loadbalancer service.
func (googlecontainer *Googlecontainer) Createservice(request interface{}) (resp interface{}, err error) {

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

	Createservicejsonmap := make(map[string]interface{})

	Createservicedictnoaryconvert(option, Createservicejsonmap)

	Createservicedict := make(map[string]interface{})

	Createservicedict["nodePool"] = Createservicejsonmap

	Createservicejson, _ := json.Marshal(Createservicedict)

	Createservicejsonstring := string(Createservicejson)

	var Createservicejsonstringbyte = []byte(Createservicejsonstring)

	url := "https://container.googleapis.com/v1/projects/" + Projectid + "/zones/" + Zone + "/clusters/" + ClusterId + "/nodePools"

	client := googleauth.SignJWT()

	Createservicerequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Createservicejsonstringbyte))

	Createservicerequest.Header.Set("Content-Type", "application/json")

	Createservicerresp, err := client.Do(Createservicerequest)

	defer Createservicerresp.Body.Close()

	body, err := ioutil.ReadAll(Createservicerresp.Body)

	Createserviceresponse := make(map[string]interface{})
	Createserviceresponse["status"] = Createservicerresp.StatusCode
	Createserviceresponse["body"] = string(body)
	resp = Createserviceresponse
	return resp, err
}

//runtask runs container.
func (googlecontainer *Googlecontainer) Runtask(request interface{}) (resp interface{}, err error) {
	fmt.Println("\nThis API is not provided by Google cloud")
	return
}

//Starttask start container.
func (googlecontainer *Googlecontainer) Starttask(request interface{}) (resp interface{}, err error) {
	fmt.Println("\nThis API is not provided by Google cloud")
	return
}

//Deleteservice crestes loadbalancer service.
func (googlecontainer *Googlecontainer) Deleteservice(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://container.googleapis.com/v1/projects/" + options["Project"] + "/zones/" + options["Zone"] + "/clusters/" + options["clusterId"] + "/nodePools/" + options["nodePoolId"]

	client := googleauth.SignJWT()

	Deleteservicerequest, err := http.NewRequest("DELETE", url, nil)
	Deleteservicerequest.Header.Set("Content-Type", "application/json")

	Deleteserviceresp, err := client.Do(Deleteservicerequest)

	defer Deleteserviceresp.Body.Close()

	body, err := ioutil.ReadAll(Deleteserviceresp.Body)

	Deleteserviceresponse := make(map[string]interface{})
	Deleteserviceresponse["status"] = Deleteserviceresp.StatusCode
	Deleteserviceresponse["body"] = string(body)
	resp = Deleteserviceresponse
	return resp, err
}

//Stoptask stops container.
func (googlecontainer *Googlecontainer) Stoptask(request interface{}) (resp interface{}, err error) {
	options := request.(map[string]string)

	url := "https://container.googleapis.com/v1/projects/" + options["Project"] + "/zones/" + options["Zone"] + "/operations/" + options["OperationId"] + ":cancel"

	client := googleauth.SignJWT()

	Stoptaskrequest, err := http.NewRequest("POST", url, nil)
	Stoptaskrequest.Header.Set("Content-Type", "application/json")

	Stoptaskresp, err := client.Do(Stoptaskrequest)

	defer Stoptaskresp.Body.Close()

	body, err := ioutil.ReadAll(Stoptaskresp.Body)

	Stoptaskrespresponse := make(map[string]interface{})
	Stoptaskrespresponse["status"] = Stoptaskresp.StatusCode
	Stoptaskrespresponse["body"] = string(body)
	resp = Stoptaskrespresponse
	return resp, err
}
