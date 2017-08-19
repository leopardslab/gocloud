package googleauth

import (
	"encoding/json"
	"fmt"
	"github.com/scorelab/gocloud-v2/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//sign() GCE signature it give URL to get Autorization code on which we it generate auth token and pass in each request in request header
func Sign() (token *oauth2.Token) {

	conf := &oauth2.Config{

		ClientID: auth.Config.GoogleClientID,

		ClientSecret: auth.Config.GoogleClientSecret,

		RedirectURL: "urn:ietf:wg:oauth:2.0:oob",

		Scopes: []string{
			"https://www.googleapis.com/auth/compute",
			"https://www.googleapis.com/auth/devstorage.full_control",
			"https://www.googleapis.com/auth/ndev.clouddns.readwrite",
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/ndev.clouddns.readonly",
		},

		Endpoint: google.Endpoint,
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("\nVisit the URL for Autorization \n: %v", url)
	fmt.Printf("\nEnter Autorization code :\n")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	return token
}

//JWT struct reperesnts JWT json.
type JWT struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}


//SignJWT reperesnts google service account authentication.
func SignJWT() (client *http.Client) {

	var home string = os.Getenv("HOME")

	data, err := ioutil.ReadFile(home + "Hello.json")//"gocloud-v2-testing.json")

	if err != nil {
		jwt := JWT{}
		jwt.PrivateKey = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDilnkJluzrS2ZW\nK6dLkG0G58IMAypcT95NNVMmLG6jr3CLuWv0nCsp7bX8cTYeT5WqaX1sRX8yVREV\nhiJpIr4aC39O44v+3cxTawHJdYvZtq43S4wFPwRbznRE4/JwPW2tA9WYEksfbCd+\nSdUddngghpDzLmG9IjPwd0nDo5tyd+5blM8KtACTMWIe9DozpH6TT507vEeviLJ1\nYEvw8W8qPiLBmeXk9G0t8VbCtKewpBpgRzxGxW5mt4gXclKIrEXTmjPBFrGKrrea\nCHnro+orlXm8TXtfmFLBSyNfBZCWeLxyneaERQlBa2zlkJ5xRx00lz1cPff1weKE\noKtZncKXAgMBAAECggEAXkz3hLkuUuKnbZdvYbBy35vWk39Bwvducj2OJGX7id36\n/fumZUAVarbGtZA8ULMYe7q5ezUVMdJbcNTJP245JcfBLQjge6q9Xb48GSb1lB9G\nNDcZyBko7DhpBT7E6Fd4bLgazdZXM3WknheGPWPPrOG0M7EK9lbkvOA4b9aIXEOr\nE6rtaDj853KQHbJGQiZULJH5/RqFHx1JVfra7K9npGPeewfpSbsuoFrxwaClGLC4\nmzHVwK/YQpJb+TxYfs8Ia7VGQYiRQ0JLbro/eqqKEy+ZbJfgKnpfbQqPYpD3t5C+\nOt5BdAUZGauoYDSRVqjcJOa+fNGHAg8NVMCjAGJnOQKBgQDxmhEyOZaC8V+v35tc\nNo+wVr7EcEyYNlThFDGFQ3XmsuIdDiJd65Y3TZ8aDH8hOiq/jJWaXbRFXY5l8xtt\ne2oUixUjuPZOpef2q9g6TAVARMcitSXAjUA63sKyKdDw0kVEPhcUoNU+gLcoujJV\nhJCmeNnmUnnELP52HlKUxownhQKBgQDwF1kVnJ3Bzp0S32kMgEWLIOKq8aHNz+L5\nGSKHRVyhtPHg5mG0urZ7OTgPEnzB306DzA8bYxPyynUIKmZT4P1h55Np+n2P2CND\niOiMPcWSEsqUcobi2WDjhsukImvNESudkIz5WBazNxJjM08tiNXvv3BOb3ZbfYtZ\nN3qp/5WmawKBgHXhHhLLVkx9SJqfF0ZeuKzpJhriT/zURYWsPJ0w1Q3CqeBTTSFW\nGOzSypYDEuQ7ZTAgf73ZimNtylSj0PYSCSgM1Duhu0uUVyFbpxyKiuVYqXCdwdW5\nFWWj2orjLDbT4Ufyo4BFGMZuNu+AKZ7gF1OiE271PsQgz+cB1HkShr+hAoGAJHsH\nuDbx6Y7hYwq7RYEAECHRDzRj54fc7wiYrIEnkBKUZh3bXsC4FYUeNXwTpMmvfms1\nKG4ni86jdbgrkDcxiPzM819yULcAtLRK0XRZXtaoHWJBiJqLFEdZDfmE88XWILzk\nDscJu/V0P8p+D2cpSqKGCAT7sO8ki0vYwZfqPz0CgYEA7DZHf8pZ0ls9gHA4SfjN\n+4x6SVKf7uaAhZQQeJiHdabod0Lv2fTL7iIrLvr6CJJuNDNabGrNUCcHRuSOFsX6\n5BxwGzu8IoZr+OdUxTNWH99M2KgKaedOzc+UWYWydNPQjKQqJvRzl72TcihCw/Gm\n4GrbT+IssTAOvwngM6TQfYk=\n-----END PRIVATE KEY-----\n"
		jwt.Type = "service_account"
		jwt.ProjectID = "sheltermap-1493101612061"
		jwt.PrivateKeyID =  "73bdb0a53ce9638b1325291416bb91e65ac20641"
		jwt.ClientID = "114946447985527565666"
		jwt.ClientEmail = "containerv2@sheltermap-1493101612061.iam.gserviceaccount.com"
		jwt.AuthURI = "https://accounts.google.com/o/oauth2/auth"
		jwt.TokenURI = "https://accounts.google.com/o/oauth2/token"
		jwt.AuthProviderX509CertURL = "https://www.googleapis.com/oauth2/v1/certs"
		jwt.ClientX509CertURL =  "https://www.googleapis.com/robot/v1/metadata/x509/containerv2%40sheltermap-1493101612061.iam.gserviceaccount.com"

		jwtjson, _ := json.Marshal(jwt)

		jwtjsonstring := string(jwtjson)

		datastr := strings.NewReader(jwtjsonstring)

		data, err = ioutil.ReadAll(datastr)
		if err != nil {
			log.Fatal(err)
		}
	}

	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/compute",
		"https://www.googleapis.com/auth/devstorage.full_control",
		"https://www.googleapis.com/auth/ndev.clouddns.readwrite",
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/ndev.clouddns.readonly",
	)
	if err != nil {
		log.Fatal(err)
	}

	client = conf.Client(oauth2.NoContext)
	return client
}
