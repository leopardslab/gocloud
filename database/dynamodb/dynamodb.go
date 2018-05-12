package dynamodb

import (
	"bytes"
	"encoding/json"
	awsauth "github.com/cloudlibz/gocloud/awsauth"
	"io/ioutil"
	"net/http"
	"fmt"
)

func (dynamodb *Dynamodb) Listtables(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})

	var ExclusiveStartTableName, Limit, Region string

	for key, value := range param {
		switch key {
			case "ExclusiveStartTableName":
				ExclusiveStartTableNameV, _ := value.(string)
				ExclusiveStartTableName = ExclusiveStartTableNameV

			case "Limit":
				LimitV, _ := value.(string)
				Limit = LimitV

			case "Region":
				RegionV, _ := value.(string)
				Region = RegionV
		}
	}

	params := make(map[string]string)

	preparelisttables(params, ExclusiveStartTableName, Limit, Region)

	listtablesjsonmap := map[string]interface{}{
	//	"ExclusiveStartTableName":ExclusiveStartTableName ,
	//	"Limit": Limit,
	}

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, listtablesjsonmap, response)
	resp = response
	return resp, err
}



func (dynamodb *Dynamodb) Deletetables(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})

	var TableName, Region string

	for key, value := range param {
		switch key {
			case "TableName":
				TableNameV, _ := value.(string)
				TableName = TableNameV

			case "Region":
				RegionV, _ := value.(string)
				Region = RegionV
		}
	}

	params := make(map[string]string)

	preparedeletetables(params, TableName, Region)

	deletetablesjsonmap := map[string]interface{}{
		"TableName": TableName ,
	}

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, deletetablesjsonmap, response)
	resp = response
	return resp, err
}

func (dynamodb *Dynamodb) Createtables(request interface{}) (resp interface{}, err error) {
	return resp, err
}


func (dynamodb *Dynamodb)Describetables(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})

	var TableName, Region string

	for key, value := range param {
		switch key {
			case "TableName":
				TableNameV, _ := value.(string)
				TableName = TableNameV

			case "Region":
				RegionV, _ := value.(string)
				Region = RegionV
		}
	}

	params := make(map[string]string)

	preparedescribetables(params, TableName, Region)

	deletetablesjsonmap := map[string]interface{}{
		"TableName": TableName ,
	}

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, deletetablesjsonmap, response)
	resp = response
	return resp, err
}


func preparedescribetables(params map[string]string, TableName string,Region string) {

	if TableName != "" {
		params["TableName"] = TableName
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "DynamoDB_20120810.DescribeTable"
}

func preparedeletetables(params map[string]string, TableName string,Region string) {

	if TableName != "" {
		params["TableName"] = TableName
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "DynamoDB_20120810.DeleteTable"
}



func preparelisttables(params map[string]string, ExclusiveStartTableName string,Limit string,  Region string) {
	if ExclusiveStartTableName != "" {
		params["ExclusiveStartTableName"] = ExclusiveStartTableName
	}

	if Limit != "" {
		params["Limit"] = Limit
	}

	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "DynamoDB_20120810.ListTables"
}



//PrepareSignatureV4query creates PrepareSignatureV4 for request.
func (dynamodb *Dynamodb) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}, response map[string]interface{}) error {
	ECSEndpoint := "https://dynamodb." + params["Region"] + ".amazonaws.com"
	fmt.Println("ECSEndpoint : ",ECSEndpoint)
	service := "dynamodb"
	method := "POST"
	host := service + "." + params["Region"] + ".amazonaws.com"
	fmt.Println("host : ",host)
	ContentType := "application/x-amz-json-1.0"
	signedheaders := "content-type;host;x-amz-date;x-amz-target"
	amztarget := params["amztarget"]

	requestparametersjson, _ := json.Marshal(paramsmap)
	requestparametersjsonstring := string(requestparametersjson)
	requestparametersjsonstringbyte := []byte(requestparametersjsonstring)
	client := new(http.Client)
	request, _ := http.NewRequest("POST", ECSEndpoint, bytes.NewBuffer(requestparametersjsonstringbyte))
	request = awsauth.SignatureV4(request, requestparametersjsonstringbyte, amztarget, method, params["Region"], service, host, ContentType, signedheaders)
	resp, err := client.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	response["body"] = string(body)
	response["status"] = resp.StatusCode
	return err
}