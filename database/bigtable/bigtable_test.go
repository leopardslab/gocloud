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
		"tableId":       "tableId",
		"table":         table,
		"initialSplits": initialSplits,
	}

	_, err := bigtable.CreateTables(createtables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDescribeTables(t *testing.T) {

	var bigtable Bigtable

	describetables := map[string]string{
		"name": "projects/adept-comfort-202709/instances/helloo/tables/bokkkya",
	}

	_, err := bigtable.DescribeTables(describetables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListTables(t *testing.T) {

	var bigtable Bigtable

	listtables := map[string]string{
		"parent":    "projects/adept-comfort-202709/instances/helloo",
		"view":      "NAME_ONLY",
		"pageToken": "",
	}

	_, err := bigtable.ListTables(listtables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteTables(t *testing.T) {

	var bigtable Bigtable

	deletetables := map[string]string{
		"name": "projects/adept-comfort-202709/instances/helloo/tables/bokkkya",
	}

	_, err := bigtable.DeleteTables(deletetables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
