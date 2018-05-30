package dynamodb

import "testing"
import awsAuth "github.com/cloudlibz/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}


func TestDescribetables(t *testing.T) {

	var dynamodb Dynamodb

	describetables := map[string]interface{}{
		"Region":    "us-east-2",
		"TableName": "hello",
	}

	_, err := bigtable.Describetables(describetables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListtables(t *testing.T) {

	var dynamodb Dynamodb

	listtables := map[string]interface{}{
		"Region":    "us-east-2",
		"TableName": "hello",
	}

	_, err := bigtable.Listtables(listtables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeletetables(t *testing.T) {

	var dynamodb Dynamodb

	deletetables := map[string]interface{}{
		"Region":    "us-east-2",
		"TableName": "hello",
	}

	_, err := dynamodb.Deletetables(deletetables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}




func TestCreatetables(t *testing.T) {

	var dynamodb Dynamodb


	keySchema :=[]map[string]interface{}{
		{
					 "AttributeName": "ForumName",
					 "KeyType": "HASH",
			 },
			 {
					 "AttributeName": "Subject",
					 "KeyType": "RANGE",
			 },
	}


	attributeDefinitions := []map[string]interface{}{

		{
            "AttributeName": "ForumName",
            "AttributeType": "S",
        },
        {
            "AttributeName": "Subject",
            "AttributeType": "S",
        },
        {
            "AttributeName": "LastPostDateTime",
            "AttributeType": "S",
        },

		}


		projection := map[string]interface{}{
					"ProjectionType": "KEYS_ONLY",
		}

		provisionedThroughput := map[string]interface{}{
        "ReadCapacityUnits": 5,
        "WriteCapacityUnits": 5,
    }

	localSecondaryIndexes :=[]map[string]interface{}{
		{
            "IndexName": "LastPostIndex",
            "KeySchema": keySchema ,
            "Projection": projection,
    },
	}



	createtables := map[string]interface{}{
		"Region":    "us-east-1",
		"TableName" : "Thread",
		"KeySchema" : keySchema,
		"AttributeDefinitions" :attributeDefinitions ,
		"LocalSecondaryIndexes" : localSecondaryIndexes,
		"ProvisionedThroughput" : provisionedThroughput,
	}

	_, err := dynamodb.Createtables(createtables)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
