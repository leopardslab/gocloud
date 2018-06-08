package dynamodb

func preparedescribetables(params map[string]string, TableName string, Region string) {

	if TableName != "" {
		params["TableName"] = TableName
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "DynamoDB_20120810.DescribeTable"
}

func preparedeletetables(params map[string]string, TableName string, Region string) {

	if TableName != "" {
		params["TableName"] = TableName
	}

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "DynamoDB_20120810.DeleteTable"
}

func preparelisttables(params map[string]string, Region string) {

	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "DynamoDB_20120810.ListTables"
}

func preparelisttablesparamsdict(listtablesjsonmap map[string]interface{}, ExclusiveStartTableName string, Limit int) {

	if ExclusiveStartTableName != "" {
		listtablesjsonmap["ExclusiveStartTableName"] = ExclusiveStartTableName
	}

	if Limit != 0 {
		listtablesjsonmap["Limit"] = Limit
	}
}

func preparecreatetableStreamSpecificationparams(createtablejsonmap map[string]interface{}, createtable Createtable) {

	if (createtable.streamSpecification != StreamSpecification{}) {

		streamSpecificationv := make(map[string]interface{})

		if createtable.streamSpecification.StreamViewType != "" {
			streamSpecificationv["StreamViewType"] = createtable.streamSpecification.StreamViewType
		}

		if createtable.streamSpecification.StreamEnabled {
			streamSpecificationv["StreamEnabled"] = createtable.streamSpecification.StreamEnabled
		}

		createtablejsonmap["StreamSpecification"] = streamSpecificationv
	}
}

func preparecreatetableSSESpecificationparams(createtablejsonmap map[string]interface{}, createtable Createtable) {
	if (createtable.sSESpecification != SSESpecification{}) {
		sSESpecificationv := make(map[string]interface{})
		sSESpecificationv["Enabled"] = createtable.sSESpecification.Enabled
		createtablejsonmap["SSESpecification"] = sSESpecificationv
	}
}

func preparecreatetableProvisionedThroughputparams(createtablejsonmap map[string]interface{}, createtable Createtable) {
	if (createtable.provisionedThroughput != ProvisionedThroughput{}) {
		provisionedThroughputv := make(map[string]interface{})
		if createtable.provisionedThroughput.ReadCapacityUnits != 0 {
			provisionedThroughputv["ReadCapacityUnits"] = createtable.provisionedThroughput.ReadCapacityUnits
		}
		if createtable.provisionedThroughput.WriteCapacityUnits != 0 {
			provisionedThroughputv["WriteCapacityUnits"] = createtable.provisionedThroughput.WriteCapacityUnits
		}
		createtablejsonmap["ProvisionedThroughput"] = provisionedThroughputv
	}
}

func prepareAttributeDefinitionsparams(createtablejsonmap map[string]interface{}, createtable Createtable) {
	if len(createtable.attributeDefinitions) != 0 {
		attributeDefinitionvs := make([]map[string]interface{}, 0)

		for i := 0; i < len(createtable.keySchema); i++ {
			attributeDefinitionv := make(map[string]interface{})

			if createtable.attributeDefinitions[i].AttributeName != "" {
				attributeDefinitionv["AttributeName"] = createtable.attributeDefinitions[i].AttributeName
			}

			if createtable.attributeDefinitions[i].AttributeType != "" {
				attributeDefinitionv["AttributeType"] = createtable.attributeDefinitions[i].AttributeType
			}

			attributeDefinitionvs = append(attributeDefinitionvs, attributeDefinitionv)
		}

		createtablejsonmap["AttributeDefinitions"] = attributeDefinitionvs
	}
}

func preparekeySchemaparams(createtablejsonmap map[string]interface{}, createtable Createtable) {
	if len(createtable.keySchema) != 0 {
		keySchemavs := make([]map[string]interface{}, 0)

		for i := 0; i < len(createtable.keySchema); i++ {
			keySchemav := make(map[string]interface{})

			if createtable.keySchema[i].AttributeName != "" {
				keySchemav["AttributeName"] = createtable.keySchema[i].AttributeName
			}

			if createtable.keySchema[i].KeyType != "" {
				keySchemav["KeyType"] = createtable.keySchema[i].KeyType
			}

			keySchemavs = append(keySchemavs, keySchemav)
		}

		createtablejsonmap["KeySchema"] = keySchemavs
	}
}

func prepareGlobalSecondaryIndexesparams(createtablejsonmap map[string]interface{}, createtable Createtable) {

	if len(createtable.globalSecondaryIndexes) != 0 {

		globalSecondaryIndexesvarrayjsonmap := make([]map[string]interface{}, 0)

		for i := 0; i < len(createtable.globalSecondaryIndexes); i++ {

			globalSecondaryIndexessvjsonmap := make(map[string]interface{})

			if createtable.globalSecondaryIndexes[i].IndexName != "" {
				globalSecondaryIndexessvjsonmap["IndexName"] = createtable.globalSecondaryIndexes[i].IndexName
			}

			p := Projection{}

			if createtable.globalSecondaryIndexes[i].projection.ProjectionType == p.ProjectionType && len(createtable.globalSecondaryIndexes[i].projection.NonKeyAttributes) == len(p.NonKeyAttributes) {

				projectionv := make(map[string]interface{})
				projectionv["ProjectionType"] = createtable.globalSecondaryIndexes[i].projection.ProjectionType
				projectionv["NonKeyAttributes"] = createtable.globalSecondaryIndexes[i].projection.NonKeyAttributes
				globalSecondaryIndexessvjsonmap["Projection"] = projectionv
			}

			if (createtable.globalSecondaryIndexes[i].provisionedThroughput != ProvisionedThroughput{}) {

				provisionedThroughputv := make(map[string]interface{})
				provisionedThroughputv["ReadCapacityUnits"] = createtable.globalSecondaryIndexes[i].provisionedThroughput.ReadCapacityUnits
				provisionedThroughputv["WriteCapacityUnits"] = createtable.globalSecondaryIndexes[i].provisionedThroughput.WriteCapacityUnits
				globalSecondaryIndexessvjsonmap["ProvisionedThroughput"] = provisionedThroughputv
			}

			if len(createtable.globalSecondaryIndexes[i].keySchema) != 0 {

				keySchemavs := make([]map[string]interface{}, 0)

				for i := 0; i < len(createtable.globalSecondaryIndexes[i].keySchema); i++ {

					keySchemav := make(map[string]interface{})

					if createtable.globalSecondaryIndexes[i].keySchema[i].AttributeName != "" {
						keySchemav["AttributeName"] = createtable.globalSecondaryIndexes[i].keySchema[i].AttributeName
					}

					if createtable.localSecondaryIndexes[i].keySchema[i].KeyType != "" {
						keySchemav["KeyType"] = createtable.globalSecondaryIndexes[i].keySchema[i].KeyType
					}

					keySchemavs = append(keySchemavs, keySchemav)
				}

				globalSecondaryIndexessvjsonmap["KeySchema"] = keySchemavs
			}

			globalSecondaryIndexesvarrayjsonmap = append(globalSecondaryIndexesvarrayjsonmap, globalSecondaryIndexessvjsonmap)
		}

		createtablejsonmap["GlobalSecondaryIndexes"] = globalSecondaryIndexesvarrayjsonmap
	}

}

func prepareLocalSecondaryIndexesparams(createtablejsonmap map[string]interface{}, createtable Createtable) {

	if len(createtable.localSecondaryIndexes) != 0 {

		localSecondaryIndexesvarrayjsonmap := make([]map[string]interface{}, 0)

		for i := 0; i < len(createtable.localSecondaryIndexes); i++ {

			localSecondaryIndexesvjsonmap := make(map[string]interface{})

			if createtable.localSecondaryIndexes[i].IndexName != "" {
				localSecondaryIndexesvjsonmap["IndexName"] = createtable.localSecondaryIndexes[i].IndexName
			}

			p := Projection{}

			if createtable.localSecondaryIndexes[i].projection.ProjectionType == p.ProjectionType && len(createtable.localSecondaryIndexes[i].projection.NonKeyAttributes) == len(p.NonKeyAttributes) {
				projectionv := make(map[string]interface{})
				projectionv["ProjectionType"] = createtable.localSecondaryIndexes[i].projection.ProjectionType
				projectionv["NonKeyAttributes"] = createtable.localSecondaryIndexes[i].projection.NonKeyAttributes
				localSecondaryIndexesvjsonmap["Projection"] = projectionv
			}

			if len(createtable.localSecondaryIndexes[i].keySchema) != 0 {

				lenv := len(createtable.localSecondaryIndexes[i].keySchema)

				keySchemavs := make([]map[string]interface{}, 0)

				for j := 0; j < lenv; i++ {

					keySchemav := make(map[string]interface{})
					keySchemav["AttributeName"] = createtable.localSecondaryIndexes[i].keySchema[j].AttributeName

					if createtable.localSecondaryIndexes[i].keySchema[j].KeyType != "" {
						keySchemav["KeyType"] = createtable.localSecondaryIndexes[i].keySchema[j].KeyType
					}

					keySchemavs = append(keySchemavs, keySchemav)
				}

				localSecondaryIndexesvjsonmap["KeySchema"] = keySchemavs
			}

			localSecondaryIndexesvarrayjsonmap = append(localSecondaryIndexesvarrayjsonmap, localSecondaryIndexesvjsonmap)
		}
		createtablejsonmap["LocalSecondaryIndexes"] = localSecondaryIndexesvarrayjsonmap
	}

}

func preparecreatetable(params map[string]string, Region string) {

	if Region != "" {
		params["Region"] = Region
	}

	params["amztarget"] = "DynamoDB_20120810.CreateTable"
}

func preparecreatetablejsonmap(createtablejsonmap map[string]interface{}, createtable Createtable) {

	if createtable.TableName != "" {
		createtablejsonmap["TableName"] = createtable.TableName
	}

	preparecreatetableStreamSpecificationparams(createtablejsonmap, createtable)
	preparecreatetableSSESpecificationparams(createtablejsonmap, createtable)
	preparecreatetableProvisionedThroughputparams(createtablejsonmap, createtable)
	preparekeySchemaparams(createtablejsonmap, createtable)
	prepareAttributeDefinitionsparams(createtablejsonmap, createtable)
	prepareGlobalSecondaryIndexesparams(createtablejsonmap, createtable)
}
