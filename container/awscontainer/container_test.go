package awscontainer

import "testing"
import awsAuth "github.com/shlokgilda/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}

func TestCreatecluster(t *testing.T) {
	var ecscontainer Ecscontainer

	createcluster := map[string]interface{}{
		"clusterName": "gocloud-test",
		"Region":      "us-east-1",
	}

	ecscontainer.Createcluster(createcluster)
}

func TestDeletecluster(t *testing.T) {
	var ecscontainer Ecscontainer

	deletecluster := map[string]interface{}{
		"clusterName": "gocloud-test",
		"Region":      "us-east-1",
	}

	_, err := ecscontainer.Deletecluster(deletecluster)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestStoptask(t *testing.T) {
	var ecscontainer Ecscontainer
	stoptask := map[string]interface{}{
		"cluster": "cluster",
		"reason":  "reason",
		"task":    "task",
		"Region":  "us-east-1",
	}
	_, err := ecscontainer.Stoptask(stoptask)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestCreateservice(t *testing.T) {
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

	_, err := ecscontainer.Createservice(createservice)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteservice(t *testing.T) {
	var ecscontainer Ecscontainer

	deleteservice := map[string]interface{}{
		"cluster": "cluster",
		"service": "service",
		"Region":  "us-east-1",
	}
	_, err := ecscontainer.Deleteservice(deleteservice)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
