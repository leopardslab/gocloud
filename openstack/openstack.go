package openstack

import (
	nova "github.com/cloudlibz/gocloud/compute/nova"
)

//openstack  struct represents openstack cloud provider.
type Openstack struct {
	nova.Nova
}
