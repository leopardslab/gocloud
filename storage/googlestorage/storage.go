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

	Creatdiskjson, _ := json.Marshal(Creatdiskjsonmap)

	Creatdiskjsonstring := string(Creatdiskjson)

	var Creatdiskjsonstringbyte = []byte(Creatdiskjsonstring)

	url := "https://www.googleapis.com/compute/v1/projects/" + Projectid + "/zones/" + Zone + "/disks"

	client := googleauth.SignJWT()

	Creatediskrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Creatdiskjsonstringbyte))

	Creatediskrequest.Header.Set("Content-Type", "application/json")

	Creatediskresp, err := client.Do(Creatediskrequest)

	defer Creatediskresp.Body.Close()

	body, err := ioutil.ReadAll(Creatediskresp.Body)

	fmt.Println(string(body))

	return
}

func (googlestorage *GoogleStorage) Deletedisk(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/disks/" + options["disk"]

	client := googleauth.SignJWT()

	Deletediskrequest, err := http.NewRequest("DELETE", url, nil)
	Deletediskrequest.Header.Set("Content-Type", "application/json")

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

	Createsnapshotjsonmap := make(map[string]interface{})

	Createsnapshotdictnoaryconvert(snapshot, Createsnapshotjsonmap)

	Createsnapshotjson, _ := json.Marshal(Createsnapshotjsonmap)
	Createsnapshotstring := string(Createsnapshotjson)

	var Createsnapshotstringbyte = []byte(Createsnapshotstring)

	url := "https://www.googleapis.com/compute/v1/projects/" + Projectid + "/zones/" + Zone + "/disks/" + Disk + "/createSnapshot"

	client := googleauth.SignJWT()

	Createsnapshotrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Createsnapshotstringbyte))
	Createsnapshotrequest.Header.Set("Content-Type", "application/json")

	Createsnapshotresp, err := client.Do(Createsnapshotrequest)

	defer Createsnapshotresp.Body.Close()

	body, err := ioutil.ReadAll(Createsnapshotresp.Body)
	fmt.Println(string(body))

	return
}

func (googlestorage *GoogleStorage) Deletesnapshot(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/global/snapshots/" + options["snapshot"]

	client := googleauth.SignJWT()

	Deletesnapshotrequest, err := http.NewRequest("DELETE", url, nil)
	Deletesnapshotrequest.Header.Set("Content-Type", "application/json")

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
			attachdisk.Type = TypeV

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

		case "InitializeParam":
			InitializeParamV, _ := value.(map[string]interface{})
			fmt.Println("InitializeParamV", InitializeParamV)
			for key, value := range InitializeParamV {
				switch key {
				case "DiskName":
					DiskNameV, _ := value.(string)
					attachdisk.InitializeParam.DiskName = DiskNameV

				case "DiskType":
					DiskTypeV, _ := value.(string)
					attachdisk.InitializeParam.DiskType = DiskTypeV

				case "DiskSizeGb":
					DiskSizeGbV, _ := value.(string)
					attachdisk.InitializeParam.DiskSizeGb = DiskSizeGbV

				case "SourceImage":
					SourceImageV, _ := value.(string)
					attachdisk.InitializeParam.SourceImage = SourceImageV

				case "SourceImageEncryptionKeys":
					SourceImageEncryptionKeysV, _ := value.(map[string]string)
					attachdisk.InitializeParam.SourceImageEncryptionKeys.RawKey = SourceImageEncryptionKeysV["RawKey"]
					attachdisk.InitializeParam.SourceImageEncryptionKeys.Sha256 = SourceImageEncryptionKeysV["Sha256"]
				}
			}

		default:
			fmt.Println("Incorrect Value")

		}
	}

	Attachdiskjsonmap := make(map[string]interface{})

	Attachdiskdictnoaryconvert(attachdisk, Attachdiskjsonmap)

	attachdiskjson, _ := json.Marshal(Attachdiskjsonmap)

	attachdiskjsonstring := string(attachdiskjson)

	var attachdiskjsonstringbyte = []byte(attachdiskjsonstring)

	url := "https://www.googleapis.com/compute/v1/projects/" + Projectid + "/zones/" + Zone + "/instances/" + Instance + "/attachDisk"

	client := googleauth.SignJWT()

	attachdiskrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(attachdiskjsonstringbyte))
	attachdiskrequest.Header.Set("Content-Type", "application/json")

	attachdiskresp, err := client.Do(attachdiskrequest)
	defer attachdiskresp.Body.Close()

	body, err := ioutil.ReadAll(attachdiskresp.Body)
	fmt.Println(string(body))

	return
}

func (googlestorage *GoogleStorage) Detachdisk(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/instances/" + options["instance"] + "/detachDisk"

	client := googleauth.SignJWT()

	detachdiskrequest, err := http.NewRequest("POST", url, nil)

	detachdiskrequestparam := detachdiskrequest.URL.Query()

	detachdiskrequestparam.Add("deviceName", options["deviceName"])

	detachdiskrequest.URL.RawQuery = detachdiskrequestparam.Encode()

	detachdiskrequest.Header.Set("Content-Type", "application/json")

	detachdiskresp, err := client.Do(detachdiskrequest)

	defer detachdiskresp.Body.Close()

	body, err := ioutil.ReadAll(detachdiskresp.Body)

	fmt.Println(string(body))

	return
}
