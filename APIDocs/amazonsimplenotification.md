PACKAGE DOCUMENTATION

package amazonsimplenotification
    import "."


TYPES

type Amazonsimplenotification struct {
}
    Amazonsimplenotification struct represents Amazon simple notification
    service and methods associates with it.

func (amazonsimplenotification *Amazonsimplenotification) CreateTopic(request interface{}) (resp interface{}, err error)
    CreateTopic create topic

func (amazonsimplenotification *Amazonsimplenotification) DeleteTopic(request interface{}) (resp interface{}, err error)
    DeleteTopic deletetopic

func (amazonsimplenotification *Amazonsimplenotification) GetTopic(request interface{}) (resp interface{}, err error)
    GetTopic gettopic

func (amazonsimplenotification *Amazonsimplenotification) ListTopic(request interface{}) (resp interface{}, err error)
    ListTopic list topic

func (amazonsimplenotification *Amazonsimplenotification) PrepareSignatureV2query(params map[string]string, Region string, response map[string]interface{}) error


