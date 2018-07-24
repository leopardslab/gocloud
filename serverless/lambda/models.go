package lambda

//Lambda struct reperesnts aws serverless service.
type Lambda struct {
}


//Deletefunction struct represents aws serverless Delete function.
type Deletefunction struct {
	FunctionName string
	Qualifier    string
}

//Getfunction struct represents aws serverless Get function.
type Getfunction struct {
	FunctionName string
	Qualifier    string
}


//Listfunction struct represents aws serverless List function.
 type Listfunction struct {
	functionVersion string
	marker    string
	masterRegion string
	maxItems string
}



//CreateFunction struct represents aws serverless Create function.
type CreateFunction struct {
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

//Tags struct represents CreateFunction parameters.
type Tags struct {
	String string `json:"string"`
}

//TracingConfig struct represents CreateFunction parameters.
type TracingConfig struct {
	Mode string `json:"Mode"`
}

//VpcConfig struct represents CreateFunction parameters.
type VpcConfig struct {
	SecurityGroupIds []string `json:"SecurityGroupIds"`
	SubnetIds        []string `json:"SubnetIds"`
}

//Variables struct represents CreateFunction parameters.
type Variables struct {
	String string `json:"string"`
}

//Environment struct represents CreateFunction parameters.
type Environment struct {
	variables Variables `json:"Variables"`
}

//DeadLetterConfig struct represents CreateFunction parameters.
type DeadLetterConfig struct {
	TargetArn string `json:"TargetArn"`
}

//Code struct represents CreateFunction parameters.
type Code struct {
	S3Bucket        string `json:"S3Bucket"`
	S3Key           string `json:"S3Key"`
	S3ObjectVersion string `json:"S3ObjectVersion"`
}
