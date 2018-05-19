package aliauth

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"sort"
	"net/url"
	"io"
	"bytes"
	"net/http"
	"time"
	"io/ioutil"
)

func ContainerSignAndDoRequest(region string, method string, path string, query map[string]interface{}, args interface{}, response map[string]interface{}) error {
	var reqBody []byte
	var err error
	var contentMD5 string

	// generate request body and Content-MD5
	if args != nil {
		reqBody, err = json.Marshal(args)
		if err != nil {
			return err
		}
		contentMD5 = createContentMD5(reqBody)
	}

	// Sort the query by key
	keys := make([]string, len(query))
	i := 0
	for k := range query {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	// generate url query
	var queryString string
	for _, k := range keys {
		v := query[k]
		queryString += "&"
		queryString += url.QueryEscape(k)
		queryString += "="
		queryString += url.QueryEscape(getString(v))
	}

	requestURL := "cs.aliyuncs.com" + path + "?" + queryString[1:]

	var bodyReader io.Reader
	if reqBody != nil {
		bodyReader = bytes.NewReader(reqBody)
	}
	httpReq, err := http.NewRequest(method, requestURL, bodyReader)
	if err != nil {
		return err
	}

	if region != "" {
		httpReq.Header["x-acs-region-id"] = []string{string(region)}
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if contentMD5 != "" {
		httpReq.Header.Set("Content-MD5", contentMD5)
	}
	httpReq.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header["x-acs-signature-version"] = []string{"1.0"}
	httpReq.Header["x-acs-signature-nonce"] = []string{createRandomString()}
	httpReq.Header["x-acs-signature-method"] = []string{"HMAC-SHA1"}

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

func createContentMD5(reqBody []byte) string {
	hasher := md5.New()
	hasher.Write(reqBody)
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
