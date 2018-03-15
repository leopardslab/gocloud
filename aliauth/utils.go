package aliauth

import (
	"time"
	"math/rand"
	"net/url"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"strings"
	"sort"
	"net/http"
	"io/ioutil"
	"strconv"
)

const formatISO8601 = "2006-01-02T15:04:05Z"

func SignAndDoRequest(action string, params map[string]interface{}, response map[string]interface{}) error {
	params = initParams(action, params)

	var canonicalizedQueryString string

	//sort map by key
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := params[k]
		canonicalizedQueryString += "&"
		canonicalizedQueryString += percentEncode(k)
		canonicalizedQueryString += "="
		canonicalizedQueryString += percentEncode(v.(string))
	}

	stringToSign := "GET" + "&%2F&" + percentEncode(canonicalizedQueryString[1:])

	base64Sign := createSignature(stringToSign, Config.AliAccessKeySecret+"&")

	query := url.Values{}
	for key, value := range params {
		query.Add(key, value.(string))
	}

	// Generate the request URL
	requestURL := "https://ecs.aliyuncs.com/" + "?" + query.Encode() + "&Signature=" + url.QueryEscape(base64Sign)

	httpReq, err := http.NewRequest("GET", requestURL, nil)

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response["body"] = string(body)
	response["status"] = resp.StatusCode
	return err
}

func initParams(action string, params map[string]interface{}) map[string]interface{} {
	params["Format"] = "XML"
	params["Version"] = "2014-05-26"
	params["AccessKeyId"] = Config.AliAccessKeyID
	params["TimeStamp"] = time.Now().UTC().Format(formatISO8601)
	params["SignatureMethod"] = "HMAC-SHA1"
	params["SignatureVersion"] = "1.0"
	randomInt := rand.Int63()
	params["SignatureNonce"] = strconv.FormatInt(randomInt, 10)

	params["Action"] = action

	return sortMap(params)
}

func sortMap(oldMap map[string]interface{}) map[string]interface{} {
	keys := make([]string, len(oldMap))
	i := 0
	for k := range oldMap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	sortedMap := make(map[string]interface{})
	for _, k := range keys {
		sortedMap[k] = oldMap[k]
	}
	return sortedMap
}

//CreateSignature creates signature for string following Aliyun rules
func createSignature(stringToSignature, accessKeySecret string) string {
	// Crypto by HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, []byte(accessKeySecret))
	hmacSha1.Write([]byte(stringToSignature))
	sign := hmacSha1.Sum(nil)

	// Encode to Base64
	base64Sign := base64.StdEncoding.EncodeToString(sign)

	return base64Sign
}

func percentEncode(str string) string {
	return percentReplace(url.QueryEscape(str))
}

func percentReplace(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)

	return str
}
