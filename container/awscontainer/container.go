package awscontainer

import (
	"bytes"
	"encoding/json"
	"fmt"
	awsauth "github.com/scorelab/gocloud-v2/awsauth"
	"io/ioutil"
	"net/http"
)

type Ecscontainer struct {
}

type Createservice struct {
	ServiceName              string
	TaskDefinition           string
	DesiredCount             int
	ClientToken              string
	Cluster                  string
	Role                     string
	DeploymentConfigurations DeploymentConfiguration
	LoadBalancers            []LoadBalancer
	PlacementConstraints     []Placementconstraint
	PlacementStrategys       []Placementstrategy
}

type Placementconstraint struct {
	Expression string
	Type       string
}

type Placementstrategy struct {
	Field string
	Type  string
}

type LoadBalancer struct {
	ContainerName    string
	ContainerPort    int
	LoadBalancerName string
	TargetGroupArn   string
}

type DeploymentConfiguration struct {
	MaximumPercent        int
	MinimumHealthyPercent int
}

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
	ecscontainer.PrepareSignatureV4query(params, Createservicejsonmap)
	return
}

func preparecreateServiceparams(params map[string]string, createservice Createservice, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.CreateService"
}

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

func preparecreateServiceloadBalancersparams(Createservicejsonmap map[string]interface{}, createservice Createservice) {
	fmt.Println("len of createservice.LoadBalancers", len(createservice.LoadBalancers))
	if len(createservice.LoadBalancers) != 0 {
		loadBalancers := make([]map[string]interface{}, 0)
		fmt.Println("loadBalancers", loadBalancers)
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

		fmt.Println("Createservicejsonmap of loadBalancers", Createservicejsonmap["loadBalancers"])
	}
}

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

func (ecscontainer *Ecscontainer) Createcontainer(request interface{}) (resp interface{}, err error) {

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
	prepareCreatcontainerparams(params, clusterName, Region)

	Creatcontainerjsonmap := map[string]interface{}{
		"clusterName": params["clusterName"],
	}
	ecscontainer.PrepareSignatureV4query(params, Creatcontainerjsonmap)
	return
}

func (ecscontainer *Ecscontainer) Deletecontainer(request interface{}) (resp interface{}, err error) {

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
	prepareDeletecontainer(params, clusterName, Region)

	Creatcontainerjsonmap := map[string]interface{}{
		"clusterName": params["clusterName"],
	}
	ecscontainer.PrepareSignatureV4query(params, Creatcontainerjsonmap)
	return
}

func prepareDeletecontainer(params map[string]string, clusterName string, Region string) {
	if clusterName != "" {
		params["clusterName"] = clusterName
	}
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.DeleteCluster"
}

func prepareCreatcontainerparams(params map[string]string, clusterName string, Region string) {
	if clusterName != "" {
		params["clusterName"] = clusterName
	}
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.CreateCluster"
}

func (ecscontainer *Ecscontainer) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}) {
	fmt.Println(paramsmap)
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
	fmt.Println("requestparametersjsonstring", requestparametersjsonstring)
	client := new(http.Client)
	request, _ := http.NewRequest("POST", ECSEndpoint, bytes.NewBuffer(requestparametersjsonstringbyte))
	request = awsauth.SignatureV4(request, requestparametersjsonstringbyte, amztarget, method, params["Region"], service, host, ContentType, signedheaders)
	resp, _ := client.Do(request)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("resp Body:", string(body))
}
