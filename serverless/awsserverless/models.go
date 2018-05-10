package lambda

type Createfunction struct {

  code Code `json:"Code"`

  Description string `json:"Description"`

  deadLetterConfig DeadLetterConfig `json:"DeadLetterConfig"`

  environment Environment `json:"Environment"`

	FunctionName string `json:"FunctionName"`
	Handler      string `json:"Handler"`
	KMSKeyArn    string `json:"KMSKeyArn"`
	Role         string `json:"Role"`
	Runtime      string `json:"Runtime"`
  tags Tags`json:"Tags"`
  tracingConfig TracingConfig`json:"TracingConfig"`
  vpcConfig VpcConfig`json:"VpcConfig"`
}

type Tags struct {
  String string `json:"string"`
}

type TracingConfig struct {
  Mode string `json:"Mode"`
}

type VpcConfig struct {
  SecurityGroupIds []string `json:"SecurityGroupIds"`
  SubnetIds        []string `json:"SubnetIds"`
}

type Variables struct {
  String string `json:"string"`
}

type Environment struct {
  variables Variables `json:"Variables"`
}

type DeadLetterConfig struct {
  TargetArn string `json:"TargetArn"`
}

type Code struct {
  S3Bucket        string `json:"S3Bucket"`
  S3Key           string `json:"S3Key"`
  S3ObjectVersion string `json:"S3ObjectVersion"`
}
