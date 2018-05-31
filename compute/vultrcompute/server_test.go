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
