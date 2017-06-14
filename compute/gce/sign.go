package gce

import (
	"fmt"
	//"io/ioutil"
	"log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	//"net"
	//"net/http"
	//"bytes"
//"errors"
)

func sign()(token *oauth2.Token) {

conf := &oauth2.Config{

	ClientID: "enterhere",

  ClientSecret: "enterhere",

	RedirectURL: "urn:ietf:wg:oauth:2.0:oob",

	Scopes: []string{
			"https://www.googleapis.com/auth/compute",
			"https://www.googleapis.com/auth/devstorage.full_control",
      "https://www.googleapis.com/auth/ndev.clouddns.readwrite",
	    "https://www.googleapis.com/auth/cloud-platform",
	},

	Endpoint: google.Endpoint,
	}


	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	fmt.Printf("Visit the URL for Autorization \n: %v", url)

  fmt.Printf("\nEnter Autorization code :\n")

	var code string

	if _, err := fmt.Scan(&code); err != nil {
		    log.Fatal(err)
	}

	//token = new(oauth2.Token)

	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
    		log.Fatal(err)
	}

//	fmt.Println(token)
	return token
}
