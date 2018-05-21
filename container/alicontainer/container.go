package alicontainer

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"strconv"
	"net/http"
	"encoding/json"
	"fmt"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"io"
	"bytes"
)

//TODO
func (alicontainer *Alicontainer) Createcluster(request interface{}) (resp interface{}, err error) {
	var options CreateCluster

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	var regionID string

	for key, value := range param {
		switch key {
		case "region_id":
			regionID, _ = value.(string)
		case "name":
			name, _ := value.(string)
			options.Name = name
		case "size":
			switch value.(type) {
			case int:
				options.Size = int64(value.(int))
			case int64:
				options.Size = value.(int64)
			case string:
				options.Size, _ = strconv.ParseInt(value.(string), 10, 64)
			}
		case "instance_type":
			instanceType, _ := value.(string)
			options.InstanceType = instanceType
		case "network_mode":
			networkMode, _ := value.(string)
			options.NetworkMode = networkMode
		case "subnet_cidr":
			subnetCIDR, _ := value.(string)
			options.SubnetCIDR = subnetCIDR
		case "vpc_id":
			vpcID, _ := value.(string)
			options.VPCID = vpcID
		case "vswitch_id":
			vSwitchID, _ := value.(string)
			options.VSwitchID = vSwitchID
		case "password":
			password, _ := value.(string)
			options.Password = password
		case "data_disk_category":
			dataDiskCategory, _ := value.(string)
			options.DataDiskCategory = dataDiskCategory
		case "data_disk_size":
			switch value.(type) {
			case int:
				options.Size = int64(value.(int))
			case int64:
				options.Size = value.(int64)
			case string:
				options.DataDiskSize, _ = strconv.ParseInt(value.(string), 10, 64)
			}
		case "ecs_image_id":
			ecsImageID, _ := value.(string)
			options.ECSImageID = ecsImageID
		case "io_optimized":
			IOOptimized, _ := value.(string)
			options.IOOptimized = IOOptimized
		case "need_slb":
			switch value.(type) {
			case bool:
				options.NeedSLB = value.(bool)
			case string:
				options.NeedSLB = value.(string) == "true" || value.(string) == "True"
			}
		case "release_eip_flag":
			switch value.(type) {
			case bool:
				options.ReleaseEipFlag = value.(bool)
			case string:
				options.ReleaseEipFlag = value.(string) == "true" || value.(string) == "True"
			}
		}
	}

	response := make(map[string]interface{})
	err = aliauth.ContainerSignAndDoRequest(regionID, http.MethodPost, "/clusters", nil, options, response)
	resp = response
	return resp, err
}

//TODO
func (alicontainer *Alicontainer) Deletecluster(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	var regionID string
	var clusterID string

	for key, value := range param {
		switch key {
		case "region_id":
			regionID, _ = value.(string)
		case "cluster_id":
			clusterID, _ = value.(string)
		}
	}

	response := make(map[string]interface{})
	err = aliauth.ContainerSignAndDoRequest(regionID, http.MethodDelete, "/clusters/"+clusterID, nil, nil, response)
	resp = response
	return resp, err
}

//TODO
func (alicontainer *Alicontainer) Createservice(request interface{}) (resp interface{}, err error) {
	return resp, err
}

//TODO
func (alicontainer *Alicontainer) Runtask(request interface{}) (resp interface{}, err error) {
	var options RunTask

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	var clusterID string

	for key, value := range param {
		switch key {
		case "cluster_id":
			clusterID = value.(string)
		case "name":
			options.Name, _ = value.(string)
		case "description":
			options.Description, _ = value.(string)
		case "template":
			options.Template, _ = value.(string)
		case "version":
			options.Version, _ = value.(string)
		case "environment":
			options.Environment, _ = value.(map[string]string)
		case "latest_image":
			switch value.(type) {
			case bool:
				options.LatestImage = value.(bool)
			case string:
				options.LatestImage = value.(string) == "true" || value.(string) == "True"
			}
		}
	}

	clusterCerts, err := getClusterCerts(clusterID)
	if err != nil {
		return
	}
	fmt.Println(clusterCerts)

	masterURL, err := getClusterMasterURL(clusterID)
	if err != nil {
		return
	}
	fmt.Println(masterURL)

	certs, err := tls.X509KeyPair([]byte(clusterCerts.Cert), []byte(clusterCerts.Key))
	if err != nil {
		return
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(clusterCerts.CA))

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Certificates:       []tls.Certificate{certs},
				ClientCAs:          caCertPool,
				ClientAuth:         tls.RequireAndVerifyClientCert,
			},
		},
	}

	var reqBody []byte

	// generate request body and Content-MD5
	reqBody, err = json.Marshal(options)
	if err != nil {
		return
	}

	// generate request url
	requestURL := masterURL + "/projects/"
	fmt.Println(requestURL)

	var bodyReader io.Reader
	if reqBody != nil {
		bodyReader = bytes.NewReader(reqBody)
	}
	httpReq, err := http.NewRequest(http.MethodGet, requestURL, bodyReader)
	if err != nil {
		return
	}

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	response := make(map[string]interface{})

	response["body"] = string(body)
	response["status"] = httpResp.StatusCode

	resp = response
	return resp, err
}

//TODO
func (alicontainer *Alicontainer) Starttask(request interface{}) (resp interface{}, err error) {
	return
}

//TODO
func (alicontainer *Alicontainer) Deleteservice(request interface{}) (resp interface{}, err error) {
	return resp, err
}

//TODO
func (alicontainer *Alicontainer) Stoptask(request interface{}) (resp interface{}, err error) {
	return resp, err
}

func getClusterCerts(clusterID string) (certs clusterCerts, err error) {
	response := make(map[string]interface{})
	err = aliauth.ContainerSignAndDoRequest("", http.MethodGet, "/clusters/"+clusterID+"/certs", nil, nil, response)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(response["body"].(string)), &certs)
	return
}

func getClusterMasterURL(clusterID string) (masterURL string, err error) {
	response := make(map[string]interface{})
	err = aliauth.ContainerSignAndDoRequest("", http.MethodGet, "/clusters/"+clusterID, nil, nil, response)
	if err != nil {
		return
	}
	var clusterInfo cluster
	err = json.Unmarshal([]byte(response["body"].(string)), &clusterInfo)
	if err != nil {
		return
	}
	masterURL = clusterInfo.MasterURL
	return
}
