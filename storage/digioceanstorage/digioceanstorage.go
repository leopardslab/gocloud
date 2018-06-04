package digioceanstorage

// Digioceanstorage struct represents a DigitalOcean storage object.
type Digioceanstorage struct {
	Name          string `json:"name"`
	Region        string `json:"region,omitempty"`
	Description   string `json:"description,omitempty"`
	SizeGigaBytes int64  `json:"size_gigabytes"`
	SnapshotID    string `json:"snapshot_id,omitempty"`
}
