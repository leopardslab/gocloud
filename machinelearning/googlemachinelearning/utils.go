package googlemachinelearning


func createMLModeldictnoaryconvert(option CreateMLModel, createMLModeljsonmap map[string]interface{}){

		if option.name != ""{
			createMLModeljsonmap["name"] =  option.name
		}

		if option.description != ""{
			createMLModeljsonmap["description"] =  option.description
		}

		if len(option.regions) != 0 {
			createMLModeljsonmap["regions"] =  option.regions
		}

		if option.onlinePredictionLogging != false {
			createMLModeljsonmap["onlinePredictionLogging"] =  option.onlinePredictionLogging
		}


}


func preparedefaultVersion(option CreateMLModel, createMLModeljsonmap map[string]interface{}){

		defaultVersiondictmap := make(map[string]interface{})

		if option.defaultVersion.name != ""{
			defaultVersiondictmap["name"] =  option.defaultVersion.name
		}

		if option.defaultVersion.description != ""{
			createMLModeljsonmap["description"] =  option.defaultVersion.description
		}

		if (option.defaultVersion.deploymentUri) != "" {
			createMLModeljsonmap["deploymentUri"] =  option.defaultVersion.deploymentUri
		}

		if (option.defaultVersion.isDefault) != false {
			defaultVersiondictmap["isDefault"] =  option.defaultVersion.isDefault
		}

		if (option.defaultVersion.createTime) != "" {
			defaultVersiondictmap["createTime"] =  option.defaultVersion.createTime
		}

		if (option.defaultVersion.lastUseTime) != "" {
				defaultVersiondictmap["lastUseTime"] =  option.defaultVersion.lastUseTime
		}

		if (option.defaultVersion.runtimeVersion) != "" {
				defaultVersiondictmap["runtimeVersion"] =  option.defaultVersion.runtimeVersion
		}


		if (option.defaultVersion.state) != "" {
				defaultVersiondictmap["state"] =  option.defaultVersion.state
		}

		if (option.defaultVersion.errorMessage) != "" {
				defaultVersiondictmap["errorMessage"] =  option.defaultVersion.errorMessage
		}

		if (option.defaultVersion.framework) != "" {
				defaultVersiondictmap["framework"] =  option.defaultVersion.framework
		}

		if (option.defaultVersion.pythonVersion) != "" {
				defaultVersiondictmap["pythonVersion"] =  option.defaultVersion.pythonVersion
		}

		autoScalingdictmap := make(map[string]interface{})

		if (option.defaultVersion.autoScaling.minNodes) != "" {
				autoScalingdictmap["minNodes"] =  option.defaultVersion.autoScaling.minNodes
				defaultVersiondictmap["autoScaling"] = autoScalingdictmap
		}


		manualScalingdictmap := make(map[string]interface{})

		if (option.defaultVersion.manualScaling.nodes) != "" {
				manualScalingdictmap["nodes"] =  option.defaultVersion.manualScaling.nodes
				defaultVersiondictmap["autoScaling"] = manualScalingdictmap
		}

		if defaultVersiondictmap != nil {
			createMLModeljsonmap["defaultVersion"]	= defaultVersiondictmap
		}

}
