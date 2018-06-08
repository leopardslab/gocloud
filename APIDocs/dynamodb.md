```
package ecs
    import "github.com/cloudlibz/gocloud/database/dynamodb"
```

TYPES

type AttributeDefinitions struct {
    AttributeName string `json:"AttributeName"`
    AttributeType string `json:"AttributeType"`
}
    AttributeDefinitions struct reperesnts Createtable AttributeDefinitions.

type Createtable struct {
    TableName string `json:"TableName"`
    // contains filtered or unexported fields
}
    Createtable struct reperesnts Createtable.

type Dynamodb struct {
}
    Dynamodb struct reperesnts aws Dynamodb service.

func (dynamodb *Dynamodb) Createtables(request interface{}) (resp interface{}, err error)
    Create tables.

func (dynamodb *Dynamodb) Deletetables(request interface{}) (resp interface{}, err error)
    Delete tables.

func (dynamodb *Dynamodb) Describetables(request interface{}) (resp interface{}, err error)
    Describe tables.

func (dynamodb *Dynamodb) Listtables(request interface{}) (resp interface{}, err error)
    List tables.

func (dynamodb *Dynamodb) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}, response map[string]interface{}) error
    PrepareSignatureV4query creates PrepareSignatureV4 for request.

type GlobalSecondaryIndexes struct {
    IndexName string `json:"IndexName"`
    // contains filtered or unexported fields
}
    GlobalSecondaryIndexes struct reperesnts Createtable
    GlobalSecondaryIndexes.

type KeySchema struct {
    AttributeName string `json:"AttributeName"`
    KeyType       string `json:"KeyType"`
}
    KeySchema struct reperesnts Createtable KeySchema.

type LocalSecondaryIndexes struct {
    IndexName string `json:"IndexName"`
    // contains filtered or unexported fields
}
    LocalSecondaryIndexes struct reperesnts Createtable
    LocalSecondaryIndexes.

type Projection struct {
    NonKeyAttributes []string `json:"NonKeyAttributes"`
    ProjectionType   string   `json:"ProjectionType"`
}
    Projection struct reperesnts Createtable Projection.

type ProvisionedThroughput struct {
    ReadCapacityUnits  int `json:"ReadCapacityUnits"`
    WriteCapacityUnits int `json:"WriteCapacityUnits"`
}
    ProvisionedThroughput struct reperesnts Createtable
    ProvisionedThroughput.

type SSESpecification struct {
    Enabled bool `json:"Enabled"`
}
    SSESpecification struct reperesnts Createtable SSESpecification.

type StreamSpecification struct {
    StreamViewType string `json:"StreamViewType"`
    StreamEnabled  bool   `json:"StreamEnabled"`
}
    StreamSpecification struct reperesnts Createtable StreamSpecification.


