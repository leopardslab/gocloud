package googlecloudfunctions

import "testing"
import "fmt"


func TestCallfunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

	callfunction := map[string]string{
		"name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
	}

	resp, err := googlecloudfunctions.Callfunction(callfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}


	response := resp.(map[string]interface{})

	fmt.Println(response["body"])

}

func TestDeletefunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

	deletefunction := map[string]string{
		"name": "projects/adept-comfort-202709/locations/us-central1/functions/function-6",
	}

	resp, err := googlecloudfunctions.Deletefunction(deletefunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])

}

func TestGetfunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

	getfunction := map[string]string{
		"name": "projects/adept-comfort-202709/locations/us-central1/functions/function-1",
	}

	resp, err := googlecloudfunctions.Getfunction(getfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}

func TestListfunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

	listfunction := map[string]string{
		"name":     "projects/adept-comfort-202709/locations/us-central1",
		"pageSize": "1",
	}

	resp, err := googlecloudfunctions.Listfunction(listfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])

}


func TestCreatefunction(t *testing.T) {

	var googlecloudfunctions Googlecloudfunctions

	httpsTrigger := make(map[string]string)
	httpsTrigger["URL"] = "https://us-central1-adept-comfort-202709.cloudfunctions.net/function-1"

	labels := make(map[string]string)
	labels["deployment-tool"] = "console-cloud"

	createfunction := map[string]interface{}{
		"Location":            "projects/adept-comfort-202709/locations/us-central1",
		"Name":                "projects/adept-comfort-202709/locations/us-central1/functions/function-3",
		"Status":              "ACTIVE",
		"HTTPSTrigger":        httpsTrigger,
		"EntryPoint":          "helloWorld",
		"Timeout":             "60s",
		"AvailableMemoryMb":   256,
		"ServiceAccountEmail": "adept-comfort-202709@appspot.gserviceaccount.com",
		"UpdateTime":          "2018-05-11T18:20:33Z",
		"Runtime":             "nodejs6",
		"SourceUploadURL":

		"https://storage.googleapis.com/gcf-upload-us-central1-f24bda97-6cd1-46cc-b37d-1f60eac4210a/f306c974-23cd-4294-ad46-8361d26488fa.zip?GoogleAccessId=service-126778294088@gcf-admin-robot.iam.gserviceaccount.com&Expires=1528366301&Signature=XcRaGgJXw71brq0eXcyQ1Dw8PVOT5Eu5E6xFzsa1uu87jd3irQr24pD48nm0PATjDD%2FttG2pRdyDTG0MpeUKUdtib8ohTWPcfM684LWvW0qQqv9q3k5DEa4DW74%2F9K2iMbda9P0oXvNEodVjToxezRGxY6BqAfNd5qqzA%2FbpL2LRlgOolg9d65vyyJy6ebewhYXlwEZHGhRHPuu3evnWLxTGxdLp4mbs1C%2FymzVsw4%2BS5Byl9GE0OWxjSNUt6YmL98l1QguLxefTERJdOU8qZvVDNnvtT4xMjIb1bju1zxPXEle7oQb2eM7sPoKYVYI5EheeIJpUvObFaklLr5qJ2Q%3D%3D",

		 	"VersionID":           "1",
		"Labels":              labels,
	}

	resp, err := googlecloudfunctions.Createfunction(createfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])

}
