package bigtable

//Createbigtable json formation.
func Createbigtabledictnoaryconvert(option Createbigtable, Createbigtablejsonmap map[string]interface{}) {

	if option.tableId != ""{
		Createbigtablejsonmap["tableId"] = option.tableId
	}

	Createbigtablejsonmap["table"] = option.table
}
