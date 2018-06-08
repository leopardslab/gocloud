```
package googlecloudfunctions
    import "github.com/cloudlibz/gocloud/serverless/googlecloudfunctions"
```


CONSTANTS

const (
    UnixDate = "Mon Jan _2 15:04:05 MST 2006"
    RFC3339  = "2006-01-02T15:04:05Z07:00"
)

FUNCTIONS

func CreateGooglecloudfunctionedictnoaryconvert(option CreateGooglecloudfunction, CreateGooglecloudfunctionjsonmap map[string]interface{})
    CreateGooglecloudfunctionedictnoaryconvert convert dict to form valid
    json.

TYPES

type CreateGooglecloudfunction struct {
    Name                string       `json:"name"`
    Status              string       `json:"status"`
    EntryPoint          string       `json:"entryPoint"`
    Timeout             string       `json:"timeout"`
    AvailableMemoryMb   int          `json:"availableMemoryMb"`
    ServiceAccountEmail string       `json:"serviceAccountEmail"`
    UpdateTime          string       `json:"updateTime"`
    VersionID           string       `json:"versionId"`
    SourceUploadURL     string       `json:"sourceUploadUrl"`
    Labels              labels       `json:"labels"`
    HTTPSTrigger        httpstrigger `json:"httpsTrigger"`
}
    CreateGooglecloudfunction struct reperesnts Create Googlecloud function.

type Googlecloudfunctions struct {
}
    Googlecloudfunctions struct reperesnts Google cloud functions service.

func (googlecloudfunctions *Googlecloudfunctions) Callfunction(request interface{}) (resp interface{}, err error)
    Callfunction call function.

func (googlecloudfunctions *Googlecloudfunctions) Createfunction(request interface{}) (resp interface{}, err error)
    Create Create Google cloud function.

func (googlecloudfunctions *Googlecloudfunctions) Deletefunction(request interface{}) (resp interface{}, err error)
    Delete delete function.

func (googlecloudfunctions *Googlecloudfunctions) Getfunction(request interface{}) (resp interface{}, err error)
    Getfunction get function.

func (googlecloudfunctions *Googlecloudfunctions) Listfunction(request interface{}) (resp interface{}, err error)
    Listfunction list function.


