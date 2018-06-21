package awsauth


import("io/ioutil"
"net/http"
"crypto/hmac"
"crypto/sha256"
"fmt"
"encoding/hex"
"bytes"
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
