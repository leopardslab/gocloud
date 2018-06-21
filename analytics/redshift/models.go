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

type DeleteCluster struct{
  clusterIdentifier string
  finalClusterSnapshotIdentifier string
  skipFinalClusterSnapshot bool
}
