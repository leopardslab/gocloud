package lambda

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"
import "fmt"

func init() {
	awsAuth.LoadConfig()
}

/*
func TestDeleteFunction(t *testing.T) {

	var lambda Lambda

	deletedunction := map[string]interface{}{
		"Region":     "us-east-1",
		"FunctionName": "gocloud",
    "Qualifier" : "Qualifier",
	}

	resp, err := lambda.DeleteFunction(deletedunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}


*/

/*
func TestCreateFunction(t *testing.T) {

	var lambda Lambda

	deletedunction := map[string]interface{}{
		"Region":       "us-east-1",
		"FunctionName": "gocloud",
		"Qualifier":    "Qualifier",
	}

	resp, err := lambda.CreateFunction(deletedunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}
*/



/*
func TestListFunction(t *testing.T) {

	var lambda Lambda

	listfunction := map[string]interface{}{
		"Region":     "us-east-1",
	}


	resp, err := lambda.ListFunction(listfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}

*/

func TestGetFunction(t *testing.T) {

	var lambda Lambda

	getfunction := map[string]interface{}{
		"Region":     "us-east-1",
		"FunctionName" : "gocloud",
	}


	resp, err := lambda.GetFunction(getfunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}


func TestDeleteFunction(t *testing.T) {

	var lambda Lambda

	deletefunction	 := map[string]interface{}{
		"Region":     "us-east-1",
		"FunctionName" : "gocloud2",
	}


	resp, err := lambda.DeleteFunction(deletefunction)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}
