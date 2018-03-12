package aliecs

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"testing"
	"fmt"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreatenode(t *testing.T) {
	var aliEcs ECS
	create := map[string]interface{}{
		"RegionId":        "cn-qingdao",
		"ImageId":         "centos_7_04_64_20G_alibase_201701015.vhd",
		"InstanceType":    "ecs.t1.small",
		"SecurityGroupId": "sg-m5egbo9s5xb21kpu6nk2",
	}
	resp, err := aliEcs.Createnode(create)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := resp.(map[string]interface{})
	fmt.Println(response["body"])
}
