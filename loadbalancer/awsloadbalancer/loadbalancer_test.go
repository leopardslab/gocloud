package awsloadbalancer

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}

func TestCreateLoadBalancer(t *testing.T) {
	var awsloadbalancer Awsloadbalancer
	Listeners := []map[string]string{{
		"InstancePort":     "80",
		"LoadBalancerPort": "80",
		"Protocol":         "http",
		"InstanceProtocol": "http",
		"SSLCertificateId": "",
	},
	}

	Subnets := []string{"subnet-b59a4599", "subnet-32f9727a"}

	creatloadbalancer := map[string]interface{}{
		"LoadBalancerName": "my-load-balancer",
		"Listeners":        Listeners,
		"Subnets":          Subnets,
	}

	awsloadbalancer.CreateLoadBalancer(creatloadbalancer)
}

func TestDeleteLoadBalancer(t *testing.T) {
	var awsloadbalancer Awsloadbalancer

	deleteloadbalancer := map[string]string{
		"LoadBalancerName": "my-load-balancer",
	}
	_, err := awsloadbalancer.DeleteLoadBalancer(deleteloadbalancer)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListLoadBalancer(t *testing.T) {
	var awsloadbalancer Awsloadbalancer

	_, err := awsloadbalancer.ListLoadBalancer(nil)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestAttachNodeWithLoadBalancer(t *testing.T) {
	var awsloadbalancer Awsloadbalancer

	attachnodewithloadbalancer := map[string]interface{}{
		"Instances":        []string{"i-05f4f2535c41b680b"},
		"LoadBalancerName": "my-load-balancer",
	}

	_, err := awsloadbalancer.AttachNodeWithLoadBalancer(attachnodewithloadbalancer)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDetachNodeWithLoadBalancer(t *testing.T) {
	var awsloadbalancer Awsloadbalancer

	detachnodewithloadbalancer := map[string]interface{}{
		"Instances":        []string{"i-05f4f2535c41b680b"},
		"LoadBalancerName": "my-load-balancer",
	}
	_, err := awsloadbalancer.DetachNodeWithLoadBalancer(detachnodewithloadbalancer)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
