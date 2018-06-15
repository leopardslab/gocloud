package bigquery

//Bigquery struct reperesnts google analytics service.
type Bigquery struct {
}

//View  struct reperesnts CreateDatasets internal struct.
type View struct {
	datasetID string `json:"datasetId"`
	projectID string `json:"projectId"`
	tableID   string `json:"tableId"`
}

//Access struct reperesnts CreateDatasets internal struct.
type Access struct {
	domain       string `json:"domain"`
	groupByEmail string `json:"groupByEmail"`
	role         string `json:"role"`
	specialGroup string `json:"specialGroup"`
	userByEmail  string `json:"userByEmail"`
	view         View   `json:"view"`
}

//DatasetReference struct reperesnts CreateDatasets internal struct.
type DatasetReference struct {
	datasetID string `json:"datasetId"`
	projectID string `json:"projectId"`
}

//CreateDatasets struct reperesnts google bigquery CreateDatasets service struct.
type CreateDatasets struct {
	access                   []Access         `json:"access"`
	datasetReference         DatasetReference `json:"datasetReference"`
	defaultTableExpirationMs string           `json:"defaultTableExpirationMs"`
	description              string           `json:"description"`
	etag                     string           `json:"etag"`
	friendlyName             string           `json:"friendlyName"`
	id                       string           `json:"id"`
	kind                     string           `json:"kind"`
	lastModifiedTime         string           `json:"lastModifiedTime"`
	creationTime						 string           `json:"creationTime"`
	location                 string           `json:"location"`
	selfLink                 string           `json:"selfLink"`
}
