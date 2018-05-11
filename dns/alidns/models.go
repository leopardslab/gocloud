package alidns

//Alidns represents Alidns attribute and method associates with it.
type Alidns struct {
}

type CreateDNS struct {
	DomainName string
	RR         string
	Type       string
	Value      string
	TTL        int
	Priority   int
	Line       string
}

type DeleteDNS struct {
	RecordId string
}

type ListDNS struct {
	PageNumber int
	PageSize   int
	KeyWord    string
	GroupId    string
}

type ListResourceDNSRecordSets struct {
	DomainName   string
	PageNumber   int
	PageSize     int
	RRKeyWord    string
	TypeKeyWord  string
	ValueKeyWord string
}
