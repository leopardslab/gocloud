package vultrcompute

import (
	"github.com/cloudlibz/gocloud/vultrauth"
	"testing"
)

func init() {
	vultrauth.LoadConfig()
}

func TestVultrCompute_Createnode(t *testing.T) {
	var vultrServer VultrCompute
	create := map[string]interface{}{
		"DCID":      1,
		"VPSPLANID": 201,
		"OSID":      127,
	}
	resp, err := vultrServer.Createnode(create)
	if err != nil {
		t.Errorf("Createnode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr node is created successfully.")
}

func TestVultrCompute_Startnode(t *testing.T) {
	var vultrServer VultrCompute
	start := map[string]interface{}{
		"SUBID": 6492936,
	}
	resp, err := vultrServer.Startnode(start)
	if err != nil {
		t.Errorf("Startnode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr node is started successfully.")
}

func TestVultrCompute_Rebootnode(t *testing.T) {
	var vultrServer VultrCompute
	reboot := map[string]interface{}{
		"SUBID": 6492936,
	}
	resp, err := vultrServer.Rebootnode(reboot)
	if err != nil {
		t.Errorf("Rebootnode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr node is rebooted successfully.")
}

func TestVultrCompute_Deletenode(t *testing.T) {
	var vultrServer VultrCompute
	destroy := map[string]interface{}{
		"SUBID": 6492936,
	}
	resp, err := vultrServer.Deletenode(destroy)
	if err != nil {
		t.Errorf("Deletenode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr node is deleted successfully.")
}

func TestVultrCompute_ListNode(t *testing.T) {
	var vultrServer VultrCompute
	resp, err := vultrServer.ListNode()
	if err != nil {
		t.Errorf("ListNode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr node list: %s", response["body"])
}
