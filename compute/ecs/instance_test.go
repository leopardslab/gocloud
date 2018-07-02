package ecs

import (
	"fmt"
	"github.com/cloudlibz/gocloud/aliauth"
	"testing"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreateNode(t *testing.T) {
	var aliEcs ECS
	create := map[string]interface{}{
		"RegionId":            "cn-qingdao",
		"ImageId":             "centos_7_04_64_20G_alibase_201701015.vhd",
		"InstanceType":        "ecs.xn4.small",
		"SecurityGroupId":     "sg-m5egbo9s5xb21kpu6nk2",
		"DataDisk.1.Size":     "20",
		"DataDisk.1.DiskName": "testName",
	}
	_, err := aliEcs.CreateNode(create)
	if err != nil {
		t.Errorf("CreateNode Test Fail: %s", err)
		return
	}
	t.Logf("Ali node is created successfully.")
}
func TestCreateNodeBuilder(t *testing.T) {
	var aliEcs ECS
	createNode, err := NewCreateNodeBuilder().
		RegionID("cn-qingdao").
		ImageID("centos_7_04_64_20G_alibase_201701015.vhd").
		InstanceType("ecs.xn4.small").
		SecurityGroupID("sg-m5egbo9s5xb21kpu6nk2").
		Build()
	if err != nil {
		t.Errorf("CreateNode Test Fail: %s", err.Error())
		return
	}
	_, err = aliEcs.CreateNode(createNode)
	if err != nil {
		t.Errorf("CreateNode Test Fail: %s", err)
		return
	}
	t.Logf("Ali node is created successfully.")
}

func TestStartNode(t *testing.T) {
	var aliEcs ECS
	start := map[string]interface{}{
		"InstanceId": "i-m5ea01zic0bg8dufo2eo",
	}
	_, err := aliEcs.StartNode(start)
	if err != nil {
		t.Errorf("StartNode Test Fail")
		return
	}
	t.Logf("Ali node is started successfully.")
}
func TestStartNodeBuilder(t *testing.T) {
	var aliEcs ECS
	startNode, err := NewStartNodeBuilder().
		InstanceID("m5ea01zic0bg8dufo2eo").
		Build()
	if err != nil {
		t.Errorf("StartNode Test Fail: %s", err.Error())
		return
	}
	_, err = aliEcs.StartNode(startNode)
	if err != nil {
		t.Errorf("StartNode Test Fail: %s", err)
		return
	}
	t.Logf("Ali node is started successfully.")
}

func TestStopNode(t *testing.T) {
	var aliEcs ECS
	stop := map[string]interface{}{
		"InstanceId": "i-m5ea01zic0bg8dufo2eo",
		"ForceStop":  false,
	}
	_, err := aliEcs.StopNode(stop)
	if err != nil {
		t.Errorf("StopNode Test Fail")
		return
	}
	t.Logf("Ali node is stoped successfully.")
}
func TestStopNodeBuilder(t *testing.T) {
	var aliEcs ECS
	stopNode, err := NewStopNodeBuilder().
		InstanceID("m5ea01zic0bg8dufo2eo").
		Build()
	if err != nil {
		t.Errorf("StopNode Test Fail: %s", err.Error())
		return
	}
	_, err = aliEcs.StopNode(stopNode)
	if err != nil {
		t.Errorf("StopNode Test Fail: %s", err)
		return
	}
	t.Logf("Ali node is stoped successfully.")
}

func TestRebootNode(t *testing.T) {
	var aliEcs ECS
	reboot := map[string]interface{}{
		"InstanceId": "i-m5ea01zic0bg8dufo2eo",
		"ForceStop":  false,
	}
	_, err := aliEcs.RebootNode(reboot)
	if err != nil {
		t.Errorf("RebootNode Test Fail")
		return
	}
	t.Logf("Ali node is rebooted successfully.")
}
func TestRebootNodeBuilder(t *testing.T) {
	var aliEcs ECS
	rebootNode, err := NewRebootNodeBuilder().
		InstanceID("m5ea01zic0bg8dufo2eo").
		Build()
	if err != nil {
		t.Errorf("RebootNode Test Fail: %s", err.Error())
		return
	}
	_, err = aliEcs.RebootNode(rebootNode)
	if err != nil {
		t.Errorf("RebootNode Test Fail: %s", err)
		return
	}
	t.Logf("Ali node is rebooted successfully.")
}

func TestDeleteNode(t *testing.T) {
	var aliEcs ECS
	delete := map[string]interface{}{
		"InstanceId": "i-m5ea01zic0bg8dufo2eo",
	}
	_, err := aliEcs.DeleteNode(delete)
	if err != nil {
		t.Errorf("DeleteNode Test Fail")
		return
	}
	t.Logf("Ali node is deleted successfully.")
}
func TestDeleteNodeBuilder(t *testing.T) {
	var aliEcs ECS
	deleteNode, err := NewDeleteNodeBuilder().
		InstanceID("m5ea01zic0bg8dufo2eo").
		Build()
	if err != nil {
		t.Errorf("DeleteNode Test Fail: %s", err.Error())
		return
	}
	_, err = aliEcs.DeleteNode(deleteNode)
	if err != nil {
		t.Errorf("DeleteNode Test Fail: %s", err)
		return
	}
	t.Logf("Ali node is deleted successfully.")
}

func TestListNodeType(t *testing.T) {
	var aliEcs ECS
	resp, err := aliEcs.ListNodeType(nil)
	if err != nil {
		t.Errorf("ListNodeType Test Fail")
		return
	}
	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
	t.Logf("Ali node type is listed successfully.")
}
