package redshift

//CreateDatasets Create Datasets.
func (redshift *Redshift) CreateDatasets(request interface{}) (resp interface{}, err error) {

	var region string

	var createCluster CreateCluster

	param := make(map[string]interface{})

	param = request.(map[string]interface{})
	for key, value := range param {
		switch key {

		case "Region":
			regionV, _ := value.(string)
			region = regionV

		case "ClusterIdentifier":
			clusterIdentifier, _ := value.(string)
			createCluster.clusterIdentifier = clusterIdentifier

		case "MasterUsername":
			masterUsername, _ := value.(string)
			createCluster.masterUsername = masterUsername

		case "MasterUserPassword":
			masterUserPassword, _ := value.(string)
			createCluster.masterUserPassword = masterUserPassword

		case "NodeType":
			nodeType, _ := value.(string)
			createCluster.nodeType = nodeType

		case "AdditionalInfo":
			additionalInfo, _ := value.(string)
			createCluster.additionalInfo = additionalInfo

		case "AllowVersionUpgrade":
			allowVersionUpgrade, _ := value.(bool)
			createCluster.allowVersionUpgrade = allowVersionUpgrade


		case "AutomatedSnapshotRetentionPeriod":
			automatedSnapshotRetentionPeriod, _ := value.(int)
			createCluster.automatedSnapshotRetentionPeriod = automatedSnapshotRetentionPeriod

		case "AvailabilityZone":
			availabilityZone, _ := value.(string)
			createCluster.availabilityZone = availabilityZone


		case "ClusterParameterGroupName":
			clusterParameterGroupName, _ := value.(string)
			createCluster.clusterParameterGroupName = clusterParameterGroupName


		case "ClusterSecurityGroups":
			clusterSecurityGroups, _ := value.([]string)
			createCluster.clusterSecurityGroups = clusterSecurityGroups

		case "ClusterType":
			clusterType, _ := value.(string)
			createCluster.clusterType = clusterType


		case "ClusterVersion":
			clusterVersion, _ := value.(string)
			createCluster.clusterVersion = clusterVersion


		case "DBName":
			DBName, _ := value.(string)
			createCluster.dBName = dBName


		case "ElasticIp":
						elasticIp, _ := value.(string)
						createCluster.elasticIp = elasticIp

					case "Encrypted":
									encrypted, _ := value.(bool)
									createCluster.encrypted = encrypted

					case "EnhancedVpcRouting":
							enhancedVpcRouting, _ := value.(bool)
							createCluster.enhancedVpcRouting = enhancedVpcRouting

						case "HsmClientCertificateIdentifier":
								hsmClientCertificateIdentifier, _ := value.(string)
								createCluster.hsmClientCertificateIdentifier = hsmClientCertificateIdentifier


							case "HsmConfigurationIdentifier":
									hsmConfigurationIdentifier, _ := value.(string)
									createCluster.hsmConfigurationIdentifier = hsmConfigurationIdentifier



						case "IamRoles":
									iamRoles, _ := value.([]string)
								  createCluster.iamRoles = iamRoles

								case "KmsKeyId":
										kmsKeyId, _ := value.(string)
										createCluster.kmsKeyId = kmsKeyId


									case "NumberOfNodes":
											numberOfNodes, _ := value.(int)
											createCluster.numberOfNodes = numberOfNodes

										case "Port":
												port, _ := value.(int)
												createCluster.port = port


											case "PreferredMaintenanceWindow":
													preferredMaintenanceWindow, _ := value.(int)
													createCluster.preferredMaintenanceWindow = preferredMaintenanceWindow

												case "TagKeys":
													tagKeys, _ := value.([]string)
													describecluster.tagKeys = tagKeys

												case "TagValues":
													tagValues, _ := value.([]string)
													describecluster.tagValues = tagValues


				case "PubliclyAccessible":
								publiclyAccessible, _ := value.(bool)
								createCluster.publiclyAccessible = PubliclyAccessible


				case "VpcSecurityGroupIds":
					vpcSecurityGroupIds, _ := value.([]string)
					createCluster.vpcSecurityGroupIds = vpcSecurityGroupIds

}
}

	deleteClusterpram := make(map[string]string)

	preparedefaultDeleteClusterpram(deleteClusterpram)
	prepareDeleteClusterpram(deleteClusterpram, deleteCluster)

	response := make(map[string]interface{})

	resp = PrepareSignaturequery(deleteClusterpram, region, response)

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
