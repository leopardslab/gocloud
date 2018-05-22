package alicontainer

import (
	"testing"
	"github.com/cloudlibz/gocloud/aliauth"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreateCluster(t *testing.T) {
	var aliContainer Alicontainer
	create := map[string]interface{}{
		"password":           "TestPwd124",
		"region_id":          "cn-beijing",
		"instance_type":      "ecs.n1.small",
		"name":               "my-test-cluster",
		"size":               1,
		"network_mode":       "vpc",
		"vpc_id":             "vpc-2ze578wokbm1ykyr7d8w6",
		"vswitch_id":         "vsw-2zeryg1zeofqj0u7o6buw",
		"subnet_cidr":        "172.28.1.0/24",
		"data_disk_category": "cloud_ssd",
		"data_disk_size":     40,
		"need_slb":           true,
		"ecs_image_id":       "centos_7_04_64_20G_alibase_201701015",
		"io_optimized":       "true",
		"release_eip_flag":   false,
	}
	_, err := aliContainer.Createcluster(create)
	if err != nil {
		t.Errorf("CreateCluster Test Fail: %s", err)
		return
	}
	t.Logf("Ali container is created successfully.")
}

func TestDeleteCluster(t *testing.T) {
	var aliContainer Alicontainer
	delete := map[string]interface{}{
		"region_id":  "cn-beijing",
		"cluster_id": "cf02b9dffa1fa45daac18cb436471ff2a",
	}
	_, err := aliContainer.Deletecluster(delete)
	if err != nil {
		t.Errorf("DeleteCluster Test Fail: %s", err)
		return
	}
	t.Logf("Ali container is deleted successfully.")
}

func TestRunTask(t *testing.T) {
	var aliContainer Alicontainer
	runTask := map[string]interface{}{
		"cluster_id":  "cc27bc9c9edbc4af5a0a369557e8da39d",
		"name":        "test",
		"description": "This is a test application",
		"template":    "web:\r\n  image: nginx",
		"version":     "1.0",
		"environment": map[string]string{"USER": "abc", "PWD": "password",},
	}
	resp, err := aliContainer.Runtask(runTask)
	if err != nil {
		t.Errorf("Runtask Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 201 {
		t.Errorf("Runtask Test Fail: %s", err)
		return
	}
	t.Logf("Ali container project is created successfully.")
}

func TestStartTask(t *testing.T) {
	var aliContainer Alicontainer
	startTask := map[string]interface{}{
		"cluster_id": "c7adbebc81e9647618a4348b8d92eac2f",
		"name":       "test",
	}
	_, err := aliContainer.Starttask(startTask)
	if err != nil {
		t.Errorf("StartTask Test Fail: %s", err)
		return
	}
	t.Logf("Ali container task is started successfully.")
}

func TestStopTask(t *testing.T) {
	var aliContainer Alicontainer
	stopTask := map[string]interface{}{
		"cluster_id": "c7adbebc81e9647618a4348b8d92eac2f",
		"name":       "test",
		"timeout":    20,
	}
	_, err := aliContainer.Stoptask(stopTask)
	if err != nil {
		t.Errorf("StopTask Test Fail: %s", err)
		return
	}
	t.Logf("Ali container task is stoped successfully.")
}
