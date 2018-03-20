package google

import (
	gce "github.com/shlokgilda/gocloud/compute/gce"
	googlecontainer "github.com/shlokgilda/gocloud/container/googlecontainer"
	googledns "github.com/shlokgilda/gocloud/dns/googledns"
	googleloadbalancer "github.com/shlokgilda/gocloud/loadbalancer/googleloadbalancer"
	googlestorage "github.com/shlokgilda/gocloud/storage/googlestorage"
)

// Google  struct represents Google Cloud provider.
type Google struct {
	gce.GCE
	googlestorage.GoogleStorage
	googleloadbalancer.Googleloadbalancer
	googlecontainer.Googlecontainer
	googledns.Googledns
}
