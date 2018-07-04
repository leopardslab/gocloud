package vultrcompute

import (
	"errors"
	"github.com/cloudlibz/gocloud/vultrauth"
)

// VultrCompute struct
type VultrCompute struct {
}

type CreateNode struct {
	DCID                   int      // Location to create this virtual machine in.  See v1/regions/list
	VPSPLANID              int      // Plan to use when creating this virtual machine.  See v1/plans/list
	OSID                   int      // Operating system to use.  See v1/os/list
	ipxe_chain_url         string   // (optional) If you've selected the 'custom' operating system, this can be set to chainload the specified URL on bootup, via iPXE
	ISOID                  string   // (optional)  If you've selected the 'custom' operating system, this is the ID of a specific ISO to mount during the deployment
	SCRIPTID               int      // (optional) If you've not selected a 'custom' operating system, this can be the SCRIPTID of a startup script to execute on boot.  See v1/startupscript/list
	SNAPSHOTID             string   // (optional) If you've selected the 'snapshot' operating system, this should be the SNAPSHOTID (see v1/snapshot/list) to restore for the initial installation.
	enable_ipv6            string   // (optional) 'yes' or 'no'.  If yes, an IPv6 subnet will be assigned to the machine (where available)
	enable_private_network string   // (optional) 'yes' or 'no'. If yes, private networking support will be added to the new server.
	NETWORKID              []string // (optional) List of private networks to attach to this server.Use either this field or enable_private_network, not both
	label                  string   // (optional) This is a text label that will be shown in the control panel
	SSHKEYID               string   // (optional) List of SSH keys to apply to this server on install (only valid for Linux/FreeBSD).See v1/sshkey/list.Separate keys with commas
	auto_backups           string   // (optional) 'yes' or 'no'.  If yes, automatic backups will be enabled for this server (these have an extra charge associated with them)
	APPID                  int      // (optional)  If launching an application (OSID 186), this is the APPID to launch.See v1/app/list.
	userdata               string   // (optional)  Base64 encoded user-data
	notify_activate        string   // (optional, default 'yes') 'yes' or 'no'. If yes, an activation email will be sent when the server is ready.
	ddos_protection        string   // (optional, default 'no') 'yes' or 'no'.  If yes, DDOS protection will be enabled on the subscription (there is an additional charge for this)
	reserved_ip_v4         string   // (optional) IP address of the floating IP to use as the main IP of this server
	hostname               string   // (optional) The hostname to assign to this server.
	tag                    string   // (optional) The tag to assign to this server.
	FIREWALLGROUPID        string   // (optional) The firewall group to assign to this server. See /v1/firewall/group_list.
}

type DeleteNode struct {
	SUBID int // Unique identifier for this subscription.  These can be found using the v1/server/list call.
}

// CreateNode builder pattern code
type CreateNodeBuilder struct {
	createNode *CreateNode
}

func NewCreateNodeBuilder() *CreateNodeBuilder {
	createNode := &CreateNode{}
	b := &CreateNodeBuilder{createNode: createNode}
	return b
}

func (b *CreateNodeBuilder) DCID(dCID int) *CreateNodeBuilder {
	b.createNode.DCID = dCID
	return b
}

func (b *CreateNodeBuilder) VPSPLANID(vPSPLANID int) *CreateNodeBuilder {
	b.createNode.VPSPLANID = vPSPLANID
	return b
}

func (b *CreateNodeBuilder) OSID(oSID int) *CreateNodeBuilder {
	b.createNode.OSID = oSID
	return b
}

func (b *CreateNodeBuilder) IpxeChainUrl(ipxe_chain_url string) *CreateNodeBuilder {
	b.createNode.ipxe_chain_url = ipxe_chain_url
	return b
}

func (b *CreateNodeBuilder) ISOID(iSOID string) *CreateNodeBuilder {
	b.createNode.ISOID = iSOID
	return b
}

func (b *CreateNodeBuilder) SCRIPTID(sCRIPTID int) *CreateNodeBuilder {
	b.createNode.SCRIPTID = sCRIPTID
	return b
}

func (b *CreateNodeBuilder) SNAPSHOTID(sNAPSHOTID string) *CreateNodeBuilder {
	b.createNode.SNAPSHOTID = sNAPSHOTID
	return b
}

func (b *CreateNodeBuilder) EnableIpv6(enable_ipv6 string) *CreateNodeBuilder {
	b.createNode.enable_ipv6 = enable_ipv6
	return b
}

func (b *CreateNodeBuilder) EnablePrivateNetwork(enable_private_network string) *CreateNodeBuilder {
	b.createNode.enable_private_network = enable_private_network
	return b
}

func (b *CreateNodeBuilder) NETWORKID(nETWORKID []string) *CreateNodeBuilder {
	b.createNode.NETWORKID = nETWORKID
	return b
}

func (b *CreateNodeBuilder) Label(label string) *CreateNodeBuilder {
	b.createNode.label = label
	return b
}

func (b *CreateNodeBuilder) SSHKEYID(sSHKEYID string) *CreateNodeBuilder {
	b.createNode.SSHKEYID = sSHKEYID
	return b
}

func (b *CreateNodeBuilder) AutoBackups(auto_backups string) *CreateNodeBuilder {
	b.createNode.auto_backups = auto_backups
	return b
}

func (b *CreateNodeBuilder) APPID(aPPID int) *CreateNodeBuilder {
	b.createNode.APPID = aPPID
	return b
}

func (b *CreateNodeBuilder) UserData(userdata string) *CreateNodeBuilder {
	b.createNode.userdata = userdata
	return b
}

func (b *CreateNodeBuilder) NotifyActivate(notify_activate string) *CreateNodeBuilder {
	b.createNode.notify_activate = notify_activate
	return b
}

func (b *CreateNodeBuilder) DdosProtection(ddos_protection string) *CreateNodeBuilder {
	b.createNode.ddos_protection = ddos_protection
	return b
}

func (b *CreateNodeBuilder) ReservedIpV4(reserved_ip_v4 string) *CreateNodeBuilder {
	b.createNode.reserved_ip_v4 = reserved_ip_v4
	return b
}

func (b *CreateNodeBuilder) Hostname(hostname string) *CreateNodeBuilder {
	b.createNode.hostname = hostname
	return b
}

func (b *CreateNodeBuilder) Tag(tag string) *CreateNodeBuilder {
	b.createNode.tag = tag
	return b
}

func (b *CreateNodeBuilder) FIREWALLGROUPID(fIREWALLGROUPID string) *CreateNodeBuilder {
	b.createNode.FIREWALLGROUPID = fIREWALLGROUPID
	return b
}

func (b *CreateNodeBuilder) Build() (map[string]interface{}, error) {
	if b.createNode.DCID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "DCID")
	}
	if b.createNode.VPSPLANID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "VPSPLANID")
	}
	if b.createNode.OSID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "OSID")
	}

	params := make(map[string]interface{})

	params["DCID"] = b.createNode.DCID
	params["VPSPLANID"] = b.createNode.VPSPLANID
	params["OSID"] = b.createNode.OSID

	if b.createNode.ipxe_chain_url != "" {
		params["ipxe_chain_url"] = b.createNode.ipxe_chain_url
	}
	if b.createNode.ISOID != "" {
		params["ISOID"] = b.createNode.ISOID
	}
	if b.createNode.SCRIPTID != 0 {
		params["SCRIPTID"] = b.createNode.SCRIPTID
	}
	if b.createNode.SNAPSHOTID != "" {
		params["SNAPSHOTID"] = b.createNode.SNAPSHOTID
	}
	if b.createNode.enable_ipv6 != "" {
		params["enable_ipv6"] = b.createNode.enable_ipv6
	}
	if b.createNode.enable_private_network != "" {
		params["enable_private_network"] = b.createNode.enable_private_network
	}
	if len(b.createNode.NETWORKID) != 0 {
		params["NETWORKID"] = b.createNode.NETWORKID
	}
	if b.createNode.label != "" {
		params["label"] = b.createNode.label
	}
	if b.createNode.SSHKEYID != "" {
		params["SSHKEYID"] = b.createNode.SSHKEYID
	}
	if b.createNode.auto_backups != "" {
		params["auto_backups"] = b.createNode.auto_backups
	}
	if b.createNode.APPID != 0 {
		params["APPID"] = b.createNode.APPID
	}
	if b.createNode.userdata != "" {
		params["userdata"] = b.createNode.userdata
	}
	if b.createNode.notify_activate != "" {
		params["notify_activate"] = b.createNode.notify_activate
	}
	if b.createNode.ddos_protection != "" {
		params["ddos_protection"] = b.createNode.ddos_protection
	}
	if b.createNode.reserved_ip_v4 != "" {
		params["reserved_ip_v4"] = b.createNode.reserved_ip_v4
	}
	if b.createNode.hostname != "" {
		params["hostname"] = b.createNode.hostname
	}
	if b.createNode.tag != "" {
		params["tag"] = b.createNode.tag
	}
	if b.createNode.FIREWALLGROUPID != "" {
		params["FIREWALLGROUPID"] = b.createNode.FIREWALLGROUPID
	}

	return params, nil
}
