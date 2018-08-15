package clouddataflow

func createstreamdictnoaryconvert(option Createstream, createstreamjsonmap map[string]interface{}) {

	if option.ID != "" {
		createstreamjsonmap["id"] = option.ID
	}

	if option.ProjectID != "" {
		createstreamjsonmap["projectId"] = option.ProjectID
	}

	if option.Name != "" {
		createstreamjsonmap["name"] = option.Name
	}

	if option.Type != "" {
		createstreamjsonmap["type"] = option.Type
	}

	if option.CurrentState != "" {
		createstreamjsonmap["currentState"] = option.CurrentState
	}

	if option.CurrentStateTime != "" {
		createstreamjsonmap["currentStateTime"] = option.CurrentStateTime
	}

	if option.RequestedState != "" {
		createstreamjsonmap["requestedState"] = option.RequestedState
	}

	if option.CreateTime != "" {
		createstreamjsonmap["createTime"] = option.CreateTime
	}

	if option.ReplaceJobID != "" {
		createstreamjsonmap["replaceJobId"] = option.ReplaceJobID
	}

	if option.Location != "" {
		createstreamjsonmap["location"] = option.Location
	}

	prepareStageStates(option, createstreamjsonmap)
	prepareEnvironment(option, createstreamjsonmap)
}

func prepareEnvironment(option Createstream, createstreamjsonmap map[string]interface{}) {

	environmentv := make(map[string]interface{})

	versionv := make(map[string]interface{})

	if option.environment.version.Major != "" {
		versionv["major"] = option.environment.version.Major
	}

	if option.environment.version.JobType != "" {
		versionv["job_type"] = option.environment.version.JobType
	}

	environmentv["version"] = versionv

	userAgentv := make(map[string]interface{})

	if option.environment.userAgent.Name != "" {
		userAgentv["name"] = option.environment.userAgent.Name
	}

	if option.environment.userAgent.BuildDate != "" {
		userAgentv["build.date"] = option.environment.userAgent.BuildDate
	}

	if option.environment.userAgent.Version != "" {
		userAgentv["version"] = option.environment.userAgent.Version
	}

	supportv := make(map[string]interface{})

	if option.environment.userAgent.support.Status != "" {
		supportv["status"] = option.environment.userAgent.support.Status
	}

	if option.environment.userAgent.support.URL != "" {
		supportv["url"] = option.environment.userAgent.support.URL
	}

	userAgentv["support"] = supportv

	environmentv["userAgent"] = userAgentv

	createstreamjsonmap["environment"] = environmentv
}

func prepareStageStates(option Createstream, createstreamjsonmap map[string]interface{}) {

	if len(option.stageStates) > 0 {

		stageStatesarray := make([]map[string]interface{}, 0)

		for i := 0; i < len(option.stageStates); i++ {

			stageState := make(map[string]interface{})
			stageState["currentStateTime"] = option.stageStates[i].CurrentStateTime
			stageState["executionStageName"] = option.stageStates[i].ExecutionStageName
			stageState["executionStageName"] = option.stageStates[i].ExecutionStageName

			stageStatesarray = append(stageStatesarray, stageState)

		}

		createstreamjsonmap["stageStates"] = stageStatesarray
	}
}
