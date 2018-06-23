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

// ListResourceDNSRecordSets to store all attribute to list resource Ali-cloud DNS record sets
type ListDNS struct {
	DomainName   string
	PageNumber   int
	PageSize     int
	RRKeyWord    string
	TypeKeyWord  string
	ValueKeyWord string
}
