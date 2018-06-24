package googlecontainer

import (
	"reflect"
)

//CreateClusterdictnoaryconvert create dictnoary for CreateCluster.
func CreateClusterdictnoaryconvert(option CreateCluster, CreateClusterjsonmap map[string]interface{}) {
	if option.Name != "" {
		CreateClusterjsonmap["name"] = option.Name
	}
	if option.Network != "" {
		CreateClusterjsonmap["network"] = option.Network
	}

	if option.LoggingService != "" {
		CreateClusterjsonmap["loggingService"] = option.LoggingService
	}

	if option.MonitoringService != "" {
		CreateClusterjsonmap["monitoringService"] = option.MonitoringService
	}

	if option.InitialClusterVersion != "" {
		CreateClusterjsonmap["initialClusterVersion"] = option.InitialClusterVersion
	}

	if option.Subnetwork != "" {
		CreateClusterjsonmap["subnetwork"] = option.Subnetwork
	}

	if option.Zone != "" {
		CreateClusterjsonmap["zone"] = option.Zone
	}

	CreateClusterLegacyAbacdictnoaryconvert(option, CreateClusterjsonmap)
	CreateClusterMasterAuthdictnoaryconvert(option, CreateClusterjsonmap)
	CreateClusterNodePoolsdictnoaryconvert(option, CreateClusterjsonmap)
}

//CreateClusterLegacyAbacdictnoaryconvert create dictnoary for CreateCluster.
func CreateClusterLegacyAbacdictnoaryconvert(option CreateCluster, CreateClusterjsonmap map[string]interface{}) {
	legacyAbac := make(map[string]interface{})

	if option.LegacyAbac.Enabled {
		legacyAbac["Enabled"] = option.LegacyAbac.Enabled
	}

	CreateClusterjsonmap["legacyAbac"] = legacyAbac
}

//CreateClusterMasterAuthdictnoaryconvert create dictnoary for CreateCluster.
func CreateClusterMasterAuthdictnoaryconvert(option CreateCluster, CreateClusterjsonmap map[string]interface{}) {
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

	CreateClusterjsonmap["legacyAbac"] = masterAuth

}

//IsEmpty check the config struct is empty or not.
func (c config) IsEmpty() bool {
	return reflect.DeepEqual(c, config{})
}

//CreateClusterNodePoolsdictnoaryconvert create dictnoary for CreateCluster.
func CreateClusterNodePoolsdictnoaryconvert(option CreateCluster, CreateClusterjsonmap map[string]interface{}) {
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
	CreateClusterjsonmap["nodePools"] = nodePools
}

//CreateServicedictnoaryconvert create dictnoary for CreateService.
func CreateServicedictnoaryconvert(option nodepool, CreateServicejsonmap map[string]interface{}) {

	if option.Name != "" {
		CreateServicejsonmap["name"] = option.Name
	}

	if option.StatusMessage != "" {
		CreateServicejsonmap["statusMessage"] = option.StatusMessage
	}

	if option.InitialNodeCount != 0 {
		CreateServicejsonmap["initialNodeCount"] = option.InitialNodeCount
	}

	if option.SelfLink != "" {
		CreateServicejsonmap["selfLink"] = option.SelfLink
	}

	if option.Version != "" {
		CreateServicejsonmap["version"] = option.Version
	}

	if option.Status != "" {
		CreateServicejsonmap["status"] = option.Status
	}
	CreateServiceConfigdictnoaryconvert(option, CreateServicejsonmap)
	CreateServiceAutoscalingdictnoaryconvert(option, CreateServicejsonmap)
}

//CreateServiceConfigdictnoaryconvert create dictnoary for CreateService.
func CreateServiceConfigdictnoaryconvert(option nodepool, CreateServicejsonmap map[string]interface{}) {

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

	CreateServicejsonmap["config"] = config
}

//CreateServiceManagementdictnoaryconvert create dictnoary for CreateService.
func CreateServiceManagementdictnoaryconvert(option nodepool, CreateServicejsonmap map[string]interface{}) {

	management := make(map[string]interface{})

	if option.Management.AutoRepair {
		management["autoRepair"] = option.Management.AutoRepair
	}

	if option.Management.AutoUpgrade {
		management["autoUpgrade"] = option.Management.AutoUpgrade
	}

	CreateServicejsonmap["management"] = management

}

//CreateServiceAutoscalingdictnoaryconvert create dictnoary for CreateService.
func CreateServiceAutoscalingdictnoaryconvert(option nodepool, CreateServicejsonmap map[string]interface{}) {
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

	CreateServicejsonmap["autoscaling"] = autoscaling
}
