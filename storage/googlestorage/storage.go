package googlestorage

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/scorelab/gocloud-v2/googleauth"
	"io/ioutil"
	"net/http"
)

func (googlestorage *GoogleStorage) Createdisk(request interface{}) (resp interface{}, err error) {

	//var googlestorage GoogleStorage
	var Projectid string
	var Zone string
	var Type string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "projectid":
			Projectid, _ = value.(string)
			fmt.Println("Projectid", Projectid)

		case "Name":
			name, _ := value.(string)
			googlestorage.Name = name

		case "Zone":
			ZoneV, _ := value.(string)
			Zone = ZoneV

		case "Type":
			TypeV, _ := value.(string)
			Type = TypeV

		default:
			fmt.Println("Incorrect Value")

		}
	}

	zonevalue := "projects/" + Projectid + "/zones/" + Zone
	googlestorage.Zone = zonevalue

	Typevalue := "projects/" + Projectid + "/zones/" + Zone + "/diskTypes/" + Type
	googlestorage.Type = Typevalue

	fmt.Println(googlestorage.Name)

	googlestoragejson, _ := json.Marshal(googlestorage)
	googlestoragejsonstring := string(googlestoragejson)
	fmt.Println(googlestoragejsonstring)
	var googlestoragejsonstringbyte = []byte(googlestoragejsonstring)

	url := "https://www.googleapis.com/compute/v1/projects/" + Projectid + "/zones/" + Zone + "/disks"

	fmt.Println(url)

	client := &http.Client{}
	Creatediskrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(googlestoragejsonstringbyte))
	Creatediskrequest.Header.Set("Content-Type", "application/json")

	token := googleauth.Sign()

	token.SetAuthHeader(Creatediskrequest)

	Creatediskresp, err := client.Do(Creatediskrequest)
	defer Creatediskresp.Body.Close()

	body, err := ioutil.ReadAll(Creatediskresp.Body)
	fmt.Println(string(body))

	return
}

func (googlestorage *GoogleStorage) Deletedisk(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/disks/" + options["disk"]

	token := googleauth.Sign()

	client := &http.Client{}

	Deletediskrequest, err := http.NewRequest("DELETE", url, nil)
	Deletediskrequest.Header.Set("Content-Type", "application/json")

	token.SetAuthHeader(Deletediskrequest)
	Deletediskresp, err := client.Do(Deletediskrequest)

	defer Deletediskresp.Body.Close()
	body, err := ioutil.ReadAll(Deletediskresp.Body)

	fmt.Println(string(body))

	return
}






func (googlestorage *GoogleStorage) createSnapshot(request interface{}) (resp interface{}, err error) {

	//var googlestorage GoogleStorage
	var Projectid string
	var Zone string
	var Type string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "projectid":
			Projectid, _ = value.(string)
			fmt.Println("Projectid", Projectid)

		case "Name":
			name, _ := value.(string)
			googlestorage.Name = name

		case "Zone":
			ZoneV, _ := value.(string)
			Zone = ZoneV

		case "Type":
			TypeV, _ := value.(string)
			Type = TypeV

		default:
			fmt.Println("Incorrect Value")

		}
	}

	zonevalue := "projects/" + Projectid + "/zones/" + Zone
	googlestorage.Zone = zonevalue

	Typevalue := "projects/" + Projectid + "/zones/" + Zone + "/diskTypes/" + Type
	googlestorage.Type = Typevalue

	fmt.Println(googlestorage.Name)

	googlestoragejson, _ := json.Marshal(googlestorage)
	googlestoragejsonstring := string(googlestoragejson)
	fmt.Println(googlestoragejsonstring)
	var googlestoragejsonstringbyte = []byte(googlestoragejsonstring)

	url := "https://www.googleapis.com/compute/v1/projects/" + Projectid + "/zones/" + Zone + "/createSnapshot"

	fmt.Println(url)

	client := &http.Client{}
	Creatediskrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(googlestoragejsonstringbyte))
	Creatediskrequest.Header.Set("Content-Type", "application/json")

	token := googleauth.Sign()

	token.SetAuthHeader(Creatediskrequest)

	Creatediskresp, err := client.Do(Creatediskrequest)
	defer Creatediskresp.Body.Close()

	body, err := ioutil.ReadAll(Creatediskresp.Body)
	fmt.Println(string(body))

	return
}
