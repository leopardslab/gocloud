package ec2


const (
	debug = false

	legacyAPIVersion = "2011-12-15"

	vpcAPIVersion = "2013-10-15"
)

type Auth struct {
	AccessKey string
	SecretKey string
}


type EC2 struct {
	Auth
	Region
	Private byte
}


func New2(auth Auth, region Region) *EC2 {
	return &EC2{auth,region, 0}
}
