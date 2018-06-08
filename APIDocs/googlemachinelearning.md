package googlemachinelearning
   import "github.com/cloudlibz/gocloud/machinelearning/googlemachinelearning"


CONSTANTS

const (
    UnixDate = "Mon Jan _2 15:04:05 MST 2006"
    RFC3339  = "2006-01-02T15:04:05Z07:00"
)

TYPES

type AutoScaling struct {
    // contains filtered or unexported fields
}
    AutoScaling struct reperesnts DefaultVersion sub-struct.

type CreateMLModel struct {
    // contains filtered or unexported fields
}
    CreateMLModel struct reperesnts CreateMLModel service api.

type DefaultVersion struct {
    // contains filtered or unexported fields
}
    DefaultVersion struct reperesnts CreateMLModel sub-struct.

type Googlemachinelearning struct {
}
    Googlemachinelearning struct reperesnts google machine learning service.

func (googlemachinelearning *Googlemachinelearning) CreateMLModel(request interface{}) (resp interface{}, err error)
    CreateMLModel creates model.

func (googlemachinelearning *Googlemachinelearning) DeleteMLModel(request interface{}) (resp interface{}, err error)
    DeleteMLModel delete model.

func (googlemachinelearning *Googlemachinelearning) GetMLModel(request interface{}) (resp interface{}, err error)
    GetMLModel get model.

func (googlemachinelearning *Googlemachinelearning) UpdateMLModel(request interface{}) (resp interface{}, err error)
    UpdateMLModel Update model.

type ManualScaling struct {
    // contains filtered or unexported fields
}
    ManualScaling struct reperesnts DefaultVersion sub-struct.


