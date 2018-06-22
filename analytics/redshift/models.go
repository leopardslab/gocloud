package redshift

//redshift struct reperesnts aws alianalytics service.
type Redshift struct {
}

type Describecluster struct {
	clusterIdentifier string
	marker            string
	maxRecords        int
	tagKeys           []string
	tagValues         []string
}

type DeleteCluster struct {
	clusterIdentifier              string
	finalClusterSnapshotIdentifier string
	skipFinalClusterSnapshot       bool
}

type CreateCluster struct {
	clusterIdentifier                string
	masterUsername                   string
	masterUserPassword               string
	nodeType                         string
	additionalInfo                   string
	AllowVersionUpgrade              bool
	automatedSnapshotRetentionPeriod string
	availabilityZone                 string
	clusterParameterGroupName        string
	clusterType                      string
	clusterVersion                   string
	dBName                           string
	elasticIp                        string
	encrypted                        bool
	enhancedVpcRouting               string
	hsmClientCertificateIdentifier   string
	hsmConfigurationIdentifier       string
	kmsKeyId                         string
	numberOfNodes                    int
	port                             int
	PreferredMaintenanceWindow       string
	PubliclyAccessible               bool
	tagKeys                          []string
	tagValues                        []string
	iamRoles                         []string
	clusterSecurityGroups            []string
	vpcSecurityGroupIds              []string
}
