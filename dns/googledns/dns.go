package googledns

import (
	"bytes"
	"encoding/json"
	googleauth "github.com/shlokgilda/gocloud/googleauth"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	RFC3339  = "2006-01-02T15:04:05Z07:00"
)

//ListResourcednsRecordSets list ListResourcednsRecordSets.
func (googledns *Googledns) ListResourcednsRecordSets(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/dns/v1/projects/" + options["project"] + "/managedZones/" + options["managedZone"] + "/rrsets"

	client := googleauth.SignJWT()

	ListResourcednsRecordSetsrequest, err := http.NewRequest("GET", url, nil)

	ListResourcednsRecordSetsparam := ListResourcednsRecordSetsrequest.URL.Query()

	if options["maxResults"] != "" {
		ListResourcednsRecordSetsparam.Add("deviceName", options["maxResults"])
	}

	if options["pageToken"] != "" {
		ListResourcednsRecordSetsparam.Add("pageToken", options["pageToken"])
	}

	if options["sortBy"] != "" {
		ListResourcednsRecordSetsparam.Add("sortBy", options["sortBy"])
	}

	if options["sortOrder"] != "" {
		ListResourcednsRecordSetsparam.Add("sortOrder", options["sortOrder"])
	}

	ListResourcednsRecordSetsrequest.URL.RawQuery = ListResourcednsRecordSetsparam.Encode()

	ListResourcednsRecordSetsrequest.Header.Set("Content-Type", "application/json")

	ListResourcednsRecordSetssresp, err := client.Do(ListResourcednsRecordSetsrequest)

	defer ListResourcednsRecordSetssresp.Body.Close()

	body, err := ioutil.ReadAll(ListResourcednsRecordSetssresp.Body)

	ListResourcednsRecordSetresponse := make(map[string]interface{})
	ListResourcednsRecordSetresponse["status"] = ListResourcednsRecordSetssresp.StatusCode
	ListResourcednsRecordSetresponse["body"] = string(body)
	resp = ListResourcednsRecordSetresponse
	return resp, err
}

//Createdns creates DNS.
func (googledns *Googledns) Createdns(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var Project string
	var option Createdns

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

	Creatednsjsonmap := make(map[string]interface{})

	Creatednsedictnoaryconvert(option, Creatednsjsonmap)

	Creatednsjson, _ := json.Marshal(Creatednsjsonmap)

	Creatednsjsonstring := string(Creatednsjson)

	var Creatednsjsonstringbyte = []byte(Creatednsjsonstring)

	url := "https://www.googleapis.com/dns/v1/projects/" + Project + "/managedZones"

	client := googleauth.SignJWT()

	Creatednsrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Creatednsjsonstringbyte))

	Creatednsrequest.Header.Set("Content-Type", "application/json")

	Creatednsrresp, err := client.Do(Creatednsrequest)

	defer Creatednsrresp.Body.Close()

	body, err := ioutil.ReadAll(Creatednsrresp.Body)

	Creatednsresponse := make(map[string]interface{})
	Creatednsresponse["status"] = Creatednsrresp.StatusCode
	Creatednsresponse["body"] = string(body)
	resp = Creatednsresponse
	return resp, err
}

//Listdns lists DNS.
func (googledns *Googledns) Listdns(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/dns/v1/projects/" + options["Project"] + "/managedZones/"

	client := googleauth.SignJWT()

	Listdnsrequest, err := http.NewRequest("GET", url, nil)

	Listdnsrequestparam := Listdnsrequest.URL.Query()

	if options["maxResults"] != "" {
		Listdnsrequestparam.Add("deviceName", options["maxResults"])
	}

	if options["pageToken"] != "0" {
		Listdnsrequestparam.Add("pageToken", options["pageToken"])
	}

	Listdnsrequest.URL.RawQuery = Listdnsrequestparam.Encode()

	Listdnsrequest.Header.Set("Content-Type", "application/json")

	Listdnsresp, err := client.Do(Listdnsrequest)

	defer Listdnsresp.Body.Close()

	body, err := ioutil.ReadAll(Listdnsresp.Body)

	Listdnsresponse := make(map[string]interface{})
	Listdnsresponse["status"] = Listdnsresp.StatusCode
	Listdnsresponse["body"] = string(body)
	resp = Listdnsresponse
	return resp, err
}

//Deletedns deletes DNS.
func (googledns *Googledns) Deletedns(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/dns/v1/projects/" + options["Project"] + "/managedZones/" + options["managedZone"]

	client := googleauth.SignJWT()

	Deletednsrequest, err := http.NewRequest("DELETE", url, nil)

	Deletednsrequest.Header.Set("Content-Type", "application/json")

	Deletednsresp, err := client.Do(Deletednsrequest)

	defer Deletednsresp.Body.Close()

	body, err := ioutil.ReadAll(Deletednsresp.Body)

	Deletednsresponse := make(map[string]interface{})
	Deletednsresponse["status"] = Deletednsresp.StatusCode
	Deletednsresponse["body"] = string(body)
	resp = Deletednsresponse
	return resp, err
}
