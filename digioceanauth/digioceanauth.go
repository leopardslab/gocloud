package digioceanauth

import (
	"encoding/json"
	"log"
	"os"
)

// TokenSource struct for representing DigiOcean credentials.
type TokenSource struct {
	AccessToken string
}

// Token is a variable of type TokenSource.
var Token TokenSource

// LoadConfig loads the DigitalOcean credentials.
func LoadConfig() {

	// Read from file first.
	var home string = os.Getenv("HOME")
	file, _ := os.Open(home + "/digioceancloudconfig.json")

	// Defer the closing of our file so that we can parse it later on.
	defer file.Close()

	// We initialize TokenSource struct.
	decoder := json.NewDecoder(file)
	Token = TokenSource{}
	_ = decoder.Decode(&Token)

	if Token.AccessToken == "" {
		// If digioceancloudconfig.json doesn't exist, look for credentials as environment variables.

		Token.AccessToken = os.Getenv("DigiOceanAccessToken")
		if Token.AccessToken == "" {
			log.Fatalln("Cannot get access token for DigitalOcean.")
		}
	}
}
