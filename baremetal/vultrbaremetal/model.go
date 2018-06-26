package vultrbaremetal

type VultrBareMetal struct {
}

type CreateBareMetal struct {
	DCID            int    // Location in which to create the server. See v1/regions/list.
	METALPLANID     int    // Plan to use when creating this server. See v1/plans/list_baremetal.
	OSID            int    // Operating system to use. See v1/os/list.
	SCRIPTID        int    // (optional) The SCRIPTID of a startup script to execute on boot. This only works when using a Vultr supplied operating system. See v1/startupscript/list.
	SNAPSHOTID      string // (optional) If you've selected the 'snapshot' operating system, this should be the SNAPSHOTID (see v1/snapshot/list) to restore for the initial installation.
	enable_ipv6     string // (optional) 'yes' or 'no'.  If yes, an IPv6 subnet will be assigned to the server.
	label           string // (optional) This is a text label that will be shown in the control panel.
	SSHKEYID        string // (optional) List of SSH keys to apply to this server on install (only valid for Linux/FreeBSD). See v1/sshkey/list. Separate keys with commas.
	APPID           int    // (optional) If launching an application (OSID 186), this is the APPID to launch. See v1/app/list.
	userdata        string // (optional) Base64 encoded user-data.
	notify_activate string // (optional, default 'yes') 'yes' or 'no'. If yes, an activation email will be sent when the server is ready.
	hostname        string // (optional) The hostname to assign to this server.
	tag             string // (optional) The tag to assign to this server.
}

type DeleteBareMetal struct {
	SUBID int // Unique identifier for this subscription.
}

type RebootBareMetal struct {
	SUBID int // Unique identifier for this subscription.
}

type ReinstallBareMetal struct {
	SUBID int // Unique identifier for this subscription.
}

type HaltBareMetal struct {
	SUBID int // Unique identifier for this subscription.
}

type ListBareMetal struct {
	SUBID   int    // (optional) Unique identifier of a subscription. Only the subscription object will be returned.
	tag     string // (optional) A tag string. Only subscription objects with this tag will be returned.
	label   string // (optional) A text label string. Only subscription objects with this text label will be returned.
	main_ip string // (optional) An IPv4 address. Only the subscription matching this IPv4 address will be returned.
}
