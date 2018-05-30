package lambda

//Lambda struct reperesnts aws serverless service.
type Lambda struct {
}

type Deletefunction struct {
	FunctionName string
	Qualifier    string
}

//Createfunction struct represents aws serverless Create function.
type Createfunction struct {
	code Code `json:"Code"`

	Description string `json:"Description"`

	deadLetterConfig DeadLetterConfig `json:"DeadLetterConfig"`

	environment Environment `json:"Environment"`

	FunctionName  string        `json:"FunctionName"`
	Handler       string        `json:"Handler"`
	KMSKeyArn     string        `json:"KMSKeyArn"`
	Role          string        `json:"Role"`
	Runtime       string        `json:"Runtime"`
	tags          Tags          `json:"Tags"`
	tracingConfig TracingConfig `json:"TracingConfig"`
	vpcConfig     VpcConfig     `json:"VpcConfig"`
}

//Tags struct represents Createfunction parameters.
type Tags struct {
	String string `json:"string"`
}

//TracingConfig struct represents Createfunction parameters.
type TracingConfig struct {
	Mode string `json:"Mode"`
}

//VpcConfig struct represents Createfunction parameters.
type VpcConfig struct {
	SecurityGroupIds []string `json:"SecurityGroupIds"`
	SubnetIds        []string `json:"SubnetIds"`
}

//Variables struct represents Createfunction parameters.
type Variables struct {
	String string `json:"string"`
}

//Environment struct represents Createfunction parameters.
type Environment struct {
	variables Variables `json:"Variables"`
}

//DeadLetterConfig struct represents Createfunction parameters.
type DeadLetterConfig struct {
	TargetArn string `json:"TargetArn"`
}

//Code struct represents Createfunction parameters.
type Code struct {
	S3Bucket        string `json:"S3Bucket"`
	S3Key           string `json:"S3Key"`
	S3ObjectVersion string `json:"S3ObjectVersion"`
}
