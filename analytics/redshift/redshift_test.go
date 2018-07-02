package redshift

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"
import "fmt"

func init() {
	awsAuth.LoadConfig()
}

/*
func TestGetDatasets(t *testing.T) {

	var redshift Redshift

	getDatasets := map[string]interface{}{
		"Region": "us-east-1",
	}

	_, err := redshift.GetDatasets(getDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

	//response := resp.(map[string]interface{})

	//fmt.Println(response["body"])

	fmt.Println("hi")
}



*/

/*

func TestDeleteDatasets(t *testing.T) {

	var redshift Redshift

	deleteDatasets := map[string]interface{}{
		"Region":            "us-east-1",
		"ClusterIdentifier": "test",
	}

	_, err := redshift.DeleteDatasets(deleteDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

	//response := resp.(map[string]interface{})

	//fmt.Println(response["body"])

	fmt.Println("hi")
}

*/


/*
func TestCreateDatasets(t *testing.T) {

	var redshift Redshift

	createDatasets := map[string]interface{}{
		"Region":             "us-east-1",
		"ClusterIdentifier":  "examplecluster",
		"MasterUsername":     "masteruser",
		"MasterUserPassword": "12345678Aa",
		"NumberOfNodes":      "1",
		"NodeType":           "ds2.xlarge",
	}

	_, err := redshift.CreateDatasets(createDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

	//response := resp.(map[string]interface{})

	//fmt.Println(response["body"])

	fmt.Println("hi")
}

*/


func TestUpdateDatasets(t *testing.T) {

	var redshift Redshift

	updateDatasets := map[string]interface{}{
		"Region":             "us-east-1",
		"ClusterIdentifier":  "examplecluster",
		"AutomatedSnapshotRetentionPeriod":     "2",
		"AllowVersionUpgrade": true,
		"ClusterParameterGroupName":      "parametergroup1",
		"PreferredMaintenanceWindow":           "wed:07:30-wed:08:00",
	}

	resp, err := redshift.UpdateDatasets(updateDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])

	fmt.Println("hi")
}
