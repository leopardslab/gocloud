package googledns

import (
	"bytes"
	"encoding/json"
	googleauth "github.com/cloudlibz/gocloud/googleauth"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	RFC3339  = "2006-01-02T15:04:05Z07:00"
)

//ListResourceDnsRecordSets list ListResourceDnsRecordSets.
func (googledns *Googledns) ListResourceDnsRecordSets(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/dns/v1/projects/" + options["project"] + "/managedZones/" + options["managedZone"] + "/rrsets"

	client := googleauth.SignJWT()

	ListResourceDnsRecordSetsrequest, err := http.NewRequest("GET", url, nil)

	ListResourceDnsRecordSetsparam := ListResourceDnsRecordSetsrequest.URL.Query()

	if options["maxResults"] != "" {
		ListResourceDnsRecordSetsparam.Add("deviceName", options["maxResults"])
	}

	if options["pageToken"] != "" {
		ListResourceDnsRecordSetsparam.Add("pageToken", options["pageToken"])
	}

	if options["sortBy"] != "" {
		ListResourceDnsRecordSetsparam.Add("sortBy", options["sortBy"])
	}

	if options["sortOrder"] != "" {
		ListResourceDnsRecordSetsparam.Add("sortOrder", options["sortOrder"])
	}

	ListResourceDnsRecordSetsrequest.URL.RawQuery = ListResourceDnsRecordSetsparam.Encode()

	ListResourceDnsRecordSetsrequest.Header.Set("Content-Type", "application/json")

	ListResourceDnsRecordSetssresp, err := client.Do(ListResourceDnsRecordSetsrequest)

	defer ListResourceDnsRecordSetssresp.Body.Close()

	body, err := ioutil.ReadAll(ListResourceDnsRecordSetssresp.Body)

	ListResourcednsRecordSetresponse := make(map[string]interface{})
	ListResourcednsRecordSetresponse["status"] = ListResourceDnsRecordSetssresp.StatusCode
	ListResourcednsRecordSetresponse["body"] = string(body)
	resp = ListResourcednsRecordSetresponse
	return resp, err
}

//CreateDns creates DNS.
func (googledns *Googledns) CreateDns(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var Project string
	var option CreateDns

	for key, value := range param {
		switch key {

		case "Project":
			ProjectV, _ := value.(string)
			Project = ProjectV

		case "CreationTime":
			CreationTimeV, _ := value.(string)
			option.CreationTime = CreationTimeV
			option.CreationTime = time.Now().UTC().Format(time.RFC3339)

		case "Description":
			DescriptionV, _ := value.(string)
			option.Description = DescriptionV

		case "DnsName":
			DNSNameV, _ := value.(string)
			option.DNSName = DNSNameV

		case "nameServers":
			nameServersV, _ := value.([]string)
			option.NameServers = nameServersV

		case "Id":
			IDV, _ := value.(string)
			option.ID = IDV

		case "Kind":
			kindV, _ := value.(string)
			option.Kind = kindV

		case "Name":
			NameV, _ := value.(string)
			option.Name = NameV

		case "nameServerSet":
			NameServerSetV, _ := value.(string)
			option.NameServerSet = NameServerSetV
		}
	}

	CreateDnsjsonmap := make(map[string]interface{})

	CreateDnsedictnoaryconvert(option, CreateDnsjsonmap)

	CreateDnsjson, _ := json.Marshal(CreateDnsjsonmap)

	CreateDnsjsonstring := string(CreateDnsjson)

	var CreateDnsjsonstringbyte = []byte(CreateDnsjsonstring)

	url := "https://www.googleapis.com/dns/v1/projects/" + Project + "/managedZones"

	client := googleauth.SignJWT()

	CreateDnsrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(CreateDnsjsonstringbyte))

	CreateDnsrequest.Header.Set("Content-Type", "application/json")

	CreateDnsrresp, err := client.Do(CreateDnsrequest)

	defer CreateDnsrresp.Body.Close()

	body, err := ioutil.ReadAll(CreateDnsrresp.Body)

	CreateDnsresponse := make(map[string]interface{})
	CreateDnsresponse["status"] = CreateDnsrresp.StatusCode
	CreateDnsresponse["body"] = string(body)
	resp = CreateDnsresponse
	return resp, err
}

//ListDns lists DNS.
func (googledns *Googledns) ListDns(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/dns/v1/projects/" + options["Project"] + "/managedZones/"

	client := googleauth.SignJWT()

	ListDnsrequest, err := http.NewRequest("GET", url, nil)

	ListDnsrequestparam := ListDnsrequest.URL.Query()

	if options["maxResults"] != "" {
		ListDnsrequestparam.Add("deviceName", options["maxResults"])
	}

	if options["pageToken"] != "0" {
		ListDnsrequestparam.Add("pageToken", options["pageToken"])
	}

	ListDnsrequest.URL.RawQuery = ListDnsrequestparam.Encode()

	ListDnsrequest.Header.Set("Content-Type", "application/json")

	ListDnsresp, err := client.Do(ListDnsrequest)

	defer ListDnsresp.Body.Close()

	body, err := ioutil.ReadAll(ListDnsresp.Body)

	ListDnsresponse := make(map[string]interface{})
	ListDnsresponse["status"] = ListDnsresp.StatusCode
	ListDnsresponse["body"] = string(body)
	resp = ListDnsresponse
	return resp, err
}

//DeleteDns deletes DNS.
func (googledns *Googledns) DeleteDns(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/dns/v1/projects/" + options["Project"] + "/managedZones/" + options["managedZone"]

	client := googleauth.SignJWT()

	DeleteDnsrequest, err := http.NewRequest("DELETE", url, nil)

	DeleteDnsrequest.Header.Set("Content-Type", "application/json")

	DeleteDnsresp, err := client.Do(DeleteDnsrequest)

	defer DeleteDnsresp.Body.Close()

	body, err := ioutil.ReadAll(DeleteDnsresp.Body)

	DeleteDnsresponse := make(map[string]interface{})
	DeleteDnsresponse["status"] = DeleteDnsresp.StatusCode
	DeleteDnsresponse["body"] = string(body)
	resp = DeleteDnsresponse
	return resp, err
}
