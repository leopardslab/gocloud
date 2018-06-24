package awsloadbalancer

//CreateLoadBalancer creates classic loadbalancer.
func (awsloadbalancer *Awsloadbalancer) CreateLoadBalancer(request interface{}) (resp interface{}, err error) {

	var options CreateLoadBalancer

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {

		case "LoadBalancerName":
			LoadBalancerNameV, _ := value.(string)
			options.LoadBalancerName = LoadBalancerNameV

		case "AvailabilityZones":
			AvailabilityZonesV, _ := value.([]string)
			options.AvailabilityZones = AvailabilityZonesV

		case "Scheme":
			SchemeV, _ := value.(string)
			options.Scheme = SchemeV

		case "Subnets":
			SubnetsV, _ := value.([]string)
			options.Subnets = SubnetsV

		case "SecurityGroups":
			SecurityGroupsV, _ := value.([]string)
			options.SecurityGroups = SecurityGroupsV

		case "Listeners":
			Listenersparam, _ := value.([]map[string]string)
			for i := 0; i < len(Listenersparam); i++ {
				var listener Listener
				listener.InstanceProtocol = Listenersparam[i]["InstanceProtocol"]
				listener.InstancePort = Listenersparam[i]["InstancePort"]
				listener.LoadBalancerPort = Listenersparam[i]["LoadBalancerPort"]
				listener.Protocol = Listenersparam[i]["Protocol"]
				listener.SSLCertificateId = Listenersparam[i]["SSLCertificateId"]
				options.Listeners = append(options.Listeners, listener)
			}
		}
	}

	params := makeParamsWithVersion("CreateLoadBalancer")

	prepareListeners(params, options.Listeners)

	prepareSubnets(params, options.Subnets)

	prepareSecurityGroups(params, options.SecurityGroups)

	prepareAvailabilityZones(params, options.AvailabilityZones)

	if options.LoadBalancerName != "" {
		params["LoadBalancerName"] = options.LoadBalancerName
	}
	if options.Scheme != "" {
		params["Scheme"] = options.Scheme
	}

	response := make(map[string]interface{})
	err = awsloadbalancer.PrepareSignatureV2query(params, response)
	resp = response
	return resp, err
}

//DeleteLoadBalancer Delete loadbalancer accept LoadBalancerName.
func (awsloadbalancer *Awsloadbalancer) DeleteLoadBalancer(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]string)

	params := makeParamsWithVersion("DeleteLoadBalancer")

	params["LoadBalancerName"] = param["LoadBalancerName"]

	response := make(map[string]interface{})
	err = awsloadbalancer.PrepareSignatureV2query(params, response)
	resp = response
	return resp, err
}

//ListLoadBalancer List running loadbalancer.
func (awsloadbalancer *Awsloadbalancer) ListLoadBalancer(request interface{}) (resp interface{}, err error) {
	params := makeParamsWithVersion("DescribeLoadBalancers")
	if request != nil {
		param := request.(map[string]string)
		if params["LoadBalancerArn"] != "" {
			params["LoadBalancerArn"] = param["LoadBalancerArn"]
		}
	}
	response := make(map[string]interface{})
	err = awsloadbalancer.PrepareSignatureV2query(params, response)
	resp = response
	return resp, err
}

//DetachNodeWithLoadBalancer Detach node with loadbalancer.
func (awsloadbalancer *Awsloadbalancer) DetachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error) {
	var options AttachNodeWithLoadBalancer

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {

		case "LoadBalancerName":
			LoadBalancerNameV, _ := value.(string)
			options.LoadBalancerName = LoadBalancerNameV

		case "Instances":
			InstancesV, _ := value.([]string)
			options.Instances = InstancesV

		}
	}
	params := makeParamsWithVersion("DeregisterInstancesFromLoadBalancer")

	prepareInstances(params, options.Instances)

	if options.LoadBalancerName != "" {
		params["LoadBalancerName"] = options.LoadBalancerName
	}

	response := make(map[string]interface{})
	err = awsloadbalancer.PrepareSignatureV2query(params, response)
	resp = response
	return resp, err
}

// AttachNodeWithLoadBalancer method Attach node with loadbalancer.
func (awsloadbalancer *Awsloadbalancer) AttachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error) {

	var options AttachNodeWithLoadBalancer

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {

		case "LoadBalancerName":
			LoadBalancerNameV, _ := value.(string)
			options.LoadBalancerName = LoadBalancerNameV

		case "Instances":
			InstancesV, _ := value.([]string)
			options.Instances = InstancesV

		}
	}
	params := makeParamsWithVersion("RegisterInstancesWithLoadBalancer")

	prepareInstances(params, options.Instances)

	if options.LoadBalancerName != "" {
		params["LoadBalancerName"] = options.LoadBalancerName
	}
	response := make(map[string]interface{})
	err = awsloadbalancer.PrepareSignatureV2query(params, response)
	resp = response
	return resp, err
}
