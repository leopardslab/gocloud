package auth

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration struct for representing AWS credentials
type Configuration struct {
	AWSAccessKeyID string
	AWSSecretKey   string
}

var Config Configuration

func AWSLoadConfig() {

// Read from file first.
	var home string = os.Getenv("HOME")
	file, _ := os.Open(home + "/amazoncloudconfig.json")

	// Defer the closing of our jsonFile so that we can parse it later on
	defer file.Close()

	// We initialize Configuration struct
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	_ = decoder.Decode(&Config)

	if Config.AWSAccessKeyID == "" || Config.AWSSecretKey == "" {
// If amazoncloudconfig.json doesn't exist, look for credentials as environment variables.

		Config.AWSAccessKeyID = os.Getenv("AWSAccessKeyID")
		Config.AWSSecretKey = os.Getenv("AWSSecretKey")

		if Config.AWSAccessKeyID == "" || Config.AWSSecretKey == "" {
			log.Fatalln("Cannot get access key and secret key")
		}
	}
}
