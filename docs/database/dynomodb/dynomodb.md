# gocloud database - dynamodb

## Configure AWS credentials.

Create `amazoncloudconfig.json` in your <b>HOME</b> directory as follows:
```js
{
  "AWSAccessKeyID": "xxxxxxxxxxxx",
  "AWSSecretKey": "xxxxxxxxxxxx"
}
```

You can also set the credentials as environment variables:
```js
export AWSAccessKeyID =  "xxxxxxxxxxxx"
export AWSSecretKey = "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/cloudlibz/gocloud/gocloud"

amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider)
```

### create tables

```js
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

resp, err := amazoncloud.Createtables(createtables)

 response := resp.(map[string]interface{})
 fmt.Println(response["body"])
```

### Delete tables

```js
deletetables := map[string]interface{}{
  "Region":    "us-east-2",
  "TableName": "hello",
}

  resp, err := amazoncloud.Deletetables(deletetables)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### List tables

```js
listtables := map[string]interface{}{
  "Region":    "us-east-2",
  "TableName": "hello",
}

  resp, err := amazoncloud.Listtables(listtables)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### describe tables

```js
  describetables := map[string]interface{}{
    "Region":    "us-east-2",
    "TableName": "hello",
  }

  resp, err := amazoncloud.Describetables(describetables)

  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```
