package awscontainer

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}

func TestCreateCluster(t *testing.T) {
	var ecscontainer Ecscontainer

	createcluster := map[string]interface{}{
		"clusterName": "gocloud-test",
		"Region":      "us-east-1",
	}

	ecscontainer.CreateCluster(createcluster)
}

func TestDeleteCluster(t *testing.T) {
	var ecscontainer Ecscontainer

	deletecluster := map[string]interface{}{
		"clusterName": "gocloud-test",
		"Region":      "us-east-1",
	}

	_, err := ecscontainer.DeleteCluster(deletecluster)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestStopTask(t *testing.T) {
	var ecscontainer Ecscontainer
	stoptask := map[string]interface{}{
		"cluster": "cluster",
		"reason":  "reason",
		"task":    "task",
		"Region":  "us-east-1",
	}
	_, err := ecscontainer.StopTask(stoptask)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestCreateService(t *testing.T) {
	var ecscontainer Ecscontainer
	LoadBalancers := []map[string]interface{}{{
		"containerName":    "rootmonk",
		"loadBalancerName": "us-east-2",
	}, {
		"containerName":    "rootmonk",
		"loadBalancerName": "us-east-2",
	},
	}
	createservice := map[string]interface{}{
		"serviceName":    "ecs-simple-service",
		"taskDefinition": "ecs-demo",
		"desiredCount":   1,
		"Region":         "us-east-1",
		"LoadBalancers":  LoadBalancers,
	}

	_, err := ecscontainer.CreateService(createservice)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteService(t *testing.T) {
	var ecscontainer Ecscontainer

	deleteservice := map[string]interface{}{
		"cluster": "cluster",
		"service": "service",
		"Region":  "us-east-1",
	}
	_, err := ecscontainer.DeleteService(deleteservice)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
