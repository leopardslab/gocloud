package alidns

//Alidns represents Alidns attribute and method associates with it.
type Alidns struct {
}

// CreateDNS to store all attribute to create Ali-cloud DNS
type CreateDNS struct {
	DomainName string
	RR         string
	Type       string
	Value      string
	TTL        int
	Priority   int
	Line       string
}

// DeleteDNS to store all attribute to delete Ali-cloud DNS
type DeleteDNS struct {
	RecordId string
}

// ListDNS to store all attribute to list Ali-cloud DNS
type ListDNS struct {
	PageNumber int
	PageSize   int
	KeyWord    string
	GroupId    string
}

// ListResourceDNSRecordSets to store all attribute to list resource Ali-cloud DNS record sets
type ListResourceDNSRecordSets struct {
	DomainName   string
	PageNumber   int
	PageSize     int
	RRKeyWord    string
	TypeKeyWord  string
	ValueKeyWord string
}
