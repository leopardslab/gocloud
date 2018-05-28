package bigtable

type Bigtable struct {
}

type InitialSplits struct {
	key string
}

type Table struct {
	granularity string
	name        string
}

type Createbigtable struct {
	tableId       string
	table         Table
	initialSplits []InitialSplits
}

type ClusterStates struct {
	replicationState string
}

type GcRule struct {
	maxNumVersions int
	maxAge         string
}
