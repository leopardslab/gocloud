PACKAGE DOCUMENTATION

package kinesis
    import "github.com/cloudlibz/gocloud/streamdataprocessing/kinesis"


TYPES

type Kinesis struct {
}
    Kinesis struct reperesnts aws streaming service.

func (kinesis *Kinesis) CreateStream(request interface{}) (resp interface{}, err error)
    CreateStream Create Stream

func (kinesis *Kinesis) DeleteStream(request interface{}) (resp interface{}, err error)
    DeleteStream Delete Stream

func (kinesis *Kinesis) DescribeStream(request interface{}) (resp interface{}, err error)
    DescribeStream describe stream

func (kinesis *Kinesis) ListStream(request interface{}) (resp interface{}, err error)
    ListStream List Stream

func (kinesis *Kinesis) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}, response map[string]interface{}) error
    PrepareSignatureV4query creates PrepareSignatureV4 for request.
