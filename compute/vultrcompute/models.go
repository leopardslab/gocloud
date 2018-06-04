package vultrcompute

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
