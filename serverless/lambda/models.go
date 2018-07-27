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
	marker          string
	masterRegion    string
	maxItems        string
}

//CreateFunction struct represents aws serverless Create function.
type Createfunction struct {
	functionName string `json:"FunctionName"`
	handler      string `json:"Handler"`
	kMSKeyArn    string `json:"KMSKeyArn"`
	memorySize   int    `json:"MemorySize"`
	publish      bool   `json:"Publish"`
	role         string `json:"Role"`
	runtime      string `json:"Runtime"`
	tags         Tags   `json:"Tags"`
	description  string `json:"Description"`
	timeout      int `json:"Timeout"`
	deadLetterConfig DeadLetterConfig `json:"DeadLetterConfig"`
	environment      Environment      `json:"Environment"`
	tracingConfig    TracingConfig    `json:"TracingConfig"`
	vpcConfig        VpcConfig        `json:"VpcConfig"`
	code             Code             `json:"Code"`
}

//Tags struct represents CreateFunction parameters.
type Tags struct {
	String string `json:"string"`
}

//TracingConfig struct represents CreateFunction parameters.
type TracingConfig struct {
	mode string `json:"Mode"`
}

//VpcConfig struct represents CreateFunction parameters.
type VpcConfig struct {
	securityGroupIds []string `json:"SecurityGroupIds"`
	subnetIds        []string `json:"SubnetIds"`
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
	targetArn string `json:"TargetArn"`
}

//Code struct represents CreateFunction parameters.
type Code struct {
	s3Bucket        string `json:"S3Bucket"`
	s3Key           string `json:"S3Key"`
	s3ObjectVersion string `json:"S3ObjectVersion"`
	//zipFile         string `json:"ZipFile"`
	repositoryType  string `json:"RepositoryType"`
	location  string `json:"Location"`

	zipFile []byte `type:"blob"`

}
