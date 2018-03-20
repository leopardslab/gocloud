package awsauth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	auth "github.com/cloudlibz/gocloud/auth"
	"net/http"
	"time"
)

func SignatureV4(request *http.Request, requestParameters []byte, amztarget string, method string, region string, service string, host string, ContentType string, signedheaders string) (signrequest *http.Request) {
	AccessKeyID := auth.Config.AWSAccessKeyID
	SecretAccessKey := auth.Config.AWSSecretKey

	t := time.Now().UTC()
	XAmzDate := t.Format("20060102T150405Z")
	dateStamp := t.Format("20060102")
	canonicalURI := "/"
	canonicalQuerystring := ""
	canonicalHeaders := "content-type:" + ContentType + "\n" + "host:" + host + "\n" + "x-amz-date:" + XAmzDate + "\n" + "x-amz-target:" + amztarget + "\n"
	payloadHashstring := sha256Hasher(requestParameters)
	canonicalRequest := method + "\n" + canonicalURI + "\n" + canonicalQuerystring + "\n" + canonicalHeaders + "\n" + signedheaders + "\n" + payloadHashstring

	algorithm := "AWS4-HMAC-SHA256"
	credentialScope := dateStamp + "/" + region + "/" + service + "/" + "aws4_request"
	stringToSign := algorithm + "\n" + XAmzDate + "\n" + credentialScope + "\n" + sha256Hasher([]byte(canonicalRequest))

	date := hmacSHA256([]byte("AWS4" + SecretAccessKey), dateStamp)
	region := hmacSHA256(date, region)
	service := hmacSHA256(region, service)
	signing := hmacSHA256(service, "aws4_request")

	signature := hmacsignatureV4(signing, stringToSign)
	authorizationHeader := algorithm + " " + "Credential=" + AccessKeyID + "/" + credentialScope + ", " + "SignedHeaders=" + signedheaders + ", " + "Signature=" + signature

	request.Header.Set("Content-Type", ContentType)
	request.Header.Add("X-Amz-Date", XAmzDate)
	request.Header.Add("X-Amz-Target", amztarget)
	request.Header.Add("Authorization", authorizationHeader)

	return request
}

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
