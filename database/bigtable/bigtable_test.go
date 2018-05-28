package bigtable

import "testing"
import "fmt"

func TestCreatetables(t *testing.T) {

	var bigtable Bigtable

	table := make(map[string]interface{})

	initialSplits := make([]map[string]interface{},0)

	createtables := map[string]interface{}{
		"parent": "projects/adept-comfort-202709/instances/helloo",
		"tableId" :"tableId",
		"table"  : table,
		"initialSplits" : initialSplits,
	}

	resp, err := bigtable.Createtables(createtables)

	response := resp.(map[string]interface{})

	fmt.Println(response["body"])


	if err != nil {
		t.Errorf("Test Fail")
	}
}


/*
func TestDescribetables(t *testing.T) {

	var bigtable Bigtable

	describetables := map[string]string{
		"name": "projects/adept-comfort-202709/instances/helloo/tables/bokkkya",
	}

	_, err := bigtable.Describetables(describetables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListtables(t *testing.T) {

	var bigtable Bigtable

	listtables := map[string]string{
		"parent":    "projects/adept-comfort-202709/instances/helloo",
		"view":      "NAME_ONLY",
		"pageToken": "",
	}

	_, err := bigtable.Listtables(listtables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeletetables(t *testing.T) {

	var bigtable Bigtable

	deletetables := map[string]string{
		"name": "projects/adept-comfort-202709/instances/helloo/tables/bokkkya",
	}

	_, err := bigtable.Deletetables(deletetables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
*/
