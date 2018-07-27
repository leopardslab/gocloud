package lambda

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	auth "github.com/cloudlibz/gocloud/auth"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func hmacsignatureV4(signingKey []byte, stringToSign string) string {
	return hex.EncodeToString(hmacSHA256(signingKey, stringToSign))
}

func hmacSHA256(key []byte, content string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

func sha256Hasher(payload []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(payload))
}

func preparepayload(request *http.Request) []byte {

	if request.Body == nil {
		return []byte{}
	}
	payload, _ := ioutil.ReadAll(request.Body)
	request.Body = ioutil.NopCloser(bytes.NewReader(payload))
	return payload
}

func Preparedeletefnrequest(params map[string]string, region string, response map[string]interface{}) (err error) {
	service := "lambda"
	method := "DELETE"
	host := "lambda" + "." + region + ".amazonaws.com"
	signedheaders := "host;x-amz-date"

	endpoint := "https://lambda.us-east-1.amazonaws.com"

	AccessKeyID := auth.Config.AWSAccessKeyID
	SecretAccessKey := auth.Config.AWSSecretKey

	t := time.Now().UTC()

	XAmzDate := t.Format("20060102T150405Z")
	date_stamp := t.Format("20060102")

	canonical_uri := "/2015-03-31/functions/" + params["FunctionName"]

	request, _ := http.NewRequest("DELETE", endpoint + canonical_uri, nil)

	payload := preparepayload(request)
	payloadHash := sha256Hasher(payload)

/*
	requestparam := request.URL.Query()

	request.URL.RawQuery = requestparam.Encode()

	fmt.Println(request.URL)

	queryString := request.URL.RawQuery

*/
	queryString := ""

	// Go encodes a space as '+' but Amazon requires '%20'. Luckily any '+' in the
	// original query string has been percent escaped so all '+' chars that are left
	// were originally spaces.
	canonical_querystring := strings.Replace(queryString, "+", "%20", -1)
	canonical_headers := "host:" + host + "\n" + "x-amz-date:" + XAmzDate + "\n"
	canonical_request := method + "\n" + canonical_uri + "\n" + canonical_querystring + "\n" + canonical_headers + "\n" + signedheaders + "\n" + payloadHash

	algorithm := "AWS4-HMAC-SHA256"
	credential_scope := date_stamp + "/" + region + "/" + service + "/" + "aws4_request"
	string_to_sign := algorithm + "\n" + XAmzDate + "\n" + credential_scope + "\n" + sha256Hasher([]byte(canonical_request))

	kDate := hmacSHA256([]byte("AWS4"+SecretAccessKey), date_stamp)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "aws4_request")

	signature := hmacsignatureV4(kSigning, string_to_sign)
	authorization_header := algorithm + " " + "Credential=" + AccessKeyID + "/" + credential_scope + ", " + "SignedHeaders=" + signedheaders + ", " + "Signature=" + signature

	request.Header.Add("X-Amz-Date", XAmzDate)
	request.Header.Add("Authorization", authorization_header)

	client := new(http.Client)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	response["body"] = string(body)
	response["status"] = resp.StatusCode

	return
}


func Preparegetfnrequest(params map[string]string, region string, response map[string]interface{}) (err error) {
	service := "lambda"
	method := "GET"
	host := "lambda" + "." + region + ".amazonaws.com"
	signedheaders := "host;x-amz-date"

	endpoint := "https://lambda.us-east-1.amazonaws.com"

	AccessKeyID := auth.Config.AWSAccessKeyID
	SecretAccessKey := auth.Config.AWSSecretKey

	t := time.Now().UTC()

	XAmzDate := t.Format("20060102T150405Z")
	date_stamp := t.Format("20060102")

	canonical_uri := "/2015-03-31/functions/" + params["FunctionName"]

	request, _ := http.NewRequest("GET", endpoint + canonical_uri, nil)

	payload := preparepayload(request)
	payloadHash := sha256Hasher(payload)

/*
	requestparam := request.URL.Query()

	request.URL.RawQuery = requestparam.Encode()

	fmt.Println(request.URL)

	queryString := request.URL.RawQuery

*/
	queryString := ""

	// Go encodes a space as '+' but Amazon requires '%20'. Luckily any '+' in the
	// original query string has been percent escaped so all '+' chars that are left
	// were originally spaces.
	canonical_querystring := strings.Replace(queryString, "+", "%20", -1)
	canonical_headers := "host:" + host + "\n" + "x-amz-date:" + XAmzDate + "\n"
	canonical_request := method + "\n" + canonical_uri + "\n" + canonical_querystring + "\n" + canonical_headers + "\n" + signedheaders + "\n" + payloadHash

	algorithm := "AWS4-HMAC-SHA256"
	credential_scope := date_stamp + "/" + region + "/" + service + "/" + "aws4_request"
	string_to_sign := algorithm + "\n" + XAmzDate + "\n" + credential_scope + "\n" + sha256Hasher([]byte(canonical_request))

	kDate := hmacSHA256([]byte("AWS4"+SecretAccessKey), date_stamp)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "aws4_request")

	signature := hmacsignatureV4(kSigning, string_to_sign)
	authorization_header := algorithm + " " + "Credential=" + AccessKeyID + "/" + credential_scope + ", " + "SignedHeaders=" + signedheaders + ", " + "Signature=" + signature

	request.Header.Add("X-Amz-Date", XAmzDate)
	request.Header.Add("Authorization", authorization_header)

	client := new(http.Client)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	response["body"] = string(body)
	response["status"] = resp.StatusCode

	return
}

func Preparegetrequest(params map[string]string, region string, response map[string]interface{}) (err error) {
	service := "lambda"
	method := "GET"
	host := "lambda" + "." + region + ".amazonaws.com"
	signedheaders := "host;x-amz-date"

	endpoint := "https://lambda.us-east-1.amazonaws.com"

	AccessKeyID := auth.Config.AWSAccessKeyID
	SecretAccessKey := auth.Config.AWSSecretKey

	t := time.Now().UTC()

	XAmzDate := t.Format("20060102T150405Z")
	date_stamp := t.Format("20060102")

	canonical_uri := "/2015-03-31/functions/"

	request, _ := http.NewRequest("GET", endpoint+canonical_uri, nil)

	payload := preparepayload(request)
	payloadHash := sha256Hasher(payload)

	requestparam := request.URL.Query()

	for key, value := range params {
		requestparam.Add(key, value)
	}

	request.URL.RawQuery = requestparam.Encode()

	fmt.Println(request.URL)

	queryString := request.URL.RawQuery

	// Go encodes a space as '+' but Amazon requires '%20'. Luckily any '+' in the
	// original query string has been percent escaped so all '+' chars that are left
	// were originally spaces.
	canonical_querystring := strings.Replace(queryString, "+", "%20", -1)
	canonical_headers := "host:" + host + "\n" + "x-amz-date:" + XAmzDate + "\n"
	canonical_request := method + "\n" + canonical_uri + "\n" + canonical_querystring + "\n" + canonical_headers + "\n" + signedheaders + "\n" + payloadHash

	algorithm := "AWS4-HMAC-SHA256"
	credential_scope := date_stamp + "/" + region + "/" + service + "/" + "aws4_request"
	string_to_sign := algorithm + "\n" + XAmzDate + "\n" + credential_scope + "\n" + sha256Hasher([]byte(canonical_request))

	kDate := hmacSHA256([]byte("AWS4"+SecretAccessKey), date_stamp)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "aws4_request")

	signature := hmacsignatureV4(kSigning, string_to_sign)
	authorization_header := algorithm + " " + "Credential=" + AccessKeyID + "/" + credential_scope + ", " + "SignedHeaders=" + signedheaders + ", " + "Signature=" + signature

	request.Header.Add("X-Amz-Date", XAmzDate)
	request.Header.Add("Authorization", authorization_header)

	client := new(http.Client)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	response["body"] = string(body)
	response["status"] = resp.StatusCode

	return
}

func PreparePostrequest(params map[string]interface{}, region string, response map[string]interface{}) (err error) {

	service := "lambda"
	method := "POST"
	host := "lambda" + "." + region + ".amazonaws.com"

	ContentType := "application/x-amz-json-1.0"
	signedheaders := "content-type;host;x-amz-date"

	endpoint := "https://lambda.us-east-1.amazonaws.com"

	AccessKeyID := auth.Config.AWSAccessKeyID
	SecretAccessKey := auth.Config.AWSSecretKey

	t := time.Now().UTC()

	XAmzDate := t.Format("20060102T150405Z")
	date_stamp := t.Format("20060102")

	canonical_uri := "/2015-03-31/functions/"

	requestparametersjson, _ := json.Marshal(params)
	requestparametersjsonstring := string(requestparametersjson)

	requestparametersjsonstring = "{}"

	requestparametersjsonstringbyte := []byte(requestparametersjsonstring)

	request, _ := http.NewRequest("POST", endpoint+canonical_uri, bytes.NewBuffer(requestparametersjsonstringbyte))

	payload := preparepayload(request)
	payloadHash := sha256Hasher(payload)

	// Go encodes a space as '+' but Amazon requires '%20'. Luckily any '+' in the
	// original query string has been percent escaped so all '+' chars that are left
	// were originally spaces.
	canonical_querystring := ""

	canonical_headers := "content-type:" + ContentType + "\n" + "host:" + host + "\n" + "x-amz-date:" + XAmzDate + "\n"

	//payload_hashstring := sha256Hasher(request_parameters)

	canonical_request := method + "\n" + canonical_uri + "\n" + canonical_querystring + "\n" + canonical_headers + "\n" + signedheaders + "\n" + payloadHash

	algorithm := "AWS4-HMAC-SHA256"
	credential_scope := date_stamp + "/" + region + "/" + service + "/" + "aws4_request"
	string_to_sign := algorithm + "\n" + XAmzDate + "\n" + credential_scope + "\n" + sha256Hasher([]byte(canonical_request))

	kDate := hmacSHA256([]byte("AWS4"+SecretAccessKey), date_stamp)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "aws4_request")

	signature := hmacsignatureV4(kSigning, string_to_sign)
	authorization_header := algorithm + " " + "Credential=" + AccessKeyID + "/" + credential_scope + ", " + "SignedHeaders=" + signedheaders + ", " + "Signature=" + signature

	request.Header.Set("Content-Type", ContentType)
	request.Header.Set("Host", "lambda.us-east-1.amazonaws.com")
	request.Header.Add("X-Amz-Date", XAmzDate)
	request.Header.Add("Authorization", authorization_header)

	client := new(http.Client)

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	response["body"] = string(body)
	response["status"] = resp.StatusCode

	return
}
