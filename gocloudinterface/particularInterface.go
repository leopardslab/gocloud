package gocloudinterface

// Serverless module unified API
type Serverless interface {
	GetFunction(request interface{}) (resp interface{}, err error)
	CreateFunction(request interface{}) (resp interface{}, err error)
	CallFunction(request interface{}) (resp interface{}, err error)
	ListFunction(request interface{}) (resp interface{}, err error)
	DeleteFunction(request interface{}) (resp interface{}, err error)
}

// Database module unified API
type Database interface {
	ListTables(request interface{}) (resp interface{}, err error)
	DeleteTables(request interface{}) (resp interface{}, err error)
	DescribeTables(request interface{}) (resp interface{}, err error)
	CreateTables(request interface{}) (resp interface{}, err error)
}

// MachineLearning module unified API
type MachineLearning interface {
	CreateMLModel(request interface{}) (resp interface{}, err error)
	DeleteMLModel(request interface{}) (resp interface{}, err error)
	GetMLModel(request interface{}) (resp interface{}, err error)
	UpdateMLModel(request interface{}) (resp interface{}, err error)
}

// Bare Metal module unified API
type BareMetal interface {
	CreateBareMetal(request interface{}) (resp interface{}, err error)
	DeleteBareMetal(request interface{}) (resp interface{}, err error)
	RebootBareMetal(request interface{}) (resp interface{}, err error)
	ReinstallBareMetal(request interface{}) (resp interface{}, err error)
	HaltBareMetal(request interface{}) (resp interface{}, err error)
	ListBareMetal(request interface{}) (resp interface{}, err error)
}


// Analytics module unified API
type Analytics interface {
	UpdateDatasets(request interface{}) (resp interface{}, err error)
	DeleteDatasets(request interface{}) (resp interface{}, err error)
	CreateDatasets(request interface{}) (resp interface{}, err error)
	ListDatasets(request interface{}) (resp interface{}, err error)
	GetDatasets(request interface{}) (resp interface{}, err error)
}

// Bare Metal module unified API
type Notification interface {
	GetTopic(request interface{}) (resp interface{}, err error)
	ListTopic(request interface{}) (resp interface{}, err error)
	DeleteTopic(request interface{}) (resp interface{}, err error)
	CreateTopic(request interface{}) (resp interface{}, err error)
}
