PACKAGE DOCUMENTATION

package awsloadbalancer
    import "github.com/cloudlibz/gocloud/machinelearning/awsmachinelearning"


TYPES

type Awsmachinelearning struct {
}
    Awsmachinelearning struct reperesnts aws machine learning service.

func (awsmachinelearning *Awsmachinelearning) CreateMLModel(request interface{}) (resp interface{}, err error)
    CreateMLModel creates model.

func (awsmachinelearning *Awsmachinelearning) DeleteMLModel(request interface{}) (resp interface{}, err error)
    DeleteMLModel delete model.

func (awsmachinelearning *Awsmachinelearning) GetMLModel(request interface{}) (resp interface{}, err error)
    GetMLModel describe model.

func (awsmachinelearning *Awsmachinelearning) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}, response map[string]interface{}) error
    PrepareSignatureV4query creates PrepareSignatureV4 for request.

func (awsmachinelearning *Awsmachinelearning) UpdateMLModel(request interface{}) (resp interface{}, err error)
    UpdateMLModel update model.

type CreateMLModel struct {
    MLModelID   string `json:"MLModelId"`
    MLModelName string `json:"MLModelName"`
    MLModelType string `json:"MLModelType"`

    Recipe               string `json:"Recipe"`
    RecipeURI            string `json:"RecipeUri"`
    TrainingDataSourceID string `json:"TrainingDataSourceId"`
    // contains filtered or unexported fields
}
    CreateMLModel struct reperesnts aws machine learning service
    CreateMLModel.

type Parameters struct {
    String string `json:"string"`
}
    Parameters struct reperesnts CreateMLModel parameters.


