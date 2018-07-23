PACKAGE DOCUMENTATION

package googlenotification
    import "."


TYPES

type Googlenotification struct {
}
    Googlenotification struct represents Google notification service and
    methods associates with it.

func (googlenotification *Googlenotification) CreateTopic(request interface{}) (resp interface{}, err error)
    CreateTopic creates tpoic

func (googlenotification *Googlenotification) DeleteTopic(request interface{}) (resp interface{}, err error)
    DeleteTopic delete topic

func (googlenotification *Googlenotification) GetTopic(request interface{}) (resp interface{}, err error)
    GetTopic gets topic

func (googlenotification *Googlenotification) ListTopic(request interface{}) (resp interface{}, err error)
    ListTopic list topic


