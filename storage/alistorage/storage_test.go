package alistorage

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"testing"
)

func init() {
	aliauth.LoadConfig()
}

func TestCreatedisk(t *testing.T) {
	var alistorage Alistorage
	createdisk := map[string]interface{}{
		"RegionId":    "cn-qingdao",
		"ZoneId":      "cn-qingdao-b",
		"Size":        20,
		"DiskName":    "ThisIsDiskName",
		"Description": "ThisIsDescription",
	}
	_, err := alistorage.Createdisk(createdisk)
	if err != nil {
		t.Errorf("Createdisk Test Fail")
		return
	}
	t.Logf("Ali disk is created successfully.")
}
