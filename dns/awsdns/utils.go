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


func multimap(p map[string]string) url.Values {
	q := make(url.Values, len(p))
	for k, v := range p {
		q[k] = []string{v}
	}
	return q
}

func CleanZoneID(ID string) string {
	if strings.HasPrefix(ID, "/hostedzone/") {
		ID = strings.TrimPrefix(ID, "/hostedzone/")
	}
	return ID
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
