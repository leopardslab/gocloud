package vultrdns

type VultrDNS struct {
}

type CreateDNS struct {
	domain   string // Domain name to add record to
	name     string // Name (subdomain) of record
	Type     string // Type (A, AAAA, MX, etc) of record
	data     string // Data for this record
	ttl      int    // (optional) TTL of this record
	priority int    // (only required for MX and SRV) Priority of this record (omit the priority from the data)
}
