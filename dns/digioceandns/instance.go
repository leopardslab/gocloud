package digioceandns

import (
	"bytes"
	"encoding/json"
	"fmt"
	digioceanAuth "github.com/cloudlibz/gocloud/digioceanauth"
	"io/ioutil"
	"net/http"
)

// dnsBasePath is the endpoint URL for digitalocean API.
const dnsBasePath = "https://api.digitalocean.com/v2/domains"

// Createdns function creates a new DNS record.
func (digioceandns *Digioceandns) Createdns(request interface{}) (resp interface{}, err error) {

	var dnsInstance Digioceandns                                     // Initialize LoadBalancer struct
	var domainName string                                            // To store domain name
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken // Fetch the DigiOceanAccessToken

	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	for key, value := range param {

		switch key {

		case "DomainName":
			domainNameValue, _ := value.(string)
			domainName = domainNameValue

		case "Type":
			typeDNS, _ := value.(string)
			dnsInstance.Type = typeDNS

		case "Name":
			name, _ := value.(string)
			dnsInstance.Name = name

		case "Data":
			data, _ := value.(string)
			dnsInstance.Data = data

		case "Priority":
			priority, _ := value.(int)
			dnsInstance.Priority = priority

		case "Port":
			port, _ := value.(int)
			dnsInstance.Port = port

		case "TimeToLive":
			ttl, _ := value.(int)
			dnsInstance.TimeToLive = ttl

		case "Weight":
			weight, _ := value.(int)
			dnsInstance.Weight = weight

		case "Flags":
			flags, _ := value.(int)
			dnsInstance.Flags = flags

		case "Tag":
			tag, _ := value.(string)
			dnsInstance.Tag = tag

		} // Closes switch-case

	} // Closes for loop

	dnsInstanceJSON, _ := json.Marshal(dnsInstance)
	dnsInstanceJSONString := string(dnsInstanceJSON)
	var dnsInstanceJSONStringbyte = []byte(dnsInstanceJSONString)

	url := dnsBasePath + "/" + domainName + "/records"

	createDNSReq, err := http.NewRequest("POST", url, bytes.NewBuffer(dnsInstanceJSONStringbyte))
	if err != nil {
		fmt.Println(err)
	}
	createDNSReq.Header.Set("Content-Type", "application/json")
	createDNSReq.Header.Set("Authorization", "Bearer "+DigiOceanAccessToken)

	createDNSResp, err := http.DefaultClient.Do(createDNSReq)
	if err != nil {
		fmt.Println(err)
	}

	defer createDNSResp.Body.Close()

	responseBody, err := ioutil.ReadAll(createDNSResp.Body)
	createDNSResponse := make(map[string]interface{})
	createDNSResponse["status"] = createDNSResp.StatusCode
	createDNSResponse["body"] = string(responseBody)
	resp = createDNSResponse

	return resp, err
}

// Deletedns function deletes a DNS record.
func (digioceandns *Digioceandns) Deletedns(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := dnsBasePath + "/" + options["DomainName"] + "/records/" + options["RecordID"]
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken // Fetch the DigiOceanAccessToken

	deleteDNSReq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	deleteDNSReq.Header.Set("Content-Type", "application/json")
	deleteDNSReq.Header.Set("Authorization", "Bearer "+DigiOceanAccessToken)

	deleteDNSResp, err := http.DefaultClient.Do(deleteDNSReq)
	if err != nil {
		fmt.Println(err)
	}

	defer deleteDNSResp.Body.Close()

	responseBody, err := ioutil.ReadAll(deleteDNSResp.Body)
	deleteDNSResponse := make(map[string]interface{})
	deleteDNSResponse["status"] = deleteDNSResp.StatusCode
	deleteDNSResponse["body"] = string(responseBody)
	resp = deleteDNSResponse

	return resp, err
}

// Listdns function lists DNS records.
func (digioceandns *Digioceandns) Listdns(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := dnsBasePath + "/" + options["DomainName"] + "/records"
	DigiOceanAccessToken := digioceanAuth.Token.DigiOceanAccessToken // Fetch the DigiOceanAccessToken

	listDNSReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	listDNSReq.Header.Set("Content-Type", "application/json")
	listDNSReq.Header.Set("Authorization", "Bearer "+DigiOceanAccessToken)

	listDNSResp, err := http.DefaultClient.Do(listDNSReq)
	if err != nil {
		fmt.Println(err)
	}

	defer listDNSResp.Body.Close()

	responseBody, err := ioutil.ReadAll(listDNSResp.Body)
	listDNSResponse := make(map[string]interface{})
	listDNSResponse["status"] = listDNSResp.StatusCode
	listDNSResponse["body"] = string(responseBody)
	resp = listDNSResponse

	return resp, err
}

// ListResourcednsRecordSets function lists DNS record sets. DigitalOcean API
// doesn't provide functionality to suppport this function.
func (digioceandns *Digioceandns) ListResourcednsRecordSets(request interface{}) (resp interface{}, err error) {
	return resp, err
}
