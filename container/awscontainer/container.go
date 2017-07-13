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

func (ecscontainer *Ecscontainer) Creatcontainer(request interface{}) (resp interface{}, err error) {

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
	prepreCreatcontainerparams(params, clusterName, Region)
	ecscontainer.PrepareSignatureV4query(params)
	return
}

func prepreCreatcontainerparams(params map[string]string, clusterName string, Region string) {
	if clusterName != "" {
		params["clusterName"] = clusterName
	}
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.CreateCluster"
}

func (ecscontainer *Ecscontainer) PrepareSignatureV4query(params map[string]string) {

	ECSEndpoint := "https://ecs." + params["Region"] + ".amazonaws.com"
	service := "ecs"
	method := "POST"
	host := service + "." + params["Region"] + ".amazonaws.com"
	ContentType := "application/x-amz-json-1.1"
	signedheaders := "content-type;host;x-amz-date;x-amz-target"
	amztarget := params["amztarget"]

	Creatcontainerjsonmap := map[string]interface{}{
		"clusterName": params["clusterName"],
	}

	Creatcontainerjson, _ := json.Marshal(Creatcontainerjsonmap)
	Creatcontainerstring := string(Creatcontainerjson)
	var Creatcontainerstringbyte = []byte(Creatcontainerstring)

	client := new(http.Client)
	Creatcontainerrequest, _ := http.NewRequest("POST", ECSEndpoint, bytes.NewBuffer(Creatcontainerstringbyte))
	request := awsauth.SignatureV4(Creatcontainerrequest, Creatcontainerstringbyte, amztarget, method, params["Region"], service, host, ContentType, signedheaders)
	Creatcontainerresp, _ := client.Do(request)
	defer Creatcontainerresp.Body.Close()
	fmt.Println(Creatcontainerresp.Status)
	body, _ := ioutil.ReadAll(Creatcontainerresp.Body)
	fmt.Println("Creatcontainerresp Body:", string(body))
}
