package awsloadbalancer

import "testing"
import "github.com/cloudlibz/gocloud/auth"

func init() {
	auth.LoadConfig()
}

func TestCreatloadbalancer(t *testing.T) {
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

	awsloadbalancer.Creatloadbalancer(creatloadbalancer)
}

func TestDeleteloadbalancer(t *testing.T) {
	var awsloadbalancer Awsloadbalancer

	deleteloadbalancer := map[string]string{
		"LoadBalancerName": "my-load-balancer",
	}
	_, err := awsloadbalancer.Deleteloadbalancer(deleteloadbalancer)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListloadbalancer(t *testing.T) {
	var awsloadbalancer Awsloadbalancer

	_, err := awsloadbalancer.Listloadbalancer(nil)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestAttachnodewithloadbalancer(t *testing.T) {
	var awsloadbalancer Awsloadbalancer

	attachnodewithloadbalancer := map[string]interface{}{
		"Instances":        []string{"i-05f4f2535c41b680b"},
		"LoadBalancerName": "my-load-balancer",
	}

	_, err := awsloadbalancer.Attachnodewithloadbalancer(attachnodewithloadbalancer)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDetachnodewithloadbalancer(t *testing.T) {
	var awsloadbalancer Awsloadbalancer

	detachnodewithloadbalancer := map[string]interface{}{
		"Instances":        []string{"i-05f4f2535c41b680b"},
		"LoadBalancerName": "my-load-balancer",
	}
	_, err := awsloadbalancer.Detachnodewithloadbalancer(detachnodewithloadbalancer)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
