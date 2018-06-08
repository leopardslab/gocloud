package openstack

import (
	nova "github.com/cloudlibz/gocloud/compute/nova"
	magnum "github.com/cloudlibz/gocloud/container/magnum"
	openstacknosql "github.com/cloudlibz/gocloud/database/openstacknosql"
	designate "github.com/cloudlibz/gocloud/dns/designate"
	neutron "github.com/cloudlibz/gocloud/loadbalancer/neutron"
	openstackmachinelearning "github.com/cloudlibz/gocloud/machinelearning/openstackmachinelearning"
	openstackserverless "github.com/cloudlibz/gocloud/serverless/openstackserverless"
	cinder "github.com/cloudlibz/gocloud/storage/cinder"
)

//openstack  struct represents openstack cloud provider.
type Openstack struct {
	nova.Nova
	cinder.Cinder
	designate.Designate
	magnum.Magnum
	neutron.Neutron
	openstackserverless.Openstackserverless
	openstacknosql.Openstacknosql
	openstackmachinelearning.Openstackmachinelearning
}
