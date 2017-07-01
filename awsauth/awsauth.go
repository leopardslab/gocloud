package awsauth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//Sign v2 method for Authenticating request

func SignatureV2(req *http.Request, Auth interface{}) (err error) {
	auth, _ := Auth.(map[string]string)
	queryVals := req.URL.Query()
	queryVals.Set("AWSAccessKeyId", auth["AccessKey"])
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

	hash := hmac.New(sha256.New, []byte(auth["SecretKey"]))

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
