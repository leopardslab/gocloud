```
package vultrbaremetal
    import "github.com/cloudlibz/gocloud/baremetal/vultrbaremetal"
```

TYPES

```
type CreateBareMetal struct {
    DCID        int    // Location in which to create the server. See v1/regions/list.
    METALPLANID int    // Plan to use when creating this server. See v1/plans/list_baremetal.
    OSID        int    // Operating system to use. See v1/os/list.
    SCRIPTID    int    // (optional) The SCRIPTID of a startup script to execute on boot. This only works when using a Vultr supplied operating system. See v1/startupscript/list.
    SNAPSHOTID  string // (optional) If you've selected the 'snapshot' operating system, this should be the SNAPSHOTID (see v1/snapshot/list) to restore for the initial installation.

    SSHKEYID string // (optional) List of SSH keys to apply to this server on install (only valid for Linux/FreeBSD). See v1/sshkey/list. Separate keys with commas.
    APPID    int    // (optional) If launching an application (OSID 186), this is the APPID to launch. See v1/app/list.
    // contains filtered or unexported fields
}

type DeleteBareMetal struct {
    SUBID int // Unique identifier for this subscription.
}

type HaltBareMetal struct {
    SUBID int // Unique identifier for this subscription.
}

type ListBareMetal struct {
    SUBID int // (optional) Unique identifier of a subscription. Only the subscription object will be returned.
    // contains filtered or unexported fields
}

type RebootBareMetal struct {
    SUBID int // Unique identifier for this subscription.
}

type ReinstallBareMetal struct {
    SUBID int // Unique identifier for this subscription.
}

type VultrBareMetal struct {
}

func (*VultrBareMetal) CreateBareMetal(request interface{}) (resp interface{}, err error)

func (*VultrBareMetal) DeleteBareMetal(request interface{}) (resp interface{}, err error)

func (*VultrBareMetal) HaltBareMetal(request interface{}) (resp interface{}, err error)

func (*VultrBareMetal) ListBareMetal(request interface{}) (resp interface{}, err error)

func (*VultrBareMetal) RebootBareMetal(request interface{}) (resp interface{}, err error)

func (*VultrBareMetal) ReinstallBareMetal(request interface{}) (resp interface{}, err error)
```
