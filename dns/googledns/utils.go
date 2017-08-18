package googledns

//Creatednsedictnoaryconvert convert Createdns parameters to Createdns dictnoary.
func Creatednsedictnoaryconvert(option Createdns, Creatednsjsonmap map[string]interface{}) {

	if len(option.NameServers) != 0 {
		Creatednsjsonmap["nameServers"] = option.NameServers
	}

	if option.Name != "" {
		Creatednsjsonmap["name"] = option.Name
	}

	if option.NameServerSet != "" {
		Creatednsjsonmap["NameServerSet"] = option.NameServerSet
	}

	if option.DNSName != "" {
		Creatednsjsonmap["dnsName"] = option.DNSName
	}

	if option.Kind != "" {
		Creatednsjsonmap["kind"] = option.Kind
	}

	if option.ID != "" {
		Creatednsjsonmap["id"] = option.ID
	}
	if option.Description != "" {
		Creatednsjsonmap["description"] = option.Description
	}
	if option.CreationTime != "" {
		Creatednsjsonmap["creationTime"] = option.CreationTime
	}
}
