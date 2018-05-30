package googlecloudfunctions

//CreateGooglecloudfunctionedictnoaryconvert convert dict to form valid json.
func CreateGooglecloudfunctionedictnoaryconvert(option CreateGooglecloudfunction, CreateGooglecloudfunctionjsonmap map[string]interface{}) {

	if option.AvailableMemoryMb != 0 {
		CreateGooglecloudfunctionjsonmap["availableMemoryMb"] = option.AvailableMemoryMb
	}

	if option.Name != "" {
		CreateGooglecloudfunctionjsonmap["name"] = option.Name
	}

	if option.Status != "" {
		CreateGooglecloudfunctionjsonmap["status"] = option.Status
	}

	if option.EntryPoint != "" {
		CreateGooglecloudfunctionjsonmap["entryPoint"] = option.EntryPoint
	}

	if option.Timeout != "" {
		CreateGooglecloudfunctionjsonmap["timeout"] = option.Timeout
	}

	if option.ServiceAccountEmail != "" {
		CreateGooglecloudfunctionjsonmap["serviceAccountEmail"] = option.ServiceAccountEmail
	}

	if option.UpdateTime != "" {
		CreateGooglecloudfunctionjsonmap["updateTime"] = option.UpdateTime
	}

	if option.VersionID != "" {
		CreateGooglecloudfunctionjsonmap["versionId"] = option.VersionID
	}

	if option.ServiceAccountEmail != "" {
		CreateGooglecloudfunctionjsonmap["sourceUploadUrl"] = option.SourceUploadURL
	}

	if option.Labels.DeploymentTool != "" {
		labels := make(map[string]string)
		labels["deployment-tool"] = option.Labels.DeploymentTool
		CreateGooglecloudfunctionjsonmap["labels"] = labels
	}

	if option.HTTPSTrigger.URL != "" {
		labels := make(map[string]string)
		labels["url"] = option.HTTPSTrigger.URL
		CreateGooglecloudfunctionjsonmap["httpsTrigger"] = labels
	}

}
