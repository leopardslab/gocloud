package aliloadbalancer

import (
	"fmt"
	"github.com/cloudlibz/gocloud/aliauth"
	"testing"
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
		"RegionId": "cn-qingdao",
	}
	resp, err := alilb.Listloadbalancer(list)
	if err != nil {
		t.Errorf("ListLoadBalancer Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
	t.Logf("Ali LoadBalancer is listed successfully.")
}

func TestAttachLoadBalancer(t *testing.T) {
	var alilb Aliloadbalancer
	attach := map[string]interface{}{
		"LoadBalancerId": "lb-m5eemmwmtmt4l6jf73d72",
		"BackendServers": "[{'ServerId':'i-m5ecx5g9m0cflv1k83zu','Weight':'100','Type':'ecs'}," +
			"{'ServerId':'i-m5eahbbwqxawpj1opww9','Weight':'100','Type':'ecs'}]",
	}
	_, err := alilb.Attachnodewithloadbalancer(attach)
	if err != nil {
		t.Errorf("Attachnodewithloadbalancer Test Fail: %s", err)
		return
	}
	t.Logf("Ali LoadBalancer is listed successfully.")
}

func TestDettachLoadBalancer(t *testing.T) {
	var alilb Aliloadbalancer
	dettach := map[string]interface{}{
		"RegionId":       "cn-qingdao",
		"LoadBalancerId": "lb-m5eemmwmtmt4l6jf73d72",
		"BackendServers": "[{'ServerId':'i-m5ecx5g9m0cflv1k83zu','Type':'ecs'}," +
			"{'ServerId':'i-m5eahbbwqxawpj1opww9','Type':'ecs'}]",
	}
	_, err := alilb.Detachnodewithloadbalancer(dettach)
	if err != nil {
		t.Errorf("DetachnodewithLoadBalancer Test Fail: %s", err)
		return
	}
	t.Logf("Ali LoadBalancer is detached with node successfully.")
}
