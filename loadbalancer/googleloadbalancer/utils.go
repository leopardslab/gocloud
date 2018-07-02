package googleloadbalancer

//CreateLoadBalancerdictnoaryconvert creates a dictnoary for CreateLoadBalancer api.
func CreateLoadBalancerdictnoaryconvert(option TargetPools, CreateLoadBalancerjsonmap map[string]interface{}) {

	if option.ID != "" {
		CreateLoadBalancerjsonmap["id"] = option.ID
	}

	if len(option.BackupPool) != 0 {
		CreateLoadBalancerjsonmap["backupPool"] = option.BackupPool
	}

	if len(option.Instances) != 0 {
		CreateLoadBalancerjsonmap["instances"] = option.Instances
	}

	if option.FailoverRatio != 0 {
		CreateLoadBalancerjsonmap["failoverRatio"] = option.FailoverRatio
	}

	if (option.CreationTimestamp) != "" {
		CreateLoadBalancerjsonmap["creationTimestamp"] = option.CreationTimestamp
	}

	if (option.Description) != "" {
		CreateLoadBalancerjsonmap["description"] = option.Description
	}

	if len(option.HealthChecks) != 0 {
		CreateLoadBalancerjsonmap["healthChecks"] = option.HealthChecks
	}

	if (option.SessionAffinity) != "" {
		CreateLoadBalancerjsonmap["sessionAffinity"] = option.SessionAffinity
	}

	if option.SelfLink != "" {
		CreateLoadBalancerjsonmap["selfLink"] = option.SelfLink
	}
	if option.Region != "" {
		CreateLoadBalancerjsonmap["region"] = option.Region
	}

	if option.Name != "" {
		CreateLoadBalancerjsonmap["name"] = option.Name
	}

	if option.Kind != "" {
		CreateLoadBalancerjsonmap["kind"] = option.Kind
	}
}
