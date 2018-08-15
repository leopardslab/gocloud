PACKAGE DOCUMENTATION

package lambda
    import "github.com/cloudlibz/gocloud/serverless/lambda"


FUNCTIONS

func PreparePostrequest(params map[string]interface{}, region string, response map[string]interface{}) (err error)

func Preparedeletefnrequest(params map[string]string, region string, response map[string]interface{}) (err error)

func Preparegetfnrequest(params map[string]string, region string, response map[string]interface{}) (err error)

func Preparegetrequest(params map[string]string, region string, response map[string]interface{}) (err error)

TYPES

type Code struct {
    // contains filtered or unexported fields
}
    Code struct represents CreateFunction parameters.

type Createfunction struct {
    // contains filtered or unexported fields
}
    CreateFunction struct represents aws serverless Create function.

type DeadLetterConfig struct {
    // contains filtered or unexported fields
}
    DeadLetterConfig struct represents CreateFunction parameters.

type Deletefunction struct {
    FunctionName string
    Qualifier    string
}
    Deletefunction struct represents aws serverless Delete function.

type Environment struct {
    // contains filtered or unexported fields
}
    Environment struct represents CreateFunction parameters.

type Getfunction struct {
    FunctionName string
    Qualifier    string
}
    Getfunction struct represents aws serverless Get function.

type Lambda struct {
}
    Lambda struct reperesnts aws serverless service.

func (lambda *Lambda) CallFunction(request interface{}) (resp interface{}, err error)
    CallFunction call lambda function.

func (lambda *Lambda) CreateFunction(request interface{}) (resp interface{}, err error)
    CreateFunction Create lambda function.

func (lambda *Lambda) DeleteFunction(request interface{}) (resp interface{}, err error)
    DeleteFunction delete lambda function.

func (lambda *Lambda) GetFunction(request interface{}) (resp interface{}, err error)
    GetFunction describe lambda function.

func (lambda *Lambda) ListFunction(request interface{}) (resp interface{}, err error)
    ListFunction list lambda function.

type Listfunction struct {
    // contains filtered or unexported fields
}
    Listfunction struct represents aws serverless List function.

type Tags struct {
    String string `json:"string"`
}
    Tags struct represents CreateFunction parameters.

type TracingConfig struct {
    // contains filtered or unexported fields
}
    TracingConfig struct represents CreateFunction parameters.

type Variables struct {
    String string `json:"string"`
}
    Variables struct represents CreateFunction parameters.

type VpcConfig struct {
    // contains filtered or unexported fields
}
    VpcConfig struct represents CreateFunction parameters.
