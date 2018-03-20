package openstack

import (
	nova "github.com/shlokgilda/gocloud/compute/nova"
	magnum "github.com/shlokgilda/gocloud/container/magnum"
	designate "github.com/shlokgilda/gocloud/dns/designate"
	neutron "github.com/shlokgilda/gocloud/loadbalancer/neutron"
	cinder "github.com/shlokgilda/gocloud/storage/cinder"
)

//openstack  struct represents openstack cloud provider.
type Openstack struct {
	nova.Nova
	cinder.Cinder
	designate.Designate
	magnum.Magnum
	neutron.Neutron
}
