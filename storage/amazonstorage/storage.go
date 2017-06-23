package amazonstorage

import (
	"encoding/xml"
	awsauth "github.com/scorelab/gocloud-v2/awsauth"
	auth "github.com/scorelab/gocloud-v2/auth"
	//"log"
	"net/http"
	//"net/http/httputil"
	"strconv"
	"time"
)

type Region struct {
	Name        string
	EC2Endpoint string
}

var USEast = Region{
	"us-east-1",
	"https://ec2.us-east-1.amazonaws.com",
}


type Amazonstorage struct {
}

/*
func (amazonstorage *Amazonstorage) Createdisk(request interface{}) (resp interface{}, err error) {

	return
}

func (amazonstorage *Amazonstorage) Deletedisk(request interface{}) (resp interface{}, err error) {
	return
}
*/

func (amazonstorage *Amazonstorage) Createsnapshot(request interface{}) (resp interface{}, err error) {
	return
}

func (amazonstorage *Amazonstorage) Deletesnapshot(request interface{}) (resp interface{}, err error) {

	return
}

/*
func (amazonstorage *Amazonstorage)Attachdisk (request interface{}) (resp interface{}, err error){

	return
}



*/

const (
	debug = false

	legacyAPIVersion = "2011-12-15"

	vpcAPIVersion = "2013-10-15"
)






type CreateVolume struct {
	AvailZone  string
	SnapshotId string
	VolumeType string
	VolumeSize int // Size is given in GB
	Encrypted  bool
	IOPS int64
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

//CreateVolume

func (amazonstorage *Amazonstorage) Createdisk(request interface{}) (resp interface{}, err error) {
	params := makeParams("CreateVolume")
	prepareVolume(params, volume)
	resp = &CreateVolumeResp{}
	err = amazonstorage.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
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

//id string
func (amazonstorage *Amazonstorage) Deletedisk(request interface{}) (resp interface{}, err error) {
	params := makeParams("DeleteVolume")
	params["VolumeId"] = id
	resp = &SimpleResp{}
	err = amazonstorage.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
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

//volumeId, instanceId, device string
func (amazonstorage *Amazonstorage)Attachdisk (request interface{}) (resp interface{}, err error){
	params := makeParams("AttachVolume")
	params["VolumeId"] = volumeId
	params["InstanceId"] = instanceId
	params["Device"] = device
	resp = &VolumeAttachmentResp{}
	err = amazonstorage.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// volumeId, instanceId, device string, force bool
func (googlestorage *GoogleStorage) Detachdisk(request interface{}) (resp interface{}, err error) {
	params := makeParams("DetachVolume")
	params["VolumeId"] = volumeId
	params["InstanceId"] = instanceId
	params["Device"] = device
	if force {
		params["Force"] = "true"
	}
	resp = &VolumeAttachmentResp{}
	err = amazonstorage.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
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



func (amazonstorage *Amazonstorage) query(params map[string]string, resp interface{}) error {

	req, err := http.NewRequest("GET", USEast.EC2Endpoint, nil)
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

	//fmt.Println(req)

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return xml.NewDecoder(r.Body).Decode(resp)
}


//volumeId, description string
func (amazonstorage *Amazonstorage) Createsnapshot(volumeId, description string) (resp *CreateSnapshotResp, err error) {
	params := makeParams("CreateSnapshot")
	params["VolumeId"] = volumeId
	params["Description"] = description

	resp = &CreateSnapshotResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return
}


//ids []string

func (amazonstorage *Amazonstorage) Deletesnapshot(ids []string) (resp *SimpleResp, err error) {
	params := makeParams("DeleteSnapshot")
	for i, id := range ids {
		params["SnapshotId."+strconv.Itoa(i+1)] = id
	}

	resp = &SimpleResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return
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
