package google
    import "github.com/scorelab/gocloud-v2/google"


TYPES

type Google struct {
    gce.GCE
    googlestorage.GoogleStorage
    googleloadbalancer.Googleloadbalancer
    googlecontainer.Googlecontainer
    googledns.Googledns
}
    Google struct represents google cloud provider.


