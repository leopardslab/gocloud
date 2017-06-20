package google

import (
	gce "github.com/scorelab/gocloud-v2/compute/gce"
)

//Google  struct represents google cloud provider.
type Google struct {
	gce.GCE
}
