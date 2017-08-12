package awsdns

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	auth "github.com/scorelab/gocloud-v2/auth"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Awsdns struct {
}

func multimap(p map[string]string) url.Values {
	q := make(url.Values, len(p))
	for k, v := range p {
		q[k] = []string{v}
	}
	return q
}

type ListResourcednsRecordSets struct {
	Name       string
	Type       string
	Identifier string
	MaxItems   int
}

type ChangeResourceRecordSetsRequest struct {
	Comment string   `xml:"ChangeBatch>Comment,omitempty"`
	Changes []Change `xml:"ChangeBatch>Changes>Change"`
}

type Change struct {
	Action string            `xml:"Action"`
	Record ResourceRecordSet `xml:"ResourceRecordSet"`
}

type AliasTarget struct {
	HostedZoneId         string
	DNSName              string
	EvaluateTargetHealth bool
}

type HostedZone struct {
	ID              string `xml:"Id"`
	Name            string `xml:"Name"`
	CallerReference string `xml:"CallerReference"`
	Comment         string `xml:"Config>Comment"`
	ResourceCount   int    `xml:"ResourceRecordSetCount"`
}

type ChangeInfo struct {
	ID          string `xml:"Id"`
	Status      string `xml:"Status"`
	SubmittedAt string `xml:"SubmittedAt"`
}

type DelegationSet struct {
	NameServers []string `xml:"NameServers>NameServer"`
}

type ResourceRecordSet struct {
	Name          string       `xml:"Name"`
	Type          string       `xml:"Type"`
	TTL           int          `xml:"TTL"`
	Records       []string     `xml:"ResourceRecords>ResourceRecord>Value,omitempty"`
	SetIdentifier string       `xml:"SetIdentifier,omitempty"`
	Weight        int          `xml:"Weight,omitempty"`
	HealthCheckId string       `xml:"HealthCheckId,omitempty"`
	Region        string       `xml:"Region,omitempty"`
	Failover      string       `xml:"Failover,omitempty"`
	AliasTarget   *AliasTarget `xml:"AliasTarget,omitempty"`
	RecordsXML    string       `xml:",innerxml"`
}

type CreateHostedZoneRequest struct {
	Name            string `xml:"Name"`
	CallerReference string `xml:"CallerReference"`
	Comment         string `xml:"HostedZoneConfig>Comment"`
}

type CreateHostedZoneResponse struct {
	HostedZone    HostedZone    `xml:"HostedZone"`
	ChangeInfo    ChangeInfo    `xml:"ChangeInfo"`
	DelegationSet DelegationSet `xml:"DelegationSet"`
}

type DeleteHostedZoneResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

type ListHostedZonesResponse struct {
	HostedZones []HostedZone `xml:"HostedZones>HostedZone"`
	Marker      string       `xml:"Marker"`
	IsTruncated bool         `xml:"IsTruncated"`
	NextMarker  string       `xml:"NextMarker"`
	MaxItems    int          `xml:"MaxItems"`
}

type ChangeResourceRecordSetsResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

type ListResourceRecordSetsResponse struct {
	Records              []ResourceRecordSet `xml:"ResourceRecordSets>ResourceRecordSet"`
	IsTruncated          bool                `xml:"IsTruncated"`
	MaxItems             int                 `xml:"MaxItems"`
	NextRecordName       string              `xml:"NextRecordName"`
	NextRecordType       string              `xml:"NextRecordType"`
	NextRecordIdentifier string              `xml:"NextRecordIdentifier"`
}

func CleanZoneID(ID string) string {
	if strings.HasPrefix(ID, "/hostedzone/") {
		ID = strings.TrimPrefix(ID, "/hostedzone/")
	}
	return ID
}

func (awsdns *Awsdns) Changednsrecordsets(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	fmt.Println(param)
	var zone string
	var Comment string
	var option *ChangeResourceRecordSetsRequest
	var change []Change
	for key, value := range param {
		switch key {
		case "zone":
			zoneV, _ := value.(string)
			zone = zoneV

		case "comment":
			commentV, _ := value.(string)
			Comment = commentV

		case "changes":
			ChangesV, _ := value.([]map[string]interface{})
			for i := 0; i < len(ChangesV); i++ {
				var changeparam Change
				for ChangesVkey, ChangesVvalue := range ChangesV[i] {
					switch ChangesVkey {
					case "action":
						changeparam.Action = ChangesVvalue.(string)

					case "record":
						recordparam, _ := ChangesVvalue.(map[string]interface{})
						var recordparams ResourceRecordSet
						for recordparamkey, recordparamvalue := range recordparam {

						    switch recordparamkey {
							      case "name":
								        recordparams.Name = recordparamvalue.(string)

									  case "type":
									      recordparams.Type = recordparamvalue.(string)

									  case "ttl":
										   recordparams.TTL = recordparamvalue.(int)

									  case "records":
											 recordparams.Records = recordparamvalue.([]string)

							}
						}

						changeparam.Record = recordparams
						fmt.Println("recordparam",recordparam)

					}
				}
				change = append(change, changeparam)
			}

		}

	}

	option = &ChangeResourceRecordSetsRequest{Comment: Comment, Changes: change}
	fmt.Println("option", option)
	reqCopy := *option

	for i, change := range reqCopy.Changes {
		if len(change.Record.Records) > 1 {
			var buf bytes.Buffer
			for _, r := range change.Record.Records {
				buf.WriteString(fmt.Sprintf(
					"<ResourceRecord><Value>%s</Value></ResourceRecord>",
					r))
			}
			change.Record.Records = nil
			change.Record.RecordsXML = fmt.Sprintf(
				"<ResourceRecords>%s</ResourceRecords>", buf.String())
			reqCopy.Changes[i] = change
		}
	}

	zone = CleanZoneID(zone)

	out := &ChangeResourceRecordSetsResponse{}

	response := make(map[string]interface{})

	awsdns.PrepareSignatureV4query("POST", fmt.Sprintf("/%s/hostedzone/%s/rrset", "2013-04-01", zone), reqCopy, out, response)

	fmt.Println("response body :", response["body"])

	resp = response

	return resp, nil
}

func (awsdns *Awsdns) ListResourcednsRecordSets(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var zone string

	var option ListResourcednsRecordSets

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

	fmt.Println("response body :", response["body"])

	resp = response

	return resp, nil
}

func (awsdns *Awsdns) Listdns(request interface{}) (resp interface{}, err error) {
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

	fmt.Println("response body :", response["body"])

	resp = response

	return resp, nil

}

func (awsdns *Awsdns) Deletedns(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]string)

	out := &DeleteHostedZoneResponse{}

	response := make(map[string]interface{})

	awsdns.PrepareSignatureV4query("DELETE", fmt.Sprintf("/%s/hostedzone/%s", "2013-04-01", param["ID"]), nil, out, response)

	fmt.Println("response body :", response["body"])

	resp = response

	return resp, nil
}

func (awsdns *Awsdns) Createdns(request interface{}) (resp interface{}, err error) {

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

	fmt.Println("response body :", response["body"])

	resp = response

	return resp, nil
}

func (awsdns *Awsdns) PrepareSignatureV4query(method, path string, req, resp interface{}, response map[string]interface{}) error {
	params := make(map[string]string)
	endpoint, err := url.Parse("https://route53.amazonaws.com")
	if err != nil {
		return err
	}
	endpoint.Path = path
	sign(endpoint.Path, params)
	fmt.Println("params : ", params)

	if queryArgs, ok := req.(url.Values); ok {
		endpoint.RawQuery = queryArgs.Encode()
		req = nil
	}

	var body io.ReadWriter
	if req != nil {
		bodyBuf := bytes.NewBuffer(nil)
		enc := xml.NewEncoder(bodyBuf)
		start := xml.StartElement{
			Name: xml.Name{
				Space: "",
				Local: reflect.Indirect(reflect.ValueOf(req)).Type().Name(),
			},
			Attr: []xml.Attr{{xml.Name{"", "xmlns"}, "https://route53.amazonaws.com/doc/2013-04-01/"}},
		}
		if err := enc.EncodeElement(req, start); err != nil {
			return err
		}

		replace := "<ResourceRecords><ResourceRecord></ResourceRecord></ResourceRecords>"
		if strings.Contains(bodyBuf.String(), replace) {
			var newBuf bytes.Buffer
			newBuf.WriteString(strings.Replace(bodyBuf.String(), replace, "", -1))
			bodyBuf = &newBuf
		}

		if reflect.Indirect(reflect.ValueOf(req)).Type().Name() == "ChangeResourceRecordSetsRequest" {
			for _, change := range req.(ChangeResourceRecordSetsRequest).Changes {
				if change.Record.AliasTarget != nil {
					replace := change.Record.Type + "</Type><TTL>0</TTL>"
					var newBuf bytes.Buffer
					newBuf.WriteString(strings.Replace(bodyBuf.String(), replace, change.Record.Type+"</Type>", -1))
					bodyBuf = &newBuf
				}
			}
		}

		body = bodyBuf
	}

	hReq, err := http.NewRequest(method, endpoint.String(), body)
	if err != nil {
		return err
	}
	for k, v := range params {
		hReq.Header.Set(k, v)
	}
	client := new(http.Client)

	re, err := client.Do(hReq)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	bodystring, _ := ioutil.ReadAll(re.Body)

	response["status"] = re.StatusCode
	response["body"] = string(bodystring)

	switch re.StatusCode {
	case 200:
	case 201:
	default:
		var body bytes.Buffer
		io.Copy(&body, re.Body)
		return fmt.Errorf("Request failed, got status code: %d. Response: %s",
			re.StatusCode, body.Bytes())
	}

	decoder := xml.NewDecoder(re.Body)
	return decoder.Decode(resp)
}

var b64 = base64.StdEncoding

func sign(path string, params map[string]string) {

	AccessKey := auth.Config.AWSAccessKeyID

	SecretKey := auth.Config.AWSSecretKey

	date := time.Now().In(time.UTC).Format(time.RFC1123)

	params["Date"] = date

	fmt.Println("SecretKey:", SecretKey)

	hash := hmac.New(sha256.New, []byte(SecretKey))

	hash.Write([]byte(date))

	signature := make([]byte, b64.EncodedLen(hash.Size()))

	b64.Encode(signature, hash.Sum(nil))

	fmt.Println("AccessKey:", AccessKey)

	header := fmt.Sprintf("AWS3-HTTPS AWSAccessKeyId=%s,Algorithm=HmacSHA256,Signature=%s", AccessKey, signature)

	params["X-Amzn-Authorization"] = string(header)

}
