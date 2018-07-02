package googlecontainer

import "testing"

func TestCreateCluster(t *testing.T) {
	var googlecontainer Googlecontainer

	nodepools := []map[string]interface{}{
		{
			"name":             "default-pool",
			"initialNodeCount": 3,
		},
	}
	createcluster := map[string]interface{}{
		"Project":   "sheltermap-1493101612061",
		"Name":      "cluster-3",
		"Zone":      "us-central1-a",
		"nodePools": nodepools,
	}
	_, err := googlecontainer.CreateCluster(createcluster)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteCluster(t *testing.T) {
	var googlecontainer Googlecontainer

	deletecluster := map[string]string{
		"Project":   "sheltermap-1493101612061",
		"clusterId": "cluster-3",
		"Zone":      "us-central1-a",
	}

	_, err := googlecontainer.DeleteCluster(deletecluster)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestCreateService(t *testing.T) {
	var googlecontainer Googlecontainer
	createservice := map[string]interface{}{
		"Project":   "sheltermap-1493101612061",
		"clusterId": "cluster-3",
		"Zone":      "us-central1-a",
		"Name":      "nodepool",
	}

	_, err := googlecontainer.CreateService(createservice)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteService(t *testing.T) {
	var googlecontainer Googlecontainer

	deleteservice := map[string]string{
		"Project":    "sheltermap-1493101612061",
		"clusterId":  "cluster-3",
		"Zone":       "us-central1-a",
		"nodePoolId": "nodepool",
	}
	_, err := googlecontainer.DeleteService(deleteservice)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
