package vultrcompute

import (
	"github.com/cloudlibz/gocloud/vultrauth"
	"testing"
)

func init() {
	vultrauth.LoadConfig()
}

func TestVultrCompute_CreateNode(t *testing.T) {
	var vultrServer VultrCompute
	create := map[string]interface{}{
		"DCID":      1,
		"VPSPLANID": 201,
		"OSID":      127,
	}
	resp, err := vultrServer.CreateNode(create)
	if err != nil {
		t.Errorf("CreateNode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr node is created successfully.")
}

func TestCreateNodeBuilder(t *testing.T) {
	var vultrServer VultrCompute
	create, err := NewCreateNodeBuilder().
		DCID(1).
		VPSPLANID(201).
		OSID(127).
		Tag("test").
		Build()
	if err != nil {
		t.Errorf("CreateNode Test Fail: %s", err)
		return
	}
	resp, err := vultrServer.CreateNode(create)
	if err != nil {
		t.Errorf("CreateNode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr node is created successfully.")
}

func TestVultrCompute_StartNode(t *testing.T) {
	var vultrServer VultrCompute
	start := map[string]interface{}{
		"SUBID": 6492936,
	}
	resp, err := vultrServer.StartNode(start)
	if err != nil {
		t.Errorf("StartNode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr node is started successfully.")
}

func TestStartNodeBuilder(t *testing.T) {
	var vultrServer VultrCompute
	start, err := NewStartNodeBuilder().
		SUBID(6492936).
		Build()
	if err != nil {
		t.Errorf("StartNode Test Fail: %s", err)
		return
	}
	resp, err := vultrServer.StartNode(start)
	if err != nil {
		t.Errorf("StartNode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr node is started successfully.")
}

func TestVultrCompute_RebootNode(t *testing.T) {
	var vultrServer VultrCompute
	reboot := map[string]interface{}{
		"SUBID": 6492936,
	}
	resp, err := vultrServer.RebootNode(reboot)
	if err != nil {
		t.Errorf("RebootNode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr node is rebooted successfully.")
}

func TestRebootNodeBuilder(t *testing.T) {
	var vultrServer VultrCompute
	reboot, err := NewRebootNodeBuilder().
		SUBID(6492936).
		Build()
	if err != nil {
		t.Errorf("RebootNode Test Fail: %s", err)
		return
	}
	resp, err := vultrServer.RebootNode(reboot)
	if err != nil {
		t.Errorf("RebootNode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr node is rebooted successfully.")
}

func TestVultrCompute_DeleteNode(t *testing.T) {
	var vultrServer VultrCompute
	destroy := map[string]interface{}{
		"SUBID": 6492936,
	}
	resp, err := vultrServer.DeleteNode(destroy)
	if err != nil {
		t.Errorf("DeleteNode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr node is deleted successfully.")
}

func TestDeleteNodeBuilder(t *testing.T) {
	var vultrServer VultrCompute
	destroy, err := NewDeleteNodeBuilder().
		SUBID(6492936).
		Build()
	if err != nil {
		t.Errorf("DeleteNode Test Fail: %s", err)
		return
	}
	resp, err := vultrServer.DeleteNode(destroy)
	if err != nil {
		t.Errorf("DeleteNode Test Fail: %s", err)
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
	resp, err := vultrServer.ListNode(nil)
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

func TestListNodeBuilder(t *testing.T) {
	var vultrServer VultrCompute
	list, err := NewListNodeBuilder().Build()
	if err != nil {
		t.Errorf("ListNode Test Fail: %s", err)
		return
	}
	resp, err := vultrServer.ListNode(list)
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

func TestParseCreateNodeResp(t *testing.T) {
	var vultrServer VultrCompute
	create := map[string]interface{}{
		"DCID":      1,
		"VPSPLANID": 201,
		"OSID":      127,
	}
	resp, err := vultrServer.CreateNode(create)
	if err != nil {
		t.Errorf("CreateNode Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	createNodeResp, err := ParseCreateNodeResp(response["body"])
	if err != nil {
		t.Errorf("CreateNode Test Fail: %s", err)
		return
	}
	t.Logf("Vultr node is created successfully.\n %+v", createNodeResp)
}
