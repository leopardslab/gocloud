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
		"DCID":        40,
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

func TestCreateBareMetalBuilder(t *testing.T) {
	var vultrBareMetal VultrBareMetal
	createBareMetal, err := NewCreateBareMetalBuilder().
		DCID(1).
		METALPLANID(100).
		OSID(127).
		Build()
	if err != nil {
		t.Errorf("CreateBareMetal Test Fail: %s", err)
		return
	}
	resp, err := vultrBareMetal.CreateBareMetal(createBareMetal)
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

func TestVultrBareMetal_RebootBareMetal(t *testing.T) {
	var vultrBareMetal VultrBareMetal
	rebootBareMetal := map[string]interface{}{
		"SUBID": 900000,
	}
	resp, err := vultrBareMetal.RebootBareMetal(rebootBareMetal)
	if err != nil {
		t.Errorf("RebootBareMetal Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr BareMetal is rebooted successfully.")
}

func TestVultrBareMetal_ReinstallBareMetal(t *testing.T) {
	var vultrBareMetal VultrBareMetal
	reinstallBareMetal := map[string]interface{}{
		"SUBID": 900000,
	}
	resp, err := vultrBareMetal.ReinstallBareMetal(reinstallBareMetal)
	if err != nil {
		t.Errorf("ReinstallBareMetal Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr BareMetal is reinstalled successfully.")
}

func TestVultrBareMetal_HaltBareMetal(t *testing.T) {
	var vultrBareMetal VultrBareMetal
	haltBareMetal := map[string]interface{}{
		"SUBID": 900000,
	}
	resp, err := vultrBareMetal.HaltBareMetal(haltBareMetal)
	if err != nil {
		t.Errorf("HaltBareMetal Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("Vultr BareMetal is halted successfully.")
}

func TestVultrBareMetal_ListBareMetal(t *testing.T) {
	var vultrBareMetal VultrBareMetal
	resp, err := vultrBareMetal.ListBareMetal(nil)
	if err != nil {
		t.Errorf("ListBareMetal Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("response body: %s\n", response["body"])
	t.Logf("Vultr BareMetal is listed successfully.")
}

func TestParseCreateBareMetalResp(t *testing.T) {
	var vultrBareMetal VultrBareMetal
	create := map[string]interface{}{
		"DCID":        40,
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
	createBareMetalResp, err := ParseCreateBareMetalResp(response["body"])
	if err != nil {
		t.Errorf("CreateBareMetal Test Fail: %s", err)
		return
	}
	t.Log(createBareMetalResp.SUBID)
}

func TestParseListBareMetalResp(t *testing.T) {
	var vultrBareMetal VultrBareMetal
	resp, err := vultrBareMetal.ListBareMetal(nil)
	if err != nil {
		t.Errorf("ListBareMetal Test Fail: %s", err)
		return
	}
	response := resp.(map[string]interface{})
	if response["status"] != 200 {
		t.Errorf("status code: %d\n response body: %s\n", response["status"], response["body"])
		return
	}
	t.Logf("response body: %s\n", response["body"])
	listBareMetalResp, err := ParseListBareMetalResp(response["body"])
	if err != nil {
		t.Errorf("CreateNode Test Fail: %s", err)
		return
	}
	t.Logf("Vultr BareMetal is listed successfully.")
	for _, bareMetal := range listBareMetalResp {
		t.Logf("%+v", bareMetal)
	}
}
