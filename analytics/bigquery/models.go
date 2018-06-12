package bigquery

//Bigquery struct reperesnts google analytics service.
type Bigquery struct {
}

type View  struct {
  DatasetID string `json:"datasetId"`
  ProjectID string `json:"projectId"`
  TableID   string `json:"tableId"`
}

type Access struct {
  Domain       string `json:"domain"`
  GroupByEmail string `json:"groupByEmail"`
  Role         string `json:"role"`
  SpecialGroup string `json:"specialGroup"`
  UserByEmail  string `json:"userByEmail"`
  view View  `json:"view"`
  }

type DatasetReference struct {
    DatasetID string `json:"datasetId"`
    ProjectID string `json:"projectId"`
  }

type CreateDatasets struct {
	access []Access `json:"access"`
	datasetReference DatasetReference `json:"datasetReference"`

	defaultTableExpirationMs string `json:"defaultTableExpirationMs"`
	description              string `json:"description"`
	etag                     string `json:"etag"`
	friendlyName             string `json:"friendlyName"`
	id                       string `json:"id"`
	kind                     string `json:"kind"`
	lastModifiedTime string `json:"lastModifiedTime"`
	location         string `json:"location"`
	selfLink         string `json:"selfLink"`
}
