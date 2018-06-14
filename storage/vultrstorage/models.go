package vultrstorage

type VultrStorage struct {
}

type CreateSnapshot struct {
	SUBID int
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
