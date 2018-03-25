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

//Sign and do request by action parameter and specific parameters
func SignAndDoRequest(action string, params map[string]interface{}, response map[string]interface{}) error {
	// Add common params and action param
	params = initParams(action, params)

	// Sort the parameters by key
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	// Generate the stringToSign
	var canonicalizedQueryString string
	for _, k := range keys {
		v := params[k]
		canonicalizedQueryString += "&"
		canonicalizedQueryString += percentEncode(k)
		canonicalizedQueryString += "="
		canonicalizedQueryString += percentEncode(getString(v))
	}
	stringToSign := "GET" + "&%2F&" + percentEncode(canonicalizedQueryString[1:])

	// Create signature
	base64Sign := createSignature(stringToSign, Config.AliAccessKeySecret+"&")

	// Init url query
	query := url.Values{}
	for key, value := range params {
		query.Add(key, getString(value))
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
	params["Action"] = action

	params["Format"] = "XML"
	params["Version"] = "2014-05-26"
	params["AccessKeyId"] = Config.AliAccessKeyID
	params["TimeStamp"] = time.Now().UTC().Format(formatISO8601)
	params["SignatureMethod"] = "HMAC-SHA1"
	params["SignatureVersion"] = "1.0"
	params["SignatureNonce"] = createRandomString()

	return params
}

//CreateSignature creates signature for string following Ali-cloud rules
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

const dictionary = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// CreateRandomString create random string
func createRandomString() string {
	b := make([]byte, 32)
	l := len(dictionary)

	_, err := srand.Read(b)

	if err != nil {
		// Fail back to insecure rand
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

// getString return string from interface{}
func getString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int:
		return strconv.Itoa(v.(int))
	case bool:
		return strconv.FormatBool(v.(bool))
	}
	return ""
}
