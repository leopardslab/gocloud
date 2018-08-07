```
package vultrbaremetal
    import "github.com/cloudlibz/gocloud/baremetal/vultrbaremetal"
```

### TYPES

```
type BareMetalInfo struct {
    SUBID           string
    OS              string  `json:"os"`
    RAM             string  `json:"ram"`
    Disk            string  `json:"disk"`
    MainIP          string  `json:"main_ip"`
    CPUCount        float64 `json:"cpu_count"`
    Location        string  `json:"location"`
    DCID            string
    DefaultPassword string `json:"default_password"`
    DateCreated     string `json:"date_created"`
    Status          string `json:"status"`
    NetmaskV4       string `json:"netmask_v4"`
    GatewayV4       string `json:"gateway_v4"`
    METALPLANID     float64
    V6Networks      []V6Network `json:"v6_networks"`
    Label           string      `json:"label"`
    Tag             string      `json:"tag"`
    OSID            string
    APPID           string
}

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

type CreateBareMetalResp struct {
    StatusCode int
    SUBID      string
}

func ParseCreateBareMetalResp(resp interface{}) (createBareMetalResp CreateBareMetalResp, err error)

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

type ListBareMetalResp struct {
    StatusCode     int
    BareMetalSlice []BareMetalInfo
}

func ParseListBareMetalResp(resp interface{}) (listBareMetalResp ListBareMetalResp, err error)


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

type V6Network struct {
    V6Network     string  `json:"v6_network"`
    V6MainIP      string  `json:"v6_main_ip"`
    V6NetworkSize float64 `json:"v6_network_size"`
}

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

