PACKAGE DOCUMENTATION

package bigquery
    import "github.com/cloudlibz/gocloud/database/bigtable"


CONSTANTS

const (
    UnixDate = "Mon Jan _2 15:04:05 MST 2006"
    RFC3339  = "2006-01-02T15:04:05Z07:00"
)

TYPES

type Access struct {
    // contains filtered or unexported fields
}
    Access struct reperesnts CreateDatasets internal struct.

type Bigquery struct {
}
    Bigquery struct reperesnts google analytics service.

func (bigquery *Bigquery) CreateDatasets(request interface{}) (resp interface{}, err error)
    CreateDatasets Create Datasets.

func (bigquery *Bigquery) DeleteDatasets(request interface{}) (resp interface{}, err error)
    DeleteDatasets delete Datasets.

func (bigquery *Bigquery) GetDatasets(request interface{}) (resp interface{}, err error)
    GetDatasets get Datasets.

func (bigquery *Bigquery) ListDatasets(request interface{}) (resp interface{}, err error)
    ListDatasets list Datasets.

func (bigquery *Bigquery) UpdateDatasets(request interface{}) (resp interface{}, err error)
    UpdateDatasets Update Datasets.

type CreateDatasets struct {
    // contains filtered or unexported fields
}
    CreateDatasets struct reperesnts google bigquery CreateDatasets service
    struct.

type DatasetReference struct {
    // contains filtered or unexported fields
}
    DatasetReference struct reperesnts CreateDatasets internal struct.

type View struct {
    // contains filtered or unexported fields
}
    View struct reperesnts CreateDatasets internal struct.
