package openstack

import (
	nova "github.com/cloudlibz/gocloud/compute/nova"
	magnum "github.com/cloudlibz/gocloud/container/magnum"
	designate "github.com/cloudlibz/gocloud/dns/designate"
	neutron "github.com/cloudlibz/gocloud/loadbalancer/neutron"
	cinder "github.com/cloudlibz/gocloud/storage/cinder"
	openstackserverless "github.com/cloudlibz/gocloud/serverless/openstackserverless"
)

//openstack  struct represents openstack cloud provider.
type Openstack struct {
	nova.Nova
	cinder.Cinder
	designate.Designate
	magnum.Magnum
	neutron.Neutron
	openstackserverless.Openstackserverless
}
