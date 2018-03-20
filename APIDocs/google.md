package google
    import "github.com/shlokgilda/gocloud/google"


TYPES

type Google struct {
    gce.GCE
    googlestorage.GoogleStorage
    googleloadbalancer.Googleloadbalancer
    googlecontainer.Googlecontainer
    googledns.Googledns
}
    Google struct represents google cloud provider.


