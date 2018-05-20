package googleloadbalancer

//Createloadbalancerdictnoaryconvert creates a dictnoary for Createloadbalancer api.
func Createloadbalancerdictnoaryconvert(option TargetPools, Createloadbalancerjsonmap map[string]interface{}) {

	if option.ID != "" {
		Createloadbalancerjsonmap["id"] = option.ID
	}

	if len(option.BackupPool) != 0 {
		Createloadbalancerjsonmap["backupPool"] = option.BackupPool
	}

	if len(option.Instances) != 0 {
		Createloadbalancerjsonmap["instances"] = option.Instances
	}

	if option.FailoverRatio != 0 {
		Createloadbalancerjsonmap["failoverRatio"] = option.FailoverRatio
	}

	if (option.CreationTimestamp) != "" {
		Createloadbalancerjsonmap["creationTimestamp"] = option.CreationTimestamp
	}

	if (option.Description) != "" {
		Createloadbalancerjsonmap["description"] = option.Description
	}

	if len(option.HealthChecks) != 0 {
		Createloadbalancerjsonmap["healthChecks"] = option.HealthChecks
	}

	if (option.SessionAffinity) != "" {
		Createloadbalancerjsonmap["sessionAffinity"] = option.SessionAffinity
	}

	if option.SelfLink != "" {
		Createloadbalancerjsonmap["selfLink"] = option.SelfLink
	}
	if option.Region != "" {
		Createloadbalancerjsonmap["region"] = option.Region
	}

	if option.Name != "" {
		Createloadbalancerjsonmap["name"] = option.Name
	}

	if option.Kind != "" {
		Createloadbalancerjsonmap["kind"] = option.Kind
	}
}
