package rackspacecompute

import "testing"

func TestCreateNode(t *testing.T) {

	var rackspacecompute Rackspacecompute

	request := map[string]interface{}{
		"Name":       "testcloudserver",
		"ImageRef":   "3afe97b2-26dc-49c5-a2cc-a2fc8d80c001",
		"FlavourRef": "2",
		"Region":     "DFW",
	}
	_, err := rackspacecompute.CreateNode(request)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
