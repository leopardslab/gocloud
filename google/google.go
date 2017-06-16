package google

import (
	gce "github.com/scorelab/gocloud/compute/gce"
)

//struct represents Google cloud
type Google struct {
	gce.GCE
}
