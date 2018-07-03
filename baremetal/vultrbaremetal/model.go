package vultrbaremetal

import (
	"errors"
	"github.com/cloudlibz/gocloud/aliauth"
	"github.com/cloudlibz/gocloud/vultrauth"
)

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

// CreateBareMetal builder pattern code
type CreateBareMetalBuilder struct {
	createBareMetal *CreateBareMetal
}

func NewCreateBareMetalBuilder() *CreateBareMetalBuilder {
	createBareMetal := &CreateBareMetal{}
	b := &CreateBareMetalBuilder{createBareMetal: createBareMetal}
	return b
}

func (b *CreateBareMetalBuilder) DCID(dCID int) *CreateBareMetalBuilder {
	b.createBareMetal.DCID = dCID
	return b
}

func (b *CreateBareMetalBuilder) METALPLANID(mETALPLANID int) *CreateBareMetalBuilder {
	b.createBareMetal.METALPLANID = mETALPLANID
	return b
}

func (b *CreateBareMetalBuilder) OSID(oSID int) *CreateBareMetalBuilder {
	b.createBareMetal.OSID = oSID
	return b
}

func (b *CreateBareMetalBuilder) SCRIPTID(sCRIPTID int) *CreateBareMetalBuilder {
	b.createBareMetal.SCRIPTID = sCRIPTID
	return b
}

func (b *CreateBareMetalBuilder) SNAPSHOTID(sNAPSHOTID string) *CreateBareMetalBuilder {
	b.createBareMetal.SNAPSHOTID = sNAPSHOTID
	return b
}

func (b *CreateBareMetalBuilder) EnableIpv6(enable_ipv6 string) *CreateBareMetalBuilder {
	b.createBareMetal.enable_ipv6 = enable_ipv6
	return b
}

func (b *CreateBareMetalBuilder) Label(label string) *CreateBareMetalBuilder {
	b.createBareMetal.label = label
	return b
}

func (b *CreateBareMetalBuilder) SSHKEYID(sSHKEYID string) *CreateBareMetalBuilder {
	b.createBareMetal.SSHKEYID = sSHKEYID
	return b
}

func (b *CreateBareMetalBuilder) APPID(aPPID int) *CreateBareMetalBuilder {
	b.createBareMetal.APPID = aPPID
	return b
}

func (b *CreateBareMetalBuilder) UserData(userdata string) *CreateBareMetalBuilder {
	b.createBareMetal.userdata = userdata
	return b
}

func (b *CreateBareMetalBuilder) NotifyActivate(notify_activate string) *CreateBareMetalBuilder {
	b.createBareMetal.notify_activate = notify_activate
	return b
}

func (b *CreateBareMetalBuilder) Hostname(hostname string) *CreateBareMetalBuilder {
	b.createBareMetal.hostname = hostname
	return b
}

func (b *CreateBareMetalBuilder) Tag(tag string) *CreateBareMetalBuilder {
	b.createBareMetal.tag = tag
	return b
}

func (b *CreateBareMetalBuilder) Build() (map[string]interface{}, error) {
	if b.createBareMetal.DCID == 0 {
		return nil, errors.New(aliauth.StrMissRequired + "DCID")
	}
	if b.createBareMetal.METALPLANID == 0 {
		return nil, errors.New(aliauth.StrMissRequired + "METALPLANID")
	}
	if b.createBareMetal.OSID == 0 {
		return nil, errors.New(aliauth.StrMissRequired + "OSID")
	}
	params := map[string]interface{}{
		"DCID":            b.createBareMetal.DCID,
		"METALPLANID":     b.createBareMetal.METALPLANID,
		"OSID":            b.createBareMetal.OSID,
		"SCRIPTID":        b.createBareMetal.SCRIPTID,
		"SNAPSHOTID":      b.createBareMetal.SNAPSHOTID,
		"enable_ipv6":     b.createBareMetal.enable_ipv6,
		"label":           b.createBareMetal.label,
		"SSHKEYID":        b.createBareMetal.SSHKEYID,
		"APPID":           b.createBareMetal.APPID,
		"userdata":        b.createBareMetal.userdata,
		"notify_activate": b.createBareMetal.notify_activate,
		"hostname":        b.createBareMetal.hostname,
		"tag":             b.createBareMetal.tag,
	}
	return params, nil
}

// DeleteBareMetal builder pattern code
type DeleteBareMetalBuilder struct {
	deleteBareMetal *DeleteBareMetal
}

func NewDeleteBareMetalBuilder() *DeleteBareMetalBuilder {
	deleteBareMetal := &DeleteBareMetal{}
	b := &DeleteBareMetalBuilder{deleteBareMetal: deleteBareMetal}
	return b
}

func (b *DeleteBareMetalBuilder) SUBID(sUBID int) *DeleteBareMetalBuilder {
	b.deleteBareMetal.SUBID = sUBID
	return b
}

func (b *DeleteBareMetalBuilder) Build() (map[string]interface{}, error) {
	if b.deleteBareMetal.SUBID == 0 {
		return nil, errors.New(aliauth.StrMissRequired + "SUBID")
	}
	params := vultrauth.PutStructIntoMap(b.deleteBareMetal)
	return params, nil
}

// RebootBareMetal builder pattern code
type RebootBareMetalBuilder struct {
	rebootBareMetal *RebootBareMetal
}

func NewRebootBareMetalBuilder() *RebootBareMetalBuilder {
	rebootBareMetal := &RebootBareMetal{}
	b := &RebootBareMetalBuilder{rebootBareMetal: rebootBareMetal}
	return b
}

func (b *RebootBareMetalBuilder) SUBID(sUBID int) *RebootBareMetalBuilder {
	b.rebootBareMetal.SUBID = sUBID
	return b
}

func (b *RebootBareMetalBuilder) Build() (map[string]interface{}, error) {
	if b.rebootBareMetal.SUBID == 0 {
		return nil, errors.New(aliauth.StrMissRequired + "SUBID")
	}
	params := vultrauth.PutStructIntoMap(b.rebootBareMetal)
	return params, nil
}

// ReinstallBareMetal builder pattern code
type ReinstallBareMetalBuilder struct {
	reinstallBareMetal *ReinstallBareMetal
}

func NewReinstallBareMetalBuilder() *ReinstallBareMetalBuilder {
	reinstallBareMetal := &ReinstallBareMetal{}
	b := &ReinstallBareMetalBuilder{reinstallBareMetal: reinstallBareMetal}
	return b
}

func (b *ReinstallBareMetalBuilder) SUBID(sUBID int) *ReinstallBareMetalBuilder {
	b.reinstallBareMetal.SUBID = sUBID
	return b
}

func (b *ReinstallBareMetalBuilder) Build() (map[string]interface{}, error) {
	if b.reinstallBareMetal.SUBID == 0 {
		return nil, errors.New(aliauth.StrMissRequired + "SUBID")
	}
	params := vultrauth.PutStructIntoMap(b.reinstallBareMetal)
	return params, nil
}

// HaltBareMetal builder pattern code
type HaltBareMetalBuilder struct {
	haltBareMetal *HaltBareMetal
}

func NewHaltBareMetalBuilder() *HaltBareMetalBuilder {
	haltBareMetal := &HaltBareMetal{}
	b := &HaltBareMetalBuilder{haltBareMetal: haltBareMetal}
	return b
}

func (b *HaltBareMetalBuilder) SUBID(sUBID int) *HaltBareMetalBuilder {
	b.haltBareMetal.SUBID = sUBID
	return b
}

func (b *HaltBareMetalBuilder) Build() (map[string]interface{}, error) {
	if b.haltBareMetal.SUBID == 0 {
		return nil, errors.New(aliauth.StrMissRequired + "SUBID")
	}
	params := vultrauth.PutStructIntoMap(b.haltBareMetal)
	return params, nil
}

// ListBareMetal builder pattern code
type ListBareMetalBuilder struct {
	listBareMetal *ListBareMetal
}

func NewListBareMetalBuilder() *ListBareMetalBuilder {
	listBareMetal := &ListBareMetal{}
	b := &ListBareMetalBuilder{listBareMetal: listBareMetal}
	return b
}

func (b *ListBareMetalBuilder) SUBID(sUBID int) *ListBareMetalBuilder {
	b.listBareMetal.SUBID = sUBID
	return b
}

func (b *ListBareMetalBuilder) Tag(tag string) *ListBareMetalBuilder {
	b.listBareMetal.tag = tag
	return b
}

func (b *ListBareMetalBuilder) Label(label string) *ListBareMetalBuilder {
	b.listBareMetal.label = label
	return b
}

func (b *ListBareMetalBuilder) MainIp(main_ip string) *ListBareMetalBuilder {
	b.listBareMetal.main_ip = main_ip
	return b
}

func (b *ListBareMetalBuilder) Build() (map[string]interface{}, error) {
	params := make(map[string]interface{})
	if b.listBareMetal.SUBID != 0 {
		params = map[string]interface{}{
			"SUBID": b.listBareMetal.SUBID,
		}
		return params, nil
	}
	if b.listBareMetal.tag != "" {
		params = map[string]interface{}{
			"tag": b.listBareMetal.tag,
		}
		return params, nil
	}
	if b.listBareMetal.label != "" {
		params = map[string]interface{}{
			"label": b.listBareMetal.label,
		}
		return params, nil
	}
	if b.listBareMetal.main_ip != "" {
		params = map[string]interface{}{
			"main_ip": b.listBareMetal.main_ip,
		}
		return params, nil
	}
	return nil, nil
}
