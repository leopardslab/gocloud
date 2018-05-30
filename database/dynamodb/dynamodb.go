package dynamodb

//List tables.
func (dynamodb *Dynamodb) Listtables(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var ExclusiveStartTableName, Limit, Region string

	for key, value := range param {
		switch key {
		case "ExclusiveStartTableName":
			ExclusiveStartTableNameV, _ := value.(string)
			ExclusiveStartTableName = ExclusiveStartTableNameV

		case "Limit":
			LimitV, _ := value.(string)
			Limit = LimitV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)

	preparelisttables(params, ExclusiveStartTableName, Limit, Region)

	listtablesjsonmap := map[string]interface{}{}

	listtablesjsonmap["ExclusiveStartTableName"] = ExclusiveStartTableName
	listtablesjsonmap["Limit"] = Limit

	response := make(map[string]interface{})

	err = dynamodb.PrepareSignatureV4query(params, listtablesjsonmap, response)
	resp = response
	return resp, err
}

//Delete tables.
func (dynamodb *Dynamodb) Deletetables(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})

	var TableName, Region string

	for key, value := range param {
		switch key {
		case "TableName":
			TableNameV, _ := value.(string)
			TableName = TableNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)

	preparedeletetables(params, TableName, Region)

	deletetablesjsonmap := map[string]interface{}{
		"TableName": TableName,
	}

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, deletetablesjsonmap, response)
	resp = response
	return resp, err
}

//Create tables.
func (dynamodb *Dynamodb) Createtables(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var createtable Createtable
	var Region string

	for key, value := range param {
		switch key {

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "TableName":
			TableNameV, _ := value.(string)
			createtable.TableName = TableNameV

		case "StreamSpecification":
			streamSpecificationparam, _ := value.(map[string]interface{})
			for streamSpecificationparamkey, streamSpecificationparamvalue := range streamSpecificationparam {
				switch streamSpecificationparamkey {
				case "StreamViewType":
					createtable.streamSpecification.StreamViewType = streamSpecificationparamvalue.(string)

				case "StreamEnabled":
					createtable.streamSpecification.StreamEnabled = streamSpecificationparamvalue.(bool)
				}
			}

		case "SSESpecification":
			sSESpecificationparam, _ := value.(map[string]interface{})
			for sSESpecificationparamkey, sSESpecificationparamparamvalue := range sSESpecificationparam {
				switch sSESpecificationparamkey {
				case "Enabled":
					createtable.sSESpecification.Enabled = sSESpecificationparamparamvalue.(bool)
				}
			}

		case "ProvisionedThroughput":
			provisionedThroughputparam, _ := value.(map[string]interface{})

			for provisionedThroughputparamkey, provisionedThroughputparamvalue := range provisionedThroughputparam {
				switch provisionedThroughputparamkey {
				case "ReadCapacityUnits":
					ReadCapacityUnitsv, _ := provisionedThroughputparamvalue.(int)
					createtable.provisionedThroughput.ReadCapacityUnits = ReadCapacityUnitsv

				case "WriteCapacityUnits":
					createtable.provisionedThroughput.WriteCapacityUnits = provisionedThroughputparamvalue.(int)

				}
			}

		case "KeySchema":
			keySchemaparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(keySchemaparam); i++ {
				var keySchema KeySchema
				for keySchemaparamkey, keySchemaparamvalue := range keySchemaparam[i] {
					switch keySchemaparamkey {
					case "AttributeName":
						keySchema.AttributeName = keySchemaparamvalue.(string)

					case "KeyType":
						keySchema.KeyType = keySchemaparamvalue.(string)
					}
				}
				createtable.keySchema = append(createtable.keySchema, keySchema)
			}
		case "LocalSecondaryIndexes":
			localSecondaryIndexesparam, _ := value.([]map[string]interface{})

			for i := 0; i < len(localSecondaryIndexesparam); i++ {
				var localSecondaryIndexes LocalSecondaryIndexes
				for localSecondaryIndexesparamkey, localSecondaryIndexesparamvalue := range localSecondaryIndexesparam[i] {
					switch localSecondaryIndexesparamkey {
					case "IndexName":
						localSecondaryIndexes.IndexName = localSecondaryIndexesparamvalue.(string)

					case "KeySchema":
						keySchemaparam, _ := localSecondaryIndexesparamvalue.([]map[string]interface{})
						for i := 0; i < len(keySchemaparam); i++ {
							var keySchema KeySchema
							for keySchemaparamkey, keySchemaparamvalue := range keySchemaparam[i] {
								switch keySchemaparamkey {
								case "AttributeName":
									keySchema.AttributeName = keySchemaparamvalue.(string)

								case "KeyType":
									keySchema.KeyType = keySchemaparamvalue.(string)
								}
							}

							localSecondaryIndexes.keySchema = append(localSecondaryIndexes.keySchema, keySchema)
						}

					case "Projection":
						projectionparam, _ := localSecondaryIndexesparamvalue.(map[string]interface{})
						for projectionparamkey, projectionparamvalue := range projectionparam {
							switch projectionparamkey {
							case "NonKeyAttributes":
								localSecondaryIndexes.projection.NonKeyAttributes = projectionparamvalue.([]string)

							case "ProjectionType":
								localSecondaryIndexes.projection.ProjectionType = projectionparamvalue.(string)
							}
						}
					}
				}
				createtable.localSecondaryIndexes = append(createtable.localSecondaryIndexes, localSecondaryIndexes)
			}

		case "globalSecondaryIndexes":
			globalSecondaryIndexesparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(globalSecondaryIndexesparam); i++ {
				var globalSecondaryIndexes GlobalSecondaryIndexes
				for globalSecondaryIndexesparamkey, globalSecondaryIndexesparamvalue := range globalSecondaryIndexesparam[i] {
					switch globalSecondaryIndexesparamkey {
					case "IndexName":
						globalSecondaryIndexes.IndexName = globalSecondaryIndexesparamvalue.(string)

					case "keySchema":
						keySchemaparam, _ := globalSecondaryIndexesparamvalue.([]map[string]interface{})
						for i := 0; i < len(keySchemaparam); i++ {
							var keySchema KeySchema
							for keySchemaparamkey, keySchemaparamvalue := range keySchemaparam[i] {
								switch keySchemaparamkey {
								case "AttributeName":
									keySchema.AttributeName = keySchemaparamvalue.(string)

								case "KeyType":
									keySchema.KeyType = keySchemaparamvalue.(string)
								}
							}

							globalSecondaryIndexes.keySchema = append(globalSecondaryIndexes.keySchema, keySchema)
						}

					case "Projection":
						projectionparam, _ := globalSecondaryIndexesparamvalue.(map[string]interface{})
						for projectionparamkey, projectionparamvalue := range projectionparam {
							switch projectionparamkey {
							case "NonKeyAttributes":
								globalSecondaryIndexes.projection.NonKeyAttributes = projectionparamvalue.([]string)

							case "ProjectionType":
								globalSecondaryIndexes.projection.ProjectionType = projectionparamvalue.(string)
							}
						}

					case "ProvisionedThroughput":
						provisionedThroughputparam, _ := globalSecondaryIndexesparamvalue.(map[string]interface{})
						for provisionedThroughputparamkey, provisionedThroughputparamvalue := range provisionedThroughputparam {
							switch provisionedThroughputparamkey {

							case "ReadCapacityUnits":
								createtable.provisionedThroughput.ReadCapacityUnits = provisionedThroughputparamvalue.(int)

							case "WriteCapacityUnits":
								createtable.provisionedThroughput.WriteCapacityUnits = provisionedThroughputparamvalue.(int)
							}
						}

					}
				}
				createtable.globalSecondaryIndexes = append(createtable.globalSecondaryIndexes, globalSecondaryIndexes)
			}

		case "AttributeDefinitions":
			attributeDefinitionsparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(attributeDefinitionsparam); i++ {
				var attributeDefinitions AttributeDefinitions
				for attributeDefinitionsparamkey, attributeDefinitionsparamvalue := range attributeDefinitionsparam[i] {
					switch attributeDefinitionsparamkey {
					case "AttributeName":
						attributeDefinitions.AttributeName = attributeDefinitionsparamvalue.(string)

					case "AttributeType":
						attributeDefinitions.AttributeType = attributeDefinitionsparamvalue.(string)
					}
				}
				createtable.attributeDefinitions = append(createtable.attributeDefinitions, attributeDefinitions)
			}
		}
	}
	params := make(map[string]string)

	preparecreatetable(params, Region)

	createtablejsonmap := map[string]interface{}{}

	preparecreatetablejsonmap(createtablejsonmap, createtable)

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, createtablejsonmap, response)
	resp = response
	return resp, err
}

//Describe tables.
func (dynamodb *Dynamodb) Describetables(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})

	var TableName, Region string

	for key, value := range param {
		switch key {
		case "TableName":
			TableNameV, _ := value.(string)
			TableName = TableNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}

	params := make(map[string]string)

	preparedescribetables(params, TableName, Region)

	deletetablesjsonmap := map[string]interface{}{
		"TableName": TableName,
	}

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, deletetablesjsonmap, response)
	resp = response
	return resp, err
}
