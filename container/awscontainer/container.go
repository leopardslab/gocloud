package awscontainer

import (
	"bytes"
	"encoding/json"
	"fmt"
	awsauth "github.com/scorelab/gocloud-v2/awsauth"
	"io/ioutil"
	"net/http"
)

type Ecscontainer struct {
}

func (ecscontainer *Ecscontainer) Createcontainer(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var clusterName, Region string
	for key, value := range param {
		switch key {
		case "clusterName":
			clusterNameV, _ := value.(string)
			clusterName = clusterNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}
	params := make(map[string]string)
	prepareCreatcontainerparams(params, clusterName, Region)

	Creatcontainerjsonmap := map[string]interface{}{
		"clusterName": params["clusterName"],
	}
	ecscontainer.PrepareSignatureV4query(params, Creatcontainerjsonmap)
	return
}

func (ecscontainer *Ecscontainer) Deletecontainer(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var clusterName, Region string
	for key, value := range param {
		switch key {
		case "clusterName":
			clusterNameV, _ := value.(string)
			clusterName = clusterNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}
	params := make(map[string]string)
	prepareDeletecontainer(params, clusterName, Region)

	Creatcontainerjsonmap := map[string]interface{}{
		"clusterName": params["clusterName"],
	}
	ecscontainer.PrepareSignatureV4query(params, Creatcontainerjsonmap)
	return
}

func prepareDeletecontainer(params map[string]string, clusterName string, Region string) {
	if clusterName != "" {
		params["clusterName"] = clusterName
	}
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.DeleteCluster"
}

func prepareCreatcontainerparams(params map[string]string, clusterName string, Region string) {
	if clusterName != "" {
		params["clusterName"] = clusterName
	}
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.CreateCluster"
}

func (ecscontainer *Ecscontainer) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}) {
	fmt.Println(paramsmap)
	ECSEndpoint := "https://ecs." + params["Region"] + ".amazonaws.com"
	service := "ecs"
	method := "POST"
	host := service + "." + params["Region"] + ".amazonaws.com"
	ContentType := "application/x-amz-json-1.1"
	signedheaders := "content-type;host;x-amz-date;x-amz-target"
	amztarget := params["amztarget"]

	requestparametersjson, _ := json.Marshal(paramsmap)
	requestparametersjsonstring := string(requestparametersjson)
	requestparametersjsonstringbyte := []byte(requestparametersjsonstring)
	fmt.Println("requestparametersjsonstring", requestparametersjsonstring)
	client := new(http.Client)
	request, _ := http.NewRequest("POST", ECSEndpoint, bytes.NewBuffer(requestparametersjsonstringbyte))
	request = awsauth.SignatureV4(request, requestparametersjsonstringbyte, amztarget, method, params["Region"], service, host, ContentType, signedheaders)
	resp, _ := client.Do(request)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("resp Body:", string(body))
}
