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

type CreateBareMetalBuilder struct {
    // contains filtered or unexported fields
}
    CreateBareMetal builder pattern code

func NewCreateBareMetalBuilder() *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) APPID(aPPID int) *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) Build() (map[string]interface{}, error)

func (b *CreateBareMetalBuilder) DCID(dCID int) *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) EnableIpv6(enable_ipv6 string) *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) Hostname(hostname string) *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) Label(label string) *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) METALPLANID(mETALPLANID int) *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) NotifyActivate(notify_activate string) *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) OSID(oSID int) *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) SCRIPTID(sCRIPTID int) *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) SNAPSHOTID(sNAPSHOTID string) *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) SSHKEYID(sSHKEYID string) *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) Tag(tag string) *CreateBareMetalBuilder

func (b *CreateBareMetalBuilder) UserData(userdata string) *CreateBareMetalBuilder

type DeleteBareMetal struct {
    SUBID int // Unique identifier for this subscription.
}

type DeleteBareMetalBuilder struct {
    // contains filtered or unexported fields
}
    DeleteBareMetal builder pattern code

func NewDeleteBareMetalBuilder() *DeleteBareMetalBuilder

func (b *DeleteBareMetalBuilder) Build() (map[string]interface{}, error)

func (b *DeleteBareMetalBuilder) SUBID(sUBID int) *DeleteBareMetalBuilder

type HaltBareMetal struct {
    SUBID int // Unique identifier for this subscription.
}

type HaltBareMetalBuilder struct {
    // contains filtered or unexported fields
}
    HaltBareMetal builder pattern code

func NewHaltBareMetalBuilder() *HaltBareMetalBuilder

func (b *HaltBareMetalBuilder) Build() (map[string]interface{}, error)

func (b *HaltBareMetalBuilder) SUBID(sUBID int) *HaltBareMetalBuilder

type ListBareMetal struct {
    SUBID int // (optional) Unique identifier of a subscription. Only the subscription object will be returned.
    // contains filtered or unexported fields
}

type ListBareMetalBuilder struct {
    // contains filtered or unexported fields
}
    ListBareMetal builder pattern code

func NewListBareMetalBuilder() *ListBareMetalBuilder

func (b *ListBareMetalBuilder) Build() (map[string]interface{}, error)

func (b *ListBareMetalBuilder) Label(label string) *ListBareMetalBuilder

func (b *ListBareMetalBuilder) MainIp(main_ip string) *ListBareMetalBuilder

func (b *ListBareMetalBuilder) SUBID(sUBID int) *ListBareMetalBuilder

func (b *ListBareMetalBuilder) Tag(tag string) *ListBareMetalBuilder

type RebootBareMetal struct {
    SUBID int // Unique identifier for this subscription.
}

type RebootBareMetalBuilder struct {
    // contains filtered or unexported fields
}
    RebootBareMetal builder pattern code

func NewRebootBareMetalBuilder() *RebootBareMetalBuilder

func (b *RebootBareMetalBuilder) Build() (map[string]interface{}, error)

func (b *RebootBareMetalBuilder) SUBID(sUBID int) *RebootBareMetalBuilder

type ReinstallBareMetal struct {
    SUBID int // Unique identifier for this subscription.
}

type ReinstallBareMetalBuilder struct {
    // contains filtered or unexported fields
}
    ReinstallBareMetal builder pattern code

func NewReinstallBareMetalBuilder() *ReinstallBareMetalBuilder

func (b *ReinstallBareMetalBuilder) Build() (map[string]interface{}, error)

func (b *ReinstallBareMetalBuilder) SUBID(sUBID int) *ReinstallBareMetalBuilder

type VultrBareMetal struct {
}

func (*VultrBareMetal) CreateBareMetal(request interface{}) (resp interface{}, err error)
    CreateBareMetal function creates a new Vultr bare metal machine.

func (*VultrBareMetal) DeleteBareMetal(request interface{}) (resp interface{}, err error)
    DeleteBareMetal function deletes a Vultr bare metal machine.

func (*VultrBareMetal) HaltBareMetal(request interface{}) (resp interface{}, err error)
    HaltBareMetal function halt a Vultr bare metal machine.

func (*VultrBareMetal) ListBareMetal(request interface{}) (resp interface{}, err error)
    ListBareMetal function list Vultr bare metal machines.

func (*VultrBareMetal) RebootBareMetal(request interface{}) (resp interface{}, err error)
    RebootBareMetal function reboots a Vultr bare metal machine.

func (*VultrBareMetal) ReinstallBareMetal(request interface{}) (resp interface{}, err error)
    ReinstallBareMetal function reinstall a Vultr bare metal machine.
```
