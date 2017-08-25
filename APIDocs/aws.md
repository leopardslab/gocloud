package aws
    import "github.com/scorelab/gocloud-v2/aws"


TYPES

type AWS struct {
    ec2.EC2
    amazonstorage.Amazonstorage
    awsloadbalancer.Awsloadbalancer
    awscontainer.Ecscontainer
    awsdns.Awsdns
}
    AWS struct reperents amazon cloud provider.


