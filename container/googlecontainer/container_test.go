package googlecontainer

import "testing"

func TestCreatecluster(t *testing.T) {
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
	_, err := googlecontainer.Createcluster(createcluster)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeletecluster(t *testing.T) {
	var googlecontainer Googlecontainer

	deletecluster := map[string]string{
		"Project":   "sheltermap-1493101612061",
		"clusterId": "cluster-3",
		"Zone":      "us-central1-a",
	}

	_, err := googlecontainer.Deletecluster(deletecluster)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestCreateservice(t *testing.T) {
	var googlecontainer Googlecontainer
	createservice := map[string]interface{}{
		"Project":   "sheltermap-1493101612061",
		"clusterId": "cluster-3",
		"Zone":      "us-central1-a",
		"Name":      "nodepool",
	}

	_, err := googlecontainer.Createservice(createservice)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteservice(t *testing.T) {
	var googlecontainer Googlecontainer

	deleteservice := map[string]string{
		"Project":    "sheltermap-1493101612061",
		"clusterId":  "cluster-3",
		"Zone":       "us-central1-a",
		"nodePoolId": "nodepool",
	}
	_, err := googlecontainer.Deleteservice(deleteservice)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
