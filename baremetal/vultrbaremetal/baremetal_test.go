package vultrbaremetal

import (
	"github.com/cloudlibz/gocloud/vultrauth"
	"testing"
)

func init() {
	vultrauth.LoadConfig()
}

func TestVultrBareMetal_CreateBareMetal(t *testing.T) {
	var vultrBareMetal VultrBareMetal
	create := map[string]interface{}{
		"DCID":        1,
		"METALPLANID": 100,
		"OSID":        127,
	}
	resp, err := vultrBareMetal.CreateBareMetal(create)
	if err != nil {
		t.Errorf("CreateBareMetal Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr BareMetal is created successfully.")
}

func TestVultrBareMetal_DeleteBareMetal(t *testing.T) {
	var vultrBareMetal VultrBareMetal
	deleteBareMetal := map[string]interface{}{
		"SUBID": 900000,
	}
	resp, err := vultrBareMetal.DeleteBareMetal(deleteBareMetal)
	if err != nil {
		t.Errorf("DeleteBareMetal Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr BareMetal is deleted successfully.")
}
