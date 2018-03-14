package googleauth

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)


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

	data, err := ioutil.ReadFile(home + "/googlecloudconfig.json")

	if err != nil {
		jwt := JWT{}

		jwt.PrivateKey = os.Getenv("PrivateKey")

		jwt.Type = os.Getenv("Type")

		jwt.ProjectID = os.Getenv("ProjectID")

		jwt.PrivateKeyID = os.Getenv("PrivateKeyID")

		jwt.ClientID = os.Getenv("ClientID")

		jwt.ClientEmail = os.Getenv("ClientEmail")

		jwt.AuthURI = os.Getenv("AuthURI")

		jwt.TokenURI = os.Getenv("TokenURI")

		jwt.AuthProviderX509CertURL = os.Getenv("AuthProviderX509CertURL")

		jwt.ClientX509CertURL = os.Getenv("ClientX509CertURL")

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
