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

func (googlestorage *GoogleStorage) Createsnapshot(request interface{}) (resp interface{}, err error) {

	var snapshot Snapshot
	var Projectid string
	var Zone string
	var Disk string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "projectid":
			Projectid, _ = value.(string)
			fmt.Println("Projectid", Projectid)

		case "Name":
			name, _ := value.(string)
			snapshot.Name = name

		case "Zone":
			ZoneV, _ := value.(string)
			Zone = ZoneV

		case "disk":
			DiskV, _ := value.(string)
			Disk = DiskV
		default:
			fmt.Println("Incorrect Value")

		}
	}

	fmt.Println(snapshot)

	snapshotjson, _ := json.Marshal(snapshot)
	snapshotjsonstring := string(snapshotjson)
	fmt.Println(snapshotjsonstring)
	var snapshotjsonstringbyte = []byte(snapshotjsonstring)

	//POST https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk/createSnapshot

	url := "https://www.googleapis.com/compute/v1/projects/" + Projectid + "/zones/" + Zone + "/disks/" + Disk + "/createSnapshot"

	fmt.Println(url)

	client := &http.Client{}
	Createsnapshotrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(snapshotjsonstringbyte))
	Createsnapshotrequest.Header.Set("Content-Type", "application/json")

	token := googleauth.Sign()

	token.SetAuthHeader(Createsnapshotrequest)

	Createsnapshotresp, err := client.Do(Createsnapshotrequest)
	defer Createsnapshotresp.Body.Close()

	body, err := ioutil.ReadAll(Createsnapshotresp.Body)
	fmt.Println(string(body))

	return
}

func (googlestorage *GoogleStorage) Deletesnapshot(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/global/snapshots/" + options["snapshot"]

	fmt.Println(url)

	token := googleauth.Sign()

	client := &http.Client{}

	Deletesnapshotrequest, err := http.NewRequest("DELETE", url, nil)
	Deletesnapshotrequest.Header.Set("Content-Type", "application/json")

	token.SetAuthHeader(Deletesnapshotrequest)

	Deletesnapshotresp, err := client.Do(Deletesnapshotrequest)

	defer Deletesnapshotresp.Body.Close()
	body, err := ioutil.ReadAll(Deletesnapshotresp.Body)

	fmt.Println(string(body))

	return
}

func (googlestorage *GoogleStorage) Attachdisk(request interface{}) (resp interface{}, err error) {

	var attachdisk Attachdisk
	var Projectid string
	var Zone string
	var Instance string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "projectid":
			Projectid, _ = value.(string)
			fmt.Println("Projectid", Projectid)

		case "Zone":
			ZoneV, _ := value.(string)
			Zone = ZoneV

		case "Source":
			source, _ := value.(string)
			attachdisk.Source = source

		case "instance":
			instanceV, _ := value.(string)
			Instance = instanceV

		default:
			fmt.Println("Incorrect Value")

		}
	}

	fmt.Println(attachdisk)

	attachdiskjson, _ := json.Marshal(attachdisk)
	attachdiskjsonstring := string(attachdiskjson)
	fmt.Println(attachdiskjsonstring)
	var attachdiskjsonstringbyte = []byte(attachdiskjsonstring)

	url := "https://www.googleapis.com/compute/v1/projects/" + Projectid + "/zones/" + Zone + "/instances/" + Instance + "/attachDisk"

	fmt.Println(url)

	client := &http.Client{}
	attachdiskrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(attachdiskjsonstringbyte))
	attachdiskrequest.Header.Set("Content-Type", "application/json")

	token := googleauth.Sign()

	token.SetAuthHeader(attachdiskrequest)

	attachdiskresp, err := client.Do(attachdiskrequest)
	defer attachdiskresp.Body.Close()

	body, err := ioutil.ReadAll(attachdiskresp.Body)
	fmt.Println(string(body))

	return
}

func (googlestorage *GoogleStorage) Detachdisk(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/instances/" + options["instance"] + "/detachDisk"

	fmt.Println(url)

	token := googleauth.Sign()

	client := &http.Client{}

	detachdiskrequest, err := http.NewRequest("POST", url, nil)

	q := detachdiskrequest.URL.Query()

	q.Add("deviceName", options["deviceName"])

	detachdiskrequest.URL.RawQuery = q.Encode()

	detachdiskrequest.Header.Set("Content-Type", "application/json")

	token.SetAuthHeader(detachdiskrequest)

	fmt.Println(detachdiskrequest)

	detachdiskresp, err := client.Do(detachdiskrequest)

	defer detachdiskresp.Body.Close()

	body, err := ioutil.ReadAll(detachdiskresp.Body)

	fmt.Println(string(body))

	return
}
