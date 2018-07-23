PACKAGE DOCUMENTATION

package redshift
    import "."


FUNCTIONS

func PrepareSignaturequery(describeclusterspram map[string]string, region string) (response map[string]interface{}, err error)

TYPES

type CreateCluster struct {
    // contains filtered or unexported fields
}

type DeleteCluster struct {
    // contains filtered or unexported fields
}

type Describecluster struct {
    // contains filtered or unexported fields
}

type Redshift struct {
}
    redshift struct reperesnts aws alianalytics service.

func (redshift *Redshift) CreateDatasets(request interface{}) (resp interface{}, err error)
    CreateDatasets Create Datasets.

func (redshift *Redshift) DeleteDatasets(request interface{}) (resp interface{}, err error)
    DeleteDatasets delete Datasets.

func (redshift *Redshift) GetDatasets(request interface{}) (resp interface{}, err error)
    GetDatasets get Datasets.

func (redshift *Redshift) ListDatasets(request interface{}) (resp interface{}, err error)
    ListDatasets list Datasets.

func (redshift *Redshift) UpdateDatasets(request interface{}) (resp interface{}, err error)
    UpdateDatasets Update Datasets.


