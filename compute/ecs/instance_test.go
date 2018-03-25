package ecs

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"testing"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreatenode(t *testing.T) {
	var aliEcs ECS
	create := map[string]interface{}{
		"RegionId":        "cn-qingdao",
		"ImageId":         "centos_7_04_64_20G_alibase_201701015.vhd",
		"InstanceType":    "ecs.xn4.small",
		"SecurityGroupId": "sg-m5egbo9s5xb21kpu6nk2",
	}
	_, err := aliEcs.Createnode(create)
	if err != nil {
		t.Errorf("Createnode Test Fail")
		return
	}
	t.Logf("Ali node is created successfully.")
}

func TestStartnode(t *testing.T) {
	var aliEcs ECS
	start := map[string]interface{}{
		"InstanceId": "i-m5e3ee3z8wdy8ktdq591",
	}
	_, err := aliEcs.Startnode(start)
	if err != nil {
		t.Errorf("Startnode Test Fail")
		return
	}
	t.Logf("Ali node is started successfully.")
}

func TestStopnode(t *testing.T)  {
	var aliEcs ECS
	stop := map[string]interface{}{
		"InstanceId": "i-m5e3ee3z8wdy8ktdq591",
		"ForceStop":  false,
	}
	_, err := aliEcs.Stopnode(stop)
	if err != nil {
		t.Errorf("Stopnode Test Fail")
		return
	}
	t.Logf("Ali node is stoped successfully.")
}

func TestRebootnode(t *testing.T)  {
	var aliEcs ECS
	reboot := map[string]interface{}{
		"InstanceId": "i-m5e3ee3z8wdy8ktdq591",
		"ForceStop":  false,
	}
	_, err := aliEcs.Stopnode(reboot)
	if err != nil {
		t.Errorf("Rebootnode Test Fail")
		return
	}
	t.Logf("Ali node is rebooted successfully.")
}

func TestDeletenode(t *testing.T)  {
	var aliEcs ECS
	delete := map[string]interface{}{
		"InstanceId": "i-m5e3ee3z8wdy8ktdq591",
	}
	_, err := aliEcs.Deletenode(delete)
	if err != nil {
		t.Errorf("Deletenode Test Fail")
		return
	}
	t.Logf("Ali node is deleted successfully.")
}