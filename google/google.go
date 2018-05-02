package google

import (
	gce "github.com/cloudlibz/gocloud/compute/gce"
	googlecontainer "github.com/cloudlibz/gocloud/container/googlecontainer"
	googledns "github.com/cloudlibz/gocloud/dns/googledns"
	googleloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/googleloadbalancer"
	googlestorage "github.com/cloudlibz/gocloud/storage/googlestorage"
	googlecloudfunctions "github.com/cloudlibz/gocloud/serverless/googlecloudfunctions"
)

// Google  struct represents Google Cloud provider.
type Google struct {
	gce.GCE
	googlestorage.GoogleStorage
	googleloadbalancer.Googleloadbalancer
	googlecontainer.Googlecontainer
	googledns.Googledns
	googlecloudfunctions.Googlecloudfunctions
}
