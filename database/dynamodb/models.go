package dynamodb


type Dynamodb struct{

}
/*

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

type ProvisionedThroughput struct {
}

type GlobalSecondaryIndexes struct {
  IndexName string `json:"IndexName"`
  keySchema []KeySchema `json:"KeySchema"`
  projection Projection`json:"Projection"`
  provisionedThroughput  ProvisionedThroughput `json:"ProvisionedThroughput"`
}



type Createtable struct {
	attributeDefinitions []AttributeDefinitions `json:"AttributeDefinitions"`
  globalSecondaryIndexes []GlobalSecondaryIndexes `json:"GlobalSecondaryIndexes"`

  type keySchema []struct {
		AttributeName string `json:"AttributeName"`
		KeyType       string `json:"KeyType"`
	} `json:"KeySchema"`


  LocalSecondaryIndexes []struct {
		IndexName string `json:"IndexName"`
		KeySchema []struct {
			AttributeName string `json:"AttributeName"`
			KeyType       string `json:"KeyType"`
		} `json:"KeySchema"`
		Projection struct {
			NonKeyAttributes []string `json:"NonKeyAttributes"`
			ProjectionType   string   `json:"ProjectionType"`
		} `json:"Projection"`
	} `json:"LocalSecondaryIndexes"`

  provisionedThroughput ProvisionedThroughput `json:"ProvisionedThroughput"`
  sSESpecification SSESpecification `json:"SSESpecification"`
  streamSpecification StreamSpecification `json:"StreamSpecification"`
	TableName string `json:"TableName"`
}

type ProvisionedThroughput struct {

}


type SSESpecification struct {
}

type StreamSpecification struct {
  StreamViewType string `json:"StreamViewType"`
}

*/
