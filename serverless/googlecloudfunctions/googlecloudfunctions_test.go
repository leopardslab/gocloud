package bigtable

import "testing"


func TestCallfunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

  callfunction := map[string]string{
    "name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
  }

	_, err := googlecloudfunctions.Callfunction(callfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}
}


func TestDeletefunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

  deletefunction := map[string]string{
    "name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
  }

	_, err := googlecloudfunctions.Deletefunction(deletefunction)

	if err != nil {
		t.Errorf("Test Fail")
	}
}


func TestGetfunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

  getfunction := map[string]string{
    "name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
  }

	_, err := googlecloudfunctions.Getfunction(getfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}
}


func TestListfunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

  listfunction := map[string]string{
    "name": "projects/adept-comfort-202709/locations/us-central1",
    "pageSize": "1",
  }


	_, err := googlecloudfunctions.Listfunction(listfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
