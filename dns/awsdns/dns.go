package awsdns

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

//ListResourceDnsRecordSets represents awsdns ListResourceDnsRecordSets.
func (awsdns *Awsdns) ListResourceDnsRecordSets(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var zone string

	var option ListResourceDnsRecordSets

	for key, value := range param {
		switch key {

		case "zone":
			zoneV, _ := value.(string)
			zone = zoneV

		case "name":
			nameV, _ := value.(string)
			option.Name = nameV

		case "type":
			typeV, _ := value.(string)
			option.Type = typeV

		case "maxItems":
			maxItemsV, _ := value.(int)
			option.MaxItems = maxItemsV

		case "identifier":
			identifierV, _ := value.(string)
			option.Identifier = identifierV

		}
	}

	params := make(map[string]string)

	if option.Name != "" {
		params["name"] = option.Name
	}

	if option.Type != "" {
		params["type"] = option.Type
	}

	if option.Identifier != "" {
		params["identifier"] = option.Identifier
	}

	if option.MaxItems != 0 {
		params["maxitems"] = strconv.FormatInt(int64(option.MaxItems), 10)
	}

	req := multimap(params)

	zone = CleanZoneID(zone)

	response := make(map[string]interface{})

	out := &ListResourceRecordSetsResponse{}

	awsdns.PrepareSignatureV4query("GET", fmt.Sprintf("/%s/hostedzone/%s/rrset", "2013-04-01", zone), req, out, response)

	resp = response

	return resp, nil
}

//ListDns ListDns awsdns.
func (awsdns *Awsdns) ListDns(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var marker string
	var maxItems int
	for key, value := range param {
		switch key {

		case "marker":
			markerV, _ := value.(string)
			marker = markerV

		case "maxItems":
			maxItemsV, _ := value.(int)
			maxItems = maxItemsV

		}
	}

	out := &ListHostedZonesResponse{}

	response := make(map[string]interface{})

	values := url.Values{}

	if marker != "" {
		values.Add("marker", marker)
	}

	if maxItems != 0 {
		values.Add("maxItems", strconv.Itoa(maxItems))
	}

	awsdns.PrepareSignatureV4query("GET", fmt.Sprintf("/%s/hostedzone/", "2013-04-01"), values, out, response)

	resp = response

	return resp, nil

}

func (awsdns *Awsdns) DeleteDns(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]string)

	out := &DeleteHostedZoneResponse{}

	response := make(map[string]interface{})

	awsdns.PrepareSignatureV4query("DELETE", fmt.Sprintf("/%s/hostedzone/%s", "2013-04-01", param["ID"]), nil, out, response)

	resp = response

	return resp, nil
}

//CreateDns creates awsdns.
func (awsdns *Awsdns) CreateDns(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var option CreateHostedZoneRequest

	for key, value := range param {
		switch key {

		case "name":
			nameV, _ := value.(string)
			option.Name = nameV

		case "callerReference":
			CallerReferenceV, _ := value.(string)
			option.CallerReference = CallerReferenceV

		case "comment":
			commentV, _ := value.(string)
			option.Comment = commentV
		}
	}

	if option.CallerReference == "" {
		option.CallerReference = time.Now().Format(time.RFC3339Nano)
	}

	out := &CreateHostedZoneResponse{}

	response := make(map[string]interface{})

	awsdns.PrepareSignatureV4query("POST", fmt.Sprintf("/%s/hostedzone", "2013-04-01"), option, out, response)

	resp = response

	return resp, nil
}
