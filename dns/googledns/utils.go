package googledns

//CreateDnsedictnoaryconvert convert CreateDns parameters to CreateDns dictnoary.
func CreateDnsedictnoaryconvert(option CreateDns, CreateDnsjsonmap map[string]interface{}) {

	if len(option.NameServers) != 0 {
		CreateDnsjsonmap["nameServers"] = option.NameServers
	}

	if option.Name != "" {
		CreateDnsjsonmap["name"] = option.Name
	}

	if option.NameServerSet != "" {
		CreateDnsjsonmap["NameServerSet"] = option.NameServerSet
	}

	if option.DNSName != "" {
		CreateDnsjsonmap["dnsName"] = option.DNSName
	}

	if option.Kind != "" {
		CreateDnsjsonmap["kind"] = option.Kind
	}

	if option.ID != "" {
		CreateDnsjsonmap["id"] = option.ID
	}
	if option.Description != "" {
		CreateDnsjsonmap["description"] = option.Description
	}
	if option.CreationTime != "" {
		CreateDnsjsonmap["creationTime"] = option.CreationTime
	}
}
