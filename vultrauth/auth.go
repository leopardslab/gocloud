package vultrauth

import (
	"encoding/json"
	"log"
	"os"
)

//Configuration struct for representing Vultr API Key
type Configuration struct {
	VultrAPIKey string
}

//Config from vultrconfig.json
var Config Configuration

//LoadConfig represents Load Vultr config from vultrconfig.json or environment variables
func LoadConfig() {
	// Read from file first.
	var home string = os.Getenv("HOME")
	file, _ := os.Open(home + "/.gocloud" + "/vultrconfig.json")

	// Defer the closing of our jsonFile so that we can parse it later on
	defer file.Close()

	// We initialize Configuration struct
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	_ = decoder.Decode(&Config)

	if Config.VultrAPIKey == "" {
		// If vultrconfig.json doesn't exist, look for credentials as environment variables.
		Config.VultrAPIKey = os.Getenv("VultrAPIKey")

		if Config.VultrAPIKey == "" {
			log.Fatalln("Cannot Get Vultr API Key")
		}
	}
}
