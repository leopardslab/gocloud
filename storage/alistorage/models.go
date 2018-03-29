package alistorage

// Alistorage struct represents Alistorage attribute and method associates with it.
type Alistorage struct {
}

// CreateDisk to store all attribute to create Ali-cloud ECS-Disk
type CreateDisk struct {
	RegionID     string
	ZoneID       string
	DiskName     string
	Description  string
	Encrypted    bool
	DiskCategory string
	Size         int
	SnapshotId   string
	ClientToken  string
}
