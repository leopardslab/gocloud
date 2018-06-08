package google

import (
	gce "github.com/cloudlibz/gocloud/compute/gce"
	googlecontainer "github.com/cloudlibz/gocloud/container/googlecontainer"
	bigtable "github.com/cloudlibz/gocloud/database/bigtable"
	googledns "github.com/cloudlibz/gocloud/dns/googledns"
	googleloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/googleloadbalancer"
	googlemachinelearning "github.com/cloudlibz/gocloud/machinelearning/googlemachinelearning"
	googlecloudfunctions "github.com/cloudlibz/gocloud/serverless/googlecloudfunctions"
	googlestorage "github.com/cloudlibz/gocloud/storage/googlestorage"
)

// Google  struct represents Google Cloud provider.
type Google struct {
	gce.GCE
	googlestorage.GoogleStorage
	googleloadbalancer.Googleloadbalancer
	googlecontainer.Googlecontainer
	googledns.Googledns
	googlecloudfunctions.Googlecloudfunctions
	bigtable.Bigtable
	googlemachinelearning.Googlemachinelearning
}
