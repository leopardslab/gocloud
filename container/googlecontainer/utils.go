package googlecontainer

import (
	"reflect"
)

//Createclusterdictnoaryconvert create dictnoary for Createcluster.
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

	CreateclusterLegacyAbacdictnoaryconvert(option, Createclusterjsonmap)
	CreateclusterMasterAuthdictnoaryconvert(option, Createclusterjsonmap)
	CreateclusterNodePoolsdictnoaryconvert(option, Createclusterjsonmap)
}

//CreateclusterLegacyAbacdictnoaryconvert create dictnoary for Createcluster.
func CreateclusterLegacyAbacdictnoaryconvert(option Createcluster, Createclusterjsonmap map[string]interface{}) {
	legacyAbac := make(map[string]interface{})

	if option.LegacyAbac.Enabled {
		legacyAbac["Enabled"] = option.LegacyAbac.Enabled
	}

	Createclusterjsonmap["legacyAbac"] = legacyAbac
}

//CreateclusterMasterAuthdictnoaryconvert create dictnoary for Createcluster.
func CreateclusterMasterAuthdictnoaryconvert(option Createcluster, Createclusterjsonmap map[string]interface{}) {
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

//IsEmpty check the config struct is empty or not.
func (c config) IsEmpty() bool {
	return reflect.DeepEqual(c, config{})
}

//CreateclusterNodePoolsdictnoaryconvert create dictnoary for Createcluster.
func CreateclusterNodePoolsdictnoaryconvert(option Createcluster, Createclusterjsonmap map[string]interface{}) {
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
			if !option.NodePools[i].Config.IsEmpty() {
				config := make(map[string]interface{})
				if (option.NodePools[i].Config.MachineType) != "" {
					config["MachineType"] = option.NodePools[i].Config.MachineType
				}

				if (option.NodePools[i].Config.DiskSizeGb) != 0 {
					config["diskSizeGb"] = option.NodePools[i].Config.DiskSizeGb
				}

				if option.NodePools[i].Config.Preemptible {
					config["preemptible"] = option.NodePools[i].Config.Preemptible
				}

				if (len(option.NodePools[i].Config.OauthScopes)) != 0 {
					config["oauthScopes"] = option.NodePools[i].Config.OauthScopes
				}

				if (option.NodePools[i].Config.ImageType) != "" {
					config["oauthScopes"] = option.NodePools[i].Config.ImageType
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

//Createservicedictnoaryconvert create dictnoary for Createservice.
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

//CreateserviceConfigdictnoaryconvert create dictnoary for Createservice.
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

//CreateserviceManagementdictnoaryconvert create dictnoary for Createservice.
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

//CreateserviceAutoscalingdictnoaryconvert create dictnoary for Createservice.
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
