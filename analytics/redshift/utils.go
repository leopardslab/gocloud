package redshift

import (
	"fmt"
	awsAuth "github.com/cloudlibz/gocloud/awsauth"
	"io/ioutil"
	"net/http"
	"strconv"
)

func preparedefaultDescribeClusterspram(describeclusterspram map[string]string) {

	describeclusterspram["Action"] = "DescribeClusters"
	describeclusterspram["Version"] = "2012-12-01"
}

func prepareDescribeClusterspram(describeclusterspram map[string]string, describecluster Describecluster) {

	if describecluster.clusterIdentifier != "" {
		describeclusterspram["ClusterIdentifier"] = describecluster.clusterIdentifier
	}

	if describecluster.marker != "" {
		describeclusterspram["Marker"] = describecluster.marker
	}

	if describecluster.maxRecords != 0 {
		describeclusterspram["MaxRecords"] = strconv.Itoa(describecluster.maxRecords)
	}

	if len(describecluster.tagKeys) != 0 {

		for i := 0; i < len(describecluster.tagKeys); i++ {

			n := strconv.Itoa(i)

			prefix := "TagKeys.TagKey." + n

			describeclusterspram[prefix] = describecluster.tagKeys[i]

		}
	}

	if len(describecluster.tagValues) != 0 {

		for i := 0; i < len(describecluster.tagValues); i++ {

			n := strconv.Itoa(i)

			prefix := "TagValues.TagValue." + n

			describeclusterspram[prefix] = describecluster.tagValues[i]
		}
	}
}

func PrepareSignaturequery(describeclusterspram map[string]string, region string) (response map[string]interface{}, err error) {

	service := "redshift"

	endpoint := "https://" + service + "." + region + ".amazonaws.com"

	req, err := http.NewRequest("GET", endpoint, nil)


	// Add the params passed in to the query string
	query := req.URL.Query()
	for varName, varVal := range describeclusterspram {
		query.Add(varName, varVal)
	}

	req.URL.RawQuery = query.Encode()


	awsAuth.Getrequestsign4(req, region, service)

	r, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	response = make(map[string]interface{})

	response["body"] = string(body)

	response["status"] = r.StatusCode

	return response,err
}


func preparedefaultCreateClusterpram(createClusterpram map[string]string) {

	createClusterpram["Action"] = "CreateCluster"
	createClusterpram["Version"] = "2012-12-01"
}

func preparecreateClusterpram(createClusterpram map[string]string, createCluster CreateCluster) {

	if createCluster.clusterIdentifier != "" {
		createClusterpram["ClusterIdentifier"] = createCluster.clusterIdentifier
	}

	if createCluster.masterUsername != "" {
		createClusterpram["MasterUsername"] = createCluster.masterUsername
	}

	if createCluster.masterUserPassword != "" {
		createClusterpram["MasterUserPassword"] = createCluster.masterUserPassword
	}

	if createCluster.nodeType != "" {
		createClusterpram["NodeType"] = createCluster.nodeType
	}

	if createCluster.additionalInfo != "" {
		createClusterpram["AdditionalInfo"] = createCluster.additionalInfo
	}

	if createCluster.allowVersionUpgrade == true {
		createClusterpram["AllowVersionUpgrade"] = "true"
	}

	if createCluster.automatedSnapshotRetentionPeriod != "" {
		createClusterpram["AutomatedSnapshotRetentionPeriod"] = createCluster.automatedSnapshotRetentionPeriod
	}

	if createCluster.availabilityZone != "" {
		createClusterpram["AvailabilityZone"] = createCluster.availabilityZone
	}

	if createCluster.clusterParameterGroupName != "" {
		createClusterpram["ClusterParameterGroupName"] = createCluster.clusterParameterGroupName
	}

	if createCluster.clusterType != "" {
		createClusterpram["ClusterType"] = createCluster.clusterType
	}

	if createCluster.clusterVersion != "" {
		createClusterpram["ClusterVersion"] = createCluster.clusterVersion
	}

	if createCluster.dBName != "" {
		createClusterpram["DBName"] = createCluster.dBName
	}

	if createCluster.elasticIp != "" {
		createClusterpram["ElasticIp"] = createCluster.elasticIp
	}

	if createCluster.encrypted == true {
		createClusterpram["Encrypted"] = "true"
	}

	if createCluster.enhancedVpcRouting == true {
		createClusterpram["EnhancedVpcRouting"] = "true"
	}

	if createCluster.hsmClientCertificateIdentifier != "" {
		createClusterpram["HsmClientCertificateIdentifier"] = createCluster.hsmClientCertificateIdentifier
	}

	if createCluster.hsmConfigurationIdentifier != "" {
		createClusterpram["HsmConfigurationIdentifier"] = createCluster.hsmConfigurationIdentifier
	}

	if createCluster.kmsKeyId != "" {
		createClusterpram["KmsKeyId"] = createCluster.kmsKeyId
	}

	if createCluster.numberOfNodes > 0 {
		createClusterpram["NumberOfNodes"] = strconv.Itoa(createCluster.numberOfNodes)
	}

	if createCluster.port != 0 {
		createClusterpram["Port"] =strconv.Itoa(createCluster.port)
	}

	if createCluster.preferredMaintenanceWindow != "" {
		createClusterpram["PreferredMaintenanceWindow"] = createCluster.preferredMaintenanceWindow
	}

	if createCluster.publiclyAccessible != true {
		createClusterpram["PubliclyAccessible"] = "true"
	}

	if len(createCluster.iamRoles) != 0 {

		for i := 0; i < len(createCluster.iamRoles); i++ {

			n := strconv.Itoa(i)

			prefix := "IamRoles.IamRoleArn." + n

			createClusterpram[prefix] = createCluster.iamRoles[i]
		}
	}

	if len(createCluster.clusterSecurityGroups) != 0 {

		for i := 0; i < len(createCluster.clusterSecurityGroups); i++ {

			n := strconv.Itoa(i)

			prefix := "ClusterSecurityGroups.ClusterSecurityGroupName." + n

			createClusterpram[prefix] = createCluster.clusterSecurityGroups[i]
		}
	}

	if len(createCluster.vpcSecurityGroupIds) != 0 {

		for i := 0; i < len(createCluster.vpcSecurityGroupIds); i++ {

			n := strconv.Itoa(i)

			prefix := "VpcSecurityGroupIds.VpcSecurityGroupId." + n

			createClusterpram[prefix] = createCluster.vpcSecurityGroupIds[i]
		}
	}

	if len(createCluster.tagKeys) != 0 {

		for i := 0; i < len(createCluster.tagKeys); i++ {

			n := strconv.Itoa(i)

			prefix := "TagKeys.TagKey." + n

			createClusterpram[prefix] = createCluster.tagKeys[i]

		}
	}

	if len(createCluster.tagValues) != 0 {

		for i := 0; i < len(createCluster.tagValues); i++ {

			n := strconv.Itoa(i)

			prefix := "TagValues.TagValue." + n

			createClusterpram[prefix] = createCluster.tagValues[i]
		}
	}

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

	if deleteCluster.skipFinalClusterSnapshot == true {
		deleteClusterpram["SkipFinalClusterSnapshot"] = "true"
	}

	if deleteCluster.skipFinalClusterSnapshot == false {
		deleteClusterpram["SkipFinalClusterSnapshot"] = "false"
	}

}

func 	preparedefaultupdateClusterpram(createClusterpram map[string]string){

createClusterpram["Action"] = "ModifyCluster"
createClusterpram["Version"] = "2012-12-01"
}
