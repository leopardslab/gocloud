package aliloadbalancer

import (
	"testing"
	"github.com/cloudlibz/gocloud/aliauth"
	"fmt"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreateLoadBalancer(t *testing.T) {
	var alilb Aliloadbalancer
	create := map[string]interface{}{
		"RegionId":           "cn-qingdao",
		"LoadBalancerName":   "abc",
		"AddressType":        "internet",
		"InternetChargeType": "paybytraffic",
	}
	_, err := alilb.Creatloadbalancer(create)
	if err != nil {
		t.Errorf("CreateLoadBalancer Test Fail: %s", err)
		return
	}
	t.Logf("Ali LoadBalancer is created successfully.")
}

func TestDeleteLoadBalancer(t *testing.T) {
	var alilb Aliloadbalancer
	delete := map[string]interface{}{
		"RegionId":       "cn-qingdao",
		"LoadBalancerId": "lb-m5ehdbs3p10a7kmq344je",
	}
	_, err := alilb.Deleteloadbalancer(delete)
	if err != nil {
		t.Errorf("DeleteLoadBalancer Test Fail: %s", err)
		return
	}
	t.Logf("Ali LoadBalancer is deleted successfully.")
}

func TestListLoadBalancer(t *testing.T) {
	var alilb Aliloadbalancer
	list := map[string]interface{}{
		"RegionId":       "cn-qingdao",
		"LoadBalancerId": "lb-m5ehdbs3p10a7kmq344je",
	}
	resp, err := alilb.Listloadbalancer(list)
	if err != nil {
		t.Errorf("DeleteLoadBalancer Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
	t.Logf("Ali LoadBalancer is listed successfully.")
}
