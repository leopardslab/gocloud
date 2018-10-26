package lambda

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"
// import "fmt"

func init() {
	awsAuth.LoadConfig()
}

func TestDeleteFunction(t *testing.T) {

	var lambda Lambda

	deletefunction := map[string]interface{}{
		"Region":       "us-east-1",
		"FunctionName": "gocloud",
		"Qualifier":    "Qualifier",
	}

	_, err := lambda.DeleteFunction(deletefunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

}

func TestCreateFunction(t *testing.T) {

	var lambda Lambda

	code := map[string]interface{}{
		"ZipFile": "gocloud-ecce2422-f215-4d44-83a8-8361b457c5d9",
	}

	createfunction := map[string]interface{}{
		"Region":       "us-east-1",
		"FunctionName": "gocloud3",
		"Role":         "arn:aws:iam::478991680879:role/service-role/bokya",
		"Runtime":      "nodejs6.10",
		"Handler":      "index.handler",
		"Code":         code,
	}

	_, err := lambda.CreateFunction(createfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

}

func TestListFunction(t *testing.T) {

	var lambda Lambda

	listfunction := map[string]interface{}{
		"Region": "us-east-1",
	}

	_, err := lambda.ListFunction(listfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestGetFunction(t *testing.T) {

	var lambda Lambda

	getfunction := map[string]interface{}{
		"Region":       "us-east-1",
		"FunctionName": "gocloud",
	}

	_, err := lambda.GetFunction(getfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

} 
