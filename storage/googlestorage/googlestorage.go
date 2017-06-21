package googlestorage

type GoogleStorage struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Zone string `json:"zone"`
}

type Snapshot struct {
	Name string `json:"name"`
}

/*
type Snapshot struct {
		CreationTimestamp string `json:"creationTimestamp"`
		Description string `json:"description"`
		DiskSizeGb string `json:"diskSizeGb"`
		ID string `json:"id"`
		Kind string `json:"kind"`
		Labels struct {
		} `json:"labels"`
		LabelFingerprint string `json:"labelFingerprint"`
		Licenses []string `json:"licenses"`
		Name string `json:"name"`
		SelfLink string `json:"selfLink"`
		SnapshotEncryptionKey struct {
			RawKey string `json:"rawKey"`
			Sha256 string `json:"sha256"`
		} `json:"snapshotEncryptionKey"`
		SourceDisk string `json:"sourceDisk"`
		SourceDiskEncryptionKey struct {
			RawKey string `json:"rawKey"`
			Sha256 string `json:"sha256"`
		} `json:"sourceDiskEncryptionKey"`
		SourceDiskID string `json:"sourceDiskId"`
		Status string `json:"status"`
		StorageBytes string `json:"storageBytes"`
		StorageBytesStatus string `json:"storageBytesStatus"`
	}
*/
