package google

import (
	gce "github.com/scorelab/gocloud/compute/gce"
)

//Google  struct represents google cloud provider.
type Google struct {
	gce.GCE
}
