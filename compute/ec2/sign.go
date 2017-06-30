package ec2

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//Sign v2 method for Authenticating request

func SignatureV2(req *http.Request, auth Auth) (err error) {

	queryVals := req.URL.Query()

	queryVals.Set("AWSAccessKeyId", auth.AccessKey)

	queryVals.Set("SignatureVersion", "2")

	queryVals.Set("SignatureMethod", "HmacSHA256")

	queryStr, err := canonicalQueryString(queryVals)
	if err != nil {
		return err
	}

	path := req.URL.Path
	if path == "" {
		path = "/"
	}

	payload := new(bytes.Buffer)

	fmt.Println("req.Method:\n", req.Method)

	payloadstring := checkrequestMethod(req.Method) + "\n" + req.Host + "\n" + path + "\n" + queryStr

	fmt.Fprintf(payload, "%s", payloadstring)

	hash := hmac.New(sha256.New, []byte(auth.SecretKey))

	hash.Write(payload.Bytes())

	signature := make([]byte, base64.StdEncoding.EncodedLen(hash.Size()))

	base64.StdEncoding.Encode(signature, hash.Sum(nil))

	queryVals.Set("Signature", string(signature))

	req.URL.RawQuery = queryVals.Encode()

	return nil
}

func canonicalQueryString(queryString url.Values) (string, error) {

	return strings.Replace(queryString.Encode(), "+", "%20", -1), nil
}

func clientToken() (string, error) {

	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

func checkrequestMethod(rawMethod string) (verb string) {
	fmt.Println(rawMethod)
	rawMethodverb := strings.SplitN(rawMethod, " ", 2)
	switch {
	case len(rawMethodverb) == 0:
		verb = "GET"
	default:
		verb = rawMethodverb[0]
	}
	return verb
}

func buildError(r *http.Response) error {
	errors := xmlErrors{}
	xml.NewDecoder(r.Body).Decode(&errors)
	var err Error
	if len(errors.Errors) > 0 {
		err = errors.Errors[0]
	}
	err.RequestId = errors.RequestId
	err.StatusCode = r.StatusCode
	if err.Message == "" {
		err.Message = r.Status
	}
	return &err
}

type xmlErrors struct {
	RequestId string  `xml:"RequestID"`
	Errors    []Error `xml:"Errors>Error"`
}

var timeNow = time.Now

func (err *Error) Error() string {
	if err.Code == "" {
		return err.Message
	}

	return fmt.Sprintf("%s (%s)", err.Message, err.Code)
}
