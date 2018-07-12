package vultrstorage

import (
	"errors"
	"github.com/cloudlibz/gocloud/vultrauth"
)

type VultrStorage struct {
}

type CreateSnapshot struct {
	SUBID       int
	description string // (optional) Description of snapshot contents
}

type DeleteSnapshot struct {
	SNAPSHOTID string
}

type CreateDisk struct {
	DCID    int    // DCID of the location to create this subscription in.  See /v1/regions/list
	size_gb int    // Size (in GB) of this subscription.
	label   string // (optional) Text label that will be associated with the subscription
}

type AttachDisk struct {
	SUBID           int // ID of the block storage subscription to attach
	attach_to_SUBID int // ID of the VPS subscription to mount the block storage subscription to
}

type DetachDisk struct {
	SUBID int // ID of the block storage subscription to detach
}

type DeleteDisk struct {
	SUBID int // ID of the block storage subscription to delete
}

// CreateSnapshot builder pattern code
type CreateSnapshotBuilder struct {
	createSnapshot *CreateSnapshot
}

func NewCreateSnapshotBuilder() *CreateSnapshotBuilder {
	createSnapshot := &CreateSnapshot{}
	b := &CreateSnapshotBuilder{createSnapshot: createSnapshot}
	return b
}

func (b *CreateSnapshotBuilder) SUBID(sUBID int) *CreateSnapshotBuilder {
	b.createSnapshot.SUBID = sUBID
	return b
}

func (b *CreateSnapshotBuilder) Description(description string) *CreateSnapshotBuilder {
	b.createSnapshot.description = description
	return b
}

func (b *CreateSnapshotBuilder) Build() (map[string]interface{}, error) {
	if b.createSnapshot.SUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "SUBID")
	}
	params := map[string]interface{}{
		"SUBID": b.createSnapshot.SUBID,
	}
	if b.createSnapshot.description != "" {
		params["description"] = b.createSnapshot.description
	}
	return params, nil
}

// DeleteSnapshot builder pattern code
type DeleteSnapshotBuilder struct {
	deleteSnapshot *DeleteSnapshot
}

func NewDeleteSnapshotBuilder() *DeleteSnapshotBuilder {
	deleteSnapshot := &DeleteSnapshot{}
	b := &DeleteSnapshotBuilder{deleteSnapshot: deleteSnapshot}
	return b
}

func (b *DeleteSnapshotBuilder) SNAPSHOTID(sNAPSHOTID string) *DeleteSnapshotBuilder {
	b.deleteSnapshot.SNAPSHOTID = sNAPSHOTID
	return b
}

func (b *DeleteSnapshotBuilder) Build() (map[string]interface{}, error) {
	if b.deleteSnapshot.SNAPSHOTID == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "SNAPSHOTID")
	}
	params := map[string]interface{}{
		"SNAPSHOTID": b.deleteSnapshot.SNAPSHOTID,
	}
	return params, nil
}

// CreateDisk builder pattern code
type CreateDiskBuilder struct {
	createDisk *CreateDisk
}

func NewCreateDiskBuilder() *CreateDiskBuilder {
	createDisk := &CreateDisk{}
	b := &CreateDiskBuilder{createDisk: createDisk}
	return b
}

func (b *CreateDiskBuilder) DCID(dCID int) *CreateDiskBuilder {
	b.createDisk.DCID = dCID
	return b
}

func (b *CreateDiskBuilder) SizeGb(size_gb int) *CreateDiskBuilder {
	b.createDisk.size_gb = size_gb
	return b
}

func (b *CreateDiskBuilder) Label(label string) *CreateDiskBuilder {
	b.createDisk.label = label
	return b
}

func (b *CreateDiskBuilder) Build() (map[string]interface{}, error) {
	if b.createDisk.DCID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "DCID")
	}
	if b.createDisk.size_gb == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "size_gb")
	}
	params := make(map[string]interface{})
	params["DCID"] = b.createDisk.DCID
	params["size_gb"] = b.createDisk.size_gb
	if b.createDisk.label != "" {
		params["label"] = b.createDisk.label
	}
	return params, nil
}

// AttachDisk builder pattern code
type AttachDiskBuilder struct {
	attachDisk *AttachDisk
}

func NewAttachDiskBuilder() *AttachDiskBuilder {
	attachDisk := &AttachDisk{}
	b := &AttachDiskBuilder{attachDisk: attachDisk}
	return b
}

func (b *AttachDiskBuilder) SUBID(sUBID int) *AttachDiskBuilder {
	b.attachDisk.SUBID = sUBID
	return b
}

func (b *AttachDiskBuilder) AttachToSUBID(attach_to_SUBID int) *AttachDiskBuilder {
	b.attachDisk.attach_to_SUBID = attach_to_SUBID
	return b
}

func (b *AttachDiskBuilder) Build() (map[string]interface{}, error) {
	if b.attachDisk.SUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "SUBID")
	}
	if b.attachDisk.attach_to_SUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "attach_to_SUBID")
	}
	params := make(map[string]interface{})
	params["SUBID"] = b.attachDisk.SUBID
	params["attach_to_SUBID"] = b.attachDisk.attach_to_SUBID
	return params, nil
}

// DetachDisk builder pattern code
type DetachDiskBuilder struct {
	detachDisk *DetachDisk
}

func NewDetachDiskBuilder() *DetachDiskBuilder {
	detachDisk := &DetachDisk{}
	b := &DetachDiskBuilder{detachDisk: detachDisk}
	return b
}

func (b *DetachDiskBuilder) SUBID(sUBID int) *DetachDiskBuilder {
	b.detachDisk.SUBID = sUBID
	return b
}

func (b *DetachDiskBuilder) Build() (map[string]interface{}, error) {
	if b.detachDisk.SUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "SUBID")
	}
	params := make(map[string]interface{})
	params["SUBID"] = b.detachDisk.SUBID
	return params, nil
}

// DeleteDisk builder pattern code
type DeleteDiskBuilder struct {
	deleteDisk *DeleteDisk
}

func NewDeleteDiskBuilder() *DeleteDiskBuilder {
	deleteDisk := &DeleteDisk{}
	b := &DeleteDiskBuilder{deleteDisk: deleteDisk}
	return b
}

func (b *DeleteDiskBuilder) SUBID(sUBID int) *DeleteDiskBuilder {
	b.deleteDisk.SUBID = sUBID
	return b
}

func (b *DeleteDiskBuilder) Build() (map[string]interface{}, error) {
	if b.deleteDisk.SUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "SUBID")
	}
	params := make(map[string]interface{})
	params["SUBID"] = b.deleteDisk.SUBID
	return params, nil
}
