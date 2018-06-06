package bigtable

//Createbigtable struct reperesnts  Google bigtable.
type Bigtable struct {
}

//InitialSplits struct reperesnts  InitialSplits.
type InitialSplits struct {
	key string
}

//Table struct reperesnts  Table.
type Table struct {
	granularity string
	name        string
}

//Createbigtable struct reperesnts  Create bigtable.
type Createbigtable struct {
	tableId       string
	table         Table
	initialSplits []InitialSplits
	ClusterStates ClusterStates
}

//ClusterStates struct reperesnts  ClusterStates.
type ClusterStates struct {
	replicationState string
}

//GcRule struct reperesnts  GcRule.
type GcRule struct {
	maxNumVersions int
	maxAge         string
}
