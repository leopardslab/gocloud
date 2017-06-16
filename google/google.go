package google

import (
	gce "github.com/scorelab/gocloud/compute/gce"
)

//struct repersents Google cloud
type Google struct {
	gce.GCE
}
