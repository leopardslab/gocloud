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

	var option Creatdisk

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
			option.Name = name

		case "Zone":
			ZoneV, _ := value.(string)
			Zone = ZoneV

		case "Type":
			TypeV, _ := value.(string)
			Type = TypeV

		case "SizeGb":
			SizeGbV, _ := value.(string)
			option.SizeGb = SizeGbV

		case "SourceImageEncryptionKeys":
			SourceImageEncryptionKeysV, _ := value.(map[string]string)
			option.SourceImageEncryptionKeys.RawKey = SourceImageEncryptionKeysV["RawKey"]
			option.SourceImageEncryptionKeys.Sha256 = SourceImageEncryptionKeysV["Sha256"]

		case "DiskEncryptionKeys":
			DiskEncryptionKeysV, _ := value.(map[string]string)
			option.DiskEncryptionKeys.RawKey = DiskEncryptionKeysV["RawKey"]
			option.DiskEncryptionKeys.Sha256 = DiskEncryptionKeysV["Sha256"]

		case "SourceSnapshotEncryptionKeys":
			SourceSnapshotEncryptionKeysV, _ := value.(map[string]string)
			option.SourceSnapshotEncryptionKeys.RawKey = SourceSnapshotEncryptionKeysV["RawKey"]
			option.SourceSnapshotEncryptionKeys.Sha256 = SourceSnapshotEncryptionKeysV["Sha256"]

		case "Licenses":
			LicensesV, _ := value.([]string)
			option.Licenses = LicensesV

		case "Users":
			UsersV, _ := value.([]string)
			option.Users = UsersV

		case "CreationTimestamp":
			CreationTimestampV, _ := value.(string)
			option.CreationTimestamp = CreationTimestampV

		case "Description":
			DescriptionV, _ := value.(string)
			option.Description = DescriptionV

		case "ID":
			IDV, _ := value.(string)
			option.ID = IDV

		case "Kind":
			KindV, _ := value.(string)
			option.Kind = KindV

		case "LabelFingerprint":
			LabelFingerprintV, _ := value.(string)
			option.LabelFingerprint = LabelFingerprintV

		case "SourceSnapshotID":
			SourceSnapshotIDV, _ := value.(string)
			option.SourceSnapshotID = SourceSnapshotIDV

		case "Status":
			StatusV, _ := value.(string)
			option.Status = StatusV

		case "LastAttachTimestamp":
			LastAttachTimestampV, _ := value.(string)
			option.LastAttachTimestamp = LastAttachTimestampV

		case "LastDetachTimestamp":
			LastDetachTimestampV, _ := value.(string)
			option.LastDetachTimestamp = LastDetachTimestampV

		case "Options":
			OptionsV, _ := value.(string)
			option.Options = OptionsV

		case "SelfLink":
			SelfLinkV, _ := value.(string)
			option.SelfLink = SelfLinkV

		case "SourceImage":
			SourceImage, _ := value.(string)
			option.SourceImage = SourceImage

		case "SourceImageID":
			SourceImageIDV, _ := value.(string)
			option.SourceImageID = SourceImageIDV

		case "SourceSnapshot":
			SourceSnapshotV, _ := value.(string)
			option.SourceSnapshot = SourceSnapshotV

		default:
			fmt.Println("Incorrect Value")

		}
	}

	zonevalue := "projects/" + Projectid + "/zones/" + Zone
	option.Zone = zonevalue

	Typevalue := "projects/" + Projectid + "/zones/" + Zone + "/diskTypes/" + Type
	option.Type = Typevalue

	Creatdiskjsonmap := make(map[string]interface{})

	Creatediskdictnoaryconvert(option, Creatdiskjsonmap)

	//fmt.Println("Creatdiskjsonmap :\n", Creatdiskjsonmap)

	Creatdiskjson, _ := json.Marshal(Creatdiskjsonmap)

	Creatdiskjsonstring := string(Creatdiskjson)

	fmt.Println("Creatdiskjsonstring:\n", Creatdiskjsonstring)

	var Creatdiskjsonstringbyte = []byte(Creatdiskjsonstring)

	url := "https://www.googleapis.com/compute/v1/projects/" + Projectid + "/zones/" + Zone + "/disks"

	fmt.Println(url)

	client := &http.Client{}

	Creatediskrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Creatdiskjsonstringbyte))

	Creatediskrequest.Header.Set("Content-Type", "application/json")

	token := googleauth.Sign()

	token.SetAuthHeader(Creatediskrequest)

	Creatediskresp, err := client.Do(Creatediskrequest)

	defer Creatediskresp.Body.Close()

	body, err := ioutil.ReadAll(Creatediskresp.Body)

	fmt.Println(string(body))

	return
}

func Creatediskdictnoaryconvert(option Creatdisk, Creatdiskjsonmap map[string]interface{}) {

	if option.Name != "" {
		Creatdiskjsonmap["name"] = option.Name
	}

	if option.Type != "" {
		Creatdiskjsonmap["type"] = option.Type
	}

	if option.Zone != "" {
		Creatdiskjsonmap["zone"] = option.Zone
	}

	if option.SizeGb != "" {
		Creatdiskjsonmap["sizeGb"] = option.SizeGb
	}

	if option.SourceImageEncryptionKeys != (SourceImageEncryptionKey{}) {
		Creatdiskjsonmap["sourceImageEncryptionKey"] = option.SourceImageEncryptionKeys
	}
	if option.DiskEncryptionKeys != (DiskEncryptionKey{}) {
		Creatdiskjsonmap["diskEncryptionKey"] = option.DiskEncryptionKeys
	}
	if option.SourceSnapshotEncryptionKeys != (SourceSnapshotEncryptionKey{}) {
		Creatdiskjsonmap["sourceSnapshotEncryptionKey"] = option.SourceSnapshotEncryptionKeys
	}

	if len(option.Licenses) != 0 {
		Creatdiskjsonmap["licenses"] = option.Licenses
	}

	if len(option.Users) != 0 {
		Creatdiskjsonmap["users"] = option.Users
	}

	if option.LastAttachTimestamp != "" {
		Creatdiskjsonmap["lastAttachTimestamp"] = option.LastAttachTimestamp
	}

	if option.SourceSnapshot != "" {
		Creatdiskjsonmap["SourceSnapshot"] = option.SourceSnapshot
	}

	if option.LastDetachTimestamp != "" {
		Creatdiskjsonmap["lastDetachTimestamp"] = option.LastDetachTimestamp
	}

	if option.Options != "" {
		Creatdiskjsonmap["options"] = option.Options
	}

	if option.SelfLink != "" {
		Creatdiskjsonmap["selfLink"] = option.SelfLink
	}

	if option.SourceImage != "" {
		Creatdiskjsonmap["sourceImage"] = option.SourceImage
	}

	if option.SourceImageID != "" {
		Creatdiskjsonmap["sourceImageID"] = option.SourceImageID
	}

	if option.Status != "" {
		Creatdiskjsonmap["status"] = option.Status
	}

	if option.SourceSnapshotID != "" {
		Creatdiskjsonmap["sourceSnapshotID"] = option.SourceSnapshotID
	}

	if option.LabelFingerprint != "" {
		Creatdiskjsonmap["labelFingerprint"] = option.LabelFingerprint
	}

	if option.Kind != "" {
		Creatdiskjsonmap["kind"] = option.Kind
	}

	if option.ID != "" {
		Creatdiskjsonmap["id"] = option.ID
	}

	if option.Description != "" {
		Creatdiskjsonmap["description"] = option.Description
	}

	if option.CreationTimestamp != "" {
		Creatdiskjsonmap["creationTimestamp"] = option.CreationTimestamp
	}

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
			NameV, _ := value.(string)
			snapshot.Name = NameV

		case "Zone":
			ZoneV, _ := value.(string)
			Zone = ZoneV

		case "disk":
			DiskV, _ := value.(string)
			Disk = DiskV

		case "CreationTimestamp":
			CreationTimestampV, _ := value.(string)
			snapshot.CreationTimestamp = CreationTimestampV

		case "Description":
			DescriptionV, _ := value.(string)
			snapshot.Description = DescriptionV

		case "DiskSizeGb":
			DiskSizeGbV, _ := value.(string)
			snapshot.DiskSizeGb = DiskSizeGbV

		case "ID":
			IDV, _ := value.(string)
			snapshot.ID = IDV

		case "Kind":
			KindV, _ := value.(string)
			snapshot.Kind = KindV

		case "LabelFingerprint":
			LabelFingerprintV, _ := value.(string)
			snapshot.LabelFingerprint = LabelFingerprintV

		case "SelfLink":
			SelfLinkV, _ := value.(string)
			snapshot.SelfLink = SelfLinkV

		case "StorageBytes":
			StorageBytesV, _ := value.(string)
			snapshot.StorageBytes = StorageBytesV

		case "Status":
			StatusV, _ := value.(string)
			snapshot.Status = StatusV

		case "SourceDiskID":
			SourceDiskIDV, _ := value.(string)
			snapshot.SourceDiskID = SourceDiskIDV

		case "SourceDisk":
			SourceDiskV, _ := value.(string)
			snapshot.SourceDisk = SourceDiskV

		case "StorageBytesStatus":
			StorageBytesStatusV, _ := value.(string)
			snapshot.StorageBytesStatus = StorageBytesStatusV

		case "Licenses":
			LicensesV, _ := value.([]string)
			snapshot.Licenses = LicensesV

		case "SourceDiskEncryptionKeys":
			SourceDiskEncryptionKeysV, _ := value.(map[string]string)
			snapshot.SourceDiskEncryptionKeys.RawKey = SourceDiskEncryptionKeysV["RawKey"]
			snapshot.SourceDiskEncryptionKeys.Sha256 = SourceDiskEncryptionKeysV["Sha256"]

		case "SnapshotEncryptionKeys":
			SnapshotEncryptionKeysV, _ := value.(map[string]string)
			snapshot.SnapshotEncryptionKeys.RawKey = SnapshotEncryptionKeysV["RawKey"]
			snapshot.SnapshotEncryptionKeys.Sha256 = SnapshotEncryptionKeysV["Sha256"]

		default:
			fmt.Println("Incorrect Value")

		}
	}

	fmt.Println(snapshot)

	Createsnapshotjsonmap := make(map[string]interface{})

	Createsnapshotdictnoaryconvert(snapshot, Createsnapshotjsonmap)

	Createsnapshotjson, _ := json.Marshal(Createsnapshotjsonmap)
	Createsnapshotstring := string(Createsnapshotjson)
	fmt.Println(Createsnapshotstring)
	var Createsnapshotstringbyte = []byte(Createsnapshotstring)

	//POST https://www.googleapis.com/compute/v1/projects/project/zones/zone/disks/disk/createSnapshot

	url := "https://www.googleapis.com/compute/v1/projects/" + Projectid + "/zones/" + Zone + "/disks/" + Disk + "/createSnapshot"

	fmt.Println(url)

	client := &http.Client{}
	Createsnapshotrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Createsnapshotstringbyte))
	Createsnapshotrequest.Header.Set("Content-Type", "application/json")

	token := googleauth.Sign()

	token.SetAuthHeader(Createsnapshotrequest)

	Createsnapshotresp, err := client.Do(Createsnapshotrequest)
	defer Createsnapshotresp.Body.Close()

	body, err := ioutil.ReadAll(Createsnapshotresp.Body)
	fmt.Println(string(body))

	return
}

func Createsnapshotdictnoaryconvert(option Snapshot, Createsnapshotjsonmap map[string]interface{}) {

	if len(option.Licenses) != 0 {
		Createsnapshotjsonmap["licenses"] = option.Licenses
	}

	if option.Name != "" {
		Createsnapshotjsonmap["name"] = option.Name
	}

	if option.CreationTimestamp != "" {
		Createsnapshotjsonmap["creationTimestamp"] = option.CreationTimestamp
	}

	if option.Description != "" {
		Createsnapshotjsonmap["description"] = option.Description
	}

	if option.DiskSizeGb != "" {
		Createsnapshotjsonmap["diskSizeGb"] = option.DiskSizeGb
	}

	if option.ID != "" {
		Createsnapshotjsonmap["id"] = option.ID
	}

	if option.Kind != "" {
		Createsnapshotjsonmap["kind"] = option.Kind
	}

	if option.Status != "" {
		Createsnapshotjsonmap["status"] = option.Status
	}

	if option.SelfLink != "" {
		Createsnapshotjsonmap["selfLink"] = option.SelfLink
	}

	if option.LabelFingerprint != "" {
		Createsnapshotjsonmap["labelFingerprint"] = option.LabelFingerprint
	}

	if option.StorageBytes != "" {
		Createsnapshotjsonmap["storageBytes"] = option.StorageBytes
	}

	if option.SourceDisk != "" {
		Createsnapshotjsonmap["sourceDisk"] = option.SourceDisk
	}

	if option.SourceDiskID != "" {
		Createsnapshotjsonmap["sourceDiskID"] = option.SourceDiskID
	}

	if option.StorageBytesStatus != "" {
		Createsnapshotjsonmap["storageBytesStatus"] = option.StorageBytesStatus
	}

	if option.SourceDiskEncryptionKeys != (SourceDiskEncryptionKey{}) {
		Createsnapshotjsonmap["sourceDiskEncryptionKey"] = option.SourceDiskEncryptionKeys
	}

	if option.SnapshotEncryptionKeys != (SnapshotEncryptionKey{}) {
		Createsnapshotjsonmap["snapshotEncryptionKey"] = option.SnapshotEncryptionKeys
	}

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

		case "Licenses":
			LicensesV, _ := value.([]string)
			attachdisk.Licenses = LicensesV

		case "DiskEncryptionKeys":
			DiskEncryptionKeysV, _ := value.(map[string]string)
			attachdisk.DiskEncryptionKeys.RawKey = DiskEncryptionKeysV["RawKey"]
			attachdisk.DiskEncryptionKeys.Sha256 = DiskEncryptionKeysV["Sha256"]


		case "Mode":
			ModeV, _ := value.(string)
			attachdisk.Mode = ModeV

		case "Type":
			TypeV, _ := value.(string)
			attachdisk.Type  =TypeV

		case "Kind":
			KindV, _ := value.(string)
			attachdisk.Kind = KindV


		case "Interface":
			InterfaceV, _ := value.(string)
			attachdisk.Interface = InterfaceV

		case "Index":
			IndexV, _ := value.(int)
			attachdisk.Index = IndexV

		case "Boot":
			BootV, _ := value.(bool)
			attachdisk.Boot = BootV

		case "AutoDelete":
			AutoDeleteV, _ := value.(bool)
			attachdisk.AutoDelete = AutoDeleteV

		case "DeviceName":
			DeviceNameV, _ := value.(string)
			attachdisk.DeviceName = DeviceNameV

		default:
			fmt.Println("Incorrect Value")

		}
	}

	//fmt.Println(attachdisk)

	Attachdiskjsonmap := make(map[string]interface{})

	Attachdiskdictnoaryconvert(attachdisk, Attachdiskjsonmap)

	attachdiskjson, _ := json.Marshal(Attachdiskjsonmap)

	attachdiskjsonstring := string(attachdiskjson)

	fmt.Println("attachdiskjsonstring", attachdiskjsonstring)

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

func Attachdiskdictnoaryconvert(option Attachdisk, Attachdiskjsonmap map[string]interface{}) {
	if len(option.Licenses) != 0 {
		Attachdiskjsonmap["licenses"] = option.Licenses
	}

	if option.DiskEncryptionKeys != (DiskEncryptionKey{}) {
		Attachdiskjsonmap["diskEncryptionKey"] = option.DiskEncryptionKeys
	}



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
