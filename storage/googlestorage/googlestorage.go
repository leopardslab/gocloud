package googlestorage

type GoogleStorage struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Zone string `json:"zone"`
}


type Snapshot struct {
	Name string `json:"name"`
	CreationTimestamp string `json:"creationTimestamp"`
	DiskSizeGb string `json:"diskSizeGb"`
	ID string `json:"id"`
	Kind string `json:"kind"`
	LabelFingerprint string `json:"labelFingerprint"`
	Labels struct {
		 string `json:""`
	} `json:"labels"`
	Description string `json:"description"`
	SelfLink string `json:"selfLink"`
	Licenses []interface{} `json:"licenses"`
	SnapshotEncryptionKey struct {
	} `json:"snapshotEncryptionKey"`
	SourceDiskEncryptionKey struct {
	} `json:"sourceDiskEncryptionKey"`
	SourceDisk string `json:"sourceDisk"`
	SourceDiskID string `json:"sourceDiskId"`
	Status string `json:"status"`
	StorageBytes string `json:"storageBytes"`
	StorageBytesStatus string `json:"storageBytesStatus"`
}
