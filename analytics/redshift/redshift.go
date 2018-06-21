package redshift

//CreateDatasets Create Datasets.
func (redshift *Redshift) CreateDatasets(request interface{}) (resp interface{}, err error) {

	return resp, err
}

//DeleteDatasets delete Datasets.
func (redshift *Redshift) DeleteDatasets(request interface{}) (resp interface{}, err error) {

	var region string

	var deleteCluster DeleteCluster

	param := make(map[string]interface{})

	param = request.(map[string]interface{})
	for key, value := range param {
		switch key {

		case "Region":
			regionV, _ := value.(string)
			region = regionV

		case "ClusterIdentifier":
			clusterIdentifier, _ := value.(string)
			deleteCluster.clusterIdentifier = clusterIdentifier

		case "FinalClusterSnapshotIdentifier":
			finalClusterSnapshotIdentifier, _ := value.(string)
			deleteCluster.finalClusterSnapshotIdentifier = finalClusterSnapshotIdentifier

		case "SkipFinalClusterSnapshot":
			skipFinalClusterSnapshot, _ := value.(bool)
			deleteCluster.skipFinalClusterSnapshot = skipFinalClusterSnapshot
		}
	}



	deleteClusterpram := make(map[string]string)

	preparedefaultDeleteClusterpram(deleteClusterpram)
	prepareDeleteClusterpram(deleteClusterpram, deleteCluster)

	response := make(map[string]interface{})

	resp = PrepareSignaturequery(deleteClusterpram, region, response)

	return resp, err
}


func preparedefaultDeleteClusterpram(deleteClusterpram map[string]string) {

	deleteClusterpram["Action"] = "DeleteCluster"
	deleteClusterpram["Version"] = "2012-12-01"
}




func prepareDeleteClusterpram(deleteClusterpram map[string]string, deleteCluster DeleteCluster) {


	if deleteCluster.clusterIdentifier != "" {
		deleteClusterpram["ClusterIdentifier"] = deleteCluster.clusterIdentifier
	}

	if deleteCluster.finalClusterSnapshotIdentifier != "" {
		deleteClusterpram["finalClusterSnapshotIdentifier"] = deleteCluster.finalClusterSnapshotIdentifier
	}

	if deleteCluster.skipFinalClusterSnapshot ==  true {
			deleteClusterpram["SkipFinalClusterSnapshot"] = "true"
	}

	if deleteCluster.skipFinalClusterSnapshot ==  false {
			deleteClusterpram["SkipFinalClusterSnapshot"] = "false"
	}

}


//GetDatasets get Datasets.
func (redshift *Redshift) GetDatasets(request interface{}) (resp interface{}, err error) {

	var region string

	var describecluster Describecluster

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {

		case "Region":
			regionV, _ := value.(string)
			region = regionV

		case "ClusterIdentifier":
			clusterIdentifier, _ := value.(string)
			describecluster.clusterIdentifier = clusterIdentifier

		case "Marker":
			marker, _ := value.(string)
			describecluster.marker = marker

		case "MaxRecords":
			maxRecords, _ := value.(int)
			describecluster.maxRecords = maxRecords

		case "TagKeys":
			tagKeys, _ := value.([]string)
			describecluster.tagKeys = tagKeys

		case "TagValues":
			tagValues, _ := value.([]string)
			describecluster.tagValues = tagValues
		}
	}

	describeclusterspram := make(map[string]string)

	preparedefaultDescribeClusterspram(describeclusterspram)
	prepareDescribeClusterspram(describeclusterspram, describecluster)

	response := make(map[string]interface{})

	resp = PrepareSignaturequery(describeclusterspram, region, response)

	return resp, err
}

//UpdateDatasets  Update Datasets.
func (redshift *Redshift) UpdateDatasets(request interface{}) (resp interface{}, err error) {

	return resp, err
}


//ListDatasets  list Datasets.
func (redshift *Redshift) ListDatasets(request interface{}) (resp interface{}, err error) {
	return resp, err

}
