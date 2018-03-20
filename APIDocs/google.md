package google
    import "github.com/cloudlibz/gocloud/google"


TYPES

type Google struct {
    gce.GCE
    googlestorage.GoogleStorage
    googleloadbalancer.Googleloadbalancer
    googlecontainer.Googlecontainer
    googledns.Googledns
}
    Google struct represents google cloud provider.


