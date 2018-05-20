package aliauth

import (
	"os"
	"encoding/json"
	"log"
)

//Configuration struct for representing Ali-cloud credentials
type Configuration struct {
	AliAccessKeyID     string
	AliAccessKeySecret string
}

//Config from alicloudconfig.json
var Config Configuration

//LoadConfig represents Load Ali-cloud config from alicloudconfig.json or environment variables
func LoadConfig() {
	// Read from file first.
	var home string = os.Getenv("HOME")
	file, _ := os.Open(home + "/.gocloud" + "/alicloudconfig.json")

	// Defer the closing of our jsonFile so that we can parse it later on
	defer file.Close()

	// We initialize Configuration struct
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	_ = decoder.Decode(&Config)

	if Config.AliAccessKeyID == "" || Config.AliAccessKeySecret == "" {
		// If alicloudconfig.json doesn't exist, look for credentials as environment variables.
		Config.AliAccessKeyID = os.Getenv("AliAccessKeyID")
		Config.AliAccessKeySecret = os.Getenv("AliAccessKeySecret")

		if Config.AliAccessKeyID == "" || Config.AliAccessKeySecret == "" {
			log.Fatalln("Cannot Get Ali access key and secret key")
		}
	}
}
