package awscontainer

import (
	"bytes"
	"encoding/json"
	awsauth "github.com/shlokgilda/gocloud/awsauth"
	"io/ioutil"
	"net/http"
)

//preparestoptaskparamsdict create dictnoary for stoptask.
func preparestoptaskparamsdict(stoptaskjsonmap map[string]interface{}, stoptask Stoptask) {
	if stoptask.Cluster != "" {
		stoptaskjsonmap["cluster"] = stoptask.Cluster
	}

	if stoptask.Reason != "" {
		stoptaskjsonmap["reason"] = stoptask.Reason
	}

	if stoptask.Task != "" {
		stoptaskjsonmap["task"] = stoptask.Task
	}
}

//preparestoptaskparams create dictnoary for stoptask.
func preparestoptaskparams(params map[string]string, stoptask Stoptask, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.StopTask"
}

//preparedeleteserviceparams Delete dictnoary for deleteservice.
func preparedeleteserviceparams(params map[string]string, deleteservice Deleteservice, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.DeleteService"
}

//preparedeleteserviceparamsdict Delete dictnoary for deleteservice.
func preparedeleteserviceparamsdict(deleteServicejsonmap map[string]interface{}, deleteservice Deleteservice) {
	if deleteservice.Cluster != "" {
		deleteServicejsonmap["cluster"] = deleteservice.Cluster
	}
	if deleteservice.Service != "" {
		deleteServicejsonmap["service"] = deleteservice.Service
	}
}

//prepareDeletecluster Delete dictnoary for Deletecluster.
func prepareDeletecluster(params map[string]string, clusterName string, Region string) {
	if clusterName != "" {
		params["clusterName"] = clusterName
	}
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.DeleteCluster"
}

//prepareCreateclusterrparams creates dictnoary for Createcluster.
func prepareCreateclusterrparams(params map[string]string, clusterName string, Region string) {
	if clusterName != "" {
		params["clusterName"] = clusterName
	}
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.CreateCluster"
}

//preparecreateServiceplacementConstraintsparams creates dictnoary for createService.
func preparecreateServiceplacementConstraintsparams(Createservicejsonmap map[string]interface{}, createservice Createservice) {
	if len(createservice.PlacementConstraints) != 0 {
		placementConstraints := make([]map[string]interface{}, 0)
		for i := 0; i < len(createservice.PlacementConstraints); i++ {
			PlacementConstraint := make(map[string]interface{})

			if createservice.PlacementConstraints[i].Expression != "" {
				PlacementConstraint["expression"] = createservice.PlacementConstraints[i].Expression
			}

			if createservice.PlacementConstraints[i].Type != "" {
				PlacementConstraint["type"] = createservice.PlacementConstraints[i].Type
			}

			placementConstraints = append(placementConstraints, PlacementConstraint)
		}

		Createservicejsonmap["placementConstraints"] = placementConstraints
	}
}

//preparecreateServiceloadBalancersparams  creates dictnoary for loadBalancer.
func preparecreateServiceloadBalancersparams(Createservicejsonmap map[string]interface{}, createservice Createservice) {
	if len(createservice.LoadBalancers) != 0 {
		loadBalancers := make([]map[string]interface{}, 0)
		for i := 0; i < len(createservice.LoadBalancers); i++ {
			loadBalancer := make(map[string]interface{})

			if createservice.LoadBalancers[i].ContainerName != "" {
				loadBalancer["containerName"] = createservice.LoadBalancers[i].ContainerName
			}

			if createservice.LoadBalancers[i].LoadBalancerName != "" {
				loadBalancer["loadBalancerName"] = createservice.LoadBalancers[i].LoadBalancerName
			}

			if createservice.LoadBalancers[i].TargetGroupArn != "" {
				loadBalancer["targetGroupArn"] = createservice.LoadBalancers[i].TargetGroupArn
			}

			if createservice.LoadBalancers[i].ContainerPort != 0 {
				loadBalancer["containerPort"] = createservice.LoadBalancers[i].ContainerPort
			}

			loadBalancers = append(loadBalancers, loadBalancer)
		}

		Createservicejsonmap["loadBalancers"] = loadBalancers

	}
}

//preparecreateServicedeploymentConfigurationparams  creates dictnoary for createService.
func preparecreateServicedeploymentConfigurationparams(Createservicejsonmap map[string]interface{}, createservice Createservice) {
	if (createservice.DeploymentConfigurations != DeploymentConfiguration{}) {
		deploymentConfiguration := make(map[string]interface{})
		if createservice.DeploymentConfigurations.MaximumPercent != 0 {
			deploymentConfiguration["maximumPercent"] = createservice.DeploymentConfigurations.MaximumPercent
		}
		if createservice.DeploymentConfigurations.MinimumHealthyPercent != 0 {
			deploymentConfiguration["minimumHealthyPercent"] = createservice.DeploymentConfigurations.MinimumHealthyPercent
		}
		Createservicejsonmap["deploymentConfiguration"] = deploymentConfiguration
	}
}

//preparestarttaskparams  creates dictnoary for starttask.
func preparestarttaskparams(params map[string]string, starttask Starttask, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.StartTask"
}

//preparestarttaskparamsdict creates dictnoary for starttask.
func preparestarttaskparamsdict(starttaskjsonmap map[string]interface{}, starttask Starttask) {
	if starttask.Cluster != "" {
		starttaskjsonmap["cluster"] = starttask.Cluster
	}
	if starttask.TaskDefinition != "" {
		starttaskjsonmap["taskDefinition"] = starttask.TaskDefinition
	}
	if len(starttask.ContainerInstances) != 0 {
		starttaskjsonmap["containerInstances"] = starttask.ContainerInstances
	}

	if starttask.Group != "" {
		starttaskjsonmap["group"] = starttask.Group
	}
	if starttask.StartedBy != "" {
		starttaskjsonmap["startedBy"] = starttask.StartedBy
	}

	preparestarttaskoverridesparams(starttaskjsonmap, starttask)
}

//preparestarttaskoverridesparams creates dictnoary for starttask.
func preparestarttaskoverridesparams(starttaskjsonmap map[string]interface{}, starttask Starttask) {
	overrides := make(map[string]interface{})
	if starttask.overrides.TaskRoleArn != "" {
		overrides["taskRoleArn"] = starttask.overrides.TaskRoleArn
	}
	if len(starttask.overrides.ContainerOverrides) != 0 {
		containerOverrides := make([]map[string]interface{}, 0)
		for i := 0; i < len(starttask.overrides.ContainerOverrides); i++ {
			containerOverride := make(map[string]interface{})
			if starttask.overrides.ContainerOverrides[i].Name != "" {
				containerOverride["name"] = starttask.overrides.ContainerOverrides[i].Name
			}
			if starttask.overrides.ContainerOverrides[i].MemoryReservation != "" {
				containerOverride["memoryReservation"] = starttask.overrides.ContainerOverrides[i].MemoryReservation
			}
			if starttask.overrides.ContainerOverrides[i].Memory != 0 {
				containerOverride["memory"] = starttask.overrides.ContainerOverrides[i].Memory
			}
			if starttask.overrides.ContainerOverrides[i].Cpu != 0 {
				containerOverride["cpu"] = starttask.overrides.ContainerOverrides[i].Cpu
			}
			if len(starttask.overrides.ContainerOverrides[i].Command) != 0 {
				containerOverride["command"] = starttask.overrides.ContainerOverrides[i].Command
			}

			if len(starttask.overrides.ContainerOverrides[i].Environments) != 0 {
				containerOverride["environment"] = starttask.overrides.ContainerOverrides[i].Environments
			}
			containerOverrides = append(containerOverrides, containerOverride)
		}
		overrides["containerOverrides"] = containerOverrides
	}
	if len(overrides) != 0 {
		starttaskjsonmap["overrides"] = overrides
	}
}

//prepareruntaskparamsdict creates dictnoary for runtask.
func prepareruntaskparamsdict(runtaskjsonmap map[string]interface{}, runtask Runtask) {
	if runtask.Cluster != "" {
		runtaskjsonmap["cluster"] = runtask.Cluster
	}
	if runtask.TaskDefinition != "" {
		runtaskjsonmap["taskDefinition"] = runtask.TaskDefinition
	}
	if runtask.Count != 0 {
		runtaskjsonmap["count"] = runtask.Count
	}

	if runtask.Group != "" {
		runtaskjsonmap["group"] = runtask.Group
	}
	if runtask.StartedBy != "" {
		runtaskjsonmap["startedBy"] = runtask.StartedBy
	}

	prepareruntaskoverridesparams(runtaskjsonmap, runtask)
	prepareruntaskplacementConstraintsparams(runtaskjsonmap, runtask)
	prepareruntaskplacementStrategyparams(runtaskjsonmap, runtask)
}

//prepareruntaskoverridesparams creates dictnoary for runtask.
func prepareruntaskoverridesparams(runtaskjsonmap map[string]interface{}, runtask Runtask) {
	overrides := make(map[string]interface{})
	if runtask.overrides.TaskRoleArn != "" {
		overrides["taskRoleArn"] = runtask.overrides.TaskRoleArn
	}
	if len(runtask.overrides.ContainerOverrides) != 0 {
		containerOverrides := make([]map[string]interface{}, 0)
		for i := 0; i < len(runtask.overrides.ContainerOverrides); i++ {
			containerOverride := make(map[string]interface{})
			if runtask.overrides.ContainerOverrides[i].Name != "" {
				containerOverride["name"] = runtask.overrides.ContainerOverrides[i].Name
			}
			if runtask.overrides.ContainerOverrides[i].MemoryReservation != "" {
				containerOverride["memoryReservation"] = runtask.overrides.ContainerOverrides[i].MemoryReservation
			}
			if runtask.overrides.ContainerOverrides[i].Memory != 0 {
				containerOverride["memory"] = runtask.overrides.ContainerOverrides[i].Memory
			}
			if runtask.overrides.ContainerOverrides[i].Cpu != 0 {
				containerOverride["cpu"] = runtask.overrides.ContainerOverrides[i].Cpu
			}
			if len(runtask.overrides.ContainerOverrides[i].Command) != 0 {
				containerOverride["command"] = runtask.overrides.ContainerOverrides[i].Command
			}

			if len(runtask.overrides.ContainerOverrides[i].Environments) != 0 {
				containerOverride["environment"] = runtask.overrides.ContainerOverrides[i].Environments
			}
			containerOverrides = append(containerOverrides, containerOverride)
		}
		overrides["containerOverrides"] = containerOverrides
	}
	if len(overrides) != 0 {
		runtaskjsonmap["overrides"] = overrides
	}
}

//prepareruntaskparams creates dictnoary for runtask.
func prepareruntaskparams(params map[string]string, runtask Runtask, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.RunTask"
}

//prepareruntaskplacementStrategyparams creates dictnoary for runtask.
func prepareruntaskplacementStrategyparams(Createservicejsonmap map[string]interface{}, runtask Runtask) {
	if len(runtask.PlacementStrategys) != 0 {
		placementstrategys := make([]map[string]interface{}, 0)
		for i := 0; i < len(runtask.PlacementStrategys); i++ {
			placementstrategy := make(map[string]interface{})

			if runtask.PlacementStrategys[i].Field != "" {
				placementstrategy["field"] = runtask.PlacementStrategys[i].Field
			}

			if runtask.PlacementStrategys[i].Type != "" {
				placementstrategy["type"] = runtask.PlacementStrategys[i].Type
			}

			placementstrategys = append(placementstrategys, placementstrategy)
		}

		Createservicejsonmap["placementstrategy"] = placementstrategys
	}
}

//prepareruntaskplacementConstraintsparams creates dictnoary for runtask.
func prepareruntaskplacementConstraintsparams(runtaskjsonmap map[string]interface{}, runtask Runtask) {
	if len(runtask.PlacementConstraints) != 0 {
		placementConstraints := make([]map[string]interface{}, 0)
		for i := 0; i < len(runtask.PlacementConstraints); i++ {
			PlacementConstraint := make(map[string]interface{})

			if runtask.PlacementConstraints[i].Expression != "" {
				PlacementConstraint["expression"] = runtask.PlacementConstraints[i].Expression
			}

			if runtask.PlacementConstraints[i].Type != "" {
				PlacementConstraint["type"] = runtask.PlacementConstraints[i].Type
			}

			placementConstraints = append(placementConstraints, PlacementConstraint)
		}

		runtaskjsonmap["placementConstraints"] = placementConstraints
	}
}

//preparecreateServiceparams creates dictnoary for createService.
func preparecreateServiceparams(params map[string]string, createservice Createservice, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.CreateService"
}

//preparecreateServiceparams creates dictnoary for createService.
func preparecreateServiceparamsdict(Createservicejsonmap map[string]interface{}, createservice Createservice) {
	if createservice.ServiceName != "" {
		Createservicejsonmap["serviceName"] = createservice.ServiceName
	}
	if createservice.TaskDefinition != "" {
		Createservicejsonmap["taskDefinition"] = createservice.TaskDefinition
	}
	if createservice.DesiredCount != 0 {
		Createservicejsonmap["desiredCount"] = createservice.DesiredCount
	}

	if createservice.ClientToken != "" {
		Createservicejsonmap["clientToken"] = createservice.ClientToken
	}
	if createservice.Cluster != "" {
		Createservicejsonmap["cluster"] = createservice.Cluster
	}
	if createservice.Role != "" {
		Createservicejsonmap["role"] = createservice.Role
	}
	preparecreateServicedeploymentConfigurationparams(Createservicejsonmap, createservice)
	preparecreateServiceloadBalancersparams(Createservicejsonmap, createservice)

	preparecreateServiceplacementConstraintsparams(Createservicejsonmap, createservice)
	preparecreateServiceplacementStrategyparams(Createservicejsonmap, createservice)

}

//preparecreateServiceplacementStrategyparams creates dictnoary for createService.
func preparecreateServiceplacementStrategyparams(Createservicejsonmap map[string]interface{}, createservice Createservice) {
	if len(createservice.PlacementStrategys) != 0 {
		placementstrategys := make([]map[string]interface{}, 0)
		for i := 0; i < len(createservice.PlacementStrategys); i++ {
			placementstrategy := make(map[string]interface{})

			if createservice.PlacementStrategys[i].Field != "" {
				placementstrategy["field"] = createservice.PlacementStrategys[i].Field
			}

			if createservice.PlacementStrategys[i].Type != "" {
				placementstrategy["type"] = createservice.PlacementStrategys[i].Type
			}

			placementstrategys = append(placementstrategys, placementstrategy)
		}

		Createservicejsonmap["placementstrategy"] = placementstrategys
	}
}

//PrepareSignatureV4query creates PrepareSignatureV4 for request.
func (ecscontainer *Ecscontainer) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}, response map[string]interface{}) error {
	ECSEndpoint := "https://ecs." + params["Region"] + ".amazonaws.com"
	service := "ecs"
	method := "POST"
	host := service + "." + params["Region"] + ".amazonaws.com"
	ContentType := "application/x-amz-json-1.1"
	signedheaders := "content-type;host;x-amz-date;x-amz-target"
	amztarget := params["amztarget"]

	requestparametersjson, _ := json.Marshal(paramsmap)
	requestparametersjsonstring := string(requestparametersjson)
	requestparametersjsonstringbyte := []byte(requestparametersjsonstring)
	client := new(http.Client)
	request, _ := http.NewRequest("POST", ECSEndpoint, bytes.NewBuffer(requestparametersjsonstringbyte))
	request = awsauth.SignatureV4(request, requestparametersjsonstringbyte, amztarget, method, params["Region"], service, host, ContentType, signedheaders)
	resp, err := client.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	response["body"] = string(body)
	response["status"] = resp.StatusCode
	return err
}
