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
