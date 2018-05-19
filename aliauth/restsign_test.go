package aliauth

import (
	"testing"
)

func TestRESTSign(t *testing.T) {
	reqBody := `{"password": "Just$test","instance_type": "ecs.m2.medium","name": "my-test-cluster-97082734","size": 1,"network_mode": "classic","data_disk_category": "cloud","data_disk_size": 10,"ecs_image_id": "m-253llee3l"}`

	contentLength := len(reqBody)
	if contentLength != 210 {
		t.Errorf("Content-Length: got %d, expected %d", contentLength, 210)
	}

	contentMD5 := createContentMD5([]byte(reqBody))
	if contentMD5 != "6U4ALMkKSj0PYbeQSHqgmA==" {
		t.Errorf("Content-MD5: got %s, expected %s", contentMD5, "6U4ALMkKSj0PYbeQSHqgmA==")
	}
}
