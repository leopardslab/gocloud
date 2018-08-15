```
package vultrcompute
    import "github.com/cloudlibz/gocloud/compute/vultrcompute"
```

### TYPES

```
type VultrCompute struct {
}
```
VultrCompute struct

```
func (vultrCompute *VultrCompute) CreateNode(request interface{}) (resp interface{}, err error)
```
CreateNode function creates a new VultrCompute instance.

```
func (vultrCompute *VultrCompute) DeleteNode(request interface{}) (resp interface{}, err error)
```
DeleteNode function deletes a VultrCompute instance.

```
func (vultrCompute *VultrCompute) ListNode() (resp interface{}, err error)
```
ListNode function lists VultrCompute instances.

```
func (vultrCompute *VultrCompute) RebootNode(request interface{}) (resp interface{}, err error)
```
RebootNode function reboots a VultrCompute instance.

```
func (vultrCompute *VultrCompute) StartNode(request interface{}) (resp interface{}, err error)
```
StartNode function starts a VultrCompute instance.

```
func (vultrCompute *VultrCompute) StopNode(request interface{}) (resp interface{}, err error)
```
StopNode function stops a VultrCompute instance.

```
type CreateNode struct {
    DCID      int // Location to create this virtual machine in.  See v1/regions/list
    VPSPLANID int // Plan to use when creating this virtual machine.  See v1/plans/list
    OSID      int // Operating system to use.  See v1/os/list

    ISOID      string // (optional)  If you've selected the 'custom' operating system, this is the ID of a specific ISO to mount during the deployment
    SCRIPTID   int    // (optional) If you've not selected a 'custom' operating system, this can be the SCRIPTID of a startup script to execute on boot.  See v1/startupscript/list
    SNAPSHOTID string // (optional) If you've selected the 'snapshot' operating system, this should be the SNAPSHOTID (see v1/snapshot/list) to restore for the initial installation.

    NETWORKID []string // (optional) List of private networks to attach to this server.Use either this field or enable_private_network, not both

    SSHKEYID string // (optional) List of SSH keys to apply to this server on install (only valid for Linux/FreeBSD).See v1/sshkey/list.Separate keys with commas

    APPID int // (optional)  If launching an application (OSID 186), this is the APPID to launch.See v1/app/list.

    FIREWALLGROUPID string // (optional) The firewall group to assign to this server. See /v1/firewall/group_list.
    // contains filtered or unexported fields
}

type CreateNodeBuilder struct {
    // contains filtered or unexported fields
}
    CreateNode builder pattern code

func NewCreateNodeBuilder() *CreateNodeBuilder

func (b *CreateNodeBuilder) APPID(aPPID int) *CreateNodeBuilder

func (b *CreateNodeBuilder) AutoBackups(auto_backups string) *CreateNodeBuilder

func (b *CreateNodeBuilder) Build() (map[string]interface{}, error)

func (b *CreateNodeBuilder) DCID(dCID int) *CreateNodeBuilder

func (b *CreateNodeBuilder) DdosProtection(ddos_protection string) *CreateNodeBuilder

func (b *CreateNodeBuilder) EnableIpv6(enable_ipv6 string) *CreateNodeBuilder

func (b *CreateNodeBuilder) EnablePrivateNetwork(enable_private_network string) *CreateNodeBuilder

func (b *CreateNodeBuilder) FIREWALLGROUPID(fIREWALLGROUPID string) *CreateNodeBuilder

func (b *CreateNodeBuilder) Hostname(hostname string) *CreateNodeBuilder

func (b *CreateNodeBuilder) ISOID(iSOID string) *CreateNodeBuilder

func (b *CreateNodeBuilder) IpxeChainUrl(ipxe_chain_url string) *CreateNodeBuilder

func (b *CreateNodeBuilder) Label(label string) *CreateNodeBuilder

func (b *CreateNodeBuilder) NETWORKID(nETWORKID []string) *CreateNodeBuilder

func (b *CreateNodeBuilder) NotifyActivate(notify_activate string) *CreateNodeBuilder

func (b *CreateNodeBuilder) OSID(oSID int) *CreateNodeBuilder

func (b *CreateNodeBuilder) ReservedIpV4(reserved_ip_v4 string) *CreateNodeBuilder

func (b *CreateNodeBuilder) SCRIPTID(sCRIPTID int) *CreateNodeBuilder

func (b *CreateNodeBuilder) SNAPSHOTID(sNAPSHOTID string) *CreateNodeBuilder

func (b *CreateNodeBuilder) SSHKEYID(sSHKEYID string) *CreateNodeBuilder

func (b *CreateNodeBuilder) Tag(tag string) *CreateNodeBuilder

func (b *CreateNodeBuilder) UserData(userdata string) *CreateNodeBuilder

func (b *CreateNodeBuilder) VPSPLANID(vPSPLANID int) *CreateNodeBuilder

type DeleteNode struct {
    SUBID int // Unique identifier for this subscription.  These can be found using the v1/server/list call.
}

type DeleteNodeBuilder struct {
    // contains filtered or unexported fields
}
    DeleteNode builder pattern code

func NewDeleteNodeBuilder() *DeleteNodeBuilder

func (b *DeleteNodeBuilder) Build() (map[string]interface{}, error)

func (b *DeleteNodeBuilder) SUBID(sUBID int) *DeleteNodeBuilder

type ListNode struct {
    SUBID int // (optional) Unique identifier of a subscription. Only the subscription object will be returned.
    // contains filtered or unexported fields
}

type ListNodeBuilder struct {
    // contains filtered or unexported fields
}
    ListNode builder pattern code

func NewListNodeBuilder() *ListNodeBuilder

func (b *ListNodeBuilder) Build() (map[string]interface{}, error)

func (b *ListNodeBuilder) SUBID(sUBID int) *ListNodeBuilder

type RebootNode struct {
    SUBID int
}

type RebootNodeBuilder struct {
    // contains filtered or unexported fields
}
    RebootNode builder pattern code

func NewRebootNodeBuilder() *RebootNodeBuilder

func (b *RebootNodeBuilder) Build() (map[string]interface{}, error)

func (b *RebootNodeBuilder) SUBID(sUBID int) *RebootNodeBuilder

type StartNode struct {
    SUBID int
}

type StartNodeBuilder struct {
    // contains filtered or unexported fields
}
    StartNode builder pattern code

func NewStartNodeBuilder() *StartNodeBuilder

func (b *StartNodeBuilder) Build() (map[string]interface{}, error)

func (b *StartNodeBuilder) SUBID(sUBID int) *StartNodeBuilder

```