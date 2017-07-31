package google

import (
	gce "github.com/scorelab/gocloud-v2/compute/gce"
	googlecontainer "github.com/scorelab/gocloud-v2/container/googlecontainer"
	googleloadbalancer "github.com/scorelab/gocloud-v2/loadbalancer/googleloadbalancer"
	googlestorage "github.com/scorelab/gocloud-v2/storage/googlestorage"
	//googledns "github.com/scorelab/gocloud-v2/dns/googledns"
)

//Google  struct represents google cloud provider.
type Google struct {
	gce.GCE
	googlestorage.GoogleStorage
	googleloadbalancer.Googleloadbalancer
	googlecontainer.Googlecontainer
	//googledns.Googledns
}
