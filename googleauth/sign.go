package googleauth

import (
	"fmt"
	"github.com/scorelab/gocloud-v2/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func SignJWT() (client *http.Client) {

	var home string = os.Getenv("HOME")

	data, err := ioutil.ReadFile(home + "/ShelterMap-3d450eb49f43_container.json")
	if err != nil {
		log.Fatal(err)
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
