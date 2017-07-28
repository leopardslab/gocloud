package googleloadbalancer

func Creatloadbalancerdictnoaryconvert(option TargetPools, Creatloadbalancerjsonmap map[string]interface{}) {

	if option.ID == "" {
		Creatloadbalancerjsonmap["id"] = option.ID
	}

	if len(option.BackupPool) != 0 {
		Creatloadbalancerjsonmap["backupPool"] = option.BackupPool
	}

	if len(option.Instances) != 0 {
		Creatloadbalancerjsonmap["instances"] = option.Instances
	}

	if option.FailoverRatio != 0 {
		Creatloadbalancerjsonmap["failoverRatio"] = option.FailoverRatio
	}

	if option.CreationTimestamp == "" {
		Creatloadbalancerjsonmap["creationTimestamp"] = option.CreationTimestamp
	}

	if option.Description == "" {
		Creatloadbalancerjsonmap["description"] = option.Description
	}

	if len(option.HealthChecks) != 0 {
		Creatloadbalancerjsonmap["healthChecks"] = option.HealthChecks
	}

	if option.SessionAffinity == "" {
		Creatloadbalancerjsonmap["sessionAffinity"] = option.SessionAffinity
	}

	if option.SelfLink == "" {
		Creatloadbalancerjsonmap["selfLink"] = option.SelfLink
	}
	if option.Region == "" {
		Creatloadbalancerjsonmap["region"] = option.Region
	}

	if option.Name == "" {
		Creatloadbalancerjsonmap["name"] = option.Name
	}

	if option.Kind == "" {
		Creatloadbalancerjsonmap["kind"] = option.Kind
	}

}
