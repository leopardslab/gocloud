package vultrstorage

type VultrStorage struct {
}

type CreateSnapshot struct {
	SUBID int
}

type DeleteSnapshot struct {
	SNAPSHOTID string
}
