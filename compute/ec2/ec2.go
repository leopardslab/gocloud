package ec2

const (
	debug = false
	legacyAPIVersion = "2011-12-15"
	vpcAPIVersion = "2013-10-15"
)

//Authentication struct to store AccessKey and SecretKey
type Auth struct {
	AccessKey string
	SecretKey string
}

//Ec2 struct

type EC2 struct {
	Auth
	Region
	Private byte
}

// Function return EC2 instance
func New2(auth Auth, region Region) *EC2 {
	return &EC2{auth, region, 0}
}
