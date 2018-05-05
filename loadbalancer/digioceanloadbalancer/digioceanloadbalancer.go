package digioceanloadbalancer

// LoadBalancer represents the configuration to be applied to an existing or a new load balancer.
type LoadBalancer struct {
	Name                string           `json:"name,omitempty"`
	Algorithm           string           `json:"algorithm,omitempty"`
	Region              string           `json:"region,omitempty"`
	ForwardingRules     []ForwardingRule `json:"forwardingRules,omitempty"`
	HealthCheck         HealthCheck      `json:"healthCheck,omitempty"`
	StickySessions      StickySessions   `json:"stickySessions,omitempty"`
	DropletIDs          []int            `json:"dropletIds,omitempty"`
	Tag                 string           `json:"tag,omitempty"`
	RedirectHttpToHttps bool             `json:"redirectHTTPToHTTPS,omitempty"`
}

// ForwardingRule represents load balancer forwarding rules.
type ForwardingRule struct {
	EntryProtocol  string `json:"entryProtocol,omitempty"`
	EntryPort      int    `json:"entryPort,omitempty"`
	TargetProtocol string `json:"targetProtocol,omitempty"`
	TargetPort     int    `json:"targetPort,omitempty"`
	CertificateID  string `json:"certificateId,omitempty"`
	TlsPassthrough bool   `json:"TLSPassthrough,omitempty"`
}

// HealthCheck represents optional load balancer health check rules.
type HealthCheck struct {
	Protocol               string `json:"protocol,omitempty"`
	Port                   int    `json:"port,omitempty"`
	Path                   string `json:"path,omitempty"`
	CheckIntervalSeconds   int    `json:"checkIntervalSeconds,omitempty"`
	ResponseTimeoutSeconds int    `json:"responseTimeoutSeconds,omitempty"`
	HealthyThreshold       int    `json:"healthyThreshold,omitempty"`
	UnhealthyThreshold     int    `json:"unhealthyThreshold,omitempty"`
}

// StickySessions represents optional load balancer session affinity rules.
type StickySessions struct {
	Type             string `json:"typeStickySessions,omitempty"`
	CookieName       string `json:"cookieName,omitempty"`
	CookieTtlSeconds int    `json:"cookieTTLSeconds,omitempty"`
}
