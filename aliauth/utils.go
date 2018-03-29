package aliauth

import "reflect"

// PutStructToMap puts key and value of struct into map[string]interface{}
// if value is string and not empty -> put
// if value is bool                 -> put   NOTE: the default value of origin Ali API 's parameter must be false, if not ,do not use this function
// if value is int and not 0        -> put   NOTE: the optional values of origin Ali API 's parameter must dose not include 0, if not ,do not use this function
func PutStructToMap(i interface{}) map[string]interface{} {
	params := make(map[string]interface{})
	e := reflect.ValueOf(i).Elem()
	typeOfOptions := e.Type()
	for i := 0; i < e.NumField(); i++ {
		switch e.Field(i).Type().String() {
		case "string":
			if e.Field(i).Interface() != "" {
				params[typeOfOptions.Field(i).Name] = e.Field(i).Interface()
			}
		case "bool":
			params[typeOfOptions.Field(i).Name] = e.Field(i).Interface()
		case "int":
			if e.Field(i).Interface() != 0 {
				params[typeOfOptions.Field(i).Name] = e.Field(i).Interface()
			}
		}
	}
	return params
}
