package alicontainer

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"testing"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreateCluster(t *testing.T) {
	var aliContainer Alicontainer
	create := map[string]interface{}{
		"password":           "TestPwd123",
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
		"need_slb":           false,
		"ecs_image_id":       "centos_7_04_64_20G_alibase_201701015",
		"io_optimized":       "true",
		"release_eip_flag":   false,
	}
	_, err := aliContainer.CreateCluster(create)
	if err != nil {
		t.Errorf("CreateCluster Test Fail: %s", err)
		return
	}
	t.Logf("Ali container is created successfully.")
}

func TestCreateClusterBuilder(t *testing.T) {
	var aliContainer Alicontainer
	create, err := NewCreateClusterBuilder().
		Password("TestPwd123").
		RegionID("cn-beijing").
		InstanceType("ecs.n1.small").
		Name("my-test-cluster").
		Size(1).NetworkMode("vpc").
		VPCID("vpc-2ze578wokbm1ykyr7d8w6").
		VSwitchID("vsw-2zeryg1zeofqj0u7o6buw").
		SubnetCIDR("172.28.1.0/24").
		DataDiskCategory("cloud_ssd").
		DataDiskSize(40).NeedSLB(false).
		ECSImageID("centos_7_04_64_20G_alibase_201701015").
		IOOptimized("true").ReleaseEipFlag(false).
		Build()
	if err != nil {
		t.Errorf("CreateCluster Test Fail: %s", err)
		return
	}
	_, err = aliContainer.CreateCluster(create)
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
	_, err := aliContainer.DeleteCluster(delete)
	if err != nil {
		t.Errorf("DeleteCluster Test Fail: %s", err)
		return
	}
	t.Logf("Ali container is deleted successfully.")
}

func TestDeleteClusterBuilder(t *testing.T) {
	var aliContainer Alicontainer
	deleteCluster, err := NewDeleteClusterBuilder().
		RegionID("cn-beijing").
		ClusterID("cf02b9dffa1fa45daac18cb436471ff2a").
		Build()
	if err != nil {
		t.Errorf("DeleteCluster Test Fail: %s", err)
		return
	}
	_, err = aliContainer.DeleteCluster(deleteCluster)
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
		"environment": map[string]string{"USER": "abc", "PWD": "password"},
	}
	resp, err := aliContainer.RunTask(runTask)
	if err != nil {
		t.Errorf("RunTask Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 201 {
		t.Errorf("RunTask Test Fail: %s", err)
		return
	}
	t.Logf("Ali container project is created successfully.")
}

func TestRunTaskBuilder(t *testing.T) {
	var aliContainer Alicontainer
	runTask, err := NewRunTaskBuilder().
		ClusterID("cc27bc9c9edbc4af5a0a369557e8da39d").
		Name("test").
		Description("This is a test application").
		Template("web:\r\n  image: nginx").
		Version("1.0").
		Environment(map[string]string{"USER": "abc", "PWD": "password"}).
		Build()
	if err != nil {
		t.Errorf("RunTask Test Fail: %s", err)
		return
	}
	resp, err := aliContainer.RunTask(runTask)
	if err != nil {
		t.Errorf("RunTask Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 201 {
		t.Errorf("RunTask Test Fail: %s", err)
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
	_, err := aliContainer.StartTask(startTask)
	if err != nil {
		t.Errorf("StartTask Test Fail: %s", err)
		return
	}
	t.Logf("Ali container task is started successfully.")
}

func TestStartTaskBuilder(t *testing.T) {
	var aliContainer Alicontainer
	startTask, err := NewStartTaskBuilder().
		ClusterID("c7adbebc81e9647618a4348b8d92eac2f").
		Name("test").
		Build()
	if err != nil {
		t.Errorf("StartTask Test Fail: %s", err)
		return
	}
	_, err = aliContainer.StartTask(startTask)
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
	_, err := aliContainer.StopTask(stopTask)
	if err != nil {
		t.Errorf("StopTask Test Fail: %s", err)
		return
	}
	t.Logf("Ali container task is stoped successfully.")
}

func TestStopTaskBuilder(t *testing.T) {
	var aliContainer Alicontainer
	stopTask, err := NewStopTaskBuilder().
		ClusterID("c7adbebc81e9647618a4348b8d92eac2f").
		Name("test").
		Timeout(20).
		Build()
	if err != nil {
		t.Errorf("StopTask Test Fail: %s", err)
		return
	}
	_, err = aliContainer.StopTask(stopTask)
	if err != nil {
		t.Errorf("StopTask Test Fail: %s", err)
		return
	}
	t.Logf("Ali container task is stoped successfully.")
}
