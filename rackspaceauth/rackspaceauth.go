package rackspaceauth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// TokenSource struct for representing RackSpace credentials.
type TokenSource struct {
	RackSpaceTenantID   string
	RackSpaceTenantName string
	RackSpaceAuthToken  string
	RackSpaceAPIKey     string
	RackSpaceUsername   string
}

// Token is a variable of type TokenSource.
var Token TokenSource

// LoadConfigAndAuthenticate loads the RackSpace credentials.
func LoadConfigAndAuthenticate() {

	// Read from file first.
	var home = os.Getenv("HOME")
	file, _ := os.Open(home + "/.gocloud" + "/rackspacecloudconfig.json")

	// Defer the closing of our file so that we can parse it later on.
	defer file.Close()

	// We initialize TokenSource struct.
	decoder := json.NewDecoder(file)
	Token = TokenSource{}
	_ = decoder.Decode(&Token)

	if Token.RackSpaceAPIKey == "" || Token.RackSpaceUsername == "" {
		// If digioceancloudconfig.json doesn't exist, look for credentials as environment variables.

		Token.RackSpaceAPIKey = os.Getenv("RackSpaceAPIKey")
		Token.RackSpaceUsername = os.Getenv("RackSpaceUsername")
		if Token.RackSpaceAPIKey == "" || Token.RackSpaceUsername == "" {
			log.Fatalln("Cannot get access token for DigitalOcean.")
		}
	}

	// prepare the authoriation request
	url := "https://identity.api.rackspacecloud.com/v2.0/tokens"
	authRequestData := map[string]interface{}{
		"auth": map[string]interface{}{
			"RAX-KSKEY:apiKeyCredentials": map[string]interface{}{
				"username": Token.RackSpaceUsername,
				"apiKey":   Token.RackSpaceAPIKey,
			},
		},
	}
	jsonStr, err := json.Marshal(authRequestData)
	if err != nil {
		log.Fatalln("Failed to marshal json")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatalln("Failed to create the request")
	}
	req.Header.Set("Content-Type", "application/json")

	// make the authorization request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Failed to make the request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Failed to read the response body")
	}

	// Extract the Auth Token and the tenant details
	respJSON := make(map[string]interface{})
	err = json.Unmarshal(body, &respJSON)
	if err != nil {
		log.Fatalln("Failed to unmarshal response JSON")
	}
	accessJSON := respJSON["access"].(map[string]interface{})
	tokenJSON := accessJSON["token"].(map[string]interface{})
	Token.RackSpaceAuthToken = tokenJSON["id"].(string)
	tenantJSON := tokenJSON["tenant"].(map[string]interface{})
	Token.RackSpaceTenantID = tenantJSON["id"].(string)
	Token.RackSpaceTenantName = tenantJSON["name"].(string)
}
