package awscontainer

import "testing"
import "fmt"

import "github.com/scorelab/gocloud-v2/auth"

func init() {
	auth.LoadConfig()
}

func TestCreatecluster(t *testing.T) {
	var ecscontainer Ecscontainer

  createcluster := map[string]interface{}{
    "clusterName": "rootmonk",
    "Region":      "us-east-1",
  }

	_, err := ecscontainer.Createcluster(createcluster)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}


func TestDeletecluster(t *testing.T) {
	var ecscontainer Ecscontainer

  deletecluster := map[string]interface{}{
    "clusterName": "rootmonk",
    "Region":      "us-east-1",
  }

	_, err := ecscontainer.Deletecluster(deletecluster)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}
