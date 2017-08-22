package awscontainer

//Stoptask stops container.
func (ecscontainer *Ecscontainer) Stoptask(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var Region string
	var stoptask Stoptask
	for key, value := range param {
		switch key {
		case "cluster":
			clusterV, _ := value.(string)
			stoptask.Cluster = clusterV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "reason":
			ReasonV, _ := value.(string)
			stoptask.Reason = ReasonV

		case "task":
			taskV, _ := value.(string)
			stoptask.Task = taskV
		}
	}
	params := make(map[string]string)
	preparestoptaskparams(params, stoptask, Region)
	stoptaskjsonmap := make(map[string]interface{})
	preparestoptaskparamsdict(stoptaskjsonmap, stoptask)
	response := make(map[string]interface{})
	err = ecscontainer.PrepareSignatureV4query(params, stoptaskjsonmap, response)
	resp = response
	return resp, err
}

//Deleteservice Delete container service.
func (ecscontainer *Ecscontainer) Deleteservice(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var Region string
	var deleteservice Deleteservice
	for key, value := range param {
		switch key {
		case "cluster":
			clusterV, _ := value.(string)
			deleteservice.Cluster = clusterV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "service":
			serviceV, _ := value.(string)
			deleteservice.Service = serviceV
		}
	}
	params := make(map[string]string)
	preparedeleteserviceparams(params, deleteservice, Region)
	deleteServicejsonmap := make(map[string]interface{})
	preparedeleteserviceparamsdict(deleteServicejsonmap, deleteservice)
	response := make(map[string]interface{})
	err = ecscontainer.PrepareSignatureV4query(params, deleteServicejsonmap, response)
	resp = response
	return resp, err
}

//Starttask start container service.
func (ecscontainer *Ecscontainer) Starttask(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var starttask Starttask
	var Region string
	for key, value := range param {
		switch key {
		case "cluster":
			clusterV, _ := value.(string)
			starttask.Cluster = clusterV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "containerInstances":
			ContainerInstancesV, _ := value.([]string)
			starttask.ContainerInstances = ContainerInstancesV

		case "group":
			GroupV, _ := value.(string)
			starttask.Group = GroupV

		case "startedBy":
			StartedByV, _ := value.(string)
			starttask.StartedBy = StartedByV

		case "taskDefinition":
			TaskDefinitionV, _ := value.(string)
			starttask.TaskDefinition = TaskDefinitionV

		case "overrides":
			overridesparam, _ := value.(map[string]interface{})
			for overridesparamkey, overridesparamvalue := range overridesparam {
				switch overridesparamkey {
				case "taskRoleArn":
					starttask.overrides.TaskRoleArn = overridesparamvalue.(string)
				case "containerOverrides":
					containerOverridesparam, _ := overridesparamvalue.([]map[string]interface{})
					for i := 0; i < len(containerOverridesparam); i++ {
						var containerOverride ContainerOverride
						for containerOverrideparamkey, containerOverrideparamvalue := range containerOverridesparam[i] {
							switch containerOverrideparamkey {
							case "name":
								containerOverride.Name = containerOverrideparamvalue.(string)

							case "memoryReservation":
								containerOverride.MemoryReservation = containerOverrideparamvalue.(string)

							case "memory":
								containerOverride.Memory = containerOverrideparamvalue.(int)

							case "cpu":
								containerOverride.Cpu = containerOverrideparamvalue.(int)

							case "command":
								containerOverride.Command = containerOverrideparamvalue.([]string)

							case "environment":
								Environmentsparam := containerOverrideparamvalue.([]map[string]string)
								for i := 0; i < len(Environmentsparam); i++ {
									var environment Environment
									for environmentsparamkey, environmentsparamvalue := range Environmentsparam[i] {
										switch environmentsparamkey {
										case "name":
											environment.Name = environmentsparamvalue
										case "value":
											environment.Value = environmentsparamvalue
										}
									}
									containerOverride.Environments = append(containerOverride.Environments, environment)
								}
							}
						}
						starttask.overrides.ContainerOverrides = append(starttask.overrides.ContainerOverrides, containerOverride)
					}
				}
			}
		}
	}
	params := make(map[string]string)
	preparestarttaskparams(params, starttask, Region)
	starttaskjsonmap := make(map[string]interface{})
	preparestarttaskparamsdict(starttaskjsonmap, starttask)

	response := make(map[string]interface{})
	err = ecscontainer.PrepareSignatureV4query(params, starttaskjsonmap, response)
	resp = response
	return resp, err
}

//Runtask runs container.
func (ecscontainer *Ecscontainer) Runtask(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var runtask Runtask
	var Region string
	for key, value := range param {
		switch key {
		case "cluster":
			clusterV, _ := value.(string)
			runtask.Cluster = clusterV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "count":
			CountV, _ := value.(int)
			runtask.Count = CountV

		case "group":
			GroupV, _ := value.(string)
			runtask.Group = GroupV

		case "startedBy":
			StartedByV, _ := value.(string)
			runtask.StartedBy = StartedByV

		case "taskDefinition":
			TaskDefinitionV, _ := value.(string)
			runtask.TaskDefinition = TaskDefinitionV

		case "placementConstraints":
			placementconstraintsparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(placementconstraintsparam); i++ {
				var placementconstraint Placementconstraint
				for placementConstraintsparamkey, placementConstraintsparamvalue := range placementconstraintsparam[i] {
					switch placementConstraintsparamkey {
					case "Expression":
						placementconstraint.Expression = placementConstraintsparamvalue.(string)
					case "Type":
						placementconstraint.Type = placementConstraintsparamvalue.(string)
					}
				}
				runtask.PlacementConstraints = append(runtask.PlacementConstraints, placementconstraint)
			}

		case "placementStrategy":
			placementstrategyparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(placementstrategyparam); i++ {
				var placementstrategy Placementstrategy
				for placementstrategyparamkey, placementstrategyparamvalue := range placementstrategyparam[i] {
					switch placementstrategyparamkey {
					case "field":
						placementstrategy.Field = placementstrategyparamvalue.(string)
					case "Type":
						placementstrategy.Type = placementstrategyparamvalue.(string)
					}
				}
				runtask.PlacementStrategys = append(runtask.PlacementStrategys, placementstrategy)
			}

		case "overrides":
			overridesparam, _ := value.(map[string]interface{})
			for overridesparamkey, overridesparamvalue := range overridesparam {
				switch overridesparamkey {
				case "taskRoleArn":
					runtask.overrides.TaskRoleArn = overridesparamvalue.(string)
				case "containerOverrides":
					containerOverridesparam, _ := overridesparamvalue.([]map[string]interface{})
					for i := 0; i < len(containerOverridesparam); i++ {
						var containerOverride ContainerOverride
						for containerOverrideparamkey, containerOverrideparamvalue := range containerOverridesparam[i] {
							switch containerOverrideparamkey {
							case "name":
								containerOverride.Name = containerOverrideparamvalue.(string)

							case "memoryReservation":
								containerOverride.MemoryReservation = containerOverrideparamvalue.(string)

							case "memory":
								containerOverride.Memory = containerOverrideparamvalue.(int)

							case "cpu":
								containerOverride.Cpu = containerOverrideparamvalue.(int)

							case "command":
								containerOverride.Command = containerOverrideparamvalue.([]string)

							case "environment":
								Environmentsparam := containerOverrideparamvalue.([]map[string]string)
								for i := 0; i < len(Environmentsparam); i++ {
									var environment Environment
									for environmentsparamkey, environmentsparamvalue := range Environmentsparam[i] {
										switch environmentsparamkey {
										case "name":
											environment.Name = environmentsparamvalue
										case "value":
											environment.Value = environmentsparamvalue
										}
									}
									containerOverride.Environments = append(containerOverride.Environments, environment)
								}
							}
						}
						runtask.overrides.ContainerOverrides = append(runtask.overrides.ContainerOverrides, containerOverride)
					}
				}
			}
		}
	}
	params := make(map[string]string)
	prepareruntaskparams(params, runtask, Region)
	runtaskjsonmap := make(map[string]interface{})
	prepareruntaskparamsdict(runtaskjsonmap, runtask)
	response := make(map[string]interface{})
	err = ecscontainer.PrepareSignatureV4query(params, runtaskjsonmap, response)
	resp = response
	return resp, err
}

//Createservice creates container service.
func (ecscontainer *Ecscontainer) Createservice(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var createservice Createservice
	var Region string
	for key, value := range param {
		switch key {
		case "serviceName":
			serviceNameV, _ := value.(string)
			createservice.ServiceName = serviceNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "taskDefinition":
			TaskDefinitionV, _ := value.(string)
			createservice.TaskDefinition = TaskDefinitionV

		case "desiredCount":
			desiredCountV, _ := value.(int)
			createservice.DesiredCount = desiredCountV

		case "clientToken":
			clientTokenV, _ := value.(string)
			createservice.ClientToken = clientTokenV

		case "cluster":
			clusterV, _ := value.(string)
			createservice.Cluster = clusterV

		case "role":
			roleV, _ := value.(string)
			createservice.Role = roleV

		case "deploymentConfiguration":
			deploymentConfigurationV, _ := value.(map[string]int)
			createservice.DeploymentConfigurations.MaximumPercent = deploymentConfigurationV["maximumPercent"]
			createservice.DeploymentConfigurations.MinimumHealthyPercent = deploymentConfigurationV["minimumHealthyPercent"]

		case "LoadBalancers":
			LoadBalancersparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(LoadBalancersparam); i++ {
				var loadBalancer LoadBalancer
				for loadBalancersparamkey, loadBalancersparamparamvalue := range LoadBalancersparam[i] {
					switch loadBalancersparamkey {
					case "containerName":
						loadBalancer.ContainerName = loadBalancersparamparamvalue.(string)
					case "containerPort":
						loadBalancer.ContainerPort = loadBalancersparamparamvalue.(int)
					case "loadBalancerName":
						loadBalancer.LoadBalancerName = loadBalancersparamparamvalue.(string)
					case "targetGroupArn":
						loadBalancer.TargetGroupArn = loadBalancersparamparamvalue.(string)
					}
				}
				createservice.LoadBalancers = append(createservice.LoadBalancers, loadBalancer)
			}

		case "placementConstraints":
			placementconstraintsparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(placementconstraintsparam); i++ {
				var placementconstraint Placementconstraint
				for placementConstraintsparamkey, placementConstraintsparamvalue := range placementconstraintsparam[i] {
					switch placementConstraintsparamkey {
					case "Expression":
						placementconstraint.Expression = placementConstraintsparamvalue.(string)
					case "Type":
						placementconstraint.Type = placementConstraintsparamvalue.(string)
					}
				}
				createservice.PlacementConstraints = append(createservice.PlacementConstraints, placementconstraint)
			}

		case "placementStrategy":
			placementstrategyparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(placementstrategyparam); i++ {
				var placementstrategy Placementstrategy
				for placementstrategyparamkey, placementstrategyparamvalue := range placementstrategyparam[i] {
					switch placementstrategyparamkey {
					case "field":
						placementstrategy.Field = placementstrategyparamvalue.(string)
					case "Type":
						placementstrategy.Type = placementstrategyparamvalue.(string)
					}
				}
				createservice.PlacementStrategys = append(createservice.PlacementStrategys, placementstrategy)
			}

		}
	}
	params := make(map[string]string)
	preparecreateServiceparams(params, createservice, Region)
	Createservicejsonmap := make(map[string]interface{})
	preparecreateServiceparamsdict(Createservicejsonmap, createservice)
	response := make(map[string]interface{})
	err = ecscontainer.PrepareSignatureV4query(params, Createservicejsonmap, response)
	resp = response
	return resp, err
}

//Createcluster creates cluster.
func (ecscontainer *Ecscontainer) Createcluster(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var clusterName, Region string
	for key, value := range param {
		switch key {
		case "clusterName":
			clusterNameV, _ := value.(string)
			clusterName = clusterNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}
	params := make(map[string]string)
	prepareCreateclusterrparams(params, clusterName, Region)

	Createclusterjsonmap := map[string]interface{}{
		"clusterName": params["clusterName"],
	}
	response := make(map[string]interface{})
	err = ecscontainer.PrepareSignatureV4query(params, Createclusterjsonmap, response)
	resp = response
	return resp, err
}

//Deletecluster delete cluster.
func (ecscontainer *Ecscontainer) Deletecluster(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var clusterName, Region string
	for key, value := range param {
		switch key {
		case "clusterName":
			clusterNameV, _ := value.(string)
			clusterName = clusterNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}
	params := make(map[string]string)
	prepareDeletecluster(params, clusterName, Region)

	Deletecontainerjsonmap := map[string]interface{}{
		"cluster": clusterName,
	}

	response := make(map[string]interface{})
	err = ecscontainer.PrepareSignatureV4query(params, Deletecontainerjsonmap, response)
	resp = response
	return resp, err
}
