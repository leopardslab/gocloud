package googlecloudfunctions

import "testing"
import "fmt"

func TestCallFunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

	callfunction := map[string]string{
		"name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
	}

	_, err := googlecloudfunctions.CallFunction(callfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteFunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

	deletefunction := map[string]string{
		"name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
	}

	_, err := googlecloudfunctions.DeleteFunction(deletefunction)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestGetFunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

	getfunction := map[string]string{
		"name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
	}

	_, err := googlecloudfunctions.GetFunction(getfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListFunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

	listfunction := map[string]string{
		"name":     "projects/adept-comfort-202709/locations/us-central1",
		"pageSize": "1",
	}

	_, err := googlecloudfunctions.ListFunction(listfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestCreateFunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

	httpsTrigger := make(map[string]string)
	httpsTrigger["URL"] = "https://us-central1-adept-comfort-202709.cloudfunctions.net/function-1"

	labels := make(map[string]string)
	labels["deployment-tool"] = "console-cloud"

	createfunction := map[string]interface{}{
		"Location":            "projects/adept-comfort-202709/locations/us-central1",
		"Name":                "projects/adept-comfort-202709/locations/us-central1/functions/function-2",
		"Status":              "ACTIVE",
		"HTTPSTrigger":        httpsTrigger,
		"EntryPoint":          "helloWorld",
		"Timeout":             "60s",
		"AvailableMemoryMb":   256,
		"ServiceAccountEmail": "adept-comfort-202709@appspot.gserviceaccount.com",
		"UpdateTime":          "2018-05-11T18:20:33Z",
		"Runtime":             "nodejs6",
		"SourceUploadURL":     "https://storage.googleapis.com/gcf-upload-us-central1-f24bda97-6cd1-46cc-b37d-1f60eac4210a/8548b011-9626-42c1-86ed-6190892b328e.zip?GoogleAccessId=126778294088@cloudservices.gserviceaccount.com&Expires=1526064618&Signature=nB%2FI6cwIap0DF5T0Uo9eYCnlmi3HLqvoRW4MfodzVI%2FXuC7HU%2BE9SwduVQKYeTRddo5iFNdm4VDmBu4A4fGQvZ5PaCuoKG4i7jZXRJgq1B4NIpocaFnHmY6ZWaCS0Av%2Bus29FHs2nTYIqp9zHWHHORSQC%2BPF8GP2mRToDOShpodkQFkxP6wsXUnkk8tDUf5mvTRkeqtgf0rX0huidbEVl7ZtGkcQiusDcS9Nhe3dwqOdsJ7xs2khl2D%2FOmch6jgrZ11MtXum3G5XnFLqMYupS0pvB%2BQiy7g7eLfIw%2BdvtRTEuEFyxWP49lCHUG8wWad0hNEVf29oAHS2x%2B4Q%2FaIGbA%3D%3D",
		"VersionID":           "1",
		"Labels":              labels,
	}

	_, err := googlecloudfunctions.CreateFunction(createfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
