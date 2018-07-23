# gocloud compute - AWS

## Configure AWS credentials.

Create `amazoncloudconfig.json` in your <b>HOME/.gocloud</b> directory as follows:
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

amazoncloud := gocloud.AmazonProvider()

```

### create Datasets

```js
createDatasets := map[string]interface{}{
  "Region":             "us-east-1",
  "ClusterIdentifier":  "examplecluster",
  "MasterUsername":     "masteruser",
  "MasterUserPassword": "12345678Aa",
  "NumberOfNodes":      "1",
  "NodeType":           "ds2.xlarge",
}

resp, err := amazoncloud.CreateDatasets(createDatasets)

if err != nil {
  t.Errorf("Test Fail")
}

response := resp.(map[string]interface{})

fmt.Println(response["body"])

```

### update Datasets

```js
	updateDatasets := map[string]interface{}{
		"Region":             "us-east-1",
		"ClusterIdentifier":  "examplecluster",
		"AutomatedSnapshotRetentionPeriod":     "2",
		"AllowVersionUpgrade": true,
		"ClusterParameterGroupName":      "parametergroup1",
		"PreferredMaintenanceWindow":           "wed:07:30-wed:08:00",
	}

	resp, err := amazoncloud.UpdateDatasets(updateDatasets)


	response := resp.(map[string]interface{})

	fmt.Println(response["body"])
```

### Delete Datasets

```js
deleteDatasets := map[string]interface{}{
  "Region":            "us-east-1",
  "ClusterIdentifier": "test",
}

resp, err := amazoncloud.DeleteDatasets(deleteDatasets)


response := resp.(map[string]interface{})

fmt.Println(response["body"])
```

### Reboot instance

```js
  Reboot := map[string]string{
		"instance-id": "i-037a9fae81c33ac30",
		"Region":      "us-east-1",
	}


  resp, err := amazoncloud.Rebootnode(Reboot)
  response := resp.(map[string]interface{})
  fmt.Println(response["body"])
```

### Get Datasets

```js
getDatasets := map[string]interface{}{
  "Region": "us-east-1",
}

resp, err := amazoncloud.GetDatasets(getDatasets)


response := resp.(map[string]interface{})

fmt.Println(response["body"])
```
