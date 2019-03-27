package rackspacecompute

// Rackspacecompute represents a request to create a Rackspace compute instance.
type Rackspacecompute struct {
	Server Cloudserver `json:"server"`
}

// Cloudserver represents a node of cloud server on rackspace
type Cloudserver struct {
	Name      string `json:"name"`
	ImageRef  string `json:"imageRef"`
	FlavorRef string `json:"flavorRef"`
}
