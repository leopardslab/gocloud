package bigtable

import "testing"
import "fmt"

//create bigtable Createtable
func TestCreateTables(t *testing.T) {

	var bigtable Bigtable

	table := make(map[string]interface{})

	initialSplits := make([]map[string]interface{}, 0)

	createtables := map[string]interface{}{
		"parent":        "projects/adept-comfort-202709/instances/helloo",
		"tableId":       "tableId3",
		"table":         table,
		"initialSplits": initialSplits,
	}

	resp, err := bigtable.Createtables(createtables)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])

}


func TestDescribetables(t *testing.T) {

	var bigtable Bigtable

	describetables := map[string]string{
		"name": "projects/adept-comfort-202709/instances/helloo/tables/bokkkya",
	}

	resp, err := bigtable.Describetables(describetables)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])

}

func TestListTables(t *testing.T) {

	var bigtable Bigtable

	listtables := map[string]string{
		"parent":    "projects/adept-comfort-202709/instances/helloo",
		"view":      "NAME_ONLY",
		"pageToken": "",
	}

	resp, err := bigtable.Listtables(listtables)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}

/*
func TestDeletetables(t *testing.T) {

	var bigtable Bigtable

	deletetables := map[string]string{
		"name": "projects/adept-comfort-202709/instances/helloo/tables/bokkkya",
	}

	response, err := bigtable.Deletetables(deletetables)

	if err != nil {
		t.Errorf("Test Fail")
	}

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
}
*/
