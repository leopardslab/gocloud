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
		t.Errorf("CreateCluster Test Fail: %s", err)
		return
	}
	t.Logf("Ali container is created successfully.")
}
