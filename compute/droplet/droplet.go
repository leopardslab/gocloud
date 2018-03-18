package droplet

import (
  "time"
  "encoding/json"
)

// DropletCreateRequest represents a request to create a Droplet.
type DropletCreateRequest struct {
	Name              string                `json:"name"`
	Region            string                `json:"region"`
	Size              string                `json:"size"`
	Image             DropletCreateImage    `json:"image"`
	SSHKeys           []DropletCreateSSHKey `json:"ssh_keys"`
	Backups           bool                  `json:"backups"`
	IPv6              bool                  `json:"ipv6"`
	PrivateNetworking bool                  `json:"private_networking"`
	Monitoring        bool                  `json:"monitoring"`
	UserData          string                `json:"user_data,omitempty"`
	Volumes           []DropletCreateVolume `json:"volumes,omitempty"`
	Tags              []string              `json:"tags"`
}

// DropletCreateImage identifies an image for the create request. It prefers slug over ID.
type DropletCreateImage struct {
	ID   int	`json:"id,omitempty"`
	Slug string	`json:"name,omitempty"`
}

// MarshalJSON returns either the slug or id of the image. It returns the id
// if the slug is empty.
func (d DropletCreateImage) MarshalJSON() ([]byte, error) {
	if d.Slug != "" {
		return json.Marshal(d.Slug)
	}

	return json.Marshal(d.ID)
}

// DropletCreateSSHKey identifies a SSH Key for the create request. It prefers fingerprint over ID.
type DropletCreateSSHKey struct {
	ID          string
	Fingerprint string
}

// MarshalJSON returns either the fingerprint or id of the ssh key. It returns
// the id if the fingerprint is empty.
func (d DropletCreateSSHKey) MarshalJSON() ([]byte, error) {
	if d.Fingerprint != "" {
		return json.Marshal(d.Fingerprint)
	}

	return json.Marshal(d.ID)
}

// DropletCreateVolume identifies a volume to attach for the create request. It
// prefers Name over ID,
type DropletCreateVolume struct {
	ID   string
	Name string
}

// MarshalJSON returns an object with either the name or id of the volume. It
// returns the id if the name is empty.
func (d DropletCreateVolume) MarshalJSON() ([]byte, error) {
	if d.Name != "" {
		return json.Marshal(struct {
			Name string `json:"name"`
		}{Name: d.Name})
	}

	return json.Marshal(struct {
		ID string `json:"id"`
	}{ID: d.ID})
}

// Droplet represents a DigitalOcean Droplet
type Droplet struct {
	ID               int           `json:"id,float64,omitempty"`
	Name             string        `json:"name,omitempty"`
	Memory           int           `json:"memory,omitempty"`
	Vcpus            int           `json:"vcpus,omitempty"`
	Disk             int           `json:"disk,omitempty"`
	Region           Region       `json:"region,omitempty"`
	Image            Image        `json:"image,omitempty"`
	Size             Size         `json:"size,omitempty"`
	SizeSlug         string        `json:"size_slug,omitempty"`
	BackupIDs        []int         `json:"backup_ids,omitempty"`
	NextBackupWindow BackupWindow `json:"next_backup_window,omitempty"`
	SnapshotIDs      []int         `json:"snapshot_ids,omitempty"`
	Features         []string      `json:"features,omitempty"`
	Locked           bool          `json:"locked,bool,omitempty"`
	Status           string        `json:"status,omitempty"`
	Networks         Networks     `json:"networks,omitempty"`
	Created          string        `json:"created_at,omitempty"`
	Kernel           Kernel       `json:"kernel,omitempty"`
	Tags             []string      `json:"tags,omitempty"`
	VolumeIDs        []string      `json:"volume_ids"`
}

// Region represents a DigitalOcean Region
type Region struct {
	Slug      string   `json:"slug,omitempty"`
	Name      string   `json:"name,omitempty"`
	Sizes     []string `json:"sizes,omitempty"`
	Available bool     `json:"available,omitempty"`
	Features  []string `json:"features,omitempty"`
}

// Image represents a DigitalOcean Image
type Image struct {
	ID           int      `json:"id,float64,omitempty"`
	Name         string   `json:"name,omitempty"`
	Type         string   `json:"type,omitempty"`
	Distribution string   `json:"distribution,omitempty"`
	Slug         string   `json:"slug,omitempty"`
	Public       bool     `json:"public,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	MinDiskSize  int      `json:"min_disk_size,omitempty"`
	Created      string   `json:"created_at,omitempty"`
}

// Size represents a DigitalOcean Size
type Size struct {
	Slug         string   `json:"slug,omitempty"`
	Memory       int      `json:"memory,omitempty"`
	Vcpus        int      `json:"vcpus,omitempty"`
	Disk         int      `json:"disk,omitempty"`
	PriceMonthly float64  `json:"price_monthly,omitempty"`
	PriceHourly  float64  `json:"price_hourly,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	Available    bool     `json:"available,omitempty"`
	Transfer     float64  `json:"transfer,omitempty"`
}

// BackupWindow object
type BackupWindow struct {
	Start time.Time `json:"start,omitempty"`
	End   time.Time `json:"end,omitempty"`
}

// Networks represents the Droplet's Networks.
type Networks struct {
	V4 []NetworkV4 `json:"v4,omitempty"`
	V6 []NetworkV6 `json:"v6,omitempty"`
}

// NetworkV4 represents a DigitalOcean IPv4 Network.
type NetworkV4 struct {
	IPAddress string `json:"ip_address,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Type      string `json:"type,omitempty"`
}

// NetworkV6 represents a DigitalOcean IPv6 network.
type NetworkV6 struct {
	IPAddress string `json:"ip_address,omitempty"`
	Netmask   int    `json:"netmask,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Type      string `json:"type,omitempty"`
}

// Kernel object
type Kernel struct {
	ID      int    `json:"id,float64,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}
