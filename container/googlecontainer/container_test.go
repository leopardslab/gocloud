package googlecontainer

import "testing"
import "fmt"


func TestCreatecluster(t *testing.T) {
	var googlecontainer Googlecontainer

  nodepools := []map[string]interface{}{
    {
      "name": "default-pool",
      "initialNodeCount": 3,
    },
  }
  createcluster := map[string]interface{}{
    "Project":"sheltermap-1493101612061",
    "Name": "cluster-2",
    "Zone": "us-central1-a",
    "nodePools":nodepools,
  }
	_, err := googlecontainer.Createcluster(createcluster)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestDeletecluster(t *testing.T) {
  var googlecontainer Googlecontainer

  deletecluster := map[string]string{
  	"Project":"sheltermap-1493101612061",
  	"clusterId": "cluster-1",
  	"Zone": "us-central1-a",
  }

	_, err := googlecontainer.Deletecluster(deletecluster)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}


func TestCreateservice(t *testing.T) {
  var googlecontainer Googlecontainer
  createservice := map[string]interface{}{
  	"Project":"sheltermap-1493101612061",
  	"clusterId": "cluster-2",
  	"Zone": "us-central1-a",
 	"Name":"nodepool",
  }

	_, err := googlecontainer.Createservice(createservice)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestDeleteservice(t *testing.T) {
  var googlecontainer Googlecontainer

  deleteservice := map[string]string{
  	"Project":"sheltermap-1493101612061",
  	"clusterId": "cluster-2",
  	"Zone": "us-central1-a",
  	"nodePoolId":"nodepool",
  }
	_, err := googlecontainer.Deleteservice(deleteservice)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}
