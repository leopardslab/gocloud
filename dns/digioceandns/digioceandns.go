package digioceandns

// Digioceandns struct represents a DigitalOcean DNS service.
type Digioceandns struct {
	Type       string `json:"type"`
	Name       string `json:"name"`
	Data       string `json:"data"`
	Priority   int    `json:priority,omitempty`
	Port       int    `json:port,omitempty`
	TimeToLive int    `json:ttl`
	Weight     int    `json:weight,omitempty`
	Flags      int    `json:flags,omitempty`
	Tag        string `json:"tag"`
}
