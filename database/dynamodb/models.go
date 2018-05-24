package dynamodb

type Dynamodb struct {
}

type AttributeDefinitions struct {
	AttributeName string `json:"AttributeName"`
	AttributeType string `json:"AttributeType"`
}

type KeySchema struct {
	AttributeName string `json:"AttributeName"`
	KeyType       string `json:"KeyType"`
}

type Projection struct {
	NonKeyAttributes []string `json:"NonKeyAttributes"`
	ProjectionType   string   `json:"ProjectionType"`
}

type GlobalSecondaryIndexes struct {
	IndexName             string                `json:"IndexName"`
	keySchema             []KeySchema           `json:"KeySchema"`
	projection            Projection            `json:"Projection"`
	provisionedThroughput ProvisionedThroughput `json:"ProvisionedThroughput"`
}

type Createtable struct {
	attributeDefinitions   []AttributeDefinitions   `json:"AttributeDefinitions"`
	globalSecondaryIndexes []GlobalSecondaryIndexes `json:"GlobalSecondaryIndexes"`
	localSecondaryIndexes  []LocalSecondaryIndexes  `json:"LocalSecondaryIndexes"`
	keySchema              []KeySchema              `json:"KeySchema"`
	provisionedThroughput  ProvisionedThroughput    `json:"ProvisionedThroughput"`
	sSESpecification       SSESpecification         `json:"SSESpecification"`
	streamSpecification    StreamSpecification      `json:"StreamSpecification"`
	TableName              string                   `json:"TableName"`
}

type LocalSecondaryIndexes struct {
	IndexName  string      `json:"IndexName"`
	keySchema  []KeySchema `json:"KeySchema"`
	projection Projection  `json:"Projection"`
}

type ProvisionedThroughput struct {
	ReadCapacityUnits  int `json:"ReadCapacityUnits"`
	WriteCapacityUnits int `json:"WriteCapacityUnits"`
}

type SSESpecification struct {
	Enabled bool `json:"Enabled"`
}

type StreamSpecification struct {
	StreamViewType string `json:"StreamViewType"`
	StreamEnabled  bool   `json:"StreamEnabled"`
}
