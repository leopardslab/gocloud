package gocloudinterface

// Serverless module unified API
type Serverless interface {
	Getfunction(request interface{}) (resp interface{}, err error)
	Createfunction(request interface{}) (resp interface{}, err error)
	Callfunction(request interface{}) (resp interface{}, err error)
	Listfunction(request interface{}) (resp interface{}, err error)
	Deletefunction(request interface{}) (resp interface{}, err error)
}

// Database module unified API
type Database interface {
	Listtables(request interface{}) (resp interface{}, err error)
	Deletetables(request interface{}) (resp interface{}, err error)
	Describetables(request interface{}) (resp interface{}, err error)
	Createtables(request interface{}) (resp interface{}, err error)
}

// MachineLearning module unified API
type MachineLearning interface {
	CreateMLModel(request interface{}) (resp interface{}, err error)
	DeleteMLModel(request interface{}) (resp interface{}, err error)
	GetMLModel(request interface{}) (resp interface{}, err error)
	UpdateMLModel(request interface{}) (resp interface{}, err error)
}
