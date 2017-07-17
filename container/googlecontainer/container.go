package googlecontainer

import (
	"bytes"
	"encoding/json"
	"fmt"
	googleauth "github.com/scorelab/gocloud-v2/googleauth"
	"io/ioutil"
	"net/http"
)

type Googlecontainer struct {
}

type Createcluster struct {
	Name                  string `json:"name"`
	Zone                  string `json:"zone"`
	Network               string `json:"network"`
	LoggingService        string `json:"loggingService"`
	MonitoringService     string `json:"monitoringService"`
	InitialClusterVersion string `json:"initialClusterVersion"`
	Subnetwork            string `json:"subnetwork"`
}

/*
type Createcluster struct {
	NodePools []struct {
		Name string `json:"name"`
		InitialNodeCount int `json:"initialNodeCount"`
		Config struct {
			MachineType string `json:"machineType"`
			ImageType string `json:"imageType"`
			DiskSizeGb int `json:"diskSizeGb"`
			Preemptible bool `json:"preemptible"`
			OauthScopes []string `json:"oauthScopes"`
		} `json:"config"`
		Autoscaling struct {
			Enabled bool `json:"enabled"`
		} `json:"autoscaling"`
		Management struct {
			AutoUpgrade bool `json:"autoUpgrade"`
			AutoRepair bool `json:"autoRepair"`
			UpgradeOptions struct {
			} `json:"upgradeOptions"`
		} `json:"management"`
	} `json:"nodePools"`

	MasterAuth struct {
		Username string `json:"username"`
		ClientCertificateConfig struct {
			IssueClientCertificate bool `json:"issueClientCertificate"`
		} `json:"clientCertificateConfig"`
	} `json:"masterAuth"`

	LegacyAbac struct {
		Enabled bool `json:"enabled"`
	} `json:"legacyAbac"`

}

*/

func (googlecontainer *Googlecontainer) Deletecontainer(request interface{}) (resp interface{}, err error) {

	return
}

func (googlecontainer *Googlecontainer) Createcontainer(request interface{}) (resp interface{}, err error) {

	var option Createcluster

	var Projectid string

	var Zone string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "projectid":
			Projectid, _ = value.(string)

		case "name":
			name, _ := value.(string)
			option.Name = name

		case "Zone":
			ZoneV, _ := value.(string)
			Zone = ZoneV

		case "network":
			NetworkV, _ := value.(string)
			option.Network = NetworkV

		case "loggingService":
			LoggingServiceV, _ := value.(string)
			option.LoggingService = LoggingServiceV

		case "monitoringService":
			MonitoringServiceV, _ := value.(string)
			option.MonitoringService = MonitoringServiceV

		case "initialClusterVersion":
			InitialClusterVersionV, _ := value.(string)
			option.InitialClusterVersion = InitialClusterVersionV

		case "subnetwork":
			SubnetworkV, _ := value.(string)
			option.Subnetwork = SubnetworkV
		}
	}

	zonevalue := "projects/" + Projectid + "/zones/" + Zone
	option.Zone = zonevalue

	Createclusterjsonmap := make(map[string]interface{})

	Createclusterdictnoaryconvert(option, Createclusterjsonmap)

	Createclusterjson, _ := json.Marshal(Createclusterjsonmap)

	Createclusterjsonstring := string(Createclusterjson)

	fmt.Println(Createclusterjsonstring)

	//var Createclusterjsonstringbyte = []byte(Createclusterjsonstring)
	var Createclusterjsonstringbyte = []byte(`{ "cluster": { "name": "cluster-1", "zone": "us-central1-a", "network": "default", "loggingService": "logging.googleapis.com", "monitoringService": "none", "nodePools": [ { "name": "default-pool", "initialNodeCount": 3, "config": { "machineType": "n1-standard-1", "imageType": "COS", "diskSizeGb": 100, "preemptible": false, "oauthScopes": [ "https://www.googleapis.com/auth/compute", "https://www.googleapis.com/auth/devstorage.read_only", "https://www.googleapis.com/auth/logging.write", "https://www.googleapis.com/auth/monitoring.write", "https://www.googleapis.com/auth/servicecontrol", "https://www.googleapis.com/auth/service.management.readonly", "https://www.googleapis.com/auth/trace.append" ] }, "autoscaling": { "enabled": false }, "management": { "autoUpgrade": false, "autoRepair": false, "upgradeOptions": {} } } ], "initialClusterVersion": "1.6.4", "masterAuth": { "username": "admin", "clientCertificateConfig": { "issueClientCertificate": true } }, "subnetwork": "default", "legacyAbac": { "enabled": true } } }`)

	url := "https://container.googleapis.com/v1/projects/" + Projectid + "/zones/" + Zone + "/clusters"

	client := googleauth.SignJWT()

	Createclusterrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(Createclusterjsonstringbyte))

	Createclusterrequest.Header.Set("Content-Type", "application/json")

	Createclusterresp, err := client.Do(Createclusterrequest)

	defer Createclusterresp.Body.Close()

	body, err := ioutil.ReadAll(Createclusterresp.Body)

	fmt.Println(string(body))

	return
}

func Createclusterdictnoaryconvert(option Createcluster, Createclusterjsonmap map[string]interface{}) {
	if option.Name != "" {
		Createclusterjsonmap["name"] = option.Name
	}
	if option.Network != "" {
		Createclusterjsonmap["network"] = option.Network
	}

	if option.LoggingService != "" {
		Createclusterjsonmap["loggingService"] = option.LoggingService
	}

	if option.MonitoringService != "" {
		Createclusterjsonmap["monitoringService"] = option.MonitoringService
	}

	if option.InitialClusterVersion != "" {
		Createclusterjsonmap["initialClusterVersion"] = option.InitialClusterVersion
	}

	if option.Zone != "" {
		Createclusterjsonmap["subnetwork"] = option.Zone
	}

}

func (googlecontainer *Googlecontainer) Createservice(request interface{}) (resp interface{}, err error) {
	return
}

func (googlecontainer *Googlecontainer) Runtask(request interface{}) (resp interface{}, err error) {
	return
}

func (googlecontainer *Googlecontainer) Deleteservice(request interface{}) (resp interface{}, err error) {
	return
}

func (googlecontainer *Googlecontainer) Stoptask(request interface{}) (resp interface{}, err error) {
	return
}

/*

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

*/
