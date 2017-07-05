package awsloadbalancer



func (awsloadbalancer *Awsloadbalancer) Creatloadbalancer(request interface{}) (resp interface{}, err error) {

	var options CreateLoadBalancer

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {

		case "Name":
			NameV, _ := value.(string)
			options.Name = NameV

		case "IpAddressType":
			IpAddressTypeV, _ := value.(string)
			options.IpAddressType = IpAddressTypeV

		case "Scheme":
			SchemeV, _ := value.(string)
			options.Scheme = SchemeV

		case "Subnets":
			SubnetsV, _ := value.([]string)
			options.Subnets = SubnetsV

		case "SecurityGroups":
			SecurityGroupsV, _ := value.([]string)
			options.SecurityGroups = SecurityGroupsV

		}
	}

	params := makeParamsWithVersion("CreateLoadBalancer")

	prepareSubnets(params, options.Subnets)

	prepareSecurityGroups(params, options.SecurityGroups)

	if options.Name != "" {
		params["Name"] = options.Name
	}

	if options.IpAddressType != "" {
		params["IpAddressType"] = options.IpAddressType
	}

	if options.Scheme != "" {
		params["Scheme"] = options.Scheme
	}

	awsloadbalancer.PrepareSignatureV2query(params)

	return
}


func (awsloadbalancer *Awsloadbalancer) Deleteloadbalancer(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]string)

	params := makeParamsWithVersion("DeleteLoadBalancer")

	params["LoadBalancerArn"] = param["LoadBalancerArn"]

	awsloadbalancer.PrepareSignatureV2query(params)

	return
}

func (awsloadbalancer *Awsloadbalancer) Listloadbalancer(request interface{}) (resp interface{}, err error) {
  params := makeParamsWithVersion("DescribeLoadBalancers")
  if(request!=nil){
    param := request.(map[string]string)
    if(params["LoadBalancerArn"] !=""){
       params["LoadBalancerArn"] = param["LoadBalancerArn"]
    }
  }
  awsloadbalancer.PrepareSignatureV2query(params)
	return
}
