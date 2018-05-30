package dynamodb

import (
	"bytes"
	"encoding/json"
	"fmt"
	awsauth "github.com/cloudlibz/gocloud/awsauth"
	"io/ioutil"
	"net/http"
)

//PrepareSignatureV4query creates PrepareSignatureV4 for request.
func (dynamodb *Dynamodb) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}, response map[string]interface{}) error {
	ECSEndpoint := "https://dynamodb." + params["Region"] + ".amazonaws.com"
	fmt.Println("ECSEndpoint : ", ECSEndpoint)
	service := "dynamodb"
	method := "POST"
	host := service + "." + params["Region"] + ".amazonaws.com"
	fmt.Println("host : ", host)
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
