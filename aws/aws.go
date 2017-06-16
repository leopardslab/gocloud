package aws

import (
	ec2 "github.com/scorelab/gocloud-v2/compute/ec2"
)

// repersents AWS struct

type AWS struct {
	ec2.EC2
}
