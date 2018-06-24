package googlestorage

func CreateDiskdictnoaryconvert(option Creatdisk, Creatdiskjsonmap map[string]interface{}) {

	if option.Name != "" {
		Creatdiskjsonmap["name"] = option.Name
	}

	if option.Type != "" {
		Creatdiskjsonmap["type"] = option.Type
	}

	if option.Zone != "" {
		Creatdiskjsonmap["zone"] = option.Zone
	}

	if option.SizeGb != "" {
		Creatdiskjsonmap["sizeGb"] = option.SizeGb
	}

	if option.SourceImageEncryptionKeys != (SourceImageEncryptionKey{}) {
		Creatdiskjsonmap["sourceImageEncryptionKey"] = option.SourceImageEncryptionKeys
	}
	if option.DiskEncryptionKeys != (DiskEncryptionKey{}) {
		Creatdiskjsonmap["diskEncryptionKey"] = option.DiskEncryptionKeys
	}
	if option.SourceSnapshotEncryptionKeys != (SourceSnapshotEncryptionKey{}) {
		Creatdiskjsonmap["sourceSnapshotEncryptionKey"] = option.SourceSnapshotEncryptionKeys
	}

	if len(option.Licenses) != 0 {
		Creatdiskjsonmap["licenses"] = option.Licenses
	}

	if len(option.Users) != 0 {
		Creatdiskjsonmap["users"] = option.Users
	}

	if option.LastAttachTimestamp != "" {
		Creatdiskjsonmap["lastAttachTimestamp"] = option.LastAttachTimestamp
	}

	if option.SourceSnapshot != "" {
		Creatdiskjsonmap["SourceSnapshot"] = option.SourceSnapshot
	}

	if option.LastDetachTimestamp != "" {
		Creatdiskjsonmap["lastDetachTimestamp"] = option.LastDetachTimestamp
	}

	if option.Options != "" {
		Creatdiskjsonmap["options"] = option.Options
	}

	if option.SelfLink != "" {
		Creatdiskjsonmap["selfLink"] = option.SelfLink
	}

	if option.SourceImage != "" {
		Creatdiskjsonmap["sourceImage"] = option.SourceImage
	}

	if option.SourceImageID != "" {
		Creatdiskjsonmap["sourceImageID"] = option.SourceImageID
	}

	if option.Status != "" {
		Creatdiskjsonmap["status"] = option.Status
	}

	if option.SourceSnapshotID != "" {
		Creatdiskjsonmap["sourceSnapshotID"] = option.SourceSnapshotID
	}

	if option.LabelFingerprint != "" {
		Creatdiskjsonmap["labelFingerprint"] = option.LabelFingerprint
	}

	if option.Kind != "" {
		Creatdiskjsonmap["kind"] = option.Kind
	}

	if option.ID != "" {
		Creatdiskjsonmap["id"] = option.ID
	}

	if option.Description != "" {
		Creatdiskjsonmap["description"] = option.Description
	}

	if option.CreationTimestamp != "" {
		Creatdiskjsonmap["creationTimestamp"] = option.CreationTimestamp
	}

}

func CreateSnapshotdictnoaryconvert(option Snapshot, CreateSnapshotjsonmap map[string]interface{}) {

	if len(option.Licenses) != 0 {
		CreateSnapshotjsonmap["licenses"] = option.Licenses
	}

	if option.Name != "" {
		CreateSnapshotjsonmap["name"] = option.Name
	}

	if option.CreationTimestamp != "" {
		CreateSnapshotjsonmap["creationTimestamp"] = option.CreationTimestamp
	}

	if option.Description != "" {
		CreateSnapshotjsonmap["description"] = option.Description
	}

	if option.DiskSizeGb != "" {
		CreateSnapshotjsonmap["diskSizeGb"] = option.DiskSizeGb
	}

	if option.ID != "" {
		CreateSnapshotjsonmap["id"] = option.ID
	}

	if option.Kind != "" {
		CreateSnapshotjsonmap["kind"] = option.Kind
	}

	if option.Status != "" {
		CreateSnapshotjsonmap["status"] = option.Status
	}

	if option.SelfLink != "" {
		CreateSnapshotjsonmap["selfLink"] = option.SelfLink
	}

	if option.LabelFingerprint != "" {
		CreateSnapshotjsonmap["labelFingerprint"] = option.LabelFingerprint
	}

	if option.StorageBytes != "" {
		CreateSnapshotjsonmap["storageBytes"] = option.StorageBytes
	}

	if option.SourceDisk != "" {
		CreateSnapshotjsonmap["sourceDisk"] = option.SourceDisk
	}

	if option.SourceDiskID != "" {
		CreateSnapshotjsonmap["sourceDiskID"] = option.SourceDiskID
	}

	if option.StorageBytesStatus != "" {
		CreateSnapshotjsonmap["storageBytesStatus"] = option.StorageBytesStatus
	}

	if option.SourceDiskEncryptionKeys != (SourceDiskEncryptionKey{}) {
		CreateSnapshotjsonmap["sourceDiskEncryptionKey"] = option.SourceDiskEncryptionKeys
	}

	if option.SnapshotEncryptionKeys != (SnapshotEncryptionKey{}) {
		CreateSnapshotjsonmap["snapshotEncryptionKey"] = option.SnapshotEncryptionKeys
	}

}

func AttachDiskdictnoaryconvert(option AttachDisk, AttachDiskjsonmap map[string]interface{}) {

	if len(option.Licenses) != 0 {
		AttachDiskjsonmap["licenses"] = option.Licenses
	}

	if option.DiskEncryptionKeys != (DiskEncryptionKey{}) {
		AttachDiskjsonmap["diskEncryptionKey"] = option.DiskEncryptionKeys
	}

	if option.Source != "" {
		AttachDiskjsonmap["source"] = option.Source
	}

	if option.DeviceName != "" {
		AttachDiskjsonmap["deviceName"] = option.DeviceName
	}

	if option.AutoDelete {
		AttachDiskjsonmap["autoDelete"] = option.AutoDelete
	}

	if option.Boot {
		AttachDiskjsonmap["boot"] = option.Boot
	}

	if option.Index != 0 {
		AttachDiskjsonmap["index"] = option.Index
	}

	if option.Interface != "" {
		AttachDiskjsonmap["interface"] = option.Interface
	}

	if option.Kind != "" {
		AttachDiskjsonmap["kind"] = option.Kind
	}

	if option.Mode != "" {
		AttachDiskjsonmap["mode"] = option.Mode
	}

	if option.Type != "" {
		AttachDiskjsonmap["type"] = option.Type
	}

	if option.Type != "" {
		AttachDiskjsonmap["type"] = option.Type
	}

	if option.InitializeParam != (InitializeParams{}) {

		InitializeParammap := make(map[string]interface{})

		if option.InitializeParam.DiskName != "" {
			InitializeParammap["diskName"] = option.InitializeParam.DiskName
		}

		if option.InitializeParam.DiskType != "" {
			InitializeParammap["diskType"] = option.InitializeParam.DiskType
		}

		if option.InitializeParam.DiskSizeGb != "" {
			InitializeParammap["diskSizeGb"] = option.InitializeParam.DiskSizeGb
		}

		if option.InitializeParam.SourceImage != "" {
			InitializeParammap["sourceImage"] = option.InitializeParam.SourceImage
		}

		if (option.InitializeParam.SourceImageEncryptionKeys != SourceImageEncryptionKey{}) {

			SourceImageEncryptionKeysmmap := make(map[string]interface{})

			if option.InitializeParam.SourceImageEncryptionKeys.RawKey != "" {
				SourceImageEncryptionKeysmmap["rawKey"] = option.InitializeParam.SourceImageEncryptionKeys.RawKey
			}
			if option.InitializeParam.SourceImageEncryptionKeys.RawKey != "" {
				SourceImageEncryptionKeysmmap["sha256"] = option.InitializeParam.SourceImageEncryptionKeys.Sha256
			}

			InitializeParammap["sourceImage"] = SourceImageEncryptionKeysmmap
		}

		AttachDiskjsonmap["initializeParams"] = InitializeParammap
	}

}
