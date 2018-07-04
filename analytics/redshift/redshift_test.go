package redshift

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}

func TestCreateDatasets(t *testing.T) {

	var redshift Redshift

	createDatasets := map[string]interface{}{
		"Region":             "us-east-1",
		"ClusterIdentifier":  "examplecluster",
		"MasterUsername":     "masteruser",
		"MasterUserPassword": "12345678Aa",
		"NumberOfNodes":      3,
		"NodeType":           "ds2.xlarge",
	}

	_, err := redshift.CreateDatasets(createDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

}

func TestGetDatasets(t *testing.T) {

	var redshift Redshift

	getDatasets := map[string]interface{}{
		"Region": "us-east-1",
	}

	_, err := redshift.GetDatasets(getDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

}

func TestUpdateDatasets(t *testing.T) {

	var redshift Redshift

	updateDatasets := map[string]interface{}{
		"Region":                           "us-east-1",
		"ClusterIdentifier":                "examplecluster",
		"AutomatedSnapshotRetentionPeriod": "2",
		"AllowVersionUpgrade":              true,
		"ClusterParameterGroupName":        "parametergroup1",
		"PreferredMaintenanceWindow":       "wed:07:30-wed:08:00",
	}

	_, err := redshift.UpdateDatasets(updateDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}

}

func TestDeleteDatasets(t *testing.T) {

	var redshift Redshift

	deleteDatasets := map[string]interface{}{
		"Region":                   "us-east-1",
		"ClusterIdentifier":        "examplecluster",
		"SkipFinalClusterSnapshot": true,
	}

	_, err := redshift.DeleteDatasets(deleteDatasets)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
