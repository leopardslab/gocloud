package ec2

//repersents Region
type Region struct {
	Name                 string
	EC2Endpoint          string
}


var USEast = Region{
	"us-east-1",
	"https://ec2.us-east-1.amazonaws.com",
}
