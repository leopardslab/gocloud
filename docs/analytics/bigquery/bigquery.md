# gocloud analytics - bigquery

## Configure Google Cloud credentials.

Download your service account credentials file from Google Cloud and save it as `googlecloudinfo.json` in your <b>HOME/.gocloud</b> directory.

You can also set the credentials as environment variables:
```js
export PrivateKey =  "xxxxxxxxxxxx"
export Type =  "xxxxxxxxxxxx"
export ProjectID = "xxxxxxxxxxxx"
export PrivateKeyID = "xxxxxxxxxxxx"
export ClientEmail = "xxxxxxxxxxxx"
export ClientID = "xxxxxxxxxxxx"
export AuthURI = "xxxxxxxxxxxx"
export TokenURI = "xxxxxxxxxxxx"
export AuthProviderX509CertURL = "xxxxxxxxxxxx"
export ClientX509CertURL =  "xxxxxxxxxxxx"
```

## Initialize library

```js

import "github.com/cloudlibz/gocloud/gocloud"

googlecloud := gocloud.GoogleProvider()
```

### List Datasets

```js

listDatasets := map[string]interface{}{
	"ProjectId": "gocloud-206919",
	"All" : true,
	"Filter" : "",
	"MaxResults" : 1,
	"PageToken" : "",
}

resp, err := googlecloud.ListDatasets(listDatasets)

response := resp.(map[string]interface{})
fmt.Println(response["body"])
  ```

### Delete Datasets

```js

deleteDatasets := map[string]string{
	"ProjectId": "gocloud-206919",
	"DatasetId" : "gocloudv2",
}

resp, err := googlecloud.DeleteDatasets(deleteDatasets)

response := resp.(map[string]interface{})

fmt.Println(response["body"])
```

### Get Datasets

```js

getDatasets := map[string]string{
	"ProjectId": "gocloud-206919",
	"DatasetId" : "gocloudv3",
}

resp, err := googlecloud.GetDatasets(getDatasets)

response := resp.(map[string]interface{})
fmt.Println(response["body"])

```

### Create Datasets

```js
datasetReference := map[string]string {
 "DatasetId": "gocloudv4",
 "ProjectId": "gocloud-206919",
}

createDatasets := map[string]interface{}{
	"ProjectId": "gocloud-206919",
	"Description" : "gocloudv4 created",
	"Kind": "bigquery#dataset",
	"Etag": "wJ6J76UJduYf9EzfNz0gJw==",
	"SelfLink": "https://www.googleapis.com/bigquery/v2/projects/gocloud-206919/datasets/gocloudv3",
	 "Id": "gocloud-206919:gocloudv4",
	"Location": "US",
	"DatasetReference" : datasetReference,
}

resp, err := googlecloud.CreateDatasets(createDatasets)


response := resp.(map[string]interface{})
fmt.Println(response["body"])

```


### Create update

```js
datasetReference := map[string]string {
 "DatasetId": "gocloudv3",
 "ProjectId": "gocloud-206919",
}

updateDatasets := map[string]interface{}{
	"ProjectId": "gocloud-206919",
	"Description" : "gocloudv3 created",
	"Kind": "bigquery#dataset",
	"Etag": "wJ6J76UJduYf9EzfNz0gJw==",
	"SelfLink": "https://www.googleapis.com/bigquery/v2/projects/gocloud-206919/datasets/gocloudv3",
	 "Id": "gocloud-206919:gocloudv3",
	"Location": "US",
	"DatasetReference" : datasetReference,
	 "DatasetId": "gocloudv3",
}

resp, err := googlecloud.UpdateDatasets(updateDatasets)


response := resp.(map[string]interface{})
fmt.Println(response["body"])


```
