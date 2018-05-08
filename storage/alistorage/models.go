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
	SnapshotID   string
	ClientToken  string
}

// DeleteDisk to store all attribute to delete Ali-cloud ECS-Disk
type DeleteDisk struct {
	DiskID string
}

// AttachDisk to store all attribute of attach Ali-cloud ECS-Disk to ECS
type AttachDisk struct {
	InstanceID         string
	DiskID             string
	DeleteWithInstance bool
}

// DetachDisk to store all attribute of detach Ali-cloud ECS-Disk from ECS
type DetachDisk struct {
	InstanceID string
	DiskID     string
}

// CreateSnapshot to store all attribute of create Ali-cloud ECS-Disk 's Snapshot
type CreateSnapshot struct {
	DiskID       string
	SnapshotName string
	Description  string
	ClientToken  string
}

// DeleteSnapshot to store all attribute of delete Ali-cloud ECS-Disk 's Snapshot
type DeleteSnapshot struct {
	SnapshotID string
}
