PACKAGE DOCUMENTATION

package bigtable
    import "."


FUNCTIONS

func Createbigtabledictnoaryconvert(option Createbigtable, Createbigtablejsonmap map[string]interface{})
    Createbigtable json formation.

TYPES

type Bigtable struct {
}
    Createbigtable struct reperesnts Google bigtable.

func (bigtable *Bigtable) Createtables(request interface{}) (resp interface{}, err error)

func (bigtable *Bigtable) Deletetables(request interface{}) (resp interface{}, err error)

func (bigtable *Bigtable) Describetables(request interface{}) (resp interface{}, err error)
    describe describe tables.

func (bigtable *Bigtable) Listtables(request interface{}) (resp interface{}, err error)
    List list tables.

type ClusterStates struct {
    // contains filtered or unexported fields
}
    ClusterStates struct reperesnts ClusterStates.

type Createbigtable struct {
    ClusterStates ClusterStates
    // contains filtered or unexported fields
}
    Createbigtable struct reperesnts Create bigtable.

type GcRule struct {
    // contains filtered or unexported fields
}
    GcRule struct reperesnts GcRule.

type InitialSplits struct {
    // contains filtered or unexported fields
}
    InitialSplits struct reperesnts InitialSplits.

type Table struct {
    // contains filtered or unexported fields
}
    Table struct reperesnts Table.


