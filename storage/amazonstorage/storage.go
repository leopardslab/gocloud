package amazonstorage

import (
	"encoding/xml"
	auth "github.com/scorelab/gocloud-v2/auth"
	awsauth "github.com/scorelab/gocloud-v2/awsauth"
	//"log"
	"net/http"
	//"net/http/httputil"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

type Amazonstorage struct {
}

const (
	debug = false

	legacyAPIVersion = "2011-12-15"

	vpcAPIVersion = "2013-10-15"
)

type Region struct {
	Name        string
	EC2Endpoint string
}

var USEast = Region{
	"us-east-1",
	"https://ec2.us-east-1.amazonaws.com",
}

type CreateVolume struct {
	AvailZone  string
	SnapshotId string
	VolumeType string
	VolumeSize int
	Encrypted  bool
	IOPS       int64
}

type Volume struct {
	Id          string             `xml:"volumeId"`
	Size        int                `xml:"size"`
	SnapshotId  string             `xml:"snapshotId"`
	Status      string             `xml:"status"`
	IOPS        int64              `xml:"iops"`
	AvailZone   string             `xml:"availabilityZone"`
	CreateTime  string             `xml:"createTime"`
	VolumeType  string             `xml:"volumeType"`
	Encrypted   bool               `xml:"encrypted"`
	Tags        []Tag              `xml:"tagSet>item"`
	Attachments []VolumeAttachment `xml:"attachmentSet>item"`
}

type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

type VolumeAttachment struct {
	VolumeId            string `xml:"volumeId"`
	Device              string `xml:"device"`
	InstanceId          string `xml:"instanceId"`
	Status              string `xml:"status"`
	DeleteOnTermination bool   `xml:"deleteOnTermination"`
}

type CreateVolumeResp struct {
	RequestId string `xml:"requestId"`
	Volume
}

func prepareVolume(params map[string]string, volume CreateVolume) {
	params["AvailabilityZone"] = volume.AvailZone
	if volume.SnapshotId != "" {
		params["SnapshotId"] = volume.SnapshotId
	}
	if volume.VolumeType != "" {
		params["VolumeType"] = volume.VolumeType
	}
	if volume.VolumeSize > 0 {
		params["Size"] = strconv.FormatInt(int64(volume.VolumeSize), 10)
	}
	if volume.IOPS > 0 {
		params["Iops"] = strconv.FormatInt(volume.IOPS, 10)
	}
	if volume.Encrypted {
		params["Encrypted"] = "true"
	}
}

type SimpleResp struct {
	XMLName   xml.Name
	RequestId string `xml:"requestId"`
}

type VolumeAttachmentResp struct {
	RequestId  string `xml:"requestId"`
	VolumeId   string `xml:"volumeId"`
	Device     string `xml:"device"`
	InstanceId string `xml:"instanceId"`
	Status     string `xml:"status"`
	AttachTime string `xml:"attachTime"`
}

func makeParams(action string) map[string]string {
	return makeParamsWithVersion(action, legacyAPIVersion)
}

func makeParamsVPC(action string) map[string]string {
	return makeParamsWithVersion(action, vpcAPIVersion)
}

//add version to params
func makeParamsWithVersion(action, version string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	params["Version"] = version
	return params
}

func (amazonstorage *Amazonstorage) PrepareSignatureV2query(params map[string]string,Region string, resp interface{}) error {

	EC2Endpoint := "https://ec2." + Region + ".amazonaws.com"

	req, err := http.NewRequest("GET", EC2Endpoint , nil)
	if err != nil {
		return err
	}

	// Add the params passed in to the query string
	query := req.URL.Query()
	for varName, varVal := range params {
		query.Add(varName, varVal)
	}
	query.Add("Timestamp", awsauth.TimeNow().In(time.UTC).Format(time.RFC3339))
	req.URL.RawQuery = query.Encode()

	auth := map[string]string{"AccessKey": auth.Config.AWSAccessKeyID, "SecretKey": auth.Config.AWSSecretKey}

	awsauth.SignV2(req, auth)

	fmt.Println(req)

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	fmt.Println(string(body))

	return xml.NewDecoder(r.Body).Decode(resp)
}

type SnapshotsResp struct {
	RequestId string     `xml:"requestId"`
	Snapshots []Snapshot `xml:"snapshotSet>item"`
}

type Snapshot struct {
	Id          string `xml:"snapshotId"`
	VolumeId    string `xml:"volumeId"`
	VolumeSize  string `xml:"volumeSize"`
	Status      string `xml:"status"`
	StartTime   string `xml:"startTime"`
	Description string `xml:"description"`
	Progress    string `xml:"progress"`
	OwnerId     string `xml:"ownerId"`
	OwnerAlias  string `xml:"ownerAlias"`
	Tags        []Tag  `xml:"tagSet>item"`
}

type CreateSnapshotResp struct {
	RequestId string `xml:"requestId"`
	Snapshot
}

func (amazonstorage *Amazonstorage) Createdisk(request interface{}) (resp interface{}, err error) {

	param, _ := request.(map[string]interface{})

	fmt.Println(param)

	var createvolume CreateVolume
	var Region string
	for key, value := range param {
		switch key {

		case "Region":
			regionV, _ := value.(string)
			Region = regionV

		case "AvailZone":
			AvailZoneV, _ := value.(string)
			createvolume.AvailZone = AvailZoneV

		case "VolumeType":
			VolumeTypeV, _ := value.(string)
			createvolume.VolumeType = VolumeTypeV

		case "VolumeSize":
			VolumeSizeV, _ := value.(int)
			createvolume.VolumeSize = VolumeSizeV

		case "IOPS":
			IOPSV, _ := value.(int64)
			createvolume.IOPS = IOPSV

		case "Encrypted":
			EncryptedV, _ := value.(bool)
			createvolume.Encrypted = EncryptedV

		case "SnapshotId":
			SnapshotIdV, _ := value.(string)
			createvolume.SnapshotId = SnapshotIdV
		}
	}

	fmt.Println(createvolume)

	volume1 := CreateVolume{
		AvailZone:  "us-east-1a",
		VolumeType: "gp2",
		VolumeSize: 100,
		IOPS:       3000,
		Encrypted:  true,
	}
	fmt.Println(volume1)

	params := makeParams("CreateVolume")
	prepareVolume(params, createvolume)
	resp = &CreateVolumeResp{}
	err = amazonstorage.PrepareSignatureV2query(params,Region,resp)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	return resp, nil
}

func (amazonstorage *Amazonstorage) Deletedisk(request interface{}) (resp interface{}, err error) {
	param, _ := request.(map[string]string)
	params := makeParams("DeleteVolume")
	params["VolumeId"] = param["VolumeId"]
	Region := param["Region"]
	resp = &SimpleResp{}
	err = amazonstorage.PrepareSignatureV2query(params,Region,resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//volumeId, description string
func (amazonstorage *Amazonstorage) Createsnapshot(request interface{}) (resp interface{}, err error) {

	param, _ := request.(map[string]string)
	params := makeParams("CreateSnapshot")

	params["VolumeId"] = param["VolumeId"]
	params["Description"] = param["Description"]
	Region := param["Region"]
	resp = &CreateSnapshotResp{}
	err = amazonstorage.PrepareSignatureV2query(params,Region, resp)
	if err != nil {
		return nil, err
	}
	return
}

func (amazonstorage *Amazonstorage) Deletesnapshot(request interface{}) (resp interface{}, err error) {
	ids := []string{}
	param, _ := request.(map[string]string)
	ids = append(ids, param["SnapshotId"])
	params := makeParams("DeleteSnapshot")
	Region := param["Region"]
	for i, id := range ids {
		params["SnapshotId."+strconv.Itoa(i+1)] = id
	}

	resp = &SimpleResp{}
	err = amazonstorage.PrepareSignatureV2query(params,Region, resp)
	if err != nil {
		return nil, err
	}
	return
}

func (amazonstorage *Amazonstorage) Attachdisk(request interface{}) (resp interface{}, err error) {
	param, _ := request.(map[string]string)
	params := makeParams("AttachVolume")
	params["VolumeId"] = param["VolumeId"]
	params["InstanceId"] = param["InstanceId"]
	params["Device"] = param["Device"]
	Region := param["Region"]
	resp = &VolumeAttachmentResp{}
	err = amazonstorage.PrepareSignatureV2query(params,Region, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (amazonstorage *Amazonstorage) Detachdisk(request interface{}) (resp interface{}, err error) {
	param, _ := request.(map[string]string)
	params := makeParams("DetachVolume")
	params["VolumeId"] = param["VolumeId"]
	params["InstanceId"] = param["InstanceId"]
	params["Device"] = param["Device"]
	if param["Force"] == "true" {
		params["Force"] = "true"
	}
	Region := param["Region"]
	resp = &VolumeAttachmentResp{}
	err = amazonstorage.PrepareSignatureV2query(params, Region, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
