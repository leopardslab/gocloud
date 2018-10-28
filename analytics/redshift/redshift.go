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
			automatedSnapshotRetentionPeriod, _ := value.(string)
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
			dBName, _ := value.(string)
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
			preferredMaintenanceWindow, _ := value.(string)
			createCluster.preferredMaintenanceWindow = preferredMaintenanceWindow

		case "TagKeys":
			tagKeys, _ := value.([]string)
			createCluster.tagKeys = tagKeys

		case "TagValues":
			tagValues, _ := value.([]string)
			createCluster.tagValues = tagValues

		case "PubliclyAccessible":
			publiclyAccessible, _ := value.(bool)
			createCluster.publiclyAccessible = publiclyAccessible

		case "VpcSecurityGroupIds":
			vpcSecurityGroupIds, _ := value.([]string)
			createCluster.vpcSecurityGroupIds = vpcSecurityGroupIds
		}
	}

	createClusterpram := make(map[string]string)

	preparedefaultCreateClusterpram(createClusterpram)
	preparecreateClusterpram(createClusterpram, createCluster)

	resp, err = PrepareSignaturequery(createClusterpram, region)

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

	resp, err = PrepareSignaturequery(deleteClusterpram, region)

	return resp, err
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

	resp, err = PrepareSignaturequery(describeclusterspram, region)

	return resp, err
}

//UpdateDatasets  Update Datasets.
func (redshift *Redshift) UpdateDatasets(request interface{}) (resp interface{}, err error) {
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
			automatedSnapshotRetentionPeriod, _ := value.(string)
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
			dBName, _ := value.(string)
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
			preferredMaintenanceWindow, _ := value.(string)
			createCluster.preferredMaintenanceWindow = preferredMaintenanceWindow

		case "TagKeys":
			tagKeys, _ := value.([]string)
			createCluster.tagKeys = tagKeys

		case "TagValues":
			tagValues, _ := value.([]string)
			createCluster.tagValues = tagValues

		case "PubliclyAccessible":
			publiclyAccessible, _ := value.(bool)
			createCluster.publiclyAccessible = publiclyAccessible

		case "VpcSecurityGroupIds":
			vpcSecurityGroupIds, _ := value.([]string)
			createCluster.vpcSecurityGroupIds = vpcSecurityGroupIds

		}
	}

	createClusterpram := make(map[string]string)

	preparedefaultupdateClusterpram(createClusterpram)
	preparecreateClusterpram(createClusterpram, createCluster)

	resp, err = PrepareSignaturequery(createClusterpram, region)

	return resp, err
}

//ListDatasets  list Datasets.
func (redshift *Redshift) ListDatasets(request interface{}) (resp interface{}, err error) {
	return resp, err
}
