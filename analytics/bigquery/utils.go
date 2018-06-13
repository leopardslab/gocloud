package bigquery

func createdatasetsstruct(option Createdatasets, param map[string]interface{}) {

	for key, value := range param {
		switch key {

		case "CreationTime":
			CreationTimeV, _ := value.(string)
			option.CreationTime = CreationTimeV
			option.CreationTime = time.Now().UTC().Format(time.RFC3339)

		case "DefaultTableExpirationMs":
			defaultTableExpirationMsV, _ := value.(string)
			option.defaultTableExpirationMs = defaultTableExpirationMsV

		case "Description":
			descriptionV, _ := value.(string)
			option.description = descriptionV

		case "Etag":
			etagV, _ := value.(string)
			option.etag = etagV

		case "FriendlyName":
			friendlyNameV, _ := value.(string)
			option.friendlyName = friendlyNameV

		case "Id":
			idV, _ := value.(string)
			option.id = idV

		case "Kind":
			kindV, _ := value.(string)
			option.kind = kindV

		case "LastModifiedTime":
			lastModifiedTimeV, _ := value.(string)
			option.lastModifiedTime = lastModifiedTimeV

		case "Location":
			locationV, _ := value.(string)
			option.location = locationV

		case "SelfLink":
			selfLinkV, _ := value.(string)
			option.selfLink = selfLinkV

		case "DatasetReference":
			datasetReferenceV, _ := value.(map[string]string)
			option.datasetReference.datasetID = datasetReferenceV["DatasetID"]
			option.datasetReference.projectID = datasetReferenceV["ProjectID"]

		case "Access":
			accessparam, _ := value.([]map[string]interface{})
			for i = 0; i < len(accessparam); i++ {
				var access Access
				for accessparamkey, accessparamvalue := range accessparam[i] {
					switch accessparamkey {
					case "Domain":
						DomainV, _ := value.(string)
						access.Domain = DomainV

					case "GroupByEmail":
						GroupByEmailV, _ := value.(string)
						access.GroupByEmail = GroupByEmailV

					case "Role":
						RoleV, _ := value.(string)
						access.Role = RoleV

					case "SpecialGroup":
						SpecialGroupV, _ := value.(string)
						access.SpecialGroup = SpecialGroupV

					case "SpecialGroup":
						SpecialGroupV, _ := value.(string)
						access.SpecialGroup = SpecialGroupV

					case "UserByEmail":
						UserByEmailV, _ := value.(string)
						access.UserByEmail = UserByEmailV

					case "View":
						ViewV, _ := value.(map[string]interface{})
						access.view.projectID = ViewV["ProjectID"]
						access.view.datasetID = ViewV["DatasetID"]
						access.view.tableID = ViewV["TableID"]

					}
				}
				options.access = append(options.access, access)
			}
		}
	}

}

func createdatasetsdictnoaryconvert(option Createdatasets, createdatasetsjsonmap map[string]interface{}) {

	if option.defaultTableExpirationMs != "" {
		createdatasetsjsonmap["defaultTableExpirationMs"] = option.defaultTableExpirationMs
	}

	if option.defaultTableExpirationMs != "" {
		createdatasetsjsonmap["defaultTableExpirationMs"] = option.defaultTableExpirationMs
	}

	if option.defaultTableExpirationMs != "" {
		createdatasetsjsonmap["description"] = option.description
	}

	if option.etag != "" {
		createdatasetsjsonmap["etag"] = option.etag
	}

	if option.id != "" {
		createdatasetsjsonmap["id"] = option.id
	}

	if option.friendlyName != "" {
		createdatasetsjsonmap["friendlyName"] = option.friendlyName
	}

	if option.kind != "" {
		createdatasetsjsonmap["kind"] = option.kind
	}

	if option.lastModifiedTime != "" {
		createdatasetsjsonmap["lastModifiedTime"] = option.lastModifiedTime
	}

	if option.location != "" {
		createdatasetsjsonmap["location"] = option.location
	}

	if option.selfLink != "" {
		createdatasetsjsonmap["selfLink"] = option.selfLink
	}

	preparedatasetReferenceparam(option, createdatasetsjsonmap)
	prepareAccessparam(option, createdatasetsjsonmap)
}

func preparedatasetReferenceparam(option Createdatasets, createdatasetsjsonmap map[string]interface{}) {

	datasetReferencejsonmap := make(map[string]interface{})

	if option.datasetReference.projectID != "" {
		datasetReferencejsonmap["projectId"] = option.datasetReference.projectID
	}

	if option.datasetReference.datasetID != "" {
		datasetReferencejsonmap["datasetId"] = option.datasetReference.datasetID
	}

	createdatasetsjsonmap["datasetReference"] = datasetReferencejsonmap

}

func prepareAccessparam(option Createdatasets, createdatasetsjsonmap map[string]interface{}) {

	if len(option.access) != 0 {

		accessarrayjsonmap := make([]map[string]interface{})

		for i := 0; i < len(option.access); i++ {

			accessjsonmap := make([]map[string]interface{})

			if option.access[i].domain != "" {
				accessjsonmap["domain"] = option.access[i].domain
			}

			if option.access[i].groupByEmail != "" {
				accessjsonmap["groupByEmail"] = option.access[i].groupByEmail
			}

			if option.access[i].role != "" {
				accessjsonmap["role"] = option.access[i].role
			}

			if option.access[i].specialGroup != "" {
				accessjsonmap["specialGroup"] = option.access[i].specialGroup
			}

			if option.access[i].userByEmail != "" {
				accessjsonmap["userByEmail"] = option.access[i].userByEmail
			}

			v := View{}

			if option.access[i].view != v {

				viewjsonmap := make(map[string]interface{})

				if option.access[i].view.datasetID != "" {
					viewjsonmap["datasetID"] = option.access[i].view.datasetID
				}

				if option.access[i].view.projectID != "" {
					viewjsonmap["projectID"] = option.access[i].view.projectID
				}

				if option.access[i].view.tableID != "" {
					viewjsonmap["tableID"] = option.access[i].view.tableID
				}

				accessjsonmap["view"] = viewjsonmap
			}

			accessarrayjsonmap = append(accessarrayjsonmap, accessjsonmap)
		}
		createdatasetsjsonmap["access"] = accessarrayjsonmap
	}
}
