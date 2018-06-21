package awsauth

import (
	auth "github.com/cloudlibz/gocloud/auth"
	"net/http"
	"time"
)

func SignatureV4(request *http.Request, request_parameters []byte, amztarget string, method string, region string, service string, host string, ContentType string, signedheaders string) (signrequest *http.Request) {
	AccessKeyID := auth.Config.AWSAccessKeyID
	SecretAccessKey := auth.Config.AWSSecretKey

	t := time.Now().UTC()
	XAmzDate := t.Format("20060102T150405Z")
	date_stamp := t.Format("20060102")
	canonical_uri := "/"
	canonical_querystring := ""
	canonical_headers := "content-type:" + ContentType + "\n" + "host:" + host + "\n" + "x-amz-date:" + XAmzDate + "\n" + "x-amz-target:" + amztarget + "\n"
	payload_hashstring := sha256Hasher(request_parameters)
	canonical_request := method + "\n" + canonical_uri + "\n" + canonical_querystring + "\n" + canonical_headers + "\n" + signedheaders + "\n" + payload_hashstring

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
	request.Header.Add("X-Amz-Date", XAmzDate)
	request.Header.Add("X-Amz-Target", amztarget)
	request.Header.Add("Authorization", authorization_header)

	return request
}
