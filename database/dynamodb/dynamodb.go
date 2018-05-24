package dynamodb

import (
	"bytes"
	"encoding/json"
	"fmt"
	awsauth "github.com/cloudlibz/gocloud/awsauth"
	"io/ioutil"
	"net/http"
)

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
				case "StreamViewType":
					createtable.provisionedThroughput.ReadCapacityUnits = provisionedThroughputparamvalue.(int)

				case "StreamEnabled":
					createtable.provisionedThroughput.WriteCapacityUnits = provisionedThroughputparamvalue.(int)
				}
			}

		case "keySchema":
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

					case "keySchema":
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
							case "StreamViewType":
								createtable.provisionedThroughput.ReadCapacityUnits = provisionedThroughputparamvalue.(int)

							case "StreamEnabled":
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

	response := make(map[string]interface{})
	err = dynamodb.PrepareSignatureV4query(params, createtablejsonmap, response)
	resp = response
	return resp, err
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

	preparecreatetableStreamSpecificationparams(createtablejsonmap,createtable)
	preparecreatetableSSESpecificationparams(createtablejsonmap,createtable)
	preparecreatetableProvisionedThroughputparams(createtablejsonmap,createtable)
	preparekeySchemaparams(createtablejsonmap,createtable)
	prepareAttributeDefinitionsparams(createtablejsonmap,createtable)
	prepareAttributeDefinitionsparams(createtablejsonmap,createtable)
	prepareGlobalSecondaryIndexesparams(createtablejsonmap,createtable)
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
			provisionedThroughputv["ReadCapacityUnits"] = createtable.provisionedThroughput.WriteCapacityUnits
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

				globalSecondaryIndexesvarrayjsonmap := make([]map[string]interface{},0)

				for(i:=0;i< len(createtable.globalSecondaryIndexes); i++){

					globalSecondaryIndexessvjsonmap := make(map[string]interface{})

					if(globalSecondaryIndexesvjsonmap["IndexName"] != ""){
							globalSecondaryIndexesvjsonmap["IndexName"]	= globalSecondaryIndexes[i].IndexName
					}

					if( createtable.globalSecondaryIndexes.projection!=Projection{}){

						projectionv := make(map[string]interface{})
						projectionv["ProjectionType"] = createtable.globalSecondaryIndexes[i].projection.ProjectionType
						projectionv["NonKeyAttributes"] =  createtable.globalSecondaryIndexes[i].projection.NonKeyAttributes
						globalSecondaryIndexesvjsonmap["Projection"] = projectionv
					}

					if( createtable.globalSecondaryIndexes.provisionedThroughput!=ProvisionedThroughput{}){

						provisionedThroughputv := make(map[string]interface{})
						projectionv["ReadCapacityUnits"] = createtable.globalSecondaryIndexes[i].provisionedThroughput.ReadCapacityUnits
						projectionv["WriteCapacityUnits"] =  createtable.globalSecondaryIndexes[i].provisionedThroughput.WriteCapacityUnits
						globalSecondaryIndexesvjsonmap["ProvisionedThroughput"] = provisionedThroughputv
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

						globalSecondaryIndexesvjsonmap["KeySchema"] = keySchemavs
					}

				 globalSecondaryIndexesvarrayjsonmap = append(globalSecondaryIndexesvarrayjsonmap,globalSecondaryIndexesvjsonmap)
				}

				createtablejsonmap["GlobalSecondaryIndexes"] = globalSecondaryIndexesvarrayjsonmap
		}
}



func prepareLocalSecondaryIndexesparams(createtablejsonmap map[string]interface{}, createtable Createtable) {

		if len(createtable.localSecondaryIndexes) != 0 {

				localSecondaryIndexesvarrayjsonmap := make([]map[string]interface{},0)

				for(i:=0;i< len(localSecondaryIndexes); i++){

					localSecondaryIndexesvjsonmap := make(map[string]interface{})

					if createtable.localSecondaryIndexes[i].IndexName != "" {
						localSecondaryIndexesvjsonmap["IndexName"]	=createtable.localSecondaryIndexes[i].IndexName
					}
					if(createtable.projection!=Projection{}){

						projectionv := make(map[string]interface{})
						projectionv["ProjectionType"] = createtable.localSecondaryIndexes[i].projection.ProjectionType
						projectionv["NonKeyAttributes"] =  createtable.localSecondaryIndexes[i].projection.NonKeyAttributes
						localSecondaryIndexesvjsonmap["Projection"] = projectionv
					}

					if len(createtable.localSecondaryIndexes[i].keySchema) != 0 {

						keySchemavs := make([]map[string]interface{}, 0)

						for i := 0; i < len(createtable.localSecondaryIndexes[i].keySchema); i++ {

							keySchemav := make(map[string]interface{})


							if createtable.localSecondaryIndexes[i].keySchema[i].AttributeName != "" {
								keySchemav["AttributeName"] = createtable.localSecondaryIndexes[i].keySchema[i].AttributeName
							}

							if createtable.localSecondaryIndexes[i].keySchema[i].KeyType != "" {
								keySchemav["KeyType"] = createtable.localSecondaryIndexes[i].keySchema[i].KeyType
							}

							keySchemavs = append(keySchemavs, keySchemav)
						}

						localSecondaryIndexesvjsonmap["KeySchema"] = keySchemavs
					}

				 localSecondaryIndexesvarrayjsonmap = append(localSecondaryIndexesvarrayjsonmap,localSecondaryIndexesvjsonmap)
				}
					createtablejsonmap["LocalSecondaryIndexes"]=localSecondaryIndexesvarrayjsonmap
		}
}


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

func preparelisttables(params map[string]string, ExclusiveStartTableName string, Limit string, Region string) {
	if ExclusiveStartTableName != "" {
		params["ExclusiveStartTableName"] = ExclusiveStartTableName
	}

	if Limit != "" {
		params["Limit"] = Limit
	}

	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "DynamoDB_20120810.ListTables"
}

//PrepareSignatureV4query creates PrepareSignatureV4 for request.
func (dynamodb *Dynamodb) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}, response map[string]interface{}) error {
	ECSEndpoint := "https://dynamodb." + params["Region"] + ".amazonaws.com"
	fmt.Println("ECSEndpoint : ", ECSEndpoint)
	service := "dynamodb"
	method := "POST"
	host := service + "." + params["Region"] + ".amazonaws.com"
	fmt.Println("host : ", host)
	ContentType := "application/x-amz-json-1.0"
	signedheaders := "content-type;host;x-amz-date;x-amz-target"
	amztarget := params["amztarget"]

	requestparametersjson, _ := json.Marshal(paramsmap)
	requestparametersjsonstring := string(requestparametersjson)
	requestparametersjsonstringbyte := []byte(requestparametersjsonstring)
	client := new(http.Client)
	request, _ := http.NewRequest("POST", ECSEndpoint, bytes.NewBuffer(requestparametersjsonstringbyte))
	request = awsauth.SignatureV4(request, requestparametersjsonstringbyte, amztarget, method, params["Region"], service, host, ContentType, signedheaders)
	resp, err := client.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	response["body"] = string(body)
	response["status"] = resp.StatusCode
	return err
}
