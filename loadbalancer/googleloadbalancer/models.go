package googleloadbalancer

//Googleloadbalancer reperents google loadbalancer methods and attributes.
type Googleloadbalancer struct {
}

//TargetPools reperents google loadbalancer.
type TargetPools struct {
	BackupPool        string   `json:"backupPool"`
	CreationTimestamp string   `json:"creationTimestamp"`
	Description       string   `json:"description"`
	HealthChecks      []string `json:"healthChecks"`
	FailoverRatio     int      `json:"failoverRatio"`
	ID                string   `json:"id"`
	Instances         []string `json:"instances"`
	Kind              string   `json:"kind"`
	Name              string   `json:"name"`
	Region            string   `json:"region"`
	SelfLink          string   `json:"selfLink"`
	SessionAffinity   string   `json:"sessionAffinity"`
}
