package awsloadbalancer

import "testing"
import "fmt"
import "github.com/scorelab/gocloud-v2/auth"

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
	_, err := awsloadbalancer.Creatloadbalancer(creatloadbalancer)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestDeleteloadbalancer(t *testing.T) {
	var awsloadbalancer Awsloadbalancer

	deleteloadbalancer := map[string]string{
		"LoadBalancerName": "my-load-balancer",
	}
	_, err := awsloadbalancer.Deleteloadbalancer(deleteloadbalancer)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}

func TestListloadbalancer(t *testing.T) {
	var awsloadbalancer Awsloadbalancer

	_, err := awsloadbalancer.Listloadbalancer(nil)

	if err != nil {
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
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
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
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
		fmt.Println("Test Fail")
	} else {
		fmt.Println("Test Pass")
	}
}
