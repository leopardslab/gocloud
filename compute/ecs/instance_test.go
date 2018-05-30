package ecs

import (
	"fmt"
	"github.com/cloudlibz/gocloud/aliauth"
	"testing"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreatenode(t *testing.T) {
	var aliEcs ECS
	create := map[string]interface{}{
		"RegionId":            "cn-qingdao",
		"ImageId":             "centos_7_04_64_20G_alibase_201701015.vhd",
		"InstanceType":        "ecs.xn4.small",
		"SecurityGroupId":     "sg-m5egbo9s5xb21kpu6nk2",
		"DataDisk.1.Size":     "20",
		"DataDisk.1.DiskName": "testName",
	}
	_, err := aliEcs.Createnode(create)
	if err != nil {
		t.Errorf("Createnode Test Fail: %s", err)
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
		t.Errorf("Createnode Test Fail: %s", err.Error())
		return
	}
	_, err = aliEcs.Createnode(createNode)
	if err != nil {
		t.Errorf("Createnode Test Fail: %s", err)
		return
	}
	t.Logf("Ali node is created successfully.")
}

func TestStartnode(t *testing.T) {
	var aliEcs ECS
	start := map[string]interface{}{
		"InstanceId": "i-m5ea01zic0bg8dufo2eo",
	}
	_, err := aliEcs.Startnode(start)
	if err != nil {
		t.Errorf("Startnode Test Fail")
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
	_, err = aliEcs.Startnode(startNode)
	if err != nil {
		t.Errorf("StartNode Test Fail: %s", err)
		return
	}
	t.Logf("Ali node is started successfully.")
}

func TestStopnode(t *testing.T) {
	var aliEcs ECS
	stop := map[string]interface{}{
		"InstanceId": "i-m5ea01zic0bg8dufo2eo",
		"ForceStop":  false,
	}
	_, err := aliEcs.Stopnode(stop)
	if err != nil {
		t.Errorf("Stopnode Test Fail")
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
		t.Errorf("Stopnode Test Fail: %s", err.Error())
		return
	}
	_, err = aliEcs.Stopnode(stopNode)
	if err != nil {
		t.Errorf("Stopnode Test Fail: %s", err)
		return
	}
	t.Logf("Ali node is stoped successfully.")
}

func TestRebootnode(t *testing.T) {
	var aliEcs ECS
	reboot := map[string]interface{}{
		"InstanceId": "i-m5ea01zic0bg8dufo2eo",
		"ForceStop":  false,
	}
	_, err := aliEcs.Rebootnode(reboot)
	if err != nil {
		t.Errorf("Rebootnode Test Fail")
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
		t.Errorf("Rebootnode Test Fail: %s", err.Error())
		return
	}
	_, err = aliEcs.Rebootnode(rebootNode)
	if err != nil {
		t.Errorf("Rebootnode Test Fail: %s", err)
		return
	}
	t.Logf("Ali node is rebooted successfully.")
}

func TestDeletenode(t *testing.T) {
	var aliEcs ECS
	delete := map[string]interface{}{
		"InstanceId": "i-m5ea01zic0bg8dufo2eo",
	}
	_, err := aliEcs.Deletenode(delete)
	if err != nil {
		t.Errorf("Deletenode Test Fail")
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
		t.Errorf("Deletenode Test Fail: %s", err.Error())
		return
	}
	_, err = aliEcs.Deletenode(deleteNode)
	if err != nil {
		t.Errorf("Deletenode Test Fail: %s", err)
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
