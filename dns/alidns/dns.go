package alidns

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"strconv"
	"fmt"
)

// ListResourcednsRecordSets list resource DNS record sets accept map[string]interface{}
func (alidns *Alidns) ListResourcednsRecordSets(request interface{}) (resp interface{}, err error) {
	fmt.Println("\nThis API is not provided by Alibaba cloud")
	return resp, err
}

// Listdns list DNS record accept map[string]interface{}
func (alidns *Alidns) Listdns(request interface{}) (resp interface{}, err error) {
	var options ListDNS

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "DomainName":
			options.DomainName = value.(string)
		case "PageNumber":
			switch value.(type) {
			case int:
				options.PageNumber = value.(int)
			case string:
				options.PageNumber, _ = strconv.Atoi(value.(string))
			}
		case "PageSize":
			switch value.(type) {
			case int:
				options.PageSize = value.(int)
			case string:
				options.PageSize, _ = strconv.Atoi(value.(string))
			}
		case "RRKeyWord":
			options.RRKeyWord = value.(string)
		case "TypeKeyWord":
			options.TypeKeyWord = value.(string)
		case "ValueKeyWord":
			options.ValueKeyWord = value.(string)
		}
	}
	// Put all of options into params
	params := aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.DNSSignAndDoRequest("DescribeDomainRecords", params, response)
	resp = response
	return resp, err
}

// Deletedns delete DNS record accept map[string]interface{}
func (alidns *Alidns) Deletedns(request interface{}) (resp interface{}, err error) {
	var options DeleteDNS

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "RecordId":
			options.RecordId = value.(string)
		}
	}
	// Put all of options into params
	params := aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.DNSSignAndDoRequest("DeleteDomainRecord", params, response)
	resp = response
	return resp, err
}

// Createdns add DNS record accept map[string]interface{}
func (alidns *Alidns) Createdns(request interface{}) (resp interface{}, err error) {
	var options CreateDNS

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "DomainName":
			options.DomainName = value.(string)
		case "RR":
			options.RR = value.(string)
		case "Type":
			options.Type = value.(string)
		case "Value":
			options.Value = value.(string)
		case "TTL":
			switch value.(type) {
			case int:
				options.TTL = value.(int)
			case string:
				options.TTL, _ = strconv.Atoi(value.(string))
			}
		case "Priority":
			switch value.(type) {
			case int:
				options.Priority = value.(int)
			case string:
				options.Priority, _ = strconv.Atoi(value.(string))
			}
		case "Line":
			options.Line = value.(string)
		}
	}
	// Put all of options into params
	params := aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.DNSSignAndDoRequest("AddDomainRecord", params, response)
	resp = response
	return resp, err
}
