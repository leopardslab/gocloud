package googlecloudfunctions

const (
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	RFC3339  = "2006-01-02T15:04:05Z07:00"
)

//Googlecloudfunctions struct reperesnts Google cloud functions service.
type Googlecloudfunctions struct {
}

//CreateGooglecloudfunction struct reperesnts Create Googlecloud function.
type CreateGooglecloudfunction struct {
	Name                string       `json:"name"`
	Status              string       `json:"status"`
	EntryPoint          string       `json:"entryPoint"`
	Timeout             string       `json:"timeout"`
	AvailableMemoryMb   int          `json:"availableMemoryMb"`
	ServiceAccountEmail string       `json:"serviceAccountEmail"`
	UpdateTime          string       `json:"updateTime"`
	VersionID           string       `json:"versionId"`
	SourceUploadURL     string       `json:"sourceUploadUrl"`
	Labels              labels       `json:"labels"`
	HTTPSTrigger        httpstrigger `json:"httpsTrigger"`
}

//labels struct reperesnts Create CreateGooglecloudfunction parameters.
type labels struct {
	DeploymentTool string `json:"deployment-tool"`
}

//httpstrigger struct reperesnts Create CreateGooglecloudfunction parameters.
type httpstrigger struct {
	URL string `json:"url"`
}
