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

func SignJWT() (client *http.Client) {

	var home string = os.Getenv("HOME")

	data, err := ioutil.ReadFile(home + "/ShelterMap-3d450eb49f43_containe.json")

	if err != nil {
		jwt := JWT{}

		PrivateKey := "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDWJKDV4ZIqH8gx\namJBiSr4rZsZg3oFmZEDuRJeRWYX//uUWl3anBPV0j/tuFojUZSzeLO2WWxjZxqT\nYi+Cqu2YR6dpvwq51y1+RD7GgR3NgoA0ABlofC2ez5yLXrklopGV7BS5mW2kzzMD\nncIb+XPJT81l5/Nbkmw5bnTHLjJTUAWyGsd3UnXGPyqPi823HaHWdwrCVrEdQfKC\n6sfbd0RMX625LMgK4cYl+z8BpcONLC2G+GEJf/yONonArW2EsExlRSWk4obKttb0\ngy/uro1rMhVVUSIn09XrjqdTYICQWgou8zWSfVeIFhU3KSNXrG5SeDucus+mZiyw\nXPcV7UVZAgMBAAECggEAX9EL1ieUSxzlapb7V++UNqkXWRcnJhJMfKn2ug+7WsnD\nO+IjVIm5plbpG+j9DAiYzKUZZGImyWN81K1+LK1mmliVZA+DMRxC0tiebhufxjC6\nRIo21DWcBA9JMGM13M0c70QvEczA7pl3bgrMq8+2K0F12j8q+LJPknMF9YmKIxQ9\nR3sFbKBzx5vHNn22pd+kZMWOxYXfErF3FiDwECRDziDZhi5WH4UnTqxIZ99nRC4N\nF80PwCZgCcF2/6k1d6gBNt6ZIwymn2ouEFSoD66UHRohw0Yt2e09AHl9tQoesx+/\nvuoVIum17S5fbPdAKg0Hpjz0T/3VajgrH5lgDeEqcQKBgQD5Q6AAPFVC4qUT95bl\nFnxjvL3ehg2IHY1aOFaOE2LkfQcTpG/OqbdxlWi1mhUzz/k1xtFeYCs/WFR+Dpbe\nkPp6WOp624q8JwJM+gJIMKNe1it/ViYaDl6lYR9zD0YH8euHrZyvcIgi4jkCeuzq\nBIMmUIp0Sf/qvibGTzwWuQUzPwKBgQDb7gpcuhULtURDwrDfy9dQbhLRIv95K7NA\nnDjj50KtzQAv3Lo2eURAWPK3CzzCsPcsW/KQbUjdVf8Xu7lql67CTx+ME8b2f6kb\nbyXvSDwbYPmbA1L4dEsM3027R8WDOUgJn4cq17FQPpoXR9uXAR+PAbbKvLGyqBXe\nOkIEyouZZwKBgD+mbBIDVgZJV7v8ijSfhE33oUhUVNpAKZszLa8D9knAP3FdmQtd\nvTEs5NsSqH1dixaXWVlPF1wKkzDJSu80eDGLyPxsWUXodCbx5GlIuj76U4sllX8r\n0jFK1rwL12cp+GniBFOsWacGu2YNu7eaRV8gS+qcBtSLj90t2Hs47cIDAoGAQ3A6\nvTRSP8TVqfGJSxa2b3NRUc/phr5fuPAugbPoPmMeK48DirCkvRaJVpebihe5s2B+\ngMsCDNzR2/U0ZfsdG3gntExcNjnvIPp0J2t/AuY5o87hIk6GtOvEaikX56Uo8cp5\nKCn7tR99IKZoL2Wox7E+2+wAkUf5bKtwkweIxwECgYBszEXz+GC7buwQSWG42st6\nEMi2BzSBq/WvFPFasv73xWn0joPh0KPXe6nIvMLONcKOHao7Wsnj98xYaIrbKapO\nOmto26vWyN61Ou53myhupfMT/hhyxln/AdVtJEtikPKEcwaKOIi6okbhqCyTcz76\nHnVHRTp+bAzjBEnoSxvTUw==\n-----END PRIVATE KEY-----\n"

		jwt.PrivateKey = os.Getenv("PrivateKey")
		jwt.Type = os.Getenv("Type")
		jwt.ProjectID = os.Getenv("ProjectID")
		jwt.PrivateKeyID = os.Getenv("PrivateKeyID")
		jwt.ClientEmail = os.Getenv("ClientEmail")
		jwt.ClientID = os.Getenv("ClientID")
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
