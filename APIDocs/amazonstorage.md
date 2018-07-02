package amazonstorage
    import "github.com/cloudlibz/gocloud/storage/amazonstorage"


TYPES

type Amazonstorage struct {
}

func (amazonstorage *Amazonstorage) AttachDisk(request interface{}) (resp interface{}, err error)

func (amazonstorage *Amazonstorage) CreateDisk(request interface{}) (resp interface{}, err error)

func (amazonstorage *Amazonstorage) CreateSnapshot(request interface{}) (resp interface{}, err error)
    volumeId, description string

func (amazonstorage *Amazonstorage) DeleteDisk(request interface{}) (resp interface{}, err error)

func (amazonstorage *Amazonstorage) DeleteSnapshot(request interface{}) (resp interface{}, err error)

func (amazonstorage *Amazonstorage) DetachDisk(request interface{}) (resp interface{}, err error)

func (amazonstorage *Amazonstorage) PrepareSignatureV2query(params map[string]string, Region string, response map[string]interface{}) error

type CreateSnapshotResp struct {
    RequestId string `xml:"requestId"`
    Snapshot
}

type CreateVolume struct {
    AvailZone  string
    SnapshotId string
    VolumeType string
    VolumeSize int
    Encrypted  bool
    IOPS       int64
}

type CreateVolumeResp struct {
    RequestId string `xml:"requestId"`
    Volume
}

type SimpleResp struct {
    XMLName   xml.Name
    RequestId string `xml:"requestId"`
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

type SnapshotsResp struct {
    RequestId string     `xml:"requestId"`
    Snapshots []Snapshot `xml:"snapshotSet>item"`
}

type Tag struct {
    Key   string `xml:"key"`
    Value string `xml:"value"`
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

type VolumeAttachment struct {
    VolumeId            string `xml:"volumeId"`
    Device              string `xml:"device"`
    InstanceId          string `xml:"instanceId"`
    Status              string `xml:"status"`
    DeleteOnTermination bool   `xml:"deleteOnTermination"`
}

type VolumeAttachmentResp struct {
    RequestId  string `xml:"requestId"`
    VolumeId   string `xml:"volumeId"`
    Device     string `xml:"device"`
    InstanceId string `xml:"instanceId"`
    Status     string `xml:"status"`
    AttachTime string `xml:"attachTime"`
}


