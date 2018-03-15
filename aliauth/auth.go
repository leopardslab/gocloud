package aliauth

import (
	"os"
	"encoding/json"
	"log"
)

//Configuration struct reperesnts.
type Configuration struct {
	AliAccessKeyID     string
	AliAccessKeySecret string
}

var Config Configuration

func LoadConfig() {
	// Read from file first.
	var home string = os.Getenv("HOME")
	file, _ := os.Open(home + "/gocloudconfig.json")

	decoder := json.NewDecoder(file)
	Config = Configuration{}
	_ = decoder.Decode(&Config)

	if Config.AliAccessKeyID == "" || Config.AliAccessKeySecret == "" {

		Config.AliAccessKeyID = os.Getenv("AliAccessKeyID")
		Config.AliAccessKeySecret = os.Getenv("AliAccessKeySecret")

		if Config.AliAccessKeyID == "" || Config.AliAccessKeySecret == "" {
			log.Fatalln("Cannot Get Ali access key and secret key")
		}
	}
}
