package main

import(
  "github.com/scorelab/gocloud-v2/gocloud"
)


const(
  Amazonprovider = "aws"
  Googleprovider = "google"
  Secretkey = "1"
  Secretid = "2"
)


type RunInstances struct {
	ImageId               string
	MinCount              int
	MaxCount              int
	KeyName               string
	InstanceType          string
	KernelId              string
	RamdiskId             string
	UserData              []byte
	AvailZone             string
	PlacementGroupName    string
	Monitoring            bool
	SubnetId              string
	DisableAPITermination bool
	ShutdownBehavior      string
	PrivateIPAddress      string
}

func main(){

 amazoncloud, _ := gocloud.CloudProvider(Amazonprovider,Secretkey,Secretid)




  //var any interface{}

	r1:=RunInstances{
        ImageId:      "ami-ccf405a5", // Ubuntu Maverick, i386, EBS store
        InstanceType: "t1.micro",
    }
  //  any = r1
  amazoncloud.Createnode(&r1)

//amazoncloud.Createcontainer()
  googlecloud, _ := gocloud.CloudProvider(Googleprovider,Secretkey,Secretid)
  googlecloud.Createnode(&r1)


}
