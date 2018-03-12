package aliauth

import (
	"time"
	srand "crypto/rand"
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
	params = InitParams(action, params)

	var canonicalizedQueryString string

	//sort map by key
	keys := make([]string, len(params))
	i := 0
	for k, _ := range params {
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

	base64Sign := CreateSignature(stringToSign, Config.AliAccessKeySecret+"&")

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

func InitParams(action string, params map[string]interface{}) map[string]interface{} {
	params["Format"] = "XML"
	params["Version"] = "2014-05-26"
	params["AccessKeyId"] = Config.AliAccessKeyId
	params["TimeStamp"] = time.Now().UTC().Format(formatISO8601)
	params["SignatureMethod"] = "HMAC-SHA1"
	params["SignatureVersion"] = "1.0"
	randomInt := rand.Int63()
	params["SignatureNonce"] = strconv.FormatInt(randomInt, 10)

	params["Action"] = action

	return SortMap(params)
}

func SortMap(oldMap map[string]interface{}) map[string]interface{} {
	keys := make([]string, len(oldMap))
	i := 0
	for k, _ := range oldMap {
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
func CreateSignature(stringToSignature, accessKeySecret string) string {
	// Crypto by HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, []byte(accessKeySecret))
	hmacSha1.Write([]byte(stringToSignature))
	sign := hmacSha1.Sum(nil)

	// Encode to Base64
	base64Sign := base64.StdEncoding.EncodeToString(sign)

	return base64Sign
}

func percentEncode(str string) string {
	return PercentReplace(url.QueryEscape(str))
}

func PercentReplace(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)

	return str
}

// CreateSignatureForRequest creates signature for query string values
func CreateSignatureForRequest(method string, values *url.Values, accessKeySecret string) string {

	canonicalizedQueryString := PercentReplace(values.Encode())

	stringToSign := method + "&%2F&" + url.QueryEscape(canonicalizedQueryString)
	return CreateSignature(stringToSign, accessKeySecret)
}

const dictionary = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

//CreateRandomString create random string
func CreateRandomString() string {
	b := make([]byte, 32)
	l := len(dictionary)

	_, err := srand.Read(b)

	if err != nil {
		// fail back to insecure rand
		rand.Seed(time.Now().UnixNano())
		for i := range b {
			b[i] = dictionary[rand.Int()%l]
		}
	} else {
		for i, v := range b {
			b[i] = dictionary[v%byte(l)]
		}
	}

	return string(b)
}
