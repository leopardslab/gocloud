package googlestorage

type GoogleStorage struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Zone   string `json:"zone"`
	SizeGb string `json:"sizeGb"`
}

type Snapshot struct {
	Name string `json:"name"`
}

type Attachdisk struct {
	Source string `json:"source"`
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

/*

type Attachdisk struct {
	InitializeParams struct {
		DiskName string `json:"diskName"`
		DiskType string `json:"diskType"`
		DiskSizeGb string `json:"diskSizeGb"`
		SourceImage string `json:"sourceImage"`
		SourceImageEncryptionKey struct {
			RawKey string `json:"rawKey"`
			Sha256 string `json:"sha256"`
		} `json:"sourceImageEncryptionKey"`
	} `json:"initializeParams"`
	DeviceName string `json:"deviceName"`
	AutoDelete bool `json:"autoDelete"`
	Boot bool `json:"boot"`
	DiskEncryptionKey struct {
		RawKey string `json:"rawKey"`
		Sha256 string `json:"sha256"`
	} `json:"diskEncryptionKey"`
	Index int `json:"index"`
	Interface string `json:"interface"`
	Kind string `json:"kind"`
	Licenses []string `json:"licenses"`
	Mode string `json:"mode"`
	Source string `json:"source"`
	Type string `json:"type"`
}
*/

type Creatdisk struct {
	Name                         string                      `json:"name"`
	Type                         string                      `json:"type"`
	Zone                         string                      `json:"zone"`
	SizeGb                       string                      `json:"sizeGb"`
	SourceImageEncryptionKeys    SourceImageEncryptionKey    `json:"sourceImageEncryptionKey"`
	DiskEncryptionKeys           DiskEncryptionKey           `json:"diskEncryptionKey"`
	SourceSnapshotEncryptionKeys SourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey"`
	Licenses                     []string                    `json:"licenses"`
	Users                        []string                    `json:"users"`

	CreationTimestamp   string `json:"creationTimestamp"`
	Description         string `json:"description"`
	ID                  string `json:"id"`
	Kind                string `json:"kind"`
	LabelFingerprint    string `json:"labelFingerprint"`
	SourceSnapshotID    string `json:"sourceSnapshotId"`
	Status              string `json:"status"`
	LastAttachTimestamp string `json:"lastAttachTimestamp"`
	LastDetachTimestamp string `json:"lastDetachTimestamp"`
	Options             string `json:"options"`
	SelfLink            string `json:"selfLink"`
	SourceImage         string `json:"sourceImage"`
	SourceImageID       string `json:"sourceImageId"`
	SourceSnapshot      string `json:"sourceSnapshot"`
}

type SourceImageEncryptionKey struct {
	RawKey string `json:"rawKey"`
	Sha256 string `json:"sha256"`
}

type DiskEncryptionKey struct {
	RawKey string `json:"rawKey"`
	Sha256 string `json:"sha256"`
}

type SourceSnapshotEncryptionKey struct {
	RawKey string `json:"rawKey"`
	Sha256 string `json:"sha256"`
}

type Creatdisk1 struct {
	CreationTimestamp string `json:"creationTimestamp"`
	Description       string `json:"description"`
	ID                string `json:"id"`
	Kind              string `json:"kind"`
	LabelFingerprint  string `json:"labelFingerprint"`

	SourceSnapshotID    string `json:"sourceSnapshotId"`
	Status              string `json:"status"`
	LastAttachTimestamp string `json:"lastAttachTimestamp"`
	LastDetachTimestamp string `json:"lastDetachTimestamp"`

	Options  string `json:"options"`
	SelfLink string `json:"selfLink"`

	SourceImage    string `json:"sourceImage"`
	SourceImageID  string `json:"sourceImageId"`
	SourceSnapshot string `json:"sourceSnapshot"`
}
