package aliauth

import (
	"os"
	"encoding/json"
	"log"
)

//Configuration struct reperesnts.
type Configuration struct {
	AliAccessKeyId     string
	AliAccessKeySecret string
}

var Config Configuration

func LoadConfig() {

	var home string = os.Getenv("HOME")
	file, _ := os.Open(home + "/gocloudconfig.json")

	decoder := json.NewDecoder(file)
	Config = Configuration{}
	_ = decoder.Decode(&Config)

	if Config.AliAccessKeyId == "" || Config.AliAccessKeySecret == "" {

		Config.AliAccessKeyId = os.Getenv("AliAccessKeyId")
		Config.AliAccessKeySecret = os.Getenv("AliAccessKeySecret")

		if Config.AliAccessKeyId == "" || Config.AliAccessKeySecret == "" {
			log.Fatalln("Cannot Get Ali access key and secret key")
		}
	}
}
