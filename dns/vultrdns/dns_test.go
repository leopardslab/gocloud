package vultrdns

import (
	"github.com/cloudlibz/gocloud/vultrauth"
	"testing"
)

func init() {
	vultrauth.LoadConfig()
}

func TestVultrDNS_Createdns(t *testing.T) {
	var vultrDNS VultrDNS
	createDNS := map[string]interface{}{
		"domain": "oddcn.cn",
		"name":   "gocloudtest",
		"type":   "A",
		"data":   "192.0.2.1",
	}
	resp, err := vultrDNS.Createdns(createDNS)
	if err != nil {
		t.Errorf("Createdns Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr DNS record is created successfully.")
}
