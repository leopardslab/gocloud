package googlestorage

func Creatediskdictnoaryconvert(option Creatdisk, Creatdiskjsonmap map[string]interface{}) {

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

func Createsnapshotdictnoaryconvert(option Snapshot, Createsnapshotjsonmap map[string]interface{}) {

	if len(option.Licenses) != 0 {
		Createsnapshotjsonmap["licenses"] = option.Licenses
	}

	if option.Name != "" {
		Createsnapshotjsonmap["name"] = option.Name
	}

	if option.CreationTimestamp != "" {
		Createsnapshotjsonmap["creationTimestamp"] = option.CreationTimestamp
	}

	if option.Description != "" {
		Createsnapshotjsonmap["description"] = option.Description
	}

	if option.DiskSizeGb != "" {
		Createsnapshotjsonmap["diskSizeGb"] = option.DiskSizeGb
	}

	if option.ID != "" {
		Createsnapshotjsonmap["id"] = option.ID
	}

	if option.Kind != "" {
		Createsnapshotjsonmap["kind"] = option.Kind
	}

	if option.Status != "" {
		Createsnapshotjsonmap["status"] = option.Status
	}

	if option.SelfLink != "" {
		Createsnapshotjsonmap["selfLink"] = option.SelfLink
	}

	if option.LabelFingerprint != "" {
		Createsnapshotjsonmap["labelFingerprint"] = option.LabelFingerprint
	}

	if option.StorageBytes != "" {
		Createsnapshotjsonmap["storageBytes"] = option.StorageBytes
	}

	if option.SourceDisk != "" {
		Createsnapshotjsonmap["sourceDisk"] = option.SourceDisk
	}

	if option.SourceDiskID != "" {
		Createsnapshotjsonmap["sourceDiskID"] = option.SourceDiskID
	}

	if option.StorageBytesStatus != "" {
		Createsnapshotjsonmap["storageBytesStatus"] = option.StorageBytesStatus
	}

	if option.SourceDiskEncryptionKeys != (SourceDiskEncryptionKey{}) {
		Createsnapshotjsonmap["sourceDiskEncryptionKey"] = option.SourceDiskEncryptionKeys
	}

	if option.SnapshotEncryptionKeys != (SnapshotEncryptionKey{}) {
		Createsnapshotjsonmap["snapshotEncryptionKey"] = option.SnapshotEncryptionKeys
	}

}

func Attachdiskdictnoaryconvert(option Attachdisk, Attachdiskjsonmap map[string]interface{}) {

	if len(option.Licenses) != 0 {
		Attachdiskjsonmap["licenses"] = option.Licenses
	}

	if option.DiskEncryptionKeys != (DiskEncryptionKey{}) {
		Attachdiskjsonmap["diskEncryptionKey"] = option.DiskEncryptionKeys
	}

	if option.Source != "" {
		Attachdiskjsonmap["source"] = option.Source
	}

	if option.DeviceName != "" {
		Attachdiskjsonmap["deviceName"] = option.DeviceName
	}

	if option.AutoDelete {
		Attachdiskjsonmap["autoDelete"] = option.AutoDelete
	}

	if option.Boot {
		Attachdiskjsonmap["boot"] = option.Boot
	}

	if option.Index != 0 {
		Attachdiskjsonmap["index"] = option.Index
	}

	if option.Interface != "" {
		Attachdiskjsonmap["interface"] = option.Interface
	}

	if option.Kind != "" {
		Attachdiskjsonmap["kind"] = option.Kind
	}

	if option.Mode != "" {
		Attachdiskjsonmap["mode"] = option.Mode
	}

	if option.Type != "" {
		Attachdiskjsonmap["type"] = option.Type
	}

	if option.Type != "" {
		Attachdiskjsonmap["type"] = option.Type
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

		Attachdiskjsonmap["initializeParams"] = InitializeParammap
	}

}
